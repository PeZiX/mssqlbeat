package main

import (
	"os"

	"github.com/PeZiX/mssqlbeat/cmd"

	_ "github.com/PeZiX/mssqlbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
