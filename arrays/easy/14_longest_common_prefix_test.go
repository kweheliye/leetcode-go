package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Test Case 2",
			args: args{
				strs: []string{"dog", "flow", "car"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixHorizontalScanning(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefixHorizontalScanning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonPrefixVerticalScanning(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "Test Case 2",
			args: args{
				strs: []string{"dog", "flow", "car"},
			},
			want: "",
		},
		{
			name: "Test Case 3 - Empty Array",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "Test Case 4 - Single String",
			args: args{
				strs: []string{"alone"},
			},
			want: "alone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixVerticalScanning(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefixVerticalScanning() = %v, want %v", got, tt.want)
			}
		})
	}
}
