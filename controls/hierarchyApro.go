package controls

import (
	"sort"
	"time"

	"go.dev/models"
)

// Comprehensive func to scan and store the intersection within the range of
// train. After getting the intersection within the range, we try to lock
// intersection with the lower id first before the high, this is to prevent
// deadlocks

// Id is the train id
// reserves are the range we are taking into account
// crossings are the pointer to the traincrossing struct
func crossingAhead(id, reserveStart, reserveEnd int, crossings []*models.TrainCrossing) {

	// Create a variable to store the list of the intersection within
	// our scope
	var allCrossings []*models.Intersection

	// Loop over all the crossings on the map to find the one in our range
	for _, intersections := range crossings {

		// The intersection we can concerned about are the intersections within
		// the reserves that is  if the intersectio position is greater than
		//  the start of the reserve and lower than the end of the reserve
		// And if the interection is not in use by us already
		if reserveStart <= intersections.Position && reserveEnd >= intersections.Position && intersections.Intersection.PresentlyUsedBy != id {

			// Any intersection that passed the test is within our range
			allCrossings = append(allCrossings, intersections.Intersection)
		}
	}

	// Use the Slice method of sort to rearrange the intersection so the
	// lower intersection is locked first since we are using the
	// RESOURCE HIERARCHY approach to prevent deadlock
	sort.Slice(allCrossings, func(i, j int) bool {
		return allCrossings[i].Id < allCrossings[j].Id
	})

	// Now lock each intersections in the [] we capture and change them
	// Id of the train using it to the present train
	for _, it := range allCrossings {
		it.Mu.Lock()
		it.PresentlyUsedBy = id

		// Sleep for a short period
		time.Sleep(10 * time.Millisecond)
	}
}

// The function simulates the movement of the train across the
// drawing board
// The train parameter is a pointer to slice of trains
// Distance is the distance the train is expected to cover in form of
// the position the train must be at the end of the program
// by increasing the pixel
// Crossings is [] of all the crossings

func Locomotive(train *models.Train, distance int, crossings []*models.TrainCrossing) {

	// While the front of the train is less than the distance to be covered
	for train.FrontPosition < distance {

		// Add a pixel to the frontpostion of the train
		train.FrontPosition += 1

		// Loop over all the crossing postions
		for _, trainCross := range crossings {

			// Test if the front of the train is in the postion as the
			// as the position of the intersection
			if train.FrontPosition == trainCross.Position {

				// The cross function is called if the train is at the point of
				// intersection so as to lock the intersection
				crossingAhead(train.Id, trainCross.Position, trainCross.Position+train.Length, crossings)
			}

			// Calculate the position of the back of the brain by
			// subtracting the length of the train from the front
			// postion of the train
			trainBack := train.FrontPosition - train.Length

			// If the back position of the train is equals to postion of the
			// intersection then we want to unlock the section and
			// change the Id of the train using the intersection
			if trainBack == trainCross.Position {

				// Reset the ID of the train using the intersection
				trainCross.Intersection.PresentlyUsedBy = -1

				// Unlock the intersection
				trainCross.Intersection.Mu.Unlock()
			}

		}

		// Sleep for some time before looping moving another pixel
		time.Sleep(30 * time.Millisecond)
	}
}
