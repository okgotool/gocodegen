package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

func StartGen(configFile string) {
	if _, err := os.Stat(configFile); err != nil {
		panic("Config file not exist: " + configFile)
	}

	config := parseYaml(configFile)

	StartGenConfig(config)
}

func StartGenConfig(config *GenConfig) {
	dbDns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.MySql.User, config.MySql.Password, config.MySql.Host, config.MySql.Port, config.MySql.Database)
	conn := ConnectDB(dbDns).Debug()

	// run db dml: create tables
	if config.Gen.RunDml.Enable {
		runDbDmls(conn, config)
	}

	// generate folders:
	GenerateFolders(config)

	GenerateDalCodes(conn, config)
	//
	dbModels := getDbTableModels(config)
	if len(dbModels) < 1 {
		fmt.Println("Not found models, exit.")
		return
	}

	// set ignore fields:
	for _, model := range dbModels {
		for _, modelField := range model.Fields {
			for _, ignoreField := range config.Gen.GenApi.ExcludeModelFieldsForQueryParameters {
				if strings.EqualFold(ignoreField, modelField.ModelFieldName) {
					modelField.IgnoreGenerateRequestModel = true
					modelField.IgnoreGenerateRequestQueryParameter = true
				}
			}
			for _, ignoreField := range config.Gen.GenApi.ExcludeModelFieldsForResponse {
				if strings.EqualFold(ignoreField, modelField.ModelFieldName) {
					modelField.IgnoreGenerateResponseModel = true
				}
			}
		}
	}

	GenerateAutomigrate(config, dbModels)

	GenerateDataSources(config)
	for _, model := range dbModels {
		GenerateBizCode(config, model)
		GenerateApis(config, model)
		VueElementAdminGen.GenerateVueCodes(config, model)
	}
	VueElementAdminGen.GenerateVueRouter(config, dbModels)
}

func getDbTableModels(config *GenConfig) []*DbModel {
	dbTables := []*DbModel{}

	modelFileDir := config.Gen.OutputRoot + "/dal/model"
	modelFileList, err := ioutil.ReadDir(modelFileDir)
	if err != nil {
		fmt.Println("List files faild: "+modelFileDir+",erro:%s", err)
		return dbTables
	}

	for _, modelFile := range modelFileList {
		if !modelFile.IsDir() && strings.HasSuffix(modelFile.Name(), ".gen.go") {
			modelTableName := strings.Replace(modelFile.Name(), ".gen.go", "", -1)

			// if len(config.Gen.GenTables) > 0 {
			// 	isGenTable := false
			// 	for _, table := range config.Gen.GenTables {
			// 		if strings.EqualFold(tableName, table.TableName) {
			// 			isGenTable = true
			// 			break
			// 		}
			// 	}
			// 	if !isGenTable {
			// 		fmt.Println("Not need gen table: " + tableName)
			// 		continue
			// 	}
			// }

			modelFilePath := modelFileDir + "/" + modelFile.Name()
			modelFileContent, err := ioutil.ReadFile(modelFilePath)
			if err != nil {
				fmt.Println("Read file failed: "+modelFilePath+",erro:%s", err)
				continue
			}

			// start to decode model file:
			astFileSet := token.NewFileSet()
			astFile, err := parser.ParseFile(astFileSet, "", string(modelFileContent), 0)
			if err != nil {
				fmt.Println("Parse file failed: "+modelFilePath+", err = %s", err)
				continue
			}
			// ast.Print(fset, f)

			ast.Inspect(astFile, func(astNode ast.Node) bool {
				if nil == astNode {
					return true
				}
				switch value := astNode.(type) {
				case *ast.TypeSpec: // is a type

					// ast.Print(fset, n)
					modelStructName := value.Name.Name
					fmt.Println()
					fmt.Println(" ---- Find model: " + modelStructName + " in " + modelFilePath)

					modelFields := getDbModeFields(astNode)

					primaryKeyPropertyName := "ID"
					for _, field := range modelFields {
						if field.IsPrimaryKey || field.IsAutoIncrement {
							primaryKeyPropertyName = field.ModelFieldName
							break
						}
					}

					table := &DbModel{
						TableName:              modelTableName,
						StructName:             modelStructName,
						PrivatePropertyName:    strings.ToLower(string(modelStructName[0])) + modelStructName[1:],
						Fields:                 modelFields,
						PrimaryKeyPropertyName: primaryKeyPropertyName,
					}

					dbTables = append(dbTables, table)
				}
				return true
			})
			// dal.DB.AutoMigrate(f.)
		} // end if
	} // end for model files

	return dbTables

}

