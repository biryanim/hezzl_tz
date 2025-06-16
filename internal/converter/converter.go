package converter

import (
	"github.com/biryanim/hezzl_tz/internal/api/dto"
	"github.com/biryanim/hezzl_tz/internal/model"
)

func FromGoodCreateReq(createReq *dto.GoodCreateReq) *model.GoodCreateParams {
	return &model.GoodCreateParams{
		ProjectID: createReq.ProjectID,
		Info: model.GoodInfo{
			Name:        createReq.Info.Name,
			Description: createReq.Info.Description,
		},
	}
}

func ToGoodsResponse(goods *model.Good) *dto.Good {
	return &dto.Good{
		ID:          goods.ID,
		ProjectID:   goods.ProjectID,
		Name:        goods.Info.Name,
		Description: goods.Info.Description,
		Priority:    goods.Priority,
		Removed:     goods.Removed,
		CreatedAt:   goods.CreatedAt,
	}
}

func FromGoodUpdateReq(updateReq *dto.GoodUpdateReq) *model.GoodUpdateParams {
	return &model.GoodUpdateParams{
		ID:        updateReq.ID,
		ProjectID: updateReq.ProjectID,
		Info: model.GoodInfo{
			Name:        updateReq.Info.Name,
			Description: updateReq.Info.Description,
		},
	}
}

func FromGoodRemoveReq(removeReq *dto.GoodDeleteReq) *model.GoodDRemoveParams {
	return &model.GoodDRemoveParams{
		ID:        removeReq.ID,
		ProjectID: removeReq.ProjectID,
	}
}

func ToGoodRemoveResponse(removingResp *model.GoodRemove) *dto.GoodRemoveResp {
	return &dto.GoodRemoveResp{
		ID:        removingResp.ID,
		ProjectID: removingResp.ProjectID,
		Removed:   removingResp.Removed,
	}
}

func FromGoodsListReq(req *dto.GoodsListReq) *model.GoodListParams {
	return &model.GoodListParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
}

func ToGoodsListResponse(goods *model.GoodsList) *dto.GoodsList {
	return &dto.GoodsList{
		Meta:  toMetaResp(goods.MetaInfo),
		Goods: toGoodsListResp(goods.Goods),
	}
}

func toMetaResp(meta model.Meta) dto.Meta {
	return dto.Meta{
		Total:   meta.Total,
		Removed: meta.Removed,
		Limit:   meta.Limit,
		Offset:  meta.Offset,
	}
}

func toGoodsListResp(goods []model.Good) []dto.Good {
	var resp []dto.Good
	for _, good := range goods {
		resp = append(resp, dto.Good{
			ID:          good.ID,
			ProjectID:   good.ProjectID,
			Name:        good.Info.Name,
			Description: good.Info.Description,
			Priority:    good.Priority,
			Removed:     good.Removed,
			CreatedAt:   good.CreatedAt,
		})
	}
	return resp
}

func FromReprioritizeReq(req *dto.GoodReprioritizeReq) *model.GoodReprioritizeParams {
	return &model.GoodReprioritizeParams{
		ProjectID:   req.ProjectID,
		ID:          req.ID,
		NewPriority: req.NewPriority,
	}
}

func ToReprioritizeResp(resp *model.GoodsPrioritize) *dto.GoodsPrioritize {
	var r dto.GoodsPrioritize
	for _, v := range resp.Priorities {
		r.Prioritise = append(r.Prioritise, dto.Prioritise{
			ID:       v.ID,
			Priority: v.Priority,
		})
	}

	return &r
}
