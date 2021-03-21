package main

func main() {
	firstName := "john"
	updateName(firstName)
	println(firstName)

	// ポインタ: 別の変数のメモリアドレスを格納する変数
	// &演算子: &の後にあるオブジェクトのメモリアドレスを取得
	// &firstName: firstNameのメモリアドレスを取得
	// firstNameが格納されたメモリアドレスを渡す
	updateNamePointer(&firstName)
	println(firstName)
}

// Goでは値渡しを行う
func updateName(name string) {
	name = "david"
}

// 意図的に参照渡しを行う
// ここでいうとnameがポインタ、*nameでポインタを逆参照する
// *演算子: ポインタに格納されたメモリアドレスにあるオブジェクトへのアクセスを付与
// ここでいうとfirstName変数のメモリアドレスに対してname変数で触れるようにしている
// メモリアドレスに格納された変数の型で受け取る
func updateNamePointer(name *string) {
	*name = "123"

	// すでにstring型でメモリを確保しているのでint型して突っ込めない
	//cannot assign int to *name (type string) in multiple assignment
	//*name, _ = strconv.Atoi(*name)
}
