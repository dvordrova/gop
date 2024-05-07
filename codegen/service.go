package codegen

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dvordrova/gop/config"
	"github.com/dvordrova/gop/constants"
)

var (
	//go:embed templates
	templatesFolder embed.FS
)

var DependenciesVersions = map[string]string{
	"github.com/go-fuego/fuego":                                       "v0.13.4",
	"go.opentelemetry.io/otel":                                        "v1.25.0",
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp": "v1.25.0",
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric":          "v1.25.0",
	"go.opentelemetry.io/otel/sdk":                                    "v1.25.0",
	"go.opentelemetry.io/otel/sdk/metric":                             "v1.25.0",
	"go.opentelemetry.io/otel/trace":                                  "v1.25.0",
}

func fillTemplate(cfg *config.ServiceConfigFile) error {
	outputDir := "."
	err := fs.WalkDir(templatesFolder, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking through templates: %v", err)
		}

		relPath, err := filepath.Rel("templates", path)
		if err != nil {
			return fmt.Errorf("error rel path: %v", err)
		}
		outputPath := strings.TrimSuffix(filepath.Join(outputDir, relPath), ".tmpl")
		if d.IsDir() {
			if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
				return fmt.Errorf("error create dir %s: %v", path, err)
			}
			return nil
		}

		templateBytes, err := templatesFolder.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error read file %s: %v", path, err)
		}

		templateContent := string(templateBytes)
		tmpl, err := template.New(filepath.Base(path)).Parse(templateContent)
		if err != nil {
			return fmt.Errorf("error parse template %s: %v", path, err)
		}
		outputFile, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("error create %s: %v", outputPath, err)
		}
		defer outputFile.Close()
		if err := tmpl.Execute(outputFile, cfg); err != nil {
			return fmt.Errorf("tmpl execute: %v", err)
		}

		slog.Info(fmt.Sprintf("Rendered and created: %s", outputPath))
		return nil
	})
	if err != nil {
		return fmt.Errorf("gen by template: %s", err)
	}
	return err
}

func initGoModule(moduleName string) error {
	// Construct the command "go mod init <module-name>"
	cmd := exec.Command("go", "mod", "init", moduleName)

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("init go module: %v", err)
	}

	return nil
}

func installDeps() error {
	for dep, version := range DependenciesVersions {
		cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", dep, version))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("install dep %s: %v", dep, err)
		}
	}
	return nil
}

func GenServiceStruct() error {
	var cfg config.ServiceConfigFile
	if _, err := toml.DecodeFile(".gop.app.toml", &cfg); err != nil {
		return fmt.Errorf("decode .gop.app.toml: %v", err)
	}
	cfg.GopVersionGen = constants.GopVersion
	cfg.GenerationDate = time.Now().Format(time.RFC1123)

	if err := fillTemplate(&cfg); err != nil {
		return err
	}
	if err := initGoModule(cfg.ModuleName); err != nil {
		return err
	}
	if err := installDeps(); err != nil {
		return err
	}

	return nil
}
