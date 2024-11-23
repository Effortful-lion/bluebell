package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	// 查询数据库表 community 查询所有数据(id、名称)
	// data, err = mysql.GetCommunityList()
	// if err != nil {
	// 	return nil, err
	// }
	// return data, nil
	return mysql.GetCommunityList()
}