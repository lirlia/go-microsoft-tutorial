package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var lineFlag = flag.Bool("n", false, "出力時の行番号の表示を行う")
var lineNum int = 1

// ファイルを開いて一行づつ出力する処理
func readFile(fn string) error {

	// ファイルが開けない場合はエラー
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	// ファイルの読み込み
	scanner := bufio.NewScanner(f)

	// lineFlagがtrueの場合は行番号を表示し
	// falseの場合は行番号を表示しない
	if *lineFlag == true {
		for scanner.Scan() {
			fmt.Printf("%d: %s\n", lineNum, scanner.Text())
			lineNum = lineNum + 1
		}
	} else {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	// ファイルが読み込めない場合はエラー
	if err := scanner.Err(); err != nil {
		return err
	}

	// 読み込みの処理を別関数にし、deferできるようにしている
	defer f.Close()
	return nil
}

// メイン
func main() {

	// 引数のパース
	flag.Parse()

	// 引数のファイルを読み込む
	for _, fn := range flag.Args() {
		err := readFile(fn)
		if err != nil {
			log.Fatal(err)
		}
	}
}
