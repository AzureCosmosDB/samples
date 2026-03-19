#:package Microsoft.Azure.Cosmos
#:package Microsoft.Extensions.Configuration.UserSecrets
#:package Newtonsoft.Json

using Microsoft.Azure.Cosmos;
using Microsoft.Extensions.Configuration;

IConfigurationRoot configuration = new ConfigurationBuilder()
    .AddUserSecrets<Program>()
    .Build();

string endpoint = configuration["COSMOS_ENDPOINT"]
    ?? throw new InvalidOperationException("COSMOS_ENDPOINT is not configured. Run: dotnet user-secrets set --file connect.cs \"COSMOS_ENDPOINT\" \"<your-endpoint>\"");

string key = configuration["COSMOS_KEY"]
    ?? throw new InvalidOperationException("COSMOS_KEY is not configured. Run: dotnet user-secrets set --file connect.cs \"COSMOS_KEY\" \"<your-key>\"");

var client = new CosmosClient(endpoint, key);

AccountProperties account = await client.ReadAccountAsync();
Console.WriteLine($"Account: {account.Id}");
