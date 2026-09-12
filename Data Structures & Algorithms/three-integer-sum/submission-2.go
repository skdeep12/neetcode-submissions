import "slices"
func threeSum(nums []int) [][]int {
	finalAns := make([][]int, 0)
	ansMap := make(map[candidate]bool)
	slices.Sort(nums)
	for i:=0;i<len(nums)-2;i+=1{
		ans := twoSum(nums, i+1, -nums[i])
		for _,j := range ans {
			c := candidate{nums[i], j[0], j[1]}
			if _, ok := ansMap[c]; !ok {
				ansMap[c] = true
				finalAns = append(finalAns, c.ToArr())
			}
		}		
	}
	return finalAns
}

type candidate struct {
	one, two, three int
}

func (c *candidate) ToArr() []int {
	return []int{c.one,c.two,c.three}
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
			start+=1
		}
	}
	return ans
}