package main

import (
	"cleaner-serve/internal/routers"
	"fmt"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
	fmt.Println("serve run 8080")
}
