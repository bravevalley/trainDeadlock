package visual

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go.dev/models"
	"image/color"
	"math"
)

var (
	colours = [4]color.RGBA{
		{233, 33, 40, 255},
		{78, 151, 210, 255},
		{251, 170, 26, 255},
		{11, 132, 54, 255},
	}

	white = color.RGBA{R: 185, G: 185, B: 185, A: 255}
)

// Declare the variable so it values can be imported from the
// main module
var Trains [4]*models.Train
var Intersections [4]*models.Intersection

func DrawIntersections(screen *ebiten.Image) {
	drawIntersection(screen, Intersections[0], 145, 145)
	drawIntersection(screen, Intersections[1], 175, 145)
	drawIntersection(screen, Intersections[2], 175, 175)
	drawIntersection(screen, Intersections[3], 145, 175)
}

func DrawTracks(screen *ebiten.Image) {
	for i := 0; i < 300; i++ {
		screen.Set(10+i, 135, white)
		screen.Set(185, 10+i, white)
		screen.Set(310-i, 185, white)
		screen.Set(135, 310-i, white)
	}
}

func DrawTrains(screen *ebiten.Image) {
	drawXTrain(screen, 0, 1, 10, 135)
	drawYTrain(screen, 1, 1, 10, 185)
	drawXTrain(screen, 2, -1, 310, 185)
	drawYTrain(screen, 3, -1, 310, 135)
}

func drawIntersection(screen *ebiten.Image, intersection *models.Intersection, x int, y int) {
	c := white
	if intersection.PresentlyUsedBy >= 0 {
		c = colours[intersection.PresentlyUsedBy]
	}
	screen.Set(x-1, y, c)
	screen.Set(x, y-1, c)
	screen.Set(x, y, c)
	screen.Set(x+1, y, c)
	screen.Set(x, y+1, c)
}

func drawXTrain(screen *ebiten.Image, id int, dir int, start int, yPos int) {
	s := start + (dir * (Trains[id].FrontPosition - Trains[id].Length))
	e := start + (dir * Trains[id].FrontPosition)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(int(i)-dir, yPos-1, colours[id])
		screen.Set(int(i), yPos, colours[id])
		screen.Set(int(i)-dir, yPos+1, colours[id])
	}
}

func drawYTrain(screen *ebiten.Image, id int, dir int, start int, xPos int) {
	s := start + (dir * (Trains[id].FrontPosition - Trains[id].Length))
	e := start + (dir * Trains[id].FrontPosition)
	for i := math.Min(float64(s), float64(e)); i <= math.Max(float64(s), float64(e)); i++ {
		screen.Set(xPos-1, int(i)-dir, colours[id])
		screen.Set(xPos, int(i), colours[id])
		screen.Set(xPos+1, int(i)-dir, colours[id])
	}
}
