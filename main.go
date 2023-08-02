package main

import (
	"fmt"

	"github.com/lixoi/prjct_validator/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
