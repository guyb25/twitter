package main

import "twitter/internal/bootstrap"

func main() {
	bootstrap.NewRoutes().Bootstrap().Run(":8080")
}
