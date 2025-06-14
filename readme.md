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
          
go work edit -replace github.com/ldez/go-monorepo-example/httpadapter@v0.0.0-00010101000000-000000000000=./pkg/httpadapter
go work edit -replace github.com/ldez/go-monorepo-example/product@v0.0.0-00010101000000-000000000000=./internal/product

go work sync

```

```console
$ cd ./cmd/product-api/
$ go run  .                    

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.13.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:3000
^Csignal: interrupt
```

```bash
$ go list -m -f '{{.Dir}}' | xargs -L 1 bash -c 'cd "$0" && pwd && golangci-lint run'
/home/ldez/sources/experimental/guidiego-monorepo-example/cmd/product-api
cmd/product-api/app/api.go:25:21: Error return value of `api.deps.Echo.Start` is not checked (errcheck)
        api.deps.Echo.Start(
                           ^
cmd/product-api/main.go:6:29: Magic number: 3000, in <argument> detected (mnd)
        app.New().Start("0.0.0.0", 3000)
                                   ^
cmd/product-api/app/api.go:1:1: package-comments: should have a package comment (revive)
package app
^
cmd/product-api/app/api.go:11:6: exported: exported type Api should have comment or be unexported (revive)
type Api struct {
     ^
cmd/product-api/app/api.go:15:1: exported: exported function New should have comment or be unexported (revive)
func New() *Api {
^
cmd/product-api/app/api.go:24:1: exported: exported method Api.Start should have comment or be unexported (revive)
func (api *Api) Start(address string, port int) {
^
cmd/product-api/app/di.go:11:2: exported: exported type Controllers should have comment or be unexported (revive)
        Controllers struct {
        ^
cmd/product-api/app/di.go:16:2: exported: exported type Dependencies should have comment or be unexported (revive)
        Dependencies struct {
        ^
cmd/product-api/app/di.go:22:1: exported: exported function InitDependencies should have comment or be unexported (revive)
func InitDependencies() *Dependencies {
^
cmd/product-api/main.go:1:1: package-comments: should have a package comment (revive)
package main
^
cmd/product-api/app/api.go:18:65: QF1008: could remove embedded field "Controllers" from selector (staticcheck)
        api.deps.Echo.GET("/product", httpadapter.EchoAdapter(api.deps.Controllers.ListProduct))
                                                                       ^
cmd/product-api/app/api.go:19:66: QF1008: could remove embedded field "Controllers" from selector (staticcheck)
        api.deps.Echo.POST("/product", httpadapter.EchoAdapter(api.deps.Controllers.SaveProduct))
                                                                        ^
cmd/product-api/app/api.go:9:5: var dependencies is unused (unused)
var dependencies *Dependencies
    ^
13 issues:
* errcheck: 1
* mnd: 1
* revive: 8
* staticcheck: 2
* unused: 1
/home/ldez/sources/experimental/guidiego-monorepo-example/internal/product
internal/product/application/controllers/ListProductController.go:22:9: Error return value of `w.Write` is not checked (errcheck)
        w.Write(response)
               ^
internal/product/application/controllers/SaveProductController.go:20:9: Error return value of `w.Write` is not checked (errcheck)
        w.Write([]byte("{ \"ok\": true }"))
               ^
internal/product/application/controllers/SaveProductController.go:15:47: dtos.Product is missing fields Id, Price, Name, Description (exhaustruct)
        err := controller.SaveProductUseCase.Execute(dtos.Product{})
                                                     ^
internal/product/application/mappers/money.go:10:13: valueobjects.Currency is missing fields Name, Token, TokenPosition, DecimalToken, ThousandToken (exhaustruct)
                Currency: valueobjects.Currency{},
                          ^
internal/product/domain/valueobjects/Money.go:30:9: Magic number: 3, in <condition> detected (mnd)
        if n > 3 {
               ^
internal/product/domain/valueobjects/Money.go:59:28: Magic number: 100, in <operation> detected (mnd)
        integerPart := absValue / 100
                                  ^
internal/product/domain/valueobjects/Money.go:60:31: Magic number: 100, in <operation> detected (mnd)
        fractionalPart := absValue % 100
                                     ^
internal/product/application/controllers/ListProductController.go:23:2: return with no blank line before (nlreturn)
        return nil
        ^
internal/product/application/controllers/SaveProductController.go:21:2: return with no blank line before (nlreturn)
        return nil
        ^
internal/product/application/usecases/SaveProductUseCase.go:28:2: return with no blank line before (nlreturn)
        return nil
        ^
internal/product/domain/valueobjects/Money.go:28:12: integer-format: fmt.Sprintf can be replaced with faster strconv.Itoa (perfsprint)
        intStr := fmt.Sprintf("%d", integerPart)
                  ^
internal/product/adapters/repositories/inmemory/ProductRepository.go:5:6: exported: exported type InMemoryProductRepository should have comment or be unexported (revive)
type InMemoryProductRepository struct {
     ^
internal/product/adapters/repositories/inmemory/ProductRepository.go:9:1: exported: exported method InMemoryProductRepository.List should have comment or be unexported (revive)
func (repo *InMemoryProductRepository) List() []*entities.Product {
^
internal/product/adapters/repositories/inmemory/ProductRepository.go:18:1: exported: exported method InMemoryProductRepository.Save should have comment or be unexported (revive)
func (repo *InMemoryProductRepository) Save(product *entities.Product) {
^
internal/product/application/controllers/ListProductController.go:1:1: package-comments: should have a package comment (revive)
package controllers
^
internal/product/application/controllers/ListProductController.go:10:6: exported: exported type ListProductController should have comment or be unexported (revive)
type ListProductController struct {
     ^
internal/product/application/controllers/ListProductController.go:14:72: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (controller *ListProductController) Handle(w http.ResponseWriter, r *http.Request) error {
                                                                       ^
internal/product/application/controllers/SaveProductController.go:10:6: exported: exported type SaveProductController should have comment or be unexported (revive)
type SaveProductController struct {
     ^
internal/product/application/controllers/SaveProductController.go:14:72: unused-parameter: parameter 'r' seems to be unused, consider removing or renaming it as _ (revive)
func (controller *SaveProductController) Handle(w http.ResponseWriter, r *http.Request) error {
                                                                       ^
internal/product/application/dtos/Product.go:3:6: exported: exported type Product should have comment or be unexported (revive)
type Product struct {
     ^
internal/product/application/dtos/Product.go:4:2: var-naming: struct field Id should be ID (revive)
        Id          string `json:"id"`
        ^
internal/product/application/mappers/money.go:7:24: unused-parameter: parameter 'priceString' seems to be unused, consider removing or renaming it as _ (revive)
func FromStringToMoney(priceString string) *valueobjects.Money {
                       ^
internal/product/application/mappers/product.go:8:1: exported: exported function FromProductDomainToDTO should have comment or be unexported (revive)
func FromProductDomainToDTO(product *entities.Product) *dtos.Product {
^
internal/product/application/usecases/ListProductUseCase.go:1:1: package-comments: should have a package comment (revive)
package usecases
^
internal/product/application/usecases/ListProductUseCase.go:9:6: exported: exported type ListProductUseCase should have comment or be unexported (revive)
type ListProductUseCase struct {
     ^
internal/product/application/usecases/ListProductUseCase.go:13:1: exported: exported method ListProductUseCase.Execute should have comment or be unexported (revive)
func (usecase *ListProductUseCase) Execute() []*dtos.Product {
^
internal/product/application/usecases/SaveProductUseCase.go:10:6: exported: exported type SaveProductUseCase should have comment or be unexported (revive)
type SaveProductUseCase struct {
     ^
internal/product/application/usecases/SaveProductUseCase.go:14:1: exported: exported method SaveProductUseCase.Execute should have comment or be unexported (revive)
func (usecase *SaveProductUseCase) Execute(productDto dtos.Product) error {
^
internal/product/domain/entities/Product.go:11:6: exported: exported type Product should have comment or be unexported (revive)
type Product struct {
     ^
internal/product/domain/entities/Product.go:12:2: var-naming: struct field Id should be ID (revive)
        Id          string
        ^
internal/product/domain/entities/Product.go:19:9: errorf: should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...) (revive)
        return errors.New(
                fmt.Sprintf("Cannot accept %s as empty string", field),
        )
internal/product/domain/entities/Product.go:24:1: exported: exported method Product.Validate should have comment or be unexported (revive)
func (product *Product) Validate() error {
^
internal/product/domain/repositories/IProductRepository.go:1:1: package-comments: should have a package comment (revive)
package repositories
^
internal/product/domain/repositories/IProductRepository.go:5:6: exported: exported type IProductRepository should have comment or be unexported (revive)
type IProductRepository interface {
     ^
internal/product/domain/valueobjects/Money.go:6:2: exported: exported type CurrencyTokenPosition should have comment or be unexported (revive)
        CurrencyTokenPosition int
        ^
internal/product/domain/valueobjects/Money.go:8:2: exported: exported type Currency should have comment or be unexported (revive)
        Currency struct {
        ^
internal/product/domain/valueobjects/Money.go:16:2: exported: exported type Money should have comment or be unexported (revive)
        Money struct {
        ^
internal/product/domain/valueobjects/Money.go:23:2: exported: exported const CurrencyTokenOnStart should have comment (or a comment on this block) or be unexported (revive)
        CurrencyTokenOnStart CurrencyTokenPosition = iota
        ^
internal/product/domain/valueobjects/Money.go:27:1: exported: exported method Currency.FormatString should have comment or be unexported (revive)
func (currency *Currency) FormatString(integerPart int, decimalPart int) string {
^
internal/product/domain/valueobjects/Money.go:57:1: exported: exported method Money.FormatString should have comment or be unexported (revive)
func (money *Money) FormatString() string {
^
internal/product/domain/valueobjects/Money.go:65:1: exported: exported method Money.Validate should have comment or be unexported (revive)
func (money *Money) Validate() error {
^
internal/product/application/controllers/ListProductController.go:19:10: error returned from external package is unwrapped: sig: func encoding/json.Marshal(v any) ([]byte, error) (wrapcheck)
                return jsonErr
                       ^
internal/product/application/controllers/SaveProductController.go:17:10: error returned from external package is unwrapped: sig: func (*github.com/ldez/go-monorepo-example/product/application/usecases.SaveProductUseCase).Execute(productDto github.com/ldez/go-monorepo-example/product/application/dtos.Product) error (wrapcheck)
                return err
                       ^
internal/product/application/usecases/SaveProductUseCase.go:24:10: error returned from external package is unwrapped: sig: func (*github.com/ldez/go-monorepo-example/product/domain/entities.Product).Validate() error (wrapcheck)
                return err
                       ^
internal/product/domain/entities/Product.go:38:10: error returned from external package is unwrapped: sig: func (*github.com/ldez/go-monorepo-example/product/domain/valueobjects.Money).Validate() error (wrapcheck)
                return priceError
                       ^
internal/product/application/usecases/ListProductUseCase.go:16:2: only one cuddle assignment allowed before range statement (wsl)
        for _, product := range products {
        ^
internal/product/domain/valueobjects/Money.go:32:3: ranges should only be cuddled with assignments used in the iteration (wsl)
                for i, c := range intStr {
                ^
internal/product/domain/valueobjects/Money.go:36:4: assignments should only be cuddled with other assignments (wsl)
                        result += string(c)
                        ^
internal/product/domain/valueobjects/Money.go:38:3: assignments should only be cuddled with other assignments (wsl)
                intStr = result
                ^
49 issues:
* errcheck: 2
* exhaustruct: 2
* mnd: 3
* nlreturn: 3
* perfsprint: 1
* revive: 30
* wrapcheck: 4
* wsl: 4
/home/ldez/sources/experimental/guidiego-monorepo-example/pkg/httpadapter
pkg/httpadapter/adapter.go:1:1: package-comments: should have a package comment (revive)
package httpadapter
^
pkg/httpadapter/adapter.go:5:6: exported: exported type IHttpAdapter should have comment or be unexported (revive)
type IHttpAdapter interface {
     ^
pkg/httpadapter/echo.go:5:1: exported: exported function EchoAdapter should have comment or be unexported (revive)
func EchoAdapter(controller IHttpAdapter) echo.HandlerFunc {
^
3 issues:
* revive: 3
$                                                                                    
```
