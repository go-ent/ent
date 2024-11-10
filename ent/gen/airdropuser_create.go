// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-ent/ent/dialect/sql/sqlgraph"
	"github.com/go-ent/ent/ent/gen/airdropuser"
	"github.com/go-ent/ent/schema/field"
	"github.com/shopspring/decimal"
)

// AirdropUserCreate is the builder for creating a AirdropUser entity.
type AirdropUserCreate struct {
	config
	mutation *AirdropUserMutation
	hooks    []Hook
}

// SetAddress sets the "address" field.
func (auc *AirdropUserCreate) SetAddress(s string) *AirdropUserCreate {
	auc.mutation.SetAddress(s)
	return auc
}

// SetPrivateKey sets the "private_key" field.
func (auc *AirdropUserCreate) SetPrivateKey(s string) *AirdropUserCreate {
	auc.mutation.SetPrivateKey(s)
	return auc
}

// SetCanClaimAirdrop sets the "can_claim_airdrop" field.
func (auc *AirdropUserCreate) SetCanClaimAirdrop(i int) *AirdropUserCreate {
	auc.mutation.SetCanClaimAirdrop(i)
	return auc
}

// SetNillableCanClaimAirdrop sets the "can_claim_airdrop" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableCanClaimAirdrop(i *int) *AirdropUserCreate {
	if i != nil {
		auc.SetCanClaimAirdrop(*i)
	}
	return auc
}

// SetLastAirdropClaimTime sets the "last_airdrop_claim_time" field.
func (auc *AirdropUserCreate) SetLastAirdropClaimTime(i int64) *AirdropUserCreate {
	auc.mutation.SetLastAirdropClaimTime(i)
	return auc
}

// SetNillableLastAirdropClaimTime sets the "last_airdrop_claim_time" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableLastAirdropClaimTime(i *int64) *AirdropUserCreate {
	if i != nil {
		auc.SetLastAirdropClaimTime(*i)
	}
	return auc
}

// SetNextAirdropClaimTime sets the "next_airdrop_claim_time" field.
func (auc *AirdropUserCreate) SetNextAirdropClaimTime(i int64) *AirdropUserCreate {
	auc.mutation.SetNextAirdropClaimTime(i)
	return auc
}

// SetNillableNextAirdropClaimTime sets the "next_airdrop_claim_time" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableNextAirdropClaimTime(i *int64) *AirdropUserCreate {
	if i != nil {
		auc.SetNextAirdropClaimTime(*i)
	}
	return auc
}

// SetClaimedAmount sets the "claimed_amount" field.
func (auc *AirdropUserCreate) SetClaimedAmount(d decimal.Decimal) *AirdropUserCreate {
	auc.mutation.SetClaimedAmount(d)
	return auc
}

// SetNillableClaimedAmount sets the "claimed_amount" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableClaimedAmount(d *decimal.Decimal) *AirdropUserCreate {
	if d != nil {
		auc.SetClaimedAmount(*d)
	}
	return auc
}

// SetScheduledAmount sets the "scheduled_amount" field.
func (auc *AirdropUserCreate) SetScheduledAmount(d decimal.Decimal) *AirdropUserCreate {
	auc.mutation.SetScheduledAmount(d)
	return auc
}

// SetNillableScheduledAmount sets the "scheduled_amount" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableScheduledAmount(d *decimal.Decimal) *AirdropUserCreate {
	if d != nil {
		auc.SetScheduledAmount(*d)
	}
	return auc
}

// SetAirdropFailedAttempts sets the "airdrop_failed_attempts" field.
func (auc *AirdropUserCreate) SetAirdropFailedAttempts(i int) *AirdropUserCreate {
	auc.mutation.SetAirdropFailedAttempts(i)
	return auc
}

// SetNillableAirdropFailedAttempts sets the "airdrop_failed_attempts" field if the given value is not nil.
func (auc *AirdropUserCreate) SetNillableAirdropFailedAttempts(i *int) *AirdropUserCreate {
	if i != nil {
		auc.SetAirdropFailedAttempts(*i)
	}
	return auc
}

