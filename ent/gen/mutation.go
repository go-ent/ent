// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-ent/ent"
	"github.com/go-ent/ent/dialect/sql"
	"github.com/go-ent/ent/ent/gen/airdropuser"
	"github.com/go-ent/ent/ent/gen/predicate"
	"github.com/shopspring/decimal"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAirdropUser = "AirdropUser"
)

// AirdropUserMutation represents an operation that mutates the AirdropUser nodes in the graph.
type AirdropUserMutation struct {
	config
	op                         Op
	typ                        string
	id                         *uint
	address                    *string
	private_key                *string
	can_claim_airdrop          *int
	addcan_claim_airdrop       *int
	last_airdrop_claim_time    *int64
	addlast_airdrop_claim_time *int64
	next_airdrop_claim_time    *int64
	addnext_airdrop_claim_time *int64
	claimed_amount             *decimal.Decimal
	scheduled_amount           *decimal.Decimal
	airdrop_failed_attempts    *int
	addairdrop_failed_attempts *int
	clearedFields              map[string]struct{}
	done                       bool
	oldValue                   func(context.Context) (*AirdropUser, error)
	predicates                 []predicate.AirdropUser
}

var _ ent.Mutation = (*AirdropUserMutation)(nil)

// airdropuserOption allows management of the mutation configuration using functional options.
type airdropuserOption func(*AirdropUserMutation)

