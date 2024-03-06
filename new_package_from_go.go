package typescript

import "go/ast"

func NewPackageFromGo(pkg *ast.Package) (*Package, error)
