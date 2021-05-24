package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/eaneto/notify"
	sns "github.com/eaneto/notify/service/amazonsns"
	sqs "github.com/eaneto/notify/service/amazonsqs"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	AWS_REGION = "us-east-1"
	QUEUE_NAME = "notification-queue"
	TOPIC_NAME = "SMSmessage"
)

var accessKey, secretKey, account string

func sqsHandler(w http.ResponseWriter, r *http.Request) {
	notifier := notify.New()
	sqsService, err := sqs.New(accessKey, secretKey, AWS_REGION)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	baseUrl := "https://sqs.%s.amazonaws.com/%s/%s"
	sqsService.AddReceivers(fmt.Sprintf(
		baseUrl, AWS_REGION, account, QUEUE_NAME,
	))

	notifier.UseServices(sqsService)

	err = notifier.Send(
		context.Background(),
		"Subject",
		"Example message",
	)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		logrus.Info("Message published succesfully")
		w.WriteHeader(http.StatusOK)
	}
}

func snsHandler(w http.ResponseWriter, r *http.Request) {
	notifier := notify.New()
	snsService, err := sns.New(accessKey, secretKey, AWS_REGION)
	fmt.Println(accessKey, secretKey)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	baseTopic := "arn:aws:sns:%s:%s:%s"
	snsService.AddReceivers(fmt.Sprintf(
		baseTopic, AWS_REGION, account, TOPIC_NAME,
	))
	notifier.UseServices(snsService)

	err = notifier.Send(
		context.Background(),
		"Subject",
		"Example message")

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		logrus.Info("Message published succesfully")
		w.WriteHeader(http.StatusOK)
	}
}

// healthHandler returns ok always, used by nagios.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Health Check OK")
	w.WriteHeader(http.StatusOK)
}

func main() {
	flag.StringVar(&accessKey, "accessKey", "", "")
	flag.StringVar(&secretKey, "secretKey", "", "")
	flag.StringVar(&account, "account", "", "")
	flag.Parse()
	fmt.Println(accessKey, secretKey)
	http.HandleFunc("/publish-sqs", sqsHandler)
	http.HandleFunc("/publish-sns", snsHandler)
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		logrus.Error("Failed to Listen and Serve")
		return
	}
}
