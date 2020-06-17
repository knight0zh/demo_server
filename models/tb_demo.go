package models

import (
	"time"
)

type TbDemo struct {
	//[0] create_time                                    datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	CreateTime time.Time `json:"create_time"`
	//[1] last_update_time                               timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	LastUpdateTime time.Time `json:"last_update_time"`
}

func (t *TbDemo) TableName() string {
	return "db.tb_demo"
}
