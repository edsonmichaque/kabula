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
		License: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
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
