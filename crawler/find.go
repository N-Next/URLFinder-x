package crawler

import (
	"github.com/pingc0y/URLFinder/cmd"
	"github.com/pingc0y/URLFinder/config"
	"github.com/pingc0y/URLFinder/mode"
	"github.com/pingc0y/URLFinder/result"
	"regexp"
	"strings"
)

// 分析内容中的js
func jsFind(cont, host, scheme, path, source string, num int) {
	var cata string
	care := regexp.MustCompile("/.*/{1}|/")
	catae := care.FindAllString(path, -1)
	if len(catae) == 0 {
		cata = "/"
	} else {
		cata = catae[0]
	}
	//js匹配正则
	host = scheme + "://" + host
	for _, re := range config.JsFind {
		reg := regexp.MustCompile(re)
		jss := reg.FindAllStringSubmatch(cont, -1)
		//return
		jss = jsFilter(jss)
		//循环提取js放到结果中
		for _, js := range jss {
			if js[0] == "" {
				continue
			}
			if strings.HasPrefix(js[0], "https:") || strings.HasPrefix(js[0], "http:") {
				switch AppendJs(js[0], source) {
				case 0:
					if num <= config.JsSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(js[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			} else if strings.HasPrefix(js[0], "//") {
				switch AppendJs(scheme+":"+js[0], source) {
				case 0:
					if num <= config.JsSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(scheme+":"+js[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			} else if strings.HasPrefix(js[0], "/") {
				switch AppendJs(host+js[0], source) {
				case 0:
					if num <= config.JsSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(host+js[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			} else {
				switch AppendJs(host+cata+js[0], source) {
				case 0:
					if num <= config.JsSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(host+cata+js[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			}
		}

	}

}

// 分析内容中的url
func urlFind(cont, host, scheme, path, source string, num int) {
	var cata string
	care := regexp.MustCompile("/.*/{1}|/")
	catae := care.FindAllString(path, -1)
	if len(catae) == 0 {
		cata = "/"
	} else {
		cata = catae[0]
	}
	host = scheme + "://" + host

	//url匹配正则

	for _, re := range config.UrlFind {
		reg := regexp.MustCompile(re)
		urls := reg.FindAllStringSubmatch(cont, -1)
		//fmt.Println(urls)
		urls = urlFilter(urls)

		//循环提取url放到结果中
		for _, url := range urls {
			if url[0] == "" {
				continue
			}
			if strings.HasPrefix(url[0], "https:") || strings.HasPrefix(url[0], "http:") {
				switch AppendUrl(url[0], source) {
				case 0:
					if num <= config.UrlSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(url[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}
			} else if strings.HasPrefix(url[0], "//") {
				switch AppendUrl(scheme+":"+url[0], source) {
				case 0:
					if num <= config.UrlSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(scheme+":"+url[0], num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			} else if strings.HasPrefix(url[0], "/") {
				urlz := ""
				if cmd.B != "" {
					urlz = cmd.B + url[0]
				} else {
					urlz = host + url[0]
				}
				switch AppendUrl(urlz, source) {
				case 0:
					if num <= config.UrlSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(urlz, num+1)
					}
				case 1:
					return
				case 2:
					continue
				}
			} else if !strings.HasSuffix(source, ".js") {
				urlz := ""
				if cmd.B != "" {
					if strings.HasSuffix(cmd.B, "/") {
						urlz = cmd.B + url[0]
					} else {
						urlz = cmd.B + "/" + url[0]
					}
				} else {
					urlz = host + cata + url[0]
				}
				switch AppendUrl(urlz, source) {
				case 0:
					if num <= config.UrlSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(urlz, num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			} else if strings.HasSuffix(source, ".js") {
				urlz := ""
				if cmd.B != "" {
					if strings.HasSuffix(cmd.B, "/") {
						urlz = cmd.B + url[0]
					} else {
						urlz = cmd.B + "/" + url[0]
					}
				} else {
					config.Lock.Lock()
					su := result.Jsinurl[source]
					config.Lock.Unlock()
					if strings.HasSuffix(su, "/") {
						urlz = su + url[0]
					} else {
						urlz = su + "/" + url[0]
					}
				}
				switch AppendUrl(urlz, source) {
				case 0:
					if num <= config.UrlSteps && (cmd.M == 2 || cmd.M == 3) {
						config.Wg.Add(1)
						config.Ch <- 1
						go Spider(urlz, num+1)
					}
				case 1:
					return
				case 2:
					continue
				}

			}
		}
	}
}

// 分析内容中的敏感信息
func infoFind(cont, source string) {
	info := []mode.Info{}
	//cont = strings.ReplaceAll(cont, "\r", "")
	//cont = strings.ReplaceAll(cont, "\n", "")
	for key, regexps := range config.Infofind {
		for _, regexpstr := range regexps {
			values := regexp.MustCompile(regexpstr).FindAllStringSubmatch(cont, -1)
			if values != nil {
				for _, value := range values {
					found := false
					for i, inf := range info {
						if inf.Key == key {
							info[i].Matches = append(info[i].Matches, value...)
							found = true
							break
						}
					}
					if !found {
						info = append(info, mode.Info{Key: key, Matches: value, Source: source})
					}
				}
			}
		}
	}
	if len(info) != 0 {
		for _, singleInfo := range info {
			AppendInfo(singleInfo)
		}
	}
}
