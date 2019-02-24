package main

import (
	"context"
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

	in := &freya.SendEmailParams{
		Subject:      "Hello",
		Params:       map[string]string{},
		TemplateName: "welcome_mail",
		To:           map[int32]string{0: "bregy.malpartida@utec.edu.pe", 1: "mateo@bombo.pe"},
		Attachment:   map[string]string{"0": "", "1": ""},
	}

	res, err := freyaClient.SendEmail(context.Background(), in)

	if err != nil {
		panic(err)
	}

	log.Printf("%v\n", res)

}
