package filters

import (
	"reflect"
	"testing"
)

func TestRemoveOutliers(t *testing.T) {
	type args struct {
		samples []float64
		scale   float64
	}
	tests := []struct {
		name  string
		args  args
		want  []float64
		want1 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RemoveOutliers(tt.args.samples, tt.args.scale)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveOutliers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RemoveOutliers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_removeOutliers(t *testing.T) {
	type args struct {
		samples []float64
	}
	tests := []struct {
		name  string
		args  args
		want  []float64
		want1 []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeOutliers(tt.args.samples)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeOutliers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeOutliers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
