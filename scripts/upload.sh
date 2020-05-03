

curl -X POST -H "Content-Type: application/openapi+json" --data-binary @openapi/test/json/uber.json http://localhost:8000/api/apis/uber:api
curl -X POST -H "Content-Type: application/openapi+json" --data-binary @openapi/test/json/petstore-simple.json http://localhost:8000/api/apis/petstore-simple:api
curl -X POST -H "Content-Type: application/openapi+json" --data-binary @openapi/test/json/petstore-expanded.json http://localhost:8000/api/apis/petstore-expanded:api
curl -X POST -H "Content-Type: application/openapi+json" --data-binary @openapi/test/json/petstore.json http://localhost:8000/api/apis/petstore:api
curl -X POST -H "Content-Type: application/openapi+json" --data-binary @openapi/test/json/api-with-examples.json http://localhost:8000/api/apis/api-with-examples:api
curl -X POST -H "Content-Type: application/openapi+yaml" --data-binary @openapi/test/v3.0/uspto.yaml http://localhost:8000/api/apis/uspto:api
curl -X POST -H "Content-Type: application/wadl+xml" --data-binary @wadl/test/storage-service.wadl http://localhost:8000/api/apis/storage:service:api
