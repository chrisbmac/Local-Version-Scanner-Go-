package main

import (
	"fmt"
	"strings"

	"github.com/chrisbmac/localVersionScanner/internal/detectors"
	"github.com/chrisbmac/localVersionScanner/internal/detectors/output"
	"github.com/chrisbmac/localVersionScanner/internal/runner"
)

func main() {
	fmt.Println("\n" + output.Cyan + "--- Developer Versions Scanner --- \n" + output.Reset)

	// List of detectors
	ds := []detectors.Detector{
		detectors.GitDetector{},
		detectors.GoDetector{},
		detectors.NodeDetector{},
		detectors.NpmDetector{},
		detectors.YarnDetector{},
		detectors.PnpmDetector{},
		detectors.PythonDetector{},
		detectors.JavaDetector{},
		detectors.DockerDetector{},
		detectors.RustDetector{},
		detectors.CargoDetector{},
		detectors.KubectlDetector{},
		detectors.TerraformDetector{},
		detectors.AwsDetector{},
		detectors.AzDetector{},
		detectors.GcloudDetector{},
		detectors.MakeDetector{},
		detectors.GitLfsDetector{},
	}

	// Run each dectector
	for _, d := range ds {
		result := runner.Run(d.Name(), d.Command())
		name := strings.ToUpper(result.Name)
		emoji := output.EmojiFor(result.Name)

		// fmt strings can only contain 4 strings, simple enough to store as below
		startError := emoji + " %s✖ %s:%s %s\n "
		startPass := emoji + " %s✔ %s:%s %s\n "

		if result.Error != "" {
			fmt.Printf(startError, output.Red, name, "Package Doesn't Exist!\n", output.Reset)
		} else {
			fmt.Printf(startPass, output.Green, name, result.Version+"\n", output.Reset)
		}
	}
	fmt.Println(output.Cyan + "--- End Of Developer Versions Scanner Results --- \n" + output.Reset)
}
