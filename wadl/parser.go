package wadl

import (
	"encoding/xml"
	"github.com/grahambrooks/scheme/search"
	"io"
	"io/ioutil"
	"path"
)

type Parser struct {
}

type Param struct {
	Type  string `xml:"type,attr"`
	Style string `xml:"style,attr"`
	Name  string `xml:"name,attr"`
}

type Requests struct {
	Param          Param          `xml:"param"`
	Representation Representation `xml:"representation"`
}

type Representation struct {
	MediaType string `xml:"mediaType,attr"`
}

type Response struct {
	Representation Representation `xml:"representation"`
}

type Method struct {
	Name      string     `xml:"name,attr"`
	Id        string     `xml:"id,attr"`
	Requests  []Requests `xml:"request"`
	Responses []Response `xml:"response"`
}

type Resource struct {
	Path      string     `xml:"path,attr"`
	Resources []Resource `xml:"resource"`
	Param     []Param    `xml:"param"`
	Methods   []Method   `xml:"methods"`
}

type Resources struct {
	Base      string     `xml:"base,attr"`
	Resources []Resource `xml:"resource"`
}

type Spec struct {
	Resources Resources `xml:"resources"`
}

func (parser *Parser) Parse(reader io.Reader) (search.Model, error) {
	model := search.Model{}

	var spec Spec
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return model, err
	}
	err = xml.Unmarshal(bytes, &spec)
	if err != nil {
		return model, err
	}

	model.Kind = search.WADL
	model.Host = spec.Resources.Base

	parseModelResources(&model, spec.Resources.Resources, "")

	return model, nil
}

func parseModelResources(model *search.Model, resources []Resource, base string) {
	for _, resource := range resources {
		resourcePath := path.Join(base, resource.Path)
		r := search.Resource{Path: resourcePath}
		for _, param := range resource.Param {
			p := search.Parameter{
				Type:     param.Type,
				Location: param.Style,
				Name:     param.Name,
			}
			r.Parameters = append(r.Parameters, p)
		}
		model.Resources = append(model.Resources, r)
		parseModelResources(model, resource.Resources, resourcePath)
	}
}
