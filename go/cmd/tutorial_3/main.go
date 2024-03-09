package main

import "fmt"

// array (fixed), slice (dynamic), map, loop

func main() {
	// array: ukuran array di go bersifat FIXED
	var numbers1 [3]int
	numbers1[0] = 9
	numbers1[1] = 10
	numbers1[2] = 11

	fmt.Println(numbers1)

	var numbers2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)

	numbers := [...]int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	fmt.Println(numbers)

	// loop array
	for i := 0; i < len(numbers2); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, numbers2[i])
	}

	// slice: fungsinya sama seperti array, hanya saja memiliki ukuran yang DINAMIS
	// slice lebih umum digunakan di GO daripada array
	sliceofNumbers := []int{}
	sliceofNumbers = append(sliceofNumbers, 100, 200)
	fmt.Println(sliceofNumbers, len(sliceofNumbers))

	// map
	var mahasiswa map[string]int
	mahasiswa = make(map[string]int)

	mahasiswa["John"] = 30
	mahasiswa["Mary"] = 40

	fmt.Println("Umur John: ", mahasiswa["John"])

	// delete element dari map
	delete(mahasiswa, "John")
	fmt.Println(mahasiswa)

	// Mengecek apakah kunci ada dalam map
	val, isExist := mahasiswa["John"]
	fmt.Println("Umur John setelah dihapus:", val, "Ada:", isExist)

	// declare map with literal map
	mahasiswa2 := map[string]int{"Shidqi": 30, "Nauval": 25}
	fmt.Println(mahasiswa2)

	// looping map
	for nama, umur := range mahasiswa {
		fmt.Printf("%s berumur %d tahun\n", nama, umur)
	}

	// Di GO tidak ada while / do-while
	// implementasi while dengan for
	i := 0
	for i < 10 {
		i++
	}
	fmt.Println(i)

	// implementasi do-while dengan for
	x := 0
	for {
		x++
		if x >= 0 {
			break
		}
	}

	fmt.Println(x)
}
