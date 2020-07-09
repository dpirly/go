package visa

import (
	"testing"
)

func TestMain(t *testing.T) {

	conn, err:= Open("tcpip0::t6290e-c00119.local::hislip1", 1000)
	if err != nil {
		t.Errorf("open failed:%s", err)
		return
	}

	resp := conn.Query("*IDN?")
	t.Log("IDN:", resp)

	conn.Close()
}
