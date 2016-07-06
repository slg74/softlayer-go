package data_types

import (
	"time"
)

type SoftLayer_Network_Vlan struct {
	AccountId       int        `json:"accountId,omitempty"`
	Id              int        `json:"Id,omitempty"`
	ModifyDate      *time.Time `json:"modifyDate,omitempty"`
	Name            string     `json:"name,omitempty"`
	NetworkVrfId    int        `json:"networkVrfId,omitempty"`
	Note            string     `json:"note,omitempty"`
	PrimarySubnetId int        `json:"primarySubnetId,omitempty"`
	VlanNumber      int        `json:"vlanNumber,omitempty"`
}
