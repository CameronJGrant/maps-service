docker run -d --net=host --name data-service -e POSTGRES_USER=game-data -e POSTGRES_PASSWORD=game-data postgres > /dev/null

echo "export PG_HOST=127.0.0.1" > /tmp/svcg.env
echo "export PG_PORT=5432" >> /tmp/svcg.env
echo "export PG_USER=game-data" >> /tmp/svcg.env
echo "export PG_PASS=game-data" >> /tmp/svcg.env
echo "export PG_DB=game-data" >> /tmp/svcg.env

echo "source /tmp/svcg.env"

read -s -n 1 -p "Press any key to exit . . ."
echo ""

docker rm -f data-service > /dev/null


