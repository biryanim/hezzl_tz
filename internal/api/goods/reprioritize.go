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

func (i *Implementation) Reprioritize(c *gin.Context) {
	var (
		reprioritizeGoodReq dto.GoodReprioritizeReq
		err                 error
	)
	reprioritizeGoodReq.ID, err = strconv.Atoi(c.Query("id"))
	if err != nil || reprioritizeGoodReq.ID <= 0 {
		fmt.Println(err)
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	reprioritizeGoodReq.ProjectID, err = strconv.Atoi(c.Query("projectId"))
	if err != nil || reprioritizeGoodReq.ProjectID <= 0 {
		fmt.Println(err, "wwwwwww")
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	if err = c.ShouldBindJSON(&reprioritizeGoodReq); err != nil {
		fmt.Println(err, "eeeeeeeeewwww")
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	resp, err := i.goodsService.Reprioritize(c.Request.Context(), converter.FromReprioritizeReq(&reprioritizeGoodReq))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, converter.ToReprioritizeResp(resp))
}
