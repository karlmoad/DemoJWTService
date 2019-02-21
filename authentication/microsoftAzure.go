package authentication

import (
	"errors"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	ErrorKeyNotFound = errors.New("key not found")
	ErrorMetadataFailure = errors.New("failure retrieving key metadata")
	ErrorKeyData = errors.New("failure getting keys")
)



type microsoftKeySource struct {
	meta 		tokenMeta
	keys 		tokenKeys
	expires 		time.Time
}

func NewMicrosoftKeySource() (*microsoftKeySource, error){
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://login.microsoftonline.com/common/.well-known/openid-configuration")
	if err != nil {
		log.Println(err)
		return nil, ErrorMetadataFailure
	}
	defer resp.Body.Close()

	var metadata tokenMeta
	err = json.NewDecoder(resp.Body).Decode(&metadata)
	if err != nil {
		log.Println(err)
		return nil,ErrorMetadataFailure
	}

	resp2, err := client.Get(metadata.JwksURI)
	if err != nil {
		log.Println(err)
		return nil, ErrorKeyData
	}
	defer resp2.Body.Close()

	var keydata tokenKeys
	err = json.NewDecoder(resp2.Body).Decode(&keydata)
	if err != nil {
		log.Println(err)
		return nil, ErrorKeyData
	}

	// Microsoft suggests to check key data for changes every 24 hours
	// to be safe we are cutting that in half with a 12 hour expiration time
	msKS := &microsoftKeySource{meta:metadata, keys:keydata, expires:time.Now().Add(time.Hour * 12)}
	return msKS, nil
}

func(m *microsoftKeySource) IsExpired() bool {
	return time.Since(m.expires).Hours() > 0
}

func(m *microsoftKeySource) GetKidEntries() []string {
	var out []string
	for _, entry := range m.keys.Keys {
		out = append(out, entry.Kid)
	}
	return out
}

func(m *microsoftKeySource) GetKeys(kid string) ([]string, error) {
	for _ , entry := range m.keys.Keys {
		if strings.Compare(kid, entry.Kid) == 0 {
			return entry.X5C, nil
		}
	}
	return nil, ErrorKeyNotFound
}

type Token struct {
	Aud        string   `json:"aud"`
	Iss        string   `json:"iss"`
	Iat        int      `json:"iat"`
	Nbf        int      `json:"nbf"`
	Exp        int      `json:"exp"`
	Aio        string   `json:"aio"`
	Amr        []string `json:"amr"`
	FamilyName string   `json:"family_name"`
	GivenName  string   `json:"given_name"`
	InCorp     string   `json:"in_corp"`
	Ipaddr     string   `json:"ipaddr"`
	Name       string   `json:"name"`
	Nonce      string   `json:"nonce"`
	Oid        string   `json:"oid"`
	OnpremSid  string   `json:"onprem_sid"`
	Sub        string   `json:"sub"`
	Tid        string   `json:"tid"`
	UniqueName string   `json:"unique_name"`
	Upn        string   `json:"upn"`
	Uti        string   `json:"uti"`
	Ver        string   `json:"ver"`
}

type tokenMeta struct {
	AuthorizationEndpoint             string      `json:"authorization_endpoint"`
	TokenEndpoint                     string      `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string    `json:"token_endpoint_auth_methods_supported"`
	JwksURI                           string      `json:"jwks_uri"`
	ResponseModesSupported            []string    `json:"response_modes_supported"`
	SubjectTypesSupported             []string    `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string    `json:"id_token_signing_alg_values_supported"`
	HTTPLogoutSupported               bool        `json:"http_logout_supported"`
	FrontchannelLogoutSupported       bool        `json:"frontchannel_logout_supported"`
	EndSessionEndpoint                string      `json:"end_session_endpoint"`
	ResponseTypesSupported            []string    `json:"response_types_supported"`
	ScopesSupported                   []string    `json:"scopes_supported"`
	Issuer                            string      `json:"issuer"`
	ClaimsSupported                   []string    `json:"claims_supported"`
	MicrosoftMultiRefreshToken        bool        `json:"microsoft_multi_refresh_token"`
	CheckSessionIframe                string      `json:"check_session_iframe"`
	UserinfoEndpoint                  string      `json:"userinfo_endpoint"`
	TenantRegionScope                 interface{} `json:"tenant_region_scope"`
	CloudInstanceName                 string      `json:"cloud_instance_name"`
	CloudGraphHostName                string      `json:"cloud_graph_host_name"`
	MsgraphHost                       string      `json:"msgraph_host"`
	RbacURL                           string      `json:"rbac_url"`
}


type tokenKeys struct {
	Keys []struct {
		Kty string   `json:"kty"`
		Use string   `json:"use"`
		Kid string   `json:"kid"`
		X5T string   `json:"x5t"`
		N   string   `json:"n"`
		E   string   `json:"e"`
		X5C []string `json:"x5c"`
	} `json:"keys"`
}

