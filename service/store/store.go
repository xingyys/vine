// Copyright 2020 lack
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package store is an interface for distributed data storage.
package store

import (
	"errors"
	"time"
)

var (
	// ErrNotFound is returned when a key doesn't exist
	ErrNotFound = errors.New("not found")
	// DefaultStore is the memory store.
	DefaultStore Store
)

// Store is a data storage interface
type Store interface {
	// Init initialises the store. It must perform any required setup on the backing storage implementation and check that it is ready for use, returning any errors.
	Init(...Option) error
	// Options allows you to view the current options.
	Options() Options
	// Read takes a single key name and optional ReadOptions. It returns matching []*Record or an error.
	Read(key string, opts ...ReadOption) ([]*Record, error)
	// Write() writes a record to the store, and returns an error if the record was not written.
	Write(r *Record, opts ...WriteOption) error
	// Delete removes the record with the corresponding key from the store.
	Delete(key string, opts ...DeleteOption) error
	// List returns any keys that match, or an empty list with no error if none matched.
	List(opts ...ListOption) ([]string, error)
	// Close the store
	Close() error
	// String returns the name of the implementation.
	String() string
}

// Record is an item stored or retrieved from a Store
type Record struct {
	// The key to store the record
	Key string `json:"key"`
	// The value within the record
	Value []byte `json:"value"`
	// Any associated metadata for indexing
	Metadata map[string]interface{} `json:"metadata"`
	// Time to expire a record: TODO: change to timestamp
	Expiry time.Duration `json:"expiry,omitempty"`
}
