package openapi

import "C"
import (
	"context"
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"net/http"
)


func ValidateRequests(path string) gin.HandlerFunc {
	spec := openapi3filter.NewRouter().WithSwaggerFromFile(path)


	errorEncoder := &openapi3filter.ValidationErrorEncoder{
		Encoder: errorEncoder,
	}

	return func(c *gin.Context) {
		httpReq := c.Request
		route, pathParams, err := spec.FindRoute(httpReq.Method, httpReq.URL)

		if err != nil {
			errorEncoder.Encode(c, err, c.Writer)
			c.Abort()
			return
		}

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    httpReq,
			PathParams: pathParams,
			Route:      route,
		}

		if err := openapi3filter.ValidateRequest(c.Request.Context(), requestValidationInput); err != nil {
			errorEncoder.Encode(c, err, c.Writer)
			c.Abort()
			return
		}
		c.Next()
	}
}

func statusCode(err error) int {
	code := http.StatusInternalServerError
	if sc, ok := err.(openapi3filter.StatusCoder); ok {
		code = sc.StatusCode()
	}
	return code
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	if headerer, ok := err.(openapi3filter.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := statusCode(err)

	w.WriteHeader(code)

	if vErr, ok := err.(*openapi3filter.ValidationError); ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8\"")
		json.NewEncoder(w).Encode(vErr)
	}  else {
		body := []byte(err.Error())
		json.NewEncoder(w).Encode(&openapi3filter.ValidationError{
			Status: 500,
			Title:string(body),
		})
	}
}
