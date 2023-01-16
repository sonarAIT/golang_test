package add

import "testing"

func TestAddFunc2(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal case",
			args: args{
				n1: 1,
				n2: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddFunc2(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("AddFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
