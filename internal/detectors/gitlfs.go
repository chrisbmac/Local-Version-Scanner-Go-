package detectors

type GitLfsDetector struct{}

func (g GitLfsDetector) Name() string { return "git-lfs" }
func (g GitLfsDetector) Command() []string {
	return []string{"git", "lfs", "version"}
}
