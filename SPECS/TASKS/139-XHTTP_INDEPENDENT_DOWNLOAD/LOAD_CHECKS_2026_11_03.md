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

The 300-second-per-mode run completed PASS, exit0, elapsed1801.674s. All six
H1/REALITY auto/packet-up/stream-up cases reached 12 concurrent workers. Total
214412 validated echoes /3512926208 payload bytes. Log SHA256:
9a05bfa123f67e9b6269592a90abbae3e2f881c35c0976ed0cf377bdb30ca751.
The compiled test commit was9b2b764; a6a9924 changed only this report while the
process was running. The per-case machine-readable receipt is alongside it.

Cloud regression41508575957 on a6a9924 completed SUCCESS. Verified downloaded
Windows artifact13173129363 SHA256
 d26a80372ea3bf687e6005ab6bbdd9948abd2bdb6c80b9c47ae8045be1c33552:
371 passing test/subtest events, no skips/failures, all three 300-second H2 cases,
no-tags build and canonical CLI build, three accepted and five correctly rejected
construction-only configurations. Linux race artifact13171323216 SHA256
4adf8ff1b44071cb90bb6f94e6b3a886ff4e5256258f36846982b6744f124188:
1550 pass events across five repeats, zero skips/failures, both packages passed.
No full untagged vet, physical runtime, license or owner-acceptance gate is promoted.

Scope: loopback transport frontends -> official Xray26.10.21 H2 VLESS echo backend.
Pinned xray.exe SHA256 c77a00eb84fb93ae044c2d771a1cabcc033b81e487253fbc35b12df6879c1e15.
This is not direct official Xray H1/REALITY ingress, Android TUN, physical network
migration, battery or linked Cronet source/license closure. PR remains draft/W;
owner acceptance and other original requirements remain open.
