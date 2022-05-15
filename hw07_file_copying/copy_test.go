package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy_Positive(t *testing.T) {
	tests := []struct {
		name     string
		fromPath string
		toPath   string
		offset   int64
		limit    int64
		wantErr  error
	}{
		{
			name:     "success copy no limit and no offset",
			fromPath: "testdata/out_offset0_limit0.txt",
			toPath:   "testdata/out_success_copy_test",
		},
		{
			name:     "success copy with limit and no offset",
			fromPath: "testdata/out_offset0_limit0.txt",
			toPath:   "testdata/out_success_copy_test",
			limit:    int64(10),
		},
		{
			name:     "success copy with limit and offset",
			fromPath: "testdata/input.txt",
			toPath:   "testdata/out_success_copy_test",
			offset:   int64(5),
			limit:    int64(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.fromPath, tt.toPath, tt.offset, tt.limit)
			require.NoError(t, err)

			err = os.Remove(tt.toPath)
			if err != nil {
				t.Errorf("failed test with error: %v", err)
			}
		})
	}
}

func TestCopy_Negative(t *testing.T) {
	tests := []struct {
		name     string
		fromPath string
		toPath   string
		offset   int64
		limit    int64
		wantErr  error
	}{
		{
			name:     "file_not_exist",
			fromPath: "",
			wantErr:  ErrFileNotExist,
		},
		{
			name:     "unsupported file",
			fromPath: "testdata/out_offset0_limit10.txt",
			wantErr:  ErrOffsetExceedsFileSize,
			offset:   int64(50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.fromPath, tt.toPath, tt.offset, tt.limit)
			require.EqualErrorf(t, tt.wantErr, err.Error(), err.Error())
		})
	}
}
