import "sort"

func nextPermutation(nums []int)  {
    n := len(nums)
    maxi := math.MinInt
    ind := -1
    for i:=n-1;i>=0;i-- {
        if nums[i] < maxi {
            ind = i+1
            break
        }
        if nums[i] > maxi {
            maxi = nums[i]
        }
    }
    if ind == -1 {
        sort.Ints(nums)
    } else {
        sort.Ints(nums[ind:])
        te := ind-1
        for i:=ind;i<len(nums);i++ {
            if nums[te] < nums[i] {
                nums[te], nums[i] = nums[i], nums[te]
                break
            }
        }
    }
}