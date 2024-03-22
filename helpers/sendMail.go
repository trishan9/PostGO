package helpers

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(otpCode string, receiverName string, receiverEmail string) {
	m := mail.NewV3Mail()

	address := "mailtotrishan@gmail.com"
	name := "PostApp by Trishan"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.SetTemplateID("d-10caee8418f54ecca263f3825112b41e")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(receiverName, receiverEmail),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("code", otpCode)

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
