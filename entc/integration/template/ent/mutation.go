// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/facebookincubator/ent/entc/integration/template/ent/group"
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGroup = "Group"
	TypePet   = "Pet"
	TypeUser  = "User"
)

// GroupMutation represents an operation that mutate the Groups
// nodes in the graph.
type GroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	max_users     *int
	addmax_users  *int
	clearedFields map[string]struct{}
	// lazy load old fields.
	oldOnce  sync.Once
	oldValue *Group
}

var _ ent.Mutation = (*GroupMutation)(nil)

// newGroupMutation creates new mutation for $n.Name.
func newGroupMutation(c config, op Op) *GroupMutation {
	return &GroupMutation{
		config:        c,
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetMaxUsers sets the max_users field.
func (m *GroupMutation) SetMaxUsers(i int) {
	m.max_users = &i
	m.addmax_users = nil
}

// MaxUsers returns the max_users value in the mutation.
func (m *GroupMutation) MaxUsers() (r int, exists bool) {
	v := m.max_users
	if v == nil {
		return
	}
	return *v, true
}

// OldMaxUsers returns the old max_users value in the database, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GroupMutation) OldMaxUsers(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMaxUsers is allowed only on UpdateOne operations")
	}
	if m.id == nil {
		return v, fmt.Errorf("OldMaxUsers requires an ID field in the mutation")
	}
	old, err := m.old(ctx, *m.id)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMaxUsers: %w", err)
	}
	return old.MaxUsers, nil
}

// AddMaxUsers adds i to max_users.
func (m *GroupMutation) AddMaxUsers(i int) {
	if m.addmax_users != nil {
		*m.addmax_users += i
	} else {
		m.addmax_users = &i
	}
}

// AddedMaxUsers returns the value that was added to the max_users field in this mutation.
func (m *GroupMutation) AddedMaxUsers() (r int, exists bool) {
	v := m.addmax_users
	if v == nil {
		return
	}
	return *v, true
}

// ResetMaxUsers reset all changes of the "max_users" field.
func (m *GroupMutation) ResetMaxUsers() {
	m.max_users = nil
	m.addmax_users = nil
}

