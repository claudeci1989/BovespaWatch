package main

//https://gobyexample.com/json
import (
    "encoding/json"
    "os"
    "fmt"
)

func main() {
    fmt.Println("The Bovespa Watch Parser")

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    enc := json.NewEncoder(os.Stdout)
    enc.Encode(mapD)
}
