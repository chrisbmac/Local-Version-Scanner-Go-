package detectors

type Detector interface {
	Name() string
	Command() []string
}
