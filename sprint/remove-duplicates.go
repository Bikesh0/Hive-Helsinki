package sprint

func RemoveDuplicates(arr []int) []int {
	// If the array is empty or nil, return nil
	if len(arr) == 0 {
		return nil
	}

	// 'seen' remembers which numbers we have already found
	// Example: seen[3] = true means we already saw 3
	seen := make(map[int]bool)

	// This will contain our final array without duplicates
	result := []int{}

	// Go through every number in the input array
	for _, n := range arr {

		// Check if we have NOT seen this number before
		if !seen[n] {

			// Add the number to our result
			result = append(result, n)

			// Remember that we have seen this number
			seen[n] = true
		}
	}

	// Return the array without duplicates
	return result
}
