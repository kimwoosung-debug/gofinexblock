package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
)

type queue struct {
	service Service

	limitAsk  chan *types.ErrReceiveContext[*grpc_order.Order] // limitAsk is channel for limit ask order
	marketAsk chan *types.ErrReceiveContext[*grpc_order.Order] // marketAsk is channel for market ask order
	limitBid  chan *types.ErrReceiveContext[*grpc_order.Order] // limitBid is channel for limit bid order
	marketBid chan *types.ErrReceiveContext[*grpc_order.Order] // marketBid is channel for market bid order

	cancel chan *types.ResultReceiveContext[string, *grpc_order.Order] // cancel is channel for cancel order
}

func newQueue(cluster *redis.ClusterClient) *queue {
	return &queue{
		service:   NewService(cluster),
		limitAsk:  make(chan *types.ErrReceiveContext[*grpc_order.Order]),
		marketAsk: make(chan *types.ErrReceiveContext[*grpc_order.Order]),
		limitBid:  make(chan *types.ErrReceiveContext[*grpc_order.Order]),
		marketBid: make(chan *types.ErrReceiveContext[*grpc_order.Order]),
		cancel:    make(chan *types.ResultReceiveContext[string, *grpc_order.Order]),
	}
}

func (q *queue) Subscribe() {

	for {
		select {
		case ctx := <-q.limitAsk:
			ctx.Tunnel <- q.service.LimitAsk(ctx.Value)
		case ctx := <-q.limitBid:
			ctx.Tunnel <- q.service.LimitBid(ctx.Value)
		case ctx := <-q.marketAsk:
			ctx.Tunnel <- q.service.MarketAsk(ctx.Value)
		case ctx := <-q.marketBid:
			ctx.Tunnel <- q.service.MarketBid(ctx.Value)
		case ctx := <-q.cancel:
			order, _ := q.service.CancelOrder(ctx.Value)
			ctx.Tunnel <- order
		}
	}
}

func (q *queue) LimitAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	q.limitAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (q *queue) LimitBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	q.limitBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (q *queue) MarketAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	q.marketAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (q *queue) MarketBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	q.marketBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (q *queue) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	ctx := &types.ResultReceiveContext[string, *grpc_order.Order]{
		Tunnel: make(chan *grpc_order.Order),
		Value:  uuid,
	}
	q.cancel <- ctx
	order = <-ctx.Tunnel
	if order == nil {
		return nil, ErrOrderCancelFailed
	}
	return order, nil
}

func (q *queue) BidOrder() (bids []*grpc_order.Order, err error) {
	return q.service.BidOrder()
}

func (q *queue) AskOrder() (asks []*grpc_order.Order, err error) {
	return q.service.AskOrder()
}
