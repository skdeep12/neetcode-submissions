import "slices"
func threeSum(nums []int) [][]int {
	finalAns := make([][]int, 0)
	
	slices.Sort(nums)
	for i:=0;i<len(nums)-2;i+=1{
		ans := twoSum(nums, i+1, -nums[i])
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for _,j := range ans {
			finalAns = append(finalAns, append([]int{nums[i]}, j...))
		}		
	}
	return finalAns
}

func twoSum(nums []int, start int, target int) [][]int{
	end := len(nums) - 1
	ans := make([][]int, 0)
	for start < end {
		sum := nums[start] + nums[end]
		if sum > target {
			end-=1
		} else if sum < target {
			start += 1
		} else {
			ans = append(ans, []int{nums[start], nums[end]})
			left := nums[start]
			right := nums[end]
			for start < end && nums[start] == left {
				start+=1
			}
			for start < end && nums[end-1] == right {
				end-=1
			}
		}
	}
	return ans
}