// SetID sets the "id" field.
func (auc *AirdropUserCreate) SetID(u uint) *AirdropUserCreate {
	auc.mutation.SetID(u)
	return auc
}

// Mutation returns the AirdropUserMutation object of the builder.
func (auc *AirdropUserCreate) Mutation() *AirdropUserMutation {
	return auc.mutation
}

// Save creates the AirdropUser in the database.
func (auc *AirdropUserCreate) Save(ctx context.Context) (*AirdropUser, error) {
	auc.defaults()
	return withHooks(ctx, auc.sqlSave, auc.mutation, auc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (auc *AirdropUserCreate) SaveX(ctx context.Context) *AirdropUser {
	v, err := auc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auc *AirdropUserCreate) Exec(ctx context.Context) error {
	_, err := auc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auc *AirdropUserCreate) ExecX(ctx context.Context) {
	if err := auc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auc *AirdropUserCreate) defaults() {
	if _, ok := auc.mutation.CanClaimAirdrop(); !ok {
		v := airdropuser.DefaultCanClaimAirdrop
		auc.mutation.SetCanClaimAirdrop(v)
	}
	if _, ok := auc.mutation.LastAirdropClaimTime(); !ok {
		v := airdropuser.DefaultLastAirdropClaimTime
		auc.mutation.SetLastAirdropClaimTime(v)
	}
	if _, ok := auc.mutation.NextAirdropClaimTime(); !ok {
		v := airdropuser.DefaultNextAirdropClaimTime
		auc.mutation.SetNextAirdropClaimTime(v)
	}
	if _, ok := auc.mutation.ClaimedAmount(); !ok {
		v := airdropuser.DefaultClaimedAmount
		auc.mutation.SetClaimedAmount(v)
	}
	if _, ok := auc.mutation.ScheduledAmount(); !ok {
		v := airdropuser.DefaultScheduledAmount
		auc.mutation.SetScheduledAmount(v)
	}
	if _, ok := auc.mutation.AirdropFailedAttempts(); !ok {
		v := airdropuser.DefaultAirdropFailedAttempts
		auc.mutation.SetAirdropFailedAttempts(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auc *AirdropUserCreate) check() error {
	if _, ok := auc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`gen: missing required field "AirdropUser.address"`)}
	}
	if v, ok := auc.mutation.Address(); ok {
		if err := airdropuser.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`gen: validator failed for field "AirdropUser.address": %w`, err)}
		}
	}
	if _, ok := auc.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "private_key", err: errors.New(`gen: missing required field "AirdropUser.private_key"`)}
	}
	if v, ok := auc.mutation.PrivateKey(); ok {
		if err := airdropuser.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf(`gen: validator failed for field "AirdropUser.private_key": %w`, err)}
		}
	}
	if _, ok := auc.mutation.CanClaimAirdrop(); !ok {
		return &ValidationError{Name: "can_claim_airdrop", err: errors.New(`gen: missing required field "AirdropUser.can_claim_airdrop"`)}
	}
	if _, ok := auc.mutation.LastAirdropClaimTime(); !ok {
		return &ValidationError{Name: "last_airdrop_claim_time", err: errors.New(`gen: missing required field "AirdropUser.last_airdrop_claim_time"`)}
	}
	if _, ok := auc.mutation.NextAirdropClaimTime(); !ok {
		return &ValidationError{Name: "next_airdrop_claim_time", err: errors.New(`gen: missing required field "AirdropUser.next_airdrop_claim_time"`)}
	}
	if _, ok := auc.mutation.ClaimedAmount(); !ok {
		return &ValidationError{Name: "claimed_amount", err: errors.New(`gen: missing required field "AirdropUser.claimed_amount"`)}
	}
	if _, ok := auc.mutation.ScheduledAmount(); !ok {
		return &ValidationError{Name: "scheduled_amount", err: errors.New(`gen: missing required field "AirdropUser.scheduled_amount"`)}
	}
	if _, ok := auc.mutation.AirdropFailedAttempts(); !ok {
		return &ValidationError{Name: "airdrop_failed_attempts", err: errors.New(`gen: missing required field "AirdropUser.airdrop_failed_attempts"`)}
	}
	if v, ok := auc.mutation.AirdropFailedAttempts(); ok {
		if err := airdropuser.AirdropFailedAttemptsValidator(v); err != nil {
			return &ValidationError{Name: "airdrop_failed_attempts", err: fmt.Errorf(`gen: validator failed for field "AirdropUser.airdrop_failed_attempts": %w`, err)}
		}
	}
	return nil
}

