package model

import (
	"time"
)

type GoodInfo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type GoodCreateParams struct {
	ProjectID int
	Info      GoodInfo
}

type GoodUpdateParams struct {
	ID        int
	ProjectID int
	Info      GoodInfo
}

type GoodDRemoveParams struct {
	ID        int
	ProjectID int
}

type GoodRemove struct {
	ID        int
	ProjectID int
	Removed   bool
}

type GoodListParams struct {
	ProjectID int
	Limit     int
	Offset    int
}

type GoodReprioritizeParams struct {
	ID          int
	ProjectID   int
	NewPriority int
}

type Prioritise struct {
	ID       int
	Priority int
}

type GoodsPrioritize struct {
	Priorities []Prioritise
}

type Meta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}

type GoodsList struct {
	MetaInfo Meta   `json:"meta"`
	Goods    []Good `json:"goods"`
}

type Good struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"projectId"`
	Info      GoodInfo  `json:"info"`
	Priority  int       `json:"priority"`
	Removed   bool      `json:"removed"`
	CreatedAt time.Time `json:"createdAt"`
}
