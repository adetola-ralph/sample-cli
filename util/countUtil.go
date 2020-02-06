package util

// CountUtil ...
type CountUtil struct {
	Text   string
	Unique bool
}

var listUtil = ListUtil{}

// Words ...
func (c *CountUtil) Words() int {
	listUtil.Text = c.Text
	listUtil.Unique = c.Unique

	wordSlice := listUtil.Words()
	return len(wordSlice)
}

// Chars ...
func (c *CountUtil) Chars() int {
	listUtil.Text = c.Text
	listUtil.Unique = c.Unique

	charSlice := listUtil.Chars()
	return len(charSlice)
}

// Lines ...
func (c *CountUtil) Lines() int {
	listUtil.Text = c.Text
	listUtil.Unique = c.Unique

	charSlice := listUtil.Lines()
	return len(charSlice)
}
