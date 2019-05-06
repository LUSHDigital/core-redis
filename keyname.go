package coreredis

import "fmt"

// Arg defines an argument for building key names
type Arg struct {
	Name  string `json:"k"`
	Value string `json:"v"`
}

func (k Arg) String() string {
	return fmt.Sprintf("%s:%v", k.Name, k.Value)
}

// KeyName returns a key name formatted using the following rules:
//
//  prefix|key:value|key2|value2
func KeyName(prefix string, args []Arg) string {
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
