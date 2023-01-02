package cache

import (
	"testing"
	"time"
)

func TestTTLInMemoryCache_Get(t *testing.T) {
	c := NewTTLInMemoryCache()
	c.Set("key", "value", time.Second*4)
	time.Sleep(time.Second * 2)
	d, err := c.Get("key")
	if err != nil {
		t.Fail()
	}
	if d != "value" {
		t.Fail()
	}
	time.Sleep(time.Second * 2)
	d, err = c.Get("key")
	if err == nil {
		t.Fail()
	}
}

func TestTTLInMemoryCache_Set(t *testing.T) {
	c := NewTTLInMemoryCache()
	c.Set("key", "value", time.Second*4)
	time.Sleep(time.Second * 2)
	d, err := c.Get("key")
	if err != nil {
		t.Fail()
	}
	if d != "value" {
		t.Fail()
	}
	time.Sleep(time.Second * 2)
	d, err = c.Get("key")
	if err == nil {
		t.Fail()
	}
}

func TestInMemoryCache_Delete(t *testing.T) {
	c := NewTTLInMemoryCache()
	c.Set("key", "value", time.Second*4)
	time.Sleep(time.Second * 2)
	d, err := c.Get("key")
	if err != nil {
		t.Fail()
	}
	if d != "value" {
		t.Fail()
	}
	err = c.Delete("key")
	if err != nil {
		t.Fail()
	}
	d, err = c.Get("key")
	if err == nil {
		t.Fail()
	}
}
