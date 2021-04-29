package main

import (
	"flag"
	"fmt"
	"log"

	im "test/modules"
)

var img_dir = flag.String("dir", "", "画像が保存されているディレクトリを指定")
var img_bf = flag.String("bef", "jpg", "変換前の画像の拡張子を指定")
var img_af = flag.String("aft", "png", "変換後の画像の拡張子を指定")

func main() {

	flag.Parse()

	// 処理前のエラー判定
	// 変換前後の拡張子が異なっているかの判定
	if *img_bf == *img_af {
		log.Fatal("変換前後の拡張子が同じです")
	}

	// 変換前後の拡張子がjpg or pngであるかのチェック
	//if img_bf != ("jpg" or "png") {
	//	log.Fatal("変換前後の拡張子が同じです")
	//}

	// 引数で指定されたディレクトリ配下の画像ファイルのパスを取得する
	err, fileList := im.ListImg(img_dir, img_bf)
	fmt.Println(fileList)
	if err != nil {
		log.Fatal(err)
	}

	// ファイルの絶対パスが含まれる配列に対して
	// 画像変換の処理を実施する
	for _, fn := range fileList {
		img := im.IMG{*img_bf, *img_af, fn}
		err = im.ChangeImg(img)
		if err != nil {
			log.Fatal(err)
		}
	}
}
