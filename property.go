package swag

import (
	"go/ast"
	"log"
	"strings"
)

// getPropertyName returns the string value for the given field if it exists, otherwise it panics.
// allowedValues: array, boolean, integer, null, number, object, string
func getPropertyName(field *ast.Field) string {
	var name string
	if astTypeSelectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {

		// Support for time.Time as a structure field
		if "Time" == astTypeSelectorExpr.Sel.Name {
			return "string"
		}

		// Support bson.ObjectId type
		if "ObjectId" == astTypeSelectorExpr.Sel.Name {
			return "string"
		}

		panic("not supported 'astSelectorExpr' yet.")

	} else if astTypeIdent, ok := field.Type.(*ast.Ident); ok {
		name = astTypeIdent.Name

		// When its the int type will transfer to integer which is goswagger supported type
		if "int" == name {
			return "integer"
		}
		// Transfer float32 and float64 to number type
		if strings.HasPrefix(name, "float") {
			return "number"
		}

	} else if _, ok := field.Type.(*ast.StarExpr); ok {
		panic("not supported astStarExpr yet.")
	} else if _, ok := field.Type.(*ast.MapType); ok { // if map
		//TODO: support map
		return "object"
	} else if _, ok := field.Type.(*ast.ArrayType); ok { // if array
		return "array"
	} else if _, ok := field.Type.(*ast.StructType); ok { // if struct
		//TODO: support nested struct
		return "object"
	} else {
		log.Fatalf("Something goes wrong: %#v", field.Type)
	}

	return name
}
