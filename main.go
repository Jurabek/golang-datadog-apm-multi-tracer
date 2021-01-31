package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go/aws"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const tableName = "Items"

func main() {
	tracer.Start()
	defer tracer.Stop()

	awsRegion := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewEnvCredentials(),
			Region:      aws.String(awsRegion),
			Endpoint:    aws.String("http://dynamodb-local:8000"),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamodbSvs := dynamodb.New(awstrace.WrapSession(sess))

	r := muxtrace.NewRouter(muxtrace.WithServiceName("demo-app.router"))
	r.HandleFunc("/getItems", func(w http.ResponseWriter, r *http.Request) {
		result, err := dynamodbSvs.Scan(&dynamodb.ScanInput{TableName: aws.String(tableName)})
		if err != nil {
			log.Fatalf("DynamoDb GetItem failed: %v", err.Error())
			http.Error(w, "Not found", 404)
			return
		}

		if result.Items == nil {
			http.Error(w, "Not found", 404)
			return
		}

		var resultItems = make([]map[string]interface{}, len(result.Items))
		_ = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resultItems)

		body, err := json.Marshal(resultItems)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)

	}).Methods("GET")

	r.HandleFunc("/saveItem", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	log.Fatal(err)
}
