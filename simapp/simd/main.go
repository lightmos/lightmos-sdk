package main

import (
	"os"

	"github.com/lightmos/lightmos-sdk/server"
	svrcmd "github.com/lightmos/lightmos-sdk/server/cmd"
	"github.com/lightmos/lightmos-sdk/simapp"
	"github.com/lightmos/lightmos-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
