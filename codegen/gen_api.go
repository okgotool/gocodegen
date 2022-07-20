package codegen

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateApis(config *GenConfig, model *DbModel) {
	if config.Gen.GenApi == nil || !config.Gen.GenApi.Enable {
		return
	}

	tableName := model.TableName
	modelName := model.StructName

	apiTemplateStr := GetTemplate(TemplateApiKey)

	apiRouterTemplateStr := ""
	if model.HasDeletedColumn() {
		apiRouterTemplateStr = GetTemplate(TemplateApirouterDeletedKey)
	} else {
		apiRouterTemplateStr = GetTemplate(TemplateApirouterKey)
	}

	responseModelTemplateStr := GetTemplate(TemplateResponseModelKey)
	if len(apiTemplateStr) < 1 || len(responseModelTemplateStr) < 1 {
		return
	}

	generateApiRouterCode(config, tableName, modelName, apiRouterTemplateStr)
	generateResponseModelCode(config, tableName, modelName, responseModelTemplateStr)
	generateApiMethodCode(config, tableName, modelName, apiTemplateStr, model)
}

func generateApiMethodCode(config *GenConfig, tableName string, modelName string, apiTemplate string, model *DbModel) {

	fileName := config.Gen.OutputRoot + "/api/" + tableName + "_api.gen.go"
	if FileExist(fileName) && !config.Gen.GenApi.OverWrite {
		fmt.Println("File exist, no need generate new: " + fileName)
	} else {
		fmt.Println("Generate " + fileName)
		codes := strings.Replace(apiTemplate, "{TableModelName}", modelName, -1)
		codes = strings.Replace(codes, "{TableModelNameLowCase}", strings.ToLower(modelName), -1)

		// swagger parameters:
		// generate query condition according to model properties:
		swaggerParameters := ""
		queryConditions := ""
		for _, modelField := range model.Fields {

			modelFieldName := modelField.ModelFieldName
			modelFieldType := modelField.ModelFieldType

			// ignore fields:
			if strings.EqualFold(modelFieldName, "id") {
				fmt.Println(" -- Ignore field: " + modelFieldName)
				continue
			}
			if modelField.IgnoreGenerateRequestQueryParameter {
				fmt.Println(" -- Ignore field by config: " + modelFieldName)
				continue
			}
			// isExcludeField := false
			// for _, excludeField := range config.Gen.GenApi.ExcludeColumnsForQueryParameters {
			// 	if strings.EqualFold(modelFieldName, excludeField) {
			// 		fmt.Println("Ignore field by config: " + modelFieldName)
			// 		isExcludeField = true
			// 		break
			// 	}
			// }
			// if isExcludeField {
			// 	continue
			// }

			paramName := modelField.PrivateModelFieldName // lowercase for first chart

			fieldQueryConditions := ""
			parameterType := "string"
			if strings.EqualFold(modelFieldType, "string") {
				fieldQueryConditions += getApiStringQueryCondition(paramName, modelName, modelFieldName)

				swaggerParameters += "// @Param " + paramName + " query " + parameterType + " false \"" + modelFieldName + "\" default()\n"
			} else if strings.EqualFold(modelFieldType, "int") ||
				strings.EqualFold(modelFieldType, "int8") ||
				strings.EqualFold(modelFieldType, "int16") ||
				strings.EqualFold(modelFieldType, "int32") ||
				strings.EqualFold(modelFieldType, "int64") ||
				strings.EqualFold(modelFieldType, "uint") ||
				strings.EqualFold(modelFieldType, "uint8") ||
				strings.EqualFold(modelFieldType, "uint16") ||
				strings.EqualFold(modelFieldType, "uint32") ||
				strings.EqualFold(modelFieldType, "uint64") {

				// parameterType = "int"
				fieldQueryConditions += getApiIntQueryCondition(paramName, modelName, modelFieldName, modelFieldType)

				swaggerParameters += "// @Param " + paramName + " query " + parameterType + " false \"" + modelFieldName + ", 数字，多个时逗号隔开\" default()\n"
			} else if strings.EqualFold(modelFieldType, "time.Time") {
				fieldQueryConditions += getApiTimeQueryCondition(paramName, modelName, modelFieldName, modelFieldType)

				swaggerParameters += "// @Param " + paramName + "Min query int64 false \"" + modelFieldName + " 起始时间, 毫秒数时间戳，查询大于等于" + paramName + "Min 该时间的数据\" default()\n"
				swaggerParameters += "// @Param " + paramName + "Max query int64 false \"" + modelFieldName + " 结束时间, 毫秒数时间戳，查询小于" + paramName + "Max 该时间的数据\" default()\n"
			} else {
				fmt.Println(" -- Ignore modelFieldName=" + modelFieldName + ", modelFieldType=" + modelFieldType)
				continue
			}

			if len(fieldQueryConditions) > 0 {
				queryConditions += fieldQueryConditions
			}
		}

		codes = strings.Replace(codes, "{ParameterQuerySwaggerParameters}", swaggerParameters, -1)
		codes = strings.Replace(codes, "{ParameterQueryConditions}", queryConditions, -1)

		codes = strings.Replace(codes, "{GenPackageRoot}", config.Gen.PackageRoot, -1)

		if model.HasDeletedColumn() {
			deletedTemplate := GetTemplate(TemplateApiDeletedKey)
			if len(deletedTemplate) > 0 {
				codes += strings.Replace(deletedTemplate, "{TableModelName}", model.StructName, -1)
			}
		}

		codes = strings.Replace(codes, "{TableModelNameLowCase}", strings.ToLower(modelName), -1)
		codes = strings.Replace(codes, "{PrimaryKeyPropertyName}", model.PrimaryKeyPropertyName, -1)

		WriteFile(fileName, codes)
	}
}

