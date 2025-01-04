package config

import (
	"fmt"
	"github.com/pingc0y/URLFinder/cmd"
	"github.com/pingc0y/URLFinder/mode"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync"
)

var Conf mode.Config
var Progress = 1
var FuzzNum int

var (
	Risks = []string{"remove", "delete", "insert", "update", "logout"}

	JsFuzzPath = []string{
		"login.js",
		"app.js",
		"main.js",
		"config.js",
		"admin.js",
		"info.js",
		"open.js",
		"user.js",
		"input.js",
		"list.js",
		"upload.js",
	}
	JsFind = []string{
		"(https{0,1}:[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{3}[.]js)",
		"[\"'‘“`]\\s{0,6}(/{0,1}[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{3}[.]js)",
		"=\\s{0,6}[\",',’,”]{0,1}\\s{0,6}(/{0,1}[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\\+.~#?&//=]{3}[.]js)",
	}
	UrlFind = []string{
		"[\"'‘“`]\\s{0,6}(https{0,1}:[-a-zA-Z0-9()@:%_\\+.~#?&//={}]{2,250}?)\\s{0,6}[\"'‘“`]",
		"=\\s{0,6}(https{0,1}:[-a-zA-Z0-9()@:%_\\+.~#?&//={}]{2,250})",
		"[\"'‘“`]\\s{0,6}([#,.]{0,2}/[-a-zA-Z0-9()@:%_\\+.~#?&//={}]{2,250}?)\\s{0,6}[\"'‘“`]",
		"\"([-a-zA-Z0-9()@:%_\\+.~#?&//={}]+?[/]{1}[-a-zA-Z0-9()@:%_\\+.~#?&//={}]+?)\"",
		"href\\s{0,6}=\\s{0,6}[\"'‘“`]{0,1}\\s{0,6}([-a-zA-Z0-9()@:%_\\+.~#?&//={}]{2,250})|action\\s{0,6}=\\s{0,6}[\"'‘“`]{0,1}\\s{0,6}([-a-zA-Z0-9()@:%_\\+.~#?&//={}]{2,250})",
	}

	JsFiler = []string{
		"www\\.w3\\.org",
		"example\\.com",
	}
	UrlFiler = []string{
		"\\.js\\?|\\.css\\?|\\.jpeg\\?|\\.jpg\\?|\\.png\\?|.gif\\?|www\\.w3\\.org|example\\.com|\\<|\\>|\\{|\\}|\\[|\\]|\\||\\^|;|/js/|\\.src|\\.replace|\\.url|\\.att|\\.href|location\\.href|javascript:|location:|application/x-www-form-urlencoded|\\.createObject|:location|\\.path|\\*#__PURE__\\*|\\*\\$0\\*|\\n",
		".*\\.js$|.*\\.css$|.*\\.scss$|.*,$|.*\\.jpeg$|.*\\.jpg$|.*\\.png$|.*\\.gif$|.*\\.ico$|.*\\.svg$|.*\\.vue$|.*\\.ts$",
	}

	Phone     = []string{"[^\\w]((?:(?:\\+|00)86)?1(?:(?:3[\\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\\d])|(?:9[189]))\\d{8})[^\\w]"}
	Email     = []string{"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"}
	IDcard    = []string{"[^0-9]((\\d{8}(0\\d|10|11|12)([0-2]\\d|30|31)\\d{3}$)|(\\d{6}(18|19|20)\\d{2}(0[1-9]|10|11|12)([0-2]\\d|30|31)\\d{3}(\\d|X|x)))[^0-9]"}
	Jwt       = []string{"['\"](ey[A-Za-z0-9_-]{10,}\\.[A-Za-z0-9._-]{10,}|ey[A-Za-z0-9_\\/+-]{10,}\\.[A-Za-z0-9._\\/+-]{10,})['\"]"}
	Other     = []string{"(access.{0,1}key|access.{0,1}Key|access.{0,1}Id|access.{0,1}id|.{0,5}密码|.{0,5}账号|默认.{0,5}|加密|解密|password:.{0,10}|username:.{0,10})"}
	Webhook   = []string{"\\bhttps://qyapi.weixin.qq.com/cgi-bin/webhook/send\\?key=[a-zA-Z0-9\\-]{25,50}\\b", "\\bhttps://oapi.dingtalk.com/robot/send\\?access_token=[a-z0-9]{50,80}\\b", "\\bhttps://open.feishu.cn/open-apis/bot/v2/hook/[a-z0-9\\-]{25,50}\\b", "\\bhttps://hooks.slack.com/services/[a-zA-Z0-9\\-_]{6,12}/[a-zA-Z0-9\\-_]{6,12}/[a-zA-Z0-9\\-_]{15,24}\\b"}
	Grafana   = []string{"\\bglc_[A-Za-z0-9\\-_+/]{32,200}={0,2}\\b", "\\bglsa_[A-Za-z0-9]{32}_[A-Fa-f0-9]{8}\\b"}
	Ip        = []string{"(?:10\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3})|(?:172\\.(?:(?:1[6-9])|(?:2\\d)|(?:3[01]))\\.\\d{1,3}\\.\\d{1,3})|(?:192\\.168\\.\\d{1,3}\\.\\d{1,3})"}
	Swaggerui = []string{"((swagger-ui.html)|(\\\"swagger\\\":)|(Swagger UI)|(swaggerUi)|(swaggerVersion))"}
	Jdbc      = []string{"(jdbc:[a-z:]+://[a-z0-9\\.\\-_:;=/@?,&]+)"}
)

