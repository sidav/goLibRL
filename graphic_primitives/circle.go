package graphic_primitives

func GetCircle(x, y, r int) *[]point {
	points := make([]point, 0)
	if r < 0 {
		return nil
	}
	// Bresenham algorithm
	x1, y1, err := -r, 0, 2-2*r
	for {
		points = append(points, point{x-x1, y+y1}) // lower right
		points = append(points, point{x-y1, y-x1}) // lower left
		points = append(points, point{x+x1, y-y1}) // upper left
		points = append(points, point{x+y1, y+x1}) // upper right
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
	return &points
}

func GetApproxCircleAroundRect(x, y, w, h, r int) *[]point { // in fact, it is more like a square with very rounded corners.
	points := make([]point, 0)
	if r < 0 {
		return nil
	}

	for x1:=x;x1<x+w;x1++{
		points = append(points, point{x1, y-r}, point{x1, y+h+r-1})
	}
	for y1:=y;y1<y+h;y1++{
		points = append(points, point{x-r, y1}, point{x+w+r-1, y1})
	}

	// Bresenham algorithm
	x1, y1, err := -r, 0, 2-2*r
	for {
		r = err
		if r > x1 {
			x1++
			err += x1*2 + 1
		}
		if r <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
		points = append(points, point{x-x1+w-1, y+y1+h-1}) // lower right
		points = append(points, point{x-y1, y-x1+h-1}) // lower left
		points = append(points, point{x+x1, y-y1}) // upper left
		points = append(points, point{x+y1+w-1, y+x1}) // upper right
	}
	return &points
}
