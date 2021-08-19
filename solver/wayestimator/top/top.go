package top

import (
	"battlecity_test/game"
	"battlecity_test/solver/tickets"
)

type TopWayEstimator struct {
	t *tickets.Ticket
	way game.Point
}

func (twe *TopWayEstimator) Estimate() int{
	twe.t.EstimateAuto(twe.way)
	return twe.t.GetAmount()
}

func (twe *TopWayEstimator) GetTicket() *tickets.Ticket{
	return twe.t
}

func New(t *tickets.Ticket)*TopWayEstimator{
	return &TopWayEstimator{
		t: t,
	}
}