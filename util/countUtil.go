package util

import (
	"strings"
)

// CountUtil ...
type CountUtil struct {
	Text      string
	Unique    bool
	Substring string
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

// SubstringCount ...
func (c *CountUtil) SubstringCount() int {
	return strings.Count(strings.ToLower(c.Text), strings.ToLower(c.Substring))
}
