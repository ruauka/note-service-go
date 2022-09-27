// Package swagger Package swagger
package swagger

import (
	"github.com/julienschmidt/httprouter"

	"web/api"
)

// Register register swagger handlers.
func Register(router *httprouter.Router) {
	// Get spec for swagger-ui
	router.GET("/openapi.yaml", api.OpenapiHandler)
	// Static file for swagger-ui
	router.ServeFiles("/swagger/*filepath", api.NewSwaggerHandler())
}
