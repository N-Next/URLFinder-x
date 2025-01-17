# 简介
代码是在[URLFinder](https://github.com/pingc0y/URLFinder)项目基础上进行优化的，[readme](https://github.com/pingc0y/URLFinder/blob/master/README.md)可直接参考。

- [x] 添加并优化了多条正则表达式，目前可以匹配到**手机号、邮箱、身份证、JWT、accesskey、Webhook、内网IP、Swagger-UI、JDBC链接**。20
- [x] 原项目对与未提取到js的网站不进行结果返回。优化后可对任意页面进行敏感信息提取。
- [x] 优化了显示逻辑。
- [ ] ~~添加阿里云/腾讯云/亚马逊云等云的key规则。~~（新版用户可以自定义规则）
- [x] 修改正则表达式添加逻辑，从yml文件中直接写好规则进行加载，like 灯塔的 [WebInfoHunter](https://tophanttechnology.github.io/ARL-doc/function_desc/web_info_hunter/)。


新版的URLFinder-x做了代码的简单重构。将粗略的完成部分灯塔的 WebInfoHunter功能与URLFinder进行融合。实现了可以通过自定义规则查找敏感信息。

可以使用-i选项自动导出config.yml或者直接使用我的yml。在`infoFiler`下增添规则。
~~~
proxy: ""
timeout: 5
thread: 50
urlSteps: 1
jsSteps: 3
max: 99999
headers:
    Accept: '*/*'
    Cookie: ""
    User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0
jsFind:
    - (https{0,1}:[-a-zA-Z0-9（）@:%_\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\+.~#?&//=]{3}[.]js)
    - '["''‘“`]\s{0,6}(/{0,1}[-a-zA-Z0-9（）@:%_\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\+.~#?&//=]{3}[.]js)'
    - =\s{0,6}[",',’,”]{0,1}\s{0,6}(/{0,1}[-a-zA-Z0-9（）@:%_\+.~#?&//=]{2,250}?[-a-zA-Z0-9（）@:%_\+.~#?&//=]{3}[.]js)
urlFind:
    - '["''‘“`]\s{0,6}(https{0,1}:[-a-zA-Z0-9()@:%_\+.~#?&//={}]{2,250}?)\s{0,6}["''‘“`]'
    - =\s{0,6}(https{0,1}:[-a-zA-Z0-9()@:%_\+.~#?&//={}]{2,250})
    - '["''‘“`]\s{0,6}([#,.]{0,2}/[-a-zA-Z0-9()@:%_\+.~#?&//={}]{2,250}?)\s{0,6}["''‘“`]'
    - '"([-a-zA-Z0-9()@:%_\+.~#?&//={}]+?[/]{1}[-a-zA-Z0-9()@:%_\+.~#?&//={}]+?)"'
    - href\s{0,6}=\s{0,6}["'‘“`]{0,1}\s{0,6}([-a-zA-Z0-9()@:%_\+.~#?&//={}]{2,250})|action\s{0,6}=\s{0,6}["'‘“`]{0,1}\s{0,6}([-a-zA-Z0-9()@:%_\+.~#?&//={}]{2,250})
infoFiler:
    Email:
        - '(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))'
    IDcard:
        - '[1-9]\d{5}(?:18|19|20)\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\d|30|31)\d{3}[\dXx]'
    Jwt:
        - '[''"](ey[A-Za-z0-9_-]{10,}\.[A-Za-z0-9._-]{10,}|ey[A-Za-z0-9_\/+-]{10,}\.[A-Za-z0-9._\/+-]{10,})[''"]'
    Other:
        - '(access.{0,1}key|access.{0,1}Key|access.{0,1}Id|access.{0,1}id|.{0,5}密码|.{0,5}账号|默认.{0,5}|加密|解密|password:.{0,10}|username:.{0,10})'
    Phone:
        - '(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[01256789]))\d{8}'
    ip:
        - '(?:10\.\d{1,3}\.\d{1,3}\.\d{1,3})|(?:172\.(?:(?:1[6-9])|(?:2\d)|(?:3[01]))\.\d{1,3}\.\d{1,3})|(?:192\.168\.\d{1,3}\.\d{1,3})'
    jdbc:
        - '(jdbc:[a-z:]+://[a-z0-9\.\-_:;=/@?,&]+)'
    swaggerui:
        - '((swagger-ui.html)|(\"swagger\":)|(Swagger UI)|(swaggerUi)|(swaggerVersion))'
    webhook:
        - '\bhttps://qyapi.weixin.qq.com/cgi-bin/webhook/send\?key=[a-zA-Z0-9\-]{25,50}\b'
        - '\bhttps://oapi.dingtalk.com/robot/send\?access_token=[a-z0-9]{50,80}\b'
        - '\bhttps://open.feishu.cn/open-apis/bot/v2/hook/[a-z0-9\-]{25,50}\b'
        - '\bhttps://hooks.slack.com/services/[a-zA-Z0-9\-_]{6,12}/[a-zA-Z0-9\-_]{6,12}/[a-zA-Z0-9\-_]{15,24}\b'
    github_access_token:
        - '[\w\-]*:[\w\-]+@github\.com*'
risks:
    - remove
    - delete
    - insert
    - update
    - logout
jsFiler:
    - www\.w3\.org
    - example\.com
urlFiler:
    - \.js\?|\.css\?|\.jpeg\?|\.jpg\?|\.png\?|.gif\?|www\.w3\.org|example\.com|\<|\>|\{|\}|\[|\]|\||\^|;|/js/|\.src|\.replace|\.url|\.att|\.href|location\.href|javascript:|location:|application/x-www-form-urlencoded|\.createObject|:location|\.path|\*#__PURE__\*|\*\$0\*|\n
    - .*\.js$|.*\.css$|.*\.scss$|.*,$|.*\.jpeg$|.*\.jpg$|.*\.png$|.*\.gif$|.*\.ico$|.*\.svg$|.*\.vue$|.*\.ts$
jsFuzzPath:
    - login.js
    - app.js
    - main.js
    - config.js
    - admin.js
    - info.js
    - open.js
    - user.js
    - input.js
    - list.js
    - upload.js

~~~
**注意：在正则表达式添加的时候需要使用`''`将表达式包裹，否则会报错**

在测试过程中发现常规敏感信息正则表达式在复杂的返回报文中无法完全适用，准确率不高。需要自行测试表达式是否准确适用，这里我提供一个简单的demo方便测试，以及一个[rule](https://github.com/any86/any-rule)库
~~~
func main() {
	text := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>示例页面</title>
</head>
<body>
    <h1>用户信息</h1>
    <p>电子邮件: user123@example.com</p>
    <p>手机号: 13800138000</p>
    <p>身份证号: 110101199001011234</p>
    <p>内网 IP: 192.168.1.100</p>
</body>
</html>`
	// 编译正则表达式
	re := regexp.MustCompile(`[1-9]\d{5}(?:18|19|20)\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\d|30|31)\d{3}[\dXx]`)
	// 替换换行符、回车符和多余的空白字符
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.TrimSpace(text)
	// 查找所有匹配的手机号码
	matches := re.FindAllString(text, -1)
	fmt.Println("匹配到的手机号码:", matches)
}

~~~
- [ ] ~~优化代码，可作为一个模块被别的工具嵌套。~~

欢迎各位大佬提交issue
# 参考
https://github.com/pingc0y/URLFinder

https://tophanttechnology.github.io/ARL-doc/function_desc/web_info_hunter/

https://xz.aliyun.com/t/13993?time__1311=GqmxnD2DyD97KGNDQ0P7KpemwcArF7a4D
# 免责声明
本工具仅面向合法授权的企业安全建设行为，如您需要测试本工具的可用性，请自行搭建靶机环境。

为避免被恶意使用，本项目不会对目标发起真实攻击和漏洞利用。

在使用本工具进行检测时，您应确保该行为符合当地的法律法规，并且已经取得了足够的授权。请勿对非授权目标进行扫描。

如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任。

在安装并使用本工具前，请您务必审慎阅读、充分理解各条款内容，限制、免责条款或者其他涉及您重大权益的条款可能会以加粗、加下划线等形式提示您重点注意。 除非您已充分阅读、完全理解并接受本协议所有条款，否则，请您不要安装并使用本工具。您的使用行为或者您以其他任何明示或者默示方式表示接受本协议的，即视为您已阅读并同意本协议的约束。
