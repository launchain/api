// Code generated by launch-base-generate DO NOT EDIT
package api

// ConsulConfig ...
type ConsulConfig struct {
	AppraisalAssetPort string `consul:"service/appraisal-asset/port"`

	AppraisalEvaluationPort string `consul:"service/appraisal-evaluation/port"`

	AppraisalKeystorePort string `consul:"service/appraisal-user-keystore/port"`

	AppraisalUserPort string `consul:"service/appraisal-user/port"`

	CommonAPI string `consul:"service/common-api"`

	DuiBaAppKey string `consul:"service/duiba/app_key"`

	DuiBaSecretKey string `consul:"service/duiba/app_secret"`

	ERC20APIKey string `consul:"service/erc20/apikey"`

	ERC20Payee string `consul:"service/erc20/payee"`

	ERC20Port string `consul:"service/erc20/port"`

	Haproxy string `consul:"haproxy"`

	KeystorePort string `consul:"service/user-keystore/port"`

	MongoDbAppraisalEvaluation string `consul:"db/mongodb/appraisal-evaluation"`

	MongoDbWeixinPaymentURI string `consul:"db/mongodb/weixin-payment"`

	MysqlURL string `consul:"db/mysql/mysqlurl"`

	NsqUrl string `consul:"nsq/nsqd"`

	RRDAppKey string `consul:"service/weixin/rrd/appkey"`

	RRDMchID string `consul:"service/weixin/rrd/mch-id"`

	RRDWxAppAuthPort string `consul:"service/weixin-app/rrd/auth/port"`

	RRDWxAppID string `consul:"service/weixin-app/rrd/app-id"`

	RRDWxAppPaymentNotifyURL string `consul:"service/weixin-app/rrd/payment/notifyurl"`

	RRDWxAppPaymentPort string `consul:"service/weixin-app/rrd/payment/port"`

	RRDWxAppRefundNotifyURL string `consul:"service/weixin-app/rrd/payment/refund-notify"`

	RRDWxAppSecret string `consul:"service/weixin-app/rrd/app-secret"`

	RRDWxAppWithdrawUrl string `consul:"service/weixin-app/rrd/withdraw-url"`

	RRDWxAuthPort string `consul:"service/weixin/rrd/weixin-auth/port"`

	RRDWxGetSessionKeyUrl string `consul:"service/weixin-app/rrd/auth/get-session-key-url"`

	RRDWxPaymentNotifyURL string `consul:"service/weixin/rrd/weixin-payment/notifyurl"`

	RRDWxRefundNotifyURL string `consul:"service/weixin/rrd/weixin-payment/refund-notify"`

	RRPointsSaasBillingMysqlDbName string `consul:"service/rrpoints-saas/billing/mysql/dbName"`

	RRPointsSaasBillingMysqlPswd string `consul:"service/rrpoints-saas/billing/mysql/pswd"`

	RRPointsSaasBillingMysqlUser string `consul:"service/rrpoints-saas/billing/mysql/user"`

	RRPointsSaasBillingPort string `consul:"service/rrpoints-saas/billing/port"`

	RRPointsSaasMongoDB string `consul:"db/mongodb/rrpoints-saas"`

	RRPointsSaasWxPayApiKey string `consul:"service/rrpoints-saas/wx-pay/apikey"`

	RRPointsSaasWxPayAppId string `consul:"service/rrpoints-saas/wx-pay/appid"`

	RRPointsSaasWxPayMchId string `consul:"service/rrpoints-saas/wx-pay/mchid"`

	RRPointsSaasWxPayNoticeAddr string `consul:"service/rrpoints-saas/wx-pay/notice-addr"`

	RRPointsSaasWxPayNoticeIsProduction string `consul:"service/rrpoints-saas/wx-pay/is-production"`

	RRPointsSaasWxPayPort string `consul:"service/rrpoints-saas/wx-pay/port"`

	RedisPassword string `consul:"redis/cache/password"`

	RedisURI string `consul:"redis/cache/host"`

	SMSPort string `consul:"service/sms/port"`

	SessionPort string `consul:"service/session/port"`

	TokenFlowPort string `consul:"service/token-flow/port"`

	UserPort string `consul:"service/user/port"`
}

// NewDefaultConsulConfig ...
func NewDefaultConsulConfig() *ConsulConfig {
	ret := &ConsulConfig{}

	return ret
}
