func sortColors(nums []int)  {
    l, r, n := 0, len(nums)-1, len(nums)
    for l<n {
        if nums[l] == 0 {
            l++
        } else {
            break
        }
    }
    for r>=0 {
        if nums[r] == 2 {
            r--
        } else {
            break
        }
    }
    ind := l
    for ind <= r {
        if nums[ind] == 0 {
            nums[l], nums[ind] = nums[ind], nums[l]
            for l < n {
                if nums[l] == 0 {
                    l++
                } else {
                    break
                }
            }
        } else if nums[ind] == 2 {
            nums[r], nums[ind] = nums[ind], nums[r]
            for r >= 0 {
                if nums[r] == 2 {
                    r--
                } else {
                    break
                }
            }
        }
        for ind <= r {
            if nums[ind] == 1 || ind < l {
                ind++
            } else {
                break
            }
        }
    }
}