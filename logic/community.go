package logic

import (
	"bluebell/Dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查找到所以的community并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
