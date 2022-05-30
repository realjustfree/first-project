package main
import (
	"fmt"
	"strconv"
)

func main(){
	v:="10"

	s, err := strconv.Atoi(v)
	k, err_k := strconv.Atoi(100)

	if err_k != nil {
		fmt.Printf("%T, %v \n", err_k, err_k)
	}

	if err == nil{
		fmt.Printf("%T, %v \n", s, s )
	}
}