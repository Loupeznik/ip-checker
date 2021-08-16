package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/slack-go/slack"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("No command line arguments supplied")
		fmt.Println("Usage: go run . --slack or go run . --email")
	}

	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the currently saved IP address from a file
	ipFile, err := ioutil.ReadFile("ip.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Assing IP variables
	currentIp := string(ipFile)
	actualIp := getIp()

	// Main logic
	if currentIp != actualIp {
		if os.Args[1] == "--email" {
			notifyByEmail(currentIp, actualIp)
		} else if os.Args[1] == "--slack" {
			notifyBySlack(currentIp, actualIp)
		} else {
			log.Fatal("Invalid flag")
		}
	}
}

// Get the actual IP address of the machine
func getIp() string {
	resp, err := http.Get("https://api64.ipify.org")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

// Write new IP address to a file
func writeIp(ip string) {
	err := ioutil.WriteFile("ip.txt", []byte(ip), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Send notification via email
func notifyByEmail(oldIp string, newIp string) {
	from := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL"))
	subject := "Server's IP has changed"
	to := mail.NewEmail(os.Getenv("RECEPIENT_EMAIL"), os.Getenv("RECEPIENT_EMAIL"))
	plainTextContent := printMessage(oldIp, newIp)
	htmlContent := "The IP of a server originally at <strong>" + oldIp + "</strong> has <strong>" + newIp + "</strong>. Update your DNS records."
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Notification was send to " + os.Getenv("RECEPIENT_EMAIL"))
		writeIp(newIp)
	}
}

// Send notification via slack
func notifyBySlack(oldIp string, newIp string) {
	api := slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))

	_, _, err := api.PostMessage(
		os.Getenv("SLACK_CHANNEL_ID"),
		slack.MsgOptionText(printMessage(oldIp, newIp), false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	writeIp(newIp)
	log.Printf("Notification was send to Slack")
}

func printMessage(oldIp string, newIp string) string {
	return "The IP of a server - " + os.Getenv("HOSTNAME") + " - originally at " + oldIp + " has changed to " + newIp + ". Update your DNS records."
}
