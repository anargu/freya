package main

import (
	"context"

	// freya "../freyacon/go"
	// freya "../freyacon/go"

	"io/ioutil"
	"log"

	freya "github.com/anargu/freya/freyacon/go"
	"google.golang.org/grpc"
)

func main() {
	endpoint := "127.0.0.1:10000"
	client, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	freyaClient := freya.NewFreyaClient(client)
	welcomeData, err := ioutil.ReadFile("/Users/macBook/Documents/welcome_template.html.txt")
	if err != nil {
		panic(err)
	}

	resp, err := freyaClient.SaveNewTemplate(context.Background(), &freya.TemplateData{
		TemplateName: "welcome_mail",
		Data:         welcomeData,
	})

	if err != nil {
		panic(err)
	}

	if resp.Saved {
		log.Println("Template saved")
	}

}
