package utils

import (
	"bee.com/conf"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"crypto/sha1"
)

func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	sort.Strings(arr)
	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}
	return Sha1(b.String()) == signature

}

func Sha1(str string) string  {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type ResAccessToken struct {
	AccessToken string	`json:"access_token"`
	ExpiresIn int64	`json:"expires_in"`
}
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

func GetTokenFromServer() (resAccessToken ResAccessToken, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", conf.C.WeChat.AppId, conf.C.WeChat.AppSecret)
	var body []byte
	body, err = HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	fmt.Println(resAccessToken)
	if err != nil {
		return
	}
	return resAccessToken, nil
}