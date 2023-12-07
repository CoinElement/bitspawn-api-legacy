package integration

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bitspawngg/bitspawn-api/server"
)

const (
	port = "8080"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	go func() {
		_ = server.Start()
	}()
	time.Sleep(3 * time.Second)
}

func shutdown() {
}

func httpHelper(httpMethod string, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(httpMethod, url, body)
	req.Header.Set("X-Auth-Token", "eyJraWQiOiJINTN4VDhGcWl3cEpwNjJDMFZaVEUrcUJYOFVnTWtRN3BsYm44Y3V1UUQ4PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJjMDI1YWNmZi01OTYxLTQ5NzMtYWQzYi1kYTgwM2M4Mjg1NDkiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiY3VzdG9tOmNvbmZpcm1fYWdlIjoiMjAyMS0wOS0wN1QxNToxMTowMi4zNThaIiwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfTVJWcXNMNEN3IiwiY3VzdG9tOmluaXRfZGlzcGxheV9uYW1lIjoiTW9oYW1tYWQiLCJjb2duaXRvOnVzZXJuYW1lIjoibW5hYmJhc2FiYWRpQGdtYWlsLmNvbSIsImN1c3RvbTphY2NlcHRfdGFuZGMiOiIyMDIxLTA5LTA3VDE1OjExOjAyLjM1OFoiLCJhdWQiOiI2aXFzZjNkdHFyZzlibWs1NHZmcGE4NnRwayIsImV2ZW50X2lkIjoiYmVhNWNiMmUtN2IzMC00NGM3LWI0ZDgtMzQ5YjRhNjYwNDg5IiwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE2MzE0NjA2NDIsImV4cCI6MTYzMTQ2NDI0MywiaWF0IjoxNjMxNDYwNjQzLCJlbWFpbCI6Im1uYWJiYXNhYmFkaUBnbWFpbC5jb20ifQ.aQ8T3FUQ6SjRN-45HN_O61PnLgpAC_3Gvygj12bKr2hiNjQWtsTY6tWNIOO0u6dI-wIwRwapgoD5-yAOIwQOTXV_4AkV_lfWDqN24iqSgcKgMgMwsAboRAeSpPE_d0FL9vuk2M12tqI9T2sPiQD4M3C5H9vdT2phXH2XZEBu68axnElMpYbdwT1tHaz6M70F8tCgwA2Dek9kXKeJ96-lqPNBSmBeTiIUp7HfmYHS3f-_n4ZjKNayOZEio8jdPHfTDHG4E6YaOIaCbMuIdojY_BxqpBpbjki9zyvDiYFKuPgN9TE-iMmc5tT6hO4JGpAVDJ2h6CfQ4KB0fpYbXfPdQQ")
	return http.DefaultClient.Do(req)
}
