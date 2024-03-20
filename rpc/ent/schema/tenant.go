package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Tenant struct {
	ent.Schema
}

func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		// 租户id
		field.Int("tenant_id").Default(0).Comment("Tenant ID | 租户ID").Annotations(entsql.WithComments(true)),
		// 租户名称
		field.String("name").Unique().Comment("Tenant name | 租户名称").Annotations(entsql.WithComments(true)),
		//// app key
		//field.String("app_key").Unique().Comment("App key | 应用密钥").Annotations(entsql.WithComments(true)),
		//// app secret
		//field.String("app_secret").Comment("App secret | 应用密钥").Annotations(entsql.WithComments(true)),
	}
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id").
			Unique(),
	}
}

func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_tenants"},
	}
}
