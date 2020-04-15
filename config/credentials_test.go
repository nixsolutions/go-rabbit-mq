package config

import "testing"

func TestCredentials_GetAMQPDialString(t *testing.T) {
	expected := "amqp://uTest:pTest@cTest:1111/"
	c := Credentials{username: "uTest", password: "pTest", clusterName: "cTest", port: "1111"}
	actual := c.GetAMQPDialString()
	if  actual != expected {
		t.Error("Dial strings not equal. Expect:", expected, "Actual:", actual)
	}
}
