package detectors

type TerraformDetector struct{}

func (t TerraformDetector) Name() string { return "terraform" }
func (t TerraformDetector) Command() []string {
	return []string{"terraform", "version"}
}
