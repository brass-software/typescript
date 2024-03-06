package typescript

import (
	"fmt"
	"io"

	"github.com/brass-software/nodejs"
)

type ImportMap struct {
	Imports []*nodejs.Import
}

func (m *ImportMap) Add(i *nodejs.Import) error {
	for _, imp := range m.Imports {
		if imp.From == i.From {
			if i.Default != "" {
				if imp.Default != "" && imp.Default != i.Default {
					return fmt.Errorf("default import %s already exists from %s", i.Default, i.From)
				}
				imp.Default = i.Default
			}
			return nil
		}
	}
	m.Imports = append(m.Imports, i)
	return nil
}

func (m *ImportMap) Write(w io.Writer) error
