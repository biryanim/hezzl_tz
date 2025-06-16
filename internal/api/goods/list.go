package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/biryanim/hezzl_tz/internal/converter"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	listReq := &dto.GoodsListReq{
		Limit:  limit,
		Offset: offset,
	}

	resp, err := i.goodsService.List(c.Request.Context(), converter.FromGoodsListReq(listReq))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, converter.ToGoodsListResponse(resp))
}
