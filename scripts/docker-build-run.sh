
echo "building container image"
docker-compose build
docker-compose start
open http://localhost:8000
