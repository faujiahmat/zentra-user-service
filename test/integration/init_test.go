package integration_test

import (
	"os"

	"github.com/faujiahmat/zentra-user-service/src/common/log"
	"github.com/sirupsen/logrus"
)

func init() {
	err := os.Chdir(os.Getenv("ZENTRA_USER_SERVICE_WORKSPACE"))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "integration_test.init", "section": "os.Chdir"}).Fatal(err)
	}
}
