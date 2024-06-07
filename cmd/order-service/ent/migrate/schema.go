// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Default: 0},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "order_id",
				Unique:  true,
				Columns: []*schema.Column{OrdersColumns[0]},
			},
			{
				Name:    "order_user_id",
				Unique:  false,
				Columns: []*schema.Column{OrdersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
	}
)

func init() {
}
