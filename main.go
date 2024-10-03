package main
import "fmt"
func secondLargest(arr []int) int {
	if len(arr) < 2 {
		
		return -1 
	}
	first := arr[0]
	second := -1 

	for _, num := range arr[1:] {
		if num > first {
			second = first
			first = num
		} else if num > second && num < first {
			second = num
		}
	}
	if second == -1 {
		
		return -1
	}
	return second
}
func main() {
	arr := []int{35, 43, 41, 22, 54}
	result := secondLargest(arr)
	if result != -1 {
		fmt.Println("The second largest element is:", result)
	} else {
		fmt.Println("Not enough elements to determine the second largest.")
	}
}
