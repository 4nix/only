// Package dao nose_item 数据组表
package dao

type NoseItem struct {
	ID       int64
	Name     string
	Content  string
	Config   string
	ModeID   int `db:"mode_id"`
	GroupID  int `db:"group_id"`
	RelateID int `db:"relate_id"`
	Ctime    string
	Utime    string
	IsDel    int `db:"is_del"`
}

type noseItemDao struct {
	TableName string
}

func NewNoseItemDao() *noseItemDao {
	return &noseItemDao{TableName: "nose_item"}
}

func (nose *noseItemDao) Add(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) (int64, error) {
	sql := `INSERT INTO ` + nose.TableName + ` (name, content, config, mode_id, group_id, relate_id) VALUES (?, ?, ?, ?, ?, ?)`
	res, _ := DB.Exec(sql, sName, sContent, sConfig, iModeID, iGroupID, iRelateID)

	return res.LastInsertId()
}

func (nose *noseItemDao) Find(iID int64) *NoseItem {
	oNose := NoseItem{}
	DB.Get(&oNose, "SELECT * FROM "+nose.TableName+" WHERE id = ?", iID)

	return &oNose
}

func (nose *noseItemDao) FindByGID(iGID int64) *[]NoseItem {
	oNoses := []NoseItem{}
	DB.Get(&oNoses, "SELECT * FROM "+nose.TableName+" WHERE group_id = ?", iGID)

	return &oNoses
}

func (nose *noseItemDao) Remove(iID int64) (int64, error) {
	sql := `DELETE FROM ` + nose.TableName + ` WHERE id = ?`
	res, _ := DB.Exec(sql, iID)

	return res.RowsAffected()
}
