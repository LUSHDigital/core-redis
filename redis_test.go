package coreredis_test

import (
	"os"
	"reflect"
	"testing"

	coreredis "github.com/LUSHDigital/core-redis"
	"github.com/go-redis/redis"
)

func TestURLFromEnv(t *testing.T) {
	a := coreredis.URLFromEnv()
	equals(t, coreredis.DefaultURL, a)
	os.Setenv("REDIS_URL", "redis://example.com:6379/0")
	b := coreredis.URLFromEnv()
	equals(t, "redis://example.com:6379/0", b)
}

func TestParse(t *testing.T) {
	cases := []struct {
		url      string
		expected *redis.Options
	}{
		{
			url: "redis://example.com:6379",
			expected: &redis.Options{
				Addr:     "example.com:6379",
				Password: "",
				DB:       0,
			},
		},
		{
			url: "redis://example.com:6379?what=1",
			expected: &redis.Options{
				Addr:     "example.com:6379",
				Password: "",
				DB:       0,
			},
		},
		{
			url: "redis://foo:bar@example.com:6379",
			expected: &redis.Options{
				Addr:     "example.com:6379",
				Password: "bar",
				DB:       0,
			},
		},
		{
			url: "redis://foo@example.com:6379",
			expected: &redis.Options{
				Addr:     "example.com:6379",
				Password: "foo",
				DB:       0,
			},
		},
		{
			url: "redis://example.com:6379/1/foobar",
			expected: &redis.Options{
				Addr:     "example.com:6379",
				Password: "",
				DB:       1,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.url, func(t *testing.T) {
			opt, err := coreredis.Parse(c.url)
			if err != nil {
				t.Error(err)
			}
			equals(t, c.expected, opt)
		})
	}
}

func equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)\n", expected, actual)
	}
}
