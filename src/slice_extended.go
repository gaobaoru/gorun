package main
import "fmt"
func main(){
    var number = [...]int{1,2,3,4,5,6,7,8,9,10}
    slice := number[4:6:8]
    length := 2
    capacity := 4
    fmt.Printf("%v, %v\n",length == len(slice), capacity == cap(slice))
    slice = slice[:cap(slice)]
    slice = append(slice, 11, 12, 13)
    length = 7
    fmt.Printf("%v\n", length == len(slice))
    slice1 := []int{0, 0, 0}
    copy(slice, slice1)
    e2 := 0
    e3 := 8
    e4 := 11
    fmt.Printf("%v, %v, %v\n", e2 == slice[2], e3 == slice[3], e4 == slice[4])
}
