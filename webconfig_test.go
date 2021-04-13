package main

import "testing"

func TestApiVersion(t *testing.T) {
	apiVersion := getApiVersion()
	if apiVersion != "v1" {
		t.Error("Api version should be v1")
	}
}
