

curl -X POST -H "Content-Type: application/json" -d @openapi-examples/v2.0/json/uber.json http://localhost:8000/api/interfaces
curl -X POST -H "Content-Type: application/json" -d @openapi-examples/v2.0/json/petstore-simple.json http://localhost:8000/api/interfaces
curl -X POST -H "Content-Type: application/json" -d @openapi-examples/v2.0/json/petstore-expanded.json http://localhost:8000/api/interfaces
curl -X POST -H "Content-Type: application/json" -d @openapi-examples/v2.0/json/petstore.json http://localhost:8000/api/interfaces
curl -X POST -H "Content-Type: application/json" -d @openapi-examples/v2.0/json/api-with-examples.json http://localhost:8000/api/interfaces
