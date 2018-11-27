package app

//属性类型
type AttrType int16

const (
	WISDOM   AttrType = 1 //智慧
	LOYAL    AttrType = 2 //忠诚
	HEROIC   AttrType = 3 //英勇
	DILIGENT AttrType = 4 //勤勉
)

type ResourceType int16

const (
	FOOD    ResourceType = 1 //食物
	SOLDIER ResourceType = 2 //士兵
	GOLD    ResourceType = 3 //金币
)

type CodeErrorType  int

const (
	ParseRequestError    		 CodeErrorType = 1 //解析请求数据出错
	SearchDataError		 	     CodeErrorType = 2 //查找数据出错
	UpdateDataError		 		 CodeErrorType = 3 //更新数据出错
	InsertDataError    	  		 CodeErrorType = 4 //插入数据出错
	RoleIdAlreadyExists  	     CodeErrorType = 5 //该账号已经村子
	RoleNameAlreadyExists   	 CodeErrorType = 6 //该用户名已经存在
	RoleIdOrPassWordNil      	 CodeErrorType = 7 //用户名或密码为空
	RoleInfoIsNotExit        	 CodeErrorType = 8 //用户不存在
	RoleIdOrPassWordError	 	 CodeErrorType = 9 //用户名或密码错误
	IntervalLessThanOneMinute	 CodeErrorType = 10//间隔时间小于60秒不能收获
	RoleIdAlreadyRegister    	 CodeErrorType = 11//用户已经注册战斗了不要重复注册
	SendTooMuchSoldier			 CodeErrorType = 12//派出士兵过多
	AlreadyPassAllLevel          CodeErrorType = 13//玩家已经完全同关了
	UserAlreadyFull       		 CodeErrorType = 14 //玩家已经满级了
	NotEnoughResources			 CodeErrorType = 15 //用户资源不够
)

var ResourceTypeList = []ResourceType{
	FOOD,
	SOLDIER,
	GOLD,
}

type ResourceInfo struct {
	Type	ResourceType
	Count 	int
}

type HeroInfo struct {
	No   int
	Lv   int

}
