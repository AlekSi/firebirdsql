/*******************************************************************************
The MIT License (MIT)

Copyright (c) 2013 Hajime Nakagami

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*******************************************************************************/

package firebirdsql

import (
    "fmt"
    "net"
    "bytes"
    "regexp"
)


var errmsgs = map[int]string {
    335544321 : "arithmetic exception, numeric overflow, or string truncation",
    335544322 : "invalid database key",
    335544323 : "file %g is not a valid database",
    335544324 : "invalid database handle (no active connection)",
    335544325 : "bad parameters on attach or create database",
    335544326 : "unrecognized database parameter block",
    335544327 : "invalid request handle",
    335544328 : "invalid BLOB handle",
    335544329 : "invalid BLOB ID",
    335544330 : "invalid parameter in transaction parameter block",
    335544331 : "invalid format for transaction parameter block",
    335544332 : "invalid transaction handle (expecting explicit transaction start)",
    335544333 : "internal Firebird consistency check (%g)",
    335544334 : "conversion error from string \"%g\"",
    335544335 : "database file appears corrupt (%g)",
    335544336 : "deadlock",
    335544337 : "attempt to start more than %g transactions",
    335544338 : "no match for first value expression",
    335544339 : "information type inappropriate for object specified",
    335544340 : "no information of this type available for object specified",
    335544341 : "unknown information item",
    335544342 : "action cancelled by trigger (%g) to preserve data integrity",
    335544343 : "invalid request BLR at offset %g",
    335544344 : "I/O error during \"%g\" operation for file \"%g\"",
    335544345 : "lock conflict on no wait transaction",
    335544346 : "corrupt system table",
    335544347 : "validation error for column %g, value \"%g\"",
    335544348 : "no current record for fetch operation",
    335544349 : "attempt to store duplicate value (visible to active transactions) in unique index \"%g\"",
    335544350 : "program attempted to exit without finishing database",
    335544351 : "unsuccessful metadata update",
    335544352 : "no permission for %g access to %g %g",
    335544353 : "transaction is not in limbo",
    335544354 : "invalid database key",
    335544355 : "BLOB was not closed",
    335544356 : "metadata is obsolete",
    335544357 : "cannot disconnect database with open transactions (%g active)",
    335544358 : "message length error (encountered %g, expected %g)",
    335544359 : "attempted update of read-only column",
    335544360 : "attempted update of read-only table",
    335544361 : "attempted update during read-only transaction",
    335544362 : "cannot update read-only view %g",
    335544363 : "no transaction for request",
    335544364 : "request synchronization error",
    335544365 : "request referenced an unavailable database",
    335544366 : "segment buffer length shorter than expected",
    335544367 : "attempted retrieval of more segments than exist",
    335544368 : "attempted invalid operation on a BLOB",
    335544369 : "attempted read of a new, open BLOB",
    335544370 : "attempted action on BLOB outside transaction",
    335544371 : "attempted write to read-only BLOB",
    335544372 : "attempted reference to BLOB in unavailable database",
    335544373 : "operating system directive %g failed",
    335544374 : "attempt to fetch past the last record in a record stream",
    335544375 : "unavailable database",
    335544376 : "table %g was omitted from the transaction reserving list",
    335544377 : "request includes a DSRI extension not supported in this implementation",
    335544378 : "feature is not supported",
    335544379 : "unsupported on-disk structure for file %g; found %g.%g, support %g.%g",
    335544380 : "wrong number of arguments on call",
    335544381 : "Implementation limit exceeded",
    335544382 : "%g",
    335544383 : "unrecoverable conflict with limbo transaction %g",
    335544384 : "internal error",
    335544385 : "internal error",
    335544386 : "too many requests",
    335544387 : "internal error",
    335544388 : "block size exceeds implementation restriction",
    335544389 : "buffer exhausted",
    335544390 : "BLR syntax error: expected %g at offset %g, encountered %g",
    335544391 : "buffer in use",
    335544392 : "internal error",
    335544393 : "request in use",
    335544394 : "incompatible version of on-disk structure",
    335544395 : "table %g is not defined",
    335544396 : "column %g is not defined in table %g",
    335544397 : "internal error",
    335544398 : "internal error",
    335544399 : "internal error",
    335544400 : "internal error",
    335544401 : "internal error",
    335544402 : "internal error",
    335544403 : "page %g is of wrong type (expected %g, found %g)",
    335544404 : "database corrupted",
    335544405 : "checksum error on database page %g",
    335544406 : "index is broken",
    335544407 : "database handle not zero",
    335544408 : "transaction handle not zero",
    335544409 : "transaction--request mismatch (synchronization error)",
    335544410 : "bad handle count",
    335544411 : "wrong version of transaction parameter block",
    335544412 : "unsupported BLR version (expected %g, encountered %g)",
    335544413 : "wrong version of database parameter block",
    335544414 : "BLOB and array data types are not supported for %g operation",
    335544415 : "database corrupted",
    335544416 : "internal error",
    335544417 : "internal error",
    335544418 : "transaction in limbo",
    335544419 : "transaction not in limbo",
    335544420 : "transaction outstanding",
    335544421 : "connection rejected by remote interface",
    335544422 : "internal error",
    335544423 : "internal error",
    335544424 : "no lock manager available",
    335544425 : "context already in use (BLR error)",
    335544426 : "context not defined (BLR error)",
    335544427 : "data operation not supported",
    335544428 : "undefined message number",
    335544429 : "bad parameter number",
    335544430 : "unable to allocate memory from operating system",
    335544431 : "blocking signal has been received",
    335544432 : "lock manager error",
    335544433 : "communication error with journal \"%g\"",
    335544434 : "key size exceeds implementation restriction for index \"%g\"",
    335544435 : "null segment of UNIQUE KEY",
    335544436 : "SQL error code = %g",
    335544437 : "wrong DYN version",
    335544438 : "function %g is not defined",
    335544439 : "function %g could not be matched",
    335544440 : "",
    335544441 : "database detach completed with errors",
    335544442 : "database system cannot read argument %g",
    335544443 : "database system cannot write argument %g",
    335544444 : "operation not supported",
    335544445 : "%g extension error",
    335544446 : "not updatable",
    335544447 : "no rollback performed",
    335544448 : "",
    335544449 : "",
    335544450 : "%g",
    335544451 : "update conflicts with concurrent update",
    335544452 : "product %g is not licensed",
    335544453 : "object %g is in use",
    335544454 : "filter not found to convert type %g to type %g",
    335544455 : "cannot attach active shadow file",
    335544456 : "invalid slice description language at offset %g",
    335544457 : "subscript out of bounds",
    335544458 : "column not array or invalid dimensions (expected %g, encountered %g)",
    335544459 : "record from transaction %g is stuck in limbo",
    335544460 : "a file in manual shadow %g is unavailable",
    335544461 : "secondary server attachments cannot validate databases",
    335544462 : "secondary server attachments cannot start journaling",
    335544463 : "generator %g is not defined",
    335544464 : "secondary server attachments cannot start logging",
    335544465 : "invalid BLOB type for operation",
    335544466 : "violation of FOREIGN KEY constraint \"%g\" on table \"%g\"",
    335544467 : "minor version too high found %g expected %g",
    335544468 : "transaction %g is %g",
    335544469 : "transaction marked invalid by I/O error",
    335544470 : "cache buffer for page %g invalid",
    335544471 : "there is no index in table %g with id %g",
    335544472 : "Your user name and password are not defined. Ask your database administrator to set up a Firebird login.",
    335544473 : "invalid bookmark handle",
    335544474 : "invalid lock level %g",
    335544475 : "lock on table %g conflicts with existing lock",
    335544476 : "requested record lock conflicts with existing lock",
    335544477 : "maximum indexes per table (%g) exceeded",
    335544478 : "enable journal for database before starting online dump",
    335544479 : "online dump failure. Retry dump",
    335544480 : "an online dump is already in progress",
    335544481 : "no more disk/tape space.  Cannot continue online dump",
    335544482 : "journaling allowed only if database has Write-ahead Log",
    335544483 : "maximum number of online dump files that can be specified is 16",
    335544484 : "error in opening Write-ahead Log file during recovery",
    335544485 : "invalid statement handle",
    335544486 : "Write-ahead log subsystem failure",
    335544487 : "WAL Writer error",
    335544488 : "Log file header of %g too small",
    335544489 : "Invalid version of log file %g",
    335544490 : "Log file %g not latest in the chain but open flag still set",
    335544491 : "Log file %g not closed properly; database recovery may be required",
    335544492 : "Database name in the log file %g is different",
    335544493 : "Unexpected end of log file %g at offset %g",
    335544494 : "Incomplete log record at offset %g in log file %g",
    335544495 : "Log record header too small at offset %g in log file %g",
    335544496 : "Log block too small at offset %g in log file %g",
    335544497 : "Illegal attempt to attach to an uninitialized WAL segment for %g",
    335544498 : "Invalid WAL parameter block option %g",
    335544499 : "Cannot roll over to the next log file %g",
    335544500 : "database does not use Write-ahead Log",
    335544501 : "cannot drop log file when journaling is enabled",
    335544502 : "reference to invalid stream number",
    335544503 : "WAL subsystem encountered error",
    335544504 : "WAL subsystem corrupted",
    335544505 : "must specify archive file when enabling long term journal for databases with round-robin log files",
    335544506 : "database %g shutdown in progress",
    335544507 : "refresh range number %g already in use",
    335544508 : "refresh range number %g not found",
    335544509 : "CHARACTER SET %g is not defined",
    335544510 : "lock time-out on wait transaction",
    335544511 : "procedure %g is not defined",
    335544512 : "Input parameter mismatch for procedure %g",
    335544513 : "Database %g: WAL subsystem bug for pid %g%g",
    335544514 : "Could not expand the WAL segment for database %g",
    335544515 : "status code %g unknown",
    335544516 : "exception %g not defined",
    335544517 : "exception %g",
    335544518 : "restart shared cache manager",
    335544519 : "invalid lock handle",
    335544520 : "long-term journaling already enabled",
    335544521 : "Unable to roll over please see Firebird log.",
    335544522 : "WAL I/O error.  Please see Firebird log.",
    335544523 : "WAL writer - Journal server communication error.  Please see Firebird log.",
    335544524 : "WAL buffers cannot be increased.  Please see Firebird log.",
    335544525 : "WAL setup error.  Please see Firebird log.",
    335544526 : "obsolete",
    335544527 : "Cannot start WAL writer for the database %g",
    335544528 : "database %g shutdown",
    335544529 : "cannot modify an existing user privilege",
    335544530 : "Cannot delete PRIMARY KEY being used in FOREIGN KEY definition.",
    335544531 : "Column used in a PRIMARY constraint must be NOT NULL.",
    335544532 : "Name of Referential Constraint not defined in constraints table.",
    335544533 : "Non-existent PRIMARY or UNIQUE KEY specified for FOREIGN KEY.",
    335544534 : "Cannot update constraints (RDB$REF_CONSTRAINTS).",
    335544535 : "Cannot update constraints (RDB$CHECK_CONSTRAINTS).",
    335544536 : "Cannot delete CHECK constraint entry (RDB$CHECK_CONSTRAINTS)",
    335544537 : "Cannot delete index segment used by an Integrity Constraint",
    335544538 : "Cannot update index segment used by an Integrity Constraint",
    335544539 : "Cannot delete index used by an Integrity Constraint",
    335544540 : "Cannot modify index used by an Integrity Constraint",
    335544541 : "Cannot delete trigger used by a CHECK Constraint",
    335544542 : "Cannot update trigger used by a CHECK Constraint",
    335544543 : "Cannot delete column being used in an Integrity Constraint.",
    335544544 : "Cannot rename column being used in an Integrity Constraint.",
    335544545 : "Cannot update constraints (RDB$RELATION_CONSTRAINTS).",
    335544546 : "Cannot define constraints on views",
    335544547 : "internal Firebird consistency check (invalid RDB$CONSTRAINT_TYPE)",
    335544548 : "Attempt to define a second PRIMARY KEY for the same table",
    335544549 : "cannot modify or erase a system trigger",
    335544550 : "only the owner of a table may reassign ownership",
    335544551 : "could not find table/procedure for GRANT",
    335544552 : "could not find column for GRANT",
    335544553 : "user does not have GRANT privileges for operation",
    335544554 : "table/procedure has non-SQL security class defined",
    335544555 : "column has non-SQL security class defined",
    335544556 : "Write-ahead Log without shared cache configuration not allowed",
    335544557 : "database shutdown unsuccessful",
    335544558 : "Operation violates CHECK constraint %g on view or table %g",
    335544559 : "invalid service handle",
    335544560 : "database %g shutdown in %g seconds",
    335544561 : "wrong version of service parameter block",
    335544562 : "unrecognized service parameter block",
    335544563 : "service %g is not defined",
    335544564 : "long-term journaling not enabled",
    335544565 : "Cannot transliterate character between character sets",
    335544566 : "WAL defined; Cache Manager must be started first",
    335544567 : "Overflow log specification required for round-robin log",
    335544568 : "Implementation of text subtype %g not located.",
    335544569 : "Dynamic SQL Error",
    335544570 : "Invalid command",
    335544571 : "Data type for constant unknown",
    335544572 : "Invalid cursor reference",
    335544573 : "Data type unknown",
    335544574 : "Invalid cursor declaration",
    335544575 : "Cursor %g is not updatable",
    335544576 : "Attempt to reopen an open cursor",
    335544577 : "Attempt to reclose a closed cursor",
    335544578 : "Column unknown",
    335544579 : "Internal error",
    335544580 : "Table unknown",
    335544581 : "Procedure unknown",
    335544582 : "Request unknown",
    335544583 : "SQLDA missing or incorrect version, or incorrect number/type of variables",
    335544584 : "Count of read-write columns does not equal count of values",
    335544585 : "Invalid statement handle",
    335544586 : "Function unknown",
    335544587 : "Column is not a BLOB",
    335544588 : "COLLATION %g for CHARACTER SET %g is not defined",
    335544589 : "COLLATION %g is not valid for specified CHARACTER SET",
    335544590 : "Option specified more than once",
    335544591 : "Unknown transaction option",
    335544592 : "Invalid array reference",
    335544593 : "Array declared with too many dimensions",
    335544594 : "Illegal array dimension range",
    335544595 : "Trigger unknown",
    335544596 : "Subselect illegal in this context",
    335544597 : "Cannot prepare a CREATE DATABASE/SCHEMA statement",
    335544598 : "must specify column name for view select expression",
    335544599 : "number of columns does not match select list",
    335544600 : "Only simple column names permitted for VIEW WITH CHECK OPTION",
    335544601 : "No WHERE clause for VIEW WITH CHECK OPTION",
    335544602 : "Only one table allowed for VIEW WITH CHECK OPTION",
    335544603 : "DISTINCT, GROUP or HAVING not permitted for VIEW WITH CHECK OPTION",
    335544604 : "FOREIGN KEY column count does not match PRIMARY KEY",
    335544605 : "No subqueries permitted for VIEW WITH CHECK OPTION",
    335544606 : "expression evaluation not supported",
    335544607 : "gen.c: node not supported",
    335544608 : "Unexpected end of command",
    335544609 : "INDEX %g",
    335544610 : "EXCEPTION %g",
    335544611 : "COLUMN %g",
    335544612 : "Token unknown",
    335544613 : "union not supported",
    335544614 : "Unsupported DSQL construct",
    335544615 : "column used with aggregate",
    335544616 : "invalid column reference",
    335544617 : "invalid ORDER BY clause",
    335544618 : "Return mode by value not allowed for this data type",
    335544619 : "External functions cannot have more than 10 parameters",
    335544620 : "alias %g conflicts with an alias in the same statement",
    335544621 : "alias %g conflicts with a procedure in the same statement",
    335544622 : "alias %g conflicts with a table in the same statement",
    335544623 : "Illegal use of keyword VALUE",
    335544624 : "segment count of 0 defined for index %g",
    335544625 : "A node name is not permitted in a secondary, shadow, cache or log file name",
    335544626 : "TABLE %g",
    335544627 : "PROCEDURE %g",
    335544628 : "cannot create index %g",
    335544629 : "Write-ahead Log with shadowing configuration not allowed",
    335544630 : "there are %g dependencies",
    335544631 : "too many keys defined for index %g",
    335544632 : "Preceding file did not specify length, so %g must include starting page number",
    335544633 : "Shadow number must be a positive integer",
    335544634 : "Token unknown - line %g, column %g",
    335544635 : "there is no alias or table named %g at this scope level",
    335544636 : "there is no index %g for table %g",
    335544637 : "table %g is not referenced in plan",
    335544638 : "table %g is referenced more than once in plan; use aliases to distinguish",
    335544639 : "table %g is referenced in the plan but not the from list",
    335544640 : "Invalid use of CHARACTER SET or COLLATE",
    335544641 : "Specified domain or source column %g does not exist",
    335544642 : "index %g cannot be used in the specified plan",
    335544643 : "the table %g is referenced twice; use aliases to differentiate",
    335544644 : "illegal operation when at beginning of stream",
    335544645 : "the current position is on a crack",
    335544646 : "database or file exists",
    335544647 : "invalid comparison operator for find operation",
    335544648 : "Connection lost to pipe server",
    335544649 : "bad checksum",
    335544650 : "wrong page type",
    335544651 : "Cannot insert because the file is readonly or is on a read only medium.",
    335544652 : "multiple rows in singleton select",
    335544653 : "cannot attach to password database",
    335544654 : "cannot start transaction for password database",
    335544655 : "invalid direction for find operation",
    335544656 : "variable %g conflicts with parameter in same procedure",
    335544657 : "Array/BLOB/DATE data types not allowed in arithmetic",
    335544658 : "%g is not a valid base table of the specified view",
    335544659 : "table %g is referenced twice in view; use an alias to distinguish",
    335544660 : "view %g has more than one base table; use aliases to distinguish",
    335544661 : "cannot add index, index root page is full.",
    335544662 : "BLOB SUB_TYPE %g is not defined",
    335544663 : "Too many concurrent executions of the same request",
    335544664 : "duplicate specification of %g - not supported",
    335544665 : "violation of PRIMARY or UNIQUE KEY constraint \"%g\" on table \"%g\"",
    335544666 : "server version too old to support all CREATE DATABASE options",
    335544667 : "drop database completed with errors",
    335544668 : "procedure %g does not return any values",
    335544669 : "count of column list and variable list do not match",
    335544670 : "attempt to index BLOB column in index %g",
    335544671 : "attempt to index array column in index %g",
    335544672 : "too few key columns found for index %g (incorrect column name?)",
    335544673 : "cannot delete",
    335544674 : "last column in a table cannot be deleted",
    335544675 : "sort error",
    335544676 : "sort error: not enough memory",
    335544677 : "too many versions",
    335544678 : "invalid key position",
    335544679 : "segments not allowed in expression index %g",
    335544680 : "sort error: corruption in data structure",
    335544681 : "new record size of %g bytes is too big",
    335544682 : "Inappropriate self-reference of column",
    335544683 : "request depth exceeded. (Recursive definition?)",
    335544684 : "cannot access column %g in view %g",
    335544685 : "dbkey not available for multi-table views",
    335544686 : "journal file wrong format",
    335544687 : "intermediate journal file full",
    335544688 : "The prepare statement identifies a prepare statement with an open cursor",
    335544689 : "Firebird error",
    335544690 : "Cache redefined",
    335544691 : "Insufficient memory to allocate page buffer cache",
    335544692 : "Log redefined",
    335544693 : "Log size too small",
    335544694 : "Log partition size too small",
    335544695 : "Partitions not supported in series of log file specification",
    335544696 : "Total length of a partitioned log must be specified",
    335544697 : "Precision must be from 1 to 18",
    335544698 : "Scale must be between zero and precision",
    335544699 : "Short integer expected",
    335544700 : "Long integer expected",
    335544701 : "Unsigned short integer expected",
    335544702 : "Invalid ESCAPE sequence",
    335544703 : "service %g does not have an associated executable",
    335544704 : "Failed to locate host machine.",
    335544705 : "Undefined service %g/%g.",
    335544706 : "The specified name was not found in the hosts file or Domain Name Services.",
    335544707 : "user does not have GRANT privileges on base table/view for operation",
    335544708 : "Ambiguous column reference.",
    335544709 : "Invalid aggregate reference",
    335544710 : "navigational stream %g references a view with more than one base table",
    335544711 : "Attempt to execute an unprepared dynamic SQL statement.",
    335544712 : "Positive value expected",
    335544713 : "Incorrect values within SQLDA structure",
    335544714 : "invalid blob id",
    335544715 : "Operation not supported for EXTERNAL FILE table %g",
    335544716 : "Service is currently busy: %g",
    335544717 : "stack size insufficent to execute current request",
    335544718 : "Invalid key for find operation",
    335544719 : "Error initializing the network software.",
    335544720 : "Unable to load required library %g.",
    335544721 : "Unable to complete network request to host \"%g\".",
    335544722 : "Failed to establish a connection.",
    335544723 : "Error while listening for an incoming connection.",
    335544724 : "Failed to establish a secondary connection for event processing.",
    335544725 : "Error while listening for an incoming event connection request.",
    335544726 : "Error reading data from the connection.",
    335544727 : "Error writing data to the connection.",
    335544728 : "Cannot deactivate index used by an integrity constraint",
    335544729 : "Cannot deactivate index used by a PRIMARY/UNIQUE constraint",
    335544730 : "Client/Server Express not supported in this release",
    335544731 : "",
    335544732 : "Access to databases on file servers is not supported.",
    335544733 : "Error while trying to create file",
    335544734 : "Error while trying to open file",
    335544735 : "Error while trying to close file",
    335544736 : "Error while trying to read from file",
    335544737 : "Error while trying to write to file",
    335544738 : "Error while trying to delete file",
    335544739 : "Error while trying to access file",
    335544740 : "A fatal exception occurred during the execution of a user defined function.",
    335544741 : "connection lost to database",
    335544742 : "User cannot write to RDB$USER_PRIVILEGES",
    335544743 : "token size exceeds limit",
    335544744 : "Maximum user count exceeded.  Contact your database administrator.",
    335544745 : "Your login %g is same as one of the SQL role name. Ask your database administrator to set up a valid Firebird login.",
    335544746 : "\"REFERENCES table\" without \"(column)\" requires PRIMARY KEY on referenced table",
    335544747 : "The username entered is too long.  Maximum length is 31 bytes.",
    335544748 : "The password specified is too long.  Maximum length is 8 bytes.",
    335544749 : "A username is required for this operation.",
    335544750 : "A password is required for this operation",
    335544751 : "The network protocol specified is invalid",
    335544752 : "A duplicate user name was found in the security database",
    335544753 : "The user name specified was not found in the security database",
    335544754 : "An error occurred while attempting to add the user.",
    335544755 : "An error occurred while attempting to modify the user record.",
    335544756 : "An error occurred while attempting to delete the user record.",
    335544757 : "An error occurred while updating the security database.",
    335544758 : "sort record size of %g bytes is too big",
    335544759 : "can not define a not null column with NULL as default value",
    335544760 : "invalid clause --- '%g'",
    335544761 : "too many open handles to database",
    335544762 : "size of optimizer block exceeded",
    335544763 : "a string constant is delimited by double quotes",
    335544764 : "DATE must be changed to TIMESTAMP",
    335544765 : "attempted update on read-only database",
    335544766 : "SQL dialect %g is not supported in this database",
    335544767 : "A fatal exception occurred during the execution of a blob filter.",
    335544768 : "Access violation.  The code attempted to access a virtual address without privilege to do so.",
    335544769 : "Datatype misalignment.  The attempted to read or write a value that was not stored on a memory boundary.",
    335544770 : "Array bounds exceeded.  The code attempted to access an array element that is out of bounds.",
    335544771 : "Float denormal operand.  One of the floating-point operands is too small to represent a standard float value.",
    335544772 : "Floating-point divide by zero.  The code attempted to divide a floating-point value by zero.",
    335544773 : "Floating-point inexact result.  The result of a floating-point operation cannot be represented as a deciaml fraction.",
    335544774 : "Floating-point invalid operand.  An indeterminant error occurred during a floating-point operation.",
    335544775 : "Floating-point overflow.  The exponent of a floating-point operation is greater than the magnitude allowed.",
    335544776 : "Floating-point stack check.  The stack overflowed or underflowed as the result of a floating-point operation.",
    335544777 : "Floating-point underflow.  The exponent of a floating-point operation is less than the magnitude allowed.",
    335544778 : "Integer divide by zero.  The code attempted to divide an integer value by an integer divisor of zero.",
    335544779 : "Integer overflow.  The result of an integer operation caused the most significant bit of the result to carry.",
    335544780 : "An exception occurred that does not have a description.  Exception number %g.",
    335544781 : "Stack overflow.  The resource requirements of the runtime stack have exceeded the memory available to it.",
    335544782 : "Segmentation Fault. The code attempted to access memory without priviledges.",
    335544783 : "Illegal Instruction. The Code attempted to perfrom an illegal operation.",
    335544784 : "Bus Error. The Code caused a system bus error.",
    335544785 : "Floating Point Error. The Code caused an Arithmetic Exception or a floating point exception.",
    335544786 : "Cannot delete rows from external files.",
    335544787 : "Cannot update rows in external files.",
    335544788 : "Unable to perform operation.  You must be either SYSDBA or owner of the database",
    335544789 : "Specified EXTRACT part does not exist in input datatype",
    335544790 : "Service %g requires SYSDBA permissions.  Reattach to the Service Manager using the SYSDBA account.",
    335544791 : "The file %g is currently in use by another process.  Try again later.",
    335544792 : "Cannot attach to services manager",
    335544793 : "Metadata update statement is not allowed by the current database SQL dialect %g",
    335544794 : "operation was cancelled",
    335544795 : "unexpected item in service parameter block, expected %g",
    335544796 : "Client SQL dialect %g does not support reference to %g datatype",
    335544797 : "user name and password are required while attaching to the services manager",
    335544798 : "You created an indirect dependency on uncommitted metadata. You must roll back the current transaction.",
    335544799 : "The service name was not specified.",
    335544800 : "Too many Contexts of Relation/Procedure/Views. Maximum allowed is 255",
    335544801 : "data type not supported for arithmetic",
    335544802 : "Database dialect being changed from 3 to 1",
    335544803 : "Database dialect not changed.",
    335544804 : "Unable to create database %g",
    335544805 : "Database dialect %g is not a valid dialect.",
    335544806 : "Valid database dialects are %g.",
    335544807 : "SQL warning code = %g",
    335544808 : "DATE data type is now called TIMESTAMP",
    335544809 : "Function %g is in %g, which is not in a permitted directory for external functions.",
    335544810 : "value exceeds the range for valid dates",
    335544811 : "passed client dialect %g is not a valid dialect.",
    335544812 : "Valid client dialects are %g.",
    335544813 : "Unsupported field type specified in BETWEEN predicate.",
    335544814 : "Services functionality will be supported in a later version  of the product",
    335544815 : "GENERATOR %g",
    335544816 : "UDF %g",
    335544817 : "Invalid parameter to FIRST.  Only integers >= 0 are allowed.",
    335544818 : "Invalid parameter to SKIP.  Only integers >= 0 are allowed.",
    335544819 : "File exceeded maximum size of 2GB.  Add another database file or use a 64 bit I/O version of Firebird.",
    335544820 : "Unable to find savepoint with name %g in transaction context",
    335544821 : "Invalid column position used in the %g clause",
    335544822 : "Cannot use an aggregate function in a WHERE clause, use HAVING instead",
    335544823 : "Cannot use an aggregate function in a GROUP BY clause",
    335544824 : "Invalid expression in the %g (not contained in either an aggregate function or the GROUP BY clause)",
    335544825 : "Invalid expression in the %g (neither an aggregate function nor a part of the GROUP BY clause)",
    335544826 : "Nested aggregate functions are not allowed",
    335544827 : "Invalid argument in EXECUTE STATEMENT - cannot convert to string",
    335544828 : "Wrong request type in EXECUTE STATEMENT '%g'",
    335544829 : "Variable type (position %g) in EXECUTE STATEMENT '%g' INTO does not match returned column type",
    335544830 : "Too many recursion levels of EXECUTE STATEMENT",
    335544831 : "Access to %g \"%g\" is denied by server administrator",
    335544832 : "Cannot change difference file name while database is in backup mode",
    335544833 : "Physical backup is not allowed while Write-Ahead Log is in use",
    335544834 : "Cursor is not open",
    335544835 : "Target shutdown mode is invalid for database \"%g\"",
    335544836 : "Concatenation overflow. Resulting string cannot exceed 32K in length.",
    335544837 : "Invalid offset parameter %g to SUBSTRING. Only positive integers are allowed.",
    335544838 : "Foreign key reference target does not exist",
    335544839 : "Foreign key references are present for the record",
    335544840 : "cannot update",
    335544841 : "Cursor is already open",
    335544842 : "%g",
    335544843 : "Context variable %g is not found in namespace %g",
    335544844 : "Invalid namespace name %g passed to %g",
    335544845 : "Too many context variables",
    335544846 : "Invalid argument passed to %g",
    335544847 : "BLR syntax error. Identifier %g... is too long",
    335544848 : "exception %g",
    335544849 : "Malformed string",
    335544850 : "Output parameter mismatch for procedure %g",
    335544851 : "Unexpected end of command - line %g, column %g",
    335544852 : "partner index segment no %g has incompatible data type",
    335544853 : "Invalid length parameter %g to SUBSTRING. Negative integers are not allowed.",
    335544854 : "CHARACTER SET %g is not installed",
    335544855 : "COLLATION %g for CHARACTER SET %g is not installed",
    335544856 : "connection shutdown",
    335544857 : "Maximum BLOB size exceeded",
    335544858 : "Can't have relation with only computed fields or constraints",
    335544859 : "Time precision exceeds allowed range (0-%g)",
    335544860 : "Unsupported conversion to target type BLOB (subtype %g)",
    335544861 : "Unsupported conversion to target type ARRAY",
    335544862 : "Stream does not support record locking",
    335544863 : "Cannot create foreign key constraint %g. Partner index does not exist or is inactive.",
    335544864 : "Transactions count exceeded. Perform backup and restore to make database operable again",
    335544865 : "Column has been unexpectedly deleted",
    335544866 : "%g cannot depend on %g",
    335544867 : "Blob sub_types bigger than 1 (text) are for internal use only",
    335544868 : "Procedure %g is not selectable (it does not contain a SUSPEND statement)",
    335544869 : "Datatype %g is not supported for sorting operation",
    335544870 : "COLLATION %g",
    335544871 : "DOMAIN %g",
    335544872 : "domain %g is not defined",
    335544873 : "Array data type can use up to %g dimensions",
    335544874 : "A multi database transaction cannot span more than %g databases",
    335544875 : "Bad debug info format",
    335544876 : "Error while parsing procedure %g's BLR",
    335544877 : "index key too big",
    335544878 : "concurrent transaction number is %g",
    335544879 : "validation error for variable %g, value \"%g\"",
    335544880 : "validation error for %g, value \"%g\"",
    335544881 : "Difference file name should be set explicitly for database on raw device",
    335544882 : "Login name too long (%g characters, maximum allowed %g)",
    335544883 : "column %g is not defined in procedure %g",
    335544884 : "Invalid SIMILAR TO pattern",
    335544885 : "Invalid TEB format",
    335544886 : "Found more than one transaction isolation in TPB",
    335544887 : "Table reservation lock type %g requires table name before in TPB",
    335544888 : "Found more than one %g specification in TPB",
    335544889 : "Option %g requires READ COMMITTED isolation in TPB",
    335544890 : "Option %g is not valid if %g was used previously in TPB",
    335544891 : "Table name length missing after table reservation %g in TPB",
    335544892 : "Table name length %g is too long after table reservation %g in TPB",
    335544893 : "Table name length %g without table name after table reservation %g in TPB",
    335544894 : "Table name length %g goes beyond the remaining TPB size after table reservation %g",
    335544895 : "Table name length is zero after table reservation %g in TPB",
    335544896 : "Table or view %g not defined in system tables after table reservation %g in TPB",
    335544897 : "Base table or view %g for view %g not defined in system tables after table reservation %g in TPB",
    335544898 : "Option length missing after option %g in TPB",
    335544899 : "Option length %g without value after option %g in TPB",
    335544900 : "Option length %g goes beyond the remaining TPB size after option %g",
    335544901 : "Option length is zero after table reservation %g in TPB",
    335544902 : "Option length %g exceeds the range for option %g in TPB",
    335544903 : "Option value %g is invalid for the option %g in TPB",
    335544904 : "Preserving previous table reservation %g for table %g, stronger than new %g in TPB",
    335544905 : "Table reservation %g for table %g already specified and is stronger than new %g in TPB",
    335544906 : "Table reservation reached maximum recursion of %g when expanding views in TPB",
    335544907 : "Table reservation in TPB cannot be applied to %g because it's a virtual table",
    335544908 : "Table reservation in TPB cannot be applied to %g because it's a system table",
    335544909 : "Table reservation %g or %g in TPB cannot be applied to %g because it's a temporary table",
    335544910 : "Cannot set the transaction in read only mode after a table reservation isc_tpb_lock_write in TPB",
    335544911 : "Cannot take a table reservation isc_tpb_lock_write in TPB because the transaction is in read only mode",
    335544912 : "value exceeds the range for a valid time",
    335544913 : "value exceeds the range for valid timestamps",
    335544914 : "string right truncation",
    335544915 : "blob truncation when converting to a string: length limit exceeded",
    335544916 : "numeric value is out of range",
    335544917 : "Firebird shutdown is still in progress after the specified timeout",
    335544918 : "Attachment handle is busy",
    335544919 : "Bad written UDF detected: pointer returned in FREE_IT function was not allocated by ib_util_malloc",
    335544920 : "External Data Source provider '%g' not found",
    335544921 : "Execute statement error at %g :%gData source : %g",
    335544922 : "Execute statement preprocess SQL error",
    335544923 : "Statement expected",
    335544924 : "Parameter name expected",
    335544925 : "Unclosed comment found near '%g'",
    335544926 : "Execute statement error at %g :%gStatement : %gData source : %g",
    335544927 : "Input parameters mismatch",
    335544928 : "Output parameters mismatch",
    335544929 : "Input parameter '%g' have no value set",
    335544930 : "BLR stream length %g exceeds implementation limit %g",
    335544931 : "Monitoring table space exhausted",
    335544932 : "module name or entrypoint could not be found",
    335544933 : "nothing to cancel",
    335544934 : "ib_util library has not been loaded to deallocate memory returned by FREE_IT function",
    335544935 : "Cannot have circular dependencies with computed fields",
    335544936 : "Security database error",
    335544937 : "Invalid data type in DATE/TIME/TIMESTAMP addition or subtraction in add_datettime()",
    335544938 : "Only a TIME value can be added to a DATE value",
    335544939 : "Only a DATE value can be added to a TIME value",
    335544940 : "TIMESTAMP values can be subtracted only from another TIMESTAMP value",
    335544941 : "Only one operand can be of type TIMESTAMP",
    335544942 : "Only HOUR, MINUTE, SECOND and MILLISECOND can be extracted from TIME values",
    335544943 : "HOUR, MINUTE, SECOND and MILLISECOND cannot be extracted from DATE values",
    335544944 : "Invalid argument for EXTRACT() not being of DATE/TIME/TIMESTAMP type",
    335544945 : "Arguments for %g must be integral types or NUMERIC/DECIMAL without scale",
    335544946 : "First argument for %g must be integral type or floating point type",
    335544947 : "Human readable UUID argument for %g must be of string type",
    335544948 : "Human readable UUID argument for %g must be of exact length %g",
    335544949 : "Human readable UUID argument for %g must have \"-\" at position %g instead of \"%g\"",
    335544950 : "Human readable UUID argument for %g must have hex digit at position %g instead of \"%g\"",
    335544951 : "Only HOUR, MINUTE, SECOND and MILLISECOND can be added to TIME values in %g",
    335544952 : "Invalid data type in addition of part to DATE/TIME/TIMESTAMP in %g",
    335544953 : "Invalid part %g to be added to a DATE/TIME/TIMESTAMP value in %g",
    335544954 : "Expected DATE/TIME/TIMESTAMP type in evlDateAdd() result",
    335544955 : "Expected DATE/TIME/TIMESTAMP type as first and second argument to %g",
    335544956 : "The result of TIME-<value> in %g cannot be expressed in YEAR, MONTH, DAY or WEEK",
    335544957 : "The result of TIME-TIMESTAMP or TIMESTAMP-TIME in %g cannot be expressed in HOUR, MINUTE, SECOND or MILLISECOND",
    335544958 : "The result of DATE-TIME or TIME-DATE in %g cannot be expressed in HOUR, MINUTE, SECOND and MILLISECOND",
    335544959 : "Invalid part %g to express the difference between two DATE/TIME/TIMESTAMP values in %g",
    335544960 : "Argument for %g must be positive",
    335544961 : "Base for %g must be positive",
    335544962 : "Argument #%g for %g must be zero or positive",
    335544963 : "Argument #%g for %g must be positive",
    335544964 : "Base for %g cannot be zero if exponent is negative",
    335544965 : "Base for %g cannot be negative if exponent is not an integral value",
    335544966 : "The numeric scale must be between -128 and 127 in %g",
    335544967 : "Argument for %g must be zero or positive",
    335544968 : "Binary UUID argument for %g must be of string type",
    335544969 : "Binary UUID argument for %g must use %g bytes",
    335544970 : "Missing required item %g in service parameter block",
    335544971 : "%g server is shutdown",
    335544972 : "Invalid connection string",
    335544973 : "Unrecognized events block",
    335544974 : "Could not start first worker thread - shutdown server",
    335544975 : "Timeout occurred while waiting for a secondary connection for event processing",
    335544976 : "Argument for %g must be different than zero",
    335544977 : "Argument for %g must be in the range [-1, 1]",
    335544978 : "Argument for %g must be greater or equal than one",
    335544979 : "Argument for %g must be in the range ]-1, 1[",
    335544980 : "Incorrect parameters provided to internal function %g",
    335544981 : "Floating point overflow in built-in function %g",
    335544982 : "Floating point overflow in result from UDF %g",
    335544983 : "Invalid floating point value returned by UDF %g",
    335544984 : "Database is probably already opened by another engine instance in another Windows session",
    335544985 : "No free space found in temporary directories",
    335544986 : "Explicit transaction control is not allowed",
    335544987 : "Use of TRUSTED switches in spb_command_line is prohibited",
    335545017 : "Asynchronous call is already running for this attachment",
    335740929 : "data base file name (%g) already given",
    335740930 : "invalid switch %g",
    335740932 : "incompatible switch combination",
    335740933 : "replay log pathname required",
    335740934 : "number of page buffers for cache required",
    335740935 : "numeric value required",
    335740936 : "positive numeric value required",
    335740937 : "number of transactions per sweep required",
    335740940 : "\"full\" or \"reserve\" required",
    335740941 : "user name required",
    335740942 : "password required",
    335740943 : "subsystem name",
    335740944 : "\"wal\" required",
    335740945 : "number of seconds required",
    335740946 : "numeric value between 0 and 32767 inclusive required",
    335740947 : "must specify type of shutdown",
    335740948 : "please retry, specifying an option",
    335740951 : "please retry, giving a database name",
    335740991 : "internal block exceeds maximum size",
    335740992 : "corrupt pool",
    335740993 : "virtual memory exhausted",
    335740994 : "bad pool id",
    335740995 : "Transaction state %g not in valid range.",
    335741012 : "unexpected end of input",
    335741018 : "failed to reconnect to a transaction in database %g",
    335741036 : "Transaction description item unknown",
    335741038 : "\"read_only\" or \"read_write\" required",
    335741042 : "positive or zero numeric value required",
    336003074 : "Cannot SELECT RDB$DB_KEY from a stored procedure.",
    336003075 : "Precision 10 to 18 changed from DOUBLE PRECISION in SQL dialect 1 to 64-bit scaled integer in SQL dialect 3",
    336003076 : "Use of %g expression that returns different results in dialect 1 and dialect 3",
    336003077 : "Database SQL dialect %g does not support reference to %g datatype",
    336003079 : "DB dialect %g and client dialect %g conflict with respect to numeric precision %g.",
    336003080 : "WARNING: Numeric literal %g is interpreted as a floating-point",
    336003081 : "value in SQL dialect 1, but as an exact numeric value in SQL dialect 3.",
    336003082 : "WARNING: NUMERIC and DECIMAL fields with precision 10 or greater are stored",
    336003083 : "as approximate floating-point values in SQL dialect 1, but as 64-bit",
    336003084 : "integers in SQL dialect 3.",
    336003085 : "Ambiguous field name between %g and %g",
    336003086 : "External function should have return position between 1 and %g",
    336003087 : "Label %g %g in the current scope",
    336003088 : "Datatypes %gare not comparable in expression %g",
    336003089 : "Empty cursor name is not allowed",
    336003090 : "Statement already has a cursor %g assigned",
    336003091 : "Cursor %g is not found in the current context",
    336003092 : "Cursor %g already exists in the current context",
    336003093 : "Relation %g is ambiguous in cursor %g",
    336003094 : "Relation %g is not found in cursor %g",
    336003095 : "Cursor is not open",
    336003096 : "Data type %g is not supported for EXTERNAL TABLES. Relation '%g', field '%g'",
    336003097 : "Feature not supported on ODS version older than %g.%g",
    336003098 : "Primary key required on table %g",
    336003099 : "UPDATE OR INSERT field list does not match primary key of table %g",
    336003100 : "UPDATE OR INSERT field list does not match MATCHING clause",
    336003101 : "UPDATE OR INSERT without MATCHING could not be used with views based on more than one table",
    336003102 : "Incompatible trigger type",
    336003103 : "Database trigger type can't be changed",
    336068740 : "Table %g already exists",
    336068784 : "column %g does not exist in table/view %g",
    336068796 : "SQL role %g does not exist",
    336068797 : "user %g has no grant admin option on SQL role %g",
    336068798 : "user %g is not a member of SQL role %g",
    336068799 : "%g is not the owner of SQL role %g",
    336068800 : "%g is a SQL role and not a user",
    336068801 : "user name %g could not be used for SQL role",
    336068802 : "SQL role %g already exists",
    336068803 : "keyword %g can not be used as a SQL role name",
    336068804 : "SQL roles are not supported in on older versions of the database.  A backup and restore of the database is required.",
    336068812 : "Cannot rename domain %g to %g.  A domain with that name already exists.",
    336068813 : "Cannot rename column %g to %g.  A column with that name already exists in table %g.",
    336068814 : "Column %g from table %g is referenced in %g",
    336068815 : "Cannot change datatype for column %g.  Changing datatype is not supported for BLOB or ARRAY columns.",
    336068816 : "New size specified for column %g must be at least %g characters.",
    336068817 : "Cannot change datatype for %g.  Conversion from base type %g to %g is not supported.",
    336068818 : "Cannot change datatype for column %g from a character type to a non-character type.",
    336068820 : "Zero length identifiers are not allowed",
    336068829 : "Maximum number of collations per character set exceeded",
    336068830 : "Invalid collation attributes",
    336068840 : "%g cannot reference %g",
    336068852 : "New scale specified for column %g must be at most %g.",
    336068853 : "New precision specified for column %g must be at least %g.",
    336068855 : "Warning: %g on %g is not granted to %g.",
    336068856 : "Feature '%g' is not supported in ODS %g.%g",
    336068857 : "Cannot add or remove COMPUTED from column %g",
    336068858 : "Password should not be empty string",
    336068859 : "Index %g already exists",
    336330753 : "found unknown switch",
    336330754 : "page size parameter missing",
    336330755 : "Page size specified (%g) greater than limit (16384 bytes)",
    336330756 : "redirect location for output is not specified",
    336330757 : "conflicting switches for backup/restore",
    336330758 : "device type %g not known",
    336330759 : "protection is not there yet",
    336330760 : "page size is allowed only on restore or create",
    336330761 : "multiple sources or destinations specified",
    336330762 : "requires both input and output filenames",
    336330763 : "input and output have the same name.  Disallowed.",
    336330764 : "expected page size, encountered \"%g\"",
    336330765 : "REPLACE specified, but the first file %g is a database",
    336330766 : "database %g already exists.  To replace it, use the -REP switch",
    336330767 : "device type not specified",
    336330772 : "gds_$blob_info failed",
    336330773 : "do not understand BLOB INFO item %g",
    336330774 : "gds_$get_segment failed",
    336330775 : "gds_$close_blob failed",
    336330776 : "gds_$open_blob failed",
    336330777 : "Failed in put_blr_gen_id",
    336330778 : "data type %g not understood",
    336330779 : "gds_$compile_request failed",
    336330780 : "gds_$start_request failed",
    336330781 : "gds_$receive failed",
    336330782 : "gds_$release_request failed",
    336330783 : "gds_$database_info failed",
    336330784 : "Expected database description record",
    336330785 : "failed to create database %g",
    336330786 : "RESTORE: decompression length error",
    336330787 : "cannot find table %g",
    336330788 : "Cannot find column for BLOB",
    336330789 : "gds_$create_blob failed",
    336330790 : "gds_$put_segment failed",
    336330791 : "expected record length",
    336330792 : "wrong length record, expected %g encountered %g",
    336330793 : "expected data attribute",
    336330794 : "Failed in store_blr_gen_id",
    336330795 : "do not recognize record type %g",
    336330796 : "Expected backup version 1..9.  Found %g",
    336330797 : "expected backup description record",
    336330798 : "string truncated",
    336330799 : "warning -- record could not be restored",
    336330800 : "gds_$send failed",
    336330801 : "no table name for data",
    336330802 : "unexpected end of file on backup file",
    336330803 : "database format %g is too old to restore to",
    336330804 : "array dimension for column %g is invalid",
    336330807 : "Expected XDR record length",
    336330817 : "cannot open backup file %g",
    336330818 : "cannot open status and error output file %g",
    336330934 : "blocking factor parameter missing",
    336330935 : "expected blocking factor, encountered \"%g\"",
    336330936 : "a blocking factor may not be used in conjunction with device CT",
    336330940 : "user name parameter missing",
    336330941 : "password parameter missing",
    336330952 : " missing parameter for the number of bytes to be skipped",
    336330953 : "expected number of bytes to be skipped, encountered \"%g.\"",
    336330965 : "character set",
    336330967 : "collation",
    336330972 : "Unexpected I/O error while reading from backup file",
    336330973 : "Unexpected I/O error while writing to backup file",
    336330985 : "could not drop database %g (database might be in use)",
    336330990 : "System memory exhausted",
    336331002 : "SQL role",
    336331005 : "SQL role parameter missing",
    336331010 : "page buffers parameter missing",
    336331011 : "expected page buffers, encountered \"%g\"",
    336331012 : "page buffers is allowed only on restore or create",
    336331014 : "size specification either missing or incorrect for file %g",
    336331015 : "file %g out of sequence",
    336331016 : "can't join -- one of the files missing",
    336331017 : " standard input is not supported when using join operation",
    336331018 : "standard output is not supported when using split operation",
    336331019 : "backup file %g might be corrupt",
    336331020 : "database file specification missing",
    336331021 : "can't write a header record to file %g",
    336331022 : "free disk space exhausted",
    336331023 : "file size given (%g) is less than minimum allowed (%g)",
    336331025 : "service name parameter missing",
    336331026 : "Cannot restore over current database, must be SYSDBA or owner of the existing database.",
    336331031 : "\"read_only\" or \"read_write\" required",
    336331033 : "just data ignore all constraints etc.",
    336331034 : "restoring data only ignoring foreign key, unique, not null & other constraints",
    336331093 : "Invalid metadata detected. Use -FIX_FSS_METADATA option.",
    336331094 : "Invalid data detected. Use -FIX_FSS_DATA option.",
    336397205 : "ODS versions before ODS%g are not supported",
    336397206 : "Table %g does not exist",
    336397207 : "View %g does not exist",
    336397208 : "At line %g, column %g",
    336397209 : "At unknown line and column",
    336397210 : "Column %g cannot be repeated in %g statement",
    336397211 : "Too many values (more than %g) in member list to match against",
    336397212 : "Array and BLOB data types not allowed in computed field",
    336397213 : "Implicit domain name %g not allowed in user created domain",
    336397214 : "scalar operator used on field %g which is not an array",
    336397215 : "cannot sort on more than 255 items",
    336397216 : "cannot group on more than 255 items",
    336397217 : "Cannot include the same field (%g.%g) twice in the ORDER BY clause with conflicting sorting options",
    336397218 : "column list from derived table %g has more columns than the number of items in its SELECT statement",
    336397219 : "column list from derived table %g has less columns than the number of items in its SELECT statement",
    336397220 : "no column name specified for column number %g in derived table %g",
    336397221 : "column %g was specified multiple times for derived table %g",
    336397222 : "Internal dsql error: alias type expected by pass1_expand_select_node",
    336397223 : "Internal dsql error: alias type expected by pass1_field",
    336397224 : "Internal dsql error: column position out of range in pass1_union_auto_cast",
    336397225 : "Recursive CTE member (%g) can refer itself only in FROM clause",
    336397226 : "CTE '%g' has cyclic dependencies",
    336397227 : "Recursive member of CTE can't be member of an outer join",
    336397228 : "Recursive member of CTE can't reference itself more than once",
    336397229 : "Recursive CTE (%g) must be an UNION",
    336397230 : "CTE '%g' defined non-recursive member after recursive",
    336397231 : "Recursive member of CTE '%g' has %g clause",
    336397232 : "Recursive members of CTE (%g) must be linked with another members via UNION ALL",
    336397233 : "Non-recursive member is missing in CTE '%g'",
    336397234 : "WITH clause can't be nested",
    336397235 : "column %g appears more than once in USING clause",
    336397236 : "feature is not supported in dialect %g",
    336397237 : "CTE \"%g\" is not used in query",
    336397238 : "column %g appears more than once in ALTER VIEW",
    336397239 : "%g is not supported inside IN AUTONOMOUS TRANSACTION block",
    336397240 : "Unknown node type %g in dsql/GEN_expr",
    336397241 : "Argument for %g in dialect 1 must be string or numeric",
    336397242 : "Argument for %g in dialect 3 must be numeric",
    336397243 : "Strings cannot be added to or subtracted from DATE or TIME types",
    336397244 : "Invalid data type for subtraction involving DATE, TIME or TIMESTAMP types",
    336397245 : "Adding two DATE values or two TIME values is not allowed",
    336397246 : "DATE value cannot be subtracted from the provided data type",
    336397247 : "Strings cannot be added or subtracted in dialect 3",
    336397248 : "Invalid data type for addition or subtraction in dialect 3",
    336397249 : "Invalid data type for multiplication in dialect 1",
    336397250 : "Strings cannot be multiplied in dialect 3",
    336397251 : "Invalid data type for multiplication in dialect 3",
    336397252 : "Division in dialect 1 must be between numeric data types",
    336397253 : "Strings cannot be divided in dialect 3",
    336397254 : "Invalid data type for division in dialect 3",
    336397255 : "Strings cannot be negated (applied the minus operator) in dialect 3",
    336397256 : "Invalid data type for negation (minus operator)",
    336397257 : "Cannot have more than 255 items in DISTINCT list",
    336723983 : "unable to open database",
    336723984 : "error in switch specifications",
    336723985 : "no operation specified",
    336723986 : "no user name specified",
    336723987 : "add record error",
    336723988 : "modify record error",
    336723989 : "find/modify record error",
    336723990 : "record not found for user: %g",
    336723991 : "delete record error",
    336723992 : "find/delete record error",
    336723996 : "find/display record error",
    336723997 : "invalid parameter, no switch defined",
    336723998 : "operation already specified",
    336723999 : "password already specified",
    336724000 : "uid already specified",
    336724001 : "gid already specified",
    336724002 : "project already specified",
    336724003 : "organization already specified",
    336724004 : "first name already specified",
    336724005 : "middle name already specified",
    336724006 : "last name already specified",
    336724008 : "invalid switch specified",
    336724009 : "ambiguous switch specified",
    336724010 : "no operation specified for parameters",
    336724011 : "no parameters allowed for this operation",
    336724012 : "incompatible switches specified",
    336724044 : "Invalid user name (maximum 31 bytes allowed)",
    336724045 : "Warning - maximum 8 significant bytes of password used",
    336724046 : "database already specified",
    336724047 : "database administrator name already specified",
    336724048 : "database administrator password already specified",
    336724049 : "SQL role name already specified",
    336789504 : "The license file does not exist or could not be opened for read",
    336789523 : "operation already specified",
    336789524 : "no operation specified",
    336789525 : "invalid switch",
    336789526 : "invalid switch combination",
    336789527 : "illegal operation/switch combination",
    336789528 : "ambiguous switch",
    336789529 : "invalid parameter, no switch specified",
    336789530 : "switch does not take any parameter",
    336789531 : "switch requires a parameter",
    336789532 : "syntax error in command line",
    336789534 : "The certificate was not added.  A duplicate ID exists in the license file.",
    336789535 : "The certificate was not added.  Invalid certificate ID / Key combination.",
    336789536 : "The certificate was not removed.  The key does not exist or corresponds to a temporary evaluation license.",
    336789537 : "An error occurred updating the license file.  Operation cancelled.",
    336789538 : "The certificate could not be validated based on the information given.  Please recheck the ID and key information.",
    336789539 : "Operation failed.  An unknown error occurred.",
    336789540 : "Add license operation failed, KEY: %g ID: %g",
    336789541 : "Remove license operation failed, KEY: %g",
    336789563 : "The evaluation license has already been used on this server.  You need to purchase a non-evaluation license.",
    336920577 : "found unknown switch",
    336920578 : "please retry, giving a database name",
    336920579 : "Wrong ODS version, expected %g, encountered %g",
    336920580 : "Unexpected end of database file.",
    336920605 : "Can't open database file %g",
    336920606 : "Can't read a database page",
    336920607 : "System memory exhausted",
    336986113 : "Wrong value for access mode",
    336986114 : "Wrong value for write mode",
    336986115 : "Wrong value for reserve space",
    336986116 : "Unknown tag (%g) in info_svr_db_info block after isc_svc_query()",
    336986117 : "Unknown tag (%g) in isc_svc_query() results",
    336986118 : "Unknown switch \"%g\"",
    336986159 : "Wrong value for shutdown mode",
    336986160 : "could not open file %g",
    336986161 : "could not read file %g",
    336986162 : "empty file %g",
    336986164 : "Invalid or missing parameter for switch %g",
    337051649 : "Switches trusted_svc and trusted_role are not supported from command line",
}


