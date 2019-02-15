package fritz

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// DoDigestAuthentication does a digist authentication request
func DoDigestAuthentication(fSR *SoapRequest) {

	headerDigsetAuth := fSR.soapResponse.Header.Get("WWW-Authenticate")

	if headerDigsetAuth != "" {

		digestHeader := internalCreateDigestAuthenticationHeader(headerDigsetAuth)

		digestHeader["username"] = fSR.Username
		digestHeader["password"] = fSR.Password
		digestHeader["uri"] = fSR.URLPath
		digestHeader["method"] = "POST"

		header1 := internalGetMD5(digestHeader["username"] + ":" + digestHeader["realm"] + ":" + digestHeader["password"])
		header2 := internalGetMD5(digestHeader["method"] + ":" + digestHeader["uri"])
		nonceCount := "00000001"
		cnonce := internalGetCnonce()
		response := internalGetMD5(fmt.Sprintf("%s:%s:%v:%s:%s:%s", header1, digestHeader["nonce"], nonceCount, cnonce, digestHeader["qop"], header2))

		fSR.soapRequest.Header.Set("Authorization", fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%s", qop="%s", response="%s", algorithm="MD5"`, digestHeader["username"], digestHeader["realm"], digestHeader["nonce"], digestHeader["uri"], cnonce, nonceCount, digestHeader["qop"], response))
	}
}

func internalCreateDigestAuthenticationHeader(digsetHeader string) map[string]string {
	result := map[string]string{}

	wanted := []string{"nonce", "realm", "qop"}
	response := strings.Split(digsetHeader, ",")

	for _, r := range response {
		for _, w := range wanted {
			if strings.Contains(r, w) {
				result[w] = strings.Split(r, `"`)[1]
			}
		}
	}

	return result
}

func internalGetMD5(str string) string {
	hash := md5.New()

	hash.Write([]byte(str))

	return hex.EncodeToString(hash.Sum(nil))
}

func internalGetCnonce() string {
	b := make([]byte, 32)

	io.ReadFull(rand.Reader, b)

	return fmt.Sprintf("%x", b)[:64]
}
