# Cosmos DB samples

One stop repo for all Cosmos DB samples.

## Validation

Every pull request and push to `main` runs the [validation workflow](.github/workflows/validate.yml), which enforces two things:

- **Sample structure** — all sample directories must follow the `NNN-kebab-case-name` naming convention (e.g. `001-quickstart`) and include a `README.md`.
- **Language correctness** — changes under a language root folder (`python/`, `javascript/`, `java/`, `dotnet/`, `go/`) trigger the corresponding build and test job for that language only. Unmodified languages are skipped.
