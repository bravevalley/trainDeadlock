package controls

import (
	"time"

	"go.dev/models"
)

func MoveTrain(train *models.Train, distance int, crossings []*models.TrainCrossing) {
	for train.FrontPosition < distance {
		train.FrontPosition += 1

		for _, trainCross := range crossings {
			if train.FrontPosition == trainCross.Position {
				trainCross.Intersection.Mu.Lock()
				trainCross.Intersection.PresentlyUsedBy = train.Id
			}
			trainBack := train.FrontPosition - train.Length

			if trainBack > trainCross.Position {
				trainCross.Intersection.PresentlyUsedBy = -1
				trainCross.Intersection.Mu.Unlock()
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}
