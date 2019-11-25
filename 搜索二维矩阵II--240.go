package main

/*
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 */
//func main() {
//	s := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
//	result := searchMatrix(s,9)
//	fmt.Println(result)
//	result2 := searchMatrix2(s,9)
//	fmt.Println(result2)
//
//}
/* 解法一：(分治法)
 * 左下角的元素是这一行中最小的元素，同时又是这一列中最大的元素。比较左下角元素和目标：
 * 若左下角元素等于目标，则找到
 * 若左下角元素大于目标，则目标不可能存在于当前矩阵的最后一行，问题规模可以减小为在去掉最后一行的子矩阵中寻找目标
 * 若左下角元素小于目标，则目标不可能存在于当前矩阵的第一列，问题规模可以减小为在去掉第一列的子矩阵中寻找目标
 * 若最后矩阵减小为空，则说明不存在
 */
func searchMatrix(matrix [][]int, target int) bool {
	b,i,j := false,len(matrix),0
	var m int
	if i==0 {
		return b
	}else {
		if len(matrix[0])==0 {
			return b
		}else {
			m = len(matrix[0])
		}
	}
	for i>=1&&j<m {
		if matrix[i-1][j]==target {
			b = true
			break
		}else if matrix[i-1][j]>target {
			i--
		}else if matrix[i-1][j]<target {
			j++
		}
	}
	return b
}
/* 解法二：(分治法)
 * 右上角的元素是这一行中最大的元素，同时又是这一列中最小的元素。比较右上角元素和目标：
 * 若右上角元素等于目标，则找到
 * 若右上角元素大于目标，则目标不可能存在于当前矩阵的最后一列，问题规模可以减小为在去掉最后一列的子矩阵中寻找目标
 * 若右上角元素小于目标，则目标不可能存在于当前矩阵的第一行，问题规模可以减小为在去掉第一行的子矩阵中寻找目标
 * 若最后矩阵减小为空，则说明不存在
 */
func searchMatrix2(matrix [][]int, target int) bool {
	b,i,m := false,0,len(matrix)
	var j int
	if m==0 {
		return b
	}else {
		if len(matrix[0])==0 {
			return b
		}else {
			j = len(matrix[0])
		}
	}
	for i<m&&j>=1 {
		if matrix[i][j-1]==target {
			b = true
			return b
		}else if matrix[i][j-1]>target {
			j--
		}else if matrix[i][j-1]<target {
			i++
		}
	}
	return b
}