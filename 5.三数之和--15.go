package main

//func main() {
//	l := list.New()
//	l.PushFront("ha")
//	l.PushBack("yy")
//	for i:=l.Front();i!=nil;i=i.Next() {
//		fmt.Println(i.Value)
//	}
//	for j:=l.Back();j!=nil;j=j.Prev() {
//		fmt.Println(j.Value)
//	}
//}
import (
	"sort"
)

//func main() {
//	nums := []int{-1,0,1,2,-1,4}
//	res := threeSum(nums)
//	fmt.Println(res)
//}
func threeSum(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := make([][]int,0)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// 避免添加重复的结果 i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

//func next(nums []int, l, r int) (int, int) {
//	for l < r { //不断循环，目的是去重，即重复的滤过不参与计算
//		switch {
//		case nums[l] == nums[l+1]:
//			l++
//		case nums[r] == nums[r-1]:
//			r--
//		default:
//			l++
//			r--
//			return l, r
//		}
//	}
//
//	return l, r
//}

//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	re := make([][]int,0)
//	if len(nums)<3 {
//		return re
//	}
//	for i:=0;nums[i]<=0&&i<=len(nums)-2;i++ {
//		if i>0&&nums[i]==nums[i-1] { //去重
//			continue
//		}
//		l,r := i+1,len(nums)-1
//		for l<r { //限制条件
//			sum := nums[i]+nums[l]+nums[r]
//			switch {
//			case sum>0:
//				r--
//			case sum<0:
//				l++
//			default:
//				re = append(re,[]int{nums[i],nums[l],nums[r]})
//				l,r = next(nums,l,r) //l,r更改位置，避免重复
//			}
//		}
//	}
//	return re
//}
//func next(nums []int,l int,r int) (int,int) {
//	for l<r {
//		switch {
//		case nums[l]==nums[l+1]:
//			l++
//		case nums[r]==nums[r-1]:
//			r--
//		default:
//			l++
//			r--
//			return l,r
//		}
//	}
//	return l,r
//}
//func threeSum(nums []int) [][]int {
//	res := [][]int{}
//	counter := map[int]int{}
//	for _, value := range nums {
//		counter[value]++
//	}
//
//	uniqNums := []int{}
//	for key := range counter {
//		uniqNums = append(uniqNums, key)
//	}
//	sort.Ints(uniqNums)
//
//	for i := 0; i < len(uniqNums); i++ {
//		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
//			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
//		}
//		for j := i + 1; j < len(uniqNums); j++ {
//			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
//				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
//			}
//			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
//				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
//			}
//			c := 0 - uniqNums[i] - uniqNums[j]
//			if c > uniqNums[j] && counter[c] > 0 {
//				res = append(res, []int{uniqNums[i], uniqNums[j], c})
//			}
//		}
//	}
//	return res
//}