func (auc *AirdropUserCreate) sqlSave(ctx context.Context) (*AirdropUser, error) {
	if err := auc.check(); err != nil {
		return nil, err
	}
	_node, _spec := auc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint(id)
	}
	auc.mutation.id = &_node.ID
	auc.mutation.done = true
	return _node, nil
}

func (auc *AirdropUserCreate) createSpec() (*AirdropUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AirdropUser{config: auc.config}
		_spec = sqlgraph.NewCreateSpec(airdropuser.Table, sqlgraph.NewFieldSpec(airdropuser.FieldID, field.TypeUint))
	)
	if id, ok := auc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := auc.mutation.Address(); ok {
		_spec.SetField(airdropuser.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := auc.mutation.PrivateKey(); ok {
		_spec.SetField(airdropuser.FieldPrivateKey, field.TypeString, value)
		_node.PrivateKey = value
	}
	if value, ok := auc.mutation.CanClaimAirdrop(); ok {
		_spec.SetField(airdropuser.FieldCanClaimAirdrop, field.TypeInt, value)
		_node.CanClaimAirdrop = value
	}
	if value, ok := auc.mutation.LastAirdropClaimTime(); ok {
		_spec.SetField(airdropuser.FieldLastAirdropClaimTime, field.TypeInt64, value)
		_node.LastAirdropClaimTime = value
	}
	if value, ok := auc.mutation.NextAirdropClaimTime(); ok {
		_spec.SetField(airdropuser.FieldNextAirdropClaimTime, field.TypeInt64, value)
		_node.NextAirdropClaimTime = value
	}
	if value, ok := auc.mutation.ClaimedAmount(); ok {
		_spec.SetField(airdropuser.FieldClaimedAmount, field.TypeOther, value)
		_node.ClaimedAmount = value
	}
	if value, ok := auc.mutation.ScheduledAmount(); ok {
		_spec.SetField(airdropuser.FieldScheduledAmount, field.TypeOther, value)
		_node.ScheduledAmount = value
	}
	if value, ok := auc.mutation.AirdropFailedAttempts(); ok {
		_spec.SetField(airdropuser.FieldAirdropFailedAttempts, field.TypeInt, value)
		_node.AirdropFailedAttempts = value
	}
	return _node, _spec
}

// AirdropUserCreateBulk is the builder for creating many AirdropUser entities in bulk.
type AirdropUserCreateBulk struct {
	config
	err      error
	builders []*AirdropUserCreate
}

// Save creates the AirdropUser entities in the database.
func (aucb *AirdropUserCreateBulk) Save(ctx context.Context) ([]*AirdropUser, error) {
	if aucb.err != nil {
		return nil, aucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aucb.builders))
	nodes := make([]*AirdropUser, len(aucb.builders))
	mutators := make([]Mutator, len(aucb.builders))
	for i := range aucb.builders {
		func(i int, root context.Context) {
			builder := aucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AirdropUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aucb *AirdropUserCreateBulk) SaveX(ctx context.Context) []*AirdropUser {
	v, err := aucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aucb *AirdropUserCreateBulk) Exec(ctx context.Context) error {
	_, err := aucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucb *AirdropUserCreateBulk) ExecX(ctx context.Context) {
	if err := aucb.Exec(ctx); err != nil {
		panic(err)
	}
}
