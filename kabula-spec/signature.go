package spec

type Signature struct {
	Ref   string          `json:"ref,omitempty" xml:"ref,attr,omitempty"`
	Value *SignatureValue `json:"value,omitempty" xml:",omitempty"`
}

type SignatureValue struct {
	MD5    string        `json:"md5,omitempty" xml:"md5,omitempty"`
	Sha1   string        `json:"sha1,omitempty" xml:"sha1,omitempty"`
	Sha256 string        `json:"sha256,omitempty" xml:"sha256,omitempty"`
	Sha512 string        `json:"sha512,omitempty" xml:"sha512,omitempty"`
	GPG    *GPGSignature `json:"gpg,omitempty" xml:"gpg,omitempty"`
}

type GPGSignature struct {
	Fingerprint string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	Signature   string `json:"signature,omitempty" xml:"signature,omitempty"`
}
