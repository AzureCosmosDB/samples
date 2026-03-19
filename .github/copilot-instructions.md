# GitHub Copilot instructions — Azure Cosmos DB samples

These instructions describe the conventions and standards for this repository. Follow them when creating or modifying samples.

## Repository structure

Samples are grouped by programming language at the root level. Each sample is a self-contained directory inside its language folder.

```text
/
├── dotnet/
├── python/
├── javascript/
├── java/
└── go/
```

## Numbering standard

Each sample directory name is prefixed with a three-digit number:

| Range | Category | Description |
| --- | --- | --- |
| `000–099` | **Connectivity** | Basic connectivity, authentication, and client configuration |
| `100–199` | **Quickstart** | Hello-world style scenarios to get started quickly |
| `200–299` | **Feature usage** | Using specific Azure Cosmos DB features such as change feed or bulk operations |
| `300–399` | **Patterns** | Common application patterns and best practices |
| `400–499` | **Advanced** | Advanced scenarios including AI and vector search integrations |
| `500–599` | **Vertical integration** | End-to-end vertical integration scenarios |

Directory names must follow the pattern `NNN-kebab-case-name` (e.g. `201-change-feed`).

## Every sample must include

- A `readme.md` at the root of the sample directory.
- All source files needed to run the sample without modification beyond setting environment variables or secrets.

## Validation

The [validation workflow](.github/workflows/validate.yml) runs on every pull request and push to `main`. It checks:

1. **Sample structure** — every directory under a language root must match `NNN-kebab-case-name` and contain a `readme.md`.
2. **Language correctness** — only the language subtrees that contain changed files are built and tested. Unmodified languages are skipped.

## C# (.NET) conventions

### Script-style samples

C# samples use the .NET 10 [script-file](https://learn.microsoft.com/dotnet/csharp/fundamentals/program-structure/top-level-statements) format (`dotnet run <file>.cs`). Package references use the `#:package` directive at the top of the file.

### Package management rule

- **Shared packages** (used in more than one sample) — declare only the version in `dotnet/Directory.Packages.props`. The `#:package` directive in the script omits the version and resolves it centrally.
- **Single-use packages** (used in exactly one sample) — embed the version directly in the `#:package` directive. Do **not** add these to `Directory.Packages.props`.

**Example — `Azure.Identity` is only used in `001-connect-passwordless`:**

```csharp
// connect.cs — version embedded because this is the only sample that uses Azure.Identity
#:package Azure.Identity@1.*
#:package Microsoft.Azure.Cosmos     // version comes from Directory.Packages.props
```

```xml
<!-- Directory.Packages.props — only packages used across multiple samples -->
<ItemGroup>
  <PackageVersion Include="Microsoft.Azure.Cosmos" Version="3.*" />
  <PackageVersion Include="Newtonsoft.Json" Version="13.*" />
</ItemGroup>
```

### `Directory.Build.props`

Use `dotnet/Directory.Build.props` only for MSBuild properties that apply to every sample (e.g. `UserSecretsId`). Do not add package references here.
