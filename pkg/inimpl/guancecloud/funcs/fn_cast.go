// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"fmt"

	"github.com/GuanceCloud/ppl/pkg/ast"
	"github.com/GuanceCloud/ppl/pkg/engine/runtime"
	"github.com/GuanceCloud/ppl/pkg/inimpl/guancecloud/input"
)

func CastChecking(ctx *runtime.Context, funcExpr *ast.CallExpr) *runtime.RuntimeError {
	if len(funcExpr.Param) != 2 {
		return runtime.NewRunError(ctx, fmt.Sprintf(
			"func `%s' expected 2 args", funcExpr.Name), funcExpr.NamePos)
	}
	if _, err := getKeyName(funcExpr.Param[0]); err != nil {
		return runtime.NewRunError(ctx, err.Error(), funcExpr.Param[1].StartPos())
	}

	switch funcExpr.Param[1].NodeType { //nolint:exhaustive
	case ast.TypeStringLiteral:
		switch funcExpr.Param[1].StringLiteral.Val {
		case "bool", "int", "float", "str", "string":
		default:
			return runtime.NewRunError(ctx, fmt.Sprintf("unsupported data type: %s",
				funcExpr.Param[1].StringLiteral.Val), funcExpr.Param[1].StartPos())
		}
	default:
		return runtime.NewRunError(ctx, fmt.Sprintf("param type expect StringLiteral, got `%s'",
			funcExpr.Param[1].NodeType), funcExpr.Param[1].StartPos())
	}
	return nil
}

func Cast(ctx *runtime.Context, funcExpr *ast.CallExpr) *runtime.RuntimeError {
	if len(funcExpr.Param) != 2 {
		return runtime.NewRunError(ctx, fmt.Sprintf(
			"func `%s' expected 2 args", funcExpr.Name), funcExpr.NamePos)
	}

	key, err := getKeyName(funcExpr.Param[0])
	if err != nil {
		return runtime.NewRunError(ctx, err.Error(), funcExpr.Param[0].StartPos())
	}

	var castType string

	switch funcExpr.Param[1].NodeType { //nolint:exhaustive
	case ast.TypeStringLiteral:
		castType = funcExpr.Param[1].StringLiteral.Val
	default:
		return runtime.NewRunError(ctx, fmt.Sprintf("param type expect StringLiteral, got `%s'",
			funcExpr.Param[1].NodeType), funcExpr.Param[1].StartPos())
	}

	v, err := ctx.GetKey(key)
	if err != nil {
		l.Debug(err)
		return nil
	}

	val, dtype := doCast(v.Value, castType)
	if err = addKey2PtWithVal(ctx.InData(), key, val, dtype,
		input.KindPtDefault); err != nil {
		l.Debug(err)
		return nil
	}

	return nil
}
