package node

import (
	"github.com/cloudfoundry/packit"
)

type BuildPlanMetadata struct {
	Build         bool   `toml:"build"`
	Launch        bool   `toml:"launch"`
}
func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "npm"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "node",
						Metadata: BuildPlanMetadata{
							Build:  true,
							Launch: true,
						},
					},
					{
						Name: "npm",
					},
				},
			},
		}, nil
	}
}
