package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emails/models"
)

type Mailgun struct {
	BaseURL string
	APIKey  string
}

func NewMailgun(apiKey, baseURL string) *Mailgun {
	return &Mailgun{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (m Mailgun) buildURL() string {
	return fmt.Sprintf("https://api:%s@%s/messages", m.APIKey, m.BaseURL)
}

func (m Mailgun) Send(e *models.EmailRequest) error {
	hc := http.Client{}
	req, err := NewRequest("POST", m.buildURL(), e)
	if err != nil {
		return err
	}

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type mailgunResp struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}

	var mgResp mailgunResp
	if err := json.Unmarshal(body, &mgResp); err != nil {
		return err
	}

	if mgResp.Message != "Queued. Thank you." {
		return fmt.Errorf(fmt.Sprintf("mailgun error: %s", mgResp.Message))
	}

	return nil
}
