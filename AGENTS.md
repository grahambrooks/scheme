# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Scheme is an API registry and search engine written in Go. It provides a web-based search UI for API discovery and an API for uploading/managing API specifications in OpenAPI 2/3 (JSON/YAML) and WADL (XML) formats. Elasticsearch is used as the backing store for full-text search.

## Build Commands

```bash
# Run all tests
bazel test //...

# Build and run the service (default port 8000)
bazel run //service:scheme

# Update Go dependencies (updates go.mod/go.sum)
bazel run @rules_go//go -- get -u ./...
bazel run @rules_go//go -- mod tidy
```

## Running Locally

Start Elasticsearch before running the service:
```bash
sh ./scripts/es-start.sh
```

Or use Docker Compose for the full stack:
```bash
docker build --tag scheme .
docker-compose up
```

Upload example API specs: `sh scripts/upload.sh`

Configure Elasticsearch URL via `ELASTICSEARCH_URL` environment variable (defaults to localhost).

## Architecture

**Core packages:**
- `service/` - Main application entry point and HTTP server
- `service/server/` - HTTP handlers using gorilla/mux, routes prefixed with `/api/`
- `service/store/` - Data persistence layer with `ApiStore` interface (Elasticsearch implementation)
- `search/` - Domain model (`Model`, `Resource`, `Method`, `Parameter`)
- `openapi/` - Parser for OpenAPI 2.0/3.0 specs (JSON/YAML)
- `wadl/` - Parser for WADL XML specs
- `rules/` - API validation rules engine

**Key patterns:**
- Dependency injection: `SchemeServer` accepts `ApiStore` interface, enabling `StubApiStore` for testing
- Both parsers convert specs to a common `search.Model` domain object
- HTTP handlers are methods on `SchemeServer` receiver
- Static web content served from `service/site/`

**API routes:** `/api/` prefix with endpoints for `/search`, `/apis`, `/apis/{id}`, `/registrations`, `/stats`, `/info`

## Testing

Tests use testify/assert for assertions and httptest for HTTP handler testing. Test fixtures are in `*/test/` and `*/testdata/` directories within each package.
