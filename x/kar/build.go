package kar

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/edsonmichaque/kabula/x/kar/builder"
)

type Options struct {
	Content   content
	Container archiveFormat
	Log       bool
}

var DefaultOptions = Options{
	Content:   JSON,
	Container: GZip,
}

func (kar Kab) Build() error {
	dirInfo, err := os.Stat(kar.src)
	if err != nil {
		return err
	}

	if !dirInfo.IsDir() {
		return fmt.Errorf("%s is not a dir", kar.src)
	}

	dir, err := os.Open(kar.src)
	if err != nil {
		return err
	}
	defer dir.Close()

	entries := make([]builder.Entry, 0)

	if err := filepath.Walk(kar.src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if kar.log {
			fmt.Println("found", info.Name())
		}

		entries = append(entries, builder.Entry{
			Name: info.Name(),
			Size: info.Size(),
		})

		return nil
	}); err != nil {
		return err
	}

	if err := kar.walk(entries); err != nil {
		return err
	}

	manifest, err := kar.decode()
	if err != nil {
		return err
	}

	dst, err := kar.create(manifest)
	if err != nil {
		return err
	}
	defer dst.Close()

	b := builders()

	build, ok := b[kar.archive]
	if !ok {
		build = builder.TGz
	}

	if err := build(kar.src, dst, entries); err != nil {
		if err := os.Remove(dst.Name()); err != nil {
			fmt.Println("error removing")
		}

		return err
	}

	if kar.log {
		fmt.Println("created", dst.Name())
	}

	return nil
}
