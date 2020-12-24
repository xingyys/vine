// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/lack-io/cli"

	"github.com/lack-io/vine/client/cli/namespace"
	"github.com/lack-io/vine/client/cli/util"
	pb "github.com/lack-io/vine/proto/auth"
	"github.com/lack-io/vine/service/client"
	"github.com/lack-io/vine/service/context"
	"github.com/lack-io/vine/service/errors"
)

func listRules(ctx *cli.Context) error {
	cli := pb.NewRulesService("auth", client.DefaultClient)

	env, err := util.GetEnv(ctx)
	if err != nil {
		return err
	}
	ns, err := namespace.Get(env.Name)
	if err != nil {
		return fmt.Errorf("Error getting namespace: %v", err)
	}

	rsp, err := cli.List(context.DefaultContext, &pb.ListRequest{
		Options: &pb.Options{Namespace: ns},
	}, client.WithAuthToken())
	if err != nil {
		return fmt.Errorf("Error listing rules: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	defer w.Flush()

	formatResource := func(r *pb.Resource) string {
		return strings.Join([]string{r.Type, r.Name, r.Endpoint}, ":")
	}

	// sort rules using resource name and priority to keep the list consistent
	sort.Slice(rsp.Rules, func(i, j int) bool {
		resI := formatResource(rsp.Rules[i].Resource) + string(rsp.Rules[i].Priority)
		resJ := formatResource(rsp.Rules[j].Resource) + string(rsp.Rules[j].Priority)
		return sort.StringsAreSorted([]string{resJ, resI})
	})

	fmt.Fprintln(w, strings.Join([]string{"ID", "Scope", "Access", "Resource", "Priority"}, "\t\t"))
	for _, r := range rsp.Rules {
		res := formatResource(r.Resource)
		if r.Scope == "" {
			r.Scope = "<public>"
		}
		fmt.Fprintln(w, strings.Join([]string{r.Id, r.Scope, r.Access.String(), res, fmt.Sprintf("%d", r.Priority)}, "\t\t"))
	}

	return nil
}

func createRule(ctx *cli.Context) error {
	env, err := util.GetEnv(ctx)
	if err != nil {
		return err
	}
	ns, err := namespace.Get(env.Name)
	if err != nil {
		return fmt.Errorf("Error getting namespace: %v", err)
	}

	rule, err := constructRule(ctx)
	if err != nil {
		return err
	}

	cli := pb.NewRulesService("auth", client.DefaultClient)
	_, err = cli.Create(context.DefaultContext, &pb.CreateRequest{
		Rule: rule, Options: &pb.Options{Namespace: ns},
	}, client.WithAuthToken())
	if verr := errors.FromErr(err); verr != nil {
		return fmt.Errorf("Error: %v", verr.Detail)
	} else if err != nil {
		return err
	}

	fmt.Println("Rule created")
	return nil
}

func deleteRule(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return fmt.Errorf("Expected one argument: ID")
	}

	env, err := util.GetEnv(ctx)
	if err != nil {
		return err
	}
	ns, err := namespace.Get(env.Name)
	if err != nil {
		return fmt.Errorf("Error getting namespace: %v", err)
	}

	cli := pb.NewRulesService("auth", client.DefaultClient)
	_, err = cli.Delete(context.DefaultContext, &pb.DeleteRequest{
		Id: ctx.Args().First(), Options: &pb.Options{Namespace: ns},
	}, client.WithAuthToken())
	if verr := errors.FromErr(err); err != nil {
		return fmt.Errorf("Error: %v", verr.Detail)
	} else if err != nil {
		return err
	}

	fmt.Println("Rule deleted")
	return nil
}

func constructRule(ctx *cli.Context) (*pb.Rule, error) {
	if ctx.Args().Len() != 1 {
		return nil, fmt.Errorf("Too many arguments, expected one argument: ID")
	}

	var access pb.Access
	switch ctx.String("access") {
	case "granted":
		access = pb.Access_GRANTED
	case "denied":
		access = pb.Access_DENIED
	default:
		return nil, fmt.Errorf("Invalid access: %v, must be granted or denied", ctx.String("access"))
	}

	resComps := strings.Split(ctx.String("resource"), ":")
	if len(resComps) != 3 {
		return nil, fmt.Errorf("Invalid resource, must be in the format type:name:endpoint")
	}

	return &pb.Rule{
		Id:       ctx.Args().First(),
		Access:   access,
		Scope:    ctx.String("scope"),
		Priority: int32(ctx.Int("priority")),
		Resource: &pb.Resource{
			Type:     resComps[0],
			Name:     resComps[1],
			Endpoint: resComps[2],
		},
	}, nil
}
