package detectors

type AwsDetector struct{}

func (a AwsDetector) Name() string { return "aws" }
func (a AwsDetector) Command() []string {
	return []string{"aws", "--version"}
}
