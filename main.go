package main

import (
	"bee.com/routers"
	_ "bee.com/conf"
)

func main() {
	r := routers.InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}