// newAirdropUserMutation creates new mutation for the AirdropUser entity.
func newAirdropUserMutation(c config, op Op, opts ...airdropuserOption) *AirdropUserMutation {
	m := &AirdropUserMutation{
		config:        c,
		op:            op,
		typ:           TypeAirdropUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAirdropUserID sets the ID field of the mutation.
func withAirdropUserID(id uint) airdropuserOption {
	return func(m *AirdropUserMutation) {
		var (
			err   error
			once  sync.Once
			value *AirdropUser
		)
		m.oldValue = func(ctx context.Context) (*AirdropUser, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AirdropUser.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAirdropUser sets the old AirdropUser of the mutation.
func withAirdropUser(node *AirdropUser) airdropuserOption {
	return func(m *AirdropUserMutation) {
		m.oldValue = func(context.Context) (*AirdropUser, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AirdropUserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AirdropUserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("gen: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of AirdropUser entities.
func (m *AirdropUserMutation) SetID(id uint) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AirdropUserMutation) ID() (id uint, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AirdropUserMutation) IDs(ctx context.Context) ([]uint, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().AirdropUser.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAddress sets the "address" field.
func (m *AirdropUserMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *AirdropUserMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ResetAddress resets all changes to the "address" field.
func (m *AirdropUserMutation) ResetAddress() {
	m.address = nil
}

// SetPrivateKey sets the "private_key" field.
func (m *AirdropUserMutation) SetPrivateKey(s string) {
	m.private_key = &s
}

// PrivateKey returns the value of the "private_key" field in the mutation.
func (m *AirdropUserMutation) PrivateKey() (r string, exists bool) {
	v := m.private_key
	if v == nil {
		return
	}
	return *v, true
}

// OldPrivateKey returns the old "private_key" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldPrivateKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrivateKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrivateKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrivateKey: %w", err)
	}
	return oldValue.PrivateKey, nil
}

// ResetPrivateKey resets all changes to the "private_key" field.
func (m *AirdropUserMutation) ResetPrivateKey() {
	m.private_key = nil
}

// SetCanClaimAirdrop sets the "can_claim_airdrop" field.
func (m *AirdropUserMutation) SetCanClaimAirdrop(i int) {
	m.can_claim_airdrop = &i
	m.addcan_claim_airdrop = nil
}

// CanClaimAirdrop returns the value of the "can_claim_airdrop" field in the mutation.
func (m *AirdropUserMutation) CanClaimAirdrop() (r int, exists bool) {
	v := m.can_claim_airdrop
	if v == nil {
		return
	}
	return *v, true
}

// OldCanClaimAirdrop returns the old "can_claim_airdrop" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldCanClaimAirdrop(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCanClaimAirdrop is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCanClaimAirdrop requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCanClaimAirdrop: %w", err)
	}
	return oldValue.CanClaimAirdrop, nil
}

// AddCanClaimAirdrop adds i to the "can_claim_airdrop" field.
func (m *AirdropUserMutation) AddCanClaimAirdrop(i int) {
	if m.addcan_claim_airdrop != nil {
		*m.addcan_claim_airdrop += i
	} else {
		m.addcan_claim_airdrop = &i
	}
}

// AddedCanClaimAirdrop returns the value that was added to the "can_claim_airdrop" field in this mutation.
func (m *AirdropUserMutation) AddedCanClaimAirdrop() (r int, exists bool) {
	v := m.addcan_claim_airdrop
	if v == nil {
		return
	}
	return *v, true
}

// ResetCanClaimAirdrop resets all changes to the "can_claim_airdrop" field.
func (m *AirdropUserMutation) ResetCanClaimAirdrop() {
	m.can_claim_airdrop = nil
	m.addcan_claim_airdrop = nil
}

// SetLastAirdropClaimTime sets the "last_airdrop_claim_time" field.
func (m *AirdropUserMutation) SetLastAirdropClaimTime(i int64) {
	m.last_airdrop_claim_time = &i
	m.addlast_airdrop_claim_time = nil
}

// LastAirdropClaimTime returns the value of the "last_airdrop_claim_time" field in the mutation.
func (m *AirdropUserMutation) LastAirdropClaimTime() (r int64, exists bool) {
	v := m.last_airdrop_claim_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastAirdropClaimTime returns the old "last_airdrop_claim_time" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldLastAirdropClaimTime(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastAirdropClaimTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastAirdropClaimTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastAirdropClaimTime: %w", err)
	}
	return oldValue.LastAirdropClaimTime, nil
}

// AddLastAirdropClaimTime adds i to the "last_airdrop_claim_time" field.
func (m *AirdropUserMutation) AddLastAirdropClaimTime(i int64) {
	if m.addlast_airdrop_claim_time != nil {
		*m.addlast_airdrop_claim_time += i
	} else {
		m.addlast_airdrop_claim_time = &i
	}
}

// AddedLastAirdropClaimTime returns the value that was added to the "last_airdrop_claim_time" field in this mutation.
func (m *AirdropUserMutation) AddedLastAirdropClaimTime() (r int64, exists bool) {
	v := m.addlast_airdrop_claim_time
	if v == nil {
		return
	}
	return *v, true
}

// ResetLastAirdropClaimTime resets all changes to the "last_airdrop_claim_time" field.
func (m *AirdropUserMutation) ResetLastAirdropClaimTime() {
	m.last_airdrop_claim_time = nil
	m.addlast_airdrop_claim_time = nil
}

// SetNextAirdropClaimTime sets the "next_airdrop_claim_time" field.
func (m *AirdropUserMutation) SetNextAirdropClaimTime(i int64) {
	m.next_airdrop_claim_time = &i
	m.addnext_airdrop_claim_time = nil
}

// NextAirdropClaimTime returns the value of the "next_airdrop_claim_time" field in the mutation.
func (m *AirdropUserMutation) NextAirdropClaimTime() (r int64, exists bool) {
	v := m.next_airdrop_claim_time
	if v == nil {
		return
	}
	return *v, true
}

// OldNextAirdropClaimTime returns the old "next_airdrop_claim_time" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldNextAirdropClaimTime(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNextAirdropClaimTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNextAirdropClaimTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNextAirdropClaimTime: %w", err)
	}
	return oldValue.NextAirdropClaimTime, nil
}

// AddNextAirdropClaimTime adds i to the "next_airdrop_claim_time" field.
func (m *AirdropUserMutation) AddNextAirdropClaimTime(i int64) {
	if m.addnext_airdrop_claim_time != nil {
		*m.addnext_airdrop_claim_time += i
	} else {
		m.addnext_airdrop_claim_time = &i
	}
}

// AddedNextAirdropClaimTime returns the value that was added to the "next_airdrop_claim_time" field in this mutation.
func (m *AirdropUserMutation) AddedNextAirdropClaimTime() (r int64, exists bool) {
	v := m.addnext_airdrop_claim_time
	if v == nil {
		return
	}
	return *v, true
}

// ResetNextAirdropClaimTime resets all changes to the "next_airdrop_claim_time" field.
func (m *AirdropUserMutation) ResetNextAirdropClaimTime() {
	m.next_airdrop_claim_time = nil
	m.addnext_airdrop_claim_time = nil
}

// SetClaimedAmount sets the "claimed_amount" field.
func (m *AirdropUserMutation) SetClaimedAmount(d decimal.Decimal) {
	m.claimed_amount = &d
}

// ClaimedAmount returns the value of the "claimed_amount" field in the mutation.
func (m *AirdropUserMutation) ClaimedAmount() (r decimal.Decimal, exists bool) {
	v := m.claimed_amount
	if v == nil {
		return
	}
	return *v, true
}

// OldClaimedAmount returns the old "claimed_amount" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldClaimedAmount(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldClaimedAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldClaimedAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldClaimedAmount: %w", err)
	}
	return oldValue.ClaimedAmount, nil
}

// ResetClaimedAmount resets all changes to the "claimed_amount" field.
func (m *AirdropUserMutation) ResetClaimedAmount() {
	m.claimed_amount = nil
}

// SetScheduledAmount sets the "scheduled_amount" field.
func (m *AirdropUserMutation) SetScheduledAmount(d decimal.Decimal) {
	m.scheduled_amount = &d
}

// ScheduledAmount returns the value of the "scheduled_amount" field in the mutation.
func (m *AirdropUserMutation) ScheduledAmount() (r decimal.Decimal, exists bool) {
	v := m.scheduled_amount
	if v == nil {
		return
	}
	return *v, true
}

// OldScheduledAmount returns the old "scheduled_amount" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldScheduledAmount(ctx context.Context) (v decimal.Decimal, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldScheduledAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldScheduledAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldScheduledAmount: %w", err)
	}
	return oldValue.ScheduledAmount, nil
}

