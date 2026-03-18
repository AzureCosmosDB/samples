#:package Azure.Identity
#:package Microsoft.Azure.Cosmos
#:package Microsoft.Extensions.Configuration.UserSecrets

using Azure.Identity;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Configuration;

IConfigurationRoot configuration = new ConfigurationBuilder()
    .AddUserSecrets<Program>()
    .Build();

string endpoint = configuration["COSMOS_ENDPOINT"]
    ?? throw new InvalidOperationException("COSMOS_ENDPOINT is not configured. Run: dotnet user-secrets set --file Program.cs \"COSMOS_ENDPOINT\" \"<your-endpoint>\"");

var credential = new DefaultAzureCredential();
var client = new CosmosClient(endpoint, credential);

Database database = await client.CreateDatabaseIfNotExistsAsync("cosmicworks");
Console.WriteLine($"Created/retrieved database: {database.Id}");

Container container = await database.CreateContainerIfNotExistsAsync("products", "/category");
Console.WriteLine($"Created/retrieved container: {container.Id}");

Product newItem = new(
    id: "70b63682-b93a-4c77-aad2-65501347265f",
    category: "gear-surf-surfboards",
    name: "Yamba Surfboard",
    quantity: 12,
    sale: false
);

ItemResponse<Product> createResponse = await container.UpsertItemAsync(newItem, new PartitionKey(newItem.category));
Console.WriteLine($"Upserted item: {createResponse.Resource.id} (status: {createResponse.StatusCode})");

ItemResponse<Product> readResponse = await container.ReadItemAsync<Product>(
    newItem.id,
    new PartitionKey(newItem.category)
);
Console.WriteLine($"Read item: {readResponse.Resource.name} | Quantity: {readResponse.Resource.quantity}");

record Product(string id, string category, string name, int quantity, bool sale);
