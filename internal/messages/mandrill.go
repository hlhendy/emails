package messages

import (
	"errors"

	"github.com/emails/models"
)

type Mandrill struct {
	BaseURL string `json:"base_url"`
	APIKey  string `json:"sending_domain"`
}

func NewMandrill(apiKey, baseURL string) *Mandrill {
	return &Mandrill{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (m Mandrill) Send(e *models.EmailRequest) error {
	return errors.New("not yet implemented")
}
