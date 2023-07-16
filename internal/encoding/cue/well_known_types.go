/*
Copyright Arata Furukawa

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cue

import (
	"context"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"google.golang.org/protobuf/compiler/protogen"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
)

func (g *Generator) wellKnownTypeEnum(ctx context.Context, e *protogen.Enum) (ast.Expr, bool, error) {
	switch e.Desc.FullName() {
	case "google.protobuf.NullValue":
		return &ast.Ident{Name: "null"}, true, nil
	}
	return nil, false, nil
}

func orNull(a ast.Expr) ast.Expr {
	return &ast.BinaryExpr{
		X: &ast.BasicLit{
			Kind:  token.NULL,
			Value: "null",
		},
		Op: token.OR,
		Y:  a,
	}
}

func (g *Generator) wellKnownTypeMessage(ctx context.Context, m *protogen.Message) (expr ast.Expr, err error) {
	switch m.Desc.FullName() {
	case "google.protobuf.Any":
		any := &ast.StructLit{
			Elts: []ast.Decl{
				&ast.Field{
					Label: &ast.BasicLit{
						Kind:  token.STRING,
						Value: `"@type"`,
					},
					Optional: token.Blank.Pos(),
					Value:    g.UseBuiltinType(ctx, "string"),
				},
				&ast.Ellipsis{},
			},
		}
		return orNull(any), nil
	case "google.protobuf.Struct":
		st := &ast.StructLit{
			Elts: []ast.Decl{
				&ast.Field{
					Label: &ast.ListLit{
						Elts: []ast.Expr{
							g.UseBuiltinType(ctx, "string"),
						},
					},
					Value: &ast.Ident{Name: "_"},
				},
			},
		}
		return orNull(st), nil
	case "google.protobuf.Value":
		return orNull(&ast.Ident{Name: "_"}), nil
	case "google.protobuf.ListValue":
		return &ast.ListLit{
			Elts: []ast.Expr{
				&ast.Ellipsis{},
			},
		}, nil
	case "google.protobuf.BoolValue":
		return orNull(g.UseBuiltinType(ctx, "bool")), nil
	case "google.protobuf.StringValue":
		return orNull(g.UseBuiltinType(ctx, "string")), nil
	case "google.protobuf.BytesValue":
		return orNull(g.UseBuiltinType(ctx, "bytes")), nil
	case "google.protobuf.Int32Value":
		return orNull(g.UseBuiltinType(ctx, "int32")), nil
	case "google.protobuf.Int64Value":
		return orNull(g.UseBuiltinType(ctx, "int64")), nil
	case "google.protobuf.UInt32Value":
		return orNull(g.UseBuiltinType(ctx, "uint32")), nil
	case "google.protobuf.UInt64Value":
		return orNull(g.UseBuiltinType(ctx, "uint64")), nil
	case "google.protobuf.FloatValue":
		return orNull(g.UseBuiltinType(ctx, "float32")), nil
	case "google.protobuf.DoubleValue":
		return orNull(g.UseBuiltinType(ctx, "float64")), nil
	case "google.protobuf.Empty":
		return orNull(&ast.CallExpr{
			Fun: &ast.Ident{Name: "close"},
			Args: []ast.Expr{
				&ast.StructLit{},
			},
		}), nil
	case "google.protobuf.Timestamp":
		// x, err := g.Import(ctx, "time", "builtin_time", "Time", false)
		// if err != nil {
		// 	return nil, false, fmt.Errorf("import: %w", err)
		// }
		x := g.UseBuiltinType(ctx, "string")
		return orNull(x), nil
	case "google.protobuf.Duration":
		// x, err := g.Import(ctx, "time", "builtin_time", "Duration", false)
		// if err != nil {
		// 	return nil, false, fmt.Errorf("import: %w", err)
		// }
		x := g.UseBuiltinType(ctx, "string")
		return orNull(x), nil
	case "google.protobuf.FieldMask":
		fm := &ast.StructLit{
			Elts: []ast.Decl{
				&ast.Field{
					Label: &ast.Ident{Name: "paths"},
					Value: &ast.ListLit{
						Elts: []ast.Expr{
							&ast.Ellipsis{
								Type: g.UseBuiltinType(ctx, "string"),
							},
						},
					},
				},
			},
		}
		return orNull(fm), nil
	}
	return nil, nil
}
