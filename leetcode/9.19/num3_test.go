package __19

import "testing"

//func Test_dfs(t *testing.T) {
//	type args struct {
//		grid [][]int
//		i    int
//		j    int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//		{
//			"岛屿的最大面积",
//			args{
//			//	grid: [
//			//	[0,0,1,0,0,0,0,1,0,0,0,0,0],
//			//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//			//[0,1,1,0,1,0,0,0,0,0,0,0,0],
//			//[0,1,0,0,1,1,0,0,1,0,1,0,0],
//			//[0,1,0,0,1,1,0,0,1,1,1,0,0],
//			//[0,0,0,0,0,0,0,0,0,0,1,0,0],
//			//[0,0,0,0,0,0,0,1,1,1,0,0,0],
//			//[0,0,0,0,0,0,0,1,1,0,0,0,0]
//			//	]
//
//				[][]int{
//			{0,0,1,0,0,0,0,1,0,0,0,0,0},
//			{0,0,0,0,0,0,0,1,1,1,0,0,0},
//			{0,1,1,0,1,0,0,0,0,0,0,0,0},
//			{0,1,0,0,1,1,0,0,1,0,1,0,0},
//			{0,1,0,0,1,1,0,0,1,1,1,0,0},
//			{0,0,0,0,0,0,0,0,0,0,1,0,0},
//			{0,0,0,0,0,0,0,1,1,1,0,0,0},
//			{0,0,0,0,0,0,0,1,1,0,0,0,0},
//				},
//					13,
//					13,
//
//			},
//			6,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := dfs(tt.args.grid, tt.args.i, tt.args.j); got != tt.want {
//				t.Errorf("dfs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_maxAreaOfIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"岛屿的最大面积",
			args{
				[][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.args.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
