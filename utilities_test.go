package main

import (
	"testing"
	"reflect"
)

func TestStrMapKeys(t *testing.T) {
	sample := map[string]string {
		"key_a": "0",
		"key_b": "1",
		"key_c": "2",
	}
	expected := []string{"key_a", "key_b", "key_c"}
	result := StrMapKeys(sample)

	eq := reflect.DeepEqual(expected, result)

	if !eq {
		t.Error("Map keys were not returned")
	}
}