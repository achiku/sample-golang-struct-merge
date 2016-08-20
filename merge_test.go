package mergestruct

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/AdrianLungu/decimal"
	"github.com/achiku/mergo"
)

const timePkgPath = "time.Time"
const decimalPkgPath = "github.com/achiku/sample-golang-struct-merge/vendor/github.com/AdrianLungu/decimal.Decimal"

func getStructPath(v reflect.Value) string {
	return fmt.Sprintf("%s.%s", v.Type().PkgPath(), v.Type().Name())
}

var zeroTime = time.Time{}
var zeroDecimal = decimal.NewFromFloat(0)

var mergeFunc = func(dst, src reflect.Value) {
	switch {
	case getStructPath(dst) == timePkgPath:
		dt, _ := dst.Interface().(time.Time)
		st, _ := src.Interface().(time.Time)
		if dt == zeroTime || st != zeroTime {
			dst.Set(src)
		}
	case getStructPath(dst) == decimalPkgPath:
		dd, _ := dst.Interface().(decimal.Decimal)
		sd, _ := src.Interface().(decimal.Decimal)
		if dd.Cmp(zeroDecimal) == 0 || sd.Cmp(zeroDecimal) != 0 {
			dst.Set(src)
		}
	}
}

func TestMerge(t *testing.T) {
	pNum := decimal.NewFromFloat(100.112)
	e := &Event{
		ID:        int64(rand.Intn(10000)),
		Name:      "critical event",
		Number:    decimal.NewFromFloat(100.01123),
		PNumber:   &pNum,
		EmittedAt: time.Now().AddDate(1, 0, 0),
	}

	defaultEvent := &Event{
		ID:     10,
		Name:   "default event",
		Number: decimal.NewFromFloat(2),
		Detail: EventDetail{
			UserID:      1,
			Description: "test desc",
			Tags:        []string{"a", "b"},
		},
		EmittedAt: time.Now(),
	}

	// MergeWithOverwrite will do the same as Merge except that non-empty
	// dst attributes will be overriden by non-empty src attribute values.
	// mergo.MergeWithOverwrite(dst, src)
	if err := mergo.MergeWithOverwrite(defaultEvent, e, mergeFunc); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", defaultEvent)
}

func TestFunc(t *testing.T) {
	double := func(i int) int {
		return i * 2
	}
	cases := []struct {
		f func(int) int
		x int
		y int
	}{
		{x: 10, f: double, y: 20},
		{x: 10, f: func(i int) int { return i * 2 }, y: 20},
		{x: 10, f: nil, y: 20},
	}

	for _, c := range cases {
		a := c.f(c.x)
		if a != c.y {
			t.Errorf("want %d got %d", c.y, a)
		}

		b := f(c.f, c.x)
		if a != c.y {
			t.Errorf("want %d got %d", c.y, b)
		}
	}
}
