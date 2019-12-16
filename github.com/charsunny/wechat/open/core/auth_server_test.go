package core

import (
	"encoding/json"

	"testing"
	"github.com/charsunny/wechat/open/oauth2"
)

func TestJsonUnmarshal(t *testing.T) {
	data := []byte(`{"authorization_info":{"authorizer_appid":"wx576e34f411251da3","authorizer_access_token":"17_qPvVolO_-OQfGH7Auq4Iwwmz4VRirtQG_IMWgGSsZRtt8sU7SZUSvER0K7ID00ztVL73_-d7D8vFRn-LcquyRVOuiFyNS2Fs-ALFI4b2i8r5bkNpq7mKnJLD4HVifOrcrKvRwVyWRaJqMWe2LRWgAMDXDI","expires_in":7200,"authorizer_refresh_token":"refreshtoken@@@acCM9b7IlUlyTkzACa0ZbrLU8KCcYM2BpxHviQmVzvQ"}}`)

	var result struct {
		AuthorizationInfo oauth2.AuthorizationInfo `json:"authorization_info"`
	}
	result1 := new(struct {
		AuthorizationInfo oauth2.AuthorizationInfo `json:"authorization_info"`
	})
	result1.AuthorizationInfo = oauth2.AuthorizationInfo{
		AppId:"xxx",
		AccessToken:"x",
		RefreshToken:"xxxx",
	}
	s, err := json.Marshal(result1)
	t.Logf("%s, %v", s, err)
	if err := json.Unmarshal(data, &result); err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", result.AuthorizationInfo.AppId)
	t.Log(result)
	t.Errorf("%s", result.AuthorizationInfo.AppId)
}
