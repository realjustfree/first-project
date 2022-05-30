package main
import "fmt"

func main(){
f1 := []string{"1", "2", "3"}
f2 :=[]string{"4","5"}
f3 := append(f1, f2...)
f4 := append(f1[:2], f2...)
f3_1 := append(f1, "x","y")

fmt.Println(f1)
fmt.Println(f2) 
fmt.Println(f3) // 1 2 3 4 5
fmt.Println(f4) // 1 2 4 5
fmt.Println(f3_1)

}