// Package xopen makes it easy to get buffered readers and writers.
// Ropen opens a (possibly gzipped) file/process/http site for buffered reading.
// Wopen opens a (possibly gzipped) file for buffered writing.
// Both will use gzip when appropriate and will user buffered IO.
package xopen

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	gzip "github.com/klauspost/pgzip"
	//"github.com/klauspost/compress/gzip"
	// "compress/gzip"
)

// ErrNoContent means nothing in the stream/file.
var ErrNoContent = errors.New("xopen: no content")

// ErrDirNotSupported means the path is a directory.
var ErrDirNotSupported = errors.New("xopen: input is a directory")

// IsGzip returns true buffered Reader has the gzip magic.
func IsGzip(b *bufio.Reader) (bool, error) {
	return CheckBytes(b, []byte{0x1f, 0x8b})
}

// IsStdin checks if we are getting data from stdin.
func IsStdin() bool {
	// http://stackoverflow.com/a/26567513
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// ExpandUser expands ~/path and ~otheruser/path appropriately
func ExpandUser(path string) (string, error) {
	if path[0] != '~' {
		return path, nil
	}
	var u *user.User
	var err error
	if len(path) == 1 || path[1] == '/' {
		u, err = user.Current()
	} else {
		name := strings.Split(path[1:], "/")[0]
		u, err = user.Lookup(name)
	}
	if err != nil {
		return "", err
	}
	home := u.HomeDir
	path = home + "/" + path[1:]
	return path, nil
}

// Exists checks if a local file exits
func Exists(path string) bool {
	path, perr := ExpandUser(path)
	if perr != nil {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

// CheckBytes peeks at a buffered stream and checks if the first read bytes match.
func CheckBytes(b *bufio.Reader, buf []byte) (bool, error) {

	m, err := b.Peek(len(buf))
	if err != nil {
		return false, ErrNoContent
	}
	for i := range buf {
		if m[i] != buf[i] {
			return false, nil
		}
	}
	return true, nil
}

// Reader is returned by Ropen
type Reader struct {
	*bufio.Reader
	rdr io.Reader
	gz  io.ReadCloser
}

// Close the associated files.
func (r *Reader) Close() error {
	if r.gz != nil {
		r.gz.Close()
	}
	if c, ok := r.rdr.(io.ReadCloser); ok {
		c.Close()
	}
	return nil
}

// Writer is returned by Wopen
type Writer struct {
	*bufio.Writer
	wtr *os.File
	gz  *gzip.Writer
}

// Close the associated files.
func (w *Writer) Close() error {
	w.Flush()
	if w.gz != nil {
		w.gz.Close()
	}
	w.wtr.Close()
	return nil
}

// Flush the writer.
func (w *Writer) Flush() {
	w.Writer.Flush()
	if w.gz != nil {
		w.gz.Flush()
	}
}

var pageSize = os.Getpagesize() * 2

// Buf returns a buffered reader from an io.Reader
// If f == "-", then it will attempt to read from os.Stdin.
// If the file is gzipped, it will be read as such.
func Buf(r io.Reader) (*Reader, error) {
	b := bufio.NewReaderSize(r, pageSize)
	var rdr io.ReadCloser
	if is, err := IsGzip(b); err != nil && err != io.EOF {
		return nil, err
	} else if is {
		// rdr, err = newFastGzReader(b)
		rdr, err = gzip.NewReader(b)
		if err != nil {
			return nil, err
		}
		b = bufio.NewReaderSize(rdr, pageSize)
	}

	// check BOM
	t, _, err := b.ReadRune()
	if err != nil {
		return nil, ErrNoContent
	}
	if t != '\uFEFF' {
		b.UnreadRune()
	}
	return &Reader{b, r, rdr}, nil
}

// XReader returns a reader from a url string or a file.
func XReader(f string) (io.Reader, error) {
	if strings.HasPrefix(f, "http://") || strings.HasPrefix(f, "https://") {
		var rsp *http.Response
		rsp, err := http.Get(f)
		if err != nil {
			return nil, err
		}
		if rsp.StatusCode != 200 {
			return nil, fmt.Errorf("http error downloading %s. status: %s", f, rsp.Status)
		}
		rdr := rsp.Body
		return rdr, nil
	}
	f, err := ExpandUser(f)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(f)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, ErrDirNotSupported
	}

	return os.Open(f)
}

// Ropen opens a buffered reader.
func Ropen(f string) (*Reader, error) {
	var err error
	var rdr io.Reader
	if f == "-" {
		if !IsStdin() {
			return nil, errors.New("stdin not detected")
		}
		b, err := Buf(os.Stdin)
		return b, err
	} else if f[0] == '|' {
		// TODO: use csv to handle quoted file names.
		cmdStrs := strings.Split(f[1:], " ")
		var cmd *exec.Cmd
		if len(cmdStrs) == 2 {
			cmd = exec.Command(cmdStrs[0], cmdStrs[1:]...)
		} else {
			cmd = exec.Command(cmdStrs[0])
		}
		rdr, err = cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		err = cmd.Start()
		if err != nil {
			return nil, err
		}
	} else {
		rdr, err = XReader(f)
	}
	if err != nil {
		return nil, err
	}
	b, err := Buf(rdr)
	return b, err
}

// Wopen opens a buffered reader.
// If f == "-", then stdout will be used.
// If f endswith ".gz", then the output will be gzipped.
func Wopen(f string) (*Writer, error) {
	var wtr *os.File
	if f == "-" {
		wtr = os.Stdout
	} else {
		dir := filepath.Dir(f)
		fi, err := os.Stat(dir)
		if err == nil && !fi.IsDir() {
			return nil, fmt.Errorf("can not write file into a non-directory path: %s", dir)
		}
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}

		wtr, err = os.Create(f)
		if err != nil {
			return nil, err
		}
	}
	if !strings.HasSuffix(f, ".gz") {
		return &Writer{bufio.NewWriterSize(wtr, pageSize), wtr, nil}, nil
	}
	gz := gzip.NewWriter(wtr)
	return &Writer{bufio.NewWriterSize(gz, pageSize), wtr, gz}, nil
}

// WopenGzip opens a buffered gzipped reader.
// If f == "-", then stdout will be used.
func WopenGzip(f string) (*Writer, error) {
	var wtr *os.File
	if f == "-" {
		wtr = os.Stdout
	} else {
		dir := filepath.Dir(f)
		fi, err := os.Stat(dir)
		if err == nil && !fi.IsDir() {
			return nil, fmt.Errorf("can not write file into a non-directory path: %s", dir)
		}
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
		wtr, err = os.Create(f)
		if err != nil {
			return nil, err
		}
	}
	gz := gzip.NewWriter(wtr)
	return &Writer{bufio.NewWriterSize(gz, pageSize), wtr, gz}, nil
}

// WopenFile opens a buffered reader.
// If f == "-", then stdout will be used.
// If f endswith ".gz", then the output will be gzipped.
func WopenFile(f string, flag int, perm os.FileMode) (*Writer, error) {
	var wtr *os.File
	if f == "-" {
		wtr = os.Stdout
	} else {
		dir := filepath.Dir(f)
		fi, err := os.Stat(dir)
		if err == nil && !fi.IsDir() {
			return nil, fmt.Errorf("can not write file into a non-directory path: %s", dir)
		}
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
		wtr, err = os.OpenFile(f, flag, perm)
		if err != nil {
			return nil, err
		}
	}
	if !strings.HasSuffix(f, ".gz") {
		return &Writer{bufio.NewWriterSize(wtr, pageSize), wtr, nil}, nil
	}
	gz := gzip.NewWriter(wtr)
	return &Writer{bufio.NewWriterSize(gz, pageSize), wtr, gz}, nil
}
