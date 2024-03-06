package typescript

import "go/ast"

type Package struct {
	Imports       *ImportMap
	Declarations  []*Declaration
	DefaultExport *Expression
	NamedExports  map[string]*Expression
}

func NewPackageFromGo(pkg *ast.Package) (*Package, error)

func (p *Package) WriteToDir(dir string) error
