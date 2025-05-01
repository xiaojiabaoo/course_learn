package mysql_model

type TPackage struct {
	PackageId   int64  `gorm:"not null pk autoincr INT"`     // 专业ID
	PackageName string `gorm:"comment('专业名称') VARCHAR(255)"` // 专业名称
}
