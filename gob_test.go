package coreredis_test

import (
	"testing"
	"time"

	coreredis "github.com/LUSHDigital/core-redis"
	miniredis "github.com/alicebob/miniredis"
	redis "github.com/go-redis/redis"
)

func TestGobClient(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	cc := coreredis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	gc := coreredis.NewGobClient(cc)
	if err := gc.SetGob("foobar", "hello world", 1*time.Minute); err != nil {
		t.Error(err)
	}
	var actual string
	if err := gc.GetGob("foobar", &actual); err != nil {
		t.Error(err)
	}
	equals(t, "hello world", actual)
}