// ResetScheduledAmount resets all changes to the "scheduled_amount" field.
func (m *AirdropUserMutation) ResetScheduledAmount() {
	m.scheduled_amount = nil
}

// SetAirdropFailedAttempts sets the "airdrop_failed_attempts" field.
func (m *AirdropUserMutation) SetAirdropFailedAttempts(i int) {
	m.airdrop_failed_attempts = &i
	m.addairdrop_failed_attempts = nil
}

// AirdropFailedAttempts returns the value of the "airdrop_failed_attempts" field in the mutation.
func (m *AirdropUserMutation) AirdropFailedAttempts() (r int, exists bool) {
	v := m.airdrop_failed_attempts
	if v == nil {
		return
	}
	return *v, true
}

// OldAirdropFailedAttempts returns the old "airdrop_failed_attempts" field's value of the AirdropUser entity.
// If the AirdropUser object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AirdropUserMutation) OldAirdropFailedAttempts(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAirdropFailedAttempts is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAirdropFailedAttempts requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAirdropFailedAttempts: %w", err)
	}
	return oldValue.AirdropFailedAttempts, nil
}

// AddAirdropFailedAttempts adds i to the "airdrop_failed_attempts" field.
func (m *AirdropUserMutation) AddAirdropFailedAttempts(i int) {
	if m.addairdrop_failed_attempts != nil {
		*m.addairdrop_failed_attempts += i
	} else {
		m.addairdrop_failed_attempts = &i
	}
}

// AddedAirdropFailedAttempts returns the value that was added to the "airdrop_failed_attempts" field in this mutation.
func (m *AirdropUserMutation) AddedAirdropFailedAttempts() (r int, exists bool) {
	v := m.addairdrop_failed_attempts
	if v == nil {
		return
	}
	return *v, true
}

// ResetAirdropFailedAttempts resets all changes to the "airdrop_failed_attempts" field.
func (m *AirdropUserMutation) ResetAirdropFailedAttempts() {
	m.airdrop_failed_attempts = nil
	m.addairdrop_failed_attempts = nil
}

