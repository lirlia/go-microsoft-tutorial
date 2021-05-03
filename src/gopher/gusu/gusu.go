package main

func gusuchk(n int) bool {
	switch n % 2 {
	// 2で割り切れる時は偶数
	case 0:
		return true
	// それ以外は奇数
	default:
		return false
	}
}

func main() {

}
