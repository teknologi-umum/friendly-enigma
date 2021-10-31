package business_test

import (
	"refrigerator/business"
	"testing"
)

func TestHasAccess(t *testing.T) {
	a := (1 << 1) | (1 << 2)
	r := business.HasAccess(a, business.UPDATE+business.READ)
	if !r {
		t.Error("should be true, got:", r)
	}
}

func TestGrantAccess(t *testing.T) {
	a := (1 << 0) | (1 << 3)
	r := business.GrantAccess(a, business.READ)
	if r != business.CREATE|business.READ|business.DELETE {
		t.Error("should be something gatau, got:", r)
	}
}

func TestRevokeAccess(t *testing.T) {
	a := (1 << 1) | (1 << 2)
	r := business.RevokeAccess(a, business.UPDATE)
	if r != business.READ {
		t.Error("should only be read, got:", r)
	}
}
