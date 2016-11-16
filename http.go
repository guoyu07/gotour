// http
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	Url "net/url"
	"strings"
)

func request(url string, data *map[string][]string, isPost bool) (*http.Response, error) {
	if isPost {
		return http.PostForm(url, Url.Values(*data))
	} else {
		if data != nil && len(*data) > 0 {
			idx := strings.Index(url, "?")
			//如果url里已经包含问号
			if idx != -1 {
				//如果仅仅有问号但无参数
				if idx == (len(url) - 1) {
					url += Url.Values(*data).Encode()
				} else { //如果有问号也有其他参数
					url += "&" + Url.Values(*data).Encode()
				}
			}
		}
		return http.Get(url)
	}
}

func getData(url string, params *map[string][]string, isPost bool) (string, error) {
	resp, ok := request(url, params, isPost)
	var data string
	if ok == nil {
		ret, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err == nil {
			data = string(ret)
		} else {
			return data, err
		}
	} else {
		return data, ok
	}

	return data, nil
}

func main() {
	url := "http://www.tuling123.com/openapi/api"
	params := &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"你好"}}
	data, _ := getData(url, params, true)
	fmt.Println(data)
}
