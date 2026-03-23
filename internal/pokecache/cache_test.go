package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		name string
		key  string
		val  []byte
	}{
		{
			name: "baseURL key",
			key:  "https://example.com",
			val:  []byte("testdata"),
		},
		{
			name: "baseURL + path key",
			key:  "https://example.com/path",
			val:  []byte("moretestdata"),
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Test case %s", tc.name), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(tc.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected not to find key")
		return
	}
}
