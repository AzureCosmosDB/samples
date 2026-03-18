# CRUD Operations with Azure Cosmos DB for NoSQL (Go)

This sample demonstrates the four fundamental **CRUD** operations (Create, Read, Update, Delete) against Azure Cosmos DB for NoSQL using the Azure SDK for Go.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [Go 1.21+](https://go.dev/dl/)

## Run

> [!IMPORTANT]
> These environment variables are set **only for the current terminal session** and do **not persist** after you close the window.

### macOS / Linux

```bash
export COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
export COSMOS_KEY="<your-primary-key>"
go run main.go
```

### Windows (PowerShell)

```powershell
$Env:COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
$Env:COSMOS_KEY="<your-primary-key>"
go run main.go
```

You can find your endpoint and primary key in the Azure portal under your Cosmos DB account → **Keys**.

## Operations demonstrated

| Operation | Description |
|-----------|-------------|
| **Create** | Inserts a new `Product` item into the `products` container |
| **Read**   | Retrieves the item by ID and partition key, prints its name and quantity |
| **Update** | Modifies the `quantity` field to `20` and replaces the item |
| **Delete** | Removes the item from the container |

The sample uses key-based authentication (`COSMOS_KEY`) and creates the `cosmicworks` database and `products` container (partitioned by `/category`) if they do not already exist.
