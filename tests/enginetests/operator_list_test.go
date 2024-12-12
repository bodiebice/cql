// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enginetests

import (
	"context"
	"strings"
	"testing"

	"github.com/google/cql/interpreter"
	"github.com/google/cql/model"
	"github.com/google/cql/parser"
	"github.com/google/cql/result"
	"github.com/google/cql/types"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantModel  model.IExpression
		wantResult result.Value
	}{
		{
			name: "Non Empty List",
			cql:  "exists({1})",
			wantModel: &model.Exists{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operand: &model.List{
						List:       []model.IExpression{model.NewLiteral("1", types.Integer)},
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
					},
				},
			},
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "Empty List",
			cql:        "exists({})",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "List Of Nulls",
			cql:        "exists({null})",
			wantResult: newOrFatal(t, false),
		},
		{
			name: "Null",
			cql:  "exists(null)",
			wantModel: &model.Exists{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operand: &model.As{
						UnaryExpression: &model.UnaryExpression{
							Expression: model.ResultType(&types.List{ElementType: types.Any}),
							Operand:    model.NewLiteral("null", types.Any),
						},
						AsTypeSpecifier: &types.List{ElementType: types.Any},
					},
				},
			},
			wantResult: newOrFatal(t, false),
		},
		{
			name: "Non Function Syntax",
			cql:  "exists {1}",
			wantModel: &model.Exists{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
			wantResult: newOrFatal(t, true),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}

		})
	}
}

func TestInList(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantModel  model.IExpression
		wantResult result.Value
	}{
		{
			name: "In List",
			cql:  "1 in {1, 2}",
			wantModel: &model.In{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("1", types.Integer),
								model.NewLiteral("2", types.Integer),
							},
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "null in list",
			cql:        "null in {1, null}",
			wantResult: newOrFatal(t, true),
		},
		{
			name:       "Operand not in list",
			cql:        "3 in {1, 2}",
			wantResult: newOrFatal(t, false),
		},
		{
			name:       "Functional syntax: In list",
			cql:        "In(1, {1, 2})",
			wantResult: newOrFatal(t, true),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}

		})
	}
}

func TestDistinctList(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantModel  model.IExpression
		wantResult result.Value
	}{
		{
			name: "Distinct list",
			cql:  "distinct {1, 2, 1}",
			wantModel: &model.Distinct{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
							model.NewLiteral("2", types.Integer),
							model.NewLiteral("1", types.Integer),
						},
					},
					Expression: model.ResultType(&types.List{ElementType: types.Integer}),
				},
			},
			wantResult: newOrFatal(t, result.List{
				Value: []result.Value{
					newOrFatal(t, int32(1)),
					newOrFatal(t, int32(2)),
				},
				StaticType: &types.List{ElementType: types.Integer},
			}),
		},
		{
			name: "Distinct list with null",
			cql:  "distinct {1, null}",
			wantResult: newOrFatal(t, result.List{
				Value: []result.Value{
					newOrFatal(t, int32(1)),
					newOrFatal(t, nil),
				},
				StaticType: &types.List{ElementType: types.Integer},
			}),
		},
		{
			name: "distinct list with no duplicates",
			cql:  "distinct {1, 2, 3}",
			wantResult: newOrFatal(t, result.List{
				Value: []result.Value{
					newOrFatal(t, int32(1)),
					newOrFatal(t, int32(2)),
					newOrFatal(t, int32(3)),
				},
				StaticType: &types.List{ElementType: types.Integer},
			}),
		},
		{
			name: "Functional syntax: In list",
			cql:  "Distinct({1, 2})",
			wantResult: newOrFatal(t, result.List{
				Value: []result.Value{
					newOrFatal(t, int32(1)),
					newOrFatal(t, int32(2)),
				},
				StaticType: &types.List{ElementType: types.Integer},
			}),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}

		})
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		name                 string
		cql                  string
		wantModel            model.IExpression
		wantResult           result.Value
		wantSourceExpression model.IExpression
		wantSourceValues     []result.Value
	}{
		{
			name: "First({1, 2}) = 1",
			cql:  "First({1, 2})",
			wantModel: &model.First{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
							model.NewLiteral("2", types.Integer),
						},
					},
				},
			},
			wantResult: newOrFatal(t, int32(1)),
			wantSourceExpression: &model.First{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
							model.NewLiteral("2", types.Integer),
						},
					},
				},
			},
			wantSourceValues: []result.Value{
				newOrFatal(t, result.List{Value: []result.Value{newOrFatal(t, int32(1)), newOrFatal(t, int32(2))}, StaticType: &types.List{ElementType: types.Integer}}),
			},
		},
		{
			name:       "First({}) = null",
			cql:        "First({})",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "First(null) = null",
			cql:        "First(null)",
			wantResult: newOrFatal(t, nil),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			gotResult := getTESTRESULTWithSources(t, results)
			if diff := cmp.Diff(tc.wantResult, gotResult, protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}
			if diff := cmp.Diff(tc.wantSourceExpression, gotResult.SourceExpression(), protocmp.Transform()); tc.wantSourceExpression != nil && diff != "" {
				t.Errorf("Eval SourceExpression diff (-want +got)\n%v", diff)
			}
			if diff := cmp.Diff(tc.wantSourceValues, gotResult.SourceValues(), protocmp.Transform()); tc.wantSourceValues != nil && diff != "" {
				t.Errorf("Eval SourceValues diff (-want +got)\n%v", diff)
			}
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name                 string
		cql                  string
		wantModel            model.IExpression
		wantResult           result.Value
		wantSourceExpression model.IExpression
		wantSourceValues     []result.Value
	}{
		{
			name: "Last({1, 2}) = 2",
			cql:  "Last({1, 2})",
			wantModel: &model.Last{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
							model.NewLiteral("2", types.Integer),
						},
					},
				},
			},
			wantResult: newOrFatal(t, int32(2)),
			wantSourceExpression: &model.Last{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
							model.NewLiteral("2", types.Integer),
						},
					},
				},
			},
			wantSourceValues: []result.Value{
				newOrFatal(t, result.List{Value: []result.Value{newOrFatal(t, int32(1)), newOrFatal(t, int32(2))}, StaticType: &types.List{ElementType: types.Integer}}),
			},
		},
		{
			name:       "Last({}) = null",
			cql:        "Last({})",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Last(null) = null",
			cql:        "Last(null)",
			wantResult: newOrFatal(t, nil),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			gotResult := getTESTRESULTWithSources(t, results)
			if diff := cmp.Diff(tc.wantResult, gotResult, protocmp.Transform()); diff != "" {
				t.Errorf("Eval returned diff (-want +got)\n%v", diff)
			}
			if diff := cmp.Diff(tc.wantSourceExpression, gotResult.SourceExpression(), protocmp.Transform()); tc.wantSourceExpression != nil && diff != "" {
				t.Errorf("Eval SourceExpression diff (-want +got)\n%v", diff)
			}
			if diff := cmp.Diff(tc.wantSourceValues, gotResult.SourceValues(), protocmp.Transform()); tc.wantSourceValues != nil && diff != "" {
				t.Errorf("Eval SourceValues diff (-want +got)\n%v", diff)
			}
		})
	}
}

