package config

type App struct {
	AppDebug          bool              `json:"AppDebug" yaml:"AppDebug"`
	HTTPServer        HTTPServer        `json:"HttpServer" yaml:"HttpServer"`
	Token             Token             `json:"Token" yaml:"Token"`
	Redis             Redis             `json:"Redis" yaml:"Redis"`
	Logs              Logs              `json:"Logs" yaml:"Logs"`
	Websocket         Websocket         `json:"Websocket" yaml:"Websocket"`
	SnowFlake         SnowFlake         `json:"SnowFlake" yaml:"SnowFlake"`
	Casbin            Casbin            `json:"Casbin" yaml:"Casbin"`
	RabbitMq          RabbitMq          `json:"RabbitMq" yaml:"RabbitMq"`
	Captcha           Captcha           `json:"Captcha" yaml:"Captcha"`
	FileUploadSetting FileUploadSetting `json:"FileUploadSetting" yaml:"FileUploadSetting"`
	Database          Database          `json:"Database" yaml:"Database"`
}

type Casbin struct {
	IsInit                int    `json:"IsInit" yaml:"IsInit"`
	AutoLoadPolicySeconds int    `json:"AutoLoadPolicySeconds" yaml:"AutoLoadPolicySeconds"`
	TablePrefix           string `json:"TablePrefix" yaml:"TablePrefix"`
	TableName             string `json:"TableName" yaml:"TableName"`
	ModelConfig           string `json:"ModelConfig" yaml:"ModelConfig"`
}

type RabbitMq struct {
	Enabled bool     `json:"Enabled" yaml:"Enabled"`
	Topics  []Topics `json:"Topics" yaml:"Topics"`
}

type Captcha struct {
	CaptchaID    string `json:"CaptchaId" yaml:"CaptchaId"`
	CaptchaValue string `json:"CaptchaValue" yaml:"CaptchaValue"`
	Length       int    `json:"Length" yaml:"Length"`
}

type FileUploadSetting struct {
	Size                 int      `json:"Size" yaml:"Size"`
	UploadFileField      string   `json:"UploadFileField" yaml:"UploadFileField"`
	UploadFileSavePath   string   `json:"UploadFileSavePath" yaml:"UploadFileSavePath"`
	UploadFileReturnPath string   `json:"UploadFileReturnPath" yaml:"UploadFileReturnPath"`
	UploadOSSTyper       string   `json:"UploadOSSTyper" yaml:"UploadOSSTyper"`
	AllowMimeType        []string `json:"AllowMimeType" yaml:"AllowMimeType"`
	OSS                  OSS      `json:"OSS" yaml:"OSS"`
}

type Database struct {
	UseDbType  string     `json:"UseDbType" yaml:"UseDbType"`
	Mysql      Mysql      `json:"Mysql" yaml:"Mysql"`
	SQLServer  SQLServer  `json:"SqlServer" yaml:"SqlServer"`
	PostgreSQL PostgreSQL `json:"PostgreSql" yaml:"PostgreSql"`
}

