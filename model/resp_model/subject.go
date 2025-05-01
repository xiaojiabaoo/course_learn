package resp_model

type GetSubjectList struct {
	KnowledgeTreeId int64     `json:"knowledge_tree_id"`
	SubjectId       int64     `json:"subject_id"`
	SubjectName     string  `json:"subject_name"`
	PackageId       int64     `json:"package_id"`
	Sort            int64     `json:"sort"`
	Status          int64     `json:"status"` // 拥有状态：1.试用；2.已购
	Desc            string  `json:"desc"`
	Progress        float64 `json:"progress"` // 学习进度，Status等于1和2时才有值
	Expire          bool    `json:"expire"`   // 是否即将过期，Status等于1和2时才有值
}

type GetSubjectDetail struct {
	KnowledgeTreeId int64    `json:"knowledge_tree_id"`
	SubjectId       int64    `json:"subject_id"`
	SubjectName     string `json:"subject_name"`
	PackageId       int64    `json:"package_id"`
	Sort            int64    `json:"sort"`
	Desc            string `json:"desc"`
}
