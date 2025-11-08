package main

import (
	"embed"
	"fmt"
)

//go:embed db/migrations/*.sql
var fs embed.FS

func main() {

	fmt.Print("this is a test")
}
