package types

import (
	"github.com/0xzer/snapper/protos"
)

var BundleTag = "12.42.0"
var AppVersion = &protos.AppVersion{
	VersionNumber: &protos.VersionNumber{
		Major: 12,
		Minor: 42,
	},
	Flavor: protos.Flavor_PRODUCTION,
}
var OSInfo = &protos.OSInfo{
	OsType: protos.OsType_OS_TYPE_OTHER,
}

func NewDevice() string {
	device := &protos.Device{
		AppVersion: AppVersion,
		OsInfo: OSInfo,
		BundleTag: BundleTag,
	}
	_, deviceEncodedStr, _ := protos.EncodeProtoMessageB64(device)

	return deviceEncodedStr
}