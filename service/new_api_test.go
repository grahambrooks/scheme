package main

import (
	"github.com/grahambrooks/apellicon/search"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParsingByContentType(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		content := `{ "swagger": "2.0"}`
		model, err := parseContent("application/openapi+json", ioutil.NopCloser(strings.NewReader(content)))
		assert.NoError(t, err)
		assert.Equal(t, search.OpenAPI2, model.Kind)
	})

	t.Run("YAML", func(t *testing.T) {
		content := `openapi: "3.0"`
		model, err := parseContent("application/openapi+yaml", ioutil.NopCloser(strings.NewReader(content)))
		assert.NoError(t, err)
		assert.Equal(t, search.OpenAPI3, model.Kind)
	})

	t.Run("WADL", func(t *testing.T) {
		content := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<application xmlns="http://research.sun.com/wadl/2006/10">
    <doc xmlns:jersey="http://jersey.dev.java.net/"
         jersey:generatedBy="Jersey: 1.0-ea-SNAPSHOT 10/02/2008 12:17 PM"/>
    <resources base="http://localhost:9998/storage/">
        <resource path="/containers">
            <method name="GET" id="getContainers">
                <response>
                    <representation mediaType="application/xml"/>
                </response>
            </method>
            <resource path="{container}">
                <param xmlns:xs="http://www.w3.org/2001/XMLSchema"
                       type="xs:string" style="template" name="container"/>
                <method name="PUT" id="putContainer">
                    <response>
                        <representation mediaType="application/xml"/>
                    </response>
                </method>
                <method name="DELETE" id="deleteContainer"/>
                <method name="GET" id="getContainer">
                    <request>
                        <param xmlns:xs="http://www.w3.org/2001/XMLSchema"
                               type="xs:string" style="query" name="search"/>
                    </request>
                    <response>
                        <representation mediaType="application/xml"/>
                    </response>
                </method>
                <resource path="{item: .+}">
                    <param xmlns:xs="http://www.w3.org/2001/XMLSchema"
                           type="xs:string" style="template" name="item"/>
                    <method name="PUT" id="putItem">
                        <request>
                            <representation mediaType="*/*"/>
                        </request>
                        <response>
                            <representation mediaType="*/*"/>
                        </response>
                    </method>
                    <method name="DELETE" id="deleteItem"/>
                    <method name="GET" id="getItem">
                        <response>
                            <representation mediaType="*/*"/>
                        </response>
                    </method>
                </resource>
            </resource>
        </resource>
    </resources>
</application>`
		model, err := parseContent("application/wadl+xml", ioutil.NopCloser(strings.NewReader(content)))
		assert.NoError(t, err)
		assert.Equal(t, search.WADL, model.Kind)
	})

	t.Run("Missing content type", func(t *testing.T) {
		content := `rubbish`
		_, err := parseContent("", ioutil.NopCloser(strings.NewReader(content)))
		assert.EqualError(t, err, "missing Content-Type. Supported content types application/openapi+json, applicatiion/openapi+yaml or application/wadl+xml")
	})
	t.Run("JUNK", func(t *testing.T) {
		content := `rubbish`
		_, err := parseContent("junk", ioutil.NopCloser(strings.NewReader(content)))
		assert.EqualError(t, err, "content type 'junk' not supported. Supported content types application/openapi+json, applicatiion/openapi+yaml or application/wadl+xml")
	})
}
