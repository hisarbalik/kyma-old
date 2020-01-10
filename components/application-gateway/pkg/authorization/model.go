package authorization

// Credentials contains OAuth or BasicAuth configuration.
type Credentials struct {
	// OAuth is OAuth configuration.
	OAuth *OAuth
	// BasicAuth is BasicAuth configuration.
	BasicAuth *BasicAuth
	// CertificateGen is CertificateGen configuration.
	CertificateGen *CertificateGen
	// CSRFTokenEndpointURL (optional) to fetch CSRF token
	CSRFTokenEndpointURL string
}

// BasicAuth contains details of BasicAuth Auth configuration
type BasicAuth struct {
	// Username to use for authentication
	Username string
	// Password to use for authentication
	Password string
}

// OAuth contains details of OAuth configuration
type OAuth struct {
	// URL to OAuth token provider.
	URL string
	// ClientID to use for authorization.
	ClientID string
	// ClientSecret to use for authorization.
	ClientSecret string
	// RequestParameters will be used with request send by the Application Gateway.
	RequestParameters *RequestParameters
}

// CertificateGen details of CertificateGen configuration
type CertificateGen struct {
	// CommonName of the certificate
	CommonName string
	// Certificate generated by Application Registry
	Certificate []byte
	// PrivateKey generated by Application Registry
	PrivateKey []byte
}

// RequestParameters contains Headers and QueryParameters
type RequestParameters struct {
	Headers         *map[string][]string `json:"headers,omitempty"`
	QueryParameters *map[string][]string `json:"queryParameters,omitempty"`
}

func (rp *RequestParameters) unpack() (*map[string][]string, *map[string][]string) {
	if rp == nil {
		return nil, nil
	}
	return rp.Headers, rp.QueryParameters
}
