package main

import (
	"fmt"
	"reflect"
)

func main() {
	array1 := [3]int{1, 2, 3} //isso é uma declaraçõa e inicialização de array, deve ser passado o tamanho

	//essa é uma declaraçõa de slice, não precisa passar tamanho
	//slice são dinãmicos podem crescer a medida que necessitar
	//slice tem um ponteiro e eponta para uma parte do array
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	//outra forma de criar slice
	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [6]int{1, 2, 3, 4, 5, 6}
	slice3 := array2[1:3] //slice3 vai conter os elementos das posições 1 à 3 de array 2, mas sem incluir a posição 3
	fmt.Println(slice3)
}
