package coreredis_test

import (
	"fmt"
	"testing"

	"github.com/LUSHDigital/core-redis"
)

func TestKeyName(t *testing.T) {
	type args struct {
		prefix string
		args   []coreredis.Arg
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single argument",
			args: args{
				prefix: "test",
				args: []coreredis.Arg{
					{
						Name:  "a",
						Value: "b",
					},
				},
			},
			want: "test|a:b",
		},
		{
			name: "n arguments",
			args: args{
				prefix: "test",
				args: []coreredis.Arg{
					{
						Name:  "a",
						Value: "b",
					},
					{
						Name:  "c",
						Value: "d",
					},
				},
			},
			want: "test|a:b|c:d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coreredis.KeyName(tt.args.prefix, tt.args.args); got != tt.want {
				t.Errorf("KeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleKeyName() {
	args := []coreredis.Arg{
		{Name: "hello", Value: "world"},
		{Name: "starwisp", Value: "probe"},
	}
	space := coreredis.KeyName("testcase", args)
	fmt.Println(space)
	// Output: testcase|hello:world|starwisp:probe
}