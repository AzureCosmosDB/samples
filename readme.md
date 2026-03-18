# Azure Cosmos DB samples

One stop repo for all Azure Cosmos DB samples.

## Overview

This monorepo contains code samples for Azure Cosmos DB for NoSQL organized by programming language. Each sample is self-contained and demonstrates a specific concept or operation using the Azure Cosmos DB SDK.

## Repository structure

Samples are organized by programming language at the root level. There is no shared `/nosql/` subfolder — each language folder is self-contained.

```text
/
├── dotnet/
│   ├── 001-connect-passwordless/
│   ├── 101-create-item/
│   └── 201-query-items/
├── python/
│   ├── 002-connect-connection-string/
│   ├── 101-create-item/
│   └── 201-query-items/
└── go/
    ├── 002-connect-connection-string/
    ├── 101-create-item/
    └── 201-query-items/
```

## Numbering standard

Each sample folder is prefixed with a three-digit number that indicates its category:

| Range | Category | Description |
| --- | --- | --- |
| `000–099` | **Getting started** | Connecting to Azure Cosmos DB, authentication, and client configuration |
| `100–199` | **Item operations** | Creating, reading, updating, and deleting items (CRUD) |
| `200–299` | **Query operations** | Querying items using the NoSQL query language, including filters and pagination |
| `300–399` | **Index management** | Configuring indexing policies and composite indexes |
| `400–499` | **Advanced features** | Transactions, bulk operations, and change feed |

## Project plan

The table below lists the planned samples for the v1 release. Each sample targets 2–3 examples per category per language.

| # | Name | Category | .NET | Python | Go |
| --- | --- | --- | --- | --- | --- |
| `001` | Connect with passwordless auth | Getting started | ☑️ | | |
| `002` | Connect with connection string | Getting started | | ☑️ | ☑️ |
| `101` | Create an item | Item operations | ☑️ | ☑️ | ☑️ |
| `102` | Read and update an item | Item operations | ☑️ | ☑️ | ☑️ |
| `201` | Query items with a filter | Query operations | ☑️ | ☑️ | ☑️ |
| `202` | Query items with pagination | Query operations | ☑️ | ☑️ | |

> **Legend:** ☑️ = planned for v1

## Validation

Each language folder includes a GitHub Actions workflow that builds and validates all samples in that language. Workflows are scoped to path filters so a change in `dotnet/` only triggers the .NET workflow, a change in `python/` only triggers the Python workflow, and so on.

```text
.github/
└── workflows/
    ├── validate-dotnet.yml   # triggered on changes to dotnet/**
    ├── validate-python.yml   # triggered on changes to python/**
    └── validate-go.yml       # triggered on changes to go/**
```

A GitHub repository ruleset is recommended that requires the relevant language workflow to pass before a pull request can be merged. Because each workflow is scoped to its own path, only the workflow for the affected language needs to pass — keeping CI fast and avoiding unnecessary cross-language failures.
