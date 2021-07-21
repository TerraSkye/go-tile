package tile

import "errors"

var InvalidSize = errors.New("invalid size")

type Coordinate struct {
	coordinate
}

type coordinate struct {
	size int
	x    int
	y    int
}

func NewCoordinate(size int, x int, y int) (*coordinate, error) {
	if size%2 == 0 {
		return nil, InvalidSize
	}
	c := &coordinate{
		size: size,
		x:    abs(size, x),
		y:    abs(size, y),
	}
	return c, nil
}

func (c *coordinate) Id() int {
	equator := c.size / 2

	return ((c.x + equator) * c.size) + c.y + equator
}

func (c coordinate) Relative(x int, y int) *coordinate {

	if x < 0 && c.x < 0 || x > 0 && c.x > 0 || x > 0 && c.x < 0 {
		c.x = abs(c.size, c.x+x)
	} else {
		c.x = abs(c.size, c.x-x)
	}

	if y < 0 && c.y < 0 || y > 0 && c.y > 0 || y > 0 && c.y < 0 {
		c.y = abs(c.size, c.y+y)
	} else {
		c.y = abs(c.size, c.y-y)
	}

	return &c
}

func abs(size int, value int) int {
	equator := size / 2
	// since we work with negative coordinates, we add half the map to it, that way we get positive coordinates only
	absoluteCoordinate := equator + value
	// we do modules of the map
	absoluteCoordinate = absoluteCoordinate % size
	// now we remove te equator again to get the negative
	absoluteCoordinate = absoluteCoordinate - equator
	// if the absolute coordinate is bigger then half the map, we have to add the size once more.
	if absoluteCoordinate < equator*-1 {
		absoluteCoordinate = absoluteCoordinate + size
	}
	// now we have resolved the overflow of a coordinate.
	return absoluteCoordinate
}
