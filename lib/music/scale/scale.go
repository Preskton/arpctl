package scale

var AllNamedScalePatterns []ScalePattern

type ScalePattern struct {
	Name            string
	Degrees         string
	Intervals       string
	IntegerNotation []int
}
