package tox

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Time(s interface{}) time.Time {
	vx := TimePtr(s)
	if vx != nil {
		return *vx
	}
	return time.Time{}
}

func TimeSAP(s string) *time.Time {

	var yyyy, mm, dd int

	if len(s) == 8 {
		// yyyymmdd
		runes := []rune(s)
		yyyy, _ = strconv.Atoi(string(runes[0:4]))
		mm, _ = strconv.Atoi(string(runes[4:6]))
		dd, _ = strconv.Atoi(string(runes[6:8]))
	} else if len(s) == 10 {
		// yyyy-mm-dd
		sx := strings.Split(s, `-`)
		if len(sx) == 3 {
			yyyy, _ = strconv.Atoi(sx[0])
			mm, _ = strconv.Atoi(sx[1])
			dd, _ = strconv.Atoi(sx[2])
		}
	}
	if yyyy > 2000 {
		t := time.Date(yyyy, time.Month(mm), dd, 0, 0, 0, 0, time.Local)
		return &t
	}

	return nil
}

func TimeSAPT(s, t string) *time.Time {

	var yyyy, mm, dd int

	if len(s) == 8 {
		// yyyymmdd
		runes := []rune(s)
		yyyy, _ = strconv.Atoi(string(runes[0:4]))
		mm, _ = strconv.Atoi(string(runes[4:6]))
		dd, _ = strconv.Atoi(string(runes[6:8]))
	} else if len(s) == 10 {
		// yyyy-mm-dd
		sx := strings.Split(s, `-`)
		if len(sx) == 3 {
			yyyy, _ = strconv.Atoi(sx[0])
			mm, _ = strconv.Atoi(sx[1])
			dd, _ = strconv.Atoi(sx[2])
		}
	}

	if yyyy > 2000 {

		var hh, mi, ss int
		if len(t) == 6 {
			// hhmiss
			runes := []rune(s)
			hh, _ = strconv.Atoi(string(runes[0:2]))
			mi, _ = strconv.Atoi(string(runes[2:4]))
			ss, _ = strconv.Atoi(string(runes[4:6]))
		} else if len(t) == 8 {
			// hh:mi:ss
			sx := strings.Split(t, `:`)
			if len(sx) == 3 {
				hh, _ = strconv.Atoi(sx[0])
				mi, _ = strconv.Atoi(sx[1])
				ss, _ = strconv.Atoi(sx[2])
			}
		}

		t := time.Date(yyyy, time.Month(mm), dd, hh, mi, ss, 0, time.Local)
		return &t
	}

	return nil
}

func TimePtr(s interface{}) *time.Time {

	// nil ?
	if s == nil {
		return nil
	}

	// pointer ?
	if reflect.TypeOf(s).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(s)
		if rv.IsNil() {
			return nil
		}
		s = reflect.Indirect(rv).Interface()
	}

	// time
	if vs, ok := s.(time.Time); ok {
		return &vs
	}

	// string
	if vs, ok := s.(string); ok {
		vx, ex := time.Parse(time.RFC3339, vs)
		if ex != nil {
			return nil
		}
		return &vx
	}

	return nil

}
