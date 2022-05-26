package kab

import (
	"errors"
	"fmt"
	"os"

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

	dstFile, err := os.Create(k.ManifestPath())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	manifest := spec.Manifest{
		Version: InitialVersion,
		Name:    k.src,
	}

	rawData, err := k.Marshal(manifest)
	if err != nil {
		return err
	}

	if _, err := dstFile.Write(rawData); err != nil {
		return err
	}

	return nil
}
