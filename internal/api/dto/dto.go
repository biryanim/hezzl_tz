package dto

import "time"

type GoodInfo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

type GoodCreateReq struct {
	Info      GoodInfo
	ProjectID int
}

type GoodUpdateReq struct {
	ID        int
	ProjectID int
	Info      GoodInfo
}

type GoodDeleteReq struct {
	ID        int
	ProjectID int
}

type GoodsListReq struct {
	Limit  int
	Offset int
}

type GoodReprioritizeReq struct {
	ID          int `json:"-"`
	ProjectID   int `json:"-"`
	NewPriority int `json:"newPriority" validate:"required"`
}

type GoodRemoveResp struct {
	ID        int  `json:"id"`
	ProjectID int  `json:"projectId"`
	Removed   bool `json:"removed"`
}

type Good struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Meta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}

type GoodsList struct {
	Meta  Meta   `json:"meta"`
	Goods []Good `json:"goods"`
}

type Prioritise struct {
	ID       int `json:"id"`
	Priority int `json:"priority"`
}

type GoodsPrioritize struct {
	Prioritise []Prioritise `json:"prioritise"`
}
