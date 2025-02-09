func MaxSubArraySum(data []int) int {
    
    //Implement your solution here
    max := 0
    sum := 0

    for i := 0; i < len(data); i++ {
        sum = sum + data[i]
        if sum < 0{
          sum = 0
        } 
        if sum > max {
          max = sum
        }
    }
    return max
}