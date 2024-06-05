package firebase

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./firebase/serviceAccountKeys.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Message client
	client, _ := app.Messaging(ctx)
	return app, ctx, client
}

func SendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	registrationToken := "d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification test",
			Body:  "Hello Huan",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
