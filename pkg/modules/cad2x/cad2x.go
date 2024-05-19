package cad2x

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/gotenberg/gotenberg/v8/pkg/gotenberg"
)

func init() {
	gotenberg.MustRegisterModule(new(Cad2X))
}

// QPdf abstracts the CLI tool QPDF and implements the [gotenberg.PdfEngine]
// interface.
type Cad2X struct {
	binPath string
}

// Descriptor returns a [Cad2X]'s module descriptor.
func (engine *Cad2X) Descriptor() gotenberg.ModuleDescriptor {
	return gotenberg.ModuleDescriptor{
		ID:  "cad2x",
		New: func() gotenberg.Module { return new(Cad2X) },
	}
}

// Provision sets the modules properties.
func (engine *Cad2X) Provision(ctx *gotenberg.Context) error {
	binPath, ok := os.LookupEnv("CAD2X_BIN_PATH")
	if !ok {
		return errors.New("CAD2X_BIN_PATH environment variable is not set")
	}

	engine.binPath = binPath

	return nil
}

// Validate validates the module properties.
func (engine *Cad2X) Validate() error {
	_, err := os.Stat(engine.binPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Cad2X binary path does not exist: %w", err)
	}

	return nil
}

// Merge combines multiple PDFs into a single PDF.
func (engine *Cad2X) Merge(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string) error {
	return fmt.Errorf("convert PDF to '%+v' with Cad2X: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}


// Convert DWG/DXF to PDF.
func (engine *Cad2X) Convert(ctx context.Context, logger *zap.Logger, formats gotenberg.PdfFormats, inputPath, outputPath string) error {
	var args []string
	args = append(args, "-o", outputPath)
	args = append(args, inputPath)
	args = append(args, "-abc")

	cmd, err := gotenberg.CommandContext(ctx, logger, engine.binPath, args...)

	if err != nil {
		return fmt.Errorf("create command: %w", err)
	}

	_, err = cmd.Exec()
	if err == nil {
		return nil
	}

	return fmt.Errorf("convert to PDFs with CAD2X: %w", err)
}

// Convert is not available in this implementation.
func (engine *Cad2X) Linearize(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string) error {
	return fmt.Errorf("Linearize PDF to '%+v' with Cad2X: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

// ReadMetadata is not available in this implementation.
func (engine *Cad2X) ReadMetadata(ctx context.Context, logger *zap.Logger, inputPath string) (map[string]interface{}, error) {
	return nil, fmt.Errorf("read PDF metadata with Cad2X: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

// WriteMetadata is not available in this implementation.
func (engine *Cad2X) WriteMetadata(ctx context.Context, logger *zap.Logger, metadata map[string]interface{}, inputPath string) error {
	return fmt.Errorf("write PDF metadata with Cad2X: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

var (
	_ gotenberg.Module      = (*Cad2X)(nil)
	_ gotenberg.Provisioner = (*Cad2X)(nil)
	_ gotenberg.Validator   = (*Cad2X)(nil)
	_ gotenberg.PdfEngine   = (*Cad2X)(nil)
)
