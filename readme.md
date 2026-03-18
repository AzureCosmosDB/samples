# Azure Cosmos DB samples

One stop repo for all Azure Cosmos DB samples.

## Overview

This repo contains code samples for Azure Cosmos DB for NoSQL organized by programming language. Each sample is self-contained and demonstrates a specific concept or operation using the Azure Cosmos DB SDK for each language.

## Repository structure

Samples are organized by programming language at the root level. Each project is self-contained.

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
| `000–099` | **Connectivity** | Basic connectivity, authentication, and client configuration |
| `100–199` | **Hello world** | Hello-world style scenarios to get started quickly |
| `200–299` | **Feature usage** | Using specific Azure Cosmos DB features such as change feed or bulk operations |
| `300–399` | **Patterns** | Common application patterns and best practices |
| `400–499` | **Advanced** | Advanced scenarios including AI and vector search integrations |
| `500–599` | **Vertical integration** | End-to-end vertical integration scenarios |
