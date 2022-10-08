package models

import "sync"

type Train struct {
	Id            int
	Length        int
	FrontPosition int
}

type Intersection struct {
	Id              int
	Mu              sync.Mutex
	PresentlyUsedBy int
}

type TrainCrossing struct {
	Position     int
	Intersection *Intersection
}
