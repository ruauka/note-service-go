package swagger

import (
	"github.com/julienschmidt/httprouter"

	"web/docs"
)

func Register(router *httprouter.Router) {
	// Get spec for swagger-ui
	router.GET("/openapi.yaml", docs.OpenapiHandler)
	// Static file for swagger-ui
	router.ServeFiles("/swagger/*filepath", docs.NewSwaggerHandler())
}
