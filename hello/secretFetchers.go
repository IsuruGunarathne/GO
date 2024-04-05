package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func fetchAzureSecret() (string, error) {
	// Azure Key Vault configuration
	// vaultName := os.Getenv("AZURE_VAULT_NAME")
	// secretName := os.Getenv("AZURE_SECRET_NAME")

	vaultName := "secretsidecar"
	secretName := "mysecret"

	vaultURI := fmt.Sprintf("https://%s.vault.azure.net", vaultName)

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)

	if err != nil {
		log.Fatalf("failed to get the Client: %v", err)
		return "", err
	}

	version := ""
	resp, err := client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
		return "", err
	}

	fmt.Printf("secretValue: %s\n", *resp.Value)

	return *resp.Value, nil

}
