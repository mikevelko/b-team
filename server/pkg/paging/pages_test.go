package paging

import "testing"

func Test_getPagesCount(t *testing.T) {
	type args struct {
		items        int
		itemsPerPage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "For divisible item count, it should create just enough pages",
			args: args{
				itemsPerPage: 10,
				items:        70,
			},
			want: 7,
		},
		{
			name: "For indivisible item count, it should create additional page for remainder",
			args: args{
				itemsPerPage: 10,
				items:        57,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPagesCount(tt.args.items, tt.args.itemsPerPage); got != tt.want {
				t.Errorf("getPagesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPageItems(t *testing.T) {
	type args struct {
		currPage     int
		itemsPerPage int
		itemsCount   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "If items count is zero, range should be 0:0, so slice[start:end] wont do error",
			args: args{
				currPage:     4,
				itemsPerPage: 10,
				itemsCount:   0,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "If items per page is zero, range should be 0:0 (avoid zero division)",
			args: args{
				currPage:     4,
				itemsPerPage: 0,
				itemsCount:   50,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "If current page is bigger than pageCount, it should point to last page",
			args: args{
				currPage:     8,
				itemsPerPage: 10,
				itemsCount:   50,
			},
			want:  40,
			want1: 50,
		},
		{
			name: "If current page is less than 1, it should point to first page",
			args: args{
				currPage:     -213,
				itemsPerPage: 10,
				itemsCount:   50,
			},
			want:  0,
			want1: 10,
		},
		{
			name: "If current page is less than 1, it should point to first page",
			args: args{
				currPage:     -213,
				itemsPerPage: 10,
				itemsCount:   50,
			},
			want:  0,
			want1: 10,
		},
		{
			name: "If current page is last, and item count is not divisible, it should return remainder of items",
			args: args{
				currPage:     2,
				itemsPerPage: 10,
				itemsCount:   11,
			},
			want:  10,
			want1: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetPageItems(tt.args.currPage, tt.args.itemsPerPage, tt.args.itemsCount)
			if got != tt.want {
				t.Errorf("GetPageItems() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetPageItems() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
