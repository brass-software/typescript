package typescript

type Declaration struct {
	Name     string
	Type     string
	Variable bool
	Value    *Expression
}
