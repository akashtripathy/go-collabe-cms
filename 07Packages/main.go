/*
Purpose of a package is to design and maintain a large number of programs by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs.

If we are developing in a stringtools directory, temporary module path might be <company-name>/stringtools
e.g go mod init <company-name>/stringtools

1.Locating and importing useful packages : from pkg.go.dev
2.Enabling dependency tracking in your code : go mod init example/mymodule
3.Naming a module : <prefix>/<descriptive-text>
4.Adding a dependency : go get command(go get . , go get example/theirmodule)
*/

package main

import (
	calculator "example/mymodule/packages"
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Print(quote.Glass())

	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))

	// use the subtract function of calculator package
	fmt.Println(calculator.Subtract(number1, number2))
}

// Resource 1: https://go.dev/doc/modules/managing-dependencies
// Resource 2: https://www.golang-book.com/books/intro/11
// Resource 3: https://www.programiz.com/golang/packages
