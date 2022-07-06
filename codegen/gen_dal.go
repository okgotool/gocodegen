package codegen

import (
	"fmt"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateDalCodes(conn *gorm.DB, config *GenConfig) {
	if !config.Gen.GenDal.Enable {
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: config.Gen.OutputRoot + "/dal/query",
	})

	g.UseDB(conn)

	// generate tables from database
	if len(config.Gen.GenTables) < 1 {
		g.ApplyBasic(g.GenerateAllTable()...)
	} else {

		for _, table := range config.Gen.GenTables {

			fmt.Println("Start to generate table: " + table.TableName + " ...")

			filedGenOps := []gen.FieldOpt{}
			for _, field := range table.Fields {
				if field.IsIgnore { // if ignore field:
					filedGenOps = append(filedGenOps, gen.FieldIgnore(field.ColumnName))
					continue
				} else { // if not ignore field:
					if len(field.FieldType) > 0 { // if reset model field type:
						filedGenOps = append(filedGenOps, gen.FieldType(field.ColumnName, field.FieldType))
					}
					if len(field.FieldName) > 0 { // if reset model field name:
						filedGenOps = append(filedGenOps, gen.FieldRename(field.ColumnName, field.FieldName))
					}
				}
			}

			// if len(table.ModelName) < 1 {
			g.ApplyBasic(g.GenerateModel(table.TableName, filedGenOps...))
			// }
			// else {
			// 	g.ApplyBasic(g.GenerateModelAs(table.TableName, table.ModelName, filedGenOps...))
			// }
		}
	}

	g.Execute()
}

func GenerateAutomigrate(config *GenConfig, models []*DbModel) {
	if config.Gen.GenDal == nil || !config.Gen.GenDal.Enable {
		return
	}

	codes := "package dal\n\n"
	codes += "import \"" + config.Gen.PackageRoot + "/dal/model\"\n\n"
	codes += "func AutoMigrateTables() {\n"

	for _, model := range models {
		codes += "\tCreateDbConn(\"" + model.StructName + "\").AutoMigrate(&model." + model.StructName + "{})\n"
	}

	codes += "}\n"

	fileName := config.Gen.OutputRoot + "/dal/automigrate_tables.gen.go"
	WriteFile(fileName, codes)
}