func TestSingletonFrom(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantModel  model.IExpression
		wantResult result.Value
	}{
		{
			name: "singleton from {1} = 1",
			cql:  "singleton from {1}",
			wantModel: &model.SingletonFrom{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
						},
					},
					Expression: model.ResultType(types.Integer),
				},
			},
			wantResult: newOrFatal(t, 1),
		},
		{
			name: "Functional syntax: SingletonFrom({1}) = 1",
			cql:  "SingletonFrom({1})",
			wantModel: &model.SingletonFrom{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
						},
					},
					Expression: model.ResultType(types.Integer),
				},
			},
			wantResult: newOrFatal(t, 1),
		},
		{
			name:       "Empty list input returns null",
			cql:        "singleton from {}",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Functional syntax with empty list",
			cql:        "SingletonFrom({})",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Null input returns null",
			cql:        "singleton from null",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Functional syntax with null",
			cql:        "SingletonFrom(null)",
			wantResult: newOrFatal(t, nil),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			gotResult := getTESTRESULT(t, results)
			if diff := cmp.Diff(tc.wantResult, gotResult, protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}
		})
	}
}

func TestListOperatorSingletonFrom_Error(t *testing.T) {
	tests := []struct {
		name                string
		cql                 string
		wantModel           model.IExpression
		wantEvalErrContains string
	}{
		{
			name:                "Length Greater Than 1",
			cql:                 "singleton from {1, 2}",
			wantEvalErrContains: "length 0 or 1",
		},
		{
			name:                "Length Greater Than 1 with functional syntax",
			cql:                 "SingletonFrom({1, 2})",
			wantEvalErrContains: "length 0 or 1",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			_, err = interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err == nil {
				t.Fatalf("Evaluate Expression expected an error to be returned, got nil instead")
			}
			if !strings.Contains(err.Error(), tc.wantEvalErrContains) {
				t.Errorf("Unexpected evaluation error contents. got: %v, want contains: %v", err.Error(), tc.wantEvalErrContains)
			}
		})
	}
}

func TestIndexerList(t *testing.T) {
	tests := []struct {
		name       string
		cql        string
		wantModel  model.IExpression
		wantResult result.Value
	}{
		{
			name: "Indexer [] syntax for List<Integer>",
			cql:  "{1, 2}[1]",
			wantModel: &model.Indexer{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewList([]string{"1", "2"}, types.Integer),
						model.NewLiteral("1", types.Integer),
					},
				},
			},
			wantResult: newOrFatal(t, 2),
		},
		{
			name: "Indexer [] syntax for List<String>",
			cql:  "{'a', 'b'}[1]",
			wantModel: &model.Indexer{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.String),
					Operands: []model.IExpression{
						model.NewList([]string{"a", "b"}, types.String),
						model.NewLiteral("1", types.Integer),
					},
				},
			},
			wantResult: newOrFatal(t, "b"),
		},
		{
			name:       "Indexer functional form",
			cql:        "Indexer({1, 2}, 1)",
			wantResult: newOrFatal(t, 2),
		},
		{
			name:       "Indexer with index too large",
			cql:        "{1, 2}[100]",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Indexer with index smaller than 0",
			cql:        "{1, 2}[-100]",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Indexer on null",
			cql:        "(null as List<Integer>)[1]",
			wantResult: newOrFatal(t, nil),
		},
		{
			name:       "Indexer with null index",
			cql:        "{1, 2}[null as Integer]",
			wantResult: newOrFatal(t, nil),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := newFHIRParser(t)
			parsedLibs, err := p.Libraries(context.Background(), wrapInLib(t, tc.cql), parser.Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantModel, getTESTRESULTModel(t, parsedLibs)); tc.wantModel != nil && diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}

			results, err := interpreter.Eval(context.Background(), parsedLibs, defaultInterpreterConfig(t, p))
			if err != nil {
				t.Fatalf("Eval returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.wantResult, getTESTRESULT(t, results), protocmp.Transform()); diff != "" {
				t.Errorf("Eval diff (-want +got)\n%v", diff)
			}

		})
	}
}
