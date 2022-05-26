package kab

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

type ContainerFormat int

const (
	KindUnknown ContainerFormat = iota + 1
	Zip
	Gzip
	KindBZip2
	Tar
	Zstd
)

type ContentFormat int

const (
	FormatJSON ContentFormat = iota
	FormatXML
)

func Info(target string) (*Options, error) {
	head := make([]byte, 512)
	if _, err := os.Stat(target); err != nil {
		return nil, err
	}

	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := io.ReadFull(f, head[:]); err != nil {
		if err != io.ErrUnexpectedEOF {
			log.Println("cannot read first bytes", err)
			return nil, err
		}
	}

	if isTar := bytes.HasPrefix(head[256:], prefixTar); isTar {
		fmt.Printf("%s is a tar file", target)
		return &Options{Container: Tar}, err
	}

	if isGzip := bytes.HasPrefix(head[:], prefixGzip); isGzip {
		fmt.Printf("%s is a gzip file", target)
		return &Options{Container: Gzip}, err
	}

	if isZip := bytes.HasPrefix(head[:], prefixZip); isZip {
		fmt.Printf("%s is a zip file", target)
		return &Options{Container: Zip}, err
	}

	return nil, errors.New("unknown file format")
}

func DetectKind(target string) (ContainerFormat, error) {
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
		return Tar, nil
	}

	if isGzip := bytes.HasPrefix(head[:], prefixGzip); isGzip {
		return Gzip, nil
	}

	if isZip := bytes.HasPrefix(head[:], prefixZip); isZip {
		return Zip, nil
	}

	return KindUnknown, nil
}
