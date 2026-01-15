package detectors

type NpmDetector struct{}

func (n NpmDetector) Name() string { return "npm" }
func (n NpmDetector) Command() []string {
	return []string{"npm", "--version"}
}
