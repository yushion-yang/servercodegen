package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// NeventGenerate NeventGenerate
type NeventGenerate struct {
}

// PreCheck PreCheck
func (g *NeventGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*NeventGenerate)(nil)

// GenCode GenCode
func (g NeventGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/nevent.tmpl"}
	filePath := fmt.Sprintf("%v%v/nevent.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
