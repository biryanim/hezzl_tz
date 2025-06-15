package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	projectID, err := strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	var (
		goodUpdateReq dto.GoodUpdateReq
	)
	if err = c.BindJSON(&goodUpdateReq.Info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	goodUpdateReq.ID = id
	goodUpdateReq.ProjectID = projectID
}
