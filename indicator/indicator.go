package indicator

type CaptureLoadingIndicator struct {
	indicators []string
	length     int
	step       int
}

func (c *CaptureLoadingIndicator) initialize(indicators []string) {
	c.indicators = indicators
	c.length = len(c.indicators)
	c.step = c.length - 1
}

func (c *CaptureLoadingIndicator) next() string {
	c.step = ((c.step + 1) % c.length)
	return c.indicators[c.step]
}

func CreateCaptureLoadingIndicator() *CaptureLoadingIndicator {
	c := (&CaptureLoadingIndicator{})
	c.initialize([]string{"▖", "▘", "▝", "▗"})

	return c
}

