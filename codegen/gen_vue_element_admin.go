package codegen

import (
	"os"
	"strings"
)

var (
	VueElementAdminGen = &VueElementAdminGenType{}
)

type (
	VueElementAdminGenType struct {
	}
)

func (v *VueElementAdminGenType) GenerateVueCodes(config *GenConfig, model *DbModel) {
	if config.Gen.GenVueElementAdmin == nil || !config.Gen.GenVueElementAdmin.Enable || len(config.Gen.GenVueElementAdmin.ProjectRoot) < 1 {
		return
	}

	v.GenVueClientApis(config, model)
	v.GenVueViews(config, model)
}

func (v *VueElementAdminGenType) GenVueClientApis(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateVueElementAdminApiKey)
	// deleted column:
	if model.HasDeletedColumn() {
		templateCodes += GetTemplate(TemplateVueElementAdminApiDeletedColumnKey)
	}

	templateCodes = strings.Replace(templateCodes, "{TableModelName}", model.StructName, -1)
	templateCodes = strings.Replace(templateCodes, "{TableModelNameLowCase}", strings.ToLower(model.StructName), -1)

	queryParams := v.getVueQueryParameters(model)
	templateCodes = strings.Replace(templateCodes, "// {其它查询参数}", queryParams, -1)

	vueFolder := config.Gen.GenVueElementAdmin.ProjectRoot + "/src/api"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/"+model.PrivatePropertyName+".js", templateCodes)
}

func (v *VueElementAdminGenType) GenVueViews(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateVueElementAdminViewKey)

	templateCodes = strings.Replace(templateCodes, "{TableModelName}", model.StructName, -1)
	templateCodes = strings.Replace(templateCodes, "{LowerFirstCharTableModelName}", model.PrivatePropertyName, -1)

	// list table view:
	columnsCodes := ""
	for _, modelField := range model.Fields {
		if modelField.IgnoreGenerateResponseModel || len(modelField.ModelFieldName) < 1 {
			continue
		}

		codes := "      <el-table-column align=\"left\" sortable label=\"" + modelField.ModelFieldName + "\">\n"
		codes += "        <template slot-scope=\"{ row }\">\n"
		codes += "          <span>{{ row." + modelField.ColumnName + " }}</span>\n"
		codes += "        </template>\n"
		codes += "      </el-table-column>\n\n"

		columnsCodes += codes
	}
	templateCodes = strings.Replace(templateCodes, "// {ModelTableColumns}", columnsCodes, -1)

	// create form view:
	createFormItems := ""
	editFormItems := ""
	createFormItemNames := ""
	for _, modelField := range model.Fields {
		if modelField.IgnoreGenerateResponseModel || len(modelField.ModelFieldName) < 1 {
			continue
		}
		codes := ""

		if !strings.EqualFold(modelField.ModelFieldName, "ID") {
			if strings.Contains(modelField.ModelFieldType, "time") {
				codes += "			<el-form-item label=\"" + modelField.ModelFieldName + "\">\n"
				codes += "			<el-date-picker\n"
				codes += "			  type=\"date\"\n"
				codes += "			  placeholder=\"选择日期\"\n"
				codes += "			  v-model=\"form." + modelField.ColumnName + "\"\n"
				codes += "			  style=\"width: 100%\"\n"
				codes += "			></el-date-picker>\n"
				codes += "		  </el-form-item>\n"
			} else {
				valueMode := ""
				if strings.HasPrefix(modelField.ModelFieldType, "int") || strings.HasPrefix(modelField.ModelFieldType, "uint") {
					valueMode = ".number"
				}

				codes += "        <el-form-item label=\"" + modelField.ModelFieldName + "\">\n"
				codes += "          <el-input v-model" + valueMode + "=\"form." + modelField.ColumnName + "\"></el-input>\n"
				codes += "        </el-form-item>\n"
			}

			createFormItems += codes
			editFormItems += codes
		}

		if strings.HasPrefix(modelField.ModelFieldType, "int") || strings.HasPrefix(modelField.ModelFieldType, "uint") {
			createFormItemNames += "        " + modelField.ColumnName + ": 0,\n"
		} else {
			createFormItemNames += "        " + modelField.ColumnName + ": '',\n"
		}
	}
	templateCodes = strings.Replace(templateCodes, "// {ModelCreateFormItems}", createFormItems, -1)
	templateCodes = strings.Replace(templateCodes, "// {ModelEditFormItems}", editFormItems, -1)
	templateCodes = strings.Replace(templateCodes, "// {ModelCreateFormItemNames}", createFormItemNames, -1)

	fileName := strings.ToLower(model.StructName)
	vueFolder := config.Gen.GenVueElementAdmin.ProjectRoot + "/src/views/gen"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/"+fileName+".vue", templateCodes)
}

func (v *VueElementAdminGenType) GenerateVueRouter(config *GenConfig, models []*DbModel) {
	if config.Gen.GenVueElementAdmin == nil || !config.Gen.GenVueElementAdmin.Enable || len(config.Gen.GenVueElementAdmin.ProjectRoot) < 1 {
		return
	}

	templateCodes := GetTemplate(TemplateVueElementAdminRouterKey)

	codes := ""
	for _, model := range models {

		codes += "    {\n"
		codes += "      path: '" + strings.ToLower(model.StructName) + "',\n"
		codes += "      component: () => import('@/views/gen/" + strings.ToLower(model.StructName) + "'),\n"
		codes += "      name: '" + model.StructName + "',\n"
		codes += "      meta: { title: '" + model.StructName + "' }\n"
		codes += "    },\n"

	}
	codes = codes[0:(len(codes) - 2)]
	templateCodes = strings.Replace(templateCodes, "// {VueRouterModelCodes}", codes, -1)

	vueFolder := config.Gen.GenVueElementAdmin.ProjectRoot + "/src/router/modules"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/goCodeGen.js", templateCodes)
}

func (v *VueElementAdminGenType) getVueQueryParameters(model *DbModel) string {
	codes := ""
	for _, modelField := range model.Fields {
		if !modelField.IgnoreGenerateRequestQueryParameter && len(modelField.PrivateModelFieldName) > 0 {
			extComments := ""
			if strings.HasPrefix(modelField.ModelFieldType, "uint") || strings.HasPrefix(modelField.ModelFieldType, "int") {
				extComments = "多个时，请用逗号隔开"
			}
			codes += "// " + modelField.PrivateModelFieldName + ": " + modelField.ModelFieldType + "," + extComments + "\n"
		}
	}
	if len(codes) > 2 {
		codes = codes[0:(len(codes) - 1)] // 去掉末尾换行符
	}

	return codes
}
