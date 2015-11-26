package temple

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestHumanize(t *testing.T) { TestingT(t) }

type HumanizeSuite struct{}

var _ = Suite(&HumanizeSuite{})

func (s *HumanizeSuite) TestBytes(c *C) {
	var str string
	var err error

	str, err = InBytes(123)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "123B")

	str, err = InBytes(204800)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "200KB")

	str, err = InBytes(1024 * 1024 * 1024 * 456)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "456GB")

	str, err = InBytes(1024*5 + 512)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "5.5KB")
}

func (s *HumanizeSuite) TestCount(c *C) {
	var str string
	var err error

	str, err = InCount(123)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "123")

	str, err = InCount(204800)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "205k")

	str, err = InCount(456700000000)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "457g")

	str, err = InCount(5500)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "5.5k")
}

func (s *HumanizeSuite) TestDuration(c *C) {
	var d time.Duration
	var str string
	var err error

	d, _ = time.ParseDuration("23s")
	str, err = Duration(d)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "23 secs")

	d, _ = time.ParseDuration("23m9s")
	str, err = Duration(d)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "23 mins")

	d, _ = time.ParseDuration("23h17m11s")
	str, err = Duration(d)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "23 hours")

	d, _ = time.ParseDuration("120h1m1s")
	str, err = Duration(d)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "5 days")

	d, _ = time.ParseDuration("2m30s")
	str, err = Duration(d)
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "2.5 mins")
}
