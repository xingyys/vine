// MIT License
//
// Copyright (c) 2020 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package exec

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lack-io/vine/core/server"
	"github.com/lack-io/vine/lib/proxy"
	"github.com/lack-io/vine/proto/apis/errors"
)

type Proxy struct {
	options proxy.Options

	// The file or directory to read from
	Endpoint string
}

func filePath(eps ...string) string {
	p := filepath.Join(eps...)
	return strings.Replace(p, "../", "", -1)
}

func getEndpoint(hdr map[string]string) string {
	ep := hdr["Vine-Endpoint"]
	if len(ep) > 0 && ep[0] == '/' {
		return ep
	}
	return ""
}

func (p *Proxy) ProcessMessage(ctx context.Context, msg server.Message) error {
	return nil
}

// ServeRequest honours the server.Router interface
func (p *Proxy) ServeRequest(ctx context.Context, req server.Request, rsp server.Response) error {
	if p.Endpoint == "" {
		exe, err := os.Executable()
		if err != nil {
			return err
		}
		// set the endpoint to the current path
		p.Endpoint = filepath.Dir(exe)
	}

	for {
		// get data
		_, err := req.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// get the header
		hdr := req.Header()

		// get endpoint
		endpoint := getEndpoint(hdr)

		// filepath
		file := filePath(p.Endpoint, endpoint)

		// exec the script or command
		// TODO: add args
		cmd := exec.Command(file)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.InternalServerError(req.Service(), err.Error())
		}

		// write back the header
		rsp.WriteHeader(hdr)
		// write the body
		err = rsp.Write(out)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return errors.InternalServerError(req.Service(), err.Error())
		}
	}

}

func (p *Proxy) String() string {
	return "exec"
}

//NewSingleHostProxy returns a router which sends requests to a single file
func NewSingleHostProxy(url string) proxy.Proxy {
	return &Proxy{
		Endpoint: url,
	}
}

// NewProxy returns a new proxy which will execute a script, binary or anything
func NewProxy(opts ...proxy.Option) proxy.Proxy {
	var options proxy.Options
	for _, o := range opts {
		o(&options)
	}
	p := new(Proxy)
	p.Endpoint = options.Endpoint

	return p
}
