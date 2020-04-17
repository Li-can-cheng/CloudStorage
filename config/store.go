package config

import (
	cmn "cloudstorage-server/common"
)

const (
	// CurrentStoreType : 设置当前文件的存储类型
	// CurrentStoreType = cmn.StoreLocal
	CurrentStoreType = cmn.StoreCeph
)
