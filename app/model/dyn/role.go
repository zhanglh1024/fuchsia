package dyn


type RoleInfoDyn struct{
	RoleId   string
	UserName string		`json:"user_name"`
	Sex		 int		`json:"sex"`
	Password  string	`json:"password"`
} 
