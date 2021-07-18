package Flags

type Flags struct {
	/** 通用 */
	Debug   bool  /** 是否开启Debug */
	Help    bool  /** 是否显示帮助文档 */
	Type    int   /** 类型 */
	Thread  int   /** 线程数 */
	TimeOut int64 /** 超时时间 */

	/** 数据库用到 */
	DataSourceName string /** root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4 */
	ShowDatabases  bool
	Database       string
	ShowTables     bool
	Table          string
	ShowColumns    bool
	Columns        string
	Sql            string
	Id             int

	/** 帐号密码 */
	Username string
	Password string
	Captcha  string

	/** 文件操作 */
	FileName  string
	Url       string

	/** HTTP请求使用 */
	UserAgent string
	Token     string

	/** 延时使用 */
	SleepTime  int64
	SleepTime1 int64
	SleepTime2 int64

	/** 文本 */
	Text string

	/** 其他 */
	Level int
}
