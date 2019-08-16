package kata

func CartesianNeighbor(x,y int) [][]int{
  var neighbors [][]int
  
  for i := -1; i < 2; i++ {
    for j := -1; j < 2; j++ {
      if i == 0 && j == 0 { continue }
      neighbors = append(neighbors, []int{x+i, y+j})
    }
  }
  return neighbors
}