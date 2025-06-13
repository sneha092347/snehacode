package main 
import "fmt"
func main () {
	fmt.Println("hi")
	caller("yes")
}

func caller(a string) {
	fmt.Println(a)
}
