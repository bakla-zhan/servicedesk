package senderchecker

import (
	"os"
	"testing"
)

func TestSenderChecker(t *testing.T) {
	os.Setenv("API_KEY", "a22bab1863650b7a4681b34b588c5af3b01c694c")
	os.Setenv("SECRET_KEY", "04956ac18d63b123d70fd563cd509af18390d0d3")

	testEmails := []string{
		"bob@mail.ru",
		"alice.brand@ibm.com",
		"vasiliy.pupkin@10minutemail.com",
		"dmitry_ivanov@rambler.ru",
		"igor@getairmail.com",
		"alex-williams@gmail.com",
		"valentina83@temp-mail.ru",
		"pfoqem@mailto.plus",
	}

	var (
		tempEmails = 4
		result     int
	)

	for _, email := range testEmails {
		err := SenderChecker(email)
		if err == errExpected {
			result++
		}
	}

	if result != tempEmails {
		t.Errorf("expected result: %v, got: %v", tempEmails, result)
	}
}
