package codegen

import (
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	QueryCtx = context.Background()
// )

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}

// func RemoveSoftDeletedCodes(hasDeletedColumn bool, codes string) string {
// 	if hasDeletedColumn {
// 		return codes
// 	}

// 	indexHasDeletedStart := strings.Index(codes, "//// hasDeletedColumn start")
// 	indexHasDeletedEnd := strings.Index(codes, "//// hasDeletedColumn end") + 25

// 	return codes[0:indexHasDeletedStart] + codes[indexHasDeletedEnd:]
// }

func CopyFile(srcFile string, dstFile string, isReplace bool) {
	// copy response common file:
	if !FileExist(dstFile) || isReplace {
		bytes, err := ioutil.ReadFile(srcFile)
		if err != nil {
			panic("Failed to read file: " + srcFile)
		}
		WriteFile(dstFile, string(bytes))
	}
}

func ReadFile(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Read file failed: "+file+",erro:%s", err)
		return ""
	}
	return string(bytes)
}

func WriteFile(fileName string, content string) {
	if err := ioutil.WriteFile(fileName, []byte(content), 0666); err != nil {
		fmt.Println("Write file failed: " + fileName + ", " + err.Error())
	} else {
		fmt.Println(" -- Write file success: " + fileName)
	}
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}
