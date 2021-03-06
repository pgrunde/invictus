package templates

import (
	"fmt"
	"text/template"
)

var apiTemplate = template.Must(
	template.New("api").Parse(apiTemplateText),
)

func CreateAPI(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/api.go", fullpath, projectName)
	writeFile(apiTemplate, path, nil)
}

const apiTemplateText = `package api

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/aodin/sol"
	"github.com/aodin/volta/config"
)

// API holds routing information for all attached Rest endpoints
type API struct {
	prefix    string
	resources map[string]Rest
	endpoints Endpoints
	routes    *node
	conn      sol.Conn
	config    config.Config
}

// Prefix returns the prefix of the API
func (api *API) Prefix() string {
	return api.prefix
}

// SetPrefix sets the prefix of the API
func (api *API) SetPrefix(prefix string) *API {
	if prefix == "" {
		prefix = "/"
	} else if prefix[0] != '/' {
		prefix = "/" + prefix
	}
	api.prefix = prefix
	return api
}

func (api API) Link(resource string, values ...url.Values) Link {
	u := api.config.URL()
	u.Path = filepath.Join(api.prefix, resource)
	if len(values) > 0 {
		u.RawQuery = values[0].Encode()
	}
	return Link(u.String())
}

func (api API) RootURL() string {
	u := api.config.URL()
	u.Path = api.prefix
	return u.String()
}

func (api *API) Add(resource Rest, params ...string) error {
	name := resource.Name()
	if _, exists := api.resources[name]; exists {
		return fmt.Errorf(
			"api: a resource named '%s' already exists",
			name,
		)
	}
	api.resources[name] = resource
	resource.SetAPI(api)

	api.endpoints[name] = resource.Options().EndpointInfo

	p := api.prefix
	api.routes.addRoute(fmt.Sprintf("%s%s", p, name), resource)

	var keys []string
	for _, param := range params {
		keys = append(keys, fmt.Sprintf(":%s", param))
		pk := strings.Join(keys, "/")
		api.routes.addRoute(fmt.Sprintf("%s%s/%s", p, name, pk), resource)
	}

	return nil
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == api.prefix {
		response := Response{
			Results: api.endpoints,
		}
		Write(w, response)
		return
	}

	resource, params, _ := api.routes.getValue(r.URL.Path)
	if resource == nil {
		response := Response{
			Results: make([]interface{}, 0),
		}
		Write(w, response)
		return
	}

	request := NewRequest(r, params...)

	var response Response
	var err *Error

	if r.Method == "OPTIONS" {
		response := Empty()
		response.Results = []Option{resource.Options()}
		Write(w, response)
		return
	}

	if len(params) == 0 {
		switch r.Method {
		case "GET":
			response, err = resource.List(request)
		case "POST":
			response, err = resource.Post(request)
		default:
			response, err = Unsupported(request)
		}
	} else {
		switch r.Method {
		case "GET":
			response, err = resource.Get(request)
		case "PATCH":
			response, err = resource.Patch(request)
		case "DELETE":
			response, err = resource.Delete(request)
		default:
			response, err = Unsupported(request)
		}
	}

	// response.Meta.Errors = err
	response.Error = err
	Write(w, response)
	return
}

func New(conf config.Config, conn sol.Conn) *API {
	return &API{
		conn:      conn,
		config:    conf,
		prefix:    "/",
		resources: make(map[string]Rest),
		endpoints: Endpoints{},
		routes:    &node{},
	}
}
`
