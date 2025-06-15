package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Create(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var (
		createReq dto.GoodCreateReq
	)
	if err = c.BindJSON(&createReq.Info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createReq.ProjectID = projectID

}
