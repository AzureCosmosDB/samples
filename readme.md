# Azure Cosmos DB samples

One stop repo for all Cosmos DB samples.

## Overview

This repo contains code samples for Azure Cosmos DB for NoSQL organized by programming language. Each sample is self-contained and demonstrates a specific concept or operation using the Azure Cosmos DB SDK for each language.

## Repository structure

Samples are organized by programming language at the root level. Each project is self-contained.

```text
/
├── dotnet/
│   ├── 001-connect-passwordless/
│   ├── 101-create-item/
│   └── 102-query-items/
├── python/
│   ├── 002-connect-connection-string/
│   ├── 101-create-item/
│   └── 102-query-items/
└── go/
    ├── 002-connect-connection-string/
    ├── 101-create-item/
    └── 201-change-feed/
```

## Numbering standard

Each sample folder is prefixed with a three-digit number that indicates its category:

| Range | Category | Description |
| --- | --- | --- |
| `000–099` | **Connectivity** | Basic connectivity, authentication, and client configuration |
| `100–199` | **Quickstart** | Hello-world style scenarios to get started quickly |
| `200–299` | **Feature usage** | Using specific Azure Cosmos DB features such as change feed or bulk operations |
| `300–399` | **Patterns** | Common application patterns and best practices |
| `400–499` | **Advanced** | Advanced scenarios including AI and vector search integrations |
| `500–599` | **Vertical integration** | End-to-end vertical integration scenarios |


## Validation

Every pull request and push to `main` runs the [validation workflow](.github/workflows/validate.yml), which enforces two things:

- **Sample structure** — all sample directories must follow the `NNN-kebab-case-name` naming convention (e.g. `001-quickstart`) and include a `README.md`.
- **Language correctness** — changes under a language root folder (`python/`, `javascript/`, `java/`, `dotnet/`, `go/`) trigger the corresponding build and test job for that language only. Unmodified languages are skipped.
One stop repo for all Azure Cosmos DB samples.
