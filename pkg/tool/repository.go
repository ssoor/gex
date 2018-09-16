package tool

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/afero"

	"github.com/izumin5210/gex/pkg/command"
)

// Repository is an interface for managing and operating tools
type Repository interface {
	Add(ctx context.Context, pkgs ...string) error
	Build(ctx context.Context, t Tool) (string, error)
	BuildAll(ctx context.Context) error
	Run(ctx context.Context, name string, args ...string) error
}

type repositoryImpl struct {
	*Config
	fs       afero.Fs
	parser   Parser
	writer   Writer
	executor command.Executor
}

// NewRepository creates a new Repository instance.
func NewRepository(fs afero.Fs, executor command.Executor, cfg *Config) Repository {
	return &repositoryImpl{
		Config:   cfg,
		fs:       fs,
		parser:   NewParser(fs),
		writer:   NewWriter(fs),
		executor: executor,
	}
}

func (r *repositoryImpl) Add(ctx context.Context, pkgs ...string) error {
	args := []string{"get"}
	if r.Verbose {
		args = append(args, "-v")
	}
	args = append(args, pkgs...)
	err := r.executor.Exec(ctx, "go", args...)
	if err != nil {
		return err
	}

	m, err := r.parser.Parse(r.ManifestPath())
	if err != nil {
		m = NewManifest([]Tool{})
	}

	for _, pkg := range pkgs {
		pkg = strings.SplitN(pkg, "@", 2)[0]
		t := Tool(pkg)
		m.AddTool(t)
		_, err = r.Build(ctx, t)
		if err != nil {
			return err
		}
	}

	err = r.writer.Write(r.ManifestPath(), m)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) Build(ctx context.Context, t Tool) (string, error) {
	binPath := r.BinPath(t.Name())

	if st, err := r.fs.Stat(binPath); err != nil {
		args := []string{"build", "-o", binPath}
		if r.Verbose {
			args = append(args, "-v")
		}
		args = append(args, string(t))
		err := r.executor.Exec(ctx, "go", args...)
		if err != nil {
			return "", err
		}
	} else if st.IsDir() {
		return "", fmt.Errorf("%q is a directory", t.Name())
	}

	return binPath, nil
}

func (r *repositoryImpl) BuildAll(ctx context.Context) error {
	m, err := r.parser.Parse(r.ManifestPath())
	if err != nil {
		return err
	}

	for _, t := range m.Tools() {
		_, err = r.Build(ctx, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repositoryImpl) Run(ctx context.Context, name string, args ...string) error {
	m, err := r.parser.Parse(r.ManifestPath())
	if err != nil {
		return err
	}

	t, ok := m.FindTool(name)
	if !ok {
		return fmt.Errorf("failed to find the tool %q", name)
	}

	bin, err := r.Build(ctx, t)
	if err != nil {
		return err
	}

	return r.executor.Exec(ctx, bin, args...)
}