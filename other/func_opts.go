package other

type Computer struct {
	CPUNum int
	Memory string
}

type Opt func(*Computer)

func WithCPUNum(num int) Opt {
	return func(c *Computer) {
		c.CPUNum = num
	}
}

func WithCPUMemory(cap string) Opt {
	return func(c *Computer) {
		c.Memory = cap
	}
}

func NewComputer(opts ...Opt) *Computer {
	c := &Computer{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
