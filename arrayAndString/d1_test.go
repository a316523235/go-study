package arrayAndString

import (
	"fmt"
	"strings"
	"testing"
)

func TestChangeStr(t *testing.T)  {
	s := "piaoyun"
	c := []byte(s)
	c[0] = 'X'
	s2 := string(c)
	fmt.Printf("%s\n", s2)


	ss := []string{
		"sh",
		" hn",
		" test",
	}

	var b strings.Builder
	for _, s := range ss {
		fmt.Fprint(&b, s)
	}

	print(b.String())
}
