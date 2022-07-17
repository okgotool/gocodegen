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

func (v *RuoyiVue3GenType) GenerateRuoyiVue3Codes(config *GenConfig, model *DbModel) {
	if config.Gen.GenRuoyiVue3 == nil || !config.Gen.GenRuoyiVue3.Enable || len(config.Gen.GenRuoyiVue3.ProjectRoot) < 1 {
		return
	}

	// v.GenMenus(config, model)
	v.GenApiJs(config, model)
	v.GenVueViews(config, model)
}

func (v *RuoyiVue3GenType) GenTopMenusSql(config *GenConfig) {
	if len(config.Gen.GenRuoyiVue3.TopMenus) < 1 {
		return
	}

	insertColumnsSql := "insert into sys_menu (menu_name, parent_id, order_num, path, component, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark)\n"

	sqlCodes := ""
	for _, menu := range config.Gen.GenRuoyiVue3.TopMenus {
		sqlCodes += "--- Top menu: " + menu.Name + "\n"
		sqlCodes += insertColumnsSql
		sqlCodes += " values('" + menu.Name + "', '0', '" + menu.OrderNum + "', '" + menu.Path + "', null, 1, 0, 'M', '0', '0', '', '" + menu.Icon + "', 'admin', sysdate(), '', null, '" + menu.Name + "');\n\n"
	}

	vueFolder := config.Gen.OutputRoot + "/db_migrate/menu"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/ruoyi_top_menus.gen.sql", sqlCodes)
}

func (v *RuoyiVue3GenType) GenChildMenusSql(config *GenConfig, models []*DbModel) {
	insertColumnsSql := "insert into sys_menu (menu_name, parent_id, order_num, path, component, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark)\n"

	sqlCodes := ""
	for _, table := range config.Gen.GenTables {
		if len(table.ParentMenu) < 1 {
			continue
		}

		menuName := table.MenuName

		for _, model := range models {
			if !strings.EqualFold(table.TableName, model.TableName) {
				continue
			}

			if len(menuName) < 1 {
				menuName = model.StructName
			}

			menuPath := model.PathName

			sqlCodes += "\n"
			sqlCodes += "--- Menu: " + table.ParentMenu + "/" + menuName + "\n"
			sqlCodes += "SELECT menu_id,path into @" + menuPath + "_parent_menu_id,@" + menuPath + "_parent_menu_path FROM sys_menu where parent_id=0 and menu_name='" + table.ParentMenu + "';\n"
			sqlCodes += insertColumnsSql
			// sqlCodes += " values('" + menuName + "', @" + menuPath + "_parent_menu_id, '1', '" + menuPath + "', concat(@" + menuPath + "_parent_menu_path,'/" + menuPath + "/index'), 1, 0, 'C', '0', '0', '', 'guide', 'admin', sysdate(), '', null, '" + menuName + "');\n"
			sqlCodes += " values('" + menuName + "', @" + menuPath + "_parent_menu_id, '1', '" + menuPath + "', '" + menuPath + "/index', 1, 0, 'C', '0', '0', '', 'guide', 'admin', sysdate(), '', null, '" + menuName + "');\n"

			break
		}
	}
	if len(sqlCodes) < 5 { // not set child menus
		return
	}

	vueFolder := config.Gen.OutputRoot + "/db_migrate/menu"
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/ruoyi_child_menus.gen.sql", sqlCodes)
}

// func (v *RuoyiVue3GenType) GenChildMenusSql(config *GenConfig, models []*DbModel) {
// 	tempSql := "insert into sys_menu (menu_name, parent_id, order_num, path, component, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark)"

// 	sqlCodes := ""
// 	for _, menu := range config.Gen.GenTables {
// 		sqlCodes += tempSql
// 		sqlCodes += " values('" + menu.Name + "', '0', '" + menu.OrderNum + "', '" + menu.Path + "', null, 1, 0, 'M', '0', '0', '', '" + menu.Icon + "', 'admin', sysdate(), '', null, '" + menu.Name + "');\n\n"
// 	}

// 	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/sql"
// 	os.MkdirAll(vueFolder, 0666)
// 	WriteFile(vueFolder+"/child_menus.gen.sql", sqlCodes)
// }

