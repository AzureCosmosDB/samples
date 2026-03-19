# Connect to Azure Cosmos DB for NoSQL — Passwordless (.NET)

This sample demonstrates how to authenticate to Azure Cosmos DB for NoSQL using **passwordless authentication** via `DefaultAzureCredential` from the Azure Identity library. No connection strings or keys are required.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [.NET 10 SDK](https://dotnet.microsoft.com/download/dotnet/10.0)
- [Azure CLI](https://learn.microsoft.com/cli/azure/install-azure-cli) (for local development login)

## Setup

### 1. Assign the required role

Grant your Azure AD identity the **Cosmos DB Built-in Data Reader** role on your Cosmos DB account:

```bash
az cosmosdb sql role assignment create \
  --account-name <cosmos-account-name> \
  --resource-group <resource-group> \
  --role-definition-name "Cosmos DB Built-in Data Reader" \
  --principal-id $(az ad signed-in-user show --query id -o tsv) \
  --scope "/"
```

## Run

### 1. Set your endpoint using user secrets

```shell
dotnet user-secrets set --file connect.cs "COSMOS_ENDPOINT" "https://<your-account>.documents.azure.com:443/"
```

### 2. Log in with Azure CLI

```shell
az login
```

### 3. Run the sample

```shell
dotnet run connect.cs
```

`DefaultAzureCredential` automatically uses your Azure CLI login locally. In production (Azure App Service, Azure Functions, Azure Container Apps, etc.) it uses the managed identity assigned to the resource — no code changes needed.

## What this sample does

1. Reads `COSMOS_ENDPOINT` from user secrets (via `Microsoft.Extensions.Configuration.UserSecrets`)
2. Authenticates using `DefaultAzureCredential` (Azure CLI / Managed Identity / etc.)
3. Reads account metadata and prints the account name
