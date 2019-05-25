package api

const (
	// AppraisalAssetPort service/appraisal-asset/port 说明
	AppraisalAssetPort = iota
	// AppraisalUserPort service/appraisal-user/port 说明
	AppraisalUserPort
	// NsqUrl nsq/nsqd
	NsqUrl
	// Haproxy haproxy
	Haproxy
	// SMSPort service/sms/port
	SMSPort
	// AppraisalEvaluationPort service/appraisal-evaluation/port
	AppraisalEvaluationPort
	// AppraisalKeystorePort service/appraisal-user-keystore/port
	AppraisalKeystorePort
	// TokenFlowPort service/token-flow/port
	TokenFlowPort
	// ERC20Port service/erc20/port
	ERC20Port
	// ERC20Payee service/erc20/payee
	ERC20Payee
	// MongoDbAppraisalEvaluation db/mongodb/appraisal-evaluation 鉴宝留言
	MongoDbAppraisalEvaluation
	// DuiBaAppKey service/duiba/app_key 兑吧appKey
	DuiBaAppKey
	// DuiBaSecretKey service/duiba/app_secret 兑吧secretKey
	DuiBaSecretKey
	// MysqlURL db/mysql/mysqlurl
	MysqlURL
	// RedisURI redis/cache/host   redis链接
	RedisURI
	// RedisPassword redis/cache/password  redis密码
	RedisPassword
	// UserPort　service/user/port
	UserPort
	// KeystorePort　service/user-keystore/port
	KeystorePort
	// SessionPort service/session/port
	SessionPort
	// RRDWxAppID service/weixin-app/rrd/app-id
	RRDWxAppID
	// RRDWxAppSecret service/weixin-app/rrd/app-secret
	RRDWxAppSecret
	// RRDWxGetSessionKeyUrl service/weixin-app/rrd/get-session-key-url
	RRDWxGetSessionKeyUrl
	// MongoDbWeixinPaymentURI db/mongodb/weixin-payment
	MongoDbWeixinPaymentURI
	// CommonAPI service/common-api
	CommonAPI
	// RRDAppKey service/weixin/rrd/appkey
	RRDAppKey
	// RRDMchID service/weixin/rrd/mch-id
	RRDMchID
	// RRDWxPaymentNotifyURL service/weixin/rrd/weixin-payment/notifyurl
	RRDWxPaymentNotifyURL
	// RRdwxRefundNotifyURL service/weixin/rrd/weixin-payment/refund-notify
	RRdwxRefundNotifyURL
	// RRDWxAppAuthPort  service/weixin-app/rrd/auth-port
	RRDWxAppAuthPort
	// RRDWxAppPaymentPort service/weixin-app/rrd/payment-port
	RRDWxAppPaymentPort
	// RRDWxAppWithdrawUrl service/weixin-app/rrd/withdraw-url
	RRDWxAppWithdrawUrl
)