# Connect to Azure Cosmos DB for NoSQL — Key Auth (.NET)

This sample demonstrates how to authenticate to Azure Cosmos DB for NoSQL using an **account key**. This is a quick way to get started, especially for local development and testing.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [.NET 10 SDK](https://dotnet.microsoft.com/download/dotnet/10.0)

## Run

### 1. Set your endpoint and key using user secrets

```shell
dotnet user-secrets set --file connect.cs "COSMOS_ENDPOINT" "https://<your-account>.documents.azure.com:443/"
dotnet user-secrets set --file connect.cs "COSMOS_KEY" "<your-primary-key>"
```

You can find your endpoint and primary key in the Azure portal under your Cosmos DB account → **Keys**.

### 2. Run the sample

```shell
dotnet run connect.cs
```

## What this sample does

1. Reads `COSMOS_ENDPOINT` and `COSMOS_KEY` from user secrets (via `Microsoft.Extensions.Configuration.UserSecrets`)
2. Creates a `CosmosClient` using the endpoint and key
3. Reads account metadata and prints the account name
