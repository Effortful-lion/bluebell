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

/*
 get返回单条数据并映射到单个结构体
 select返回多条数据并映射到切片
*/

// 根据id查询社区详情
func GetCommunityDetailByID(id int64)(communityDetail *models.CommunityDetail,err error){
    sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
    communityDetail = new(models.CommunityDetail)
    err = db.Get(communityDetail, sqlStr, id)
    if err != nil {
        if err == sql.ErrNoRows {
            zap.L().Warn("该社区不存在")
            err = ErrorInvalidID
            return
        }
        // 查询数据库失败
        zap.L().Error("查询数据库失败", zap.Error(err))
        return
    }
    return
}
