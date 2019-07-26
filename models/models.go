package models

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

const (
	MaxEmailAddressLength = 254
	MaxSubjectLineLength  = 78
)

type EmailRequest struct {
	RecipientEmail string `json:"to"`
	RecipientName  string `json:"to_name"`
	SenderEmail    string `json:"from"`
	SenderName     string `json:"from_name"`
	Subject        string `json:"subject"`
	HTMLBody       string `json:"body"`

	// PlainTextBody stores the body of the email with HTML stripped out
	PlainTextBody string `json:"plain_text_body,omitempty"`
}

// Validate does some basic checks on the data and populates the PlainTextBody field
func (e *EmailRequest) Validate() error {

	if len(e.RecipientEmail) > MaxEmailAddressLength {
		return errors.New("INVALID_ARG_TO_EXCEEDS_MAX_LENGTH")
	}

	if len(e.SenderEmail) > MaxEmailAddressLength {
		return errors.New("INVALID_ARG_FROM_EXCEEDS_MAX_LENGTH")
	}

	if len(e.Subject) > MaxSubjectLineLength {
		return errors.New("INVALID_ARG_SUBJECT_EXCEEDS_MAX_LENGTH")
	}

	e.PlainTextBody = getPlainText(e.HTMLBody)

	return nil

}

func getPlainText(htmlBody string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlBody))
	prevToken := tokenizer.Token()

	var plainText string

tokenizerLoop:
	for {
		t := tokenizer.Next()
		switch {
		case t == html.ErrorToken:
			break tokenizerLoop // end of html
		case t == html.TextToken:
			if prevToken.Data == "script" {
				continue
			}
			text := html.UnescapeString(string(tokenizer.Text()))
			plainText += strings.TrimSpace(text)
		}
	}
	return plainText
}
