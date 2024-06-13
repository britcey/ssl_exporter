package prober

import (
	"testing"
)

func TestDecodeCertificates(t *testing.T) {
	data := []byte(`
-----BEGIN CERTIFICATE-----
MIIFszCCA5ugAwIBAgIUdpzowWDU/AI7QBhLSRB9DPqpWvcwDQYJKoZIhvcNAQEL
BQAwaTELMAkGA1UEBhMCWFgxEjAQBgNVBAgMCVN0YXRlTmFtZTERMA8GA1UEBwwI
Q2l0eU5hbWUxFDASBgNVBAoMC0NvbXBhbnlOYW1lMQwwCgYDVQQLDANGb28xDzAN
BgNVBAMMBkZvb2JhcjAeFw0yNDA0MjgxNjIzMzNaFw0zNDA0MjYxNjIzMzNaMGkx
CzAJBgNVBAYTAlhYMRIwEAYDVQQIDAlTdGF0ZU5hbWUxETAPBgNVBAcMCENpdHlO
YW1lMRQwEgYDVQQKDAtDb21wYW55TmFtZTEMMAoGA1UECwwDRm9vMQ8wDQYDVQQD
DAZGb29iYXIwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCmxZBW/Ays
VBxt7jJQeTrPdLQpUdxnGVXOa4M54FHWA2DwwmZ2DZth5Eioq1wC9WCrByWkd8px
mvU0XDUT5ESceEKcwmDhKgYHAcJ4qXEyk1jYuBy6zw95cmV2BiTf0Xoo/8JxiQR4
YBd7Tbm52eV5Hw5oaqgVEdaOCVMnO8S57AuQGfeC5AO18ty0cZ7mKsXQje2celMH
QipujXrhRwdLBgu6FISTuS0XtqiuJnp+vllMjiTMF/uCmJUTUCpmayWSpM1FRpKe
lM8B9666uuEmVJ4V5gzy4Oe0i5Apfh4qXX6pj+Y/oeOfXRc3NfYqcIJ2hm82ghJL
5Kt6FZ9fkQ6pyk7nXMAOaf8WX+JkhqaWvlzTme8Z+6DBXLPfyyoi3VR2kG3lRida
qfyh2EPvEXip810s4f8EOHi+sehmjWLbsgn2HAHQmE7zYTEMp2Xsh0M7lGfUiMfP
P1RU/RNDkZK59yXsG9RmoDMD03qI8M4990TL7BW0FQZvBg9Rfr3KgFpWngsn84mN
l6/MKyc5X1e+RGyYJPbi0S1j+fhBeNzFqQPFZZUnWIPTxVdqF91SgY9EwE6MHuD6
t7G+eANWWQkolvwACMT+0GqiRMVHVWeIqF2SQ7Dx2tiNXjE9rSX3lMpMNt4PvWcM
PGuao3FMqXts0j2L8FSljcVU/hiLOGC2mwIDAQABo1MwUTAdBgNVHQ4EFgQU+231
i0rY0CwQ0JwT2EUX4cq81LIwHwYDVR0jBBgwFoAU+231i0rY0CwQ0JwT2EUX4cq8
1LIwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAgEAVRqUVxovcZ9C
BYUUhvyT8BMAC6PS8U59BoDhnbrp1JDiPDbMnWlRp+450kZlEkSdN9PRLGWoGi84
QqT7RIry/jA4z62B6N+IDDDcuWWstDTTC/gRA46eeHGRnIz7GJeseqrxb9z8KoO3
/c8Pdb4lu/cfG5/m2X8hkYzjLZrIQ8qz53nUAag4uael19uy3HPQ3EEr936y+vQW
rH8itgQ+5kOTr9d3ihcTSwTJGTyWq7j3T0xPu9tCzoPragpEDjoUCqCjU7uv6Tdq
UhaQ8vneimSf9VdbDfEuEU9S2Xbdg6e+BsdV9uMeWkhtD1WgtSLHcliYktwnh/kh
r4vQG86xn+LDxTrdTssjt+UmJnXzRiGOEZCDz7kBchYLiWSQn9tqcn28KK2/YobD
AghlR24hUL1uslJOjLFc4BwLmlv/4Iy+b/3iQWY+yoban/OLZo6gCZx/4Nvg5R9e
tcuxgn5Jhm2T/e8REucXQfiqKgxQVibFmqXeLH3Yj7ussA5t/ozo4VuypXUnnil0
hB+PQsViaHFId1LyBKFwMoA6JzlNle6cbWVtXJswffkjk0UUr26SLHMV87lrk2kn
IQ7ROZcT8K7gzzTxuiC2g6npmG1fDfgmgJeS7IqgiFcRTKa0hj6bvc/WoHoCpLNs
v5KXxrGlqibeNjyc3Wng6S0Kpg6YNqU=
-----END CERTIFICATE-----
-----BEGIN TRUSTED CERTIFICATE-----
MIIF7zCCA9egAwIBAgIUZXsnB0gxoPSFaUjnQeLwTW+sqbYwDQYJKoZIhvcNAQEL
BQAwgYYxCzAJBgNVBAYTAlhYMRIwEAYDVQQIDAlTdGF0ZU5hbWUxETAPBgNVBAcM
CENpdHlOYW1lMRQwEgYDVQQKDAtDb21wYW55TmFtZTEbMBkGA1UECwwSQ29tcGFu
eVNlY3Rpb25OYW1lMR0wGwYDVQQDDBRDb21tb25OYW1lT3JIb3N0bmFtZTAeFw0y
NDA0MjgxNjE3MzZaFw0zNDA0MjYxNjE3MzZaMIGGMQswCQYDVQQGEwJYWDESMBAG
A1UECAwJU3RhdGVOYW1lMREwDwYDVQQHDAhDaXR5TmFtZTEUMBIGA1UECgwLQ29t
cGFueU5hbWUxGzAZBgNVBAsMEkNvbXBhbnlTZWN0aW9uTmFtZTEdMBsGA1UEAwwU
Q29tbW9uTmFtZU9ySG9zdG5hbWUwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQDJ+g1z+nZFj0DFc+wV1AH4giAqDFhwx25yziaQJnHkoXpllN0bRpr50Fhx
QsMuMqjMCM0eFPo7kDWSrAUusdh3j1jdF7KidZRZdyUgioQvbmNtqtwdU/38Gq0s
yTspOTZTWziIMFsLYxvFb0Ia8dwQFvyDB3pM1/eXFJMr2Cz/nF4/g4IUFEtPY2pe
4sQDMhszvROMfI5LB4SmGiim/zUS7+ZDEQigf10/CrbLntjXHiFy+V0hcrDQXkuY
gXU9HsiwRpCx+GRqcI6SWPIZ2rYa7goAVBENv5KWFXdWaOFCpi55XldYYzDeaXWa
On9EOjyylLbJ2pv/yhpt/VYDC6Lsqki1XcU/zeETcBRlqzCwZhMqFErSPRBGxaTN
lEGKNS2qQnE7I7bmUksRcCxA0WIn4V2xxxGHkLFnEsCipFGgMm7LaGSukYZGQ2MF
0ELqaaRMaTBfvPhFRtolkFChm7JQlCrk9j47b07OTMj4KUoBkhYPNoK26sHjvKue
Iks7BTGy8NdZk3jVDNEgGwngj+bNyJcjDvq+cHhWhS/N8c8Tg2tFb3ZOnwbyFPLf
1DdsP/SW6TybMc6NtpWvmSCFMvxqRpr+7LscpBGThepAUaoPkATnT6UdYGqr+l5N
z1+cZEuHBF7Jk5xThXCCFwuBC/KREvRuU2LyAhm9RV1ci8XyYwIDAQABo1MwUTAd
BgNVHQ4EFgQUa1FF7jssOpGY0UY1qbFQH83rxOQwHwYDVR0jBBgwFoAUa1FF7jss
OpGY0UY1qbFQH83rxOQwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC
AgEAhRPDPmVetjE5w/igfnGiNUKhQiqB3E6yT1wMwF5LbNIC+h8E8WqF0mzF3XQC
t9+Eib+mFptHGfRGdA8JmJ79xBgch6JLagw8Ot6hA1oEqWy/PAzcFLemccvhonhZ
+N7umRO27agIYVJ71mxlS7rD1SItBKZ+g4/Lt3sr/iDM0kIhWjWdQiYgJMNhH2Mt
XAd53tmPSom1Vsca1ZorB0HJNIw8RB7QYwceAi0xk0Tui7Z6pHg8p99XHCK/2vD+
+u5a1nHjtrAvLdNti8o82EUScK4j91CODCm5bv2KDMIUUv+1O83UMicK8DLSIi0P
1acuIdbRrY5JiYewCC+NFCfODZT6y00QL7aF9cbT+ZySiYtpBUiLFlIZCw4dcJ6r
yYJca1gv+kqZugdgLD7nXWIzXgfreQ+Q/BiVouXroXaWqU/9dqzGw4LlPSpyqI9o
nFdu970j4BvVc809i1aNo3odGz9yRx7wOOMyiyHoRpYenlf3hFbqJHS2FUgJ/Ajl
TC1joWDPdwdqQY56WOpk2LMK1KJuMYQEEHb3Ib0ZaxXAT6Nd12VashG5gYhw9vGn
aZ3nCnZFahgB2tB5DtkK+p0PQVpZOv6/0CKSsfQkBFLTjeVIRnlL/UlK3VlwHH9h
0LfsLFabedOS9R1XIaQ7aKXPK4tKmNb5H/ONzH2zxU5Oqoo=
-----END TRUSTED CERTIFICATE-----
`)

	certs, err := decodeCertificates(data)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(certs) != 2 {
		t.Errorf("unexpected number of certs: %d", len(certs))
	}
}