package scalyr

import (
	"strings"
	"testing"
)

func TestConfigureInstanceGenerateASessionUUID(t *testing.T) {
	instance := &Instance{}
	err := instance.Configure()

	if err != nil {
		t.Error(err)
	}

	if strings.Compare(instance.Session.String(), "") == 0 {
		t.FailNow()
	}
}

func TestConfigureMultipleTimeWontChangeSession(t *testing.T) {
	instance := &Instance{}
	instance.Configure()

	session := instance.Session.String()

	instance.Configure()

	if strings.Compare(session, instance.Session.String()) != 0 {
		t.FailNow()
	}
}
