package main

//import "github.com/myuser/calculator"
//import "github.com/lirlia/calculator-go"
import (
	"github.com/myuser/calculator"
	//"github.com/lirlia/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
	println(quote.Hello())
}
