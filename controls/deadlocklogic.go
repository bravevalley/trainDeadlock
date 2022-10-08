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
				trainCross.Intersection.Mu.Unlock()
				trainCross.Intersection.PresentlyUsedBy = -1
			}

			time.Sleep(1900 * time.Millisecond)
		}
	}
}
