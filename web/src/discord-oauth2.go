package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/*
OAuth2 Scopes
* guilds.join => '/guilds/{guild.id}/members/{user.id}'
*/

const (
	// API
	DISCORD_ENDPOINT_API            = "https://discord.com/api/v8" //' v10'?
	DISCORD_ENDPOINT_API_USERS_ME   = DISCORD_ENDPOINT_API + "/users/@me"
	DISCORD_ENDPOINT_API_GUILD_JOIN = DISCORD_ENDPOINT_API + "'/guilds/{guild.id}/members/{user.id}"

	// OAuth2
	DISCORD_ENDPOINT_OAUTH2       = DISCORD_ENDPOINT_API + "/oauth2"
	DISCORD_ENDPOINT_OAUTH2_TOKEN = DISCORD_ENDPOINT_OAUTH2 + "/token"
	DISCORD_ENDPOINT_OAUTH2_AUTH  = DISCORD_ENDPOINT_OAUTH2 + "/authorize"
)

/*
HEADERS:
	'Content-Type': 'application/x-www-form-urlencoded'
	'User-Agent': 'XXX'
*/

// 'Content-Type': 'application/json'
type DiscordAPI_Req_Users struct {
	Authorization string // ? HEADER ODER SO ("Bot TOKEN")
}

// ? was noch
type DiscordAPI_Resp_Users struct {
	Username      string
	Discriminator string
	ID            string
	Locale        string
	Refresh_Token string
}

// url=f"{API_ENDPOINT}/guilds/{str(GUILD_ID)}/members/{user_id}",
type DiscordAPI_Req_Guild_Join struct {
	Access_Token string
	Roles        []string // role_ID
}
type DiscordAPI_Resp_Guild_Join struct {
	//?
	// status code?
}

// Access Token Exchange
type DiscordOAuth2_Req_Token struct {
	Client_ID     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Grant_Type    string `json:"grant_type"` // 'authorization_code'
	Code          string `json:"code"`       // Code from Discord callback
	Redirect_URI  string `json:"redirect_uri"`
}

// Refresh Token Exchange Example
type DiscordOAuth2_Req_Refresh_Token struct {
	Client_ID     string
	Client_Secret string
	Grant_Type    string // 'refresh_token'
	Refresh_Token string // stored refresh token
}

// Access Token Response
type DiscordOAuth2_Resp_Token struct {
	Access_Token  string
	Token_Type    string // 'Bearer'
	Expires_In    uint   // time?
	Refresh_Token string // New Refresh Token
	Scope         string //[]string? (identify,guilds.join)
}

type DiscordOAuth2_Req_Auth struct {
	Client_ID        string
	Response_Type    string //'code'
	Scope            string //[]string? (identify,guilds.join) && getrennt mit '%20', URL encode?
	State            string
	Redirect_URI     string
	Prompt           string // 'none'
	Integration_Type uint   // '0' || '1'
}

type DiscordOAuth2_Resp_Auth struct {
	Code  string
	State string
}

type DiscordOAuth2Client struct {
	clientID     string
	clientSecret string
	guildID      string
	redirectURI  string
	//	stateSigningKey []byte ???

	onSuccess OnSuccessFunc
	onError   OnErrorFunc

	client *http.Client
}

type OnSuccessFunc func() error // ctx,route?,resp
type OnErrorFunc func() error

func NewDiscordOAuth2(client_id, client_secret, guild_id, redirURL string, OnSuccess OnSuccessFunc, OnError OnErrorFunc) (*DiscordOAuth2Client, error) {
	if OnSuccess == nil {
		OnSuccess = func() error { return nil }
	}

	if OnError == nil {
		OnError = func() error { return nil }
	}

	return &DiscordOAuth2Client{
		clientID:     client_id,
		clientSecret: client_secret,
		guildID:      guild_id,
		redirectURI:  url.QueryEscape(redirURL),
		onSuccess:    OnSuccess,
		onError:      OnError,
		client:       &http.Client{},
	}, nil
}

