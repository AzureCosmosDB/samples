import os

from azure.cosmos import CosmosClient

endpoint = os.environ.get("COSMOS_ENDPOINT")
if not endpoint:
    raise ValueError("COSMOS_ENDPOINT environment variable is not set.")

key = os.environ.get("COSMOS_KEY")
if not key:
    raise ValueError("COSMOS_KEY environment variable is not set.")

client = CosmosClient(url=endpoint, credential=key)

databases = list(client.list_databases())
for database in databases:
    print(f"Database: {database['id']}")
