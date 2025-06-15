package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Reprioritize(c *gin.Context) {
	var (
		reprioritizeGoodReq dto.GoodReprioritizeReq
		err                 error
	)
	reprioritizeGoodReq.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	reprioritizeGoodReq.ProjectID, err = strconv.Atoi(c.Param("projectID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project id"})
		return
	}

	if err = c.BindJSON(&reprioritizeGoodReq.NewPriority); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
