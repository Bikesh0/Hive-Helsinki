package sprint 

func FilterBySum(arr [][]int, limit int) [][]int {

	var result [][]int
	for _, subarr := range arr {
		sum := 0

		for _, value := range subarr {
			sum += value
		}
		if sum >= limit {
			result = append(result, subarr)
		}
	}
	return result



}
