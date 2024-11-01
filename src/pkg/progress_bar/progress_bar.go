package progress_bar

import (
	"fmt"
	"io"
	"strings"
)

type ProgressBar struct {
	Writer io.Writer
	Length int
	Max    int
	state  int
}

func (p *ProgressBar) Tick() {
	if p.state < p.Max {
		p.state += 1
	}

	progress := int(float32(p.state) / float32(p.Max) * float32(p.Length-2))

	bar := fmt.Sprintf("\r[%s%s] (%d / %d)",
		strings.Repeat("█", progress),
		strings.Repeat("░", (p.Length-2)-progress),
		p.state,
		p.Max,
	)

	p.Writer.Write([]byte(bar))
}
