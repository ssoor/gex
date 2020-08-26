package tool

import (
	"go/parser"
	"go/token"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/ssoor/gex/pkg/manager"
)

// Parser retrieve tool packages from given paths.
type Parser interface {
	Parse(path string) (*Manifest, error)
}

// NewParser creates a new parser instance.
func NewParser(fs afero.Fs, mType manager.Type) Parser {
	return &parserImpl{
		fs:    fs,
		mType: mType,
	}
}

type parserImpl struct {
	fs    afero.Fs
	mType manager.Type
}

func (p *parserImpl) Parse(path string) (*Manifest, error) {
	data, err := afero.ReadFile(p.fs, path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %q", path)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", string(data), parser.ParseComments)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse %q", path)
	}

	tools := make([]Tool, 0, len(f.Imports))

	for _, s := range f.Imports {
		if pkg, err := strconv.Unquote(s.Path.Value); err == nil {
			comment := ""
			if s.Comment != nil {
				comment = s.Comment.Text()
			}
			tools = append(tools, Tool{path: pkg, Global: comment == "Global Tools\n"})
		}
	}

	return NewManifest(tools, p.mType), nil
}
