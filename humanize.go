package tmpl

import (
	"fmt"
	"reflect"
	"time"
)

type unitStruct struct {
	Value float64
	More  float64
	Unit  string
}

var bytesUnits []unitStruct = []unitStruct{
	unitStruct{1024, 0, "B"},
	unitStruct{1024, 10, "KB"},
	unitStruct{1024, 10, "MB"},
	unitStruct{1024, 10, "GB"},
	unitStruct{1024, 10, "TB"},
	unitStruct{0, 10, "PB"},
}

var countUnits []unitStruct = []unitStruct{
	unitStruct{1000, 0, ""},
	unitStruct{1000, 10, "k"},
	unitStruct{1000, 10, "m"},
	unitStruct{1000, 10, "g"},
	unitStruct{1000, 10, "t"},
	unitStruct{0, 10, "p"},
}

var timeUnits []unitStruct = []unitStruct{
	unitStruct{1000, 0, " nsecs"},
	unitStruct{1000, 10, " usecs"},
	unitStruct{1000, 10, " msecs"},
	unitStruct{60, 5, " secs"},
	unitStruct{60, 5, " mins"},
	unitStruct{24, 4, " hours"},
	unitStruct{0, 3, " days"},
}

var timeLayout = "2006-01-02-15:04:05"
var floatType = reflect.TypeOf(float64(0))

func toFloat64(v interface{}) (float64, error) {
	value := reflect.ValueOf(v)
	value = reflect.Indirect(value)
	if !value.Type().ConvertibleTo(floatType) {
		return 0, fmt.Errorf("cannot convert %v to float64", value.Type())
	}

	floatValue := value.Convert(floatType)
	return floatValue.Float(), nil
}

func inUnit(v interface{}, factor float64, units []unitStruct, suffix string) (string, error) {
	x, err := toFloat64(v)
	if err != nil {
		return "-", err
	}
	x *= factor

	for _, v := range units {
		if v.Value == 0 || x < v.Value {
			if x <= v.More {
				return fmt.Sprintf("%.1f%s%s", x, v.Unit, suffix), nil
			} else {
				return fmt.Sprintf("%.0f%s%s", x, v.Unit, suffix), nil
			}
		}
		x /= v.Value
	}

	panic("missing zero value in units slice")
}

func InBytes(bytes interface{}) (string, error) {
	return inUnit(bytes, 1, bytesUnits, "")
}

func InKBytes(kbytes interface{}) (string, error) {
	return inUnit(kbytes, 1024, bytesUnits, "")
}

func InMBytes(mbytes interface{}) (string, error) {
	return inUnit(mbytes, 1024*1024, bytesUnits, "")
}

func InGBytes(gbytes interface{}) (string, error) {
	return inUnit(gbytes, 1024*1024*1024, bytesUnits, "")
}

func InTBytes(tbytes interface{}) (string, error) {
	return inUnit(tbytes, 1024*1024*1024*1024, bytesUnits, "")
}

func InPBytes(pbytes interface{}) (string, error) {
	return inUnit(pbytes, 1024*1024*1024*1024*1024, bytesUnits, "")
}

func InCount(count interface{}) (string, error) {
	return inUnit(count, 1, countUnits, "")
}

func InKCount(kcount interface{}) (string, error) {
	return inUnit(kcount, 1.0e+3, countUnits, "")
}

func InMCount(mcount interface{}) (string, error) {
	return inUnit(mcount, 1.0e+6, countUnits, "")
}
func InGCount(gcount interface{}) (string, error) {
	return inUnit(gcount, 1.0e+9, countUnits, "")
}
func InTCount(tcount interface{}) (string, error) {
	return inUnit(tcount, 1.0e+12, countUnits, "")
}
func InPCount(pcount interface{}) (string, error) {
	return inUnit(pcount, 1.0e+15, countUnits, "")
}

func Datetime(t time.Time) (string, error) {
	return t.Format(timeLayout), nil
}

func Elapsed(t time.Time) (string, error) {
	return inUnit(time.Now().Sub(t), 1, timeUnits, " ago")
}

func Duration(t time.Duration) (string, error) {
	return inUnit(t, 1, timeUnits, "")
}

func InNanoSeconds(t interface{}) (string, error) {
	return inUnit(t, 1, timeUnits, "")
}

func InMicroSeconds(t interface{}) (string, error) {
	return inUnit(t, 1.0e+3, timeUnits, "")
}

func InMiliSeconds(t interface{}) (string, error) {
	return inUnit(t, 1.0e+6, timeUnits, "")
}

func InSeconds(t interface{}) (string, error) {
	return inUnit(t, 1.0e+9, timeUnits, "")
}

func InMinutes(t interface{}) (string, error) {
	return inUnit(t, 6.0e+10, timeUnits, "")
}

func InHours(t interface{}) (string, error) {
	return inUnit(t, 3.6e+12, timeUnits, "")
}

func InDays(t interface{}) (string, error) {
	return inUnit(t, 8.64e+13, timeUnits, "")
}
