package goods

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/biryanim/hezzl_tz/internal/converter"
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (i *Implementation) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	projectID, err := strconv.Atoi(c.Query("projectId"))
	if err != nil {
		c.Error(apperrors.ErrInvalidInput)
		return
	}

	var removeReq dto.GoodDeleteReq
	removeReq.ID = id
	removeReq.ProjectID = projectID

	resp, err := i.goodsService.Remove(c.Request.Context(), converter.FromGoodRemoveReq(&removeReq))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, converter.ToGoodRemoveResponse(resp))
}
