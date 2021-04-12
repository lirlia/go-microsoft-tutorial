package main

import (
	// https://github.com/sirupsen/logrus
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.WithFields(logger.Fields{
		"animal": "dog",
	}).Info("dog is here")
}

/* Output
% go run logrus_test1.go
INFO[0000] dog is here                                   animal=dog
*/
