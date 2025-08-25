/*
? 'User-Agent': 'XXX'
*/
package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// TODO: discordgo kann auch bearer

const (
	// API
	DISCORD_ENDPOINT_API          = "https://discord.com/api/v8" //' v10'?
	DISCORD_ENDPOINT_API_USERS_ME = DISCORD_ENDPOINT_API + "/users/@me"
	DISCORD_ENDPOINT_API_GUILD    = DISCORD_ENDPOINT_API + "'/guilds"

	// OAuth2
	DISCORD_ENDPOINT_OAUTH2       = DISCORD_ENDPOINT_API + "/oauth2"
	DISCORD_ENDPOINT_OAUTH2_TOKEN = DISCORD_ENDPOINT_OAUTH2 + "/token"
	DISCORD_ENDPOINT_OAUTH2_AUTH  = DISCORD_ENDPOINT_OAUTH2 + "/authorize"
)

// Error Types
var (
	ErrHTTPRespStatusCode              = errors.New("response status-code is not 200")
	ErrHTTPRespStatusCodeNotAsExpected = errors.New("response status-code is not 201 or 204 as expected")
)

// Resp: status-code: [201, 204]
type DiscordAPI_Req_Guild_Join struct {
	Access_Token string   `json:"access_token"`
	Nick         string   `json:"nick"`
	Roles        []string `json:"roles"`
	Mute         bool     `json:"mute"`
	Deaf         bool     `json:"deaf"`
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
// Resp: DiscordOAuth2_Resp_Token
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

type DiscordOAuth2Client struct {
	Token        string
	clientID     string
	clientSecret string
	guildID      string
	scope        string
	redirectURI  string

	onSuccess OnSuccessFunc
	onError   OnErrorFunc

	client *http.Client
}

type OnSuccessFunc func(*DiscordOAuth2_Resp_Token) error // ctx,route?,resp
type OnErrorFunc func(error) error

func NewDiscordOAuth2(token, client_id, client_secret, guild_id, redirURL string, scopes []string, onSuccess OnSuccessFunc, onError OnErrorFunc) *DiscordOAuth2Client {
	if onSuccess == nil {
		onSuccess = func(*DiscordOAuth2_Resp_Token) error { return nil }
	}

	if onError == nil {
		onError = func(error) error { return nil }
	}

	// &scopeBuilder=identify%20guilds.join
	var scopeBuilder = strings.Builder{}
	for _, s := range scopes {
		scopeBuilder.WriteString(s)
		scopeBuilder.WriteRune(' ')
	}

	return &DiscordOAuth2Client{
		Token:        token,
		clientID:     client_id,
		clientSecret: client_secret,
		guildID:      guild_id,
		scope:        url.PathEscape(scopeBuilder.String()),
		redirectURI:  url.QueryEscape(redirURL),
		onSuccess:    onSuccess,
		onError:      onError,
		client:       &http.Client{},
	}
}

// redirects to discord's oauth2/authorize
func (oauth2 *DiscordOAuth2Client) HandlerAuthorize(w http.ResponseWriter, r *http.Request) {
	var url = strings.Builder{}
	url.WriteString(DISCORD_ENDPOINT_OAUTH2_AUTH)
	url.WriteString(fmt.Sprintf(
		"?response_type=code&client_id=%s&scope=%s&redirect_uri=%s",
		oauth2.clientID,
		oauth2.scope,
		oauth2.redirectURI,
	))

	http.Redirect(w, r, url.String(), http.StatusPermanentRedirect)
}

// handle discord callback and get Access-Token
func (oauth2 *DiscordOAuth2Client) HandlerAuthorizeCallback(w http.ResponseWriter, r *http.Request) error {
	// wir bekommen: GET /callback?code=XXX&state=XXX
	var code string

	//!
	//r.URL.Query()
	//r.Form
	//r.FormValue()
	// 'code', 'state'

	// Get Access Token
	access_token, err := oauth2.ExchangeCode_for_accessToken(code)
	if err != nil {
		return oauth2.onError(err)
	}

	return oauth2.onSuccess(access_token)
}

/*
> Discord Documentation
* https://discord.com:2053/developers/docs/topics/oauth2#authorization-code-grant-access-token-response
*/
func (oauth2 *DiscordOAuth2Client) ExchangeCode_for_accessToken(code string) (*DiscordOAuth2_Resp_Token, error) {
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
		return nil, ErrHTTPRespStatusCode
	}

	var accessTokenResp DiscordOAuth2_Resp_Token
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenResp); err != nil {
		return nil, err
	}

	return &accessTokenResp, nil
}

/*
> Discord Documentation
* https://discord.com:2053/developers/docs/topics/oauth2#authorization-code-grant-refresh-token-exchange-example
*/
func (oauth2 *DiscordOAuth2Client) ExchangeRefreshToken_for_AccessToken(refresh_token string) (*DiscordOAuth2_Resp_Token, error) {
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
		return nil, ErrHTTPRespStatusCode
	}

	var accessTokenResp DiscordOAuth2_Resp_Token
	if err := json.NewDecoder(resp.Body).Decode(&accessTokenResp); err != nil {
		return nil, err
	}

	return &accessTokenResp, nil
}

/*
> Discord Documentation
* https://discord.com:2053/developers/docs/resources/user#get-current-user
Returns the user object of the requester's account.
For OAuth2, this requires the identify scope,
which will return the object without an email,
and optionally the email scope,
which returns the object with an email if the user has one.
*/
func (oauth2 *DiscordOAuth2Client) GetUser(access_token string) (*discordgo.User, error) {
	// HTTP GET
	req, err := http.NewRequest(
		"GET",
		DISCORD_ENDPOINT_API_USERS_ME,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Header
	req.Header.Set(
		"Authorization", fmt.Sprintf("Bearer %s", access_token),
	)

	// Do Request
	resp, err := oauth2.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// valid Response?
	if resp.StatusCode != http.StatusOK {
		return nil, ErrHTTPRespStatusCode
	}

	var discordUser discordgo.User
	if err := json.NewDecoder(resp.Body).Decode(&discordUser); err != nil {
		return nil, err
	}

	return &discordUser, nil
}

/*
> Discord Documentation
* https://discord.com:2053/developers/docs/resources/guild#add-guild-member
Adds a user to the guild,
provided you have a valid oauth2 access token
for the user with the guilds.join scope.
Returns a 201 Created with the guild member as the body,
or 204 No Content if the user is already a member of the guild.
*/
func (oauth2 *DiscordOAuth2Client) AddGuildMember(user_id, member_role_id string) error {
	var url = fmt.Sprintf(
		DISCORD_ENDPOINT_API_GUILD+"/%s/members/%s",
		oauth2.guildID,
		user_id,
	)

	// JSON Body
	jsonBody, err := json.Marshal(
		&DiscordAPI_Req_Guild_Join{
			Access_Token: oauth2.Token,
			Roles:        []string{member_role_id},
		},
	)
	if err != nil {
		return err
	}

	// HTTP PUT
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return err
	}

	// Header
	req.Header.Set(
		"Authorization", oauth2.Token,
	)
	req.Header.Set(
		"Content-Type", "application/json",
	)

	// Do Request
	resp, err := oauth2.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// valid Response?
	switch resp.StatusCode {
	case http.StatusCreated:
		// user added
		return nil

	case http.StatusNoContent:
		// user already added
		return nil

	default:
		return ErrHTTPRespStatusCodeNotAsExpected
	}
}
