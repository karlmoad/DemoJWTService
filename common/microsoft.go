package common

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
	meta 		microsoftTokenMeta
	keys 		microsoftTokenKeys
}

func NewMicrosoftKeySource() (*microsoftKeySource, error){
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://login.microsoftonline.com/common/.well-known/openid-configuration")
	if err != nil {
		log.Println(err)
		return nil, ErrorMetadataFailure
	}
	defer resp.Body.Close()

	var metadata microsoftTokenMeta
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

	var keydata microsoftTokenKeys
	err = json.NewDecoder(resp2.Body).Decode(&keydata)
	if err != nil {
		log.Println(err)
		return nil, ErrorKeyData
	}

	msKS := &microsoftKeySource{meta:metadata, keys:keydata}
	return msKS, nil
}



func(m *microsoftKeySource) GetKeys(kid string) ([]string, error) {
	for i , entry := range m.keys.Keys {
		if strings.Compare(kid, entry.Kid) == 0 {
			return entry.X5C, nil
		}
	}
	return nil, ErrorKeyNotFound
}

type microsoftTokenMeta struct {
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


type microsoftTokenKeys struct {
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


