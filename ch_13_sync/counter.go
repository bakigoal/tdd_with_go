package ch_13_sync

type Counter struct {
	value int
}

func (c *Counter) Inc() {
	c.value += 1
}

func (c Counter) Value() int {
	return c.value
}
