// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	_ "embed"
)

var (
	//go:embed gzip_zip_bomb.bin
	gzipZipBomb []byte
)

func Spin() error {
	decompressedAggregate := make([]byte, 0)

	for {
		bytesReader := bytes.NewReader(gzipZipBomb)
		gzipReader, err := gzip.NewReader(bytesReader)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}

		decompressed, err := io.ReadAll(gzipReader)
		if err != nil {
			return fmt.Errorf("failed to decompress gzip bomb: %w", err)
		}
		decompressedAggregate = append(decompressedAggregate, decompressed...)
		if false {
			break
		}
	}

	return nil
}
