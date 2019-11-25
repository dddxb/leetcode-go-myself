package main

//func main() {
//	nums := []int{4,3,2,7,8,2,3,1}
//	result1 := findDuplicates1(nums)
//	fmt.Println(result1)
//	result2 := findDuplicates2(nums)
//	fmt.Println(result2)
//}
//方案一：利用map键值的唯一性，判断map里是否含有某键值，只需o(1)时间
//时间复杂度o(n),空间复杂度就是新建map所占用的空间
func findDuplicates1(nums []int) []int {
	m := make(map[int]int)
	result := make([]int,0)
	for i:=0;i<len(nums);i++ {
		_,ok := m[nums[i]]
		if ok {
			result = append(result,nums[i])
		}else {}
		m[nums[i]] = nums[i]
	}
	return result
}
//方案二：利用每一位上数字的正负号来表示此位置坐标大小的数之前是否出现过
//时间复杂度o(n),空间复杂度o(1)
func findDuplicates2(nums []int) []int{
	result := make([]int,0)
	var temp int
	for i:=0;i<len(nums);i++{
		if nums[i]>0 {
			temp = nums[i]
		}else {
			temp = nums[i]*(-1)
		}
		if nums[temp-1]>0 {
			nums[temp-1] *= -1
		}else {
			result = append(result,temp)
		}
	}
	return result
}