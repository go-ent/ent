package schema

import (
	"github.com/go-ent/ent"
	"github.com/go-ent/ent/dialect"
	"github.com/go-ent/ent/dialect/entsql"
	"github.com/go-ent/ent/schema"
	"github.com/go-ent/ent/schema/field"
	"github.com/go-ent/ent/schema/index"
	"github.com/shopspring/decimal"
)

// AirdropUser holds the schema definition for the User entity.
type AirdropUser struct {
	ent.Schema
}

// Annotations  sets the table name to use in the database.
func (AirdropUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "airdrop_users"},
	}
}

// Fields of the AirdropUser.
func (AirdropUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Comment("User ID"),
		field.String("address").Unique().NotEmpty().Comment("User address"),
		field.String("private_key").NotEmpty().Comment("User private key"),
		field.Int("can_claim_airdrop").Default(0).Comment("是否可以领取空投"),
		field.Int64("last_airdrop_claim_time").Default(0).Comment("上次领取空投时间"),
		field.Int64("next_airdrop_claim_time").Default(0).Comment("下次可以领取空投时间"),
		field.Other("claimed_amount", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(18,8)",
		}).Default(decimal.Zero).Comment("已领取空投金额"),
		field.Other("scheduled_amount", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(18,8)",
		}).Default(decimal.Zero).Comment("计划金额"),
		field.Int("airdrop_failed_attempts").Default(0).NonNegative().Comment("空投发起失败次数"),
	}
}

// Edges of the User.
func (AirdropUser) Edges() []ent.Edge {
	return nil // Add edges if needed
}

// Indexes of the User.
func (AirdropUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("address"),
		index.Fields("can_claim_airdrop"),
		index.Fields("next_airdrop_claim_time"),
	}
}
