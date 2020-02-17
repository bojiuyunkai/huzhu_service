package main

import (
	"huzhu_service/cmd/server"
	"fmt"
	"os"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
