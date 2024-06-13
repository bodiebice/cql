// Copyright 2024 Google LLC
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

package parser

import (
	"context"
	"testing"

	"github.com/google/cql/model"
	"github.com/google/cql/types"
	"github.com/google/go-cmp/cmp"
	"github.com/lithammer/dedent"
)

func TestBuiltInFunctions(t *testing.T) {
	tests := []struct {
		name string
		cql  string
		want model.IExpression
	}{
		// LOGICAL OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#logical-operators-3
		{
			name: "Not",
			cql:  "Not(true)",
			want: &model.Not{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("true", types.Boolean),
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "And",
			cql:  "And(true, false)",
			want: &model.And{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operands: []model.IExpression{
						model.NewLiteral("true", types.Boolean),
						model.NewLiteral("false", types.Boolean),
					},
				},
			},
		},
		{
			name: "Or",
			cql:  "Or(false, true)",
			want: &model.Or{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operands: []model.IExpression{
						model.NewLiteral("false", types.Boolean),
						model.NewLiteral("true", types.Boolean),
					},
				},
			},
		},
		{
			name: "XOr",
			cql:  "Xor(true, false)",
			want: &model.XOr{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operands: []model.IExpression{
						model.NewLiteral("true", types.Boolean),
						model.NewLiteral("false", types.Boolean),
					},
				},
			},
		},
		{
			name: "Implies",
			cql:  "Implies(false, true)",
			want: &model.Implies{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operands: []model.IExpression{
						model.NewLiteral("false", types.Boolean),
						model.NewLiteral("true", types.Boolean),
					},
				},
			},
		},
		// TYPE OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#type-operators-1
		{
			name: "CanConvertQuantity",
			cql:  dedent.Dedent(`CanConvertQuantity(1 'cm', 'm')`),
			want: &model.CanConvertQuantity{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						&model.Quantity{Value: 1, Unit: "cm", Expression: model.ResultType(types.Quantity)},
						model.NewLiteral("m", types.String),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "ToBoolean",
			cql:  "ToBoolean(5)",
			want: &model.ToBoolean{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Boolean),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToDateTime with Date",
			cql:  "ToDateTime(@1999-10-01)",
			want: &model.ToDateTime{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.DateTime),
					Operand:    model.NewLiteral("@1999-10-01", types.Date),
				},
			},
		},
		{
			name: "ToDateTime with string",
			cql:  "ToDateTime('@2014-01-25T14:30:14.559')",
			want: &model.ToDateTime{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.DateTime),
					Operand:    model.NewLiteral("@2014-01-25T14:30:14.559", types.String),
				},
			},
		},
		{
			name: "ToDate with DateTime",
			cql:  "ToDate(@2014-01-25T14:30:14.559)",
			want: &model.ToDate{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Date),
					Operand:    model.NewLiteral("@2014-01-25T14:30:14.559", types.DateTime),
				},
			},
		},
		{
			name: "ToDate with string",
			cql:  "ToDate('@2014-01-25')",
			want: &model.ToDate{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Date),
					Operand:    model.NewLiteral("@2014-01-25", types.String),
				},
			},
		},
		{
			name: "ToDecimal",
			cql:  "ToDecimal(5)",
			want: &model.ToDecimal{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Decimal),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToLong",
			cql:  "ToLong(5)",
			want: &model.ToLong{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Long),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToInteger",
			cql:  "ToInteger(5)",
			want: &model.ToInteger{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToQuantity",
			cql:  "ToQuantity(5)",
			want: &model.ToQuantity{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Quantity),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToConcept",
			cql:  "ToConcept(Code{code: 'foo', system: 'bar', version: '1.0', display: 'severed leg' })",
			want: &model.ToConcept{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Concept),
					Operand: &model.Instance{
						ClassType: types.Code,
						Elements: []*model.InstanceElement{
							&model.InstanceElement{Name: "code", Value: model.NewLiteral("foo", types.String)},
							&model.InstanceElement{Name: "system", Value: model.NewLiteral("bar", types.String)},
							&model.InstanceElement{Name: "version", Value: model.NewLiteral("1.0", types.String)},
							&model.InstanceElement{Name: "display", Value: model.NewLiteral("severed leg", types.String)},
						},
						Expression: model.ResultType(types.Code),
					},
				},
			},
		},
		{
			name: "ToString",
			cql:  "ToString(5)",
			want: &model.ToString{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.String),
					Operand:    model.NewLiteral("5", types.Integer),
				},
			},
		},
		{
			name: "ToTime",
			cql:  "ToTime('hello')",
			want: &model.ToTime{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Time),
					Operand:    model.NewLiteral("hello", types.String),
				},
			},
		},
		// NULLOGICAL OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#nullological-operators-3
		{
			name: "Coalesce 2 Operands",
			cql:  "Coalesce(4, 5.0)",
			want: &model.Coalesce{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						&model.ToDecimal{
							UnaryExpression: &model.UnaryExpression{
								Expression: model.ResultType(types.Decimal),
								Operand:    model.NewLiteral("4", types.Integer),
							},
						},
						model.NewLiteral("5.0", types.Decimal),
					},
					Expression: model.ResultType(types.Decimal),
				},
			},
		},
		{
			name: "Coalesce 3 Operands",
			cql:  "Coalesce(4, null, 5)",
			want: &model.Coalesce{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("4", types.Integer),
						&model.As{
							UnaryExpression: &model.UnaryExpression{
								Expression: model.ResultType(types.Integer),
								Operand:    model.NewLiteral("null", types.Any),
							},
							AsTypeSpecifier: types.Integer,
						},
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Coalesce 4 Operands",
			cql:  "Coalesce(3, 4, 5, 6)",
			want: &model.Coalesce{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("3", types.Integer),
						model.NewLiteral("4", types.Integer),
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("6", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Coalesce 5 Operands",
			cql:  "Coalesce(3, 4, 5, 6, 7)",
			want: &model.Coalesce{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("3", types.Integer),
						model.NewLiteral("4", types.Integer),
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("6", types.Integer),
						model.NewLiteral("7", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Coalesce List Operand",
			cql:  "Coalesce({4, 5})",
			want: &model.Coalesce{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("4", types.Integer),
								model.NewLiteral("5", types.Integer),
							},
						},
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "IsNull",
			cql:  "IsNull(5)",
			want: &model.IsNull{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("5", types.Integer),
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "IsFalse",
			cql:  "IsFalse(false)",
			want: &model.IsFalse{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("false", types.Boolean),
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "IsTrue",
			cql:  "IsTrue(false)",
			want: &model.IsTrue{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("false", types.Boolean),
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		// COMPARISON OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#comparison-operators-4
		{
			name: "Equal",
			cql:  "Equal(5, 5)",
			want: &model.Equal{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Equivalent",
			cql:  "Equivalent(5, 5)",
			want: &model.Equivalent{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Less",
			cql:  "Less(5, 5)",
			want: &model.Less{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Greater",
			cql:  "Greater(5, 5)",
			want: &model.Greater{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "LessOrEqual",
			cql:  "LessOrEqual(5, 5)",
			want: &model.LessOrEqual{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "GreaterOrEqual",
			cql:  "GreaterOrEqual(5, 5)",
			want: &model.GreaterOrEqual{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("5", types.Integer),
						model.NewLiteral("5", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		// ARITHMETIC OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#arithmetic-operators-4
		{
			name: "Arithmetic Addition",
			cql:  "Add(1, 2)",
			want: &model.Add{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewLiteral("2", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Arithmetic Subtraction",
			cql:  "Subtract(1, 2)",
			want: &model.Subtract{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewLiteral("2", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Arithmetic Multiplication",
			cql:  "Multiply(1, 2)",
			want: &model.Multiply{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewLiteral("2", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Arithmetic Modulo with different types",
			cql:  "Modulo(40L, 3)",
			want: &model.Modulo{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("40L", types.Long),
						&model.ToLong{
							UnaryExpression: &model.UnaryExpression{
								Operand:    model.NewLiteral("3", types.Integer),
								Expression: model.ResultType(types.Long),
							},
						},
					},
					Expression: model.ResultType(types.Long),
				},
			},
		},
		{
			name: "Arithmetic Truncated Divide",
			cql:  "TruncatedDivide(40, 3)",
			want: &model.TruncatedDivide{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("40", types.Integer),
						model.NewLiteral("3", types.Integer),
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Negate",
			cql:  "Negate(4)",
			want: &model.Negate{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("4", types.Integer),
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "Predecessor for Date",
			cql:  "Predecessor(@2023)",
			want: &model.Predecessor{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("@2023", types.Date),
					Expression: model.ResultType(types.Date),
				},
			},
		},
		{
			name: "Successor for Integer",
			cql:  "Successor(41)",
			want: &model.Successor{
				UnaryExpression: &model.UnaryExpression{
					Operand:    model.NewLiteral("41", types.Integer),
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		// STRING OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#string-operators-3
		{
			name: "Add Strings",
			cql:  "Add('Hi', 'Hello')",
			want: &model.Concatenate{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("Hi", types.String),
						model.NewLiteral("Hello", types.String),
					},
					Expression: model.ResultType(types.String),
				},
			},
		},
		{
			name: "Concatenate",
			cql:  "Concatenate('Hi', 'Hello')",
			want: &model.Concatenate{
				NaryExpression: &model.NaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("Hi", types.String),
						model.NewLiteral("Hello", types.String),
					},
					Expression: model.ResultType(types.String),
				},
			},
		},
		// DATE AND TIME OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#datetime-operators-2
		{
			name: "After",
			cql:  "After(1, Interval[2, 3])",
			want: &model.After{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewInclusiveInterval("2", "3", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "After With Precision",
			cql:  "AfterDays(@2023-01-01, @2023-01-01)",
			want: &model.After{
				Precision: model.DAY,
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2023-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Before",
			cql:  "Before(1, Interval[2, 3])",
			want: &model.Before{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewInclusiveInterval("2", "3", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Before With Precision",
			cql:  "BeforeDays(@2023-01-01, @2023-01-01)",
			want: &model.Before{
				Precision: model.DAY,
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2023-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Date",
			cql:  "Date(2014, 10, 3)",
			want: &model.Date{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{model.NewLiteral("2014", types.Integer), model.NewLiteral("10", types.Integer), model.NewLiteral("3", types.Integer)},
					Expression: model.ResultType(types.Date),
				},
			},
		},
		{
			name: "DateTime",
			cql:  "DateTime(2014, 10, 3, 6, 30, 15, 500, 7.3)",
			want: &model.DateTime{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{model.NewLiteral("2014", types.Integer), model.NewLiteral("10", types.Integer), model.NewLiteral("3", types.Integer), model.NewLiteral("6", types.Integer), model.NewLiteral("30", types.Integer), model.NewLiteral("15", types.Integer), model.NewLiteral("500", types.Integer), model.NewLiteral("7.3", types.Decimal)},
					Expression: model.ResultType(types.DateTime),
				},
			},
		},
		{
			name: "DifferenceBetween",
			cql:  "DifferenceBetweenYears(@2023-01-01, @2023-01-01)",
			want: &model.DifferenceBetween{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2023-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
					Expression: model.ResultType(types.Integer),
				},
				Precision: model.YEAR,
			},
		},
		{
			name: "Now()",
			cql:  "Now()",
			want: &model.Now{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{},
					Expression: model.ResultType(types.DateTime),
				},
			},
		},
		{
			name: "TimeOfDay()",
			cql:  "TimeOfDay()",
			want: &model.TimeOfDay{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{},
					Expression: model.ResultType(types.Time),
				},
			},
		},
		{
			name: "SameOrAfter",
			cql:  "SameOrAfter(1, Interval[2, 3])",
			want: &model.SameOrAfter{
				BinaryExpression: &model.BinaryExpression{

					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewInclusiveInterval("2", "3", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "SameOrAfter With Precision",
			cql:  "SameOrAfterDays(@2023-01-01, @2023-01-01)",
			want: &model.SameOrAfter{
				Precision: model.DAY,
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2023-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "SameOrBefore",
			cql:  "SameOrBefore(1, Interval[2, 3])",
			want: &model.SameOrBefore{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						model.NewInclusiveInterval("2", "3", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "SameOrBefore With Precision",
			cql:  "SameOrBeforeDays(@2023-01-01, @2023-01-01)",
			want: &model.SameOrBefore{
				Precision: model.DAY,
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2023-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "Time",
			cql:  "Time(6, 30, 15, 500)",
			want: &model.Time{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{model.NewLiteral("6", types.Integer), model.NewLiteral("30", types.Integer), model.NewLiteral("15", types.Integer), model.NewLiteral("500", types.Integer)},
					Expression: model.ResultType(types.Time),
				},
			},
		},
		{
			name: "Today()",
			cql:  "Today()",
			want: &model.Today{
				NaryExpression: &model.NaryExpression{
					Operands:   []model.IExpression{},
					Expression: model.ResultType(types.Date),
				},
			},
		},
		// INTERVAL OPERATORS - https://cql.hl7.org/04-logicalspecification.html#interval-operators
		{
			name: "Contains",
			cql:  "Contains({3}, 1)",
			want: &model.Contains{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("3", types.Integer),
							},
						},
						model.NewLiteral("1", types.Integer),
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "End",
			cql:  "End(Interval[1, 4])",
			want: &model.End{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.Interval{
						Low:           model.NewLiteral("1", types.Integer),
						High:          model.NewLiteral("4", types.Integer),
						Expression:    model.ResultType(&types.Interval{PointType: types.Integer}),
						LowInclusive:  true,
						HighInclusive: true,
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		{
			name: "In",
			cql:  "In(1, {3})",
			want: &model.In{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("1", types.Integer),
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("3", types.Integer),
							},
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "InMonths",
			cql:  "InMonths(@2020-03, Interval[@2020-03, @2022-03])",
			want: &model.In{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2020-03", types.Date),
						&model.Interval{
							Low:           model.NewLiteral("@2020-03", types.Date),
							High:          model.NewLiteral("@2022-03", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
				Precision: model.MONTH,
			},
		},
		{
			name: "InSeconds",
			cql:  "InSeconds(@2024-03-31T00:00:05.000Z, Interval[@2024-03-31T00:00:00.000Z, @2025-03-31T00:00:00.000Z])",
			want: &model.In{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2024-03-31T00:00:05.000Z", types.DateTime),
						&model.Interval{
							Low:           model.NewLiteral("@2024-03-31T00:00:00.000Z", types.DateTime),
							High:          model.NewLiteral("@2025-03-31T00:00:00.000Z", types.DateTime),
							Expression:    model.ResultType(&types.Interval{PointType: types.DateTime}),
							LowInclusive:  true,
							HighInclusive: true,
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
				Precision: model.SECOND,
			},
		},
		{
			name: "IncludedIn for point type",
			cql:  "IncludedIn(@2010, Interval[@2010, @2020])",
			want: &model.In{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewLiteral("@2010", types.Date),
						&model.Interval{
							Low:           model.NewLiteral("@2010", types.Date),
							High:          model.NewLiteral("@2020", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "IncludedIn interval overload",
			cql:  "IncludedIn(Interval[@2015, @2016], Interval[@2010, @2020])",
			want: &model.IncludedIn{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						&model.Interval{
							Low:           model.NewLiteral("@2015", types.Date),
							High:          model.NewLiteral("@2016", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
						&model.Interval{
							Low:           model.NewLiteral("@2010", types.Date),
							High:          model.NewLiteral("@2020", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
			},
		},
		{
			name: "IncludedInYears interval overload",
			cql:  "IncludedInYears(Interval[@2015, @2016], Interval[@2010, @2020])",
			want: &model.IncludedIn{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						&model.Interval{
							Low:           model.NewLiteral("@2015", types.Date),
							High:          model.NewLiteral("@2016", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
						&model.Interval{
							Low:           model.NewLiteral("@2010", types.Date),
							High:          model.NewLiteral("@2020", types.Date),
							Expression:    model.ResultType(&types.Interval{PointType: types.Date}),
							LowInclusive:  true,
							HighInclusive: true,
						},
					},
					Expression: model.ResultType(types.Boolean),
				},
				Precision: model.YEAR,
			},
		},
		{
			name: "Start",
			cql:  "Start(Interval[1, 4])",
			want: &model.Start{
				UnaryExpression: &model.UnaryExpression{
					Operand: &model.Interval{
						Low:           model.NewLiteral("1", types.Integer),
						High:          model.NewLiteral("4", types.Integer),
						Expression:    model.ResultType(&types.Interval{PointType: types.Integer}),
						LowInclusive:  true,
						HighInclusive: true,
					},
					Expression: model.ResultType(types.Integer),
				},
			},
		},
		// LIST OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#list-operators-2
		{
			name: "Except",
			cql:  "Except({1}, {1})",
			want: &model.Except{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{&model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
						},
					},
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("1", types.Integer),
							},
						},
					},
					Expression: model.ResultType(&types.List{ElementType: types.Integer}),
				},
			},
		},
		{
			name: "First",
			cql:  "First({1})",
			want: &model.First{
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
		},
		{
			name: "Intersect",
			cql:  "Intersect({1}, {1})",
			want: &model.Intersect{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{&model.List{
						Expression: model.ResultType(&types.List{ElementType: types.Integer}),
						List: []model.IExpression{
							model.NewLiteral("1", types.Integer),
						},
					},
						&model.List{
							Expression: model.ResultType(&types.List{ElementType: types.Integer}),
							List: []model.IExpression{
								model.NewLiteral("1", types.Integer),
							},
						},
					},
					Expression: model.ResultType(&types.List{ElementType: types.Integer}),
				},
			},
		},
		{
			name: "Last",
			cql:  "Last({1})",
			want: &model.Last{
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
		},
		{
			name: "SingletonFrom",
			cql:  "SingletonFrom({1})",
			want: &model.SingletonFrom{
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
		},
		{
			name: "Union",
			cql:  "Union({1}, {'hi'})",
			want: &model.Union{
				BinaryExpression: &model.BinaryExpression{
					Operands: []model.IExpression{
						model.NewList([]string{"1"}, types.Integer),
						model.NewList([]string{"hi"}, types.String),
					},
					Expression: model.ResultType(&types.List{ElementType: &types.Choice{ChoiceTypes: []types.IType{types.Integer, types.String}}}),
				},
			},
		},
		// CLINICAL OPERATORS - https://cql.hl7.org/09-b-cqlreference.html#clinical-operators-3
		{
			name: "AgeInYears",
			cql:  "AgeInYears()",
			want: &model.CalculateAge{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.ToDateTime{
						UnaryExpression: &model.UnaryExpression{
							Expression: model.ResultType(types.DateTime),
							Operand: &model.Property{
								Expression: model.ResultType(types.Date),
								Source: &model.Property{
									Expression: model.ResultType(&types.Named{TypeName: "FHIR.date"}),
									Source: &model.ExpressionRef{
										Expression: model.ResultType(&types.Named{TypeName: "FHIR.Patient"}),
										Name:       "Patient",
									},
									Path: "birthDate",
								},
								Path: "value",
							},
						},
					},
				},
				Precision: model.YEAR,
			},
		},
		{
			name: "AgeInDays",
			cql:  "AgeInDays()",
			want: &model.CalculateAge{
				UnaryExpression: &model.UnaryExpression{
					Expression: model.ResultType(types.Integer),
					Operand: &model.ToDateTime{
						UnaryExpression: &model.UnaryExpression{
							Expression: model.ResultType(types.DateTime),
							Operand: &model.Property{
								Expression: model.ResultType(types.Date),
								Source: &model.Property{
									Expression: model.ResultType(&types.Named{TypeName: "FHIR.date"}),
									Source: &model.ExpressionRef{
										Expression: model.ResultType(&types.Named{TypeName: "FHIR.Patient"}),
										Name:       "Patient",
									},
									Path: "birthDate",
								},
								Path: "value",
							},
						},
					},
				},
				Precision: model.DAY,
			},
		},
		{
			name: "AgeInDaysAt Date Overload",
			cql:  "AgeInDaysAt(@2023-01-01)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						&model.Property{
							Expression: model.ResultType(types.Date),
							Source: &model.Property{
								Expression: model.ResultType(&types.Named{TypeName: "FHIR.date"}),
								Source: &model.ExpressionRef{
									Expression: model.ResultType(&types.Named{TypeName: "FHIR.Patient"}),
									Name:       "Patient",
								},
								Path: "birthDate",
							},
							Path: "value",
						},
						model.NewLiteral("@2023-01-01", types.Date),
					},
				},
				Precision: model.DAY,
			},
		},
		{
			name: "AgeInDaysAt DateTime Overload",
			cql:  "AgeInDaysAt(@2022-01-01T12:00:00.000)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						&model.ToDateTime{
							UnaryExpression: &model.UnaryExpression{
								Expression: model.ResultType(types.DateTime),
								Operand: &model.Property{
									Expression: model.ResultType(types.Date),
									Source: &model.Property{
										Expression: model.ResultType(&types.Named{TypeName: "FHIR.date"}),
										Source: &model.ExpressionRef{
											Expression: model.ResultType(&types.Named{TypeName: "FHIR.Patient"}),
											Name:       "Patient",
										},
										Path: "birthDate",
									},
									Path: "value",
								},
							},
						},
						model.NewLiteral("@2022-01-01T12:00:00.000", types.DateTime),
					},
				},
				Precision: model.DAY,
			},
		},
		{
			name: "AgeInMonthsAt",
			cql:  "AgeInMonthsAt(@2023-01-01)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						&model.Property{
							Expression: model.ResultType(types.Date),
							Source: &model.Property{
								Expression: model.ResultType(&types.Named{TypeName: "FHIR.date"}),
								Source: &model.ExpressionRef{
									Expression: model.ResultType(&types.Named{TypeName: "FHIR.Patient"}),
									Name:       "Patient",
								},
								Path: "birthDate",
							},
							Path: "value",
						},
						model.NewLiteral("@2023-01-01", types.Date),
					},
				},
				Precision: model.MONTH,
			},
		},
		{
			name: "CalculateAgeInDaysAt",
			cql:  "CalculateAgeInDaysAt(@2022-01-01, @2023-01-01)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewLiteral("@2022-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
				},
				Precision: model.DAY,
			},
		},
		{
			name: "CalculateAgeInYearsAt",
			cql:  "CalculateAgeInYearsAt(@2022-01-01, @2023-01-01)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewLiteral("@2022-01-01", types.Date),
						model.NewLiteral("@2023-01-01", types.Date),
					},
				},
				Precision: model.YEAR,
			},
		},
		{
			name: "CalculateAgeInHoursAt",
			cql:  "CalculateAgeInHoursAt(@2022-01-01T12:00:00.000, @2023-01-01T12:00:00.000)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewLiteral("@2022-01-01T12:00:00.000", types.DateTime),
						model.NewLiteral("@2023-01-01T12:00:00.000", types.DateTime),
					},
				},
				Precision: model.HOUR,
			},
		},
		{
			name: "CalculateAgeInMinutesAt",
			cql:  "CalculateAgeInMinutesAt(@2022-01-01T12:00:00.000, @2023-01-01T12:00:00.000)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewLiteral("@2022-01-01T12:00:00.000", types.DateTime),
						model.NewLiteral("@2023-01-01T12:00:00.000", types.DateTime),
					},
				},
				Precision: model.MINUTE,
			},
		},
		{
			name: "CalculateAgeInSecondsAt",
			cql:  "CalculateAgeInSecondsAt(@2022-01-01T12:00:00.000, @2023-01-01T12:00:00.000)",
			want: &model.CalculateAgeAt{
				BinaryExpression: &model.BinaryExpression{
					Expression: model.ResultType(types.Integer),
					Operands: []model.IExpression{
						model.NewLiteral("@2022-01-01T12:00:00.000", types.DateTime),
						model.NewLiteral("@2023-01-01T12:00:00.000", types.DateTime),
					},
				},
				Precision: model.SECOND,
			},
		},
		// ERRORS AND MESSAGES - https://cql.hl7.org/09-b-cqlreference.html#errors-and-messaging
		{
			name: "Message with all args",
			cql:  "Message(1.0, true, '100', 'Message', 'Test Message')",
			want: &model.Message{
				Source:     model.NewLiteral("1.0", types.Decimal),
				Condition:  model.NewLiteral("true", types.Boolean),
				Code:       model.NewLiteral("100", types.String),
				Severity:   model.NewLiteral("Message", types.String),
				Message:    model.NewLiteral("Test Message", types.String),
				Expression: model.ResultType(types.Decimal),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parsedLibs, err := newFHIRParser(t).Libraries(context.Background(), wrapInLib(t, tc.cql), Config{})
			if err != nil {
				t.Fatalf("Parse returned unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.want, getTESTRESULTModel(t, parsedLibs)); diff != "" {
				t.Errorf("Parse diff (-want +got):\n%s", diff)
			}
		})
	}
}
