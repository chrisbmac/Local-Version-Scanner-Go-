package output

var Emojis = map[string]string{
	"git":       "🐙",  // Octocat vibe
	"go":        "🐹",  // Gopher
	"node":      "🟢",  // Node green circle
	"npm":       "📦",  // Package
	"yarn":      "🧶",  // Yarn ball
	"pnpm":      "📦",  // Same idea
	"python":    "🐍",  // Snake
	"java":      "☕",  // Coffee cup
	"docker":    "🐳",  // Whale
	"rustc":     "🦀",  // Ferris crab
	"cargo":     "📦",  // Crates
	"kubectl":   "☸️", // Kubernetes wheel
	"terraform": "🟪",  // Purple square (Terraform color)
	"aws":       "☁️", // Cloud
	"az":        "🔷",  // Azure diamond
	"gcloud":    "🌥️", // Cloud
	"make":      "🛠️", // Tools
	"git-lfs":   "📁",  // File
}

func EmojiFor(name string) string {
	if e, ok := Emojis[name]; ok {
		return e
	}
	return "📌" // default icon
}
