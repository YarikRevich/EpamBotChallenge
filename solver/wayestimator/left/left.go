package left

import (
	"battlecity_test/game"
	"battlecity_test/solver/tickets"
)

type LeftWayEstimator struct {
	t *tickets.Ticket
	way game.Point
}

func (lwe *LeftWayEstimator) Estimate() int{
	lwe.t.EstimateAuto(lwe.way)
	return lwe.t.GetAmount()
}

func (lwe *LeftWayEstimator) GetTicket() *tickets.Ticket{
	return lwe.t
}

func New(t *tickets.Ticket)*LeftWayEstimator{
	return &LeftWayEstimator{
		t: t,
	}
}