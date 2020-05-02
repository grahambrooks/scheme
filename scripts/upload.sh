

curl -X POST -H "Content-Type: application/openapi+json" -d @openapi/test/json/uber.json http://localhost:8000/api/interfaces/uber:api
curl -X POST -H "Content-Type: application/openapi+json" -d @openapi/test/json/petstore-simple.json http://localhost:8000/api/interfaces/petstore-simple:api
curl -X POST -H "Content-Type: application/openapi+json" -d @openapi/test/json/petstore-expanded.json http://localhost:8000/api/interfaces/petstore-expanded:api
curl -X POST -H "Content-Type: application/openapi+json" -d @openapi/test/json/petstore.json http://localhost:8000/api/interfaces/petstore:api
curl -X POST -H "Content-Type: application/openapi+json" -d @openapi/test/json/api-with-examples.json http://localhost:8000/api/interfaces/api-with-examples:api
curl -X POST -H "Content-Type: application/openapi+yaml" --data-binary @openapi/test/v3.0/uspto.yaml http://localhost:8000/api/interfaces/uspto:api
curl -X POST -H "Content-Type: application/wadl+xml" -d @wadl/test/storage-service.wadl http://localhost:8000/api/interfaces/storage:service:api
