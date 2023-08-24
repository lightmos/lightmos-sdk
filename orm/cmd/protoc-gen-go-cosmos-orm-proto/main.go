package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/lightmos/lightmos-sdk/orm/internal/codegen"
)

func main() {
	protogen.Options{}.Run(codegen.QueryProtoPluginRunner)
}
