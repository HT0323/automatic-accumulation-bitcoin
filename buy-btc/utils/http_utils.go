package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// HTTPリクエストを送信する関数
func DoHttpRequest(method, url string, header, query map[string]string, data []byte) ([]byte, error) {
	if method != "GET" && method != "POST" {
		return nil, errors.New("method's GET nor POST")
	}

	// リクエスト情報の作成
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// queryParaに引数で渡された情報を設定
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// headerに引数で渡された情報を設定
	for key, value := range header {
		req.Header.Add(key, value)
	}

	// httpクライアントを作成しリクエストを投げレスポンスを受け取る
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
