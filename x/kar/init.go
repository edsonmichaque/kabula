package kar

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/edsonmichaque/kabula/x/spec"
)

const (
	InitialVersion = "0.1.0"
)

func (k Kab) Init() error {
	if err := os.Mkdir(k.src, 0755); err != nil {
		if errors.Is(err, os.ErrExist) {
			return fmt.Errorf("%s already exist", k.src)
		}

		return err
	}

	karFile, err := os.Create(k.ManifestPath())
	if err != nil {
		return err
	}
	defer karFile.Close()

	manifest := spec.Manifest{
		Version: InitialVersion,
		Name:    filepath.Base(k.src),
	}

	rawData, err := k.encode(manifest)
	if err != nil {
		return err
	}

	if _, err := karFile.Write(rawData); err != nil {
		return err
	}

	return nil
}
