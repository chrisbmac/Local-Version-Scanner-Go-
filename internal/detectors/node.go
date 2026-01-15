package detectors

type NodeDetector struct{}

func (n NodeDetector) Name() string { return "node" }
func (n NodeDetector) Command() []string {
	return []string{"node", "--version"}
}
