package mail

import "testing"

func TestIsValidOkToName(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkToName(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkFromName(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkFromName(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkToEmail(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkToEmail(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkFromEmail(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkFromEmail(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME PLATFORM TEAM", Addr: "joseedwin84gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkSubject(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkSubject(t *testing.T) {
	var tooLongSubject string
	for i := 0; i < 79; i++ {
		tooLongSubject += "e"
	}
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME PLATFORM TEAM", Addr: "joseedwin84@gmail.com"},
		Subject: tooLongSubject,
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkBody(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkBody(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME PLATFORM TEAM", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte(""),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkClient(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkClient(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME PLATFORM TEAM", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  12,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsValidOkAuthKey(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME_PLATFORM Team", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "awesome key",
	}
	if err := m.isValid(); err != nil {
		t.Errorf("should return: nil but got: %s", err)
	}
}

func TestIsValidNotOkAuthKey(t *testing.T) {
	m := Mail{
		To:      Acct{Name: "Edwin J", Addr: "joseedwin84@gmail.com"},
		From:    Acct{Name: "AWESOME PLATFORM TEAM", Addr: "joseedwin84@gmail.com"},
		Subject: "hello world email",
		Body:    []byte("hello:D"),
		Client:  SendgridClient,
		AuthKey: "",
	}
	if err := m.isValid(); err == nil {
		t.Error("should return: err but got: nil")
	}
}

func TestIsAccountValidEmptyName(t *testing.T) {
	if err := isNameValid(""); err != nil {
		if err.Error() != "cannot be empty: " {
			t.Errorf("should return: 'err empty' but got: %#v", err)
		}
		return
	}
	t.Error("should return: 'err empty' but got: nil")
}

func TestIsAccountValidNameTooLong(t *testing.T) {
	var tooLongName string
	for i := 0; i < 73; i++ {
		tooLongName += "e"
	}
	if err := isNameValid(tooLongName); err != nil {
		if err.Error() != "cannot be longer than 72 chars: "+tooLongName {
			t.Errorf("should return: 'err too long' but got: %#v", err)
		}
		return
	}
	t.Error("should return: 'err too long' but got: nil")
}

func TestIsAccountValidNameValid(t *testing.T) {
	if err := isNameValid("Edwin"); err != nil {
		t.Errorf("should return: nil  but got: %#v", err)
	}
}

func TestIsSubjectValidEmpty(t *testing.T) {
	err := isSubjectValid("")
	if err != nil {
		if err.Error() != "cannot be empty: " {
			t.Errorf("should return: 'err empty' but got: %#v", err)
		}
		return
	}
	t.Error("should return: 'err empty' but got: nil")
}

func TestIsSubjectValidTooLong(t *testing.T) {
	var tooLongSubject string
	for i := 0; i < 79; i++ {
		tooLongSubject += "e"
	}
	err := isSubjectValid(tooLongSubject)
	if err != nil {
		if err.Error() != "cannot be longer than 78 chars: "+tooLongSubject {
			t.Errorf("should return: 'err too long' but got: %#v", err)
		}
		return
	}
	t.Error("should return: 'err too long' but got: nil")
}

func TestIsSubjectValidOkSubject(t *testing.T) {
	err := isSubjectValid("Very Cool Email Subject")
	if err != nil {
		t.Errorf("should return: nil but got: %#v", err)
	}
}

func TestIsClientValid(t *testing.T) {
	err := isClientValid(10)
	if err != nil {
		if err.Error() != "not valid: 10" {
			t.Errorf("should return: 'invalid client' but got: %#v", err)
		}
		return
	}
	t.Error("should return: 'err invalid client' but got: nil")
}

func TestIsClientValidOkClient(t *testing.T) {
	err := isClientValid(SendgridClient)
	if err != nil {
		t.Errorf("should return: nil but got: %#v", err)
	}
}
