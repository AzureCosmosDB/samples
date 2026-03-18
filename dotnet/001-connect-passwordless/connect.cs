#:package Azure.Identity@1.13.2
#:package Microsoft.Azure.Cosmos@3.*
#:package Microsoft.Extensions.Configuration.UserSecrets@10.*
#:property UserSecretsId=azurecosmosdb-samples-dotnet-001-connect-passwordless

using Azure.Identity;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Configuration;

IConfigurationRoot configuration = new ConfigurationBuilder()
    .AddUserSecrets<Program>()
    .Build();

string endpoint = configuration["COSMOS_ENDPOINT"]
    ?? throw new InvalidOperationException("COSMOS_ENDPOINT is not configured. Run: dotnet user-secrets set --file connect.cs \"COSMOS_ENDPOINT\" \"<your-endpoint>\"");

var credential = new DefaultAzureCredential();
var client = new CosmosClient(endpoint, credential);

using FeedIterator<DatabaseProperties> iterator = client.GetDatabaseQueryIterator<DatabaseProperties>();
while (iterator.HasMoreResults)
{
    FeedResponse<DatabaseProperties> databases = await iterator.ReadNextAsync();
    foreach (DatabaseProperties db in databases)
    {
        Console.WriteLine($"Database: {db.Id}");
    }
}
