# Connect to Azure Cosmos DB for NoSQL — Key Auth (TypeScript)

This sample demonstrates how to authenticate to Azure Cosmos DB for NoSQL using an **account key**. This is a quick way to get started, especially for local development and testing.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [Node.js 18+](https://nodejs.org/)

## Run

> [!IMPORTANT]
> These environment variables are set **only for the current terminal session** and do **not persist** after you close the window.

### macOS / Linux

```bash
export COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
export COSMOS_KEY="<your-primary-key>"
npm install && npm start
```

### Windows (PowerShell)

```powershell
$Env:COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
$Env:COSMOS_KEY="<your-primary-key>"
npm install; npm start
```

You can find your endpoint and primary key in the Azure portal under your Cosmos DB account → **Keys**.

## What this sample does

1. Reads `COSMOS_ENDPOINT` and `COSMOS_KEY` from the environment
2. Creates a `CosmosClient` using the endpoint and key
3. Reads account metadata and prints the account name
