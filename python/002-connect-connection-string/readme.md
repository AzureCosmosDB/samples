# Connect to Azure Cosmos DB for NoSQL — Connection String (Python)

This sample demonstrates how to authenticate to Azure Cosmos DB for NoSQL using a **connection string** (account endpoint + key). This is a quick way to get started, especially for local development and testing.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [Python 3.8+](https://www.python.org/downloads/)

## Setup

### 1. Copy the example env file

```bash
cp .env.example .env
```

### 2. Fill in your connection string

Open `.env` and replace the placeholder with your actual connection string:

```
COSMOS_CONNECTION_STRING=AccountEndpoint=https://<your-account>.documents.azure.com:443/;AccountKey=<your-key>==;
```

You can find the connection string in the Azure portal under your Cosmos DB account → **Keys** → **PRIMARY CONNECTION STRING**.

> **Security note:** Never commit your `.env` file to source control. It is already listed in `.gitignore` by convention.

## Run

```bash
pip install -r requirements.txt && python app.py
```

## What this sample does

1. Loads environment variables from a `.env` file (if present) using `python-dotenv`
2. Reads `COSMOS_CONNECTION_STRING` from the environment
3. Creates a `CosmosClient` using the connection string
4. Creates (or retrieves) a database named `cosmicworks`
5. Creates (or retrieves) a container named `products` with partition key `/category`
6. Creates a product item
7. Queries items filtered by category and prints the results
