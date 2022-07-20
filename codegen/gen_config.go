package codegen

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type (
	GenConfig struct {
		MySql *GenConfigMysql `yaml:"mysql"`
		Gen   *GenConfigGen   `yaml:"gen"`
	}

	GenConfigMysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Coon     string `yaml:"conn"`
	}

	GenConfigGen struct {
		DmlFolder          string                       `yaml:"dmlFolder"`
		RunDml             *GenConfigRunDml             `yaml:"runDml"`
		GenRuoyiVue3       *GenConfigGenRuoyiVue3       `yaml:"genRuoyiVue3"`
		GenVueElementAdmin *GenConfigGenVueElementAdmin `yaml:"genVueElementAdmin"`
		GenApi             *GenConfigGenApi             `yaml:"genApi"`
		GenBiz             *GenConfigGenBiz             `yaml:"genBiz"`
		GenDataSource      *GenConfigGenDataSource      `yaml:"genDataSource"`
		GenDal             *GenConfigGenDal             `yaml:"genDal"`
		DataSources        []*GenConfigDataSource       `yaml:"dataSources"`
		GenTables          []*GenConfigTable            `yaml:"genTables"`
		OutputRoot         string                       `yaml:"outputRoot"`
		PackageRoot        string                       `yaml:"packageRoot"`
	}

	GenConfigRunDml struct {
		Enable bool `yaml:"enable"`
	}

	GenConfigGenRuoyiVue3 struct {
		Enable              bool                        `yaml:"enable"`
		ProjectRoot         string                      `yaml:"projectRoot"`
		ApiProxyServiceName string                      `yaml:"apiProxyServiceName"`
		TopMenus            []*GenConfigGenRuoyiVueMenu `yaml:"topMenus"`
	}

	GenConfigGenRuoyiVueMenu struct {
		Name       string                      `yaml:"name"`
		Path       string                      `yaml:"path"`
		Icon       string                      `yaml:"icon"`
		OrderNum   string                      `yaml:"orderNum"`
		ChildMenus []*GenConfigGenRuoyiVueMenu `yaml:"childMenus"`
	}

	GenConfigGenVueElementAdmin struct {
		Enable      bool   `yaml:"enable"`
		ProjectRoot string `yaml:"projectRoot"`
	}

	GenConfigGenApi struct {
		Enable                               bool     `yaml:"enable"`
		OverWrite                            bool     `yaml:"overWrite"`
		ApiGroup                             string   `yaml:"apiGroup"`
		GenApis                              []string `yaml:"genApis"`
		ExcludeModelFieldsForQueryParameters []string `yaml:"excludeModelFieldsForQueryParameters"`
		ExcludeModelFieldsForResponse        []string `yaml:"excludeModelFieldsForResponse"`
	}

	GenConfigGenBiz struct {
		Enable    bool `yaml:"enable"`
		OverWrite bool `yaml:"overWrite"`
	}

	GenConfigGenDataSource struct {
		Enable    bool `yaml:"enable"`
		OverWrite bool `yaml:"overWrite"`
	}

	GenConfigGenDal struct {
		Enable    bool `yaml:"enable"`
		OverWrite bool `yaml:"overWrite"`
	}

	GenConfigDataSource struct {
		DataSourceName string `yaml:"dataSourceName"`
		Dsn            string `yaml:"dsn"`
		Code           string `yaml:"code"`
	}

	GenConfigTable struct {
		TableName      string            `yaml:"tableName"`
		DataSourceName string            `yaml:"dataSourceName"`
		ModelName      string            `yaml:"modelName"`
		Fields         []*GenConfigField `yaml:"fields"`
		GenApis        []string          `yaml:"genApis"`
		ParentMenu     string            `yaml:"parentMenu"`
		MenuName       string            `yaml:"menuName"`
	}

	GenConfigField struct {
		ColumnName string `yaml:"columnName"`
		FieldName  string `yaml:"fieldName"`
		FieldType  string `yaml:"fieldType"`
		IsIgnore   bool   `yaml:"isIgnore"`
	}
)

func parseYaml(configPath string) *GenConfig {
	config := &GenConfig{}
	fileBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Failed to read config file: " + configPath)
	}

	err = yaml.Unmarshal(fileBytes, config)
	if err != nil {
		panic("Failed to parse config file: " + configPath)
	}

	// set default:
	if config.Gen.RunDml == nil {
		config.Gen.RunDml = &GenConfigRunDml{
			Enable: false,
		}
	}
	if config.Gen.GenDal == nil {
		config.Gen.GenDal = &GenConfigGenDal{
			Enable:    true,
			OverWrite: true,
		}
	}
	if config.Gen.GenBiz == nil {
		config.Gen.GenBiz = &GenConfigGenBiz{
			Enable:    true,
			OverWrite: true,
		}
	}
	if config.Gen.GenDataSource == nil {
		config.Gen.GenDataSource = &GenConfigGenDataSource{
			Enable:    true,
			OverWrite: true,
		}
	}
	if config.Gen.GenApi == nil {
		config.Gen.GenApi = &GenConfigGenApi{
			Enable:    true,
			OverWrite: true,
			ApiGroup:  "v1",
		}
	}

	return config
}
