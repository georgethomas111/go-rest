package main

import (
	"testing"

	"github.com/stretchrcom/testify/assert"
)

func TestPostKey(t *testing.T) {
	keyIns := &KeyResponse{
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDSJZZmFq+OThseFy+h9xv5baXtLfzC/LnWOTjtAzgisvT4yk/vcMV1QxsUkZ+xpriJGqDqhBS/vpOEIStwd3oPpp24UwF+kzgE2f+CFTxJvVFTs5G2gnxfAjtvyykHU0yO2vpHIqHfVQMGaXGNPyEtZ2+DnkuvCZDNyOBda8lbFPAQdiS6AHq9EOfZ/Zq9Ia3jPXZc0IwCSOL0D90jJsRNpFjCtQEBWoPyTAX06YEi3JFNDjK7OKkwBVqr63rfBx8ggz0cAQIQaicOqgmRJRM8exnpKCt2Lf5Brwf+Jn8o1WN2mm2j5R+z8ofAdsALx01alUXJfbD7111uFIzohYuR user@server",
	}

	keyIns.Populate()
	assert.Equal(t, keyIns.PublicKey, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDSJZZmFq+OThseFy+h9xv5baXtLfzC/LnWOTjtAzgisvT4yk/vcMV1QxsUkZ+xpriJGqDqhBS/vpOEIStwd3oPpp24UwF+kzgE2f+CFTxJvVFTs5G2gnxfAjtvyykHU0yO2vpHIqHfVQMGaXGNPyEtZ2+DnkuvCZDNyOBda8lbFPAQdiS6AHq9EOfZ/Zq9Ia3jPXZc0IwCSOL0D90jJsRNpFjCtQEBWoPyTAX06YEi3JFNDjK7OKkwBVqr63rfBx8ggz0cAQIQaicOqgmRJRM8exnpKCt2Lf5Brwf+Jn8o1WN2mm2j5R+z8ofAdsALx01alUXJfbD7111uFIzohYuR user@server", "public Key should not get altered")

	assert.Equal(t, "user@server", keyIns.Email, "Check for email")

	assert.Equal(t, "f7:ee:58:36:d8:20:38:6c:de:e2:1b:f0:48:ec:08:4c", keyIns.FingerPrint, "Check for fingerprints")

	// Add other tests when everting works
}
