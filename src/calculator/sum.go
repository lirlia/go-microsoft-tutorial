package calculator

// プライベートにする要素は、名前の先頭を小文字にします。
// パブリックにする要素は、名前の先頭を大文字にします。

var logMessage = "[LOG]"
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
