package lark

// UserIDType 用户ID类型
type UserIDType string

const (
	UserIDTypeOpenID  UserIDType = "open_id"
	UserIDTypeUnionID UserIDType = "union_id"
	UserIDTypeUserID  UserIDType = "user_id"
)

// DepartmentIDType 用户ID类型
type DepartmentIDType string

const (
	DepartmentIDTypeDepartment     DepartmentIDType = "department_id"
	DepartmentIDTypeOpenDepartment DepartmentIDType = "open_department_id"
)

// UserRequest 单个用户查询参数
type UserRequest struct {
	UserID           string           `json:"user_id"`
	UserIDType       UserIDType       `json:"user_id_type"`
	DepartmentIDType DepartmentIDType `json:"department_id_type"`
}

// User 单个用户信息
type User struct {
	UnionID       string `json:"union_id"`
	UserID        string `json:"user_id"`
	OpenID        string `json:"open_id"`
	Name          string `json:"name"`
	EnName        string `json:"en_name"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	MobileVisible bool   `json:"mobile_visible"`
	Gender        int    `json:"gender"`
	Avatar        struct {
		Avatar72     string `json:"avatar_72"`
		Avatar240    string `json:"avatar_240"`
		Avatar640    string `json:"avatar_640"`
		AvatarOrigin string `json:"avatar_origin"`
	} `json:"avatar"`
	Status struct {
		IsFrozen    bool `json:"is_frozen"`
		IsResigned  bool `json:"is_resigned"`
		IsActivated bool `json:"is_activated"`
	} `json:"status"`
	DepartmentIds   []string `json:"department_ids"`
	LeaderUserID    string   `json:"leader_user_id"`
	City            string   `json:"city"`
	Country         string   `json:"country"`
	WorkStation     string   `json:"work_station"`
	JoinTime        int64    `json:"join_time"`
	IsTenantManager bool     `json:"is_tenant_manager"`
	EmployeeNo      string   `json:"employee_no"`
	EmployeeType    int      `json:"employee_type"`
	Orders          []struct {
		DepartmentID    string `json:"department_id"`
		UserOrder       int    `json:"user_order"`
		DepartmentOrder int    `json:"department_order"`
	} `json:"orders"`
	CustomAttrs []struct {
		Type  string `json:"type"`
		ID    string `json:"id"`
		Value struct {
			Text        string `json:"text"`
			URL         string `json:"url"`
			PcURL       string `json:"pc_url"`
			OptionValue string `json:"option_value"`
			Name        string `json:"name"`
			PictureURL  string `json:"picture_url"`
			GenericUser struct {
				ID   string `json:"id"`
				Type int    `json:"type"`
			} `json:"generic_user"`
		} `json:"value"`
	} `json:"custom_attrs"`
	EnterpriseEmail string `json:"enterprise_email"`
	JobTitle        string `json:"job_title"`
}

// UserResponse 但用户查询返回结果
type UserResponse struct {
	Response
	Data struct {
		User `json:"user"`
	} `json:"data"`
}
