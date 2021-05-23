package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/eaneto/notify"
	sns "github.com/eaneto/notify/service/amazonsns"
	sqs "github.com/eaneto/notify/service/amazonsqs"
	"github.com/sirupsen/logrus"
)

func sqsHandler(w http.ResponseWriter, r *http.Request) {
	notifier := notify.New()

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")
	sqsService, err := sqs.New(accessKey, secretKey, region)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	account := os.Getenv("AWS_ACCOUNT")
	baseUrl := "https://sqs.%s.amazonaws.com/%s/notification-queue"
	sqsService.AddReceivers(fmt.Sprintf(baseUrl, region, account))

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
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")
	topic := os.Getenv("AWS_SNS_TOPIC")
	account := os.Getenv("AWS_ACCOUNT")

	snsService, err := sns.New(accessKey, secretKey, region)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	baseTopic := "arn:aws:sns:%s:%s:%s"
	snsService.AddReceivers(fmt.Sprintf(baseTopic, region, account, topic))
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
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/publish-sqs", sqsHandler)
	http.HandleFunc("/publish-sns", snsHandler)
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		logrus.Error("Failed to Listen and Serve")
		return
	}
}
