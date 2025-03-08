func setZeroes(matrix [][]int)  {
    n := len(matrix)
    m := len(matrix[0])
    // 0 -> down, 1 -> right , 2 -> both
    x := -1
    y := -1
    if matrix[0][0] == 0 {
        x, y = 1,1
    } else {
        for i:=1;i<n;i++ {
            if matrix[i][0] == 0 {
                x = 1
            }
        }
        for j:=1;j<m;j++ {
            if matrix[0][j] == 0 {
                y = 1
            }
        }
    }
    for i:=1;i<n;i++ {
        for j:=1;j<m;j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i:=1;i<n;i++ {
        for j:=1;j<m;j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if x == 1 {
        for i:=0;i<n;i++ {
            matrix[i][0] = 0
        }
    }
    if y == 1 {
        for j:=0;j<m;j++ {
            matrix[0][j] = 0
        }
    }
}