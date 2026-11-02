//go:build windows && with_xhttp && with_utls

package v2rayxhttp

import "testing"

// Both directions have independent transport identities and XMUX limits.
// The VLESS echo backend is official Xray over H2 behind the two frontends;
// these tests do not claim direct official Xray H1/REALITY ingress or Android TUN.
func TestIndependentDownloadOfficialXrayH1Concurrent(t *testing.T) {
	runIndependentDownloadOfficialXrayConcurrent(t, 1)
}

func TestIndependentDownloadOfficialXrayRealityConcurrent(t *testing.T) {
	runIndependentDownloadOfficialXrayConcurrent(t, 0)
}
