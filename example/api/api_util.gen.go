package api

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getPagerValuesFromQuery(g *gin.Context) (int, int, map[string]string) {

	pageStr := g.Query("page")
	pageSizeStr := g.Query("pageSize")

	errMsgs := map[string]string{}

	page, err := parseInt64(pageStr)
	if err != nil {
		errMsgs["page"] = err.Error()
	} else if page < 1 {
		page = 1
	}

	pageSize, err := parseInt64(pageSizeStr)
	if err != nil {
		errMsgs["pageSize"] = err.Error()
	} else if pageSize < 1 {
		pageSize = 10
	}

	return int(page), int(pageSize), errMsgs
}

func parseInt64(str string) (int64, error) {

	s := strings.Replace(str, "\"", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "'", "", -1)

	if len(s) < 1 {
		return 0, nil
	}

	value, err := strconv.ParseInt(s, 10, 64)
	return value, err
}

func parseUint64(str string) (uint64, error) {

	s := strings.Replace(str, "\"", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "'", "", -1)

	if len(s) < 1 {
		return 0, nil
	}

	value, err := strconv.ParseUint(s, 10, 64)
	return value, err
}
