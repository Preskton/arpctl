package scale

var AllScalePatterns []ScalePattern

type ScalePattern struct {
	Name            string
	Degrees         string
	Intervals       string
	IntegerNotation []int
}

func GetScaleByName(name string) *ScalePattern {
	for _, scale := range AllScalePatterns {
		if isNameMatch(name, &scale) {
			return &scale
		}
	}

	return nil
}

func isNameMatch(name string, scale *ScalePattern) bool {
	return scale != nil && scale.Name == name
}

func GetScaleBySearch(target string) *ScalePattern {
	for _, scale := range AllScalePatterns {
		if isNameMatch(target, &scale) {
			return &scale
		}
	}

	return nil
}
