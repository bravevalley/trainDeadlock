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
	Trains        [4]*models.Train
	Intersections [4]*models.Intersection
)

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

	for i := 0; i < 4; i++ {
		Trains[i] = &models.Train{Id: i,
			Length:        trainLength,
			FrontPosition: 0}
	}

	for i := 0; i < 4; i++ {
		Intersections[i] = &models.Intersection{Id: i,
			Mu:              sync.Mutex{},
			PresentlyUsedBy: -1}
	}

	visual.Trains = Trains
	visual.Intersections = Intersections

	go controls.MoveTrain(Trains[0], 400, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[0]}, {Position: 175, Intersection: Intersections[1]}})

	go controls.MoveTrain(Trains[1], 400, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[1]}, {Position: 175, Intersection: Intersections[2]}})

	go controls.MoveTrain(Trains[2], 400, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[2]}, {Position: 175, Intersection: Intersections[3]}})

	go controls.MoveTrain(Trains[3], 400, []*models.TrainCrossing{{Position: 125, Intersection: Intersections[3]}, {Position: 175, Intersection: Intersections[1]}})

	ebiten.SetWindowSize(320*3, 320*3)
	ebiten.SetWindowTitle("Trains in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
