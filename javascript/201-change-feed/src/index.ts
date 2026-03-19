import { CosmosClient, ChangeFeedStartFrom } from "@azure/cosmos";

interface Product {
  id: string;
  category: string;
  name: string;
  quantity: number;
  sale: boolean;
}

async function main(): Promise<void> {
  const endpoint = process.env.COSMOS_ENDPOINT;
  if (!endpoint) {
    throw new Error("COSMOS_ENDPOINT environment variable is not set.");
  }

  const key = process.env.COSMOS_KEY;
  if (!key) {
    throw new Error("COSMOS_KEY environment variable is not set.");
  }

  const client = new CosmosClient({ endpoint, key });

  // Create or retrieve database
  const { database } = await client.databases.createIfNotExists({ id: "cosmicworks" });
  console.log(`Database:\t${database.id}`);

  // Create or retrieve container
  const { container } = await database.containers.createIfNotExists({
    id: "products",
    partitionKey: { paths: ["/category"] },
  });
  console.log(`Container:\t${container.id}`);

  // Upsert items to generate change feed events
  const products: Product[] = [
    {
      id: "70b63682-b93a-4c77-aad2-65501347265f",
      category: "gear-surf-surfboards",
      name: "Yamba Surfboard",
      quantity: 12,
      sale: false,
    },
    {
      id: "25a68543-b90c-439d-8332-7ef41e06a0e0",
      category: "gear-surf-surfboards",
      name: "Kiama Classic Surfboard",
      quantity: 25,
      sale: true,
    },
  ];

  for (const product of products) {
    await container.items.upsert(product);
    console.log(`Upserted item:\t${product.id}`);
  }

  // Read the change feed from the beginning of the container's history
  console.log("\nReading change feed...");

  const iterator = container.items.getChangeFeedIterator<Product>({
    changeFeedStartFrom: ChangeFeedStartFrom.Beginning(),
  });

  while (iterator.hasMoreResults) {
    const { result, statusCode } = await iterator.readNext();

    // 304 Not Modified means the feed is caught up — no new changes
    if (statusCode === 304) {
      break;
    }

    for (const item of result) {
      console.log(`Changed item:\t${item.id} | ${item.name} | Qty: ${item.quantity}`);
    }
  }

  console.log("\nChange feed complete.");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
