package __20

//todo:fix this error
//搜索二维矩阵
//func searchMatrix(matrix [][]int, target int) bool {
//	if len(matrix)==0 || len(matrix[0])==0{
//		return false
//	}
//	slice:=[]int{}
//
//	for i:=0; i < len(matrix)-1; i++{
//		for j := 0; j < len(matrix[0])-1;j++ {
//			slice=append(slice,matrix[i][j])
//		}
//	}
//	l,r:=0,len(slice)-1
//	mid:=(l+r)/2
//	for l<r {
//		if target == slice[mid] {
//			return true
//			break
//		}
//		if slice[mid]< target{
//			l=mid+1
//		}else {
//			r=mid-1
//		}
//	}
//	return false
//
//}
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1 // 转一维
	for l <= r {
		mid := (r-l)>>1 + l // mid是一维的索引

		curRow := mid / cols        // 整除，得二维的当前行索引
		curCol := mid - curRow*cols // 一维mid减去它头顶上行的元素个数，得二维的当前列索引

		if matrix[curRow][curCol] == target {
			return true
		} else if matrix[curRow][curCol] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
