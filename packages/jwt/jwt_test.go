package jwt_test

import (
	"refrigerator/business"
	"refrigerator/packages/jwt"
	"testing"
)

var member = business.Member{
	ID: "169a88f9-fb2f-4121-b012-dde7796cd195",
	Permission: business.CREATE | business.READ,
	Name: "Test",
}

var secret = []byte("shh this is a secret")

func TestJWT(t *testing.T) {
	token, err := jwt.GenerateJWT(secret, member)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if token == "" {
		t.Error("token should not be empty")
	}

	m, err := jwt.VerifyJWT(secret, token)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if m.ID != member.ID {
		t.Error("member id value is not equal, got:", m.ID)
	}
	
	if m.Permission != member.Permission {
		t.Error("member permission value is not equal, got:", m.Permission)
	}
}