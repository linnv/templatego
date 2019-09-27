package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/linnv/logx"
)

func HttpPostJson(ctx string, targetUrl string, dataStruct interface{}, respStruct interface{}, client *http.Client) error {
	if client == nil {
		return fmt.Errorf("nil client")
	}

	postdata, err := json.Marshal(dataStruct)
	if err != nil {
		logx.Errorf(ctx+"err %v\n", err)
		return err
	}
	// logx.Debugfln(ctx+"->NLP postdata %s", postdata)

	req, err := http.NewRequest("POST", targetUrl, bytes.NewReader(postdata))
	if err != nil {
		logx.Errorf(ctx+"err %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logx.Errorf(ctx+"err %v\n", err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		bs, _ := httputil.DumpRequest(req, true)
		bsResp, _ := httputil.DumpResponse(resp, true)
		logx.Debugf("req %s \n resp: %s", bs, bsResp)
	}

	if respStruct != nil {
		if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
			logx.Errorf(ctx+"err %v\n", err)
			return err
		}
		resp.Body.Close()
	}
	return nil
}

func HttpGet(ctx string, targetUrl string, data string, respStruct interface{}, client *http.Client) error {
	if data != "" {
		targetUrl = targetUrl + data
	}
	resp, err := client.Get(targetUrl)
	if err != nil {
		logx.Errorf(ctx+"err %v\n", err)
		logx.Warnf("err: %+v\n", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		bsResp, _ := httputil.DumpResponse(resp, true)
		logx.Debugfln(ctx+"req %s \n resp: %s", targetUrl, bsResp)
	}

	if respStruct != nil {
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logx.Errorf(ctx+"err %v\n", err)
			logx.Warnf("err: %+v\n", err)
			return err
		}
		resp.Body.Close()
		logx.Debugfln(ctx+"ai resp [%s]", bs)
		if len(bs) < 10 {
			logx.Debugfln(ctx+"invalid bs %s", bs)
			return nil
		}
		err = json.Unmarshal(bs, &respStruct)
		if err != nil {
			logx.Errorf(ctx+"err %v\n", err)
			return err
		}
	}
	return nil
}
