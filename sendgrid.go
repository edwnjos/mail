package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type sendgridAcct struct {
	Name string `json:"name"`
	Addr string `json:"email"`
}

type sendgridContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type sendgridPersonalizations struct {
	To      []sendgridAcct `json:"to"`
	Subject string         `json:"subject"`
}
type sendgridMail struct {
	Personalizations []sendgridPersonalizations `json:"personalizations"`
	From             sendgridAcct               `json:"from"`
	Content          []sendgridContent          `json:"content"`
}

func (m *Mail) parseSengridRequestBody() (*bytes.Buffer, error) {
	sm := sendgridMail{
		Personalizations: []sendgridPersonalizations{
			sendgridPersonalizations{
				To: []sendgridAcct{
					sendgridAcct{Name: m.To.Name, Addr: m.To.Addr},
				},
				Subject: m.Subject,
			},
		},
		From:    sendgridAcct{Name: m.From.Name, Addr: m.From.Addr},
		Content: []sendgridContent{sendgridContent{Type: "text/html", Value: string(m.Body)}},
	}

	body, err := json.Marshal(sm)
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %s", err)
	}

	return bytes.NewBuffer(body), nil
}

func sendSengridRequest(authKey string, b *bytes.Buffer) error {
	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", b)
	if err != nil {
		return fmt.Errorf("failed to create http request: %s", err)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authKey)

	rs, err := (&http.Client{Timeout: time.Second * 6}).Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %s", err)
	}
	defer rs.Body.Close()

	if rs.StatusCode != http.StatusAccepted {
		return fmt.Errorf("request returned non ok status: %d", rs.StatusCode)
	}
	return nil
}
