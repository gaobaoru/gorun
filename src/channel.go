package main
import "fmt"
func main(){
    ch := make(chan string, 1)
    go func(){
	ch <- ("已到达！")
    }()
    var value string = "数据"
    value = value + (<-ch)
    fmt.Println(value)
}
