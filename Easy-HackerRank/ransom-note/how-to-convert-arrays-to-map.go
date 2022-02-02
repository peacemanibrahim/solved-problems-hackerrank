package main

import "fmt"

func main() {
	arrayOfMagazine := []string{"attack", "at", "dawn"}

	mapOfMagazine := make(map[string]string)
	for _, s := range arrayOfMagazine {
		// key(string)	|	value(string)
		mapOfMagazine[s] = s
	}
	fmt.Println(mapOfMagazine)

	// mapOfMagazine := make(map[int]string)
	// for i, s := range arrayOfMagazine {
	// 	// key(int)	|	value(string)
	// 	mapOfMagazine[i] = s
	// }
	// fmt.Println(mapOfMagazine)

	// Note: j is the index and k is the value
	// for j, k := range arrayOfMagazine {
	// 	fmt.Println(j)
	// 	fmt.Println(arrayOfMagazine[j])
	// 	fmt.Println(k)
	// }
}
