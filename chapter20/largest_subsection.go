package chapter20

// LargestSubsection returns the largest sum a continuous subsection of a list could produce.
func LargestSubsection(list []int) int {
	var (
		currentSum  int
		greatestSum int
	)

	for _, v := range list {
		if currentSum+v < 0 {
			currentSum = 0
		} else {
			currentSum += v

			if currentSum > greatestSum {
				greatestSum = currentSum
			}
		}
	}

	return greatestSum
}
