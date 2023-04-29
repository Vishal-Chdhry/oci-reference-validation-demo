package main

import "testing"

var (
	certificate = `-----BEGIN CERTIFICATE-----
	MIIDVzCCAj+gAwIBAgICAL8wDQYJKoZIhvcNAQELBQAwWjELMAkGA1UEBhMCVVMx
	CzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxlMQ8wDQYDVQQKEwZOb3Rhcnkx
	GzAZBgNVBAMTEndhYmJpdC1uZXR3b3Jrcy5pbzAeFw0yMzA0MTkwODE0MjBaFw0y
	NDA0MTkwODE0MjBaMFoxCzAJBgNVBAYTAlVTMQswCQYDVQQIEwJXQTEQMA4GA1UE
	BxMHU2VhdHRsZTEPMA0GA1UEChMGTm90YXJ5MRswGQYDVQQDExJ3YWJiaXQtbmV0
	d29ya3MuaW8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCsLni/3fDd
	MUKwzfSkc8orPrCLCRoiampqeBIQNlKZGP5UL7BP6neav/oy8p09b1ArBwKh1coO
	LBLZh2BPVmNyoSmq8Ip9md7OQi5ApaC0cys/GTmJQov/tmFq+o9GOdK6yVjX8bS/
	NPrlaM5XCzY9WaTvW/ODu7MjVqBFGHQa9AFjEJZrxtPkOF0WV7ffPiQ5RZnuccK0
	V0ycKCtsadXFdQZl2w/Li5DZHTVyc6mhtIyPBSlsIbpQXB5SLPVb7jsDKk/XoNd6
	DwB6lYBKgUtNL1cFhjwdWaTrqy5ymiuMKjyeeileZinpQH129jfTjL6Rx8Tv7PMf
	Cu0hfgOMQ5TPAgMBAAGjJzAlMA4GA1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggr
	BgEFBQcDAzANBgkqhkiG9w0BAQsFAAOCAQEARnay50VbglZqudCIbsp96HHHkf+p
	lvRkFOo148UaeeN12k803yRx2PxYzmUHvNKjX76eAWT3KzQEG9FRId5mSr2gO8sp
	JJs9WITfKAXnNGSz/rd8tN6eSUKgN+32A9reKPO+0ntN1dJDb8R9892ze88UpDva
	wgFXsl+jKC0y4HSEeHufIqgyI6freRFvUEAAMJ0SKPfNROc/1WcW5F4u6Sbh2UNn
	lXOXSo6yg2xRJxATkHc9PE7tSBnDRIwjJG3ZyRLdyTeuzOl95ARImUtPAn2ElODf
	21yAXY47fCnTSkTw1ZrSQa6jWG7Er/mgsTVT98PcKIGyx3iqPFNJjUkj/Q==
	-----END CERTIFICATE-----`

	repo_name = "jimnotarytest.azurecr.io/jim/net-monitor:v1"

	artifact_type = "application/vnd.cncf.notary.signature"
)

func Test_ORAS(t *testing.T) {
	ORAS(repo_name, certificate, artifact_type)
}

func Test_Regclient(t *testing.T) {
	Regclient(repo_name, certificate, artifact_type)
}

func Test_Flux(t *testing.T) {
	Flux(repo_name, certificate, artifact_type)
}

func Test_ACR_Regclient(t *testing.T) {
	ACRRegclient()
}

func Test_GCRCrane(t *testing.T) {
	GCRCrane(repo_name, certificate, artifact_type)
}
