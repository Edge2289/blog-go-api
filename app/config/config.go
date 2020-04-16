package config

var (
	ApiAuthConfig = map[string] map[string]string {

		// 调用方
		"DEMO" : {
			"md5" : "IgkibX71IEf382PT",
			"aes" : "IgkibX71IEf382PT",
			"rsa" : "rsa/public.pem",
		},
	}
)

const (

	// 网站配置
	AppMode = "release" //debug or release
	AppPort = ":9999"
	AppName = "go-gin-api"
	AppSalt = "blog-api-go-salt"  // 管理员加密密匙
	AppLang = "cn" // en 英文  cn  中文


	// redis
	RedisHost = "127.0.0.1" // redis 连接ip
	RedisPwd   = "" // redis 连接密码
	RedisPort = "6379" // redis 端口

	// database 数据库连接
	DbType = "mysql"
	DbHost = "127.0.0.1"
	DbPort = "3306"
	DbDatabase = "v_blog"
	DbUserName = "root"
	DbPassword = "123456"




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

	/**
	 * 这个是Jaeger的配置
	 */
	// Jaeger 配置信息
	JaegerHostPort = "127.0.0.1:6831"
	// Jaeger 配置开关 1=开通 -1=关闭
	JaegerOpen = 1
)