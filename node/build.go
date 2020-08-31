package node

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/cloudfoundry/packit"
)

func Build() packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		file, err := os.Open(filepath.Join(context.CNBPath, "buildpack.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		var m struct {
			Metadata struct {
				Dependencies []struct {
					URI string `toml:"uri"`
				} `toml:"dependencies"`
			} `toml:"metadata"`
		}
		_, err = toml.DecodeReader(file, &m)
		if err != nil {
			return packit.BuildResult{}, err
		}

		uri := m.Metadata.Dependencies[0].URI
		fmt.Printf("URI -> %s\n", uri)

		nodeLayer, err := context.Layers.Get("node", packit.LaunchLayer)
		if err != nil {
			return packit.BuildResult{}, err
		}

		err = nodeLayer.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Running npm...")
		cmd := exec.Command("npm",
			"install",
		)
		cmd.Dir = filepath.Join(context.WorkingDir, "web/themes/custom/lakshminp_theme")
		err2 := cmd.Run()
		
		if err2 != nil {
			return packit.BuildResult{}, err2
		}
		return packit.BuildResult{
			Plan: context.Plan,
			Layers: []packit.Layer{
				nodeLayer,
			},
		}, nil
	}
}
