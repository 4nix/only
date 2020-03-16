package entity

type GroupEntity struct {
	info interface{}
}

// func NewGroup(iGID int64) *group {
// 	oGroup := repository.FindGroup(iGID)

// 	return &group{info: oGroup}
// }

// func (entity *GroupEntity) GetInfo() interface{} {
// 	return GroupEntity
// }

// func (entity *GroupEntity) AddItem(sName string, sContent string, sConfig string, iModeID int, iRelateID int) bool {
// 	repository.AddItem(sName, sContent, sConfig, iModeID, entity.info.ID, iRelateID)
// 	return true
// }
