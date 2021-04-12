package main

import (
	// https://github.com/sirupsen/logrus
	"fmt"

	log "github.com/sirupsen/logrus"
)

// 特殊な関数、main関数よりも前に呼び出される
func init() {
	fmt.Println("call init function")
}

func main() {
	// JSON化すると time/levelが勝手につく

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("dog is here")

}

/* Output

call init function
{"animal":"dog","level":"info","msg":"dog is here","time":"2021-04-13T01:27:23+09:00"}

*/
