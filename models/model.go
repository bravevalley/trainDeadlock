package models

import "sync"

// The model of a train which includes
// Id for identification
// Length for calc
// FrontPostion for movement
type Train struct {
	Id            int
	Length        int
	FrontPosition int
}

// The Intersection struct containing the ID, mutex to lock the intersection
// and the train currently using the intersection
type Intersection struct {
	Id              int
	Mu              sync.Mutex
	PresentlyUsedBy int
}

// More details about an intersection
type TrainCrossing struct {
	Position     int
	Intersection *Intersection
}
