// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"

	"github.com/go-ent/ent"
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/ent/gen/airdropuser"
	"github.com/shopspring/decimal"
)

// AirdropUser is the model entity for the AirdropUser schema.
type AirdropUser struct {
	config `json:"-"`
	// ID of the ent.
	// User ID
	ID uint `json:"id,omitempty"`
	// User address
	Address string `json:"address,omitempty"`
	// User private key
	PrivateKey string `json:"private_key,omitempty"`
	// 是否可以领取空投
	CanClaimAirdrop int `json:"can_claim_airdrop,omitempty"`
	// 上次领取空投时间
	LastAirdropClaimTime int64 `json:"last_airdrop_claim_time,omitempty"`
	// 下次可以领取空投时间
	NextAirdropClaimTime int64 `json:"next_airdrop_claim_time,omitempty"`
	// 已领取空投金额
	ClaimedAmount decimal.Decimal `json:"claimed_amount,omitempty"`
	// 计划金额
	ScheduledAmount decimal.Decimal `json:"scheduled_amount,omitempty"`
	// 空投发起失败次数
	AirdropFailedAttempts int `json:"airdrop_failed_attempts,omitempty"`
	selectValues          sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AirdropUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case airdropuser.FieldClaimedAmount, airdropuser.FieldScheduledAmount:
			values[i] = new(decimal.Decimal)
		case airdropuser.FieldID, airdropuser.FieldCanClaimAirdrop, airdropuser.FieldLastAirdropClaimTime, airdropuser.FieldNextAirdropClaimTime, airdropuser.FieldAirdropFailedAttempts:
			values[i] = new(sql.NullInt64)
		case airdropuser.FieldAddress, airdropuser.FieldPrivateKey:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AirdropUser fields.
func (au *AirdropUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case airdropuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			au.ID = uint(value.Int64)
		case airdropuser.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				au.Address = value.String
			}
		case airdropuser.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				au.PrivateKey = value.String
			}
		case airdropuser.FieldCanClaimAirdrop:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field can_claim_airdrop", values[i])
			} else if value.Valid {
				au.CanClaimAirdrop = int(value.Int64)
			}
		case airdropuser.FieldLastAirdropClaimTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_airdrop_claim_time", values[i])
			} else if value.Valid {
				au.LastAirdropClaimTime = value.Int64
			}
		case airdropuser.FieldNextAirdropClaimTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field next_airdrop_claim_time", values[i])
			} else if value.Valid {
				au.NextAirdropClaimTime = value.Int64
			}
		case airdropuser.FieldClaimedAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field claimed_amount", values[i])
			} else if value != nil {
				au.ClaimedAmount = *value
			}
		case airdropuser.FieldScheduledAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field scheduled_amount", values[i])
			} else if value != nil {
				au.ScheduledAmount = *value
			}
		case airdropuser.FieldAirdropFailedAttempts:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field airdrop_failed_attempts", values[i])
			} else if value.Valid {
				au.AirdropFailedAttempts = int(value.Int64)
			}
		default:
			au.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AirdropUser.
// This includes values selected through modifiers, order, etc.
func (au *AirdropUser) Value(name string) (ent.Value, error) {
	return au.selectValues.Get(name)
}

// Update returns a builder for updating this AirdropUser.
// Note that you need to call AirdropUser.Unwrap() before calling this method if this AirdropUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (au *AirdropUser) Update() *AirdropUserUpdateOne {
	return NewAirdropUserClient(au.config).UpdateOne(au)
}

// Unwrap unwraps the AirdropUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (au *AirdropUser) Unwrap() *AirdropUser {
	_tx, ok := au.config.driver.(*txDriver)
	if !ok {
		panic("gen: AirdropUser is not a transactional entity")
	}
	au.config.driver = _tx.drv
	return au
}

// String implements the fmt.Stringer.
func (au *AirdropUser) String() string {
	var builder strings.Builder
	builder.WriteString("AirdropUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", au.ID))
	builder.WriteString("address=")
	builder.WriteString(au.Address)
	builder.WriteString(", ")
	builder.WriteString("private_key=")
	builder.WriteString(au.PrivateKey)
	builder.WriteString(", ")
	builder.WriteString("can_claim_airdrop=")
	builder.WriteString(fmt.Sprintf("%v", au.CanClaimAirdrop))
	builder.WriteString(", ")
	builder.WriteString("last_airdrop_claim_time=")
	builder.WriteString(fmt.Sprintf("%v", au.LastAirdropClaimTime))
	builder.WriteString(", ")
	builder.WriteString("next_airdrop_claim_time=")
	builder.WriteString(fmt.Sprintf("%v", au.NextAirdropClaimTime))
	builder.WriteString(", ")
	builder.WriteString("claimed_amount=")
	builder.WriteString(fmt.Sprintf("%v", au.ClaimedAmount))
	builder.WriteString(", ")
	builder.WriteString("scheduled_amount=")
	builder.WriteString(fmt.Sprintf("%v", au.ScheduledAmount))
	builder.WriteString(", ")
	builder.WriteString("airdrop_failed_attempts=")
	builder.WriteString(fmt.Sprintf("%v", au.AirdropFailedAttempts))
	builder.WriteByte(')')
	return builder.String()
}

// AirdropUsers is a parsable slice of AirdropUser.
type AirdropUsers []*AirdropUser
