package detectors

type CargoDetector struct{}

func (c CargoDetector) Name() string { return "cargo" }
func (c CargoDetector) Command() []string {
	return []string{"cargo", "--version"}
}
