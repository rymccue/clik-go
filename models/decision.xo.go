// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"

	"github.com/lib/pq"
)

// Decision represents a row from 'public.decisions'.
type Decision struct {
	FromUserID int         `json:"from_user_id"` // from_user_id
	ToUserID   int         `json:"to_user_id"`   // to_user_id
	Likes      bool        `json:"likes"`        // likes
	CreatedAt  pq.NullTime `json:"created_at"`   // created_at
	UpdatedAt  pq.NullTime `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Decision exists in the database.
func (d *Decision) Exists() bool {
	return d._exists
}

// Deleted provides information if the Decision has been deleted from the database.
func (d *Decision) Deleted() bool {
	return d._deleted
}

// Insert inserts the Decision to the database.
func (d *Decision) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.decisions (` +
		`from_user_id, to_user_id, likes, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)`

	// run query
	XOLog(sqlstr, d.FromUserID, d.ToUserID, d.Likes, d.CreatedAt, d.UpdatedAt)
	err = db.QueryRow(sqlstr, d.FromUserID, d.ToUserID, d.Likes, d.CreatedAt, d.UpdatedAt).Scan(&d.ToUserID)
	if err != nil {
		return err
	}

	// set existence
	d._exists = true

	return nil
}

// Update updates the Decision in the database.
func (d *Decision) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if d._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.decisions SET (` +
		`from_user_id, likes, created_at, updated_at` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE to_user_id = $5`

	// run query
	XOLog(sqlstr, d.FromUserID, d.Likes, d.CreatedAt, d.UpdatedAt, d.ToUserID)
	_, err = db.Exec(sqlstr, d.FromUserID, d.Likes, d.CreatedAt, d.UpdatedAt, d.ToUserID)
	return err
}

// Save saves the Decision to the database.
func (d *Decision) Save(db XODB) error {
	if d.Exists() {
		return d.Update(db)
	}

	return d.Insert(db)
}

// Upsert performs an upsert for Decision.
//
// NOTE: PostgreSQL 9.5+ only
func (d *Decision) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.decisions (` +
		`from_user_id, to_user_id, likes, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (to_user_id) DO UPDATE SET (` +
		`from_user_id, to_user_id, likes, created_at, updated_at` +
		`) = (` +
		`EXCLUDED.from_user_id, EXCLUDED.to_user_id, EXCLUDED.likes, EXCLUDED.created_at, EXCLUDED.updated_at` +
		`)`

	// run query
	XOLog(sqlstr, d.FromUserID, d.ToUserID, d.Likes, d.CreatedAt, d.UpdatedAt)
	_, err = db.Exec(sqlstr, d.FromUserID, d.ToUserID, d.Likes, d.CreatedAt, d.UpdatedAt)
	if err != nil {
		return err
	}

	// set existence
	d._exists = true

	return nil
}

// Delete deletes the Decision from the database.
func (d *Decision) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return nil
	}

	// if deleted, bail
	if d._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.decisions WHERE to_user_id = $1`

	// run query
	XOLog(sqlstr, d.ToUserID)
	_, err = db.Exec(sqlstr, d.ToUserID)
	if err != nil {
		return err
	}

	// set deleted
	d._deleted = true

	return nil
}

// UserByFromUserID returns the User associated with the Decision's FromUserID (from_user_id).
//
// Generated from foreign key 'decisions_from_user_id_fkey'.
func (d *Decision) UserByFromUserID(db XODB) (*User, error) {
	return UserByID(db, d.FromUserID)
}

// UserByToUserID returns the User associated with the Decision's ToUserID (to_user_id).
//
// Generated from foreign key 'decisions_to_user_id_fkey'.
func (d *Decision) UserByToUserID(db XODB) (*User, error) {
	return UserByID(db, d.ToUserID)
}

// DecisionByFromUserIDToUserID retrieves a row from 'public.decisions' as a Decision.
//
// Generated from index 'pk_decision'.
func DecisionByFromUserIDToUserID(db XODB, fromUserID int, toUserID int) (*Decision, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`from_user_id, to_user_id, likes, created_at, updated_at ` +
		`FROM public.decisions ` +
		`WHERE from_user_id = $1 AND to_user_id = $2`

	// run query
	XOLog(sqlstr, fromUserID, toUserID)
	d := Decision{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, fromUserID, toUserID).Scan(&d.FromUserID, &d.ToUserID, &d.Likes, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &d, nil
}