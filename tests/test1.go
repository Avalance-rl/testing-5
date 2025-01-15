package tests


func PositivePow(number int, degree uint) int {
    res := 1
    for _ = range degree {
        res *= number 
    }
    return res
}
