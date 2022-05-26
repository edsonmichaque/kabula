package archive

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	prefixZip  = []byte("PK\x03\x04")
	prefixGzip = []byte("\x1F\x8B\x08")
	prefixTar  = []byte("\x00\x75\x73\x74\x61\x72")
)

type ArchiveFormat int

const (
	KindUnknown ArchiveFormat = iota + 1
	KindZip
	KindGZip
	KindBZip2
	KindTar
)

type ContentFormat int

const (
	FormatJSON ContentFormat = iota
	FormatXML
)

func Info(target string) error {
	head := make([]byte, 512)
	if _, err := os.Stat(target); err != nil {
		return err
	}

	f, err := os.Open(target)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := io.ReadFull(f, head[:]); err != nil {
		if err != io.ErrUnexpectedEOF {
			log.Println("cannot read first bytes", err)
			return nil
		}
	}

	if isTar := bytes.HasPrefix(head[256:], prefixTar); isTar {
		fmt.Printf("%s is a tar file", target)
		return nil
	}

	if isGzip := bytes.HasPrefix(head[:], prefixGzip); isGzip {
		fmt.Printf("%s is a gzip file", target)
		return nil
	}

	if isZip := bytes.HasPrefix(head[:], prefixZip); isZip {
		fmt.Printf("%s is a zip file", target)
		return nil
	}

	return errors.New("unknown file format")
}

func DetectKind(target string) (ArchiveFormat, error) {
	head := make([]byte, 512)
	if _, err := os.Stat(target); err != nil {
		return KindUnknown, err
	}

	f, err := os.Open(target)
	if err != nil {
		return KindUnknown, err
	}

	defer f.Close()

	if _, err := io.ReadFull(f, head[:]); err != nil {
		if err != io.ErrUnexpectedEOF {
			return KindUnknown, err
		}
	}

	if isTar := bytes.HasPrefix(head[256:], prefixTar); isTar {
		return KindTar, nil
	}

	if isGzip := bytes.HasPrefix(head[:], prefixGzip); isGzip {
		return KindGZip, nil
	}

	if isZip := bytes.HasPrefix(head[:], prefixZip); isZip {
		return KindZip, nil
	}

	return KindUnknown, nil
}
