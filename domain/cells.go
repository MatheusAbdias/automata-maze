package domain

var cellColor = map[int]string{
	0: "WHITE",
	1: "GREEN",
	3: "RED",
	4: "BLUE",
}

type Location struct {
	x int
	y int
}

type Cell struct {
	color    string
	location Location
}

func (c *Cell) setColor(value int) {
	c.color = cellColor[value]
}

func (c *Cell) setLocation(x int, y int) {
	c.location.x = x
	c.location.y = y
}

func InitCell(x int, y int, value int) *Cell {
	cell := Cell{}
	cell.setColor(value)
	cell.setLocation(x, y)

	return &cell
}
