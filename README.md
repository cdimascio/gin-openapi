## gin-openapi

Automatically validates API requests against an OpenAPI 3 spec. 

## Usage
```go
package main

import (
	openapi "github.com/cdimascio/gin-openapi/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	
    // Add gin-openapi middleware 
    router.Use(openapi.ValidateRequests("spec.yml"))

    // Add routes
    router.POST("/v1/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{ "id": "pong" })
    })

    router.Run(":8080")
}
```

## License MIT
