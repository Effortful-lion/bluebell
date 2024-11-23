package mysql

import (
    "bluebell/models"
    "database/sql"
    "go.uber.org/zap"
)

// 获取社区列表
func GetCommunityList() (communityList []*models.Community, err error) {
    sqlStr := "select community_id, community_name from community"
    err = db.Select(&communityList, sqlStr)
    if err != nil {
        if err == sql.ErrNoRows {
            // 没有查询到记录
            zap.L().Warn("没有社区列表信息")
            err = nil
            return
        }
		// 查询数据库失败（非数据的其他原因）
        return
    }
    return
}