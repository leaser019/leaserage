package main

import (
	"os"

	"github.com/vomkhang/leaserage/internal/app"
)

func main() {
	os.Exit(app.New(os.Stdout, os.Stderr).Run(os.Args[1:]))
}
