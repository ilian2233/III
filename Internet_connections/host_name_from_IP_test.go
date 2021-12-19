package Internet_connections

import "testing"

func TestGetHostNameByIP( t *testing.T){

	name, err := getHostNameByIP("195.216.228.3")
	if err != nil {
		t.Error(err)
	}

	if name != "www.tu-varna.bg"{
		t.Errorf("Name does not match. \n%v", name)
	}
}