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
    "database/sql/driver"
)

type firebirdsqlConn struct{
    wp *wireProtocol
    tx *firebirdsqlTx
    addr string
    dbName string
    user string
    passwd string
}

func (fc *firebirdsqlConn) Begin() (driver.Tx, error) {
    tx, err := newFirebirdsqlTx(fc.wp)
    return tx, err
}


func (fc *firebirdsqlConn) Close() (err error) {
    return
}

func (fc *firebirdsqlConn) Prepare(query string) (driver.Stmt, error) {
    var err error
    return nil, err
}

func (fc *firebirdsqlConn) Exec(query string, args []driver.Value) (driver.Result, error) {
    var err error
    return nil, err
}

func (fc *firebirdsqlConn) Query(query string, args []driver.Value) (driver.Rows, error) {
    var err error
    return nil, err
}

func newFirebirdsqlConn(wp *wireProtocol, addr string, dbName string, user string, passwd string) (*firebirdsqlConn, error) {
    fc, err := newFirebirdsqlConn(wp, addr, dbName, user, passwd)
    fc.tx, err = newFirebirdsqlTx(wp)

    return fc, err
}
