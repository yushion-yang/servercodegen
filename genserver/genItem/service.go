package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// ServiceGenerate ServiceGenerate
type ServiceGenerate struct {
}

// PreCheck PreCheck
func (g *ServiceGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*ServiceGenerate)(nil)

// GenCode GenCode
func (g ServiceGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/service.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/service/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		gencore.GenProto(filePath, v.ModelName, inputFiles, env)
	}
}
