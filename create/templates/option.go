package templates

import (
	"fmt"
	"text/template"
)

var optionTemplate = template.Must(
	template.New("option").Parse(optionTemplateText),
)

func CreateOption(projectName, fullpath string) {
	path := fmt.Sprintf("%s/%s/server/api/option.go", fullpath, projectName)
	writeFile(optionTemplate, path, nil)
}

const optionTemplateText = `package api

type Link string

type ParamType string

type EndpointInfo struct {
	About string `+"`json:\"about\"`"+`
	Link  Link   `+"`json:\"url\"`"+`
}

type Endpoints map[string]EndpointInfo

type Option struct {
	EndpointInfo
	List   OptionMethod `+"`json:\"list,omitempty\"`"+`
	Get    OptionMethod `+"`json:\"get,omitempty\"`"+`
	Post   OptionMethod `+"`json:\"post,omitempty\"`"+`
	Put    OptionMethod `+"`json:\"put,omitempty\"`"+`
	Patch  OptionMethod `+"`json:\"patch,omitempty\"`"+`
	Delete OptionMethod `+"`json:\"delete,omitempty\"`"+`
}

func NewOption(link Link, about string) Option {
	return Option{
		EndpointInfo: EndpointInfo{About: about, Link: link},
		List:         methodUnsupported(link),
		Get:          methodUnsupported(link),
		Post:         methodUnsupported(link),
		Put:          methodUnsupported(link),
		Patch:        methodUnsupported(link),
		Delete:       methodUnsupported(link),
	}
}

func (o *Option) SetList(description string) {
	o.List.Supported = true
	o.List.Description = description
	o.List.ExampleURL = o.EndpointInfo.Link
}

func (o *Option) SetGet(exampleParam, description string) {
	o.Get.Supported = true
	o.Get.Description = description
	o.Get.ExampleURL = o.EndpointInfo.Link + "/" + Link(exampleParam)
}

func (o *Option) SetPost(description string) {
	o.Post.Supported = true
	o.Post.Description = description
	o.Post.ExampleURL = o.EndpointInfo.Link
}

func (o *Option) SetPut(description string) {
	o.Put.Supported = true
	o.Put.Description = description
	o.Put.ExampleURL = o.EndpointInfo.Link
}

func (o *Option) SetPatch(description string) {
	o.Put.Supported = true
	o.Put.Description = description
	o.Put.ExampleURL = o.EndpointInfo.Link
}

func (o *Option) SetDelete(exampleParam, description string) {
	o.Get.Supported = true
	o.Get.Description = description
	o.Get.ExampleURL = o.EndpointInfo.Link + "/" + Link(exampleParam)
}

func methodUnsupported(link Link) OptionMethod {
	return OptionMethod{
		Supported:  false,
		Parameters: make(map[string]OptionParam, 0),
	}
}
`
