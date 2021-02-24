package account_test

import (
	"testing"

	"github.com/kevinschoon/bankocr/pkg/account"
)

func TestAccount(t *testing.T) {
	an := account.Number{6, 6, 4, 3, 7, 1, 4, 9, 5}
	if account.IsValid(an) {
		t.Fatalf("account number should be invalid")
	}
	an = account.Number{4, 5, 7, 5, 0, 8, 0, 0, 0}
	if !account.IsValid(an) {
		t.Fatalf("account number should be valid")
	}
	an = account.Number{0, 0, 0, 0, 0, 0, 0, 0, 0}
	if an.String() != "00000000" {
		t.Fatalf("expected 00000000, got %s", an.String())
	}
	an = account.Number{0, -1, 0, 0, 0, 0, 0, 0, 0}
	if an.String() != "0?000000 ILL" {
		t.Fatalf("expected 0?000000 ILL, got %s", an.String())
	}
}
