package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TenantMixin for tenant
type TenantMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").Default(0).Comment("Tenant ID | 租户ID").Annotations(entsql.WithComments(true)),
	}
}