var (
	UrlSteps = 1
	JsSteps  = 3
)

var (
	Lock  sync.Mutex
	Wg    sync.WaitGroup
	Mux   sync.Mutex
	Ch    = make(chan int, 50)
	Jsch  = make(chan int, 50/10*3)
	Urlch = make(chan int, 50/10*7)
)

// 读取配置文件
func GetConfig(path string) {
	if f, err := os.Open(path); err != nil {
		if strings.Contains(err.Error(), "The system cannot find the file specified") || strings.Contains(err.Error(), "no such file or directory") {
			Conf.Headers = map[string]string{"Cookie": cmd.C, "User-Agent": `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0`, "Accept": "*/*"}
			Conf.Proxy = ""
			Conf.JsFind = JsFind
			Conf.UrlFind = UrlFind
			Conf.JsFiler = JsFiler
			Conf.UrlFiler = UrlFiler
			Conf.JsFuzzPath = JsFuzzPath
			Conf.JsSteps = JsSteps
			Conf.UrlSteps = UrlSteps
			Conf.Risks = Risks
			Conf.Timeout = cmd.TI
			Conf.Thread = cmd.T
			Conf.Max = cmd.MA
			Conf.InfoFind = map[string][]string{"Phone": Phone, "Email": Email, "IDcard": IDcard, "Jwt": Jwt, "webhook": Webhook, "grafana": Grafana, "ip": Ip, "swaggerui": Swaggerui, "jdbc": Jdbc, "Other": Other}
			data, err2 := yaml.Marshal(Conf)
			err2 = os.WriteFile(path, data, 0644)
			if err2 != nil {
				fmt.Println(err)
			} else {
				fmt.Println("未找到配置文件,已在当面目录下创建配置文件: config.yaml")
			}
		} else {
			fmt.Println("配置文件错误,请尝试重新生成配置文件")
			fmt.Println(err)
		}
		os.Exit(1)
	} else {
		yaml.NewDecoder(f).Decode(&Conf)
		JsFind = Conf.JsFind
		UrlFind = Conf.UrlFind
		JsFiler = Conf.JsFiler
		UrlFiler = Conf.UrlFiler
		JsFuzzPath = Conf.JsFuzzPath
		Phone = Conf.InfoFind["Phone"]
		Email = Conf.InfoFind["Email"]
		IDcard = Conf.InfoFind["IDcard"]
		Jwt = Conf.InfoFind["Jwt"]
		Other = Conf.InfoFind["Other"]
		Webhook = Conf.InfoFind["webhook"]
		Grafana = Conf.InfoFind["grafana"]
		Ip = Conf.InfoFind["ip"]
		Swaggerui = Conf.InfoFind["swaggerui"]
		Jdbc = Conf.InfoFind["jdbc"]
		JsSteps = Conf.JsSteps
		UrlSteps = Conf.UrlSteps
		Risks = Conf.Risks
		cmd.T = Conf.Thread
		cmd.MA = Conf.Max
		cmd.TI = Conf.Timeout
	}

}
