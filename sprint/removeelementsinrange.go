/*Write a Go function that takes an array of float64, two indices, and removes elements 
that lie between these indices from the array. The lower index should be removed, and the 
upper index should be kept. The function should return the resulting modified array.
 The indices can be negative or larger than the length of the array, but the function 
 should still remove the elements that fall within the specified range. The indices might 
 also be in wrong order.*/
 
package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64{
	if from > to {
		from, to = to, from
	}

	if from < 0 {
		from = 0
	}
	if to > len(arr) {
		to = len(arr)
	}

	if from >= to {
		return arr
	}
	return append(arr[:from], arr[to:]...)


}
