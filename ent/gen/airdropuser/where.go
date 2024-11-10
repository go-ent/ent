// Code generated by ent, DO NOT EDIT.

package airdropuser

import (
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/ent/gen/predicate"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldID, id))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldAddress, v))
}

// PrivateKey applies equality check predicate on the "private_key" field. It's identical to PrivateKeyEQ.
func PrivateKey(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldPrivateKey, v))
}

// CanClaimAirdrop applies equality check predicate on the "can_claim_airdrop" field. It's identical to CanClaimAirdropEQ.
func CanClaimAirdrop(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldCanClaimAirdrop, v))
}

// LastAirdropClaimTime applies equality check predicate on the "last_airdrop_claim_time" field. It's identical to LastAirdropClaimTimeEQ.
func LastAirdropClaimTime(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldLastAirdropClaimTime, v))
}

// NextAirdropClaimTime applies equality check predicate on the "next_airdrop_claim_time" field. It's identical to NextAirdropClaimTimeEQ.
func NextAirdropClaimTime(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldNextAirdropClaimTime, v))
}

// ClaimedAmount applies equality check predicate on the "claimed_amount" field. It's identical to ClaimedAmountEQ.
func ClaimedAmount(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldClaimedAmount, v))
}

// ScheduledAmount applies equality check predicate on the "scheduled_amount" field. It's identical to ScheduledAmountEQ.
func ScheduledAmount(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldScheduledAmount, v))
}

// AirdropFailedAttempts applies equality check predicate on the "airdrop_failed_attempts" field. It's identical to AirdropFailedAttemptsEQ.
func AirdropFailedAttempts(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldAirdropFailedAttempts, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldContainsFold(FieldAddress, v))
}

// PrivateKeyEQ applies the EQ predicate on the "private_key" field.
func PrivateKeyEQ(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldPrivateKey, v))
}

// PrivateKeyNEQ applies the NEQ predicate on the "private_key" field.
func PrivateKeyNEQ(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldPrivateKey, v))
}

// PrivateKeyIn applies the In predicate on the "private_key" field.
func PrivateKeyIn(vs ...string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldPrivateKey, vs...))
}

// PrivateKeyNotIn applies the NotIn predicate on the "private_key" field.
func PrivateKeyNotIn(vs ...string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldPrivateKey, vs...))
}

// PrivateKeyGT applies the GT predicate on the "private_key" field.
func PrivateKeyGT(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldPrivateKey, v))
}

// PrivateKeyGTE applies the GTE predicate on the "private_key" field.
func PrivateKeyGTE(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldPrivateKey, v))
}

// PrivateKeyLT applies the LT predicate on the "private_key" field.
func PrivateKeyLT(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldPrivateKey, v))
}

// PrivateKeyLTE applies the LTE predicate on the "private_key" field.
func PrivateKeyLTE(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldPrivateKey, v))
}

// PrivateKeyContains applies the Contains predicate on the "private_key" field.
func PrivateKeyContains(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldContains(FieldPrivateKey, v))
}

// PrivateKeyHasPrefix applies the HasPrefix predicate on the "private_key" field.
func PrivateKeyHasPrefix(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldHasPrefix(FieldPrivateKey, v))
}

// PrivateKeyHasSuffix applies the HasSuffix predicate on the "private_key" field.
func PrivateKeyHasSuffix(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldHasSuffix(FieldPrivateKey, v))
}

// PrivateKeyEqualFold applies the EqualFold predicate on the "private_key" field.
func PrivateKeyEqualFold(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEqualFold(FieldPrivateKey, v))
}

// PrivateKeyContainsFold applies the ContainsFold predicate on the "private_key" field.
func PrivateKeyContainsFold(v string) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldContainsFold(FieldPrivateKey, v))
}

// CanClaimAirdropEQ applies the EQ predicate on the "can_claim_airdrop" field.
func CanClaimAirdropEQ(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldCanClaimAirdrop, v))
}

// CanClaimAirdropNEQ applies the NEQ predicate on the "can_claim_airdrop" field.
func CanClaimAirdropNEQ(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldCanClaimAirdrop, v))
}

// CanClaimAirdropIn applies the In predicate on the "can_claim_airdrop" field.
func CanClaimAirdropIn(vs ...int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldCanClaimAirdrop, vs...))
}

// CanClaimAirdropNotIn applies the NotIn predicate on the "can_claim_airdrop" field.
func CanClaimAirdropNotIn(vs ...int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldCanClaimAirdrop, vs...))
}

// CanClaimAirdropGT applies the GT predicate on the "can_claim_airdrop" field.
func CanClaimAirdropGT(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldCanClaimAirdrop, v))
}

// CanClaimAirdropGTE applies the GTE predicate on the "can_claim_airdrop" field.
func CanClaimAirdropGTE(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldCanClaimAirdrop, v))
}

// CanClaimAirdropLT applies the LT predicate on the "can_claim_airdrop" field.
func CanClaimAirdropLT(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldCanClaimAirdrop, v))
}

// CanClaimAirdropLTE applies the LTE predicate on the "can_claim_airdrop" field.
func CanClaimAirdropLTE(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldCanClaimAirdrop, v))
}

// LastAirdropClaimTimeEQ applies the EQ predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeEQ(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldLastAirdropClaimTime, v))
}

// LastAirdropClaimTimeNEQ applies the NEQ predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeNEQ(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldLastAirdropClaimTime, v))
}

// LastAirdropClaimTimeIn applies the In predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeIn(vs ...int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldLastAirdropClaimTime, vs...))
}

// LastAirdropClaimTimeNotIn applies the NotIn predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeNotIn(vs ...int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldLastAirdropClaimTime, vs...))
}

