package codegen

import (
	"fmt"
	"io/ioutil"

	"gorm.io/gorm"
)

func runDbDmls(db *gorm.DB, config *GenConfig) {
	fileList, err := ioutil.ReadDir(config.Gen.DmlFolder)
	if err != nil {
		fmt.Printf("List files faild: "+config.Gen.DmlFolder+",erro:%s", err)
		return
	}

	for _, file := range fileList {
		if !file.IsDir() {
			filePath := config.Gen.DmlFolder + "/" + file.Name()
			ret, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Read file failed: "+filePath+",erro:%s", err)
				continue
			}
			tx := db.Exec(string(ret))
			if tx.Error != nil {
				fmt.Printf("Execute sql faild: "+filePath+",erro:%s", err)
			}
		}
	}

}