// Op returns the operation name.
func (m *GroupMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.max_users != nil {
		fields = append(fields, group.FieldMaxUsers)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldMaxUsers:
		return m.MaxUsers()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *GroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case group.FieldMaxUsers:
		return m.OldMaxUsers(ctx)
	}
	return nil, fmt.Errorf("unknown Group field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldMaxUsers:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMaxUsers(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GroupMutation) AddedFields() []string {
	var fields []string
	if m.addmax_users != nil {
		fields = append(fields, group.FieldMaxUsers)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case group.FieldMaxUsers:
		return m.AddedMaxUsers()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	case group.FieldMaxUsers:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMaxUsers(v)
		return nil
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldMaxUsers:
		m.ResetMaxUsers()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

// old returns the old Group from the database.
func (m *GroupMutation) old(ctx context.Context, id int) (*Group, error) {
	var err error
	m.oldOnce.Do(func() {
		m.oldValue, err = m.Client().Group.Get(ctx, id)
	})
	return m.oldValue, err
}

// PetMutation represents an operation that mutate the Pets
// nodes in the graph.
type PetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	licensed_at   *time.Time
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	// lazy load old fields.
	oldOnce  sync.Once
	oldValue *Pet
}

var _ ent.Mutation = (*PetMutation)(nil)

// newPetMutation creates new mutation for $n.Name.
func newPetMutation(c config, op Op) *PetMutation {
	return &PetMutation{
		config:        c,
		op:            op,
		typ:           TypePet,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *PetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the age field.
func (m *PetMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *PetMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old age value in the database, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PetMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is allowed only on UpdateOne operations")
	}
	if m.id == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	old, err := m.old(ctx, *m.id)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return old.Age, nil
}

// AddAge adds i to age.
func (m *PetMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *PetMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge reset all changes of the "age" field.
func (m *PetMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetLicensedAt sets the licensed_at field.
func (m *PetMutation) SetLicensedAt(t time.Time) {
	m.licensed_at = &t
}

// LicensedAt returns the licensed_at value in the mutation.
func (m *PetMutation) LicensedAt() (r time.Time, exists bool) {
	v := m.licensed_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLicensedAt returns the old licensed_at value in the database, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PetMutation) OldLicensedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLicensedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil {
		return v, fmt.Errorf("OldLicensedAt requires an ID field in the mutation")
	}
	old, err := m.old(ctx, *m.id)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLicensedAt: %w", err)
	}
	return old.LicensedAt, nil
}

// ClearLicensedAt clears the value of licensed_at.
func (m *PetMutation) ClearLicensedAt() {
	m.licensed_at = nil
	m.clearedFields[pet.FieldLicensedAt] = struct{}{}
}

// LicensedAtCleared returns if the field licensed_at was cleared in this mutation.
func (m *PetMutation) LicensedAtCleared() bool {
	_, ok := m.clearedFields[pet.FieldLicensedAt]
	return ok
}

// ResetLicensedAt reset all changes of the "licensed_at" field.
func (m *PetMutation) ResetLicensedAt() {
	m.licensed_at = nil
	delete(m.clearedFields, pet.FieldLicensedAt)
}

// SetOwnerID sets the owner edge to User by id.
func (m *PetMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the owner edge to User.
func (m *PetMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *PetMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the owner id in the mutation.
func (m *PetMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the owner ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *PetMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner reset all changes of the "owner" edge.
func (m *PetMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *PetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Pet).
func (m *PetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *PetMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, pet.FieldAge)
	}
	if m.licensed_at != nil {
		fields = append(fields, pet.FieldLicensedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *PetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case pet.FieldAge:
		return m.Age()
	case pet.FieldLicensedAt:
		return m.LicensedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *PetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case pet.FieldAge:
		return m.OldAge(ctx)
	case pet.FieldLicensedAt:
		return m.OldLicensedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Pet field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case pet.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case pet.FieldLicensedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLicensedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *PetMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, pet.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *PetMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case pet.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PetMutation) AddField(name string, value ent.Value) error {
	switch name {
	case pet.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Pet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *PetMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(pet.FieldLicensedAt) {
		fields = append(fields, pet.FieldLicensedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *PetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *PetMutation) ClearField(name string) error {
	switch name {
	case pet.FieldLicensedAt:
		m.ClearLicensedAt()
		return nil
	}
	return fmt.Errorf("unknown Pet nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *PetMutation) ResetField(name string) error {
	switch name {
	case pet.FieldAge:
		m.ResetAge()
		return nil
	case pet.FieldLicensedAt:
		m.ResetLicensedAt()
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *PetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, pet.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *PetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case pet.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *PetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *PetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *PetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, pet.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *PetMutation) EdgeCleared(name string) bool {
	switch name {
	case pet.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *PetMutation) ClearEdge(name string) error {
	switch name {
	case pet.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Pet unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *PetMutation) ResetEdge(name string) error {
	switch name {
	case pet.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Pet edge %s", name)
}

// old returns the old Pet from the database.
func (m *PetMutation) old(ctx context.Context, id int) (*Pet, error) {
	var err error
	m.oldOnce.Do(func() {
		m.oldValue, err = m.Client().Pet.Get(ctx, id)
	})
	return m.oldValue, err
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op             Op
	typ            string
	id             *int
	name           *string
	clearedFields  map[string]struct{}
	pets           map[int]struct{}
	removedpets    map[int]struct{}
	friends        map[int]struct{}
	removedfriends map[int]struct{}
	// lazy load old fields.
	oldOnce  sync.Once
	oldValue *User
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value in the database, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	old, err := m.old(ctx, *m.id)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return old.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddPetIDs adds the pets edge to Pet by ids.
func (m *UserMutation) AddPetIDs(ids ...int) {
	if m.pets == nil {
		m.pets = make(map[int]struct{})
	}
	for i := range ids {
		m.pets[ids[i]] = struct{}{}
	}
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (m *UserMutation) RemovePetIDs(ids ...int) {
	if m.removedpets == nil {
		m.removedpets = make(map[int]struct{})
	}
	for i := range ids {
		m.removedpets[ids[i]] = struct{}{}
	}
}

// RemovedPets returns the removed ids of pets.
func (m *UserMutation) RemovedPetsIDs() (ids []int) {
	for id := range m.removedpets {
		ids = append(ids, id)
	}
	return
}

// PetsIDs returns the pets ids in the mutation.
func (m *UserMutation) PetsIDs() (ids []int) {
	for id := range m.pets {
		ids = append(ids, id)
	}
	return
}

// ResetPets reset all changes of the "pets" edge.
func (m *UserMutation) ResetPets() {
	m.pets = nil
	m.removedpets = nil
}

// AddFriendIDs adds the friends edge to User by ids.
func (m *UserMutation) AddFriendIDs(ids ...int) {
	if m.friends == nil {
		m.friends = make(map[int]struct{})
	}
	for i := range ids {
		m.friends[ids[i]] = struct{}{}
	}
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (m *UserMutation) RemoveFriendIDs(ids ...int) {
	if m.removedfriends == nil {
		m.removedfriends = make(map[int]struct{})
	}
	for i := range ids {
		m.removedfriends[ids[i]] = struct{}{}
	}
}

// RemovedFriends returns the removed ids of friends.
func (m *UserMutation) RemovedFriendsIDs() (ids []int) {
	for id := range m.removedfriends {
		ids = append(ids, id)
	}
	return
}

// FriendsIDs returns the friends ids in the mutation.
func (m *UserMutation) FriendsIDs() (ids []int) {
	for id := range m.friends {
		ids = append(ids, id)
	}
	return
}

// ResetFriends reset all changes of the "friends" edge.
func (m *UserMutation) ResetFriends() {
	m.friends = nil
	m.removedfriends = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.pets != nil {
		edges = append(edges, user.EdgePets)
	}
	if m.friends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]ent.Value, 0, len(m.pets))
		for id := range m.pets {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.friends))
		for id := range m.friends {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedpets != nil {
		edges = append(edges, user.EdgePets)
	}
	if m.removedfriends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]ent.Value, 0, len(m.removedpets))
		for id := range m.removedpets {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.removedfriends))
		for id := range m.removedfriends {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePets:
		m.ResetPets()
		return nil
	case user.EdgeFriends:
		m.ResetFriends()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}

// old returns the old User from the database.
func (m *UserMutation) old(ctx context.Context, id int) (*User, error) {
	var err error
	m.oldOnce.Do(func() {
		m.oldValue, err = m.Client().User.Get(ctx, id)
	})
	return m.oldValue, err
}
