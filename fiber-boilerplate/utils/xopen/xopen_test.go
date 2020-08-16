package xopen

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type XopenTest struct{}

var _ = Suite(&XopenTest{})

func gzFromString(s string) string {
	var c bytes.Buffer
	gz := gzip.NewWriter(&c)
	gz.Write([]byte(s))
	return c.String()
}

var gzTests = []struct {
	isGz bool
	data string
}{
	{false, "asdf"},
	{true, gzFromString("asdf")},
}

func (s *XopenTest) TestIsGzip(c *C) {
	for _, t := range gzTests {
		isGz, err := IsGzip(bufio.NewReader(strings.NewReader(t.data)))
		c.Assert(err, IsNil)
		c.Assert(t.isGz, Equals, isGz)
	}
}

func (s *XopenTest) TestIsStdin(c *C) {
	r := IsStdin()
	c.Assert(r, Equals, false)
}

func (s *XopenTest) TestRopen(c *C) {
	rdr, err := Ropen("-")
	c.Assert(err, ErrorMatches, ".* stdin not detected")
	c.Assert(rdr, IsNil)
}

func (s *XopenTest) TestWopen(c *C) {
	for _, f := range []string{"t.gz", "t"} {
		testString := "ASDF1234"
		wtr, err := Wopen(f)
		c.Assert(err, IsNil)
		_, err = os.Stat(f)
		c.Assert(err, IsNil)
		c.Assert(wtr.wtr, NotNil)
		fmt.Fprintf(wtr, testString)
		wtr.Close()

		rdr, err := Ropen(f)
		c.Assert(err, IsNil)

		str, err := rdr.ReadString(99)
		c.Assert(str, Equals, testString)
		c.Assert(err, Equals, io.EOF)
		str, err = rdr.ReadString(99)
		c.Assert(str, Equals, "")

		rdr.Close()
		os.Remove(f)
	}
}

var httpTests = []struct {
	url         string
	expectError bool
}{
	{"https://raw.githubusercontent.com/brentp/xopen/master/README.md", false},
	{"http://raw.githubusercontent.com/brentp/xopen/master/README.md", false},
	{"http://raw.githubusercontent.com/brentp/xopen/master/BAD.md", true},
}

func (s *XopenTest) TestReadHttp(c *C) {
	for _, t := range httpTests {
		rdr, err := Ropen(t.url)
		if !t.expectError {
			c.Assert(err, IsNil)
			v, err := rdr.ReadString(byte('\n'))
			c.Assert(err, IsNil)
			c.Assert(len(v), Not(Equals), 0)
		} else {
			c.Assert(err, ErrorMatches, ".* 404 Not Found")
		}
	}
}

func (s *XopenTest) TestReadProcess(c *C) {
	for _, cmd := range []string{"|ls -lh", "|ls", "|ls -lh xopen_test.go"} {
		rdr, err := Ropen(cmd)
		c.Assert(err, IsNil)
		b := make([]byte, 1000)
		_, err = rdr.Read(b)
		if err != io.EOF {
			c.Assert(err, IsNil)
		}
		lines := strings.Split(string(b), "\n")
		has := false
		for _, line := range lines {
			if strings.Contains(line, "xopen_test.go") {
				has = true
			}
		}
		c.Assert(has, Equals, true)
	}
}

func (s *XopenTest) TestOpenStdout(c *C) {
	w, err := Wopen("-")
	c.Assert(err, IsNil)
	c.Assert(w.wtr, Equals, os.Stdout)
}

func (s *XopenTest) TestOpenBadFile(c *C) {
	r, err := Ropen("XXXXXXXXXXXXXXXXXXXXXXX")
	c.Assert(r, IsNil)
	c.Assert(err, ErrorMatches, ".* no such file .*")
}

func (s *XopenTest) TestWOpenBadFile(c *C) {
	w, err := Wopen("XX/XXX/XXX/XXX/XXX/XXXXXXXXX")
	c.Assert(w, IsNil)
	c.Assert(err, ErrorMatches, ".* no such file .*")
}

func (s *XopenTest) TestExists(c *C) {
	c.Assert(Exists("xopen.go"), Equals, true)
	c.Assert(Exists("____xx"), Equals, false)
}

func (s *XopenTest) TestUser(c *C) {
	c.Assert(Exists("~"), Equals, true)
}

func (s *XopenTest) TestExpand(c *C) {
	_, err := ExpandUser("~baduser66")
	c.Assert(err, Not(IsNil))
}
