package codegen

import (
	"fmt"
	"strings"
)

func GenerateBizCode(config *GenConfig, model *DbModel) {
	if config.Gen.GenBiz == nil || !config.Gen.GenBiz.Enable {
		return
	}

	bizTemplate := GetTemplate(TemplateBizKey)
	if len(bizTemplate) < 1 {
		return
	}

	tableName := model.TableName
	modelName := model.StructName
	// hasDeletedColumn bool

	fileName := config.Gen.OutputRoot + "/biz/" + tableName + ".gen.go"
	if FileExist(fileName) && !config.Gen.GenBiz.OverWrite {
		fmt.Println("File exist, no need generate new: " + fileName)
	} else {
		fmt.Println("Generate " + fileName)
		codes := strings.Replace(bizTemplate, "{TableModelName}", modelName, -1)

		codes = strings.Replace(codes, "{GenPackageRoot}", config.Gen.PackageRoot, -1)

		if model.HasDeletedColumn() {
			deletedTemplate := GetTemplate(TemplateBizDeletedKey)
			if len(deletedTemplate) > 0 {
				codes += strings.Replace(deletedTemplate, "{TableModelName}", model.StructName, -1)
			}

			// for query:
			whereStr := "tx = tx.Where(" + model.StructName + "Dao.Deleted.Eq(0))"
			codes = strings.Replace(codes, "// {QueryDeletedWhereCondition}", whereStr, -1)
		}

		WriteFile(fileName, codes)

	}
}