func getApiStringQueryCondition(paramName string, modelName string, modelFieldName string) string {
	str := "\t\tif len(g.Query(\"" + paramName + "\")) > 0 {\n"

	str += "\t\t\tqueryValue := g.Query(\"" + paramName + "\")\n"
	str += "\t\t\twhereConditions = append(whereConditions, biz." + modelName + "Dao." + modelFieldName + ".Eq(queryValue))\n"

	str += "\t\t}\n"
	return str
}

func getApiIntQueryCondition(paramName string, modelName string, modelFieldName string, fieldType string) string {

	str := "\n"
	str += "\t\tif len(g.Query(\"" + paramName + "\")) > 0 {\n"

	str += "\t\t\tqueryValues := []" + fieldType + "{}\n"
	str += "\t\t\tqueryStrs := strings.Split(g.Query(\"" + paramName + "\"), \",\")\n"
	str += "\t\t\tfor _, queryStr := range queryStrs {\n"

	// for uint and  int:
	if strings.HasPrefix(strings.ToLower(fieldType), "uint") {
		str += "\t\t\t\tqueryValue, err := parseUint64(queryStr)\n"
	} else {
		str += "\t\t\t\tqueryValue, err := parseInt64(queryStr)\n"
	}

	str += "\t\t\t\tif err != nil {\n"
	str += "\t\t\t\t} else {\n"
	str += "\t\t\t\t\tqueryValues = append(queryValues, " + fieldType + "(queryValue))\n"
	str += "\t\t\t\t}\n"
	str += "\t\t\t}\n"
	str += "\t\t\tif len(queryValues) > 0 {\n"
	str += "\t\t\t\twhereConditions = append(whereConditions, biz." + modelName + "Dao." + modelFieldName + ".In(queryValues...))\n"
	str += "\t\t\t}\n"

	str += "\t\t}\n"
	str += "\n"

	return str
}

