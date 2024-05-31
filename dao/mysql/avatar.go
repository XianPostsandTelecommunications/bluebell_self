package mysql

// UploadAvatar 保存头像路径到数据库
func UploadAvatar(id string, fileName string) {
	sqlstr := `insert into user (user_id,avatar)values (?,?)`
	_, err := db.Exec(sqlstr, id, fileName)
	if err != nil {
		return
	}
}
