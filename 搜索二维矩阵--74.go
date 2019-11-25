package main

//func main() {
//	a := [][]int{{1},{3}}
//	target := 6
//	result := searchMatrix3(a,target)
//	fmt.Println(result)
//	result2 := searchMatrix4(a,target)
//	fmt.Println(result2)
//	result3 := searchMatrix5(a,target)
//	fmt.Println(result3)
//}
/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
 */
/**
  *这道题的关键点在于利用矩阵升序的特性，
  *需要选择合适的切入点与目标值相比较以缩小数据返围，
  *很明显如果选择左上角或右下角当做切入点并不能起到只能分别排除一个选项，
  *而如果选择右上角或者左下角进行比较即可起到更明显的效果。
 */
//方案一：从左上角开始
func searchMatrix3(matrix [][]int, target int) bool {
	b,m := false,len(matrix)
	var n int
	if m==0 {
		return b
	}else {
		if len(matrix[0])==0 {
			return b
		}else {
			n = len(matrix[0])
		}
	}
	if m==1 {
		for j:=0;j<n;j++ {
			if matrix[0][j]==target {
				b = true
				return b
			}
		}
	}else {
		for i:=0;i<m-1;i++ {
			if matrix[i][0]>target {
				return b
			}else if matrix[i][0]==target {
				b = true
				return b
			}else if matrix[i][0]<target&&matrix[i+1][0]>target {
				//遍历第i行
				for j:=0;j<n;j++ {
					if matrix[i][j]==target {
						b = true
						return b
					}
				}
			}else if i==m-2&&matrix[i+1][0]<target {
				//遍历第i+1行
				for j:=0;j<n;j++ {
					if matrix[i+1][j]==target {
						b = true
						return b
					}
				}
			}else if i==m-2&&matrix[i+1][0]==target {
				b = true
				return b
			}
		}
	}
	return b
}
//方案二：从右上角开始
func searchMatrix4(matrix [][]int, target int) bool {
	b,m := false,len(matrix)
	var n int
	if m==0 {
		return b
	}else {
		if len(matrix[0])==0 {
			return b
		}else {
			n = len(matrix[0])
		}
	}
	for i:=0;i<m;i++ {
		if matrix[i][n-1]==target {
			b = true
			break
		}else if matrix[i][n-1]>target {
			//遍历第i行
			for j:=0;j<n;j++ {
				if matrix[i][j]==target {
					b = true
					return b
				}
			}
			break
		}
	}
	return b
}
//方案三：注意到输入的 m x n 矩阵可以视为长度为 m x n的有序数组。
/*
由于该虚数组的序号可以由下式方便地转化为原矩阵中的行和列 (我们当然不会真的创建一个新数组) ，该有序数组非常适合二分查找。
row = idx / n ， col = idx % n。
这是一个标准二分查找算法 :
初始化左右序号
left = 0 和 right = m x n - 1。
While left < right :
选取虚数组最中间的序号作为中间序号: pivot_idx = (left + right) / 2。
该序号对应于原矩阵中的 row = pivot_idx / n行 col = pivot_idx % n 列, 由此可以拿到中间元素pivot_element。该元素将虚数组分为两部分。
比较 pivot_element 与 target 以确定在哪一部分进行进一步查找。
 */
func searchMatrix5(matrix [][]int, target int) bool {
	if len(matrix)==0 {
		return false
	}else {
		if len(matrix[0])==0 {
			return false
		}
	}
	b,left,right := false,0,len(matrix)*len(matrix[0])-1
	for left<=right {
		middle := (left+right)/2
		row,col := middle/len(matrix[0]),middle%len(matrix[0])
		if matrix[row][col]==target {
			b = true
			break
		}else if matrix[row][col]<target {
			left = middle + 1
		}else {
			right = middle - 1
		}
	}
	return b
}