package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 標準入力を引数に関数呼び出し
	// チャネルを変数に格納
	ch := input(os.Stdin)
	for {
		// 改行コードをなしにして入力を受け付けているように表現
		fmt.Print(">")
		// チャネルの内容を出力
		// チャネル経由でデータが送られるまで待機
		fmt.Println(<-ch)
		fmt.Println("ここの文字列は↑が実行されるまで表示されない")
	}
}

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)
	// ここをgo routineにすることで、きちんとreturnできるようになっている
	// 仮にここをなくすとbufio.NewScannerにひっかかり続けるので処理が進まなくなる
	// その時にstdinするとエラーで落ちる
	go func() {

		// io.Readerを引数にScannerオブジェクトを生成する
		s := bufio.NewScanner(r)
		// TODO: チャネルを閉じる
		defer close(ch)

		// 入力の終わりまでのデータをつかってループを行う
		// 入力される限り無限ループする
		for s.Scan() {
			// scanで進めた入力のTextをチャネルに突っ込む
			ch <- s.Text()
			// TODO: チャネルに読み込んだ文字列を送る
		}
	}()
	// TODO: チャネルを返す
	return ch
}
