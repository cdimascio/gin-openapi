package main

import (
	// "fmt"
	// "io/ioutil"
	// "github.com/getkin/kin-openapi/openapi3filter"
	openapi "github.com/cdimascio/go-oas/pkg"
	"github.com/gin-gonic/gin"
)

// func RequestValidation(spec *openapi3filter.Router) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		httpReq := c.Request
// 		route, pathParams, err := spec.FindRoute(httpReq.Method, httpReq.URL)

// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"message": err.Error(),
// 			})
// 		}
// 		// Validate request
// 		requestValidationInput := &openapi3filter.RequestValidationInput{
// 			Request:    httpReq,
// 			PathParams: pathParams,
// 			Route:      route,
// 		}

// 		x, _ := ioutil.ReadAll(httpReq.Body)

// 		fmt.Printf("%s \n", string(x))

// 		if err := openapi3filter.ValidateRequest(c.Request.Context(), requestValidationInput); err != nil {
// 			c.JSON(400, gin.H{
// 				"message": err.Error(),
// 			})
// 			c.Abort()
// 		}
// 		c.Next()
// 	}
// }

func main() {
	router := gin.Default()

	router.Use(openapi.ValidateRequests("spec.yaml"))

	router.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": "pong",
		})
	})

	router.POST("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": "pong",
		})
	})

	router.Run(":8080")
}
