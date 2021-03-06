package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// RepositoryGenerate RepositoryGenerate
type RepositoryGenerate struct {
}

// PreCheck PreCheck
func (g *RepositoryGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*RepositoryGenerate)(nil)

// GenCode GenCode
func (g RepositoryGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/repository.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/repository/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		gencore.GenProto(filePath, v.ModelName, inputFiles, env)
	}
}
