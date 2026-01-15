package detectors

type DockerDetector struct{}

func (d DockerDetector) Name() string { return "docker" }
func (d DockerDetector) Command() []string {
	return []string{"docker", "--version"}
}
