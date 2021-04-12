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

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("dog is here")

}

/* Output
% go run logrus_test2.go
call init function
INFO[0000] Something noteworthy happened!
WARN[0000] You should probably take a look at this.
ERRO[0000] Something failed but I'm not quitting.
FATA[0000] Bye.
exit status 1
*/
