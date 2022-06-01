package music

import (
	"github.com/preskton/arpctl/lib/music/scale"

	log "github.com/sirupsen/logrus"
)

// Holds current state about progress through a pattern.
type PatternContext struct {
	// The starting note of the scale
	RootNote *Note
	// The next note to play
	NextNote *Note
	// The scale to follow when selecting notes
	Scale *scale.ScalePattern
	// The current index of the pattern (Scale.IntegerNotation)
	PatternIndex int
	// Current direction of the pattern, -1 (decreasing) or 1 (increasing)
	PatternDirection int
	// The function used to choose the next note in the scale
	Advancer func()
}

// Arps in an increasing pattern and resets to the root
func (pc *PatternContext) ToTheMoon() {
	pc.PatternIndex++

	if pc.PatternIndex > len(pc.Scale.IntegerNotation)-1 {
		pc.PatternIndex = 0
	}

	pc.NextNote = GetNoteByRootAndOffset(pc.RootNote, pc.Scale.IntegerNotation[pc.PatternIndex])
}

// Arps in a decreasing pattern and resets to the last note in the scale
func (pc *PatternContext) Descent() {
	pc.PatternIndex--

	if pc.PatternIndex < 0 {
		pc.PatternIndex = len(pc.Scale.IntegerNotation) - 1
	}

	pc.NextNote = GetNoteByRootAndOffset(pc.RootNote, pc.Scale.IntegerNotation[pc.PatternIndex])
}

// Arps in an increasing pattern, then decreasing after it hits the last note in the scale, then repeat. Whee!
func (pc *PatternContext) BouncyCastle() {
	pc.PatternIndex = pc.PatternIndex + pc.PatternDirection

	if pc.PatternIndex > len(pc.Scale.IntegerNotation)-1 {
		pc.PatternDirection = -1
		pc.PatternIndex = len(pc.Scale.IntegerNotation) - 2
		log.WithField("index", pc.PatternIndex).Info("bounce down")
	} else if pc.PatternIndex < 0 {
		pc.PatternIndex = 1
		pc.PatternDirection = 1
		log.WithField("index", pc.PatternIndex).Info("bounce up")
	}

	log.WithField("index", pc.PatternIndex).Info("curreent index")

	pc.NextNote = GetNoteByRootAndOffset(pc.RootNote, pc.Scale.IntegerNotation[pc.PatternIndex])
}
