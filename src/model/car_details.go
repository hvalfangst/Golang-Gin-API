package model

type CarDetails struct {
	Car       *Car
	Owner     *Owner
	Tires     []*Tire
	Insurance []*Insurance
	Repairs   []*Repair
	Engine    *Engine
}
