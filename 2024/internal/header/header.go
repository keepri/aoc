package header

import (
	"fmt"
	"strings"
)

type printer struct {
	hz  string
	vt  string
	pad int
	w   int
}

func Print(day, part int) {
	p := printer{
		hz: "=",
		vt: "",
		w:  75,
	}

	var challenge []string
	switch day {
	case 1:
		challenge = p.dayOne(part)
	}

	msgs := []string{"AOC 2024 LETS GO!"}
	ok := len(challenge) > 0
	var body string
	if ok {
		body = challenge[len(challenge)-1]
		challenge = challenge[:len(challenge)-1]
		challenge = append([]string{""}, challenge...)
		msgs = append(msgs, challenge...)
	}

	p.header(msgs)
	fmt.Println()

	if ok {
		fmt.Println(body)
		fmt.Println()
	}
}

func (p *printer) header(msg []string) {
	fmt.Println(p.repeat(p.hz, p.w))

	for _, m := range msg {
		m = strings.TrimSpace(m)
		x := (p.w - len(m)) / 2

		fmt.Println(p.vt + p.repeat(" ", x) + m + p.repeat(" ", x) + p.vt)
	}

	fmt.Println(p.repeat(p.hz, p.w))
}

func (_ *printer) repeat(s string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
