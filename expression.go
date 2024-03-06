package typescript

type Expression struct {
	IsLiteral bool
	IsName    bool
	Value     string

	IsCall bool
	Fn     string
	Inputs []*Expression
}
