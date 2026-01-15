package detectors

type KubectlDetector struct{}

func (k KubectlDetector) Name() string { return "kubectl" }
func (k KubectlDetector) Command() []string {
	return []string{"kubectl", "version", "--client", "--short"}
}
