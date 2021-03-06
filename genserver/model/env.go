package model

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// MyEnv 环境定义
type MyEnv struct {
	ServerName  string      // 微服务的名字
	UsePort     string      // 端口号
	EntityList  []*MyEntity // 新加的实体列表
	ModelName   string      // 引用 第一个实体的名称（由于可能只有一个实体，为了便捷所做的冗余）
	ModelZH     string      // 引用 第一个实体的备注
	ShowExample bool        // 是否显示示例在生成代码中

	ProjectBasePath string // 项目路径
	ClusterPath     string // 集群文件夹路径
	DeployPath      string // 部署文夹路径
	ProtoPath       string // pb文夹路径
	GraphqlPath     string // graphql文夹路径
	GatePath        string // gate文夹路径
	ConfigPath      string // 配置文夹路径
	BundlePath      string // Bundle文件夹路径
}

// Encode Encode
func (m MyEnv) Encode(filePath string) error {
	bs, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, bs, 0600)
	if err != nil {
		return err
	}
	return nil
}

// Decode Decode
func Decode(filePath string) (*MyEnv, error) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	ret := &MyEnv{}
	err = yaml.Unmarshal(bs, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
