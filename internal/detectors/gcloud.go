package detectors

type GcloudDetector struct{}

func (g GcloudDetector) Name() string { return "gcloud" }
func (g GcloudDetector) Command() []string {
	return []string{"gcloud", "--version"}
}
