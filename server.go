package main

import (
	"fmt"
	"net/http"
	"os"
)

const ENCRYPTED_KEY = `
-----BEGIN CMS-----
MIIC0wYJKoZIhvcNAQcCoIICxDCCAsACAQExCTAHBgUrDgMCGjCCAawGCSqGSIb3
DQEHAaCCAZ0EggGZc3NoLXJzYUFBQUFCM056YUMxeWMyRUFBQUFEQVFBQkFBQUJB
UURZaWpxOUs0VnF1TTNjOThPQ0RNV0docTlodzU0UDJ4Qm9XRUgxbGlESlp4eis2
WVdBeHJpSzVrdnpFS0JHRW04dU42UkNESEU5VHRFcmVPYXE4WWNNRnhtVU5FSm5u
QVMxZWFCRmNhTFJ0Y1VZRDMvSjY2bSswNFA5Nll0SkFYb2p2T2xMdHh5d0d3RnE0
VVdmR3ZMUjZOT2ExK3pJa2VMSXJuMlV1ZGhpMEhUYU80eEN5ZE43SFVUKzArblY4
bHZDQWpGTVBOSnN5cUVIdE8yUnVSamJwb1huTk8yQ0pBd1NaZW9RTG9jR1RuVTZB
akNVZFZITmZmcWVkMDBNdm85anNZbEdUczIzRWZIOGxhWEpkUTdNTHVGREYvenUr
RXpjdFArSHRMZi9vdFhzOGpGU05UVUdzY0lzbU9teFM4T2RJNVJBNmxrc1k0UXd6
YXduRlRBWmRhbkBEYW5zLU1hY0Jvb2stUHJvLmxvY2FsDQoNCjGB/zCB/AIBATBc
MFcxCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVVbnNldDEOMAwGA1UEBwwFVW5zZXQx
DjAMBgNVBAoMBVVuc2V0MRgwFgYDVQQDDA93d3cuZXhhbXBsZS5jb20CAQEwBwYF
Kw4DAhowDQYJKoZIhvcNAQEBBQAEgYCdfbUIyX/5wgWAMRxWb03ZOg17BXmkuTYm
foM1dgj3AUXjDR6iMO3E2/rgOJPS/+ZXAkyTKTk4c0hnC5u1JBKRUPOx9riavfYr
VDenf9HHlwClQcjE2m9snhvx2n9gyaYlPr16cHcTdydKPCTdaugB7KvQQvxTsyap
qS1Gx9N3OQ==
-----END CMS-----
`

func main() {
	http.HandleFunc("/", returnKey)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func returnKey(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, ENCRYPTED_KEY)
}
