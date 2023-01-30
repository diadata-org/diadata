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
		{
        name: "Test case 1",
        args: args{
            samples: []float64{1, 2, 3, 4, 5},
            scale:   1,
        },
        want:  []float64{1, 2, 3, 4, 5},
        want1: []int{},
    },
    {
        name: "Test case 2",
        args: args{
            samples: []float64{1, 2, 3, 4, 5, 100},
            scale:   1,
        },
        want:  []float64{1, 2, 3, 4, 5},
        want1: []int{5},
    },
    {
        name: "Test case 3",
        args: args{
            samples: []float64{1, 2, 3, 4, 5, -100},
            scale:   1,
        },
        want:  []float64{1, 2, 3, 4, 5},
        want1: []int{5},
    },
    {
        name: "Test case 4",
        args: args{
            samples: []float64{1, 2, 3, 4, 5, 100, 200},
            scale:   2,
        },
        want:  []float64{1, 2, 3, 4, 5},
        want1: []int{5, 6},
    },
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
		{
			name: "normal case with no outliers",
			args: args{samples: []float64{1, 2, 3, 4, 5}},
			want: []float64{1, 2, 3, 4, 5},
			want1: []int{},
		},
		{
			name: "normal case with one outlier",
			args: args{samples: []float64{1, 2, 3, 4, 5, 100}},
			want: []float64{1, 2, 3, 4, 5},
			want1: []int{5},
		},
		{
			name: "normal case with two outliers",
			args: args{samples: []float64{1, 2, 3, 4, 5, 100, 200}},
			want: []float64{1, 2, 3, 4, 5},
			want1: []int{5, 6},
		},
		{
			name: "empty input",
			args: args{samples: []float64{}},
			want: []float64{},
			want1: []int{},
		},
		{
			name: "input with negative numbers",
			args: args{samples: []float64{-1, -2, -3, -4, -5}},
			want: []float64{-1, -2, -3, -4, -5},
			want1: []int{},
		},
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

func Test_removeOutliersScaled(t *testing.T) {
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
		{
		name: "normal case",
		args: args{
			samples: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			scale:   1,
		},
		want:  []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		want1: []int{},
	},
	{
		name: "outliers present",
		args: args{
			samples: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30},
			scale:   1,
		},
		want:  []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		want1: []int{10, 11},
	},
	{
		name: "all outliers",
		args: args{
			samples: []float64{1, 20, 30, 40, 50},
			scale:   1,
		},
		want:  []float64{},
		want1: []int{0, 1, 2, 3, 4},
	},
	{
		name: "empty input",
		args: args{
			samples: []float64{},
			scale:   1,
		},
		want:  []float64{},
		want1: []int{},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeOutliersScaled(tt.args.samples, tt.args.scale)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeOutliersScaled() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeOutliersScaled() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_computeMean(t *testing.T) {
	type args struct {
		samples []float64
		weights []float64
	}
	tests := []struct {
		name     string
		args     args
		wantMean float64
		wantErr  bool
	}{
		{
			name: "compute mean of empty samples",
			args: args{
				samples: []float64{},
				weights: []float64{},
			},
			wantMean: 0,
			wantErr:  true,
		},
		{
			name: "compute mean of single value",
			args: args{
				samples: []float64{5.5},
				weights: []float64{},
			},
			wantMean: 5.5,
			wantErr:  false,
		},
		{
			name: "compute mean of multiple values",
			args: args{
				samples: []float64{1, 2, 3, 4, 5},
				weights: []float64{},
			},
			wantMean: 3,
			wantErr:  false,
		},
		{
			name: "compute weighted mean of multiple values",
			args: args{
				samples: []float64{1, 2, 3, 4, 5},
				weights: []float64{2, 3, 1, 2, 1},
			},
			wantMean: 2.3636363636363638,
			wantErr:  false,
		},
		{
			name: "compute mean with different lenght of samples and weights",
			args: args{
				samples: []float64{1, 2, 3, 4, 5},
				weights: []float64{2, 3},
			},
			wantMean: 0,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMean, err := computeMean(tt.args.samples, tt.args.weights)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeMean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMean != tt.wantMean {
				t.Errorf("computeMean() = %v, want %v", gotMean, tt.wantMean)
			}
		})
	}
}

func Test_computeQuartiles(t *testing.T) {
	type args struct {
		samples []float64
	}
	tests := []struct {
		name   string
		args   args
		wantQ1 float64
		wantQ3 float64
	}{
		{name: "Test case 1", args: args{samples: []float64{1, 2, 3, 4, 5}}, wantQ1: 1.5, wantQ3: 4.5},
{name: "Test case 2", args: args{samples: []float64{5, 4, 3, 2, 1}}, wantQ1: 1.5, wantQ3: 4.5},
{name: "Test case 3", args: args{samples: []float64{1, 1, 1, 1, 1, 2}}, wantQ1: 1, wantQ3: 2},
{name: "Test case 4", args: args{samples: []float64{1, 2, 3, 4, 5, 6}}, wantQ1: 2, wantQ3: 5},
{name: "Test case 5", args: args{samples: []float64{}}, wantQ1: 0, wantQ3: 0},
{name: "Test case 6", args: args{samples: []float64{-1, -2, -3, -4, -5}}, wantQ1: -3, wantQ3: -2},
{name: "Test case 7", args: args{samples: []float64{-5, -4, -3, -2, -1}}, wantQ1: -3, wantQ3: -2},
{name: "Test case 8", args: args{samples: []float64{-1, -1, -1, -1, -1, -2}}, wantQ1: -1, wantQ3: -1},
{name: "Test case 9", args: args{samples: []float64{-1, -2, -3, -4, -5, -6}}, wantQ1: -3, wantQ3: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ1, gotQ3 := computeQuartiles(tt.args.samples)
			if gotQ1 != tt.wantQ1 {
				t.Errorf("computeQuartiles() gotQ1 = %v, want %v", gotQ1, tt.wantQ1)
			}
			if gotQ3 != tt.wantQ3 {
				t.Errorf("computeQuartiles() gotQ3 = %v, want %v", gotQ3, tt.wantQ3)
			}
		})
	}
}

func Test_computeMedian(t *testing.T) {
	type args struct {
		samples []float64
	}
	tests := []struct {
		name       string
		args       args
		wantMedian float64
	}{
		{name: "Test case 1", args: args{samples: []float64{1, 2, 3, 4, 5}}, wantMedian: 3},
{name: "Test case 2", args: args{samples: []float64{5, 4, 3, 2, 1}}, wantMedian: 3},
{name: "Test case 3", args: args{samples: []float64{1, 1, 1, 1, 1, 2}}, wantMedian: 1},
{name: "Test case 4", args: args{samples: []float64{1, 2, 3, 4, 5, 6}}, wantMedian: 3.5},
{name: "Test case 5", args: args{samples: []float64{}}, wantMedian: 0},
{name: "Test case 6", args: args{samples: []float64{-1, -2, -3, -4, -5}}, wantMedian: -3},
{name: "Test case 7", args: args{samples: []float64{-5, -4, -3, -2, -1}}, wantMedian: -3},
{name: "Test case 8", args: args{samples: []float64{-1, -1, -1, -1, -1, -2}}, wantMedian: -1},
{name: "Test case 9", args: args{samples: []float64{-1, -2, -3, -4, -5, -6}}, wantMedian: -3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMedian := computeMedian(tt.args.samples); gotMedian != tt.wantMedian {
				t.Errorf("computeMedian() = %v, want %v", gotMedian, tt.wantMedian)
			}
		})
	}
}
