package redis

// 目的：使得用户登录时，判断用户请求携带的access token和目前存储库中的access token相同嘛？


import "fmt"

func SetUserToken(token string, ID int64) error {
	// err := rdb.Set("token_user" + fmt.Sprint(userID),token,0)
	// if err != nil {
	// 	return err.Err()
	// }
	// // 查看是否修改成功
	// fmt.Printf("rdb.Get(\"token_user\" + fmt.Sprint(userID)): %v\n", rdb.Get("token_user" + fmt.Sprint(userID)))
	// return nil
	return rdb.Set("token_user"+fmt.Sprint(ID), token, 0).Err()
}

func GetUserToken(ID int64) (string, error) {
	return rdb.Get("token_user" + fmt.Sprint(ID)).Result()
}

