package main

import (
	"math"
	"sort"
)

//func main() {
//	nums := []int{-1,2,1,-4}
//	target := 1
//	res := threeSumClosest1(nums,target)
//	fmt.Println(res)
//}
//方案一：暴力破解，o(n^3)
func threeSumClosest1(nums []int, target int) int {
	res,diff := 0,math.MaxInt16
	for i:=0;i<len(nums)-2;i++ {
		for j:=i+1;j<len(nums)-1;j++ {
			for k:=j+1;k<len(nums);k++ {
				if abs(nums[i]+nums[j]+nums[k]-target)<diff {
					diff = abs(nums[i]+nums[j]+nums[k]-target)
					res = nums[i]+nums[j]+nums[k]
				}
			}
		}
	}
	return res
}
func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}
//方案二：o(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) //[-4,-1,1,2]
	res,diff := 0,math.MaxInt32
	for i:=0;i<len(nums)-2;i++ {
		// 避免重复计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j,k:=i+1,len(nums)-1;
		for j<k {
			sum := nums[i]+nums[j]+nums[k]
			if abs(sum-target)<diff {
				res,diff = sum,abs(sum-target)
			}
			if sum==target {
				return res
			}else if sum<target {
				j++
			}else {
				k--
			}
		}
	}
	return res
}
