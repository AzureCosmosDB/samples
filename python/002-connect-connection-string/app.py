import os

from dotenv import load_dotenv
from azure.cosmos import CosmosClient, PartitionKey

load_dotenv()

connection_string = os.environ.get("COSMOS_CONNECTION_STRING")
if not connection_string:
    raise ValueError("COSMOS_CONNECTION_STRING environment variable is not set.")

client = CosmosClient.from_connection_string(connection_string)

database = client.create_database_if_not_exists("cosmicworks")
print(f"Created/retrieved database: {database.id}")

container = database.create_container_if_not_exists(
    id="products",
    partition_key=PartitionKey(path="/category"),
)
print(f"Created/retrieved container: {container.id}")

new_item = {
    "id": "70b63682-b93a-4c77-aad2-65501347265f",
    "category": "gear-surf-surfboards",
    "name": "Yamba Surfboard",
    "quantity": 12,
    "sale": False,
}

created = container.upsert_item(body=new_item)
print(f"Upserted item: {created['id']}")

query = "SELECT * FROM products p WHERE p.category = @category"
parameters = [{"name": "@category", "value": "gear-surf-surfboards"}]

items = list(container.query_items(query=query, parameters=parameters, enable_cross_partition_query=False))
for item in items:
    print(f"Read item: {item['name']} | Quantity: {item['quantity']} | Sale: {item['sale']}")
