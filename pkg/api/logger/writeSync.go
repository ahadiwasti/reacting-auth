package logger

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/kjk/dailyrotate"
)

const compressSuffix = ".gz"

type dailyFile struct {
	*dailyrotate.File
}

func (d *dailyFile) Sync() error {
	return d.Flush()
}

// compressLogFile compresses the given log file, removing the
// uncompressed log file if successful.
func compressLogFile(src, dst string) (err error) {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer f.Close()

	fi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat log file: %v", err)
	}

	if err := chown(dst, fi); err != nil {
		return fmt.Errorf("failed to chown compressed log file: %v", err)
	}

	// If this file already exists, we presume it was created by
	// a previous attempt to compress the log file.
	gzf, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
	if err != nil {
		return fmt.Errorf("failed to open compressed log file: %v", err)
	}
	defer gzf.Close()

	gz := gzip.NewWriter(gzf)

	defer func() {
		if err != nil {
			os.Remove(dst)
			err = fmt.Errorf("failed to compress log file: %v", err)
		}
	}()

	if _, err := io.Copy(gz, f); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	if err := gzf.Close(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}

func onClose(compress bool, maxBackups int) func(string, bool) {
	return func(path string, didRotate bool) {
		if didRotate {
			// compress
			fFile := path
			if compress {
				if err := compressLogFile(path, path+compressSuffix); err != nil {
					Error(err)
					return
				}
				fFile += compressSuffix
			}

			// MaxBackups is the maximum number of old log files to retain.  The default
			// is to retain all old log files
			if maxBackups > 0 {
				// todo
			}
		}
	}
}
