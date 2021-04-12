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

	// loggingでpanicやexit 1を制御できるのかなるほど便利。
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

	log.WithFields(log.Fields{
		"animal": "dog",
	}).Info("dog is here")

}

/* Output
% go run logrus_test3.go
call init function
INFO[0000] Something noteworthy happened!
WARN[0000] You should probably take a look at this.
ERRO[0000] Something failed but I'm not quitting.
PANI[0000] I'm bailing.
panic: (*logrus.Entry) 0xc0000b8230

goroutine 1 [running]:
github.com/sirupsen/logrus.(*Entry).log(0xc0000b8070, 0x0, 0xc0000ac0a0, 0xc)
        /Users/xxx/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:259 +0x2e5
github.com/sirupsen/logrus.(*Entry).Log(0xc0000b8070, 0xc000000000, 0xc0000c9dd8, 0x1, 0x1)
        /Users/xxx/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:293 +0x86
github.com/sirupsen/logrus.(*Logger).Log(0xc0000b8000, 0x0, 0xc0000c9dd8, 0x1, 0x1)
        /Users/xxx/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:198 +0x7e
github.com/sirupsen/logrus.(*Logger).Panic(...)
        /Users/xxx/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:247
github.com/sirupsen/logrus.Panic(...)
        /Users/xxx/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/exported.go:129
main.main()
        /Users/xxx/Cording/go-microsoft-tutorial/src/logrus_test/logrus_test3.go:25 +0x239
exit status 2
*/
