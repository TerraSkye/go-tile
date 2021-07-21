package tile

import (
	"testing"
)

func TestCoordinateOverflowing(t *testing.T) {
	d, _ := NewCoordinate(11, -5, -5)

	if d.x != -5 {
		t.Errorf("tile id should be 0 but got %d", d.x)
	}

	e, _ := NewCoordinate(11, -6, -5)

	if e.x != 5 {
		t.Errorf("tile id should be 0 but got %d", e.x)
	}

}

func TestTileIDResolving(t *testing.T) {
	c, _ := NewCoordinate(11, -5, -5)

	if c.Id() != 0 {
		t.Errorf("tile id should be 0 but got %d", c.Id())
	}

	d, _ := NewCoordinate(21, -5, -5)
	if d.Id() != 110 {
		t.Errorf("tile id should be 110 but got %d", d.Id())
	}

	e, _ := NewCoordinate(21, 0, 0)

	if e.Id() != 220 {
		t.Errorf("tile id should be 110 but got %d", e.Id())
	}

	f, _ := NewCoordinate(21, 10, 10)

	if f.Id() != 440 {
		t.Errorf("tile id should be 441 but got %d", f.Id())
	}

	g, _ := NewCoordinate(21, -10, -10)

	if g.Id() != 0 {
		t.Errorf("tile id should be 441 but got %d", g.Id())
	}
}

func TestCoordinateCalculations(t *testing.T) {

	c, _ := NewCoordinate(11, -5, -5)

	if outcome := c.Relative(5, 0); outcome.x != 0 {
		t.Errorf("expected 0 but got %d", outcome)
	}

	if outcome := c.Relative(1, 0); outcome.x != -4 {
		t.Errorf("expected -4 but got %d", outcome)
	}

	if outcome := c.Relative(-1, 0); outcome.x != 5 {
		t.Errorf("expected 5 but got %d", outcome)
	}

	d, _ := NewCoordinate(11, 5, 0)
	if outcome := d.Relative(1, 0); outcome.x != -5 {
		t.Errorf("expected -5 but got %d", outcome)
	}

	if outcome := d.Relative(-17, 0); outcome.x != 0 {
		t.Errorf("expected 0 but got %d", outcome)
	}

}
