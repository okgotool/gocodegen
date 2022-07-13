package codegen

import (
	"os"
	"strings"
)

var (
	RuoyiVue3Gen = &RuoyiVue3GenType{}
)

type (
	RuoyiVue3GenType struct {
	}
)

func (v *RuoyiVue3GenType) GenerateRuoyiCodes(config *GenConfig, model *DbModel) {
	if config.Gen.GenRuoyiVue3 == nil || !config.Gen.GenRuoyiVue3.Enable || len(config.Gen.GenRuoyiVue3.ProjectRoot) < 1 {
		return
	}

	v.GenMenus(config, model)
	v.GenApiJs(config, model)
	v.GenVueViews(config, model)
}

func (v *RuoyiVue3GenType) GenMenus(config *GenConfig, model *DbModel) {

}

func (v *RuoyiVue3GenType) GenApiJs(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateRuoyiVue3ApiKey)
	// deleted column:
	if model.HasDeletedColumn() {
		templateCodes += GetTemplate(TemplateRuoyiVue3ApiDeletedColumnKey)
	}

	templateCodes = strings.Replace(templateCodes, "{TableModelName}", model.StructName, -1)
	templateCodes = strings.Replace(templateCodes, "{TableModelNameLowCase}", strings.ToLower(model.StructName), -1)

	// queryParams := v.getVueQueryParameters(model)
	// templateCodes = strings.Replace(templateCodes, "// {其它查询参数}", queryParams, -1)

	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/api"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/"+model.PrivatePropertyName+".js", templateCodes)
}

func (v *RuoyiVue3GenType) GenVueViews(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateRuoyiVue3ViewKey)

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
	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/views/gen"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/"+fileName+".vue", templateCodes)
}

func (v *RuoyiVue3GenType) GenerateVueRouter(config *GenConfig, models []*DbModel) {
	if config.Gen.GenRuoyiVue3 == nil || !config.Gen.GenRuoyiVue3.Enable || len(config.Gen.GenRuoyiVue3.ProjectRoot) < 1 {
		return
	}

	templateCodes := GetTemplate(TemplateRuoyiVue3MenuKey)

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

	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/router/modules"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/goCodeGen.js", templateCodes)
}

func (v *RuoyiVue3GenType) getVueQueryParameters(model *DbModel) string {
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
