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
	"errors"
	"fmt"
	"testing"
)

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/

func TestErrt_case1(t *testing.T) {
	fmt.Println("case 1")
	err0 := errors.New("err0")
	err1 := New(err0, "err1")
	err2 := New(err1, "err2")
	err3 := New(err2, "err3")
	msg := "\n\t\tTrace: Error_test.go:22\n\t\t\tError: err3" +
		"\n\t\tTrace: Error_test.go:21\n\t\t\tError: err2" +
		"\n\t\tTrace: Error_test.go:20\n\t\t\tError: err1" +
		"\n\t\t\tError: err0\n"
	if msg != err3.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err3.Error())
	}
	//fmt.Println(err3.Error())
}

func TestErrt_case2(t *testing.T) {
	fmt.Println("case 2")
	err0 := NewDesc("err0")
	err1 := New(err0, "err1")
	err2 := New(err1, "err2")
	err3 := New(err2, "err3")
	msg := "\n\t\tTrace: Error_test.go:38\n\t\t\tError: err3" +
		"\n\t\tTrace: Error_test.go:37\n\t\t\tError: err2" +
		"\n\t\tTrace: Error_test.go:36\n\t\t\tError: err1" +
		"\n\t\tTrace: Error_test.go:35\n\t\t\tError: err0\n"
	if msg != err3.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err3.Error())
	}
	//fmt.Println(err3.Error())
}

func TestErrt_case3(t *testing.T) {
	fmt.Println("case 3")
	EnableSourceNamePrint = false
	err0 := NewDesc("err0")
	err1 := New(err0, "err1")
	err2 := New(err1, "err2")
	err3 := New(err2, "err3")
	msg := "\n\t\tError: err3" +
		"\n\t\tError: err2" +
		"\n\t\tError: err1" +
		"\n\t\tError: err0\n"
	if msg != err3.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err3.Error())
	}
	//fmt.Println(err3.Error())
}

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/

func TestErrt_case4(t *testing.T) {
	fmt.Println("case 4")
	EnableSourceNamePrint = false
	err0 := NewDescf("%s", "err0")
	msg := "\n\t\tError: [err0]\n"
	if msg != err0.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err0.Error())
	}
	//fmt.Println(err0.Error())
}

func TestErrt_case5(t *testing.T) {
	fmt.Println("case 5")
	EnableSourceNamePrint = false
	err0 := Newf(errors.New("err0"), "%s", "err1")
	msg := "\n\t\tError: [err1]\n" + "\t\tError: err0\n"
	if msg != err0.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err0.Error())
	}
	//fmt.Println(err0.Error())
}

func TestErrt_case6(t *testing.T) {
	fmt.Println("case 5")
	EnableSourceNamePrint = false
	err0 := NewFrom(errors.New("err0"))
	msg := "\n\t\tError: err0\n"
	if msg != err0.Error() {
		t.Errorf("\nExpected: %s \ngot: %s", msg, err0.Error())
	}
	//fmt.Println(err0.Error())
}

/*********************************************************************************************************/
///////////////////////////////////////////////////////////////////////////////////////////////////////////
/*********************************************************************************************************/
