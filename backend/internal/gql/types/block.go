package types

import (
	"github.com/graphql-go/graphql"
	"strconv"
	"strings"
	"encoding/json"
	"github.com/graphql-go/graphql/language/ast"

)
var BlockType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Block",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.String},
		"pageId":      &graphql.Field{Type: graphql.String},
		"parentblockId": &graphql.Field{Type: graphql.String},
		"Type":   &graphql.Field{Type: BlockTypeEnum},
		"data":        &graphql.Field{Type: JSONScalar},
	},
})
var BlockTypeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "BlockType",
	Values: graphql.EnumValueConfigMap{
		"PARAGRAPH": &graphql.EnumValueConfig{Value: "PARAGRAPH"},
		"HEADING1":  &graphql.EnumValueConfig{Value: "HEADING1"},
		"HEADING2":  &graphql.EnumValueConfig{Value: "HEADING2"},
		"TODO":      &graphql.EnumValueConfig{Value: "TODO"},
		"BULLET":    &graphql.EnumValueConfig{Value: "BULLET"},
		"NUMBER":    &graphql.EnumValueConfig{Value: "NUMBER"},
		"QUOTE":     &graphql.EnumValueConfig{Value: "QUOTE"},
		"CODE":      &graphql.EnumValueConfig{Value: "CODE"},
		"IMAGE":     &graphql.EnumValueConfig{Value: "IMAGE"},
		"DIVIDER":   &graphql.EnumValueConfig{Value: "DIVIDER"},
	},
})


var JSONScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "JSON",
	Description: "Arbitrary JSON",
	Serialize: func(value any) any {
		// ce qui sort vers le client
		switch v := value.(type) {
		case json.RawMessage:
			var out any
			_ = json.Unmarshal(v, &out)
			return out
		case []byte:
			var out any
			_ = json.Unmarshal(v, &out)
			return out
		default:
			// map[string]any, []any, string, int, etc.
			return value
		}
	},
	ParseValue: func(value any) any {
		// ce qui entre via variables
		return value
	},
	ParseLiteral: func(valueAST ast.Value) any {
		// ce qui entre directement dans la query (sans variables)
		switch v := valueAST.(type) {
		case *ast.ObjectValue:
			m := map[string]any{}
			for _, f := range v.Fields {
				m[f.Name.Value] = parseASTValue(f.Value)
			}
			return m
		case *ast.ListValue:
			arr := make([]any, 0, len(v.Values))
			for _, it := range v.Values {
				arr = append(arr, parseASTValue(it))
			}
			return arr
		default:
			return parseASTValue(valueAST)
		}
	},
})

func parseASTValue(v ast.Value) any {
	switch x := v.(type) {
	case *ast.StringValue:
		return x.Value
	case *ast.BooleanValue:
		return x.Value
	case *ast.IntValue:
		// x.Value est une string
		i, err := strconv.ParseInt(x.Value, 10, 64)
		if err != nil {
			return x.Value
		}
		return i
	case *ast.FloatValue:
		f, err := strconv.ParseFloat(x.Value, 64)
		if err != nil {
			return x.Value
		}
		return f
	case *ast.EnumValue:
		// compat: parfois "null" arrive ici
		if strings.EqualFold(x.Value, "null") {
			return nil
		}
		return x.Value
	case *ast.ObjectValue:
		m := map[string]any{}
		for _, f := range x.Fields {
			m[f.Name.Value] = parseASTValue(f.Value)
		}
		return m
	case *ast.ListValue:
		arr := make([]any, 0, len(x.Values))
		for _, it := range x.Values {
			arr = append(arr, parseASTValue(it))
		}
		return arr
	default:
		return nil
	}
}