// LastAirdropClaimTimeGT applies the GT predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeGT(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldLastAirdropClaimTime, v))
}

// LastAirdropClaimTimeGTE applies the GTE predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeGTE(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldLastAirdropClaimTime, v))
}

// LastAirdropClaimTimeLT applies the LT predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeLT(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldLastAirdropClaimTime, v))
}

// LastAirdropClaimTimeLTE applies the LTE predicate on the "last_airdrop_claim_time" field.
func LastAirdropClaimTimeLTE(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldLastAirdropClaimTime, v))
}

// NextAirdropClaimTimeEQ applies the EQ predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeEQ(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldNextAirdropClaimTime, v))
}

// NextAirdropClaimTimeNEQ applies the NEQ predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeNEQ(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldNextAirdropClaimTime, v))
}

// NextAirdropClaimTimeIn applies the In predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeIn(vs ...int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldNextAirdropClaimTime, vs...))
}

// NextAirdropClaimTimeNotIn applies the NotIn predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeNotIn(vs ...int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldNextAirdropClaimTime, vs...))
}

// NextAirdropClaimTimeGT applies the GT predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeGT(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldNextAirdropClaimTime, v))
}

// NextAirdropClaimTimeGTE applies the GTE predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeGTE(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldNextAirdropClaimTime, v))
}

// NextAirdropClaimTimeLT applies the LT predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeLT(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldNextAirdropClaimTime, v))
}

// NextAirdropClaimTimeLTE applies the LTE predicate on the "next_airdrop_claim_time" field.
func NextAirdropClaimTimeLTE(v int64) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldNextAirdropClaimTime, v))
}

// ClaimedAmountEQ applies the EQ predicate on the "claimed_amount" field.
func ClaimedAmountEQ(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldClaimedAmount, v))
}

// ClaimedAmountNEQ applies the NEQ predicate on the "claimed_amount" field.
func ClaimedAmountNEQ(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldClaimedAmount, v))
}

// ClaimedAmountIn applies the In predicate on the "claimed_amount" field.
func ClaimedAmountIn(vs ...decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldClaimedAmount, vs...))
}

// ClaimedAmountNotIn applies the NotIn predicate on the "claimed_amount" field.
func ClaimedAmountNotIn(vs ...decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldClaimedAmount, vs...))
}

// ClaimedAmountGT applies the GT predicate on the "claimed_amount" field.
func ClaimedAmountGT(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldClaimedAmount, v))
}

// ClaimedAmountGTE applies the GTE predicate on the "claimed_amount" field.
func ClaimedAmountGTE(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldClaimedAmount, v))
}

// ClaimedAmountLT applies the LT predicate on the "claimed_amount" field.
func ClaimedAmountLT(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldClaimedAmount, v))
}

// ClaimedAmountLTE applies the LTE predicate on the "claimed_amount" field.
func ClaimedAmountLTE(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldClaimedAmount, v))
}

// ScheduledAmountEQ applies the EQ predicate on the "scheduled_amount" field.
func ScheduledAmountEQ(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldScheduledAmount, v))
}

// ScheduledAmountNEQ applies the NEQ predicate on the "scheduled_amount" field.
func ScheduledAmountNEQ(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldScheduledAmount, v))
}

// ScheduledAmountIn applies the In predicate on the "scheduled_amount" field.
func ScheduledAmountIn(vs ...decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldScheduledAmount, vs...))
}

// ScheduledAmountNotIn applies the NotIn predicate on the "scheduled_amount" field.
func ScheduledAmountNotIn(vs ...decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldScheduledAmount, vs...))
}

// ScheduledAmountGT applies the GT predicate on the "scheduled_amount" field.
func ScheduledAmountGT(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldScheduledAmount, v))
}

// ScheduledAmountGTE applies the GTE predicate on the "scheduled_amount" field.
func ScheduledAmountGTE(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldScheduledAmount, v))
}

// ScheduledAmountLT applies the LT predicate on the "scheduled_amount" field.
func ScheduledAmountLT(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldScheduledAmount, v))
}

// ScheduledAmountLTE applies the LTE predicate on the "scheduled_amount" field.
func ScheduledAmountLTE(v decimal.Decimal) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldScheduledAmount, v))
}

// AirdropFailedAttemptsEQ applies the EQ predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsEQ(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldEQ(FieldAirdropFailedAttempts, v))
}

// AirdropFailedAttemptsNEQ applies the NEQ predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsNEQ(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNEQ(FieldAirdropFailedAttempts, v))
}

// AirdropFailedAttemptsIn applies the In predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsIn(vs ...int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldIn(FieldAirdropFailedAttempts, vs...))
}

// AirdropFailedAttemptsNotIn applies the NotIn predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsNotIn(vs ...int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldNotIn(FieldAirdropFailedAttempts, vs...))
}

// AirdropFailedAttemptsGT applies the GT predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsGT(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGT(FieldAirdropFailedAttempts, v))
}

// AirdropFailedAttemptsGTE applies the GTE predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsGTE(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldGTE(FieldAirdropFailedAttempts, v))
}

// AirdropFailedAttemptsLT applies the LT predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsLT(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLT(FieldAirdropFailedAttempts, v))
}

// AirdropFailedAttemptsLTE applies the LTE predicate on the "airdrop_failed_attempts" field.
func AirdropFailedAttemptsLTE(v int) predicate.AirdropUser {
	return predicate.AirdropUser(sql.FieldLTE(FieldAirdropFailedAttempts, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AirdropUser) predicate.AirdropUser {
	return predicate.AirdropUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AirdropUser) predicate.AirdropUser {
	return predicate.AirdropUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AirdropUser) predicate.AirdropUser {
	return predicate.AirdropUser(sql.NotPredicates(p))
}