// Package models contains the types for schema 'trackit'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// UserUpdateTaggingComplianceJob represents a row from 'trackit.user_update_tagging_compliance_job'.
type UserUpdateTaggingComplianceJob struct {
	ID        int       `json:"id"`        // id
	Created   time.Time `json:"created"`   // created
	UserID    int       `json:"user_id"`   // user_id
	Completed time.Time `json:"completed"` // completed
	WorkerID  string    `json:"worker_id"` // worker_id
	JobError  string    `json:"job_error"` // job_error

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserUpdateTaggingComplianceJob exists in the database.
func (uutcj *UserUpdateTaggingComplianceJob) Exists() bool {
	return uutcj._exists
}

// Deleted provides information if the UserUpdateTaggingComplianceJob has been deleted from the database.
func (uutcj *UserUpdateTaggingComplianceJob) Deleted() bool {
	return uutcj._deleted
}

// Insert inserts the UserUpdateTaggingComplianceJob to the database.
func (uutcj *UserUpdateTaggingComplianceJob) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if uutcj._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO trackit.user_update_tagging_compliance_job (` +
		`created, user_id, completed, worker_id, job_error` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, uutcj.Created, uutcj.UserID, uutcj.Completed, uutcj.WorkerID, uutcj.JobError)
	res, err := db.Exec(sqlstr, uutcj.Created, uutcj.UserID, uutcj.Completed, uutcj.WorkerID, uutcj.JobError)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	uutcj.ID = int(id)
	uutcj._exists = true

	return nil
}

// Update updates the UserUpdateTaggingComplianceJob in the database.
func (uutcj *UserUpdateTaggingComplianceJob) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uutcj._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if uutcj._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE trackit.user_update_tagging_compliance_job SET ` +
		`created = ?, user_id = ?, completed = ?, worker_id = ?, job_error = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, uutcj.Created, uutcj.UserID, uutcj.Completed, uutcj.WorkerID, uutcj.JobError, uutcj.ID)
	_, err = db.Exec(sqlstr, uutcj.Created, uutcj.UserID, uutcj.Completed, uutcj.WorkerID, uutcj.JobError, uutcj.ID)
	return err
}

// Save saves the UserUpdateTaggingComplianceJob to the database.
func (uutcj *UserUpdateTaggingComplianceJob) Save(db XODB) error {
	if uutcj.Exists() {
		return uutcj.Update(db)
	}

	return uutcj.Insert(db)
}

// Delete deletes the UserUpdateTaggingComplianceJob from the database.
func (uutcj *UserUpdateTaggingComplianceJob) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !uutcj._exists {
		return nil
	}

	// if deleted, bail
	if uutcj._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM trackit.user_update_tagging_compliance_job WHERE id = ?`

	// run query
	XOLog(sqlstr, uutcj.ID)
	_, err = db.Exec(sqlstr, uutcj.ID)
	if err != nil {
		return err
	}

	// set deleted
	uutcj._deleted = true

	return nil
}

// User returns the User associated with the UserUpdateTaggingComplianceJob's UserID (user_id).
//
// Generated from foreign key 'user_update_tagging_compliance_job_ibfk_1'.
func (uutcj *UserUpdateTaggingComplianceJob) User(db XODB) (*User, error) {
	return UserByID(db, uutcj.UserID)
}

// UserUpdateTaggingComplianceJobsByUserID retrieves a row from 'trackit.user_update_tagging_compliance_job' as a UserUpdateTaggingComplianceJob.
//
// Generated from index 'foreign_user'.
func UserUpdateTaggingComplianceJobsByUserID(db XODB, userID int) ([]*UserUpdateTaggingComplianceJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, created, user_id, completed, worker_id, job_error ` +
		`FROM trackit.user_update_tagging_compliance_job ` +
		`WHERE user_id = ?`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*UserUpdateTaggingComplianceJob{}
	for q.Next() {
		uutcj := UserUpdateTaggingComplianceJob{
			_exists: true,
		}

		// scan
		err = q.Scan(&uutcj.ID, &uutcj.Created, &uutcj.UserID, &uutcj.Completed, &uutcj.WorkerID, &uutcj.JobError)
		if err != nil {
			return nil, err
		}

		res = append(res, &uutcj)
	}

	return res, nil
}

// UserUpdateTaggingComplianceJobByID retrieves a row from 'trackit.user_update_tagging_compliance_job' as a UserUpdateTaggingComplianceJob.
//
// Generated from index 'user_update_tagging_compliance_job_id_pkey'.
func UserUpdateTaggingComplianceJobByID(db XODB, id int) (*UserUpdateTaggingComplianceJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, created, user_id, completed, worker_id, job_error ` +
		`FROM trackit.user_update_tagging_compliance_job ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	uutcj := UserUpdateTaggingComplianceJob{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&uutcj.ID, &uutcj.Created, &uutcj.UserID, &uutcj.Completed, &uutcj.WorkerID, &uutcj.JobError)
	if err != nil {
		return nil, err
	}

	return &uutcj, nil
}
