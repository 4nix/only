package admin

import "only/src/domain/service"

func AddGroup(sName string, mConfig string) int64 {
	return service.AddGroup(sName, mConfig)
}
