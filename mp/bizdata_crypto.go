package mp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shenghui0779/gochat/utils"
)

// UserData 用户数据
type UserData struct {
	OpenID    string     `json:"openId"`
	Language  string     `json:"language"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	AvatarURL string     `json:"avatarUrl"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"`
	Country   string     `json:"country"`
	UnionID   string     `json:"unionId"`
	WaterMark *WaterMark `json:"watermark"`
}

// PhoneData 用户手机号绑定数据
type PhoneData struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	WaterMark       *WaterMark `json:"watermark"`
}

// WaterMark 水印
type WaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// BizDataCrypto 小程序数据解析
type BizDataCrypto struct {
	mp            *WXMP
	encryptedData string
	sessionKey    string
	iv            string
	data          []byte
}

// Decrypt 数据解密
func (b *BizDataCrypto) Decrypt() error {
	cipherText, err := base64.StdEncoding.DecodeString(b.encryptedData)

	if err != nil {
		return err
	}

	key, err := base64.StdEncoding.DecodeString(b.sessionKey)

	if err != nil {
		return err
	}

	iv, err := base64.StdEncoding.DecodeString(b.iv)

	if err != nil {
		return err
	}

	cbc := utils.NewAESCBCCrypto(key, iv)
	plainText, err := cbc.Decrypt(cipherText, utils.PKCS7)

	if err != nil {
		return err
	}

	b.data = plainText

	return nil
}

// GetUserData 获取用户数据（需先解密）
func (b *BizDataCrypto) GetUserData() (*UserData, error) {
	if len(b.data) == 0 {
		return nil, errors.New("gochat: empty data, check whether decrypted")
	}

	userData := new(UserData)

	if err := json.Unmarshal(b.data, &userData); err != nil {
		return nil, err
	}

	if userData.WaterMark.AppID != b.mp.appid {
		return nil, fmt.Errorf("gochat: wxmp user bizdata appid mismatch, want: %s, got: %s", b.mp.appid, userData.WaterMark.AppID)
	}

	return userData, nil
}

// GetPhoneData 获取用户手机号绑定数据（需先解密）
func (b *BizDataCrypto) GetPhoneData() (*PhoneData, error) {
	if len(b.data) == 0 {
		return nil, errors.New("gochat: empty data, check whether decrypted")
	}

	phoneData := new(PhoneData)

	if err := json.Unmarshal(b.data, &phoneData); err != nil {
		return nil, err
	}

	if phoneData.WaterMark.AppID != b.mp.appid {
		return nil, fmt.Errorf("gochat: wxmp phone bizdata appid mismatch, want: %s, got: %s", b.mp.appid, phoneData.WaterMark.AppID)
	}

	return phoneData, nil
}