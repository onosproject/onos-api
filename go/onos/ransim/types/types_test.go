// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlmnIDEncode(t *testing.T) {
	plmnID := EncodePlmnID("208", "20")
	assert.Equal(t, "20820", plmnID)

	mcc, mnc := DecodePlmnID(plmnID)
	assert.Equal(t, "208", mcc)
	assert.Equal(t, "20", mnc)

	plmnID = EncodePlmnID("310", "170")
	assert.Equal(t, "310170", plmnID)

	mcc, mnc = DecodePlmnID(plmnID)
	assert.Equal(t, "310", mcc)
	assert.Equal(t, "170", mnc)

	assert.Equal(t, PlmnID(0x310170), ToPlmnID("310", "170"))
	assert.Equal(t, PlmnID(0x310170), PlmnIDFromHexString("310170"))

	assert.Equal(t, PlmnID(0x310170), PlmnIDFromString("310170"))
	assert.Equal(t, "310170", PlmnIDToString(PlmnID(0x310170)))
}

func TestFullShiftBasics(t *testing.T) {
	plmnID := PlmnID(0xBBBCCC)
	cellID := CellID(0xAA)
	enbID := EnbID(0xFFFFF)

	assert.Equal(t, ECI(0xFFFFFAA), ToECI(enbID, cellID))
	assert.Equal(t, ECGI(0xBBBCCCFFFFFAA), ToECGI(plmnID, ToECI(enbID, cellID)))
}

func TestTypesWithFullShift(t *testing.T) {
	plmnID := PlmnID(0xbbbccc)
	cellID := CellID(0x12)
	enbID := EnbID(0xf8f8f)

	eci := ToECI(enbID, cellID)
	ecgi := ToECGI(plmnID, eci)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, Get4GPlmnID(uint64(ecgi)), "incorrect PLMNID")
	assert.Equal(t, enbID, GetEnbID(uint64(ecgi)), "incorrect EnbID")
	assert.Equal(t, eci, GetECI(uint64(ecgi)), "incorrect ECI")
	assert.Equal(t, cellID, Get4GCellID(uint64(ecgi)), "incorrect CID")
}

func TestTypesWithPartialShift(t *testing.T) {
	plmnID := PlmnID(0xbbbccc)
	cellID := CellID(0x12)
	enbID := EnbID(0xf8f8f)

	eci := ToECI(enbID, cellID)
	ecgi := ToECGI(plmnID, eci)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, Get4GPlmnID(uint64(ecgi)), "incorrect PLMNID")
	assert.Equal(t, enbID, GetEnbID(uint64(ecgi)), "incorrect EnbID")
	assert.Equal(t, eci, GetECI(uint64(ecgi)), "incorrect ECI")
	assert.Equal(t, cellID, Get4GCellID(uint64(ecgi)), "incorrect CID")
}

func Test5GTypes_22_14(t *testing.T) {
	err := SetNCIBitSplit(22, 14)

	fmt.Printf("%040b\n%040b\n", gnbMask, cidMask)
	assert.NoError(t, err)
	plmnID := PlmnID(0xbbbaaa)
	gnbID := GnbID(0xfffff)
	cellID := CellID(0x0f)

	fmt.Printf("%040b\n%040b\n%040b\n", plmnID, gnbID, cellID)

	nci := ToNCI(gnbID, cellID)
	ncgi := ToNCGI(plmnID, nci)

	fmt.Printf("%064b\n%064b\n", nci, ncgi)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, GetPlmnID(uint64(ncgi)), "incorrect PLMNID")
	assert.Equal(t, nci, GetNCI(ncgi), "incorrect NCI")
	assert.Equal(t, gnbID, GetGnbID(uint64(ncgi)), "incorrect NCGI GnbID")
	assert.Equal(t, cellID, GetCellID(uint64(ncgi)), "incorrect NCGI CID")

	assert.Equal(t, gnbID, GetGnbID(uint64(nci)), "incorrect NCI GnbID")
	assert.Equal(t, cellID, GetCellID(uint64(nci)), "incorrect NCI CID")
}

func Test5GTypes_32_4(t *testing.T) {
	err := SetNCIBitSplit(32, 4)

	fmt.Printf("%040b\n%040b\n", gnbMask, cidMask)
	assert.NoError(t, err)
	plmnID := PlmnID(0xbbbaaa)
	gnbID := GnbID(0xfffff)
	cellID := CellID(0x0f)

	fmt.Printf("%040b\n%040b\n%040b\n", plmnID, gnbID, cellID)

	nci := ToNCI(gnbID, cellID)
	ncgi := ToNCGI(plmnID, nci)

	fmt.Printf("%064b\n%064b\n", nci, ncgi)

	// NOTE: These work for the given values of cellID and enbID, but may fail with "lesser" values that result
	// in shifts less than 8 or 28 bits respectively.
	assert.Equal(t, plmnID, GetPlmnID(uint64(ncgi)), "incorrect PLMNID")
	assert.Equal(t, nci, GetNCI(ncgi), "incorrect NCI")
	assert.Equal(t, gnbID, GetGnbID(uint64(ncgi)), "incorrect NCGI GnbID")
	assert.Equal(t, cellID, GetCellID(uint64(ncgi)), "incorrect NCGI CID")

	assert.Equal(t, gnbID, GetGnbID(uint64(nci)), "incorrect NCI GnbID")
	assert.Equal(t, cellID, GetCellID(uint64(nci)), "incorrect NCI CID")
}

func TestSimValues(t *testing.T) {
	plmnID := ToPlmnID("310", "170")
	assert.Equal(t, PlmnID(0x310170), plmnID)

	enb1 := EnbID(144470)
	enb2 := EnbID(144471)
	ecgi11 := ToECGI(plmnID, ToECI(enb1, CellID(1)))
	ecgi12 := ToECGI(plmnID, ToECI(enb1, CellID(2)))
	ecgi21 := ToECGI(plmnID, ToECI(enb2, CellID(1)))
	ecgi22 := ToECGI(plmnID, ToECI(enb2, CellID(2)))

	fmt.Printf("%d\n%d\n%d\n%d\n", ecgi11, ecgi12, ecgi21, ecgi22)
}
