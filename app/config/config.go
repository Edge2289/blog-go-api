package config

const (

	// 网站配置
	AppMode = "release" //debug or release
	AppPort = ":9999"
	AppName = "go-gin-api"
	AppSalt = "blog-api-go-salt"  // 管理员加密密匙
	AppLang = "cn" // en 英文  cn  中文
	AppBlogRouter = "/api/blog/"

	// redis
	RedisHost = "127.0.0.1" // redis 连接ip
	RedisPwd   = "" // redis 连接密码
	RedisPort = "6379" // redis 端口

	// database 数据库连接
	DbType = "mysql"
	DbHost = "127.0.0.1"
	DbPort = "3306"
	DbDatabase = "blog_refactoring"
	DbUserName = "root"
	DbPassword = "root"

	// 綫上數據庫
	//DbHost = "39.106.158.255"
	//DbPort = "3306"
	//DbDatabase = "blog_ refactoring"
	//DbUserName = "root"
	//DbPassword = "xie19970326"

	// RabbitMQ
	MqHost = "10.10.107.36:15672" // 连接名
	MqUser = "admin"
	MqPwd = "123456"

	// 七牛云配置
	ACCESS_KEY = "Q0zPp-M3h3Ct5-p5IFCmAmyX89ZpvQbu-DVb1dmM"
	SECRET_KEY = "dzbhcSBBF0PnilMcMDVKz-ykLAxP5jkP0tRyKXJV"
	BUCKET     = "treenblogimg" // 存储的空间名称
	ImgDomain = "http://q9onx63t3.bkt.clouddn.com/"

	// es
	ES_HOST = "http://139.9.38.4:9200/"

	// JWT加密参数
	JwtSecretKey = "xiaoxiao-blog-gp-api"
	AppSignExpiry = "120"

	// 超时时间
	AppReadTimeout  = 120
	AppWriteTimeout = 120

	// 日志文件
	LogDirName = "storage/logs/api/"
	LogDrive = "file"  // file 文件日志   es  MQ  英文逗号隔开，推送日志

	/**
	 * 邮箱配置
	 */
	SystemEmailUser = "xinliangnote@163.com"
	SystemEmailPass = "" //密码或授权码
	SystemEmailHost = "smtp.163.com"
	SystemEmailPort = 465

	/**
	 * 这个是告警的配置
	 */
	// 告警接收人
	ErrorNotifyUser = "xinliangnote@163.com"
	// 告警开关 1=开通 -1=关闭
	ErrorNotifyOpen = -1


)