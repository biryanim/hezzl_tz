package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
		return
	}

	listReq := &dto.GoodsListReq{
		Limit:  limit,
		Offset: offset,
	}

}
