package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

const defaultCertificateValidityTime = 90 * 24 * time.Hour

type options struct {
	appName                       string
	externalAPIPort               int
	internalAPIPort               int
	namespace                     string
	tokenLength                   int
	appTokenExpirationMinutes     int
	runtimeTokenExpirationMinutes int
	caSecretName                  string
	requestLogging                bool
	connectorServiceHost          string
	gatewayHost                   string
	certificateProtectedHost      string
	appsInfoURL                   string
	runtimesInfoURL               string
	certificateValidityTime       time.Duration
	central                       bool
}

type environment struct {
	country            string
	organization       string
	organizationalUnit string
	locality           string
	province           string
}

func parseArgs() *options {
	appName := flag.String("appName", "connector-service", "Name of the Certificate Service, used by k8s deployments and services.")
	externalAPIPort := flag.Int("externalAPIPort", 8081, "External API port.")
	internalAPIPort := flag.Int("internalAPIPort", 8080, "Internal API port.")
	namespace := flag.String("namespace", "kyma-integration", "Namespace used by Certificate Service")
	tokenLength := flag.Int("tokenLength", 64, "Length of a registration tokens.")
	appTokenExpirationMinutes := flag.Int("appTokenExpirationMinutes", 5, "Time to Live of application tokens expressed in minutes.")
	runtimeTokenExpirationMinutes := flag.Int("runtimeTokenExpirationMinutes", 10, "Time to Live of runtime tokens expressed in minutes.")
	caSecretName := flag.String("caSecretName", "nginx-auth-ca", "Name of the secret which contains root CA.")
	requestLogging := flag.Bool("requestLogging", false, "Flag for logging incoming requests.")
	connectorServiceHost := flag.String("connectorServiceHost", "cert-service.wormhole.cluster.kyma.cx", "Host at which this service is accessible.")
	gatewayHost := flag.String("gatewayHost", "gateway.wormhole.cluster.kyma.cx", "Host at which gateway service is accessible.")
	certificateProtectedHost := flag.String("certificateProtectedHost", "gateway.wormhole.cluster.kyma.cx", "Host secured with client certificate, used for certificate renewal.")
	appsInfoURL := flag.String("appsInfoURL", "", "URL at which management information is available.")
	runtimesInfoURL := flag.String("runtimesInfoURL", "", "URL at which management information is available.")
	certificateValidityTime := flag.String("certificateValidityTime", "90d", "Validity time of certificates issued by this service.")
	central := flag.Bool("central", false, "Determines whether connector works as the central")

	flag.Parse()

	validityTime, err := parseDuration(*certificateValidityTime)
	if err != nil {
		logrus.Infof("Failed to parse certificate validity time: %s, using default value.", err)
	}

	return &options{
		appName:                       *appName,
		externalAPIPort:               *externalAPIPort,
		internalAPIPort:               *internalAPIPort,
		namespace:                     *namespace,
		tokenLength:                   *tokenLength,
		appTokenExpirationMinutes:     *appTokenExpirationMinutes,
		runtimeTokenExpirationMinutes: *runtimeTokenExpirationMinutes,
		caSecretName:                  *caSecretName,
		requestLogging:                *requestLogging,
		connectorServiceHost:          *connectorServiceHost,
		gatewayHost:                   *gatewayHost,
		certificateProtectedHost:      *certificateProtectedHost,
		central:                       *central,
		appsInfoURL:                   *appsInfoURL,
		runtimesInfoURL:               *runtimesInfoURL,
		certificateValidityTime:       validityTime,
	}
}

func (o *options) String() string {
	return fmt.Sprintf("--appName=%s --externalAPIPort=%d --internalAPIPort=%d --namespace=%s --tokenLength=%d "+
		"--appTokenExpirationMinutes=%d --runtimeTokenExpirationMinutes=%d --caSecretName=%s --requestLogging=%t "+
		"--connectorServiceHost=%s --certificateProtectedHost=%s --gatewayHost=%s "+
		"--appsInfoURL=%s --runtimesInfoURL=%s --central=%t --certificateValidityTime=%s",
		o.appName, o.externalAPIPort, o.internalAPIPort, o.namespace, o.tokenLength,
		o.appTokenExpirationMinutes, o.runtimeTokenExpirationMinutes, o.caSecretName, o.requestLogging,
		o.connectorServiceHost, o.certificateProtectedHost, o.gatewayHost,
		o.appsInfoURL, o.runtimesInfoURL, o.central, o.certificateValidityTime)
}

func parseEnv() *environment {
	return &environment{
		country:            os.Getenv("COUNTRY"),
		organization:       os.Getenv("ORGANIZATION"),
		organizationalUnit: os.Getenv("ORGANIZATIONALUNIT"),
		locality:           os.Getenv("LOCALITY"),
		province:           os.Getenv("PROVINCE"),
	}
}

func (e *environment) String() string {
	return fmt.Sprintf("COUNTRY=%s ORGANIZATION=%s ORGANIZATIONALUNIT=%s LOCALITY=%s PROVINCE=%s",
		e.country, e.organization, e.organizationalUnit, e.locality, e.province)
}

func parseDuration(durationString string) (time.Duration, error) {
	unitsMap := map[string]time.Duration{"m": time.Minute, "h": time.Hour, "d": 24 * time.Hour}

	timeUnit := durationString[len(durationString)-1:]
	_, ok := unitsMap[timeUnit]
	if !ok {
		return defaultCertificateValidityTime, fmt.Errorf("unrecognized time unit provided: %s", timeUnit)
	}

	timeLength, err := strconv.Atoi(durationString[:len(durationString)-1])
	if err != nil {
		return defaultCertificateValidityTime, err
	}

	return time.Duration(timeLength) * unitsMap[timeUnit], nil
}
