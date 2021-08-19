package right

import (
	"battlecity_test/game"
	"battlecity_test/solver/tickets"
)

type RightWayEstimator struct {
	t *tickets.Ticket
	way game.Point
}

func (rwe *RightWayEstimator) Estimate() int{
	rwe.t.EstimateAuto(rwe.way)
	return rwe.t.GetAmount()
}

func (rwe *RightWayEstimator) GetTicket() *tickets.Ticket{
	return rwe.t
}

func New(t *tickets.Ticket)*RightWayEstimator{
	return &RightWayEstimator{
		t: t,
	}
}