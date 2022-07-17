package codegen

import "strings"

type (
	DbModel struct {
		StructName             string
		TableName              string
		PrivatePropertyName    string
		PrimaryKeyPropertyName string
		PathName               string
		Fields                 []*DbModelFieldAndColumn
	}

	DbModelFieldAndColumn struct {
		ModelFieldName                      string
		PrivateModelFieldName               string
		ColumnName                          string
		ModelFieldType                      string
		ColumnType                          string
		IsPrimaryKey                        bool
		IsAutoIncrement                     bool
		IgnoreGenerateRequestModel          bool
		IgnoreGenerateRequestQueryParameter bool
		IgnoreGenerateResponseModel         bool
	}
)

func (t *DbModel) HasDeletedColumn() bool {
	for _, field := range t.Fields {
		if strings.EqualFold(field.ModelFieldName, "deleted") {
			return true
		}
	}
	return false
}
