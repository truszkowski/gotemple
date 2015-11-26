package temple

import (
	"bytes"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestHtml(t *testing.T) { TestingT(t) }

type HtmlSuite struct{}

var _ = Suite(&HtmlSuite{})

func (s *HtmlSuite) TestExterns(c *C) {
	t := NewHtml("test")

	t, err := t.Parse(`{{ define "a"}}Extends: {{extends .Name .Data}}{{ end }}{{define "b"}}this is B, {{ . }}{{end}}{{define "c"}}this is C, {{ . }}{{end}}`)
	c.Assert(err, IsNil)

	var buf *bytes.Buffer

	buf = bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "a", struct{ Name, Data string }{"b", "100"})
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "Extends: this is B, 100")

	buf = bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "a", struct{ Name, Data string }{"c", "xyz"})
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "Extends: this is C, xyz")

	buf = bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "a", struct{ Name, Data string }{"d", "aaa"})
	c.Assert(err, Not(IsNil))
}

func (s *HtmlSuite) TestBytes(c *C) {
	t := NewHtml("test")

	t, err := t.Parse(`{{ define "a"}}{{.|inBytes}},{{.|inKBytes}},{{.|inMBytes}}{{end}}`)
	c.Assert(err, IsNil)

	var buf *bytes.Buffer

	buf = bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "a", 23*1024*1024)
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "23MB,23GB,23TB")
}

func (s *HtmlSuite) TestElapsed(c *C) {
	t := NewHtml("test")
	now := time.Now()

	t, err := t.Parse(`{{ define "a"}}{{.|elapsed}}{{end}}`)
	c.Assert(err, IsNil)

	var buf *bytes.Buffer

	buf = bytes.NewBuffer([]byte{})
	err = t.ExecuteTemplate(buf, "a", now.Add(-31*time.Minute))
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "31 mins ago")
}
