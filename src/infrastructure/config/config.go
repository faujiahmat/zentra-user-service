package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/faujiahmat/zentra-user-service/src/common/log"
	"github.com/sirupsen/logrus"
)

type currentApp struct {
	GrpcPort string
}

type apiGateway struct {
	BaseUrl           string
	BasicAuth         string
	BasicAuthUsername string
	BasicAuthPassword string
}

type postgres struct {
	Url      string
	Dsn      string
	User     string
	Password string
}

type Config struct {
	CurrentApp *currentApp
	Postgres   *postgres
	ApiGateway *apiGateway
}

var Conf *Config

// *config ini hanya berisi env variable
func init() {
	appStatus := os.Getenv("ZENTRA_APP_STATUS")

	if appStatus == "DEVELOPMENT" {
		Conf = setUpForDevelopment()
		return
	}

}

func loadRSAPrivateKey(privateKey string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "pem.Decode",
		}).Fatal("failed to parse pem block containing the key")
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "x509.ParsePKCS1PrivateKey",
		}).Fatalf("failed to parse rsa private key: %v", err)
	}

	return rsaPrivateKey
}

func loadRSAPublicKey(publicKey string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "config.loadRSAPrivateKey",
			"section":  "pem.Decode",
		}).Fatal("failed to parse pem block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"location": "loadRSAPublicKey",
			"section":  "x509.ParsePKCS1PublicKey",
		}).Fatalf("failed to parse rsa public key: %v", err)
	}

	rsaPublicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Logger.WithFields(logrus.Fields{
			"location": "loadRSAPublicKey",
			"section":  "type assertion",
		}).Fatal("failed to assert type to *rsa.PublicKey")
	}

	return rsaPublicKey
}