func getApiTimeQueryCondition(paramName string, modelName string, modelFieldName string, fieldType string) string {

	str := "\n"
	str += "\t\t// query data of " + paramName + " between " + paramName + "Min and " + paramName + "Max:\n"

	str += "\t\tif len(g.Query(\"" + paramName + "Min\")) > 0 {\n"
	str += "\t\t\t" + paramName + "Mills, err := strconv.ParseInt(g.Query(\"" + paramName + "Min\"), 10, 64)\n"
	str += "\t\t\tif err == nil {\n"
	str += "\t\t\t\t" + paramName + "Min := time.Unix(" + paramName + "Mills/1000, 0)\n"
	str += "\t\t\t\twhereConditions = append(whereConditions, biz." + modelName + "Dao." + modelFieldName + ".Gte(" + paramName + "Min))\n"
	str += "\t\t\t}\n"
	str += "\t\t}\n"

	str += "\t\tif len(g.Query(\"" + paramName + "Max\")) > 0 {\n"
	str += "\t\t\t" + paramName + "Mills, err := strconv.ParseInt(g.Query(\"" + paramName + "\"), 10, 64)\n"
	str += "\t\t\tif err == nil {\n"
	str += "\t\t\t\t" + paramName + "Max := time.Unix(" + paramName + "Mills/1000, 0)\n"
	str += "\t\t\t\twhereConditions = append(whereConditions, biz." + modelName + "Dao." + modelFieldName + ".Lt(" + paramName + "Max))\n"
	str += "\t\t\t}\n"
	str += "\t\t}\n"

	str += "\n"

	return str
}

func generateApiRouterCode(config *GenConfig, tableName string, modelName string, ginRouterTemplate string) {
	if !config.Gen.GenApi.Enable {
		return
	}

	codes := ""
	fileName := config.Gen.OutputRoot + "/api/gin_router.gen.go"

	if FileExist(fileName) {
		fileBytes, err := ioutil.ReadFile(fileName)
		if err == nil {
			codes = string(fileBytes)
		}
	}

	if !strings.Contains(codes, "package api") {
		codes = "package api\n\n" + codes
	}
	if !strings.Contains(codes, "github.com/gin-gonic/gin") {
		codes = strings.Replace(codes, "package api", "package api\n\nimport \"github.com/gin-gonic/gin\"", -1)
	}

	// add main add api method if not exist:
	if !strings.Contains(codes, "func AddApis(api *gin.RouterGroup)") {
		codes += "func AddApis(api *gin.RouterGroup) {\n"
		codes += "\tv1 := api.Group(\"" + config.Gen.GenApi.ApiGroup + "\")\n\n"
		codes += "\t// {otherApisPlaceHolder}\n\n"
		codes += "}\n"
	}

	// add model api add to router method if not exist:
	if !strings.Contains(codes, "add"+modelName+"Apis(v1)") {
		codes = strings.Replace(codes, "// {otherApisPlaceHolder}", "add"+modelName+"Apis(v1)\n\t// {otherApisPlaceHolder}", -1)
	}

	// old codes not include the method, then add it:
	if !strings.Contains(codes, "func add"+modelName+"Apis(v1 *gin.RouterGroup)") {
		templateCodes := strings.Replace(ginRouterTemplate, "{TableModelName}", modelName, -1)
		templateCodes = strings.Replace(templateCodes, "{TableModelNameLowCase}", strings.ToLower(modelName), -1)

		codes += "\n\n" + templateCodes
	}

	codes = strings.Replace(codes, "{GenPackageRoot}", config.Gen.PackageRoot, -1)

	WriteFile(fileName, codes)
}

func generateResponseModelCode(config *GenConfig, tableName string, modelName string, bizTemplate string) {
	fileName := config.Gen.OutputRoot + "/model/response/" + tableName + ".gen.go"
	if FileExist(fileName) && !config.Gen.GenApi.OverWrite {
		fmt.Println("File exist, no need generate new: " + fileName)
	} else {
		fmt.Println("Generate " + fileName)
		codes := strings.Replace(bizTemplate, "{TableModelName}", modelName, -1)

		codes = strings.Replace(codes, "{GenPackageRoot}", config.Gen.PackageRoot, -1)

		WriteFile(fileName, codes)
	}
}
