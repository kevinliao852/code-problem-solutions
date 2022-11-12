func countPoints(points [][]int, queries [][]int) []int {
	var res []int

	for i := 0; i < len(queries); i++ {
		cnt := 0
		x, y, r := queries[i][0], queries[i][1], queries[i][2]
		for j := 0; j < len(points); j++ {
			xt, yt := points[j][0], points[j][1]
			d := (x-xt)*(x-xt) + (y-yt)*(y-yt)

			if d <= r*r {
				cnt++
			}
		}
		res = append(res, cnt)
	}
	return res
} 
