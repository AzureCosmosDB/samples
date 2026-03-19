#:package Azure.Identity

using Azure.Identity;
using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Configuration;

IConfigurationRoot configuration = new ConfigurationBuilder()
    .AddUserSecrets<Program>()
    .Build();

string endpoint = configuration["COSMOS_ENDPOINT"]
    ?? throw new InvalidOperationException("COSMOS_ENDPOINT is not configured. Run: dotnet user-secrets set --file connect.cs \"COSMOS_ENDPOINT\" \"<your-endpoint>\"");

var credential = new DefaultAzureCredential();
using var client = new CosmosClient(endpoint, credential);

AccountProperties account = await client.ReadAccountAsync();
Console.WriteLine($"Account: {account.Id}");
