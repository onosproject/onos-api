// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Package control defines an API for the flow control reconciliation library
package control

// FlowControl defines an abstract interface for interacting with the flow control reconciliation library
type FlowControl interface {
	// Write(deviceId, updates) error
	// Read(deviceId, entitites) (chan *entities, error)
	// PacketOut(deviceId, packet) error
	// Listen(deviceId) (chan *packetIn, error)
}
