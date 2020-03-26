// Package dao nose_group 数据组表
package dao

import "fmt"

type NoseGroup struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Config string `json:"config"`
	Ctime  string `json:"ctime"`
	Utime  string `json:"utime"`
	IsDel  int    `db:"is_del" json:"is_del"`
}

func (nose *NoseGroup) Add(sName string, sConfig string) (int64, error) {
	sql := `INSERT INTO nose_group (name, config) VALUES (?, ?)`
	res, err := DB.Exec(sql, sName, sConfig)
	fmt.Println(err)

	return res.LastInsertId()
}

func (nose *NoseGroup) Find(iID int64) *NoseGroup {
	oNose := NoseGroup{}
	DB.Get(&oNose, "SELECT * FROM nose_group WHERE id = ?", iID)

	return &oNose
}

func (nose *NoseGroup) FindAll() *[]NoseGroup {
	oNose := []NoseGroup{}
	err := DB.Select(&oNose, "SELECT * FROM nose_group")
	fmt.Println(err)

	return &oNose
}

func (nose *NoseGroup) Remove(iID int64) (int64, error) {
	sql := `DELETE FROM nose_group WHERE id = ?`
	res, _ := DB.Exec(sql, iID)

	return res.RowsAffected()
}

func (nose *NoseGroup) Update(iID int64, sName string) (int64, error) {
	sql := `UPDATE nose_group SET name = ?`
	res, _ := DB.Exec(sql, sName)

	return res.RowsAffected()
}
