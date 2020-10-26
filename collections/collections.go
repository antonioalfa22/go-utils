package collections

import "github.com/thoas/go-funk"


/**
	Example:
		Find([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
*/
func Find(arr interface{}, predicate interface{}) interface{} {
	return funk.Find(arr, predicate)
}

/**
	Example:
		Map([]int{1, 2, 3, 4}, func(x int) int { return x * 2 })
*/
func Map(arr interface{}, predicate interface{}) interface{} {
	return funk.Map(arr, predicate)
}

/**
	Example:
		ForEach([]int{1, 2, 3, 4}, func(x int) { fmt.Println(x) })
*/
func Filter(arr interface{}, predicate interface{}) interface{} {
	return funk.Filter(arr, predicate)
}

/**
	Example:
		Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
 */
func ForEach(arr interface{}, predicate interface{}) {
	funk.ForEach(arr, predicate)
}