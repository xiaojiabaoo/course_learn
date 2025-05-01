package resp_model

type GetUserLeanRecord struct{}

type GetUserData struct {
	NickName      string `json:"nick_name"`
	Gender        int64  `json:"gender"`
	Logo          string `json:"logo"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Area          string `json:"area"`
	AddTime       int64  `json:"add_time"`
	AccountStatus int64  `json:"account_status"`
}

type GetLoginRecordData struct {
	Count int64             `json:"count"`
	List  []LoginRecordData `json:"list"`
}

type LoginRecordData struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	LoginTime   int64  `json:"login_time"`
	LoginIp     string `json:"login_ip"`
	LoginArea   string `json:"login_area"`
	LoginMethod int64  `json:"login_method"`
	DateTime    string `json:"date_time"`
}

type GetOrderListData struct {
	Count int64           `json:"count"`
	List  []OrderListData `json:"list"`
}

type OrderListData struct {
	Id             int64   `json:"id"`
	OrderNo        string  `json:"order_no"`
	Type           int64   `json:"type"`
	IsPayment      int64   `json:"is_payment"`
	Amount         float64 `json:"amount"`
	BuyType        int64   `json:"buy_type"`
	IsValid        int64   `json:"is_valid"`
	ExpirationTime int64   `json:"expiration_time"`
	UserId         int64   `json:"user_id"`
	SubjectId      int64   `json:"subject_id"`
	SubjectName    string  `json:"subject_name"`
	AddTime        int64   `json:"add_time"`
}

type GetLearnRecordData struct {
	Count int64             `json:"count"`
	List  []LearnRecordData `json:"list"`
}

type LearnRecordData struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	PackageId   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
	SubjectId   int64  `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	ChapterId   int64  `json:"chapter_id"`
	ChapterName string `json:"chapter_name"`
	SectionId   int64  `json:"section_id"`
	SectionName string `json:"section_name"`
	PieceId     int64  `json:"piece_id"`
	PieceName   string `json:"piece_name"`
	AddTime     int64  `json:"add_time"`
}
