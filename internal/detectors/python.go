package detectors

type PythonDetector struct{}

func (p PythonDetector) Name() string { return "python" }
func (p PythonDetector) Command() []string {
	return []string{"python3", "--version"}
}