type API struct {
	Port string `json:"Port" yaml:"Port"`
}
type Web struct {
	Port string `json:"Port" yaml:"Port"`
}
type TrustProxies struct {
	IsOpen          int      `json:"IsOpen" yaml:"IsOpen"`
	ProxyServerList []string `json:"ProxyServerList" yaml:"ProxyServerList"`
}
type IPLimit struct {
	Count int `json:"Count" yaml:"Count"`
	Time  int `json:"Time" yaml:"Time"`
}
type HTTPServer struct {
	API              API          `json:"Api" yaml:"Api"`
	Web              Web          `json:"Web" yaml:"Web"`
	AllowCrossDomain bool         `json:"AllowCrossDomain" yaml:"AllowCrossDomain"`
	TrustProxies     TrustProxies `json:"TrustProxies" yaml:"TrustProxies"`
	IPLimit          IPLimit      `json:"IPLimit" yaml:"IPLimit"`
}
type Token struct {
	JwtTokenSignKey         string `json:"JwtTokenSignKey" yaml:"JwtTokenSignKey"`
	JwtTokenOnlineUsers     int    `json:"JwtTokenOnlineUsers" yaml:"JwtTokenOnlineUsers"`
	JwtTokenCreatedExpireAt int    `json:"JwtTokenCreatedExpireAt" yaml:"JwtTokenCreatedExpireAt"`
	JwtTokenRefreshAllowSec int    `json:"JwtTokenRefreshAllowSec" yaml:"JwtTokenRefreshAllowSec"`
	JwtTokenRefreshExpireAt int    `json:"JwtTokenRefreshExpireAt" yaml:"JwtTokenRefreshExpireAt"`
	BindContextKeyName      string `json:"BindContextKeyName" yaml:"BindContextKeyName"`
	IsCacheToRedis          int    `json:"IsCacheToRedis" yaml:"IsCacheToRedis"`
}
type Redis struct {
	Host               string `json:"Host" yaml:"Host"`
	Port               int    `json:"Port" yaml:"Port"`
	Auth               string `json:"Auth" yaml:"Auth"`
	MaxIdle            int    `json:"MaxIdle" yaml:"MaxIdle"`
	MaxActive          int    `json:"MaxActive" yaml:"MaxActive"`
	IdleTimeout        int    `json:"IdleTimeout" yaml:"IdleTimeout"`
	IndexDb            int    `json:"IndexDb" yaml:"IndexDb"`
	ConnFailRetryTimes int    `json:"ConnFailRetryTimes" yaml:"ConnFailRetryTimes"`
	ReConnectInterval  int    `json:"ReConnectInterval" yaml:"ReConnectInterval"`
}
type Logs struct {
	GinLogName        string `json:"GinLogName" yaml:"GinLogName"`
	GoSkeletonLogName string `json:"GoSkeletonLogName" yaml:"GoSkeletonLogName"`
	TextFormat        string `json:"TextFormat" yaml:"TextFormat"`
	TimePrecision     string `json:"TimePrecision" yaml:"TimePrecision"`
	Prefix            string `json:"Prefix" yaml:"Prefix"`
	MaxSize           int    `json:"MaxSize" yaml:"MaxSize"`
	MaxBackups        int    `json:"MaxBackups" yaml:"MaxBackups"`
	MaxAge            int    `json:"MaxAge" yaml:"MaxAge"`
	EncodeLevel       string `json:"EncodeLevel" yaml:"EncodeLevel"`
	FileMaxBackups    string `json:"FileMaxBackups" yaml:"FileMaxBackups"`
	FileMaxAge        bool   `json:"FileMaxAge" yaml:"FileMaxAge"`
	Compress          bool   `json:"Compress" yaml:"Compress"`
}
type Websocket struct {
	Enabled               bool `json:"Enabled" yaml:"Enabled"`
	WriteReadBufferSize   int  `json:"WriteReadBufferSize" yaml:"WriteReadBufferSize"`
	MaxMessageSize        int  `json:"MaxMessageSize" yaml:"MaxMessageSize"`
	PingPeriod            int  `json:"PingPeriod" yaml:"PingPeriod"`
	HeartbeatFailMaxTimes int  `json:"HeartbeatFailMaxTimes" yaml:"HeartbeatFailMaxTimes"`
	ReadDeadline          int  `json:"ReadDeadline" yaml:"ReadDeadline"`
	WriteDeadline         int  `json:"WriteDeadline" yaml:"WriteDeadline"`
}
type SnowFlake struct {
	SnowFlakeMachineID int `json:"SnowFlakeMachineId" yaml:"SnowFlakeMachineId"`
}

type Topics struct {
	Name                        string `json:"Name" yaml:"Name"`
	Addr                        string `json:"Addr" yaml:"Addr"`
	ExchangeType                string `json:"ExchangeType" yaml:"ExchangeType"`
	ExchangeName                string `json:"ExchangeName" yaml:"ExchangeName"`
	DelayedExchangeName         string `json:"DelayedExchangeName" yaml:"DelayedExchangeName"`
	Durable                     bool   `json:"Durable" yaml:"Durable"`
	QueueName                   string `json:"QueueName" yaml:"QueueName"`
	ConsumerChanNumber          int    `json:"ConsumerChanNumber" yaml:"ConsumerChanNumber"`
	OffLineReconnectIntervalSec int    `json:"OffLineReconnectIntervalSec" yaml:"OffLineReconnectIntervalSec"`
	RetryCount                  int    `json:"RetryCount" yaml:"RetryCount"`
}

