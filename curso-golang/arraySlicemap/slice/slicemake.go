package main

import "fmt"

func main() {
	s := make([]int, 10) //cria com 10 posições zeradas
	fmt.Println(s)
	s2 := []int{} //cria slice nulo
	fmt.Println(s2)

	s[8] = 5
	fmt.Println(s)
	//Nesse caso não posso fazer essa atribuição
	// s2[1] = 4
	// fmt.Println(s2)

	// s3 := make([]int, 5, 10)
	// fmt.Println(s3, len(s3), cap(s3))
	s3 := append(s, 1, 2, 3)
	fmt.Println(s3, len(s3), cap(s3))
	//Não estão apontando para o mesmo lugar, mas acho que deveria
	s[1] = 100
	s3[0] = 1
	fmt.Println(s, s3)

}
