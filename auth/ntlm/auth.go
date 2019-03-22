package ntlm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	ntlmssp "github.com/Azure/go-ntlmssp"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/cpass"
)

// AuthCnfg : auth config structure
type AuthCnfg struct {
	SiteURL  string `json:"siteUrl"`
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`

	masterKey string
}

// ReadConfig : reads private config with auth options
func (c *AuthCnfg) ReadConfig(privateFile string) error {
	jsonFile, err := os.Open(privateFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &c)

	crypt := cpass.Cpass(c.masterKey)
	pass, err := crypt.Decode(c.Password)
	if err == nil {
		c.Password = pass
	}

	if c.Domain != "" && !strings.Contains(c.Username, "\\") && !strings.Contains(c.Username, "@") {
		c.Username = c.Domain + "\\" + c.Username
	}

	return nil
}

// SetMasterkey : defines custom masterkey
func (c *AuthCnfg) SetMasterkey(masterKey string) {
	c.masterKey = masterKey
}

// GetAuth : authenticates, receives access token
func (c *AuthCnfg) GetAuth() (string, error) {
	return GetAuth(c)
}

// GetSiteURL : gets siteURL
func (c *AuthCnfg) GetSiteURL() string {
	return c.SiteURL
}

// GetStrategy : gets auth strategy name
func (c *AuthCnfg) GetStrategy() string {
	return "ntlm"
}

// SetAuth : authenticate request
func (c *AuthCnfg) SetAuth(req *http.Request, httpClient *gosip.SPClient) error {
	// NTML + Negotiation
	httpClient.Transport = ntlmssp.Negotiator{
		RoundTripper: &http.Transport{},
	}
	req.SetBasicAuth(c.Username, c.Password)
	return nil
}