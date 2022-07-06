package codegen

import (
	"os"

	"github.com/okgotool/gocodegen/app"
)

func GenerateFolders(conf *GenConfig) {
	rootPath := conf.Gen.OutputRoot

	// api folder:
	os.Mkdir(rootPath+"/api", 0666)

	// biz folder:
	os.Mkdir(rootPath+"/biz", 0666)

	// dao folder:
	os.Mkdir(rootPath+"/dal", 0666)
	os.Mkdir(rootPath+"/dal/model", 0666)
	os.Mkdir(rootPath+"/dal/query", 0666)

	// model folder:
	os.Mkdir(rootPath+"/model", 0666)
	os.Mkdir(rootPath+"/model/db", 0666)
	os.Mkdir(rootPath+"/model/request", 0666)
	os.Mkdir(rootPath+"/model/response", 0666)

	// config folder:
	os.Mkdir(rootPath+"/config", 0666)

	// gen files:
	generateCommonFiles(conf)
}

func generateCommonFiles(config *GenConfig) {
	// model:

	// dal utils:
	CopyFile(app.RUNTIME_PATH+"/template/dal/dal_util.go", config.Gen.OutputRoot+"/dal/dal_util.gen.go", config.Gen.GenDal.OverWrite)

	// biz:
	CopyFile("./template/biz/login_util.go", config.Gen.OutputRoot+"/biz/login_util.gen.go", config.Gen.GenDal.OverWrite)

	// api utils:
	CopyFile(app.RUNTIME_PATH+"/template/model/response/common.go", config.Gen.OutputRoot+"/model/response/common.gen.go", config.Gen.GenApi.OverWrite)
	CopyFile(app.RUNTIME_PATH+"/template/model/response/status_code.go", config.Gen.OutputRoot+"/model/response/status_code.gen.go", config.Gen.GenApi.OverWrite)
	CopyFile(app.RUNTIME_PATH+"/template/api/api_util.go", config.Gen.OutputRoot+"/api/api_util.gen.go", config.Gen.GenApi.OverWrite)

	// vue:
	CopyFile("./template/model/request/vue_user_login.go", config.Gen.OutputRoot+"/model/request/vue_user_login.gen.go", config.Gen.GenApi.OverWrite)
	CopyFile("./template/model/response/vue_user_login.go", config.Gen.OutputRoot+"/model/response/vue_user_login.gen.go", config.Gen.GenApi.OverWrite)

}
