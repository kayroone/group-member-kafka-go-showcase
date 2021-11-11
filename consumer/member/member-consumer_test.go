package main

import "testing"

func TestMessageToMembers(t *testing.T) {

	messageBytes := []byte("{\"Name\": \"foo\", \"Group\": \"bar\"}")
	want := Member{Name: "foo", Group: "bar"}

	got := MessageToMember(messageBytes)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
