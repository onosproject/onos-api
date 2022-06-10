// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"errors"
	"fmt"
	"strconv"
)

// PlmnID is a globally unique network identifier (Public Land Mobile Network)
type PlmnID uint32

// GnbID is a 5G gNodeB Identifier
type GnbID uint64

// EnbID is an eNodeB Identifier
type EnbID uint32

// CellID is a node-local cell identifier; 4 bits for 4G; 14 bits fo 5G
type CellID uint16

// NodeID is a general abstraction of a global E2 node identifier.
// It holds appropriately bit-shifted concatenation of either:
//		- [PLMNID + GnbID] or
//		- [PLMNID + EnbID]
// To extract the corresponding components, application must use the
// appropriate 4G or 5G method provided below.
type NodeID uint64

// NCI is a NR Cell Identifier; a 36-bit value (gNBID + CID)
type NCI uint64

// NCGI is NR Cell Global Identity (MCC+MNC+NCI)
type NCGI uint64

// ECI is a E-UTRAN Cell Identifier (gNBID + CID)
type ECI uint32

// ECGI is E-UTRAN Cell Global Identifier (MCC+MNC+ECI)
type ECGI uint64

// CRNTI is a cell-specific UE identifier
type CRNTI uint32

// MSIN is Mobile Subscriber Identification Number
type MSIN uint32

// IMSI is International Mobile Subscriber Identity
type IMSI uint64

// AmfUENgapID is AMF UE NGAP ID
type AmfUENgapID uint64

const (
	mask36 = 0xfffffffff
	mask28 = 0xfffffff
	mask20 = 0xfffff00
)

// EncodePlmnID encodes MCC and MNC strings into a PLMNID hex string
func EncodePlmnID(mcc string, mnc string) string {
	return mcc + mnc
}

// DecodePlmnID decodes MCC and MNC strings from PLMNID hex string
func DecodePlmnID(plmnID string) (mcc string, mnc string) {
	return plmnID[0:3], plmnID[3:]
}

// ToPlmnID encodes the specified MCC and MNC strings into a numeric PLMNID
func ToPlmnID(mcc string, mnc string) PlmnID {
	s := EncodePlmnID(mcc, mnc)
	n, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return 0
	}
	return PlmnID(n)
}

// PlmnIDFromHexString converts string form of PLMNID in its hex form into a numeric one suitable for APIs
func PlmnIDFromHexString(plmnID string) PlmnID {
	n, err := strconv.ParseUint(plmnID, 16, 32)
	if err != nil {
		return 0
	}
	return PlmnID(n)
}

// PlmnIDFromString converts string form of PLMNID given as a simple MCC-MCN catenation into a numeric one suitable for APIs
func PlmnIDFromString(plmnID string) PlmnID {
	return ToPlmnID(plmnID[0:3], plmnID[3:])
}

// PlmnIDToString generates the MCC-MCN catenation format from the specified numeric PLMNID
func PlmnIDToString(plmnID PlmnID) string {
	hexString := fmt.Sprintf("%x", plmnID)
	mcc, mnc := DecodePlmnID(hexString)
	return mcc + mnc
}

// 5G Identifiers
var (
	gnbBits uint8  = 22
	cidBits uint8  = 14
	gnbMask uint64 = 0b111111111111111111111100000000000000
	cidMask uint64 = 0b000000000000000000000011111111111111
)

// SetNCIBitSplit sets how the NCI bits are split between gNBID and CID
func SetNCIBitSplit(gnb uint8, cid uint8) error {
	if (gnb+cid) == 36 && 4 <= cid && cid <= 14 {
		cidBits = cid
		gnbMask = 0
		cidMask = 0
		for i := 0; i < 64; i++ {
			b := uint8(i)
			if b < cid {
				cidMask |= 1 << i
			}
			if cid <= b && b < (cid+gnb) {
				gnbMask |= 1 << i
			}
		}
		return nil
	}
	return errors.New("invalid bit split")
}

// ToNCI produces NCI from the specified components
func ToNCI(gnbID GnbID, cid CellID) NCI {
	return NCI(uint(gnbID)<<cidBits | uint(cid))
}

// ToNCGI produces NCGI from the specified components
func ToNCGI(plmnID PlmnID, nci NCI) NCGI {
	return NCGI(uint(plmnID)<<36 | (uint(nci) & mask36))
}

// ToNodeID produces a 5G global node ID as a catenation of PLMNID + GnbID
func ToNodeID(plmnID PlmnID, gnbID GnbID) NodeID {
	return NodeID(uint(plmnID)<<gnbBits | uint(gnbID))
}

// To5GNodeID produces a 5G global node ID as a catenation of PLMNID + GnbID
func To5GNodeID(plmnID PlmnID, gnbID GnbID) NodeID {
	return ToNodeID(plmnID, gnbID)
}

// GetPlmnID extracts PLMNID from the specified NCGI
func GetPlmnID(ncgi uint64) PlmnID {
	return PlmnID(ncgi >> 36)
}

// Get5GPlmnID extracts PLMNID from the specified NCGI
func Get5GPlmnID(ncgi uint64) PlmnID {
	return GetPlmnID(ncgi)
}

// GetNCI extracts NCI from the specified NCGI
func GetNCI(ncgi NCGI) NCI {
	return NCI(ncgi & mask36)
}

// GetGnbID extracts gNodeB ID from the specified NCGI or NCI
func GetGnbID(id uint64) GnbID {
	return GnbID((id & gnbMask) >> cidBits)
}

// GetCellID extracts Cell ID from the specified NCGI or NCI
func GetCellID(id uint64) CellID {
	return CellID(id & cidMask)
}

// Get5GCellID extracts Cell ID from the specified NCGI or NCI
func Get5GCellID(id uint64) CellID {
	return GetCellID(id)
}

// 4G Identifiers

// ToECI produces ECI from the specified components
func ToECI(enbID EnbID, cid CellID) ECI {
	return ECI(uint(enbID)<<8 | uint(cid))
}

// ToECGI produces ECGI from the specified components
func ToECGI(plmnID PlmnID, eci ECI) ECGI {
	return ECGI(uint(plmnID)<<28 | (uint(eci) & mask28))
}

// To4GNodeID produces a 54 global node ID as a catenation of PLMNID + EnbID
func To4GNodeID(plmnID PlmnID, enbID EnbID) NodeID {
	return NodeID(uint(plmnID)<<28 | uint(enbID))
}

// Get4GPlmnID extracts PLMNID from the specified ECGI or IMSI
func Get4GPlmnID(id uint64) PlmnID {
	return PlmnID(id >> 28)
}

// Get4GCellID extracts Cell ID from the specified ECGI
func Get4GCellID(id uint64) CellID {
	return CellID(id & 0xff)
}

// GetEnbID extracts Enb ID from the specified ECGI
func GetEnbID(id uint64) EnbID {
	return EnbID((id & mask20) >> 8)
}

// GetECI extracts ECI from the specified ECGI
func GetECI(id uint64) ECI {
	return ECI(id & mask28)
}
