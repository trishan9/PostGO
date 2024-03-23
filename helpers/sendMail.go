package helpers

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(otpCode string, receiverName string, receiverEmail string) error {
	m := mail.NewV3Mail()

	address := "mailtotrishan@gmail.com"
	name := "PostApp by Trishan"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)

	m.SetTemplateID(os.Getenv("EMAIL_TEMPLATE_ID"))

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
	_, err := sendgrid.API(request)

	if err != nil {
		return err
	}
	return nil
}
