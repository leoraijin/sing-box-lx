# Additional persistent transport checks — 3 November 2026

Base: fe2e2f52ae0b360b95b6cd82900208f1d393f357.
Test implementation: 9b2b764a1273316ee05ed888d870b562038df3bb.

Changes are test-only. The existing H2 load test shares a transport-version-aware
helper with new H1/H1 and REALITY/REALITY tests. Parsed request HTTP versions are
checked. Each direction retains separate endpoint, XMUX limits and identities.
The REALITY fixture's 10-second deadline now bounds only the handshake, rather
than killing an otherwise healthy persistent session ten seconds after accept.
No production transport, dependencies, module pins or build tags were changed.

Verified locally using pinned Go1.25.5, offline readonly modules, Windows:
- Existing independent endpoint matrix passed against unmodified base.
- New H1 and REALITY tests passed at 5 seconds per mode, 12 concurrent sessions,
  auto/packet-up/stream-up. First H1 attempt exposed a test-only nil TLS assumption;
  helper now handles the existing plaintext H1 endpoint as other matrix tests do.
- go vet with with_xhttp,with_quic,with_utls,with_reality passed for
  ./transport/v2rayxhttp and ./option. gofmt/diff checks passed.

A 300-second-per-mode run on test commit 9b2b764 is running. H1 auto and packet-up
have completed with validated persistent echo; full H1/REALITY soak acceptance is
NOT claimed until the process exits and all six cases are inspected. No checklist
item is promoted on the basis of a still-running process.

Scope: loopback transport frontends -> official Xray26.10.21 H2 VLESS echo backend.
Pinned xray.exe SHA256 c77a00eb84fb93ae044c2d771a1cabcc033b81e487253fbc35b12df6879c1e15.
This is not direct official Xray H1/REALITY ingress, Android TUN, physical network
migration, battery or linked Cronet source/license closure. PR remains draft/W;
owner acceptance and other original requirements remain open.