func getDbModeFields(f ast.Node) []*DbModelFieldAndColumn {
	modelFields := []*DbModelFieldAndColumn{}

	ast.Inspect(f, func(astNode ast.Node) bool {
		if nil == astNode {
			return true
		}
		switch modelNode := astNode.(type) {
		case *ast.StructType: // is a struct
			// ast.Print(fset, n)
			for _, modelField := range modelNode.Fields.List {
				// ast.Print(fset, l.Type)
				modelFieldName := modelField.Names[0].Name

				privateModelFieldName := "id"
				if !strings.EqualFold(modelFieldName, "ID") {
					privateModelFieldName = strings.ToLower(string(modelFieldName[0])) + modelFieldName[1:] // lowercase for first chart
				}

				tags := modelField.Tag.Value // eg. "`gorm:\"column:id;primaryKey;autoIncrement:true\" json:\"id\"`"
				columnName := ""
				columnNames := strings.Split(tags, "column:")
				if len(columnNames) > 1 {
					columnName = strings.Trim(strings.Split(columnNames[1], ";")[0], " ")
					columnName = strings.Trim(strings.Split(columnName, "\"")[0], " ")
					columnName = strings.Trim(strings.Split(columnName, " ")[0], " ")
					columnName = strings.ToLower(columnName)
				}

				field := &DbModelFieldAndColumn{
					ModelFieldName:        modelFieldName,
					PrivateModelFieldName: privateModelFieldName,
					ColumnName:            columnName,
					IsPrimaryKey:          strings.Contains(strings.ToLower(tags), "primarykey"),
					IsAutoIncrement:       strings.Contains(strings.ToLower(tags), "autoincrement"),
				}
				modelFields = append(modelFields, field)

				switch modelFieldType := modelField.Type.(type) {
				case *ast.Ident: // is simple time
					field.ModelFieldType = modelFieldType.Name
					// ss := l.Type.(*ast.Ident)
					// ast.Print(fset, l.Type)
					// modelFieldName := modelField.Names[0].Name
					fmt.Println(" ---- -- Find Field: " + modelFieldName + "-->" + modelFieldType.Name)
					// privateModelFieldName := "id"
					// if !strings.EqualFold(modelFieldName, "ID") {
					// 	privateModelFieldName = strings.ToLower(string(modelFieldName[0])) + modelFieldName[1:] // lowercase for first chart
					// }

					// field := &DbModelFieldAndColumn{
					// 	ModelFieldName:        modelFieldName,
					// 	ModelFieldType:        modelFieldType.Name,
					// 	PrivateModelFieldName: privateModelFieldName,
					// }
					// modelFields = append(modelFields, field)

				case *ast.SelectorExpr: // is complex type, eg. time.Time
					fieldType := ""
					switch fieldTypeXType := modelFieldType.X.(type) {
					case *ast.Ident:
						fieldType += fieldTypeXType.Name
					}

					fieldType += "." + modelFieldType.Sel.Name

					fmt.Println(" ---- -- Find Field: "+modelField.Names[0].Name+"-->", fieldType)
					field.ModelFieldType = fieldType
					// field := &DbModelFieldAndColumn{
					// 	ModelFieldName: modelField.Names[0].Name,
					// 	ModelFieldType: fieldType,
					// }
					// modelFields = append(modelFields, field)
				default:
					field.ModelFieldType = "string"
					// fmt.Println(" ---- Find ignore Field: "+modelField.Names[0].Name+"-->", modelFieldType)
				} // end case
			} // end for
		} // end case
		return true
	})

	return modelFields
}
