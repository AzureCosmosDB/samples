import { CosmosClient } from "@azure/cosmos";
import { DefaultAzureCredential } from "@azure/identity";

const endpoint = process.env["COSMOS_ENDPOINT"];
if (!endpoint) {
  throw new Error("COSMOS_ENDPOINT environment variable is not set.");
}

const credential = new DefaultAzureCredential();
const client = new CosmosClient({ endpoint, aadCredentials: credential });

const { resource: account } = await client.getDatabaseAccount();
console.log(`Connected. Consistency: ${account?.consistencyPolicy}`);
