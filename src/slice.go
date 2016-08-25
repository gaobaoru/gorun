package main
import "fmt"
func main(){
    var number = [5]int{1,2,3,4,5}
    slice := number[2 : len(number)]
    length := 3
    capacity := 3
    fmt.Printf("%v, %v\n", (length == len(slice)), (capacity == cap(slice)))
}
