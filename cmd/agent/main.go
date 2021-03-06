package main

import (
	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/wolfeidau/buildkite-serverless-agent/pkg/agent"
	"github.com/wolfeidau/buildkite-serverless-agent/pkg/config"
	"github.com/wolfeidau/buildkite-serverless-agent/pkg/registration"
)

func main() {
	logrus.AddHook(filename.NewHook())
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := config.New()
	if err != nil {
		log.WithError(err).Fatal("failed to load configuration")
	}

	sess := session.Must(session.NewSession())

	bkw := agent.New(cfg, sess)

	err = xray.Configure(xray.Config{LogLevel: "info"})
	if err != nil {
		log.WithError(err).Fatal("failed to xray configuration")
	}

	regService := registration.New(cfg, sess)

	_, err = regService.RegisterAgent()
	if err != nil {
		log.WithError(err).Fatal("failed to xray configuration")
	}

	lambda.Start(bkw.Handler)
}
