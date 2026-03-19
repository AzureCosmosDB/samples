---
applyTo: "python/**"
---

# Python sample conventions

## File structure

- Single `app.py` as the entry point
- `requirements.txt` for dependencies — use floating versions (e.g., `azure-cosmos>=4.9`)
- Do **not** use `python-dotenv` — read config with `os.environ` directly

## Authentication

- Key-based: `COSMOS_ENDPOINT` + `COSMOS_KEY` environment variables
- Validate env vars at startup with `raise ValueError(...)` if missing

## Code style

- Use `dict` literals for item data (no dataclasses or Pydantic models)
- Use `PartitionKey` from `azure.cosmos` for container creation
- Use parameterized queries with `@parameter` syntax for SQL queries
- Print operations and results to stdout for verification
