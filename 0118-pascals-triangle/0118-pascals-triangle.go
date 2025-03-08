func generate(numRows int) [][]int {
    ret := [][]int{}
    ret = append(ret, []int{1})
    for i:=1;i<numRows;i++ {
        arr := []int{1}
        for j:=1;j<i;j++ {
            arr = append(arr, ret[i-1][j-1] + ret[i-1][j])
        }
        arr = append(arr, 1)
        ret = append(ret, arr)

    }
    return ret
}