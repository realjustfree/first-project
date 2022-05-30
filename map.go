package main
import "fmt"

func main(){
	myMap := map[string]string{
		"A":"Apple",
		"B":"Banana",
		"c":"Charlie",
	}
	for key, val := range myMap{
		fmt.Println(key, val)
	}
}