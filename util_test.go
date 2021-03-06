package btcapi

import "testing"

func TestExtendedPublicKeyDetails(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.ExtendedPublicKeyDetails("xpub6EuV33a2DXxAhoJTRTnr8qnysu81AA4YHpLY6o8NiGkEJ8KADJ35T64eJsStWsmRf1xXkEANVjXFXnaUKbRtFwuSPCLfDdZwYNZToh4LBCd")
	if err != nil {
		t.Error(err)
	}
}

func TestExtendedPublicKeyDetailsPage(t *testing.T) {
	btcapi := Config{
		ExplorerURL: TestingEndpoint,
	}
	_, err := btcapi.ExtendedPublicKeyDetailsPage("xpub6EuV33a2DXxAhoJTRTnr8qnysu81AA4YHpLY6o8NiGkEJ8KADJ35T64eJsStWsmRf1xXkEANVjXFXnaUKbRtFwuSPCLfDdZwYNZToh4LBCd", 20, 0)
	if err != nil {
		t.Error(err)
	}
}
