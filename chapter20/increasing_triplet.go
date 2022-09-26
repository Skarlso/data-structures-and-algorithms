package chapter20

import "math"

// IncreasingTriplet finds three items in the list can result in an upwards trend.
func IncreasingTriplet(list []int) bool {
	if len(list) == 0 {
		return false
	}
	lowest := list[0]
	middle := math.MaxInt
	for _, v := range list {
		if v < lowest {
			lowest = v
		} else if v <= middle {
			middle = v
		} else {
			return true
		}
	}

	return false
}
