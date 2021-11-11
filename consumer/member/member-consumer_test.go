package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageToMembersValidJson(t *testing.T) {

	validMessageBytes := []byte("{\"Name\": \"foo\", \"Group\": \"bar\"}")
	got := MessageToMember(validMessageBytes)
	want := Member{Name: "foo", Group: "bar"}

	assert.Equal(t, got, want)
}

func TestMessageToMembersInvalidJson(t *testing.T) {

	invalidMessageBytes := []byte("{\"Name\": \"foo\", \"Group\": \"bar\"")

	assert.Panics(t, func() { MessageToMember(invalidMessageBytes) }, "MessageToMember did not panic")
}
