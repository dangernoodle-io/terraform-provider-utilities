# terraform-provider-utilities

Terraform provider exposing utility functions. Built on Terraform Plugin Framework (Protocol 6).

## Build & Test

```shell
make          # fmt, lint, install, generate
make test     # unit tests
make testacc  # acceptance tests (requires terraform)
```

## Adding a New Function

1. Create `internal/provider/function_<name>.go` implementing `function.Function`
2. Register in `provider.go` `Functions()`
3. Add example in `examples/functions/<name>/function.tf`
4. Add tests in `internal/provider/function_<name>_test.go`
5. `make generate` to create docs

`docs/` is generated — do not edit manually.

## Linting

`golangci-lint run ./...` — config in `.golangci.yml`.
Do not leave unused unexported functions or variables (including test helpers).
