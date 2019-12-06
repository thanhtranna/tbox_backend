package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"tbox_backend/src/entity"
)

type ITwilioHelper interface {
	SendOTP(string, string) error
}

type twilioHelper struct {
	config entity.TwilioConfig
}

func NewTwilioHelper(config entity.TwilioConfig) ITwilioHelper {
	return &twilioHelper{
		config: config,
	}
}

func (t *twilioHelper) SendOTP(OTP string, phoneNumber string) error {
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.config.AccountSID)
	// Build out the data for our message
	value := url.Values{}
	value.Set("To", phoneNumber)
	value.Set("From", t.config.PhoneNumber)
	value.Set("Body", OTP)
	rb := *strings.NewReader(value.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(t.config.AccountSID, t.config.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("something went wrong")
}
