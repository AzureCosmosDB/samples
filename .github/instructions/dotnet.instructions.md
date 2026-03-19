---
applyTo: "dotnet/**"
---

# .NET sample conventions

## File-based apps

- Use `.cs` script files with `#:package` directives — do **not** create `.csproj` files
- Run samples with `dotnet run {filename}.cs`
- Use `#:property` for per-file MSBuild properties when needed

## Package management

- **`Directory.Packages.props`** (central): only packages used by **2+ samples** (e.g., `Microsoft.Azure.Cosmos`, `Newtonsoft.Json`). Use floating major versions (`Version="3.*"`).
- **`#:package` in the `.cs` file**: packages used by **only one sample** (e.g., `Azure.Identity` is only in `001-connect-passwordless`). Use `#:package Foo@3.*` syntax.
- When a second sample needs a previously single-use package, promote it to `Directory.Packages.props`.

## Build properties

- `Directory.Build.props` holds shared properties like `UserSecretsId`
- `.editorconfig` enforces code style — do not remove or override it

## Authentication

- **00x connect samples**: `DefaultAzureCredential` via `Azure.Identity`. Store endpoint in user secrets (`dotnet user-secrets set --file {file}.cs "COSMOS_ENDPOINT" "..."`)
- **All other samples**: key-based via `COSMOS_ENDPOINT` + `COSMOS_KEY` user secrets

## Code style

- Top-level statements (no `Main` method, no namespace)
- Use `record` types for the `Product` data model
- Use `var` where the type is obvious
- Suffix async calls with `Async` — use `await` directly in top-level code
