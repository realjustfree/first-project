package main
import "fmt"

func main(){
	var m map[int]string

	m=make(map[int]string)

	m[901]= "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)

	delete(m, 777)
	fmt.Println(m)
}