type wirepPotocol struct {
    buf []byte
    buffer_len int
    bufCount int

    conn net.Conn
    dbHandle int32
    addr string
    dbname string
    user string
    password string
}

func NewWireProtocol (dsn string) *wireProtocol {
    p := new(wireProtocol)
    p.buffer_len = 1024
    p.buf, err = make([]byte, p.buffer_len)

    dsnPattern := regexp.MustCompile(
        `^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
            `(?:\((?P<addr>[^\)]*)\)?` +            // [(addr)]
            `\/(?P<dbname>.*?)`)                    // /dbname

    p.addr = "127.0.0.1"
    for i, match := range matches {
        switch names[i] {
        case "user":
            p.user = match
        case "passwd":
            p.passwd = match
        case "addr":
            p.addr = match
        case "dbname":
            p.dbname = match
        }
    }
    if strings.ContainsRune(p.addr, ':') {
        p.addr += ":3050"
    }

    var err error
    p.conn, err = net.Dial("tcp", p.addr)

    return p, err
}

func (p *wireProtocol) packInt(i int32) {
    // pack big endian int32
    p.buf[p.bufCount+0] = byte(i >> 24 & 0xFF)
    p.buf[p.bufCount+1] = byte(i >> 16 & 0xFF)
    p.buf[p.bufCount+2] = byte(i >> 8 & 0xFF)
    p.buf[p.bufCount+3] = byte(i & 0xFF)
    p.bufCount += 4
}

func (p *wireProtocol) packBytes(b []byte) {
    for _, b := range xdrBytes(b) {
        p.buf[p.bufCount] = b
        p.bufCount++
    }
}

func (p *wireProtocol) packString(s string) {
    for _, b := range xdrString(s) {
        p.buf[p.bufCount] = b
        p.bufCount++
    }
}

func (p *wireProtocol) appendBytes(bs [] byte) {
    for _, b := range bs {
        p.buf[p.bufCount] = b
        p.bufCount++
    }
}

func (p *wireProtocol) uid() string {
    // TODO:
    return "Firebird Go Driver"
}

func (p *wireProtocol) sendPackets() (n int, err error) {
    n, err = p.conn.Write(p.buf)
    return
}

func (p *wireProtocol) recvPackets(n int) ([]byte, error) {
    buf, err := make([]byte, n)
    return p.conn.Read(buf)
}


    def _parse_status_vector(self):
        sql_code = 0
        gds_codes = set()
        message = ''
        n = bytes_to_bint(recv_channel(self.sock, 4))
        while n != isc_arg_end:
            if n == isc_arg_gds:
                gds_code = bytes_to_bint(recv_channel(self.sock, 4))
                if gds_code:
                    gds_codes.add(gds_code)
                    message += messages.get(gds_code, '@1')
                    num_arg = 0
            elif n == isc_arg_number:
                num = bytes_to_bint(recv_channel(self.sock, 4))
                if gds_code == 335544436:
                    sql_code = num
                num_arg += 1
                message = message.replace('@' + str(num_arg), str(num))
            elif n == isc_arg_string or n == isc_arg_interpreted:
                nbytes = bytes_to_bint(recv_channel(self.sock, 4))
                n = str(recv_channel(self.sock, nbytes, word_alignment=True))
                num_arg += 1
                message = message.replace('@' + str(num_arg), n)
            n = bytes_to_bint(recv_channel(self.sock, 4))

        return (gds_codes, sql_code, message)



func (p *wireProtocol) _parse_op_response() {
    b = recv_channel(self.sock, 16)
    h = bytes_to_bint(b[0:4])         # Object handle
    oid = b[4:12]                       # Object ID
    buf_len = bytes_to_bint(b[12:])   # buffer length
    buf = recv_channel(self.sock, buf_len, word_alignment=True)

    (gds_codes, sql_code, message) = self._parse_status_vector()
    if sql_code or message:
        raise OperationalError(message, gds_codes, sql_code)

    return (h, oid, buf)


func (p *wireProtocol) opConnect() {
    p.packInt(op_connect)
    p.packInt(op_attach)
    p.packInt(2)   // CONNECT_VERSION2
    p.packInt(1)   // Arch type (Generic = 1)
    p.packString(bytes.NewBufferString(p.dbname))
    p.packInt(1)   // Protocol version understood count.
    p.pack_bytes(p.uid())
    p.packInt(10)  // PROTOCOL_VERSION10
    p.packInt(1)   // Arch type (Generic = 1)
    p.packInt(2)   // Min type
    p.packInt(3)   // Max type
    p.packInt(2)   // Preference weight
    p.sendPackets()
}


func (p *wireProtocol) opCreate() {
    page_size := 4096

    encode := bytes.NewBufferString("UTF8").Bytes()
    user := bytes.NewBufferString(p.user).Bytes()
    password := bytes.NewBufferString(p.password).Bytes()
    dpb := bytes.Join([][]byte{
        []byte{1},
        []byte{68, len(encode)}, encode,
        []byte{48, len(encode)}, encode,
        []byte{28, len(user)}, user,
        []byte{29, len(password)}, password,
        []byte{63, 4}, int32_to_byte(3),
        []byte{24, 4}, bint32_to_byte(1),
        []byte{54, 4}, bint32_to_byte(1),
        []byte{4, 4}, int32_to_byte(page_size),
    }, nil)

    p = xdrlib.Packer()
    p.packInt(op_create)
    p.packInt(0)                       // Database Object ID
    p.packString(p.dbName)
    p.packBytes(dpb)
    p.sendPackets()
}

func (p *wireProtocol) opAccept() {
    b, err = p.recvPackets(4)
    for {
        if bytes_to_bint(b) == op_dummy {
            b, err = p.recvPackets(4)
        }
    }

    // assert bytes_to_bint(b) == op_accept
    b = p.recvPackets(12)
    // assert up.unpack_int() == 10
    // assert  up.unpack_int() == 1
    // assert up.unpack_int() == 3
}

func (p *wireProtocol) opAttach() {
    encode := bytes.NewBufferString("UTF8").Bytes()
    user := bytes.NewBufferString(p.user).Bytes()
    password := bytes.NewBufferString(p.password).Bytes()

    dbp := bytes.Join([][]byte{
        []byte{1},
        []byte{48, len(encode)}, encode,
        []byte{28, len(user)}, user,
        []byte{29, len(password)}, password,
    })
    p.packInt(op_attach)
    p.packInt(0)                       // Database Object ID
    p.packString(p.dbName)
    p.pack_bytes(dpb)
    p.sendPackets()
}

func (p *wireProtocol) opDropDatabase() {
    p.packInt(op_drop_database)
    p.packInt(p.dbHandle)
    p.sendPackets()
}


func (p *wireProtocol) opTransaction(tpb []byte) {
    p.packInt(op_transaction)
    p.packInt(p.dbHandle)
    p.packBytes(tpb)
    p.sendPackets()
}

func (p *wireProtcol) opCommit(transHandle int32) {
    p.pack_int(op_commit)
    p.pack_int(transHandle)
    p.sendPackets()
}

func (p *wireProtcol) opCommitRetaining(transHandle int32) {
    p.pack_int(op_commit_retaining)
    p.pack_int(transHandle)
    p.sendPackets()
}

func (p *wireProtocol) opRollback(transHandle int32) {
    p.pack_int(op_rollback)
    p.pack_int(transHandle)
    p.sendPackets()
}

func (p *wireProtocol) opRollbackRetaining(transHandle int32) {
    p.packInt(op_rollback_retaining)
    p.packInt(transHandle)
    p.sendPackets()
}

func (p *wireProtocol) opAallocateStatement() {
    p.packInt(op_allocate_statement)
    p.packInt(p.dbHandle)
    p.sendPackets()
}

func (p *wireProtocol) opInfoTransaction(transHandle int32 , b []byte) {
    p.packInt(op_info_transaction)
    p.packInt(transHandle)
    p.packInt(0)
    p.packBytes(b)
    p.packInt(p.buffer_length)
    p.sendPackets()
}

func (p *wireProtocol) opInfoDatabase(bs []byte) {
    p.packInt(op_info_database)
    p.packInt(p.dbHandle)
    p.packInt(0)
    p.packBytes(bs)
    p.packInt(p.buffer_length)
    p.sendPackets()
}

func (p *wireProtocol) opFreeStatement(stmtHandle int32, mode int32) {
    p.packInt(op_free_statement)
    p.packInt(stmtHandle)
    p.packInt(mode)
    p.sendPackets()
}

func (p *wireProtocol) opPrepareStatement(stmtHandle int32, transHandle int32, query string) {

    descItems := []byte{
        isc_info_sql_stmt_type,
        isc_info_sql_select,
        isc_info_sql_describe_vars,
        isc_info_sql_sqlda_seq,
        isc_info_sql_type,
        isc_info_sql_sub_type,
        isc_info_sql_scale,
        isc_info_sql_length,
        isc_info_sql_null_ind,
        isc_info_sql_field,
        isc_info_sql_relation,
        isc_info_sql_owner,
        isc_info_sql_alias,
        isc_info_sql_describe_end,
    }

    p.packInt(op_prepare_statement)
    p.packInt(transHandle)
    p.packInt(stmtHandle)
    p.packInt(3)                        // dialect = 3
    p.packString(query)
    p.packBytes(descItems)
    p.packInt(p.buffer_length)
    p.sendPackets()
}

func (p *wireProtocol) opInfoSql(stmtHandle int32, vars []byte) {
    p.pack_int(self.op_info_sql)
    p.pack_int(stmtHandle)
    p.pack_int(0)
    p.pack_bytes(vars)
    p.pack_int(p.buffer_length)
    p.sendPackets()
}

func (p *wireProtocol) opExecute(stmtHandle int32, transHandle int32, params []interface{}) {
    p.pack_int(op_execute)
    p.pack_int(stmtHandle)
    p.pack_int(transHandle)

    if len(params) == 0 {
        p.packInt(0)        // packBytes([])
        p.packInt(0)
        p.packInt(0)
        p.sendPackets()
    } else {
        blr, values := params_to_blr(params)
        p.packBytes(blr)
        p.packInt(0)
        p.packInt(1)
        p.appendBytes(values)
        p.sendPackets()
    }
}

func (p *wireProtocol) opExecute2(stmtHandle int32, transHandle int32, params []interface{}, outputBlr []byte) {
    p.pack_int(op_execute2)
    p.pack_int(stmtHandle)
    p.pack_int(transHandle)

    if len(params) == 0 {
        p.packInt(0)        // packBytes([])
        p.packInt(0)
        p.packInt(0)
    } else {
        blr, values := params_to_blr(params)
        p.packBytes(blr)
        p.packInt(0)
        p.packInt(1)
        p.appendBytes(values)
    }

    p.packBytes(outputBlr)
    p.packInt(0)
    p.sendPackets()
}

func (p *wireProtocol)  opFetch(stmtHandle int32, blr []byte) {
    p.packInt(op_fetch)
    p.packInt(stmtHandle)
    p.packBytes(blr)
    p.packInt(0)
    p.packInt(400)
    p.sendPackets()
}

func (p *wireProtocol) opFetchResponse(stmtHandle int32, xsqlda []xSQLVAR) {
    b, err = p.recvPackets(4)
    for {
        if bytes_to_bint(b) == op_dummy {
            b, err = p.recvPackets(4)
        }
    }


    // TODO:
    if bytes_to_bint(b) == self.op_response:
        return self._parse_op_response()    # error occured
    if bytes_to_bint(b) != self.op_fetch_response:
        raise InternalError
    b = recv_channel(self.sock, 8)
    status = bytes_to_bint(b[:4])
    count = bytes_to_bint(b[4:8])
    rows = []
    while count:
        r = [None] * len(xsqlda)
        for i in range(len(xsqlda)):
            x = xsqlda[i]
            if x.io_length() < 0:
                b = recv_channel(self.sock, 4)
                ln = bytes_to_bint(b)
            else:
                ln = x.io_length()
            raw_value = recv_channel(self.sock, ln, word_alignment=True)
            if recv_channel(self.sock, 4) == bytes([0]) * 4: # Not NULL
                r[i] = x.value(raw_value)
        rows.append(r)
        b = recv_channel(self.sock, 12)
        op = bytes_to_bint(b[:4])
        status = bytes_to_bint(b[4:8])
        count = bytes_to_bint(b[8:])
    return rows, status != 100
}

func (p *wireProtocol) opDetach() {
    p.packInt(self.op_detach)
    p.packInt(self.db_handle)
    p.sendPackets()
}

func (p *wireProtocol)  opOpenBlob(blob_id int32, transHandle int32) {
    p.packInt(self.op_open_blob)
    p.packInt(transHandle)
    p.appendPacket(blog_id)
    p.sendPackets()
}

func (p *wireProtocol)  opCreateBlob2(transHandle int32) {
    p.packInt(op_create_blob2)
    p.packInt(0)
    p.packInt(transHandle)
    p.packInt(0)
    p.packInt(0)
    p.sendPackets()
}

func (p *wireProtocol) opGetSegment(blobHandle int32) {
    p.pack_int(self.op_get_segment)
    p.pack_int(blobHandle)
    p.pack_int(self.buffer_length)
    p.pack_int(0)
    p.sendPackets()
}

func (p *wireProtocol) opBatchSegments(blobHandle, seg_data) {
    ln = len(seg_data)
    p.packInt(self.op_batch_segments)
    p.packInt(blobHandle)
    p.packInt(ln + 2)
    p.packInt(ln + 2)
    pad_length = ((4-(ln+2)) & 3)
    send_channel(self.sock, p.get_buffer() 
            + int_to_bytes(ln, 2) + seg_data + bytes([0])*pad_length)
}

func (p *wireProtocol)  opCloseBlob(blobHandle) {
    p.packInt(op_close_blob)
    p.packInt(blobHandle)
    p.sendPackets()
}

func (p *wireProtocol) opConnectRequest() {
    p.packInt(op_connect_request)
    p.packInt(1)    // async
    p.packInt(p.dbHandle)
    p.packInt(0)
    p.sendPackets()

    b, err = p.recvPackets(4)
    for bytes_to_bint(b) == op_dummy {
        b, err = p.recvPackets(4)
    }

    if bytes_to_bint(b) != self.op_response:
        raise InternalError

    h = bytes_to_bint(recv_channel(self.sock, 4))
    recv_channel(self.sock, 8)  # garbase
    ln = bytes_to_bint(recv_channel(self.sock, 4))
    ln += ln % 4    # padding
    family = bytes_to_bint(recv_channel(self.sock, 2))
    port = bytes_to_bint(recv_channel(self.sock, 2), u=True)
    b = recv_channel(self.sock, 4)
    ip_address = '.'.join([str(byte_to_int(c)) for c in b])
    ln -= 8
    recv_channel(self.sock, ln)

    (gds_codes, sql_code, message) = self._parse_status_vector()
    if sql_code or message:
        raise OperationalError(message, gds_codes, sql_code)

    return (h, port, family, ip_address)
}

func (p *wireProtocol) _op_response() {
    b = recv_channel(self.sock, 4)
    while bytes_to_bint(b) == self.op_dummy:
        b = recv_channel(self.sock, 4)
    if bytes_to_bint(b) != self.op_response:
        raise InternalError
    return self._parse_op_response()
}

func (p *wireProtocol) _op_sql_response(xsqlda) {
    b = recv_channel(self.sock, 4)
    while bytes_to_bint(b) == self.op_dummy:
        b = recv_channel(self.sock, 4)
    if bytes_to_bint(b) != self.op_sql_response:
        raise InternalError

    b = recv_channel(self.sock, 4)
    count = bytes_to_bint(b[:4])

    r = []
    for i in range(len(xsqlda)):
        x = xsqlda[i]
        if x.io_length() < 0:
            b = recv_channel(self.sock, 4)
            ln = bytes_to_bint(b)
        else:
            ln = x.io_length()
        raw_value = recv_channel(self.sock, ln, word_alignment=True)
        if recv_channel(self.sock, 4) == bytes([0]) * 4: # Not NULL
            r.append(x.value(raw_value))
        else:
            r.append(None)

    b = recv_channel(self.sock, 32)     # ??? why 32 bytes skip

    return r
}
