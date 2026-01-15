package detectors

type MakeDetector struct{}

func (m MakeDetector) Name() string { return "make" }
func (m MakeDetector) Command() []string {
	return []string{"make", "--version"}
}
