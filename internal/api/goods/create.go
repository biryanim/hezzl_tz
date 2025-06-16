package goods

import (
	"fmt"
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/biryanim/hezzl_tz/internal/converter"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Create(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Query("projectId"))
	if err != nil || projectID <= 0 {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	var (
		createReq dto.GoodCreateReq
	)
	if err = c.ShouldBindJSON(&createReq.Info); err != nil {
		fmt.Println("DDDDDDDDDDDDDD")
		c.Error(apperrors.ErrInvalidInput)
		return
	}
	createReq.ProjectID = projectID

	goods, err := i.goodsService.Create(c.Request.Context(), converter.FromGoodCreateReq(&createReq))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToGoodsResponse(goods))
}
