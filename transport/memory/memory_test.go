// Copyright 2020 The vine Authors
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

package memory

import (
	"os"
	"testing"

	"github.com/lack-io/vine/transport"
)

func TestMemoryTransport(t *testing.T) {
	tr := NewTransport()

	// bind / listen
	l, err := tr.Listen("127.0.0.1:8080")
	if err != nil {
		t.Fatalf("Unexpected error listening %v", err)
	}
	defer l.Close()

	// accept
	go func() {
		if err := l.Accept(func(sock transport.Socket) {
			for {
				var m transport.Message
				if err := sock.Recv(&m); err != nil {
					return
				}
				if len(os.Getenv("IN_TRAVIS_CI")) == 0 {
					t.Logf("Server Received %s", string(m.Body))
				}
				if err := sock.Send(&transport.Message{
					Body: []byte(`pong`),
				}); err != nil {
					return
				}
			}
		}); err != nil {
			t.Fatalf("Unexpected error accepting %v", err)
		}
	}()

	// dial
	c, err := tr.Dial("127.0.0.1:8080")
	if err != nil {
		t.Fatalf("Unexpected error dialing %v", err)
	}
	defer c.Close()

	// send <=> receive
	for i := 0; i < 3; i++ {
		if err := c.Send(&transport.Message{
			Body: []byte(`ping`),
		}); err != nil {
			return
		}
		var m transport.Message
		if err := c.Recv(&m); err != nil {
			return
		}
		if len(os.Getenv("IN_TRAVIS_CI")) == 0 {
			t.Logf("Client Received %s", string(m.Body))
		}
	}

}

func TestListener(t *testing.T) {
	tr := NewTransport()

	// bind / listen on random port
	l, err := tr.Listen(":0")
	if err != nil {
		t.Fatalf("Unexpected error listening %v", err)
	}
	defer l.Close()

	// try again
	l2, err := tr.Listen(":0")
	if err != nil {
		t.Fatalf("Unexpected error listening %v", err)
	}
	defer l2.Close()

	// now make sure it still fails
	l3, err := tr.Listen(":8080")
	if err != nil {
		t.Fatalf("Unexpected error listening %v", err)
	}
	defer l3.Close()

	if _, err := tr.Listen(":8080"); err == nil {
		t.Fatal("Expected error binding to :8080 got nil")
	}
}
