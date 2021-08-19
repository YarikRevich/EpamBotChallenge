package bottom

import (
	"battlecity_test/game"
	"battlecity_test/solver/tickets"
)

type BottomWayEstimator struct {
	t *tickets.Ticket
	way game.Point
}

func (bwe *BottomWayEstimator) Estimate() int{
	bwe.t.EstimateAuto(bwe.way)
	return bwe.t.GetAmount()
}

func (bwe *BottomWayEstimator) GetTicket() *tickets.Ticket{
	return bwe.t
}

func New(t *tickets.Ticket)*BottomWayEstimator{
	return &BottomWayEstimator{
		t: t,
	}
}