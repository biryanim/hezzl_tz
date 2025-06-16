package goods

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/model"
	"log"
)

func (s *serv) Reprioritize(ctx context.Context, goodReprioritizingParams *model.GoodReprioritizeParams) (*model.GoodsPrioritize, error) {
	var goods []*model.Good
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var err error
		goods, err = s.goodsRepository.Reprioritize(ctx, goodReprioritizingParams)
		return err
	})
	if err != nil {
		return nil, err
	}

	if err = s.cache.DeleteByPattern(ctx, "goods:list:*"); err != nil {
		log.Printf("delete goods cache err: %v", err)
	}

	if goods != nil && len(goods) > 0 {
		for _, good := range goods {
			goodToLog := good
			go func() {
				s.publishLogEvent(ctx, goodToLog)
			}()
		}
	}

	var (
		resp       *model.GoodsPrioritize
		priorities []model.Prioritise
	)
	for _, good := range goods {
		p := model.Prioritise{
			ID:       good.ID,
			Priority: good.Priority,
		}
		priorities = append(priorities, p)
	}
	resp = &model.GoodsPrioritize{
		Priorities: priorities,
	}

	return resp, nil
}
