package codegen

import (
	"fmt"

	"github.com/okgotool/gocodegen/app"
)

const (
	TemplateResponseModelKey       = "responseModel"
	TemplateDatasourceKey          = "datasource"
	TemplateBizKey                 = "biz"
	TemplateBizDeletedKey          = "bizDeleted"
	TemplateApiKey                 = "api"
	TemplateApirouterKey           = "apirouter"
	TemplateApirouterDeletedKey    = "apiRouterDeleted"
	TemplateApiQueryAllKey         = "apiQueryAll"
	TemplateApiQueryByConditionKey = "apiQueryByCondition"
	TemplateApiCreateBatchKey      = "apiCreateBatch"
	TemplateApiCreateKey           = "apiCreate"
	TemplateApiUpdateBatchKey      = "apiUpdateBatch"
	TemplateApiUpdateKey           = "apiUpdate"
	TemplateApiRemoveKey           = "apiRemove"
	TemplateApiDeletedKey          = "apiDeleted"
	TemplateVueApiKey              = "vueApi"
	TemplateVueApiDeletedColumnKey = "vueApiDeletedColumn"
	TemplateVueViewKey             = "vueView"
	TemplateVueRouterKey           = "vueRouter"
)

var (
	TemplateFiles map[string]string = map[string]string{
		TemplateResponseModelKey:       app.RUNTIME_PATH + "/template/model/response/model_list_response.template",
		TemplateDatasourceKey:          app.RUNTIME_PATH + "/template/dal/datasources.template",
		TemplateBizKey:                 app.RUNTIME_PATH + "/template/biz/biz.template",
		TemplateBizDeletedKey:          app.RUNTIME_PATH + "/template/biz/biz_deleted_column.template",
		TemplateApiKey:                 app.RUNTIME_PATH + "/template/api/api.template",
		TemplateApirouterKey:           app.RUNTIME_PATH + "/template/api/api_router.template",
		TemplateApirouterDeletedKey:    app.RUNTIME_PATH + "/template/api/api_router_deleted.template",
		TemplateApiQueryAllKey:         app.RUNTIME_PATH + "/template/api/api_query_all.template",
		TemplateApiQueryByConditionKey: app.RUNTIME_PATH + "/template/api/api_query_by_condition.template",
		TemplateApiCreateBatchKey:      app.RUNTIME_PATH + "/template/api/api_create_batch.template",
		TemplateApiCreateKey:           app.RUNTIME_PATH + "/template/api/api_create.template",
		TemplateApiUpdateBatchKey:      app.RUNTIME_PATH + "/template/api/api_update_batch.template",
		TemplateApiUpdateKey:           app.RUNTIME_PATH + "/template/api/api_update.template",
		TemplateApiRemoveKey:           app.RUNTIME_PATH + "/template/api/api_remove.template",
		TemplateApiDeletedKey:          app.RUNTIME_PATH + "/template/api/api_deleted_column.template",
		TemplateVueApiKey:              app.RUNTIME_PATH + "/template/vue/vue_api.template",
		TemplateVueApiDeletedColumnKey: app.RUNTIME_PATH + "/template/vue/vue_api_deleted_column.template",
		TemplateVueViewKey:             app.RUNTIME_PATH + "/template/vue/vue_view.template",
		TemplateVueRouterKey:           app.RUNTIME_PATH + "/template/vue/vue_router.template",
	}

	templateContents map[string]string = map[string]string{}
)

func GetTemplate(key string) string {
	// lowKey := strings.ToLower(key)
	if str, ok := templateContents[key]; ok && len(str) > 0 {
		return str
	}

	if templateFile, ok := TemplateFiles[key]; !ok || len(templateFile) < 1 {
		fmt.Println("Not set template file:" + templateFile)
		return ""
	} else {
		content := ReadFile(templateFile)
		if len(content) > 0 {
			templateContents[key] = content
		}
		return content
	}
}
