package main



import (
	// "fmt"
	// "github.com/cvenkman/data_structures/datastruct"
    "github.com/cvenkman/data_structures/basic"
)



func main(){
    // basic.CheckBrackets("{[}");
    // basic.CheckBrackets("([](){([])})") // Success
    // basic.CheckBrackets("{*}") // Success
    // basic.CheckBrackets("{}") // Success
    // basic.CheckBrackets("") // Success
    // basic.CheckBrackets("*{}") // Success

    // basic.CheckBrackets("()[]}") // 5
    // basic.CheckBrackets("{{[()]]") // 7

    // basic.CheckBrackets("{[]") // 3
    basic.CheckBrackets("{{{[][][]") // 3
    basic.CheckBrackets("{*{{}") // 3

    // basic.CheckBrackets("[[*") // 2
    // basic.CheckBrackets("{{") // 2
    // basic.CheckBrackets("}") // 1
    // basic.CheckBrackets("{{{**[][][]") // 3
}

