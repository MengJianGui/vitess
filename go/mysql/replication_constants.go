/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysql

// This file contains the constant definitions for this package.

// This is the data type for a field.
// Values taken from include/mysql/mysql_com.h
const (
	// TypeDecimal is MYSQL_TYPE_DECIMAL. It is deprecated.
	TypeDecimal = 0

	// TypeTiny is MYSQL_TYPE_TINY
	TypeTiny = 1

	// TypeShort is MYSQL_TYPE_SHORT
	TypeShort = 2

	// TypeLong is MYSQL_TYPE_LONG
	TypeLong = 3

	// TypeFloat is MYSQL_TYPE_FLOAT
	TypeFloat = 4

	// TypeDouble is MYSQL_TYPE_DOUBLE
	TypeDouble = 5

	// TypeNull is MYSQL_TYPE_NULL
	TypeNull = 6

	// TypeTimestamp is MYSQL_TYPE_TIMESTAMP
	TypeTimestamp = 7

	// TypeLongLong is MYSQL_TYPE_LONGLONG
	TypeLongLong = 8

	// TypeInt24 is MYSQL_TYPE_INT24
	TypeInt24 = 9

	// TypeDate is MYSQL_TYPE_DATE
	TypeDate = 10

	// TypeTime is MYSQL_TYPE_TIME
	TypeTime = 11

	// TypeDateTime is MYSQL_TYPE_DATETIME
	TypeDateTime = 12

	// TypeYear is MYSQL_TYPE_YEAR
	TypeYear = 13

	// TypeNewDate is MYSQL_TYPE_NEWDATE
	TypeNewDate = 14

	// TypeVarchar is MYSQL_TYPE_VARCHAR
	TypeVarchar = 15

	// TypeBit is MYSQL_TYPE_BIT
	TypeBit = 16

	// TypeTimestamp2 is MYSQL_TYPE_TIMESTAMP2
	TypeTimestamp2 = 17

	// TypeDateTime2 is MYSQL_TYPE_DATETIME2
	TypeDateTime2 = 18

	// TypeTime2 is MYSQL_TYPE_TIME2
	TypeTime2 = 19

	// TypeJSON is MYSQL_TYPE_JSON
	TypeJSON = 245

	// TypeNewDecimal is MYSQL_TYPE_NEWDECIMAL
	TypeNewDecimal = 246

	// TypeEnum is MYSQL_TYPE_ENUM
	TypeEnum = 247

	// TypeSet is MYSQL_TYPE_SET
	TypeSet = 248

	// TypeTinyBlob is MYSQL_TYPE_TINY_BLOB
	TypeTinyBlob = 249

	// TypeMediumBlob is MYSQL_TYPE_MEDIUM_BLOB
	TypeMediumBlob = 250

	// TypeLongBlob is MYSQL_TYPE_LONG_BLOB
	TypeLongBlob = 251

	// TypeBlob is MYSQL_TYPE_BLOB
	TypeBlob = 252

	// TypeVarString is MYSQL_TYPE_VAR_STRING
	TypeVarString = 253

	// TypeString is MYSQL_TYPE_STRING
	TypeString = 254

	// TypeGeometry is MYSQL_TYPE_GEOMETRY
	TypeGeometry = 255
)

// Constants for the type of an INTVAR_EVENT.
const (
	// IntVarInvalidInt is INVALID_INT_EVENT
	IntVarInvalidInt = 0

	// IntVarLastInsertID is LAST_INSERT_ID_EVENT
	IntVarLastInsertID = 1

	// IntVarInsertID is INSERT_ID_EVENT
	IntVarInsertID = 2
)

// Name of the variable represented by an IntVar.
var (
	// IntVarNames maps a InVar type to the variable name it represents.
	IntVarNames = map[byte]string{
		IntVarLastInsertID: "LAST_INSERT_ID",
		IntVarInsertID:     "INSERT_ID",
	}
)

// Constants about the type of checksum in a packet.
// These constants are common between MariaDB 10.0 and MySQL 5.6.
const (
	// BinlogChecksumAlgOff indicates that checksums are supported but off.
	BinlogChecksumAlgOff = 0

	// BinlogChecksumAlgCRC32 indicates that CRC32 checksums are used.
	BinlogChecksumAlgCRC32 = 1

	// BinlogChecksumAlgUndef indicates that checksums are not supported.
	BinlogChecksumAlgUndef = 255
)

// These constants describe the event types.
// See: http://dev.mysql.com/doc/internals/en/binlog-event-type.html
const (
	//eUnknownEvent           = 0
	// Unused
	//eStartEventV3           = 1
	eQueryEvent  = 2
	eStopEvent   = 3
	eRotateEvent = 4
	eIntVarEvent = 5
	// Unused
	//eLoadEvent              = 6
	// Unused
	//eSlaveEvent             = 7
	// Unused
	//eCreateFileEvent        = 8
	// Unused
	//eAppendBlockEvent       = 9
	//eExecLoadEvent          = 10
	// Unused
	//eDeleteFileEvent        = 11
	// Unused
	//eNewLoadEvent           = 12
	eRandEvent = 13
	// Unused
	//eUserVarEvent           = 14
	eFormatDescriptionEvent = 15
	eXIDEvent               = 16
	//Unused
	//eBeginLoadQueryEvent    = 17
	//Unused
	//eExecuteLoadQueryEvent  = 18
	eTableMapEvent     = 19
	eWriteRowsEventV0  = 20
	eUpdateRowsEventV0 = 21
	eDeleteRowsEventV0 = 22
	eWriteRowsEventV1  = 23
	eUpdateRowsEventV1 = 24
	eDeleteRowsEventV1 = 25
	// Unused
	//eIncidentEvent          = 26
	//eHeartbeatEvent         = 27
	// Unused
	//eIgnorableEvent         = 28
	// Unused
	//eRowsQueryEvent         = 29
	eWriteRowsEventV2   = 30
	eUpdateRowsEventV2  = 31
	eDeleteRowsEventV2  = 32
	eGTIDEvent          = 33
	eAnonymousGTIDEvent = 34
	ePreviousGTIDsEvent = 35

	// MySQL 5.7 events. Unused.
	//eTransactionContextEvent = 36
	//eViewChangeEvent         = 37
	//eXAPrepareLogEvent       = 38

	// Transaction_payload_event when binlog compression is turned on
	eCompressedEvent = 40

	// MariaDB specific values. They start at 160.
	//eMariaAnnotateRowsEvent = 160
	// Unused
	//eMariaBinlogCheckpointEvent = 161
	eMariaGTIDEvent     = 162
	eMariaGTIDListEvent = 163
	// Unused
	//eMariaStartEncryptionEvent  = 164
)

// These constants describe the type of status variables in q Query packet.
const (
	// QFlags2Code is Q_FLAGS2_CODE
	QFlags2Code = 0

	// QSQLModeCode is Q_SQL_MODE_CODE
	QSQLModeCode = 1

	// QCatalog is Q_CATALOG
	QCatalog = 2

	// QAutoIncrement is Q_AUTO_INCREMENT
	QAutoIncrement = 3

	// QCharsetCode is Q_CHARSET_CODE
	QCharsetCode = 4

	// QTimeZoneCode is Q_TIME_ZONE_CODE
	QTimeZoneCode = 5

	// QCatalogNZCode is Q_CATALOG_NZ_CODE
	QCatalogNZCode = 6
)
