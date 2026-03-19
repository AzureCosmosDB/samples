# Change Feed with Azure Cosmos DB for NoSQL (TypeScript)

This sample demonstrates how to read the **change feed** of an Azure Cosmos DB for NoSQL container using the Azure SDK for JavaScript/TypeScript. It upserts a set of items and then reads all changes from the beginning of the container's history using the change feed iterator.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [Node.js 20+](https://nodejs.org/)

## Run

> [!IMPORTANT]
> These environment variables are set **only for the current terminal session** and do **not persist** after you close the window.

### macOS / Linux

```bash
export COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
export COSMOS_KEY="<your-primary-key>"
npm install
npm start
```

### Windows (PowerShell)

```powershell
$Env:COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
$Env:COSMOS_KEY="<your-primary-key>"
npm install
npm start
```

You can find your endpoint and primary key in the Azure portal under your Cosmos DB account → **Keys**.

## What this sample does

1. Connects to Azure Cosmos DB using key-based authentication
2. Creates (or retrieves) the `cosmicworks` database
3. Creates (or retrieves) the `products` container partitioned by `/category`
4. Upserts two `Product` items to generate change feed events
5. Reads all changes from the **beginning** of the container's history using `getChangeFeedIterator`
6. Prints each changed item until the feed is caught up (`304 Not Modified`)
