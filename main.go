package main

import "gin-fleamarket/routers"

func main() {
	routers.NewRouter().Run("localhost:8080")
}
