package main
import "fmt"
func main(){

    var num1 int = -0x1000
    fmt.Printf("16进制 %X 表示的是 %d \n", num1, (-4096))
}
