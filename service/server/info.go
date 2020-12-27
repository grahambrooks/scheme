package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PathInfo struct {
	PathTemplate  string
	PathRegexp    string
	Queries       []string
	QueriesRegexp []string
	Methods       []string
}

func (s *SchemeServer) InfoHandler(writer http.ResponseWriter, _ *http.Request) {
	var pathInfo []*PathInfo

	_ = s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pi := PathInfo{}
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			pi.PathTemplate = pathTemplate
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			pi.PathRegexp = pathRegexp
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			pi.Queries = queriesTemplates
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			pi.QueriesRegexp = queriesRegexps
		}
		methods, err := route.GetMethods()
		if err == nil {
			pi.Methods = methods
		}
		pathInfo = append(pathInfo, &pi)
		return nil
	})

	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(pathInfo)

}
