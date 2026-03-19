import os

from azure.cosmos import CosmosClient
from azure.identity import DefaultAzureCredential

endpoint = os.environ.get("COSMOS_ENDPOINT")
if not endpoint:
    raise ValueError("COSMOS_ENDPOINT environment variable is not set.")

credential = DefaultAzureCredential()
client = CosmosClient(url=endpoint, credential=credential)

account = client.get_database_account()
print(f"Account: {account['id']}")
