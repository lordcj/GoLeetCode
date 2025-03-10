func maxProfit(prices []int) int {
    ans := 0.0
    mini := float64(math.MaxInt)
    for _, val := range prices {
        ans = max(ans, float64(val) - mini)
        mini = min(mini, float64(val))
    }
    return int(ans)
}