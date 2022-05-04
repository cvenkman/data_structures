package basic

/*
	Проверить, правильно ли расставлены скобки в данном коде.
	[()] or foo(bar); - Success
	{[} or foo(bar[i); - 3

	Формат входа.
	Строка s[1 . . . n], состоящая из заглавных и прописных букв
	латинского алфавита, цифр, знаков препинания и скобок из множества []{}().

	Формат выхода.
	Если скобки в s расставлены правильно, выведите
	строку “Success". В противном случае выведите индекс (используя индексацию с единицы)
	первой закрывающей скобки, для которой нет соответствующей открывающей. Если такой нет,
	выведите индекс первой открывающей скобки, для которой нет соответствующей закрывающей.
*/

import (
	"fmt"
	"github.com/cvenkman/data_structures/datastruct"
)

//"[(}]"
func CheckBrackets(input string) {
	// var s datastruct.Stack
	s := make(datastruct.Stack, 0)
	var p, elem rune
	var err error

	// s, p, err = s.Pop()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(p)


	closeBrackets := make(map[rune]rune, 3)
	closeBrackets[']'] = '['
	closeBrackets[')'] = '('
	closeBrackets['}'] = '{'
	
	i := 0
	f := 0
	for i, elem = range input {
		switch elem {
		case '[', '(', '{':
			s = s.Push(elem)
			f++
		case ']', ')', '}':
			s, p, err = s.Pop()
			if err != nil {
				fmt.Println(err)
				fmt.Println(i + 1)
				return
			}
			// fmt.Println(string(elem), string(closeBrackets[elem]), string(p))
			if closeBrackets[elem] != p {
				fmt.Println(i + 1)
				return
			}
			f--
		}
	}
	// s, p = s.Pop()
	if s.Empty() == true {
		fmt.Println("Success")
	}
	fmt.Println(f + 1)
}