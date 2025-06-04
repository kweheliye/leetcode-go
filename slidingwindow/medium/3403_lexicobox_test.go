package medium

import "testing"

func Test_answerStringV1(t *testing.T) {
	type args struct {
		word       string
		numFriends int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				word:       "dbca",
				numFriends: 2,
			},
			want: "dbc",
		},
		{
			name: "Test Case 2",
			args: args{
				word:       "gggg",
				numFriends: 4,
			},
			want: "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerStringV1(tt.args.word, tt.args.numFriends); got != tt.want {
				t.Errorf("answerStringV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
