package music

import (
	"github.com/preskton/arpctl/lib/music/scale"
	log "github.com/sirupsen/logrus"
)

type PatternContext struct {
	RootNote     *Note
	NextNote     *Note
	Scale        *scale.ScalePattern
	PatternIndex int
}

func (pc *PatternContext) Advance() {
	pc.PatternIndex++

	// TODO bounce modes
	if pc.PatternIndex > len(pc.Scale.IntegerNotation)-1 {
		pc.PatternIndex = 0
	}

	pc.NextNote = GetNoteByRootAndOffset(pc.RootNote, pc.Scale.IntegerNotation[pc.PatternIndex])

	log.WithField("pc", pc).Info("advancing pattern")
}
