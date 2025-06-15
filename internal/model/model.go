package model

import (
	"time"
)

type GoodInfo struct {
	Name        string
	Description string
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
	Total   int
	Removed int
	Limit   int
	Offset  int
}

type GoodsList struct {
	MetaInfo Meta
	Goods    []Good
}

type Good struct {
	ID        int
	ProjectID int
	Info      GoodInfo
	Priority  int
	Removed   bool
	CreatedAt time.Time
}
