package task

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/fghwett/toolsign/config"
	"github.com/fghwett/toolsign/util"
)

type Task struct {
	cookie    string
	userAgent string
	client    *http.Client
	result    []string
}

func New(config *config.Config) *Task {
	return &Task{
		cookie:    config.Cookie,
		userAgent: config.UserAgent,
		client:    &http.Client{},
		result:    []string{"==== Tool.lu签到任务 ===="},
	}
}

func (t *Task) Do() {
	if err := t.signTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【签到任务】：失败 %s", err))
		return
	}

	util.SmallSleep(1000, 3000)

	if err := t.getPointTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【积分查询】：失败 %s", err))
		return
	}
}

func (t *Task) signTask() error {
	reqUrl := "https://plus.tool.lu/user/sign"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("referer", "https://tool.lu/")
	req.Header.Set("User-Agent", t.userAgent)
	req.Header.Set("cookie", t.cookie)

	resp, err := t.client.Do(req)

	body, err := util.GetHTTPResponseOrg(resp, reqUrl, err)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return err
	}

	var result []string
	doc.Find("section[class=panel-body] p").Each(func(i int, s *goquery.Selection) {
		result = append(result, strings.TrimSpace(s.Text()))
	})

	t.result = append(t.result, fmt.Sprintf("【签到任务】：成功 %s", strings.Join(result, " ")))

	return nil
}

func (t *Task) getPointTask() error {
	reqUrl := "https://plus.tool.lu/user/credits"
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("referer", "https://plus.tool.lu/user/sign")
	req.Header.Set("User-Agent", t.userAgent)
	req.Header.Set("cookie", t.cookie)

	resp, err := t.client.Do(req)

	body, err := util.GetHTTPResponseOrg(resp, reqUrl, err)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return err
	}

	var result []string
	doc.Find(".badge").Each(func(i int, s *goquery.Selection) {
		result = append(result, strings.TrimSpace(s.Text()))
	})

	t.result = append(t.result, fmt.Sprintf("【积分查询】：成功 余额%s", strings.Join(result, " ")))

	return nil
}

func (t *Task) GetResult() string {
	return strings.Join(t.result, " \n\n ")
}
