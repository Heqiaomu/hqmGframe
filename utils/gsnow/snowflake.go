package gsnow

import (
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/Heqiaomu/hqmGframe/utils/gconsts"
	"sync"
	"time"
)

type snowflakeInter interface {
	GetId() int64
}

var once sync.Once

func New() snowflakeInter {
	once.Do(func() {
		sF = &snowflake{
			timestamp: 0,
			machineId: config.GetConfig().SnowFlake.SnowFlakeMachineID,
			sequence:  0,
		}
	})
	return sF
}

var sF snowflakeInter

func GetSnowFlake() snowflakeInter {
	return sF
}

type snowflake struct {
	sync.Mutex
	timestamp int64
	machineId int64
	sequence  int64
}

// GetId 生成分布式ID
func (s *snowflake) GetId() int64 {
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & gconsts.SequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	r := (now-gconsts.StartTimeStamp)<<gconsts.TimestampShift | (s.machineId << gconsts.MachineIdShift) | (s.sequence)
	return r
}
