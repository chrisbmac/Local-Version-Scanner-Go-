package detectors

type GoDetector struct{}

func (g GoDetector) Name() string { return "go" }
func (g GoDetector) Command() []string {
	return []string{"go", "version"}
}
