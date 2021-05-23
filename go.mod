module gitihub.com/lab-end-soft-p2/example-app

go 1.16

require (
	github.com/aws/aws-sdk-go-v2/config v1.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.4.1 // indirect
	github.com/eaneto/notify v0.18.0 // indirect
	github.com/nikoksr/notify v0.16.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
)

replace github.com/nikiksr/notify v0.16.1 => github.com/eaneto/notify v0.17.1

replace github.com/nikiksr/notify/amazonsqs v0.16.1 => github.com/eaneto/notify/amazonsqs v0.17.1
