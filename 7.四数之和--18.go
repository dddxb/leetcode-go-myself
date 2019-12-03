package main

import (
	"sort"
)

//func main()  {
//	nums := []int{1,0,-1,0,-2,2}
//	target := 0
//	res := fourSum(nums,target)
//	fmt.Println(res)
//}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) //[-2,-1,0,0,1,2]
	s := make([][]int,0)
	for i:=0;i<len(nums)-3;i++ {
		if i>0&&nums[i]==nums[i-1] {
			continue
		}
		for j:=i+1;j<len(nums)-2;j++ {
			if j>i+1&&nums[j]==nums[j-1] {
				continue
			}
			l,r := j+1,len(nums)-1
			for l<r {
				sum := nums[i]+nums[j]+nums[l]+nums[r]
				switch {
				case sum<target:
					l++
				case sum>target:
					r--
				default:
					s = append(s,[]int{nums[i],nums[j],nums[l],nums[r]})
					l,r = next(nums,l,r)
				}
			}
		}
	}
	return s
}
func next(nums []int,l int,r int) (int,int) {
	for l<r {
		switch {
		case nums[l]==nums[l+1]:
			l++
		case nums[r]==nums[r-1]:
			r--
		default:
			l++
			r--
			return l,r
		}
	}
	return l,r
}