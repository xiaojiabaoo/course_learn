package req_model

type GetSubjectList struct {
	UserId int64  `json:"user_id"` // 用户ID，非客户端参数
	Type   int64  `json:"type"`    // 查询类型：1.主页菜单栏；2.所有课程；3.我的课程；4.热门课程推荐
	Action string `form:"action"`  // MANAGE_LOGIN.管理员查询；USER_LOGIN.用户查询
}

type TrySubject struct {
	SubjectId int64 `json:"subject_id"` // 课程ID
	UserId    int64 `json:"user_id"`    // 用户ID，非客户端参数
}

type GetSubjectDesc struct {
	SubjectId int64 `json:"subject_id"` // 课程ID
}
