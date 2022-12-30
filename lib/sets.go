package lib

func listLens[A any](lists ...[]A) ([]int, int) {
	var lengths []int
	size := 1
	for _, list := range lists {
		lengths = append(lengths, len(list))
		size *= len(list)
	}
	return lengths, size
}
func Prod[A any](lists ...[]*A) [][]*A {
	lengths, size := listLens(lists...)

	var allItems [][]*A
	for point := 0; point < size; point++ {
		var items []*A
		tempPoint := point
		for i, length := range lengths {
			index := tempPoint % length
			item := lists[i][index]
			items = append(items, item)
			tempPoint /= length
		}
		allItems = append(allItems, items)
	}
	return allItems
}
