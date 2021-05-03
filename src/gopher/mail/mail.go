package main

import (
	"strings"
)

func chkAddress(m string) bool {

	// @マークの数を確認(1以外ならエラー)
	// 仕様:メールアドレスは「**ローカルパート@ドメインパート**」で構成されるより
	if strings.Count(m, "@") != 1 {
		return false
	}

	// ローカルパートとドメインパートの分割
	m_split := strings.Split(m, "@")

	// 文字数チェック
	// 仕様:ローカルパート/ドメインパートはそれぞれ「**2文字以上5文字以下**」で構成されるより
	for _, s := range m_split {
		if len(string(s)) < 2 || len(string(s)) > 5 {
			return false
		}
	}

	// 文字種別チェック
	// 仕様:ローカルパート/ドメインパートでは「**半角小文字英数字と記号のみ**」で構成されるより
	// 仕様:使用可能な記号は「**-_.のみ**」とする
	lowerA := rune('a')
	lowerZ := rune('z')
	num0 := rune('0')
	num9 := rune('9')

	kigo := []rune{
		rune('-'),
		rune('.'),
		rune('_'),
		rune('@'),
	}

	// 入力された文字がkigoに合致するかどうかをチェックする関数
	kigoFunc := func(c rune) bool {
		for _, v := range kigo {
			if v == c {
				return true
			}
		}
		return false
	}

	rb := false
	for i, s := range m {

		r := kigoFunc(s)

		// 要件:ローカルパート/ドメインパートの**最初と最後に記号は使用できない**ものとする
		// ここではメールアドレス全体の最初と最後が記号ではないかをチェックする
		if (i == 0 || i == len(m)-1) && (r == true) {
			return false
		}

		// 要件:**記号は連続して利用できない**ものとする(@の前後も不可とする)
		// 記号が連続していないかのチェック(前回も今回も記号がokになっているか?)
		if i > 0 && r == true && rb == true {
			return false
		}

		// 結果を複製
		rb = r

		// 条件外の特殊文字の利用有無をチェック
		if (r == true) ||
			// 半角小文字英数字の利用有無をチェック
			(s >= lowerA && s <= lowerZ) ||
			(s >= num0 && s <= num9) {
			continue
		} else {
			return false
		}
	}

	return true
}
