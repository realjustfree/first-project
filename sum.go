package main
import "fmt"

func main(){
	c, t := sum(1,7,5,3,9)

	fmt.Println(c, t)
}

func sum(nums ...int)(int, int){
	s:=0
	c:=0

	for _, n := range(nums){
		s += n
		c++
	}
	return c, s
}