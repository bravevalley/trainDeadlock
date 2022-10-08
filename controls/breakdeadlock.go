package controls

import (
	"sort"
	"time"

	"go.dev/models"
)

func crossingAhead(id, reserveStart, reserveEnd int, crossings []*models.TrainCrossing) {
	var allCrossings []*models.Intersection

	for _, intersections := range crossings {
		if reserveStart <= intersections.Position && reserveEnd >= intersections.Position && intersections.Intersection.PresentlyUsedBy != id {
			allCrossings = append(allCrossings, intersections.Intersection)
		}
	}

	sort.Slice(allCrossings, func(i, j int) bool {
		return allCrossings[i].Id < allCrossings[j].Id
	})

	for _, it := range allCrossings {
		it.Mu.Lock()
		it.PresentlyUsedBy = id
		time.Sleep(10 * time.Millisecond)
	}
}

func Locomotive(train *models.Train, distance int, crossings []*models.TrainCrossing) {
	for train.FrontPosition < distance {
		train.FrontPosition += 1

		for _, trainCross := range crossings {
			if train.FrontPosition == trainCross.Position {
				crossingAhead(train.Id, trainCross.Position, trainCross.Position+train.Length, crossings)
			}
			trainBack := train.FrontPosition - train.Length

			if trainBack == trainCross.Position {
				trainCross.Intersection.PresentlyUsedBy = -1
				trainCross.Intersection.Mu.Unlock()
			}

		}
		time.Sleep(30 * time.Millisecond)
	}
}
