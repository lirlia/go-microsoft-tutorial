package imageChange

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// 変換するイメージ用の構造体
type IMG struct {
	Bef_ext  string
	Aft_ext  string
	Filepath string
}

// 指定したディレクトリ配下にある img_bf で指定された拡張子に
// 一致する画像ファイルを全て集めてきて配列で返す関数
func ListImg(dir *string, img_bf *string) (error, []string) {
	var fileList []string
	err := filepath.Walk(*dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+*img_bf {
				fileList = append(fileList, path)
			}
			return nil
		})
	if err != nil {
		return err, nil
	}

	return nil, fileList
}

// 入力されたファイルパスに存在するjpgファイルをpngに変換します
func ChangeImg(i IMG) error {

	// 画像ファイルを開く
	file, err := os.Open(i.Filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	// ファイルを画像オブジェクトに変換し変数に格納する
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 出力用のファイルを作成
	path := filepath.Dir(i.Filepath)
	filename := filepath.Base(i.Filepath[:len(i.Filepath)-len(filepath.Ext(i.Filepath))])
	out, err := os.Create(filepath.Join(path, filename) + "." + i.Aft_ext)
	defer out.Close()
	if err != nil {
		return err
	}

	// aftの形式に変換して保存
	switch i.Aft_ext {
	case "png":
		err = png.Encode(out, img)
	case "jpg":
		opts := &jpeg.Options{}
		err = jpeg.Encode(out, img, opts)
	case "default":
		errors.New(i.Aft_ext + " is invalid extention.")
	}

	if err != nil {
		return err
	}
	fmt.Printf("変換を行いました :%s\n", i.Filepath)
	return nil
}
