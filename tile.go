package tile

import (
	"errors"
)

var (
	TileOutOfBounds = errors.New("tile is out of bound")
)

type Tile interface {
}

type world struct {
	size  int
	tiles map[int]*Tile
}

func NewTileIndex(size int) *world {
	return &world{
		size: size,
	}
}

func (t *world) CoordinateFromID(id int) (*coordinate, error) {
	if id > (t.size*t.size - 1) {
		return nil, TileOutOfBounds
	}
	return &coordinate{
		size: t.size,
		x:    id % t.size,
		y:    id / t.size,
	}, nil
}

func (t *world) GetTile(id int) (*Tile, error) {
	if id > (t.size*t.size - 1) {
		return nil, TileOutOfBounds
	}

	return t.tiles[id], nil
}

func (t *world) Zoom(id int, level int) (map[int]map[int]*Tile, error) {
	coordinate, err := t.CoordinateFromID(id)

	tiles := make(map[int]map[int]*Tile)
	if err != nil {
		return nil, err
	}

	for x := coordinate.Id() - (level / 2); x < (coordinate.Id()%t.size)+(level/2); x++ {
		for y := coordinate.Id() - (level / 2); y < (coordinate.Id()/t.size)+(level/2); y++ {
			tile, err := NewCoordinate(t.size, x, y)

			if err != nil {
				return nil, err
			}
			tiles[x][y] = t.tiles[tile.Id()]
		}
	}
	return tiles, nil
}
