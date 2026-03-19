---
applyTo: ".github/workflows/**"
---

# CI workflow conventions

## Workflow structure

- Single workflow file: `validate.yml` — do not split into per-language workflows
- Trigger on both `push` to `main` AND `pull_request` to `main`
- Set `permissions: contents: read` at the workflow level
- Use `concurrency` group with `cancel-in-progress: true`

## Structure validation job

- Scan all language folders: `python`, `javascript`, `typescript`, `java`, `dotnet`, `go`
- Enforce naming regex: `^[0-9]{3}-[a-z0-9]+(-[a-z0-9]+)*$`
- Check for `readme.md` using **case-insensitive** find: `find "${sample}" -maxdepth 1 -type f -iname "readme.md"`
- Do **not** use exact filename match (e.g., `[ -f "README.md" ]`) — this breaks on case variations

## Per-language build jobs

- Each language job runs independently and only scans its own folder
- Exit cleanly (`exit 0`) if the language folder doesn't exist
- Do **not** use `matrix` strategy with single versions — hardcode versions in `with:` to avoid version-suffixed check names
- Use latest major action versions (`actions/checkout@v6`, `actions/setup-node@v6`, etc.)

## Language-specific rules

- **.NET**: handle both `.csproj`/`.sln` projects AND file-based `.cs` apps. Do **not** use `-q` quiet flag on `dotnet build` (it suppresses errors)
- **Python**: `pip install -r requirements.txt` then `py_compile` or `pytest`
- **Go**: `go build ./...` and `go test ./...`
- **JavaScript/TypeScript**: scan both `javascript/` and `typescript/` folders. `npm install` + `npm test --if-present`

## Emulator testing

- Samples should be testable against the Azure Cosmos DB emulator
- Do not require a live Azure account for CI validation
