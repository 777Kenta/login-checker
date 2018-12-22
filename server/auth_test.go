package main

import (
	"login-checker/server/model"
	"testing"
)

func TestEncodePassword(t *testing.T) {
	password := "1111"
	encodedPassword := model.EncodeStringMD5(password)
	if encodedPassword != "ab59c67bf196a4758191e42f76670ceba" {
		t.Error("unValid Password :", password)
	}
}
