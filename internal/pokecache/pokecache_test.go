package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
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
		t.Errorf("expected to not find key")
		return
	}
}
```
Now that the basics pass, here are some ideas for extra tests that actually add value:

    Get on missing key
        Arrange: cache := NewCache(...)
        Act: val, ok := cache.Get("does-not-exist")
        Assert: !ok, and val == nil.

    Overwrite existing key
        Add a key with one value.
        Add the same key with a different value.
        Get should return the most recent value.

    Reap actually respects interval
        Use a very short interval (like 5 * time.Millisecond).
        Add two entries with a small delay:
            Add key "a".
            time.Sleep(interval / 2)
            Add key "b".
        Sleep interval more, then:
            "a" may be gone, "b" should still be present.
        This checks that you’re using createdAt correctly, not nuking everything too aggressively.

    Concurrency sanity check
        Create a cache with a somewhat larger interval.
        In a loop:
            go cache.Add(fmt.Sprintf("key-%d", i), []byte("val"))
        After a short sleep, iterate a few of those keys and ensure Get doesn’t panic and often returns ok == true.
        This is more about making sure your mutex usage is sound.

Pick one or two of these (no need to do all of them) and implement them in pokecache_test.go.

Would you like help sketching one specific test (e.g., the “overwrite” one) in more detail, or do you want to try writing one yourself first and then have me review it?
```