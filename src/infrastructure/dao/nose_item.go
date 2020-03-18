// Package dao nose_item 数据组表
package dao

type NoseItem struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Config   string `json:"config"`
	ModeID   int    `db:"mode_id" json:"mode_id"`
	GroupID  int    `db:"group_id" json:"group_id"`
	RelateID int    `db:"relate_id" json:"relate_id"`
	Ctime    string `json:"ctime"`
	Utime    string `json:"utime"`
	IsDel    int    `db:"is_del" json:"is_del"`
}

func (nose *NoseItem) Add(sName string, sContent string, sConfig string, iModeID int, iGroupID int, iRelateID int) (int64, error) {
	sql := `INSERT INTO nose_item (name, content, config, mode_id, group_id, relate_id) VALUES (?, ?, ?, ?, ?, ?)`
	res, _ := DB.Exec(sql, sName, sContent, sConfig, iModeID, iGroupID, iRelateID)

	return res.LastInsertId()
}

func (nose *NoseItem) Find(iID int64) *NoseItem {
	oNose := NoseItem{}
	DB.Get(&oNose, "SELECT * FROM nose_item WHERE id = ?", iID)

	return &oNose
}

func (nose *NoseItem) FindByGroupID(iGID int64) *[]NoseItem {
	oNoses := []NoseItem{}
	DB.Select(&oNoses, "SELECT * FROM nose_item WHERE group_id = ?", iGID)

	return &oNoses
}

func (nose *NoseItem) FindByRelateID(iGroupID int64, iRelateID int64) *[]NoseItem {
	oNoses := []NoseItem{}
	DB.Select(&oNoses, "SELECT * FROM nose_item WHERE group_id = ? AND relate_id = ?", iGroupID, iRelateID)

	return &oNoses
}

func (nose *NoseItem) Remove(iID int64) (int64, error) {
	sql := `DELETE FROM nose_item WHERE id = ?`
	res, _ := DB.Exec(sql, iID)

	return res.RowsAffected()
}
