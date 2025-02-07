package handlers

import (
	"context"
	"github.com/glennliao/apijson-go/action"
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/consts"
	"github.com/glennliao/apijson-go/query"
	"github.com/gogf/gf/v2/frame/g"
)

func Get(ctx context.Context, req g.Map) (res g.Map, err error) {
	q := query.New(ctx, req)
	q.AccessVerify = config.AccessVerify
	q.AccessCondition = config.AccessConditionFunc
	return q.Result()
}

func Head(ctx context.Context, req g.Map) (res g.Map, err error) {
	return nil, err
}

func Post(ctx context.Context, req g.Map) (res g.Map, err error) {
	act := action.New(ctx, consts.MethodPost, req)
	return act.Result()
}

func Put(ctx context.Context, req g.Map) (res g.Map, err error) {
	act := action.New(ctx, consts.MethodPut, req)
	return act.Result()
}

func Delete(ctx context.Context, req g.Map) (res g.Map, err error) {
	act := action.New(ctx, consts.MethodDelete, req)
	return act.Result()
}
