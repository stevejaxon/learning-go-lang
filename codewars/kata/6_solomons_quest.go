package kata

import "math"
import "math/big"

func SolomonsQuest(ar [][3]int) [2]int {
	layer := 0
	var x big.Int
	var y big.Int
	
	for _, tracks := range ar {
		layer += tracks[0]
		scale := math.Pow(2, float64(layer))
		direction := tracks[1]
		var distance big.Int
		distance.Mul(big.NewInt(int64(scale)), big.NewInt(int64(tracks[2])))
		switch direction {
			case 0:
				y.Add(&y, &distance)
			case 1:
				x.Add(&x, &distance)
			case 2:
				y.Sub(&y, &distance)
			case 3:
				x.Sub(&x, &distance)
		}
	}
	return [2]int{int(x.Int64()),int(y.Int64())}
}