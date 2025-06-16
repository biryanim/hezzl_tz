package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/biryanim/hezzl_tz/internal/converter"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	projectID, err := strconv.Atoi(c.Query("projectId"))
	if err != nil || projectID < 1 {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	var (
		goodUpdateReq dto.GoodUpdateReq
	)
	if err = c.ShouldBindJSON(&goodUpdateReq.Info); err != nil {
		c.Error(apperrors.ErrInvalidInput)
		return
	}
	goodUpdateReq.ID = id
	goodUpdateReq.ProjectID = projectID

	resp, err := i.goodsService.Update(c.Request.Context(), converter.FromGoodUpdateReq(&goodUpdateReq))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, converter.ToGoodsResponse(resp))
}
