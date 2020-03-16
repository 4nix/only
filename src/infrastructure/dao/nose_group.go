// Package dao nose_group 数据组表
package dao

type NoseGroup struct {
	ID     int64
	Name   string
	Config string
	Ctime  string
	Utime  string
	IsDel  int `db:"is_del"`
}

type noseGroupDao struct {
	TableName string
}

func NewNoseGroupDao() *noseGroupDao {
	return &noseGroupDao{TableName: "nose_group"}
}

func (nose *noseGroupDao) Add(sName string, sConfig string) (int64, error) {
	sql := `INSERT INTO ` + nose.TableName + ` (name, config) VALUES (?, ?)`
	res, _ := DB.Exec(sql, sName, sConfig)

	return res.LastInsertId()
}

func (nose *noseGroupDao) Find(iID int64) *NoseGroup {
	oNose := NoseGroup{}
	DB.Get(&oNose, "SELECT * FROM nose_group WHERE id = ?", iID)

	return &oNose
}

func (nose *noseGroupDao) Remove(iID int64) (int64, error) {
	sql := `DELETE FROM ` + nose.TableName + ` WHERE id = ?`
	res, _ := DB.Exec(sql, iID)

	return res.RowsAffected()
}

func (nose *noseGroupDao) Update(iID int64, sName string) (int64, error) {
	sql := `UPDATE ` + nose.TableName + ` SET name = ?`
	res, _ := DB.Exec(sql, sName)

	return res.RowsAffected()
}