// func (v *RuoyiVue3GenType) GenMenus(config *GenConfig, model *DbModel) {

// }

func (v *RuoyiVue3GenType) GenApiJs(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateRuoyiVue3ApiKey)
	// deleted column:
	if model.HasDeletedColumn() {
		templateCodes += GetTemplate(TemplateRuoyiVue3ApiDeletedColumnKey)
	}

	templateCodes = strings.Replace(templateCodes, "{TableModelName}", model.StructName, -1)
	templateCodes = strings.Replace(templateCodes, "{TableModelNameLowCase}", model.PathName, -1)

	// queryParams := v.getVueQueryParameters(model)
	// templateCodes = strings.Replace(templateCodes, "// {其它查询参数}", queryParams, -1)

	pathName := strings.ToLower(model.StructName)
	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/api/" + pathName
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/"+pathName+".js", templateCodes)
}

func (v *RuoyiVue3GenType) GenVueViews(config *GenConfig, model *DbModel) {
	templateCodes := GetTemplate(TemplateRuoyiVue3ViewKey)

	templateCodes = strings.Replace(templateCodes, "{TableModelName}", model.StructName, -1)
	templateCodes = strings.Replace(templateCodes, "{LowerFirstCharTableModelName}", model.PrivatePropertyName, -1)
	templateCodes = strings.Replace(templateCodes, "{TableModelNameLowCase}", model.PathName, -1)

	primaryKey := model.PrimaryKeyPropertyName
	if strings.EqualFold(primaryKey, "ID") {
		primaryKey = "id"
	} else {
		primaryKey = strings.ToLower(string(primaryKey[0])) + primaryKey[1:]
	}
	templateCodes = strings.Replace(templateCodes, "{PrimaryKeyPropertyName}", primaryKey, -1)

	// search view:
	searchCodes := ""
	for _, modelField := range model.Fields {
		if modelField.IgnoreGenerateResponseModel || len(modelField.ModelFieldName) < 1 {
			continue
		}

		codes := ""
		if strings.Contains(modelField.ModelFieldType, "time") {

			codes += "    <el-form-item label=\"" + modelField.ModelFieldName + "\" style=\"width: 308px\">\n"
			codes += "    <el-date-picker\n"
			codes += "       v-model=\"dateRange\"\n"
			codes += "       value-format=\"YYYY-MM-DD\"\n"
			codes += "       type=\"daterange\"\n"
			codes += "       range-separator=\"-\"\n"
			codes += "       start-placeholder=\"开始时间\"\n"
			codes += "       end-placeholder=\"结束时间\"\n"
			codes += "    ></el-date-picker>\n"
			codes += "    </el-form-item>\n"
		} else {
			codes += "      <el-form-item label=\"" + modelField.ModelFieldName + "\" prop=\"" + modelField.PrivateModelFieldName + "\">\n"
			codes += "      <el-input\n"
			codes += "         v-model=\"queryParams." + modelField.PrivateModelFieldName + "\"\n"
			codes += "         placeholder=\"请输入" + modelField.ModelFieldName + "\"\n"
			codes += "         clearable\n"
			codes += "         style=\"width: 240px\"\n"
			codes += "         @keyup.enter=\"handleQuery\"\n"
			codes += "      />\n"
			codes += "    </el-form-item>\n"
		}

		// codes := "      <el-table-column align=\"left\" sortable label=\"" + modelField.ModelFieldName + "\">\n"
		// codes += "        <template slot-scope=\"{ row }\">\n"
		// codes += "          <span>{{ row." + modelField.ColumnName + " }}</span>\n"
		// codes += "        </template>\n"
		// codes += "      </el-table-column>\n\n"

		searchCodes += codes
	}
	templateCodes = strings.Replace(templateCodes, "<!-- {ModelSearchFormItems} -->", searchCodes, -1)

	// list table view:
	columnsCodes := ""
	for _, modelField := range model.Fields {
		if modelField.IgnoreGenerateResponseModel || len(modelField.ModelFieldName) < 1 {
			continue
		}

		codes := ""
		if strings.Contains(modelField.ModelFieldType, "time") {
			codes += "    <el-table-column label=\"" + modelField.ModelFieldName + "\" align=\"center\" prop=\"" + modelField.PrivateModelFieldName + "\" width=\"180\">\n"
			codes += "      <template #default=\"scope\">\n"
			codes += "         <span>{{ parseTime(scope.row." + modelField.PrivateModelFieldName + ") }}</span>\n"
			codes += "      </template>\n"
			codes += "      </el-table-column>\n"
		} else {
			codes += "      <el-table-column label=\"" + modelField.ModelFieldName + "\" align=\"center\" prop=\"" + modelField.PrivateModelFieldName + "\" />\n"
		}

		// codes := "      <el-table-column align=\"left\" sortable label=\"" + modelField.ModelFieldName + "\">\n"
		// codes += "        <template slot-scope=\"{ row }\">\n"
		// codes += "          <span>{{ row." + modelField.ColumnName + " }}</span>\n"
		// codes += "        </template>\n"
		// codes += "      </el-table-column>\n\n"

		columnsCodes += codes
	}
	templateCodes = strings.Replace(templateCodes, "<!-- {ModelTableColumns} -->", columnsCodes, -1)

	// create/edit form view:
	// createFormItems := ""
	editFormItems := ""
	// createFormItemNames := ""
	queryParamNames := ""
	for _, modelField := range model.Fields {
		if modelField.IgnoreGenerateResponseModel || len(modelField.ModelFieldName) < 1 {
			continue
		}
		codes := ""

		if !strings.EqualFold(modelField.ModelFieldName, "ID") {
			if strings.Contains(modelField.ModelFieldType, "time") {
				// 	codes += "			<el-form-item label=\"" + modelField.ModelFieldName + "\">\n"
				// 	codes += "			<el-date-picker\n"
				// 	codes += "			  type=\"date\"\n"
				// 	codes += "			  placeholder=\"选择日期\"\n"
				// 	codes += "			  v-model=\"form." + modelField.ColumnName + "\"\n"
				// 	codes += "			  style=\"width: 100%\"\n"
				// 	codes += "			></el-date-picker>\n"
				// 	codes += "		  </el-form-item>\n"
			} else {
				// 	valueMode := ""
				// 	if strings.HasPrefix(modelField.ModelFieldType, "int") || strings.HasPrefix(modelField.ModelFieldType, "uint") {
				// 		valueMode = ".number"
				// }

				codes += "        <el-form-item label=\"" + modelField.ModelFieldName + "\" prop=\"" + modelField.PrivateModelFieldName + "\">\n"
				codes += "          <el-input v-model=\"form." + modelField.PrivateModelFieldName + "\" placeholder=\"请输入" + modelField.ModelFieldName + "\" />\n"
				codes += "        </el-form-item>\n"

				// 	codes += "        <el-form-item label=\"" + modelField.ModelFieldName + "\">\n"
				// 	codes += "          <el-input v-model" + valueMode + "=\"form." + modelField.ColumnName + "\"></el-input>\n"
				// 	codes += "        </el-form-item>\n"
			}

			// createFormItems += codes
			editFormItems += codes
		}

		// if strings.HasPrefix(modelField.ModelFieldType, "int") || strings.HasPrefix(modelField.ModelFieldType, "uint") {
		// 	createFormItemNames += "        " + modelField.ColumnName + ": 0,\n"
		// } else {
		// 	createFormItemNames += "        " + modelField.ColumnName + ": '',\n"
		// }
		queryParamNames += "        " + modelField.PrivateModelFieldName + ": undefined,\n"
	}
	// templateCodes = strings.Replace(templateCodes, "// {ModelCreateFormItems}", createFormItems, -1)
	templateCodes = strings.Replace(templateCodes, "<!-- {ModelEditFormItems} -->", editFormItems, -1)
	templateCodes = strings.Replace(templateCodes, "// {ModelQueryParamNames}", queryParamNames, -1)
	// templateCodes = strings.Replace(templateCodes, "// {ModelCreateFormItemNames}", createFormItemNames, -1)

	fileName := strings.ToLower(model.StructName)
	vueFolder := config.Gen.GenRuoyiVue3.ProjectRoot + "/src/views/" + fileName
	os.MkdirAll(vueFolder, 0666)
	WriteFile(vueFolder+"/index.vue", templateCodes)
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
