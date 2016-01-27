package templates

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var apiTemplate = template.Must(
	template.New("api").Parse(apiTemplateText),
)

func CreateAPI(projectName, fullpath string) {
	file, err := os.Create(fmt.Sprintf("%s/%s/server/api/api.go", fullpath, projectName))
	if err != nil {
		log.Fatalf("cannot create a server/api/api.go file: %s", err)
	}
	defer file.Close()
	apiTemplate.Execute(file, nil)
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
	conn      sol.Connection
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
	api.routes.addRoute(fmt.Sprintf("%s%s/", p, name), resource)

	var keys []string
	for _, param := range params {
		keys = append(keys, fmt.Sprintf(":%s", param))
		pk := strings.Join(keys, "/")
		api.routes.addRoute(fmt.Sprintf("%s%s/%s", p, name, pk), resource)
		api.routes.addRoute(fmt.Sprintf("%s%s/%s/", p, name, pk), resource)
	}

	return nil
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoding := JSON{}

	if r.URL.Path == api.prefix {
		response := Response{
			Meta:    Meta{Limit: 250},
			Links:   Links{"self": Link(api.RootURL())},
			Results: api.endpoints,
		}
		encoding.Write(w, response)
		return
	}

	resource, params, _ := api.routes.getValue(r.URL.Path)
	if resource == nil {
		response := Response{
			Meta: Meta{
				Limit:  1,
				Errors: MetaError(http.StatusNotFound, "Resource not found"),
			},
			Links:   Links{},
			Results: make([]interface{}, 0),
		}
		encoding.Write(w, response)
		return
	}

	request := NewRequest(r, params...)

	var response Response
	var err *Error

	method := Method(r.Method)

	if method == OPTIONS {
		response := Empty()
		response.Results = []Option{resource.Options()}
		encoding.Write(w, response)
		return
	}

	if len(params) == 0 {
		switch method {
		case GET:
			response, err = resource.List(request)
		case POST:
			response, err = resource.Post(request)
		default:
			response, err = Unsupported(request)
		}
	} else {
		switch method {
		case GET:
			response, err = resource.Get(request)
		case PATCH:
			response, err = resource.Patch(request)
		case DELETE:
			response, err = resource.Delete(request)
		default:
			response, err = Unsupported(request)
		}
	}

	response.Meta.Errors = err
	encoding.Write(w, response)
	return
}

func New(conf config.Config, conn sol.Connection) *API {
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
