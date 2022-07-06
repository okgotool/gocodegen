package dal

import "github.com/okgotool/gocodegen/example/dal/model"

func AutoMigrateTables() {
	CreateDbConn("SysRole").AutoMigrate(&model.SysRole{})
	CreateDbConn("TestCat").AutoMigrate(&model.TestCat{})
}