// Where appends a list predicates to the AirdropUserMutation builder.
func (m *AirdropUserMutation) Where(ps ...predicate.AirdropUser) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AirdropUserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AirdropUserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.AirdropUser, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AirdropUserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AirdropUserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (AirdropUser).
func (m *AirdropUserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AirdropUserMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.address != nil {
		fields = append(fields, airdropuser.FieldAddress)
	}
	if m.private_key != nil {
		fields = append(fields, airdropuser.FieldPrivateKey)
	}
	if m.can_claim_airdrop != nil {
		fields = append(fields, airdropuser.FieldCanClaimAirdrop)
	}
	if m.last_airdrop_claim_time != nil {
		fields = append(fields, airdropuser.FieldLastAirdropClaimTime)
	}
	if m.next_airdrop_claim_time != nil {
		fields = append(fields, airdropuser.FieldNextAirdropClaimTime)
	}
	if m.claimed_amount != nil {
		fields = append(fields, airdropuser.FieldClaimedAmount)
	}
	if m.scheduled_amount != nil {
		fields = append(fields, airdropuser.FieldScheduledAmount)
	}
	if m.airdrop_failed_attempts != nil {
		fields = append(fields, airdropuser.FieldAirdropFailedAttempts)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AirdropUserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case airdropuser.FieldAddress:
		return m.Address()
	case airdropuser.FieldPrivateKey:
		return m.PrivateKey()
	case airdropuser.FieldCanClaimAirdrop:
		return m.CanClaimAirdrop()
	case airdropuser.FieldLastAirdropClaimTime:
		return m.LastAirdropClaimTime()
	case airdropuser.FieldNextAirdropClaimTime:
		return m.NextAirdropClaimTime()
	case airdropuser.FieldClaimedAmount:
		return m.ClaimedAmount()
	case airdropuser.FieldScheduledAmount:
		return m.ScheduledAmount()
	case airdropuser.FieldAirdropFailedAttempts:
		return m.AirdropFailedAttempts()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AirdropUserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case airdropuser.FieldAddress:
		return m.OldAddress(ctx)
	case airdropuser.FieldPrivateKey:
		return m.OldPrivateKey(ctx)
	case airdropuser.FieldCanClaimAirdrop:
		return m.OldCanClaimAirdrop(ctx)
	case airdropuser.FieldLastAirdropClaimTime:
		return m.OldLastAirdropClaimTime(ctx)
	case airdropuser.FieldNextAirdropClaimTime:
		return m.OldNextAirdropClaimTime(ctx)
	case airdropuser.FieldClaimedAmount:
		return m.OldClaimedAmount(ctx)
	case airdropuser.FieldScheduledAmount:
		return m.OldScheduledAmount(ctx)
	case airdropuser.FieldAirdropFailedAttempts:
		return m.OldAirdropFailedAttempts(ctx)
	}
	return nil, fmt.Errorf("unknown AirdropUser field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AirdropUserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case airdropuser.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case airdropuser.FieldPrivateKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrivateKey(v)
		return nil
	case airdropuser.FieldCanClaimAirdrop:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCanClaimAirdrop(v)
		return nil
	case airdropuser.FieldLastAirdropClaimTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastAirdropClaimTime(v)
		return nil
	case airdropuser.FieldNextAirdropClaimTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNextAirdropClaimTime(v)
		return nil
	case airdropuser.FieldClaimedAmount:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetClaimedAmount(v)
		return nil
	case airdropuser.FieldScheduledAmount:
		v, ok := value.(decimal.Decimal)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetScheduledAmount(v)
		return nil
	case airdropuser.FieldAirdropFailedAttempts:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAirdropFailedAttempts(v)
		return nil
	}
	return fmt.Errorf("unknown AirdropUser field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AirdropUserMutation) AddedFields() []string {
	var fields []string
	if m.addcan_claim_airdrop != nil {
		fields = append(fields, airdropuser.FieldCanClaimAirdrop)
	}
	if m.addlast_airdrop_claim_time != nil {
		fields = append(fields, airdropuser.FieldLastAirdropClaimTime)
	}
	if m.addnext_airdrop_claim_time != nil {
		fields = append(fields, airdropuser.FieldNextAirdropClaimTime)
	}
	if m.addairdrop_failed_attempts != nil {
		fields = append(fields, airdropuser.FieldAirdropFailedAttempts)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AirdropUserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case airdropuser.FieldCanClaimAirdrop:
		return m.AddedCanClaimAirdrop()
	case airdropuser.FieldLastAirdropClaimTime:
		return m.AddedLastAirdropClaimTime()
	case airdropuser.FieldNextAirdropClaimTime:
		return m.AddedNextAirdropClaimTime()
	case airdropuser.FieldAirdropFailedAttempts:
		return m.AddedAirdropFailedAttempts()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AirdropUserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case airdropuser.FieldCanClaimAirdrop:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCanClaimAirdrop(v)
		return nil
	case airdropuser.FieldLastAirdropClaimTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLastAirdropClaimTime(v)
		return nil
	case airdropuser.FieldNextAirdropClaimTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNextAirdropClaimTime(v)
		return nil
	case airdropuser.FieldAirdropFailedAttempts:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAirdropFailedAttempts(v)
		return nil
	}
	return fmt.Errorf("unknown AirdropUser numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AirdropUserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AirdropUserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AirdropUserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown AirdropUser nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AirdropUserMutation) ResetField(name string) error {
	switch name {
	case airdropuser.FieldAddress:
		m.ResetAddress()
		return nil
	case airdropuser.FieldPrivateKey:
		m.ResetPrivateKey()
		return nil
	case airdropuser.FieldCanClaimAirdrop:
		m.ResetCanClaimAirdrop()
		return nil
	case airdropuser.FieldLastAirdropClaimTime:
		m.ResetLastAirdropClaimTime()
		return nil
	case airdropuser.FieldNextAirdropClaimTime:
		m.ResetNextAirdropClaimTime()
		return nil
	case airdropuser.FieldClaimedAmount:
		m.ResetClaimedAmount()
		return nil
	case airdropuser.FieldScheduledAmount:
		m.ResetScheduledAmount()
		return nil
	case airdropuser.FieldAirdropFailedAttempts:
		m.ResetAirdropFailedAttempts()
		return nil
	}
	return fmt.Errorf("unknown AirdropUser field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AirdropUserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AirdropUserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AirdropUserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AirdropUserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AirdropUserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AirdropUserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AirdropUserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AirdropUser unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AirdropUserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AirdropUser edge %s", name)
}
