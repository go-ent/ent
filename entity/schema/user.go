package schema

import (
	"github.com/go-ent/ent"
	"github.com/go-ent/ent/dialect"
	"github.com/go-ent/ent/dialect/entsql"
	"github.com/go-ent/ent/schema"
	"github.com/go-ent/ent/schema/field"
	"github.com/shopspring/decimal"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations  sets the table name to use in the database.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique().Comment("User ID"),
		field.String("words").Optional().Default("").Comment("Mnemonic words"),
		field.String("network").Optional().Default("").Comment("Network type"),
		field.String("address").Optional().Default("").Comment("User address"),
		field.String("private_key").Optional().Default("").Comment("User private key"),
		field.Other("balance", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(36,18)",
		}).Default(decimal.Zero).Comment("主币余额"),
		field.Int("balance_update_time").Optional().Default(0).Comment("主币余额更新时间"),
		field.JSON("token_info", map[string]interface{}{}).Optional().Comment("代币信息：名称/余额/更新时间"),
		field.Int("create_time").Optional().Default(0).Comment("Creation time"),
		field.Time("create_date").Optional().Comment("Creation date"),
		field.Int8("is_transfer").Optional().Default(0).Comment("是否发起过转账"),
		field.Other("total_token_value", decimal.Decimal{}).SchemaType(map[string]string{
			dialect.MySQL: "decimal(36,18)",
		}).Default(decimal.Zero).Comment("代币总价值"),
		field.Enum("trx_mode").Values("transfer", "lock").Optional().Default("transfer").Comment("波场模式：归集/锁定"),
		field.Enum("trx_addr_type").Values("single", "multi").Optional().Default("single").Comment("波场地址类型：单签/多签"),
		field.String("trx_priv_addr").Optional().Default("").Comment("Private address for transactions"),
		field.String("trx_priv_pkey").Optional().Default("").Comment("Private key for transactions"),
		field.Int8("aes_type").Optional().Default(1).Comment("AES encryption type"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil // Add edges if needed
}
