package main

import "fmt"

func main() {
	fmt.Println("Slices")

	// Creating an array
	arr := [4]string{"Learning", "how", "slice", "works"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2] // cria um novo um slice a partir da posição 1 do slice existente (arr) até a posição 2, mas não a 2 (1 < 2)
	my_slice_2 := arr[0:]     // cria um novo um slice a partir da posição 0 do slice existente (arr) e pega todas as posições
	my_slice_3 := arr[:2]     // cria um novo slice sem dar posição inicial e pega até a posição 2
	my_slice_4 := arr[:]      // cria um novo slice com todas as posições

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}
