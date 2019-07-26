package messages

import (
	"testing"
)

func TestNew(t *testing.T) {
	var testcases = []struct {
		description     string
		providerName    string
		apiKey          string
		baseURL         string
		expectedErr     string
		expectedMessage Message
	}{
		{
			"invalid provider",
			"test",
			"fake_key",
			"fake_base_url",
			"must pass in valid email provider name",
			nil,
		},
		{
			"success - mailgun",
			"mailgun",
			"fake_key",
			"fake_base_url",
			"",
			&Mailgun{
				BaseURL: "fake_base_url",
				APIKey:  "fake_key",
			},
		},
		{
			"success - mandrill",
			"mandrill",
			"fake_key",
			"fake_base_url",
			"",
			&Mandrill{
				BaseURL: "fake_base_url",
				APIKey:  "fake_key",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := New(tc.providerName, tc.apiKey, tc.baseURL)
			if err != nil {
				if err.Error() != tc.expectedErr {
					t.Errorf("expected error was: %q, but actual error was: %q", tc.expectedErr, err)
				}
			}
		})
	}
}
