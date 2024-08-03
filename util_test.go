package helper

import (
	"reflect"
	"testing"
)

func TestStringToIntArr(t *testing.T) {
	type args struct {
		inp       string
		separator string
	}
	tests := []struct {
		name     string
		args     args
		wantResp []int
		wantErr  bool
	}{
		{
			name: "Test 1",
			args: args{
				inp:       "1,2,3,4,5",
				separator: ",",
			},
			wantResp: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test 2",
			args: args{
				inp:       "1 2 3 4 5",
				separator: " ",
			},
			wantResp: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test 3",
			args: args{
				inp:       "1,2,3,4,5",
				separator: " ",
			},
			wantResp: []int{},
			wantErr:  true,
		},
		{
			name: "Test 4",
			args: args{
				inp:       "a,b,c,d,e",
				separator: ",",
			},
			wantResp: []int{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := StringToIntArr(tt.args.inp, tt.args.separator)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToIntArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("StringToIntArr() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNumArrToStr(t *testing.T) {
	type args[T comparable] struct {
		inp []T
	}
	tests := []struct {
		name    string
		args    args[int]
		wantRes string
	}{
		{
			name: "Test 1",
			args: args[int]{
				inp: []int{1, 2, 3, 4, 5},
			},
			wantRes: "1,2,3,4,5",
		},
		{
			name: "Test 2",
			args: args[int]{
				inp: []int{2, 6},
			},
			wantRes: "2,6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := NumArrToStr(tt.args.inp); gotRes != tt.wantRes {
				t.Errorf("NumArrToStr() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		data   []T
		toFind T
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "Test 1",
			args: args[int]{
				data:   []int{1, 2, 3, 4, 5},
				toFind: 3,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args[int]{
				data:   []int{1, 2, 3, 4, 5},
				toFind: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.data, tt.args.toFind); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 9,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{-11, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		nums []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Test 1",
			args: args{
				nums: []float32{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			want: 1.1,
		},
		{
			name: "Test 2",
			args: args{
				nums: []float32{1.1, 2.2, 3.3, -4.4, 5.5, 6.6, 7.7, 8.8, 9.9},
			},
			want: -4.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