// /authorize?response_type=code&client_id=157730590492196864&scope=identify%20guilds.join&state=15773059ghq9183habn&redirect_uri=https%3A%2F%2Fnicememe.website&prompt=consent&integration_type=0
// Redir: /callback?code=XXX&state=XXX
func (oauth2 *DiscordOAuth2Client) HandlerAuthorize(w http.ResponseWriter, r *http.Request) {
	var url = fmt.Sprintf(
		DISCORD_ENDPOINT_OAUTH2_AUTH+"?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		oauth2.clientID,
		oauth2.redirectURI,
		"SCOPE",
	)

	// redir to discord auth
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func (oauth2 *DiscordOAuth2Client) HandlerAuthorizeCallback(w http.ResponseWriter, r *http.Request) {
	// wir bekommen: GET /callback?code=XXX&state=XXX
	//r.URL.Query()
	//r.Form
	//r.FormValue()
	// 'code', 'state'

	// code zu token machen (exchange_code)

	// user_info?

	// add user to guild
}

func (oauth2 *DiscordOAuth2Client) exchangeCode_for_accessToken(code string) (*DiscordOAuth2_Resp_Token, error) {
	// JSON Body
	jsonBody, err := json.Marshal(
		&DiscordOAuth2_Req_Token{
			Client_ID:     oauth2.clientID,
			Client_secret: oauth2.clientSecret,
			Grant_Type:    "authorization_code",
			Code:          code,
			Redirect_URI:  oauth2.redirectURI,
		},
	)
	if err != nil {
		return nil, err
	}

	// HTTP POST
	req, err := http.NewRequest(
		"POST",
		DISCORD_ENDPOINT_OAUTH2_TOKEN,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	// Header
	req.Header.Set(
		"Content-Type", "application/x-www-form-urlencoded",
	)

	// Do Request
	resp, err := oauth2.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// valid Response ?
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status-Code != 200")
	}

	var accessTokenResp DiscordOAuth2_Resp_Token
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenResp); err != nil {
		return nil, err
	}

	return &accessTokenResp, nil
}

func (oauth2 *DiscordOAuth2Client) exchangeRefreshToken_for_AccessToken(refresh_token string) (*DiscordOAuth2_Resp_Token, error) {
	// JSON Body
	jsonBody, err := json.Marshal(
		&DiscordOAuth2_Req_Refresh_Token{
			Client_ID:     oauth2.clientID,
			Client_Secret: oauth2.clientSecret,
			Grant_Type:    "refresh_token",
			Refresh_Token: refresh_token,
		},
	)
	if err != nil {
		return nil, err
	}

	// HTTP POST
	req, err := http.NewRequest(
		"POST",
		DISCORD_ENDPOINT_OAUTH2_TOKEN,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	// Header
	req.Header.Set(
		"Content-Type", "application/x-www-form-urlencoded",
	)

	// Do Request
	resp, err := oauth2.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// valid Response ?
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status-Code != 200")
	}

	var accessTokenResp DiscordOAuth2_Resp_Token
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenResp); err != nil {
		return nil, err
	}

	return &accessTokenResp, nil
}

/* REFRESH_TOKEN (:REFRESH_TOKEN)
=> f'{OAuth2.api_endpoint}/oauth2/token'
	POST:
		HEADER:
			{
            'Content-Type': 'application/x-www-form-urlencoded'
            }
		DATA:
			{
            'client_id': OAuth2.client_id,
            'client_secret': OAuth2.client_secret,
            'grant_type': 'refresh_token',
            'refresh_token': refresh_token
            }

	RESPONSE:
		{
            "access_token": r.json().get("access_token"),
            "refresh_token": r.json().get("refresh_token")
        }
*/

/* GET USER INFO (:ACCESS_TOKEN)
=> OAuth2.api_endpoint + "/users/@me"

	GET:
		HEADERS:
			{
            "Authorization": f"Bearer {access_token}"
            }

	RESPONSE:
		{
            "username": user_object.get("username") + "#" + user_object.get("discriminator"),
            "id": user_object.get("id"),
            "email": user_object.get("email"),
            "locale": user_object.get("locale")
    	}
*/

/*
@app.route("/callback", methods=["GET"])
def callback():
    code = request.args.get("code")
    tokens = OAuth2.exchange_code(code)
    user_object = OAuth2.get_user_info(tokens.get("access_token"))

    to_save_data = {
        "username": user_object.get("username"),
        "id": user_object.get("id"),
        "email": user_object.get("email"),
        "locale": user_object.get("locale"),
        "refresh_token": tokens.get("refresh_token")
        }

    # save(to_save_data)

    return render_template("callback.html")
*/
