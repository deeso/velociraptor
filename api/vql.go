/*
   Velociraptor - Hunting Evil
   Copyright (C) 2019 Velocidex Innovations.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package api

import (
	"encoding/json"
	"io"

	"github.com/Velocidex/ordereddict"
	errors "github.com/pkg/errors"
	context "golang.org/x/net/context"
	api_proto "www.velocidex.com/golang/velociraptor/api/proto"
	"www.velocidex.com/golang/velociraptor/artifacts"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/file_store/csv"
	"www.velocidex.com/golang/velociraptor/logging"
	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	"www.velocidex.com/golang/vfilter"
)

func RunVQL(
	ctx context.Context,
	config_obj *config_proto.Config,
	principal string,
	env *ordereddict.Dict,
	query string) (*api_proto.GetTableResponse, error) {

	result := &api_proto.GetTableResponse{}

	scope := artifacts.ScopeBuilder{
		Config:     config_obj,
		Env:        env,
		ACLManager: vql_subsystem.NewServerACLManager(config_obj, principal),
		Logger:     logging.NewPlainLogger(config_obj, &logging.ToolComponent),
	}.Build()
	defer scope.Close()

	vql, err := vfilter.Parse(query)
	if err != nil {
		return nil, err
	}

	sub_ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for row := range vql.Eval(sub_ctx, scope) {
		if len(result.Columns) == 0 {
			result.Columns = scope.GetMembers(row)
		}

		new_row := &api_proto.Row{}
		for _, column := range result.Columns {
			value, pres := scope.Associative(row, column)
			if !pres {
				value = ""
			}
			new_row.Cell = append(new_row.Cell, csv.AnyToString(value))
		}

		result.Rows = append(result.Rows, new_row)
	}

	return result, nil
}

func StoreVQLAsCSVAndJsonFile(
	ctx context.Context,
	config_obj *config_proto.Config,
	principal string,
	env *ordereddict.Dict,
	query string,
	csv_fd io.Writer,
	json_fd io.Writer) error {

	scope := artifacts.ScopeBuilder{
		Config:     config_obj,
		ACLManager: vql_subsystem.NewServerACLManager(config_obj, principal),
		Logger: logging.NewPlainLogger(config_obj,
			&logging.ToolComponent),
		Env: env,
	}.Build()
	defer scope.Close()

	vql, err := vfilter.Parse(query)
	if err != nil {
		return err
	}

	csv_writer := csv.GetCSVAppender(scope, csv_fd, true /* write_headers */)
	defer csv_writer.Close()

	sub_ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for row := range vql.Eval(sub_ctx, scope) {
		csv_writer.Write(row)

		serialized, err := json.Marshal(row)
		if err != nil {
			continue
		}
		_, err = json_fd.Write(serialized)
		if err != nil {
			return errors.WithStack(err)
		}
		json_fd.Write([]byte("\n"))
	}

	return nil
}
