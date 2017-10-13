/*
**  Copyright(C) 2017, StepToSky
**
**  Redistribution and use in source and binary forms, with or without
**  modification, are permitted provided that the following conditions are met:
**
**  1.Redistributions of source code must retain the above copyright notice, this
**    list of conditions and the following disclaimer.
**  2.Redistributions in binary form must reproduce the above copyright notice,
**    this list of conditions and the following disclaimer in the documentation
**    and / or other materials provided with the distribution.
**  3.Neither the name of StepToSky nor the names of its contributors
**    may be used to endorse or promote products derived from this software
**    without specific prior written permission.
**
**  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
**  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
**  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
**  DISCLAIMED.IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
**  ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
**  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
**  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
**  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
**  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
**  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
**  Contacts: www.steptosky.com
*/

package errt

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/

// Enable/disable printing source file name
var EnableSourceNamePrint bool = true

// Enable/disable printing full source file name
var EnableFullSourceName bool = false

// Represents traceable errors.
// It contains information where the error occurred.
type Error struct {
	desc string
	err  error
	file string
	line int
	ok   bool
}

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/

// Creates new error from an origin go error and a description.
func New(err error, desc string) *Error {
	return _newErr(err, desc)
}

// Creates new error from an origin go error and a formatted description.
func Newf(err error, format string, a ...interface{}) *Error {
	return _newErr(err, fmt.Sprintf(format, a))
}

//----------------------

// Creates new error from a description.
func NewDesc(desc string) *Error {
	return _newErr(nil, desc)
}

// Creates new error from a formatted description.
func NewDescf(format string, a ...interface{}) *Error {
	return _newErr(nil, fmt.Sprintf(format, a))
}

//----------------------

// Creates new error from an origin go error.
func NewFrom(err error) *Error {
	return _newErr(err, "")
}

//----------------------

func _newErr(err error, desc string) *Error {
	const CallerLevelUp int = 2
	_, file, line, ok := runtime.Caller(CallerLevelUp)
	return &Error{err: err, desc: desc, file: simplifySourcePath(&file), line: line, ok: ok}
}

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/

func simplifySourcePath(path *string) string {
	if !EnableFullSourceName {
		idx := strings.LastIndex(*path, "/")
		if idx == -1 {
			return *path
		}
		return (*path)[idx+1 : len(*path)]
	} else {
		return *path
	}
}

func (err *Error) format() *string {
	var out string
	descAvailable := len(err.desc) != 0
	if EnableSourceNamePrint {
		if err.ok {
			out = "\n\t\tTrace: " + err.file + ":" + strconv.FormatInt(int64(err.line), 10)
		}
		if descAvailable {
			out += "\n\t\t\tError: "
		}
	} else {
		if descAvailable {
			out += "\n\t\tError: "
		}
	}

	if descAvailable {
		out = out + err.desc
	}
	return &out
}

func (err *Error) Error() string {
	out := err.format()
	if err.err != nil {
		_, ok := err.err.(*Error)
		if ok {
			return *out + err.err.Error()
		} else {
			if EnableSourceNamePrint {
				return *out + "\n\t\t\tError: " + err.err.Error() + "\n"
			} else {
				return *out + "\n\t\tError: " + err.err.Error() + "\n"
			}
		}
	}
	return *out + "\n"
}

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/
