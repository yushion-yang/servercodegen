package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type MainGenerate struct {
}

func (g *MainGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*MainGenerate)(nil)

func (g MainGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/main.tmpl"}
	filePath := fmt.Sprintf("%v%v/main.go", env.ClusterPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
