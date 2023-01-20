package main

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var CheckUrl = "https://debug-vas.intlgame.com/sub-api/validateRepoUrl"

func main() {
	DownLoad("https://beta-intl-market-artifacts-1300342648.file.myqcloud.com/repo-files/2022-11-11/a20221031exchangeshop.zip", "E:\\a20221031exchangeshop.zip")
	//ValidUrl("https://e.coding.intlgame.com/ptc/morikomorilifeactcgi_protocol/a20221031exchangeshop.git", CheckUrl)
}

type CheckTag struct {
	RepoUrl  string `json:"repo_url"`
	RepoType string `json:"repo_type"`
}

func ValidUrl(url string, path string) {
	jsonData := fmt.Sprintf(`{"repo_url":%q,"repo_type":%q}`, url, "coding")
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader([]byte(jsonData)))
	if err != nil {
		log.Errorf(context.Background(), "http new request err:%v", err.Error())
		return
	}
	req.Header.Set("token", "wang_Shao") // 无实际意义,携带即可
	req.Header.Set("Content-Type", "application/json")
	// 设置数据
	client := &http.Client{Timeout: time.Second * 5}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf(context.Background(), "send request do for http client err:%v", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf(context.Background(), "ioutil read all err:%v", err.Error())
		return
	}

	fmt.Println(string(body))

}

func DownLoad(url string, path string) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// Create output file
	out, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// copy stream
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
