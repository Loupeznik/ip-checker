package main

import (
	"log"
	"os"
  	"net/http"
	"io"
	"io/ioutil"
  
 	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
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
		from := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL"))
		subject := "Server's IP has changed"
		to := mail.NewEmail(os.Getenv("RECEPIENT_EMAIL"), os.Getenv("RECEPIENT_EMAIL"))
		plainTextContent := "The server's IP has changed to " + actualIp + ". Update your DNS records."
		htmlContent := "The server's IP has changed to <strong>" + actualIp + "</strong>. Update your DNS records."
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		
		// Send email
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		
		if err != nil {
			log.Println(err)
		} else {
			log.Println(response.Headers)
			
			// Write new IP address to a file
			err := ioutil.WriteFile("ip.txt", []byte(actualIp), 0644)
			if err != nil {
				log.Fatal(err)
			}
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
