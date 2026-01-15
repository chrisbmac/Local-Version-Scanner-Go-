package detectors

type GitDetector struct{}

func (g GitDetector) Name() string {
	return "Git"
}

func (g GitDetector) Command() []string {
	return []string{"git", "--version"}
}
