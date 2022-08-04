package docs

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	// AppName - application name.
	AppName string
	// AppVersion - version application.
	AppVersion string
)

// BuildInfo holds build information about the app.
var BuildInfo = buildInfo{
	AppName,
	AppVersion,
}

type buildInfo struct {
	AppName    string `json:"app_name,omitempty"`
	AppVersion string `json:"app_version,omitempty"`
}

//go:embed swagger-ui
var swaggerFS embed.FS

//go:embed openapi
var apidocs embed.FS

// OpenapiHandler get dynamic spec for http-server.
func OpenapiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Copy source template
	tmpl, _ := template.ParseFS(apidocs, "openapi/openapi.yaml")
	tmpl.Execute(w, BuildInfo) //nolint:errcheck
}

// NewSwaggerHandler returns Handler for endpoint `/swagger/*`.
func NewSwaggerHandler() http.FileSystem {
	fswagger, _ := fs.Sub(swaggerFS, "swagger-ui")
	return http.FS(fswagger)
}
