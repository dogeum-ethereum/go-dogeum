// Copyright 2023 The go-dogeum Authors
// This file is part of the go-dogeum library.
//
// The go-dogeum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-dogeum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-dogeum library. If not, see <http://www.gnu.org/licenses/>

//go:build (arm64 || amd64) && !openbsd

package rawdb

import (
	"github.com/dogeum-network/go-dogeum/ethdb"
	"github.com/dogeum-network/go-dogeum/ethdb/pebble"
)

// Pebble is unsuported on 32bit architecture
const PebbleEnabled = true

// NewPebbleDBDatabase creates a persistent key-value database without a freezer
// moving immutable chain segments into cold storage.
func NewPebbleDBDatabase(file string, cache int, handles int, namespace string, readonly bool) (ethdb.Database, error) {
	db, err := pebble.New(file, cache, handles, namespace, readonly)
	if err != nil {
		return nil, err
	}
	return NewDatabase(db), nil
}
