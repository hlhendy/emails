package messages

import (
	"errors"
	"net/http"

	"github.com/emails/models"
)

type Message interface {
	Send(*models.EmailRequest) error
}

func New(providerName, apiKey, baseURL string) (Message, error) {
	switch providerName {
	case "mandrill":
		return NewMandrill(apiKey, baseURL), nil
	case "mailgun":
		return NewMailgun(apiKey, baseURL), nil
	default:
		return nil, errors.New("must pass in valid email provider name")
	}
}

func NewRequest(method, url string, params *models.EmailRequest) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("to", params.RecipientEmail)
	q.Add("from", params.SenderEmail)
	q.Add("subject", params.Subject)
	q.Add("text", params.PlainTextBody)
	q.Add("html", params.HTMLBody)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
