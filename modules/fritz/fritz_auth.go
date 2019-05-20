package fritz

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// DoDigestAuthentication does a digist authentication request
func doDigestAuthentication(soapResponse *http.Response, soapRequest *SoapData) (string, error) {
	authHeader := soapResponse.Header.Get("WWW-Authenticate")

	if authHeader == "" {
		return "", errors.New("HTTP Header WWW-Authenticate is empty, cant do a Digist Authentication")
	}

	digestHeader := createDigistAuthenticationHeader(authHeader)
	digestHeader["username"] = string(soapRequest.Username)
	digestHeader["uri"] = soapRequest.URLPath
	digestHeader["method"] = "POST"

	header1 := getMD5(digestHeader["username"] + ":" + digestHeader["realm"] + ":" + string(soapRequest.Password))
	header2 := getMD5(digestHeader["method"] + ":" + digestHeader["uri"])
	nonceCount := "00000001"
	cnonce := getCnonce()
	response := getMD5(fmt.Sprintf("%s:%s:%v:%s:%s:%s", header1, digestHeader["nonce"], nonceCount, cnonce, digestHeader["qop"], header2))

	r := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%s", qop="%s", response="%s", algorithm="MD5"`, digestHeader["username"], digestHeader["realm"], digestHeader["nonce"], digestHeader["uri"], cnonce, nonceCount, digestHeader["qop"], response)

	return r, nil
}

func createDigistAuthenticationHeader(header string) map[string]string {
	digestHeader := map[string]string{}

	wanted := []string{"nonce", "realm", "qop"}
	response := strings.Split(header, ",")

	for _, r := range response {
		for _, w := range wanted {
			if strings.Contains(r, w) {
				digestHeader[w] = strings.Split(r, `"`)[1]
			}
		}
	}

	return digestHeader
}

func getMD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}

func getCnonce() string {
	b := make([]byte, 32)
	io.ReadFull(rand.Reader, b)

	return fmt.Sprintf("%x", b)[:64]
}
