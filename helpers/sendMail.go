package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(otpCode int, receiverName string, receiverEmail string) {
	var body bytes.Buffer
	t, err := template.ParseFiles("./assets/email_template.html")

	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	t.Execute(&body, struct {
		Code string
	}{Code: "991991"})

	from := mail.NewEmail("PostApp by Trishan", "mailtotrishan@gmail.com")
	subject := "Verify your account!"
	to := mail.NewEmail(receiverName, receiverEmail)
	message := mail.NewSingleEmail(from, subject, to, body.String(), body.String())
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
