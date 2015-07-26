package main

import "fmt"

func main() {

	val := 123

	// 関数リテラルの記述と呼び出しを同時に行う
	func(i int) {
		// 関数リテラル外の変数valにアクセス可能
		fmt.Println(i * val)
	}(10)

	// 関数リテラルを変数に代入
	f := func(s string) {
		fmt.Println(s)
	}

	// 関数の呼び出し    
	f("hoge")
}
