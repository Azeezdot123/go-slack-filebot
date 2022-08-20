package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main(){
	// Load environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	// put the name of the file you want to upload int string slice
	fileArr := []string{"upload.txt"}

	// checking if the number of file is more than 3
	if len(fileArr) > 3 {
		fmt.Println("The file limit per user have been reached")
	}else{
		for i := 0; i < len(fileArr); i++{
			params := slack.FileUploadParameters{
				Channels: channelArr,
				File: fileArr[i],
			}
			file, err := api.UploadFile(params)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
		}
	}
}