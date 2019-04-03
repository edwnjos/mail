package mail

import (
	"fmt"

	"github.com/badoux/checkmail"
)

const (
	// SendgridClient is used by the user to choose sengrid as the email client
	SendgridClient = 1 + iota
)

// Acct represents an email account
type Acct struct {
	Name string
	Addr string
}

// Mail is the main struct of the package which contains everything
// needed to send an email
type Mail struct {
	To      Acct
	From    Acct
	Subject string
	Body    []byte
	Client  int
	AuthKey string
}

// Send calls the send client
func (m *Mail) Send() error {
	if err := m.isValid(); err != nil {
		return fmt.Errorf("mail validation failed: %s", err)
	}
	switch m.Client {
	case SendgridClient:
		body, err := m.parseSengridRequestBody()
		if err != nil {
			return fmt.Errorf("failed to create sendgrid request body: %s", err)
		}
		if err := sendSengridRequest(m.AuthKey, body); err != nil {
			return err
		}
	}
	return nil
}

func (m *Mail) isValid() error {
	if err := isNameValid(m.To.Name); err != nil {
		return fmt.Errorf("`to` name not valid: %s", err.Error())
	}
	if err := isEmailValid(m.To.Addr); err != nil {
		return fmt.Errorf("`to` email not valid: %s", err.Error())
	}

	if err := isNameValid(m.From.Name); err != nil {
		return fmt.Errorf("`from` name not valid: %s", err.Error())
	}

	if err := isSubjectValid(m.Subject); err != nil {
		return fmt.Errorf("subject not valid: %s", err.Error())
	}

	if err := isClientValid(m.Client); err != nil {
		return fmt.Errorf("client not valid: %s", err.Error())
	}

	if len(m.Body) == 0 {
		return fmt.Errorf("body not valid: cannot be empty: %#v", m.Body)
	}
	if m.AuthKey == "" {
		return fmt.Errorf("`auth_key` not valid: cannot be empty: %s", m.AuthKey)
	}
	return nil
}

func isNameValid(n string) error {
	if n == "" {
		return fmt.Errorf("cannot be empty: %s", n)
	}
	if len(n) > 72 {
		return fmt.Errorf("cannot be longer than 72 chars: %s", n)
	}
	return nil
}

func isEmailValid(e string) error {
	if err := checkmail.ValidateFormat(e); err != nil {
		return err
	}
	return checkmail.ValidateHost(e)
}

func isSubjectValid(s string) error {
	if s == "" {
		return fmt.Errorf("cannot be empty: %s", s)
	}
	if len(s) > 78 {
		return fmt.Errorf("cannot be longer than 78 chars: %s", s)
	}
	return nil
}

func isClientValid(c int) error {
	switch c {
	case SendgridClient:
		return nil
	}
	return fmt.Errorf("not valid: %d", c)
}
