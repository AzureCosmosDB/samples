# Connect to Azure Cosmos DB for NoSQL — Passwordless (Python)

This sample demonstrates how to authenticate to Azure Cosmos DB for NoSQL using **passwordless authentication** via `DefaultAzureCredential` from the Azure Identity library. No connection strings or keys are required.

## Prerequisites

- An [Azure subscription](https://azure.microsoft.com/free/)
- An Azure Cosmos DB for NoSQL account
- [Python 3.8+](https://www.python.org/downloads/)
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

> [!IMPORTANT]
> This environment variable is set **only for the current terminal session** and does **not persist** after you close the window.

### macOS / Linux

```bash
export COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
az login
pip install -r requirements.txt && python app.py
```

### Windows (PowerShell)

```powershell
$Env:COSMOS_ENDPOINT="https://<your-account>.documents.azure.com:443/"
az login
pip install -r requirements.txt; python app.py
```

`DefaultAzureCredential` automatically uses your Azure CLI login locally. In production (Azure App Service, Azure Functions, Azure Container Apps, etc.) it uses the managed identity assigned to the resource — no code changes needed.

## What this sample does

1. Reads `COSMOS_ENDPOINT` from the environment
2. Authenticates using `DefaultAzureCredential` (Azure CLI / Managed Identity / etc.)
3. Reads account metadata and prints the account name
