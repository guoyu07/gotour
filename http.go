// http
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	Url "net/url"
	"strings"
	"time"
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

type Response struct {
	data string
	err  error
}

type Payload struct {
	url    string
	param  *map[string][]string
	isPost bool
}

func (resp Response) String() string {
	err := "nil"
	if nil != resp.err {
		err = resp.err.Error()
	}
	return "{" + resp.data + "," + err + "}"
}

func parallelGetData(ch chan *Response, payloads []*Payload) []*Response {
	for _, payload := range payloads {
		go func(ch chan *Response, url string, params *map[string][]string, isPost bool) {
			data, err := getData(url, params, isPost)
			ch <- &Response{data, err}
		}(ch, (*payload).url, (*payload).param, (*payload).isPost)
	}
	var resps = []*Response{}
	len := len(payloads)
	for i := 0; i < len; i++ {
		resps = append(resps, <-ch)
	}
	close(ch)
	return resps
}

func main() {
	url := "http://www.tuling123.com/openapi/api"
	params := &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"你好"}}
	ta1 := time.Now().UnixNano()
	data, _ := getData(url, params, true)
	fmt.Println((time.Now().UnixNano() - ta1) / 1e6)
	fmt.Println(data)

	ch := make(chan *Response)
	payloads := []*Payload{
		&Payload{
			url:    "http://www.tuling123.com/openapi/api",
			param:  &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"你好"}},
			isPost: true,
		},
		&Payload{
			url:    "http://www.tuling123.com/openapi/api",
			param:  &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"吃饭"}},
			isPost: true,
		},
		&Payload{
			url:    "http://www.tuling123.com/openapi/api",
			param:  &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"睡觉"}},
			isPost: true,
		},
		&Payload{
			url:    "http://www.tuling123.com/openapi/api",
			param:  &map[string][]string{"key": {"f59ec55a36784b648e6979e6b1ba4679"}, "info": {"打豆豆"}},
			isPost: true,
		},
	}
	tb1 := time.Now().UnixNano()
	resps := parallelGetData(ch, payloads)
	fmt.Println((time.Now().UnixNano() - tb1) / 1e6)
	for _, resp := range resps {
		fmt.Println(resp)
	}
}
