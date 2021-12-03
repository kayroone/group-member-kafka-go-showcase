package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageToMembersValidJson(t *testing.T) {

	validMessageBytes := []byte("{\"Name\": \"foo\", \"Group\": \"bar\"}")
	got, err := MessageToMember(validMessageBytes)
	want := Member{Name: "foo", Group: "bar"}

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestMessageToMembersInvalidJson(t *testing.T) {

	invalidMessageBytes := []byte("{\"Name\": \"foo\", \"Group\": \"bar\"")
	_, err := MessageToMember(invalidMessageBytes)

	assert.NotNil(t, err)
}
