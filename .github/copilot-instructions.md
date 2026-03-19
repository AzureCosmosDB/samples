# Azure Cosmos DB Samples — Copilot Instructions

## Product naming

- Always say **Azure Cosmos DB** — never "Azure Cosmos DB for NoSQL"
- The NoSQL API is the only API in this repository and is always inferred
- Do not generate samples for MongoDB, Cassandra, Gremlin, or Table APIs

## Repository structure

- Root folders by language: `dotnet/`, `python/`, `go/`, `typescript/`
- Sample folders use the pattern `NNN-kebab-case-name/` (regex: `^[0-9]{3}-[a-z0-9]+(-[a-z0-9]+)*$`)
- Every sample folder **must** contain a `readme.md`
- Each sample is self-contained and independently runnable

## Numbering standard

Sample numbers are **shared across languages** — the same number means the same concept in every language.

| Range | Category | Description |
|-------|----------|-------------|
| `000–099` | Connectivity | Client setup, authentication, emulator |
| `100–199` | Quickstart | Hello-world style getting-started scenarios |
| `200–299` | Features | Change feed, bulk, batch, patch, indexing |
| `300–399` | Patterns | Multi-tenancy, pagination, retry, hierarchical partition keys |
| `400–499` | Advanced | Vector search, AI integration, semantic cache |
| `500–599` | Vertical | End-to-end apps (todo, IoT, e-commerce) |

## Shared data model

All samples use the same database, container, and item shape:

- **Database:** `cosmicworks`
- **Container:** `products`
- **Partition key:** `/category`
- **Product shape:** `{ id, category, name, quantity, sale }`
- **Canonical item:** id `70b63682-b93a-4c77-aad2-65501347265f`, category `gear-surf-surfboards`, name `Yamba Surfboard`, quantity `12`, sale `false`

## Authentication

- **00x range only** (connect samples): passwordless via `DefaultAzureCredential`
- **All other samples**: key-based via `COSMOS_ENDPOINT` + `COSMOS_KEY` environment variables

## Package versions

Do **not** pin specific package versions. Use floating or major-version ranges:

- .NET: `Version="3.*"` in `Directory.Packages.props` or `#:package Foo@3.*` in script files
- Python: `azure-cosmos>=4.9` in `requirements.txt`
- Go: standard `go get` versioning
- TypeScript: caret ranges (`^`) in `package.json`

## README conventions

- **Title pattern:** `{Operation} with Azure Cosmos DB ({Language})`
- **Sections:** Prerequisites, Run (with macOS/Linux + Windows PowerShell blocks), What this sample does
- **Environment variable callout:** use `> [!IMPORTANT]` note
- **Connect samples (00x):** only read account metadata — no CRUD operations
- Do not include `.env.example` files — put env var instructions inline in the readme

## Testing and CI

- Samples are validated on **every push to `main`** and on **every pull request**
- Samples should be runnable against the **Azure Cosmos DB emulator**
- The CI workflow (`.github/workflows/validate.yml`) runs per-language build jobs independently
- Structure validation enforces the naming regex and readme presence

## Anti-patterns

- Do not say "for NoSQL" — just "Azure Cosmos DB"
- Do not add `.env.example` files
- Do not use `python-dotenv` — use `os.environ` directly
- Do not pin exact package versions
- Do not mix CRUD operations into connect samples
- Do not use `matrix` strategy in CI with single versions (causes version-suffixed check names)
