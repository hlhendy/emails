package testclient

import "github.com/emails/models"

type TestClient struct {
	SendError       error
	ValidationError error
}

func (t *TestClient) Send(er *models.EmailRequest) error {
	return t.SendError
}

func (t *TestClient) Validation(er *models.EmailRequest) error {
	return t.ValidationError
}
