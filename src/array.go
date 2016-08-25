package main
import "fmt"

func main(){
    var number [5]int
    number[0] = 2
    number[3] = number[0] - 3
    number[1] = number[2] + 5
    number[4] = len(number)
    sum := (11)
    fmt.Printf("%v\n",(sum == number[0]+number[1]+number[2]+number[3]+number[4]))
}
