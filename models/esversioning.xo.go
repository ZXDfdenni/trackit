// Package models contains the types for schema 'trackit'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// EsVersioning represents a row from 'trackit.es_versioning'.
type EsVersioning struct {
	ID             int    `json:"id"`              // id
	CurrentVersion int    `json:"current_version"` // current_version
	TemplateName   string `json:"template_name"`   // template_name
	IndexName      string `json:"index_name"`      // index_name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the EsVersioning exists in the database.
func (ev *EsVersioning) Exists() bool {
	return ev._exists
}

// Deleted provides information if the EsVersioning has been deleted from the database.
func (ev *EsVersioning) Deleted() bool {
	return ev._deleted
}

// Insert inserts the EsVersioning to the database.
func (ev *EsVersioning) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ev._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO trackit.es_versioning (` +
		`current_version, template_name, index_name` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ev.CurrentVersion, ev.TemplateName, ev.IndexName)
	res, err := db.Exec(sqlstr, ev.CurrentVersion, ev.TemplateName, ev.IndexName)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ev.ID = int(id)
	ev._exists = true

	return nil
}

// Update updates the EsVersioning in the database.
func (ev *EsVersioning) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ev._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ev._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE trackit.es_versioning SET ` +
		`current_version = ?, template_name = ?, index_name = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ev.CurrentVersion, ev.TemplateName, ev.IndexName, ev.ID)
	_, err = db.Exec(sqlstr, ev.CurrentVersion, ev.TemplateName, ev.IndexName, ev.ID)
	return err
}

// Save saves the EsVersioning to the database.
func (ev *EsVersioning) Save(db XODB) error {
	if ev.Exists() {
		return ev.Update(db)
	}

	return ev.Insert(db)
}

// Delete deletes the EsVersioning from the database.
func (ev *EsVersioning) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ev._exists {
		return nil
	}

	// if deleted, bail
	if ev._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM trackit.es_versioning WHERE id = ?`

	// run query
	XOLog(sqlstr, ev.ID)
	_, err = db.Exec(sqlstr, ev.ID)
	if err != nil {
		return err
	}

	// set deleted
	ev._deleted = true

	return nil
}

// EsVersioningByID retrieves a row from 'trackit.es_versioning' as a EsVersioning.
//
// Generated from index 'es_versioning_id_pkey'.
func EsVersioningByID(db XODB, id int) (*EsVersioning, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, current_version, template_name, index_name ` +
		`FROM trackit.es_versioning ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ev := EsVersioning{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ev.ID, &ev.CurrentVersion, &ev.TemplateName, &ev.IndexName)
	if err != nil {
		return nil, err
	}

	return &ev, nil
}