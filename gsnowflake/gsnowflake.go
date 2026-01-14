package gsnowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/lsq51201314/go-utils/gsha"
	"github.com/lsq51201314/go-utils/gtime"
)

type Snowflake struct {
	sid *snowflake.Node
}

func New(date string, id int64) (sf *Snowflake, err error) {
	var t time.Time
	if t, err = time.ParseInLocation("2006-01-02", date, gtime.Location); err != nil {
		return nil, err
	}
	snowflake.Epoch = t.UnixNano() / 1000000
	var obj Snowflake
	if obj.sid, err = snowflake.NewNode(id); err != nil {
		return nil, err
	}
	return &obj, nil
}

func (sf *Snowflake) Get() int64 {
	return sf.sid.Generate().Int64()
}

func (sf *Snowflake) GetSha1() string {
	return gsha.Sha1Str([]byte(sf.sid.Generate().String()))
}
