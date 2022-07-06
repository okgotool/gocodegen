package codegen

import (
	"fmt"
	"strings"
)

var (
	DataSourceTemplateContent = ""
)

func GenerateDataSources(config *GenConfig) {
	if config.Gen.GenDataSource == nil || !config.Gen.GenDataSource.Enable {
		return
	}

	fileName := config.Gen.OutputRoot + "/dal/datasources.gen.go"
	if FileExist(fileName) && !config.Gen.GenDataSource.OverWrite {
		fmt.Println("File exist, no need generate new: " + fileName)
	} else {
		dataSourceTemplate := GetTemplate(TemplateDatasourceKey)
		if len(dataSourceTemplate) < 1 {
			return
		}

		dsList := ""
		defaultDsName := ""
		for _, ds := range config.Gen.DataSources {
			if len(defaultDsName) < 1 {
				defaultDsName = ds.DataSourceName
			}
			if len(ds.Dsn) > 0 {
				dsList += ds.DataSourceName + " = ConnectDB(\"" + ds.Dsn + "\").Debug()\n\t"
			} else if len(ds.Code) > 0 {
				dsList += ds.DataSourceName + " = " + ds.Code + "\n\t"
			} else {
				fmt.Println("datasource code not set: " + ds.DataSourceName)
			}
		}
		codes := strings.Replace(dataSourceTemplate, "// {GenDataSources}", dsList, -1)

		modelDsMap := map[string]string{}
		for _, table := range config.Gen.GenTables {
			if len(table.ModelName) > 0 && len(table.DataSourceName) > 0 {
				modelDsMap[table.ModelName] = table.DataSourceName
			}
		}

		getDsCodes := ""
		datasourceImports := ""
		if len(modelDsMap) < 1 { // not set special tabel datasource, return default
			getDsCodes = "return " + defaultDsName
		} else {
			for modelName, dsName := range modelDsMap {
				getDsCodes += "if strings.EqualFold(modelName, \"" + modelName + "\") {\n\t\treturn " + dsName + "\n\t} else "
			}
			getDsCodes += " {\n\t\treturn " + defaultDsName + "\n\t}\n"
			datasourceImports += "\"strings\""
		}

		codes = strings.Replace(codes, "// {GenDataSourcesImports}", datasourceImports, -1)
		codes = strings.Replace(codes, "// {GenGetDataSourceByModelName}", getDsCodes, -1)

		fmt.Println("Generate " + fileName)
		WriteFile(fileName, codes)
	}
}
