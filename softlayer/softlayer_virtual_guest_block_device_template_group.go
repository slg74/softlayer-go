package softlayer

import (
	datatypes "github.com/maximilien/softlayer-go/data_types"
)

type SoftLayer_Virtual_Guest_Block_Device_Template_Group_Service interface {
	Service

	DeleteObject(id int) (bool, error)

	GetObject(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error)
	GetDatacenters(id int) ([]datatypes.SoftLayer_Location, error)
	GetSshKeys(id int) ([]datatypes.SoftLayer_Security_Ssh_Key, error)
	GetStatus(id int) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group_Status, error)

	GetStorageLocations(id int) ([]datatypes.SoftLayer_Location, error)

	GetImageType(id int) (datatypes.SoftLayer_Image_Type, error)

	CreateFromExternalSource(configuration datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration) (datatypes.SoftLayer_Virtual_Guest_Block_Device_Template_Group, error)
	CopyToExternalSource(configuration datatypes.SoftLayer_Container_Virtual_Guest_Block_Device_Template_Configuration) (bool, error)
}
