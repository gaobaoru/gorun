package main
import "fmt"
func main(){
    mm := map[string]int{"golang": 42, "java": 1, "python": 8}
    mm["scala"] = 25
    mm["erlang"] = 50
    mm["python"] = 0
    fmt.Printf("%d, %d, %d\n", mm["scala"], mm["erlang"], mm["python"])
}
