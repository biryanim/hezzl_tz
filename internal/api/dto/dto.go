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
	ID          int
	ProjectID   int
	NewPriority int `json:"newPriority" validate:"required"`
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
