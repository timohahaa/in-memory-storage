package store

import (
	"testing"
	"time"
)

func TestSetGetKey(t *testing.T) {
	store := GetStore()
	store.Set("key1", "val1", time.Second*3)
	store.Set("key2", "val2", time.Second*3)

	got1, err := store.Get("key1")
	if err != nil {
		t.Errorf("Should not have got an error.")
	}
	if got1 != "val1" {
		t.Errorf("Incorrect result. Got: %s, expected: val1", got1)
	}

	got2, err := store.Get("key2")
	if err != nil {
		t.Errorf("Should not have got an error.")
	}
	if got2 != "val2" {
		t.Errorf("Incorrect result. Got: %s, expected: val2", got2)
	}
}

func TestTTL(t *testing.T) {
	store := GetStore()
	store.Set("key", "value", time.Second*3)
	time.Sleep(time.Second * 4)
	_, err := store.Get("key")
	if err == nil {
		t.Errorf("Should have gotten non-nil error.")
	}
}
