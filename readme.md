# Notes

## References

- https://github.com/golang/go/issues/50750
- https://github.com/golang/go/issues/73654
- https://github.com/golang/go/issues/54264
- https://go.dev/doc/tutorial/workspaces
- https://go.dev/blog/get-familiar-with-workspaces
- https://go.dev/ref/mod#workspaces
- https://go.dev/doc/modules/managing-dependencies#local_directory
- https://github.com/googleapis/google-cloud-go

## Example

```bash
go work init
go work use cmd/product-api
go work use internal/product
go work use pkg/httpadapter

cd cmd/product-api

go mod edit -replace github.com/ldez/go-monorepo-example/httpadapter=../../pkg/httpadapter
go mod edit -replace github.com/ldez/go-monorepo-example/product=../../internal/product

go mod tidy

cd -

go work sync

```
