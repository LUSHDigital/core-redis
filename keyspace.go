package coreredis

import "fmt"

type Arg struct {
	Name  string `json:"k"`
	Value string `json:"v"`
}

func (k Arg) String() string {
	return fmt.Sprintf("%s:%v", k.Name, k.Value)
}

func Keyspace(prefix string, args []Arg) string {
	const sep = "|"
	prefix += sep
	max := len(args) - 1
	for k, arg := range args {
		prefix += arg.String()
		if k == max {
			continue
		}
		prefix += sep
	}
	return prefix
}
