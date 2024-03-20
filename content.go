package crt

// The "pageContent" type represents a pageContent with a map of rows and columns.
// @property row - The "row" property is a map that stores the values of each row in the pageContent. The keys
// of the map are integers representing the row numbers, and the values are strings representing the
// content of each row.
// @property {int} cols - The "cols" property represents the number of columns in the pageContent.
// @property {int} rows - The "rows" property represents the number of rows in the pageContent.
type pageContent struct {
	row  map[int]string
	cols int
	rows int
}

// newPageDefinition initializes a new page with the specified number of columns and rows.
func (c *Crt) newPageDefinition(cols, rows int) {
	p := pageContent{}
	p.cols = cols
	p.rows = rows
	p.row = make(map[int]string)
	c.scr = &p
}
