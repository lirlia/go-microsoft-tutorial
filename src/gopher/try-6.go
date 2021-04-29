package main

type UserResult struct {
	amountResult int
	gameResult   []int
}

// 何を対象にされているかを考える
// 今回の場合は1回のゲームスコアを表現する

type Score struct {
	GameID int
	UserID int
	Point  int
}
