package main

import "fmt"

// func f1() {
// 	fmt.Println("f1")
// }

// func f2(p1 string, p2 string) {
// 	fmt.Printf("f2: %s %s\n", p1, p2)
// }

func f3() string {
	return "f3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F$: %s %s", p1, p2) //a função sprintf retorna uma string formatada
}

func f5() (string, string) {
	return "retorno1", "retorno2"
}

func main() {
	rf3, rf4 := f3(), f4("Parametro1", "prametro2")
	fmt.Println(rf3)
	fmt.Println(rf4)

	//pode retornar mais de um parãmetro e recebo assim,
	//posso ignora um dos retornos colocando _ mas não deixar de considerar um dos retornos
	ret1, ret2 := f5()
	_, ret := f5()
	res, _ := f5()
	reti := f5()

}
