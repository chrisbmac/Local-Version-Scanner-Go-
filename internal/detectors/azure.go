package detectors

type AzDetector struct{}

func (a AzDetector) Name() string { return "az" }
func (a AzDetector) Command() []string {
	return []string{"az", "--version"}
}
