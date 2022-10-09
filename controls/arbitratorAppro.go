package controls

import (
	"sync"
	"time"

	"go.dev/models"
)

var (
	controller = sync.Mutex{}
	condition  = sync.NewCond(&controller)
)

// Function test if the upcoming intersections are in used
func allFree(intersection []*models.Intersection) bool {
	for _, it := range intersection {
		if it.PresentlyUsedBy > -1 {
			return false
		}
	}
	return true
}

/// Comprehensive func to scan and store the intersection within the range of
// train. After getting the intersection within the range, we try to lock
// intersection with the lower id first before the high, this is to prevent
// deadlocks

// Id is the train id
// reserves are the range we are taking into account
// crossings are the pointer to the traincrossing struct
func crossingsAhead(id, reserveStart, reserveEnd int, crossings []*models.TrainCrossing) {

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

	// Lock the mutex so the following codes can not be accessed by another
	// thread

	controller.Lock()

	// Throw the thread in loop if the upcoming intersections are in use
	for !allFree(allCrossings) {
		condition.Wait()
	}

	//  Now that the we have locked these steps, all other threads that
	// do not have its lane free can not reach these codes
	for _, it := range allCrossings {

		it.PresentlyUsedBy = id

		// Sleep for a short period
		time.Sleep(10 * time.Millisecond)
	}

	// Unlock the mutex
	controller.Unlock()
}

// The function simulates the movement of the train across the
// drawing board
// The train parameter is a pointer to slice of trains
// Distance is the distance the train is expected to cover in form of
// the position the train must be at the end of the program
// by increasing the pixel
// Crossings is [] of all the crossings

func Movement(train *models.Train, distance int, crossings []*models.TrainCrossing) {

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
				crossingsAhead(train.Id, trainCross.Position, trainCross.Position+train.Length, crossings)
			}

			// Calculate the position of the back of the brain by
			// subtracting the length of the train from the front
			// postion of the train
			trainBack := train.FrontPosition - train.Length

			// If the back position of the train is equals to postion of the
			// intersection then we want to unlock the section and
			// change the Id of the train using the intersection
			if trainBack == trainCross.Position {

				// Lock the following codes so other threads can not reach it
				controller.Lock()

				// Reset the ID of the train using the intersection
				trainCross.Intersection.PresentlyUsedBy = -1

				// Unlock the mutex
				controller.Unlock()

				// Broadcast so other threads can coutinue there process
				condition.Broadcast()
			}

		}

		// Sleep for some time before looping moving another pixel
		time.Sleep(30 * time.Millisecond)
	}
}
