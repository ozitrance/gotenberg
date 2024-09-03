package pdftocairo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/gotenberg/gotenberg/v8/pkg/gotenberg"
)

func init() {
	gotenberg.MustRegisterModule(new(PdfToCairo))
}

// PdfToCairo abstracts the CLI tool PdfToCairo and implements the [gotenberg.PdfEngine]
// interface.
type PdfToCairo struct {
	binPath string
}

// Descriptor returns a [PdfToCairo]'s module descriptor.
func (engine *PdfToCairo) Descriptor() gotenberg.ModuleDescriptor {
	return gotenberg.ModuleDescriptor{
		ID:  "pdftocairo",
		New: func() gotenberg.Module { return new(PdfToCairo) },
	}
}

// Provision sets the modules properties.
func (engine *PdfToCairo) Provision(ctx *gotenberg.Context) error {
	binPath, ok := os.LookupEnv("PDFTOCAIRO_BIN_PATH")
	if !ok {
		return errors.New("PDFTOCAIRO_BIN_PATH environment variable is not set")
	}

	engine.binPath = binPath

	return nil
}

// Validate validates the module properties.
func (engine *PdfToCairo) Validate() error {
	_, err := os.Stat(engine.binPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("PdfToCairo binary path does not exist: %w", err)
	}

	return nil
}

// Merge combines multiple PDFs into a single PDF.
func (engine *PdfToCairo) Merge(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string) error {
	return fmt.Errorf("Merge PDF to '%+v' with PdfToCairo: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}


// Linearize the PDF for Fast Web.
func (engine *PdfToCairo) Linearize(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string) error {
	return fmt.Errorf("Linearize PDF to '%+v' with PdfToCairo: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

// Thumbnail the PDF for Fast Web.
func (engine *PdfToCairo) Thumbnail(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string, page string) error {
	var args []string
	// args = append(args, "--pages")
	out := strings.Split(outputPath, ".jpg")
	args = append(args, inputPaths...)
	args = append(args, out[0])
	args = append(args, "-jpeg")
	args = append(args, "-f", page)
	args = append(args, "-l", page)
	args = append(args, "-scale-to", "400")
	args = append(args, "-singlefile")

	cmd, err := gotenberg.CommandContext(ctx, logger, engine.binPath, args...)

	if err != nil {
		return fmt.Errorf("create command: %w", err)
	}

	_, err = cmd.Exec()
	if err == nil {
		return nil
	}

	return fmt.Errorf("Thumbnail PDFs with pdftocairo: %w", err)
}

// Thumbnail the PDF for Fast Web.
func (engine *PdfToCairo) PNG(ctx context.Context, logger *zap.Logger, inputPaths []string, outputPath string, page string, monochrome bool) error {
	var args []string
	// args = append(args, "--pages")
	out := strings.Split(outputPath, ".png")
	args = append(args, inputPaths...)
	args = append(args, out[0])
	args = append(args, "-png")
	args = append(args, "-f", page)
	args = append(args, "-l", page)
	args = append(args, "-singlefile")
	if monochrome == true {
		args = append(args, "-mono")
	}

	cmd, err := gotenberg.CommandContext(ctx, logger, engine.binPath, args...)

	if err != nil {
		return fmt.Errorf("create command: %w", err, monochrome)
	}

	_, err = cmd.Exec()
	if err == nil {
		return nil
	}

	return fmt.Errorf("PNG PDFs with pdftocairo: %w", err)
}

// Convert is not available in this implementation.
func (engine *PdfToCairo) Convert(ctx context.Context, logger *zap.Logger, formats gotenberg.PdfFormats, inputPath, outputPath string) error {
	return fmt.Errorf("convert PDF to '%+v' with PdfToCairo: %w", formats, gotenberg.ErrPdfEngineMethodNotSupported)
}

// ReadMetadata is not available in this implementation.
func (engine *PdfToCairo) ReadMetadata(ctx context.Context, logger *zap.Logger, inputPath string) (map[string]interface{}, error) {
	return nil, fmt.Errorf("read PDF metadata with PdfToCairo: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

// WriteMetadata is not available in this implementation.
func (engine *PdfToCairo) WriteMetadata(ctx context.Context, logger *zap.Logger, metadata map[string]interface{}, inputPath string) error {
	return fmt.Errorf("write PDF metadata with PdfToCairo: %w", gotenberg.ErrPdfEngineMethodNotSupported)
}

var (
	_ gotenberg.Module      = (*PdfToCairo)(nil)
	_ gotenberg.Provisioner = (*PdfToCairo)(nil)
	_ gotenberg.Validator   = (*PdfToCairo)(nil)
	_ gotenberg.PdfEngine   = (*PdfToCairo)(nil)
)
