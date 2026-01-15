package detectors

type JavaDetector struct{}

func (g JavaDetector) Name() string {
	return "Java"
}

func (g JavaDetector) Command() []string {
	return []string{"java", "-version"}
}
