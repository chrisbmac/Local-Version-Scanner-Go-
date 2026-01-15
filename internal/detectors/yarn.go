package detectors

type YarnDetector struct{}

func (y YarnDetector) Name() string { return "yarn" }
func (y YarnDetector) Command() []string {
	return []string{"yarn", "--version"}
}
