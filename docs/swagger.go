package docs

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
)

var (
	// AppName - application name.
	AppName = "Notes service"
	// AppVersion - version application.
	AppVersion = "1.0"
	// GoVersion - golang version.
	GoVersion = runtime.Version()
)

// BuildInfo holds build information about the app.
var BuildInfo = buildInfo{
	AppName,
	AppVersion,
	GoVersion,
}

type buildInfo struct {
	AppName    string `json:"app_name,omitempty"`
	AppVersion string `json:"app_version,omitempty"`
	Language   string `json:"go_version,omitempty"`
}

// Print information about the app to stdout.
func (b *buildInfo) Print() {
	i, _ := json.MarshalIndent(b, "", "   ")
	fmt.Printf("Info:\n%v\n\n", string(i))
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
