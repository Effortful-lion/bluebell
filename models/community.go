package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64  `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"` // omitempty 是说如果这个字段为空，那么就不返回这个字段
	CreateTime   time.Time `json:"create_time" db:"create_time"`	// 这里如果要使用 time 类型的时间，注意连接库的时候，需要设置参数 parseTime=true
	// 或者这里直接传回一个时间戳，然后前端自己转换（一般）
}