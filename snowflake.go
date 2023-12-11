package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

type Snowflake struct {
	sfid *snowflake.Node
}

func NewSnowflake(date string, id int64) (sfid Snowflake, err error) {
	var t time.Time
	if t, err = time.ParseInLocation("2006-01-02", date, location); err != nil {
		return
	}
	snowflake.Epoch = t.UnixNano() / 1000000
	sfid.sfid, err = snowflake.NewNode(id)
	return
}

func (s *Snowflake) GetID() int64 {
	return s.sfid.Generate().Int64()
}