type Local struct {
	Path      string `json:"Path" yaml:"Path"`
	StorePath string `json:"StorePath" yaml:"StorePath"`
}
type Qinniu struct {
	Zone       string `json:"Zone" yaml:"Zone"`
	Bucket     string `json:"Bucket" yaml:"Bucket"`
	ImagePath  string `json:"ImagePath" yaml:"ImagePath"`
	HTTPS      bool   `json:"Https" yaml:"Https"`
	AccessKey  string `json:"AccessKey" yaml:"AccessKey"`
	Secretkey  string `json:"Secretkey" yaml:"Secretkey"`
	CdnDomains bool   `json:"CdnDomains" yaml:"CdnDomains"`
}
type AliyunOSS struct {
	Endpoint  string `json:"Endpoint" yaml:"Endpoint"`
	AccessKey string `json:"AccessKey" yaml:"AccessKey"`
	Secretkey string `json:"Secretkey" yaml:"Secretkey"`
	Bucket    string `json:"Bucket" yaml:"Bucket"`
	BucketURL string `json:"BucketUrl" yaml:"BucketUrl"`
	BasePath  string `json:"BasePath" yaml:"BasePath"`
}
type TencentCOS struct {
	Bucket    string `json:"Bucket" yaml:"Bucket"`
	Region    string `json:"Region" yaml:"Region"`
	SecretID  string `json:"SecretID" yaml:"SecretID"`
	SecretKey string `json:"SecretKey" yaml:"SecretKey"`
	BaseURL   string `json:"BaseUrl" yaml:"BaseUrl"`
	Prefix    string `json:"Prefix" yaml:"Prefix"`
}
type AwsS3 struct {
	Bucket           string `json:"Bucket" yaml:"Bucket"`
	Region           string `json:"Region" yaml:"Region"`
	Endpoint         string `json:"Endpoint" yaml:"Endpoint"`
	S3ForcePathStyle bool   `json:"S3ForcePathStyle" yaml:"S3ForcePathStyle"`
	SSL              bool   `json:"SSL" yaml:"SSL"`
	AccessKey        string `json:"AccessKey" yaml:"AccessKey"`
	SecretKey        string `json:"SecretKey" yaml:"SecretKey"`
	BaseURL          string `json:"BaseUrl" yaml:"BaseUrl"`
	Prefix           string `json:"Prefix" yaml:"Prefix"`
}
type HuaweiOBS struct {
	Path      string `json:"Path" yaml:"Path"`
	Bucket    string `json:"Bucket" yaml:"Bucket"`
	Endpoint  string `json:"Endpoint" yaml:"Endpoint"`
	AccessKey string `json:"AccessKey" yaml:"AccessKey"`
	SecretKey string `json:"SecretKey" yaml:"SecretKey"`
}
type OSS struct {
	Local      Local      `json:"Local" yaml:"Local"`
	Qinniu     Qinniu     `json:"Qinniu" yaml:"Qinniu"`
	AliyunOSS  AliyunOSS  `json:"AliyunOSS" yaml:"AliyunOSS"`
	TencentCOS TencentCOS `json:"TencentCOS" yaml:"TencentCOS"`
	AwsS3      AwsS3      `json:"AWSS3" yaml:"AWSS3"`
	HuaweiOBS  HuaweiOBS  `json:"HuaweiOBS" yaml:"HuaweiOBS"`
}

type Write struct {
	Host               string `json:"Host" yaml:"Host"`
	DataBase           string `json:"DataBase" yaml:"DataBase"`
	Port               int    `json:"Port" yaml:"Port"`
	Prefix             string `json:"Prefix" yaml:"Prefix"`
	User               string `json:"User" yaml:"User"`
	Pass               string `json:"Pass" yaml:"Pass"`
	Charset            string `json:"Charset" yaml:"Charset"`
	SetMaxIdleConns    int    `json:"SetMaxIdleConns" yaml:"SetMaxIdleConns"`
	SetMaxOpenConns    int    `json:"SetMaxOpenConns" yaml:"SetMaxOpenConns"`
	SetConnMaxLifetime int    `json:"SetConnMaxLifetime" yaml:"SetConnMaxLifetime"`
	ReConnectInterval  int    `json:"ReConnectInterval" yaml:"ReConnectInterval"`
	PingFailRetryTimes int    `json:"PingFailRetryTimes" yaml:"PingFailRetryTimes"`
}
type Read struct {
	Host               string `json:"Host" yaml:"Host"`
	DataBase           string `json:"DataBase" yaml:"DataBase"`
	Port               int    `json:"Port" yaml:"Port"`
	Prefix             string `json:"Prefix" yaml:"Prefix"`
	User               string `json:"User" yaml:"User"`
	Pass               string `json:"Pass" yaml:"Pass"`
	Charset            string `json:"Charset" yaml:"Charset"`
	SetMaxIdleConns    int    `json:"SetMaxIdleConns" yaml:"SetMaxIdleConns"`
	SetMaxOpenConns    int    `json:"SetMaxOpenConns" yaml:"SetMaxOpenConns"`
	SetConnMaxLifetime int    `json:"SetConnMaxLifetime" yaml:"SetConnMaxLifetime"`
}

type Mysql struct {
	IsInitGlobalGormMysql int   `json:"IsInitGlobalGormMysql" yaml:"IsInitGlobalGormMysql"`
	SlowThreshold         int   `json:"SlowThreshold" yaml:"SlowThreshold"`
	Write                 Write `json:"Write" yaml:"Write"`
	IsOpenReadDb          int   `json:"IsOpenReadDb" yaml:"IsOpenReadDb"`
	Read                  Read  `json:"Read" yaml:"Read"`
}

type SQLServer struct {
	IsInitGlobalGormSqlserver int   `json:"IsInitGlobalGormSqlserver" yaml:"IsInitGlobalGormSqlserver"`
	SlowThreshold             int   `json:"SlowThreshold" yaml:"SlowThreshold"`
	Write                     Write `json:"Write" yaml:"Write"`
	IsOpenReadDb              int   `json:"IsOpenReadDb" yaml:"IsOpenReadDb"`
	Read                      Read  `json:"Read" yaml:"Read"`
}
type PostgreSQL struct {
	IsInitGlobalGormPostgreSQL int   `json:"IsInitGlobalGormPostgreSql" yaml:"IsInitGlobalGormPostgreSql"`
	SlowThreshold              int   `json:"SlowThreshold" yaml:"SlowThreshold"`
	Write                      Write `json:"Write" yaml:"Write"`
	IsOpenReadDb               int   `json:"IsOpenReadDb" yaml:"IsOpenReadDb"`
	Read                       Read  `json:"Read" yaml:"Read"`
}
