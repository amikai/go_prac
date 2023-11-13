package other

type Opt func(*Computer)

type Computer struct {
	CPUNum int
	Memory string
}

func NewComputer(opts ...Opt) *Computer {
	c := &Computer{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithCPUNum(num int) Opt {
	return func(c *Computer) {
		c.CPUNum = num
	}
}

func WithMemory(cap string) Opt {
	return func(c *Computer) {
		c.Memory = cap
	}
}
