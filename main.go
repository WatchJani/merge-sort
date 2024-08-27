package main

import "fmt"

func main() {
	memTable := []Read{{"M", 5}, {"M", 152}, {"M", 547}, {"M", 4114}, {"M", 5451}}
	memSort := Sort{
		memTable, 0,
	}

	freezeTable := []Read{{"M", 2}, {"M", 100}, {"M", 647}, {"M", 3222}, {"M", 6451}}
	freezeSort := Sort{
		freezeTable, 0,
	}

	discTable := []Read{{"D", 2000}, {"C", 2020}, {"D", 3201}, {"D", 3222}, {"D", 4201}}
	discSort := Sort{
		discTable, 0,
	}

	fmt.Println(MergeSort([]Sort{memSort, freezeSort, discSort}))

	// fmt.Println(memTable)
}

type Read struct {
	source string
	offset int
}

type Sort struct {
	list    []Read
	pointer int
}

func MergeSort(tro []Sort) []Read {
	var size int
	for i := 0; i < len(tro); {
		ln := len(tro[i].list)
		if ln != 0 {
			size += ln
			i++
		} else {
			tro = append(tro[:i], tro[i+1:]...)
		}
	}

	res := make([]Read, 0, size)

	for index := 0; index < size; index++ {
		var (
			minKey   Read = tro[0].list[tro[0].pointer]
			minIndex int  = 0
		)

		for i := 1; i < len(tro); i++ {
			v := &tro[i]

			if v.list[v.pointer].offset < minKey.offset {
				minKey = v.list[v.pointer]
				minIndex = i
			}
		}

		res = append(res, minKey)
		tro[minIndex].pointer++

		if tro[minIndex].pointer >= len(tro[minIndex].list) {
			tro = append(tro[:minIndex], tro[minIndex+1:]...)
		}
	}

	return res
}

// func MergeSort(tro []Sort) []Read {
// 	var size int
// 	for index, obj := range tro {
// 		ln := len(obj.list)
// 		if ln != 0 {
// 			size += ln
// 		} else {
// 			tro = append(tro[:index], tro[index+1:]...)
// 		}
// 	}

// 	res := make([]Read, 0, size)

// 	for index := 0; index < size; index++ {
// 		var (
// 			minKey   Read = tro[0].list[tro[0].pointer]
// 			previous Read
// 		)

// 		for i := 0; i < len(tro); i++ {
// 			v := &tro[i]
// 			if v.list[v.pointer].offset == previous.offset || v.list[v.pointer].offset == minKey.offset {
// 				v.pointer++ //why is this pointer is not updated
// 				if len(v.list) == v.pointer {
// 					tro = append(tro[:i], tro[i+1:]...)
// 					//check if tro is empty [empty is if just one element inside tro]
// 					continue
// 				}

// 				//check if the next key is that key
// 				if v.list[v.pointer].offset < minKey.offset {
// 					minKey = v.list[v.pointer]
// 				}
// 			} else if v.list[v.pointer].offset < minKey.offset {
// 				minKey = v.list[v.pointer]
// 			}
// 		}

// 		res = append(res, minKey)
// 		previous = minKey
// 	}

// 	return res
// }
