package main

import (
	"testing"
)

func TestLookupRunfile(t *testing.T) {
	_, err := LookupRunfile()
	if err != nil {
		t.Errorf("%s", err)
	}
}