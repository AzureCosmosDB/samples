import { CosmosClient } from "@azure/cosmos";

const endpoint = process.env["COSMOS_ENDPOINT"];
if (!endpoint) {
  throw new Error("COSMOS_ENDPOINT environment variable is not set.");
}

const key = process.env["COSMOS_KEY"];
if (!key) {
  throw new Error("COSMOS_KEY environment variable is not set.");
}

const client = new CosmosClient({ endpoint, key });

const { resource: account } = await client.getDatabaseAccount();
console.log(`Connected. Consistency: ${account?.consistencyPolicy}`);
