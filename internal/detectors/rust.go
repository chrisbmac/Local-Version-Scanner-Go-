package detectors

type RustDetector struct{}

func (r RustDetector) Name() string { return "rustc" }
func (r RustDetector) Command() []string {
	return []string{"rustc", "--version"}
}
