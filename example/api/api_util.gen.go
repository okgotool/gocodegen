package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPagerValuesFromQuery(g *gin.Context) (int, int, map[string]string) {

	pageStr := g.Query("page")
	pageSizeStr := g.Query("pageSize")

	errMsgs := map[string]string{}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		errMsgs["page"] = err.Error()
	} else if page < 1 {
		page = 1
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		errMsgs["pageSize"] = err.Error()
	} else if pageSize < 1 {
		pageSize = 10
	}

	return int(page), int(pageSize), errMsgs
}
