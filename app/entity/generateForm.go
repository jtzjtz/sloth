package entity

type GenerateForm struct {
	DBHost        string `json:"db_host" form:"db_host" `
	DBPort        string `json:"db_port" form:"db_port" `
	DBUserName    string `json:"db_user" form:"db_user" `
	Tabels        string `json:"tables" form:"tables" `
	DBPwd         string `json:"db_pwd" form:"db_pwd" `
	DB            string `json:"db" form:"db" binding:"required"`
	GitUrl        string `json:"git_url" form:"git_url" binding:"required"`
	GitPwd        string `json:"git_password" form:"git_password"`
	GitUser       string `json:"git_username" form:"git_username"`
	Tag           string `json:"version" form:"version" binding:"required"`
	GitMsg        string `json:"git_msg" form:"git_msg" binding:"required"`
	IsHttpService int    `json:"is_http_service" form:"is_http_service" `
}
