package main

import (
	"fmt"
	"os"
	"strings"
    "github.com/zzStone/ProgrammingInGo/chapter1/stacker"
)

func main() {
	who := "world"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello", who)

    var haystack stacker.Stack
    haystack.Push("hay")
    haystack.Push(-15)
    haystack.Push([]string{"ala", "bit", "da"})
    haystack.Push(81.52)
    for {
        item, err := haystack.Pop()
        if err != nil {
            break
        }
        fmt.Println(item)
    }    
}
