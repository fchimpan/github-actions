package fizzbazz

import "testing"

func Test_fizzbazz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "return FizzBuzz when value is divisible by 15",
			args: args{n: 30},
			want: "FizzBuzz",
		},
		{
			name: "return Fizz when value is divisible by 3",
			args: args{n: 9},
			want: "Fizz",
		},
		{
			name: "return Buzz when value is divisible by 5",
			args: args{n: 10},
			want: "Buzz",
		},
		{
			name:    "return error",
			args:    args{n: -1},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fizzbazz(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("fizzbazz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fizzbazz() = %v, want %v", got, tt.want)
			}
		})
	}
}
