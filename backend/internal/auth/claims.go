package auth

type claims struct {
	Issuer          string `json:"iss"`
	AuthorizedParty string `json:"azp"`
	Audience        string `json:"aud"`
	Subject         string `json:"sub"`
	Domain          string `json:"hd"`
	Email           string `json:"email"`
	EmailVerified   bool   `json:"email_verified"`
	Hash            string `json:"at_hash"`
	IssuedAtSeconds int64  `json:"iat"`
	ExpirySeconds   int64  `json:"exp"`
}
