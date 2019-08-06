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
	// ERC20APIKey service/erc20/apikey
	ERC20APIKey
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
	// UserPort service/user/port
	UserPort
	// KeystorePort service/user-keystore/port
	KeystorePort
	// SessionPort service/session/port
	SessionPort
	// RRDWxAppID service/weixin-app/rrd/app-id
	RRDWxAppID
	// RRDWxAppSecret service/weixin-app/rrd/app-secret
	RRDWxAppSecret
	// RRDWxGetSessionKeyUrl service/weixin-app/rrd/auth/get-session-key-url
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
	// RRDWxRefundNotifyURL service/weixin/rrd/weixin-payment/refund-notify
	RRDWxRefundNotifyURL
	// RRDWxAppAuthPort  service/weixin-app/rrd/auth/port
	RRDWxAppAuthPort
	// RRDWxAppPaymentPort service/weixin-app/rrd/payment/port
	RRDWxAppPaymentPort
	// RRDWxAppWithdrawUrl service/weixin-app/rrd/withdraw-url
	RRDWxAppWithdrawUrl
	// RRDWxAppPaymentNotifyURL service/weixin-app/rrd/payment/notifyurl
	RRDWxAppPaymentNotifyURL
	// RRDWxAppRefundNotifyURL service/weixin-app/rrd/payment/refund-notify
	RRDWxAppRefundNotifyURL
	// RRDWxAuthPort service/weixin/rrd/weixin-auth/port
	RRDWxAuthPort
	// RRPointsSaasBillingPort service/rrpoints-saas/billing/port
	RRPointsSaasBillingPort
	// RRPointsSaasBillingMysqlUser service/rrpoints-saas/billing/mysql/user
	RRPointsSaasBillingMysqlUser
	// RRPointsSaasBillingMysqlPswd service/rrpoints-saas/billing/mysql/pswd
	RRPointsSaasBillingMysqlPswd
	// RRPointsSaasBillingMysqlDbName service/rrpoints-saas/billing/mysql/dbName
	RRPointsSaasBillingMysqlDbName
	// RRPointsSaasWxPayPort service/rrpoints-saas/wx-pay/port
	RRPointsSaasWxPayPort
	// RRPointsSaasWxPayApiKey service/rrpoints-saas/wx-pay/apikey
	RRPointsSaasWxPayApiKey
	// RRPointsSaasWxPayAppId service/rrpoints-saas/wx-pay/appid
	RRPointsSaasWxPayAppId
	// RRPointsSaasWxPayMchId service/rrpoints-saas/wx-pay/mchid
	RRPointsSaasWxPayMchId
	// RRPointsSaasWxPayNoticeAddr service/rrpoints-saas/wx-pay/notice-addr
	RRPointsSaasWxPayNoticeAddr
	// RRPointsSaasWxPayNoticeIsProduction service/rrpoints-saas/wx-pay/is-production
	RRPointsSaasWxPayNoticeIsProduction
	// RRPointsSaasMongoDB db/mongodb/rrpoints-saas
	RRPointsSaasMongoDB
	// RRPointsSaasAppPort service/rrpoints-saas/app/port
	RRPointsSaasAppPort
	// RRPointsSaasAppBanners service/rrpoints-saas/app/banners
	RRPointsSaasAppBanners
	// RRPointsSaasResourcesAccessKey service/rrpoints-saas/resources/access-key
	RRPointsSaasResourcesAccessKey
	// RRPointsSaasResourcesSecretKey service/rrpoints-saas/resources/secret-key
	RRPointsSaasResourcesSecretKey
	// RRPointsSaasResourcesBucket service/rrpoints-saas/resources/bucket
	RRPointsSaasResourcesBucket
	// RRPointsSaasResourcesDomain service/rrpoints-saas/resources/domain
	RRPointsSaasResourcesDomain

	// TokenPort service/token/port
	TokenPort
)
