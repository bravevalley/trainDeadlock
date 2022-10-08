package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"go.dev/controls"
	"go.dev/models"
	"go.dev/visuals"
)

var (

	// Define an array for the trains and intersections which implicitly define the numbers
	// of trains and intersections
	Trains        [4]*models.Train
	Intersections [4]*models.Intersection
)

// The unanimous lenght of the train
const trainLength = 80

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	visual.DrawTracks(screen)
	visual.DrawIntersections(screen)
	visual.DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return 320, 320
}

func main() {

	// Define the each train in a loop and assign the
	// index value as the train ID
	for i := 0; i < 4; i++ {
		Trains[i] = &models.Train{Id: i,
			Length:        trainLength,
			FrontPosition: 0}
	}
	// Define the each intersection in a loop and assign the
	// index value as the intersection ID
	for i := 0; i < 4; i++ {
		Intersections[i] = &models.Intersection{Id: i,
			Mu:              sync.Mutex{},
			PresentlyUsedBy: -1}
	}

	// Send the variable to the visual module for processing
	visual.Trains = Trains
	visual.Intersections = Intersections

	// Create go routine to handle each trains, the parameters have been defined
	// in the module
	go controls.Locomotive(Trains[0], 300, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[0]}, {Position: 175, Intersection: Intersections[1]}})

	go controls.Locomotive(Trains[1], 300, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[1]}, {Position: 175, Intersection: Intersections[2]}})

	go controls.Locomotive(Trains[2], 300, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[2]}, {Position: 175, Intersection: Intersections[3]}})

	go controls.Locomotive(Trains[3], 300, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[3]}, {Position: 175, Intersection: Intersections[0]}})

	ebiten.SetWindowSize(320*3, 320*3)
	ebiten.SetWindowTitle("Trains in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
