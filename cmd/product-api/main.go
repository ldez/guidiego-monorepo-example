package main

import "github.com/ldez/go-monorepo-example/product-api/app"

func main() {
	app.New().Start("0.0.0.0", 3000)
}
