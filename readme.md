docker build -t goimage:1.0 .

# Run the container

docker run -d --rm goimage:1.0

# Check the logs

docker logs <container_id>

# Stop the container

docker stop <container_id>

# docker exec

docker exec -it <container_id> /bin/sh
