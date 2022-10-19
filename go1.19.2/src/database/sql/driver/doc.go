// Package driver defines interfaces to be implemented by database
// drivers as used by package sql.
//
// Most code should use package sql.
//
// The driver interface has evolved over time. Drivers should implement
// Connector and DriverContext interfaces.
// The Connector.Connect and Driver.Open methods should never return ErrBadConn.
// ErrBadConn should only be returned from Validator, SessionResetter, or
// a query method if the connection is already in an invalid (e.g. closed) state.
//
// All Conn implementations should implement the following interfaces:
// Pinger, SessionResetter, and Validator.
//
// If named parameters or context are supported, the driver's Conn should implement:
// ExecerContext, QueryerContext, ConnPrepareContext, and ConnBeginTx.
//
// To support custom data types, implement NamedValueChecker. NamedValueChecker
// also allows queries to accept per-query options as a parameter by returning
// ErrRemoveArgument from CheckNamedValue.
//
// If multiple result sets are supported, Rows should implement RowsNextResultSet.
// If the driver knows how to describe the types present in the returned result
// it should implement the following interfaces: RowsColumnTypeScanType,
// RowsColumnTypeDatabaseTypeName, RowsColumnTypeLength, RowsColumnTypeNullable,
// and RowsColumnTypePrecisionScale. A given row value may also return a Rows
// type, which may represent a database cursor value.
//
// Before a connection is returned to the connection pool after use, IsValid is
// called if implemented. Before a connection is reused for another query,
// ResetSession is called if implemented. If a connection is never returned to the
// connection pool but immediately reused, then ResetSession is called prior to
// reuse but IsValid is not called.
package driver
