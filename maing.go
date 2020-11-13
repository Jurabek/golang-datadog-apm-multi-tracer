package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start()
	defer tracer.Stop()

	awsRegion := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewEnvCredentials(),
			Region:      aws.String(awsRegion),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))
	_ = dynamodb.New(awstrace.WrapSession(sess))

	r := muxtrace.NewRouter(muxtrace.WithServiceName("demo-app.router"))
	r.HandleFunc("/getItems", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("GET")
}
