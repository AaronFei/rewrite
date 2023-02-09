package rewrite

import (
	"fmt"
	"strings"

	"github.com/k0kubun/go-ansi"
)

type rewrite_t struct {
	lineCount    int
	lastLinCount int
}

func Create() *rewrite_t {
	return &rewrite_t{
		lineCount: 0,
	}
}

func (p *rewrite_t) PrintMultiln(lines []string) {
	p.lineCount += len(lines)
	for i, _ := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\n", "")
		lines[i] = strings.ReplaceAll(lines[i], "\r", "")
	}

	printStr := strings.Join(lines, "\n")
	fmt.Println(printStr)
}

func (p *rewrite_t) Println(line ...interface{}) {
	p.lineCount += 1
	fmt.Println(line...)
}

func (p *rewrite_t) MoveCursorBack() {
	str := fmt.Sprintf("\033[%dE", p.lineCount)
	ansi.Printf(str)
	p.lastLinCount = p.lineCount
	p.lineCount = 0
}

func (p *rewrite_t) Stop() {
	str := fmt.Sprintf("\033[%dF", p.lastLinCount)
	ansi.Printf(str)
}
