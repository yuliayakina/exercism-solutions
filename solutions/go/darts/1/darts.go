package darts

func Score(x, y float64) int {
	radiusSquared := (x * x) + (y * y)
    if radiusSquared > 100 {
        return 0
    }else if radiusSquared > 25 {
        return 1
    }else if radiusSquared > 1 {
        return 5
    }
    return 10
}
