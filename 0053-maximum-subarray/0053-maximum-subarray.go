func maxSubArray(nums []int) int {
    ret := math.MinInt
    run := 0
    for i:=0;i<len(nums);i++ {
        if nums[i] > ret {
            ret = nums[i]
        }
        run = run + nums[i]
        if run > ret {
            ret = run
        }
        if run < 0 {
            run = 0
        }
    }
    return ret
}