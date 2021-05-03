package main

import (
	"testing"
)

func TestMail(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		// メールアドレスは「**ローカルパート@ドメインパート**」で構成される
		{name: "unit test", input: "aa", expected: false},
		{name: "unit test", input: "aa@@aa", expected: false},
		// ローカルパート/ドメインパートはそれぞれ「**2文字以上50文字以下**」で構成される
		{name: "unit test", input: "a@a", expected: false},
		{name: "unit test", input: "aa@a", expected: false},
		{name: "unit test", input: "a@aa", expected: false},
		{name: "unit test", input: "aa@aa", expected: true},
		{name: "unit test", input: "aaaaa@aaaaa", expected: true},
		{name: "unit test", input: "aaaaaa@a", expected: false},
		{name: "unit test", input: "a@aaaaaa", expected: false},
		// ローカルパート/ドメインパートでは「**半角小文字英数字と記号のみ**」で構成される
		{name: "unit test", input: "1.a@z-4_5", expected: true},
		{name: "unit test", input: "あa@aa", expected: false},
		{name: "unit test", input: "Aa@aa", expected: false},
		{name: "unit test", input: "ａa@aa", expected: false},
		{name: "unit test", input: "０a@aa", expected: false},
		{name: "unit test", input: "aa@あa", expected: false},
		{name: "unit test", input: "aa@Aa", expected: false},
		{name: "unit test", input: "aa@０a", expected: false},
		{name: "unit test", input: "aa@ａa", expected: false},
		// 使用可能な記号は「**-_.のみ**」とする
		{name: "unit test", input: "a,a@aa", expected: false},
		{name: "unit test", input: "a¥a@aa", expected: false},
		{name: "unit test", input: "a\"a@aa", expected: false},
		{name: "unit test", input: "a!a@aa", expected: false},
		{name: "unit test", input: "a\a@aa", expected: false},
		{name: "unit test", input: "a$a@aa", expected: false},
		{name: "unit test", input: "a%a@aa", expected: false},
		{name: "unit test", input: "a&a@aa", expected: false},
		{name: "unit test", input: "a'a@aa", expected: false},
		{name: "unit test", input: "a(a@aa", expected: false},
		{name: "unit test", input: "a)a@aa", expected: false},
		{name: "unit test", input: "a=a@aa", expected: false},
		{name: "unit test", input: "a~a@aa", expected: false},
		// **記号は連続して利用できない**ものとする(@の前後も不可とする)
		{name: "unit test", input: "a..a@aa", expected: false},
		{name: "unit test", input: "aa.@aa", expected: false},
		{name: "unit test", input: "aa@.aa", expected: false},
		// ローカルパート/ドメインパートの**最初と最後に記号は使用できない**ものとする
		{name: "unit test", input: ".aa@aa", expected: false},
		{name: "unit test", input: "aa@aa.", expected: false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := chkAddress(c.input); c.expected != actual {
				t.Errorf("want chkAddress(%v) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
