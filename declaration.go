package typescript

import "io"

type Declaration struct {
	Name     string
	Type     string
	Variable bool
	Value    *Expression
}

func (decl *Declaration) Write(w io.Writer) error
