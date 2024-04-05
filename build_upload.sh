registry_name="acrasgardeomainrnd001"

# Ask for the tag input
read -p "Enter tag for the Docker image: " tag

# Build Docker image
docker build -t secret-sidecar:$tag .

# Login to Azure Container Registry (ACR)
az acr login --name $registry_name

# Tag the Docker image
docker tag secret-sidecar:$tag $registry_name.azurecr.io/secret-sidecar:$tag

# Push Docker image to ACR
docker push $registry_name.azurecr.io/secret-sidecar:$tag