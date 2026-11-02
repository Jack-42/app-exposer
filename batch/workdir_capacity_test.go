package batch

import (
	"testing"

	"github.com/cyverse-de/model/v10"
)

func TestWorkdirCapacity(t *testing.T) {
	const gi = int64(1) << 30
	tests := []struct {
		name     string
		minDisks []int64
		want     int64
	}{
		{"no steps", nil, 5 * gi},
		{"unset", []int64{0}, 5 * gi},
		{"below default", []int64{1 * gi}, 5 * gi},
		{"above default", []int64{32 * gi}, 32 * gi},
		{"max across steps", []int64{8 * gi, 0, 20 * gi}, 20 * gi},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &model.Analysis{}
			for _, d := range tt.minDisks {
				step := model.Step{}
				step.Component.Container.MinDiskSpace = d
				analysis.Steps = append(analysis.Steps, step)
			}
			w := &WorkflowMaker{analysis: analysis}
			got := w.workdirCapacity()
			if got.Value() != tt.want {
				t.Errorf("workdirCapacity() = %s, want %d", got.String(), tt.want)
			}
		})
	}
}
