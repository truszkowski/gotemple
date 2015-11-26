This is not another temple.
This is golang templates!

# Example

```golang
package main

import (
	"fmt"
	"github.com/truszkowski/gotemple"
	"os"
	"strconv"
)

func main() {
	s := "{{define \"example\"}}"
	s += "{{ range $i, $n := . }}"
	s += "{{$i}}: {{$n|inCount}}\n"
	s += "{{end}}"
	s += "{{end}}"
	tmpl := temple.NewText("tmpl")

	if _, err := tmpl.Parse(s); err != nil {
		fmt.Println("parse error:", err)
		os.Exit(1)
	}

	args := []int{}
	for _, v := range os.Args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("invalid argument", n, "error:", err)
			os.Exit(1)
		}
		args = append(args, n)
	}

	if err := tmpl.ExecuteTemplate(os.Stdout, "example", args); err != nil {
		fmt.Println("execute error:", err)
		os.Exit(1)
	}
}
```

Run:
```
$ ./example 12 33333 1230000000
0: 12
1: 33k
2: 1.2g
```
