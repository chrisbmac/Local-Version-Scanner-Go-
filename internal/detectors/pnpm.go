package detectors

type PnpmDetector struct{}

func (p PnpmDetector) Name() string { return "pnpm" }
func (p PnpmDetector) Command() []string {
	return []string{"pnpm", "--version"}
}
