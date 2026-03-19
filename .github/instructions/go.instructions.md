---
applyTo: "go/**"
---

# Go sample conventions

## File structure

- Single `main.go` as the entry point
- `go.mod` with module path `github.com/AzureCosmosDB/samples/go/{sample-name}`
- Use standard `go get` versioning — do not pin exact versions

## Authentication

- Key-based: `COSMOS_ENDPOINT` + `COSMOS_KEY` environment variables
- Use `azcosmos.NewKeyCredential(key)` and `azcosmos.NewClientWithKey(endpoint, cred, nil)`
- Validate env vars at startup with `log.Fatalf(...)` if missing

## Code style

- Use a `Product` struct with JSON tags (`json:"id"`, etc.)
- Use `json.Marshal` / `json.Unmarshal` for item serialization
- Handle 409 Conflict errors with an `isConflict()` helper function:
  ```go
  func isConflict(err error) bool {
      var responseErr *azcore.ResponseError
      return errors.As(err, &responseErr) && responseErr.StatusCode == http.StatusConflict
  }
  ```
- Use `context.Background()` for all Cosmos DB operations
- Use `log.Fatalf` for fatal errors, `fmt.Printf` for output
