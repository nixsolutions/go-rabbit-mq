package connection

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRConnectionInterface(ctrl)
	m.EXPECT().Close().Return(nil)
	c := Connection{connection: m}
	err := c.Close()
	if err != nil {
		t.Error("Close connection error: ", err)
	}
}

func TestGetConnector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRConnectionInterface(ctrl)
	c := Connection{connection: m}
	if c.GetConnector() != m {
		t.Error("Get connector return incorrect connector")
	}
}
