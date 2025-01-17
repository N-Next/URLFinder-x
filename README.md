代码是在[URLFinder](https://github.com/pingc0y/URLFinder)项目基础上进行优化的，[readme](https://github.com/pingc0y/URLFinder/blob/master/README.md)可直接参考。

- [x] 添加并优化了多条正则表达式，目前可以匹配到**手机号、邮箱、身份证、JWT、accesskey、Webhook、内网IP、Swagger-UI、JDBC链接**。
- [x] 原项目对与未提取到js的网站不进行结果返回。优化后可对任意页面进行敏感信息提取。
- [x] 优化了显示逻辑。
- [ ] 添加阿里云/腾讯云/亚马逊云等云的key规则。
- [ ] 修改正则表达式添加逻辑，从yml文件中直接写好规则进行加载，like 灯塔的 WebInfoHunter。
- [ ] 优化代码，可作为一个模块被别的工具嵌套。

欢迎各位大佬提交issue
