package gopcua

/*
#cgo CFLAGS: -std=c99 -I .
#include <stdlib.h>
#include <gopcua_client.h>
*/
import "C"

import "time"

type Client struct {
	c_client *C.UA_Client
}

type Point struct {
	ServerTs    time.Time
	SourceTs    time.Time
	HasValue    bool
	GoodValue   bool
	HasSourceTs bool
	HasServerTs bool
}

type Point_Int32 struct {
	Value int32
	Point
}

func (p C.gopcua_point_int32) toPoint() Point_Int32 {
	return Point_Int32{
		Value: int32(p.value),
		Point: Point{
			ServerTs:    p.server_ts.toTime(),
			SourceTs:    p.source_ts.toTime(),
			HasValue:    bool(p.has_value),
			GoodValue:   bool(p.good_value),
			HasSourceTs: bool(p.has_source_ts),
			HasServerTs: bool(p.has_server_ts),
		},
	}
}

type Point_Float32 struct {
	Value float32
	Point
}

func (p C.gopcua_point_float32) toPoint() Point_Float32 {
	return Point_Float32{
		Value: float32(p.value),
		Point: Point{
			ServerTs:    p.server_ts.toTime(),
			SourceTs:    p.source_ts.toTime(),
			HasValue:    bool(p.has_value),
			GoodValue:   bool(p.good_value),
			HasSourceTs: bool(p.has_source_ts),
			HasServerTs: bool(p.has_server_ts),
		},
	}
}

func (d C.struct_UA_DateTimeStruct) toTime() time.Time {
	return time.Date(
		int(d.year),
		time.Month(int(d.month)),
		int(d.day),
		int(d.hour),
		int(d.min),
		int(d.sec),
		int(d.milliSec)*1000000+int(d.microSec)*1000+int(d.nanoSec),
		time.UTC,
	)
}
