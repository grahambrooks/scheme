package search

import (
	"encoding/json"
	"io"
)

type InterfaceKind string

const (
	OpenAPI2 InterfaceKind = "OpenAPI 2.0"
	OpenAPI3 InterfaceKind = "OpenAPI 3.0"
	WADL     InterfaceKind = "WADL"
)

type Parameter struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
}

type Request struct {
}

type Response struct{}

type Method struct {
	Name      string     `json:"name,omitempty"`
	Id        string     `json:"id,omitempty"`
	Requests  []Request  `json:"requests,omitempty"`
	Responses []Response `json:"responses,omitempty"`
}

type Resource struct {
	Path       string      `json:"path,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Methods    []Method    `json:"methods,omitempty"`
}

type Model struct {
	Kind        InterfaceKind `json:"kind"`
	Host        string        `json:"host,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Version     string        `json:"version,omitempty"`
	Schemes     []string      `json:"schemes,omitempty"`
	Resources   []Resource    `json:"resources,omitempty"`
}

func (m Model) AsJson(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(m)
}
