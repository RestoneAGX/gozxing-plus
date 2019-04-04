package gozxing

import (
	"testing"

	errors "golang.org/x/xerrors"
)

func testFormatExceptionType(t *testing.T, e error) {
	var fe FormatException
	if !errors.As(e, &fe) {
		t.Fatalf("Type must be FormatException")
	}
	var re ReaderException
	if !errors.As(e, &re) {
		t.Fatalf("Type must be ReaderException")
	}
	var ne NotFoundException
	if errors.As(e, &ne) {
		t.Fatalf("Type must not be NotFoundException")
	}

	if _, ok := e.(FormatException); !ok {
		t.Fatalf("Type must be FormatException")
	}
	if _, ok := e.(ReaderException); !ok {
		t.Fatalf("Type must be ReaderException")
	}
	if _, ok := e.(ChecksumException); ok {
		t.Fatalf("Type must not be ChecksumException")
	}

	fe.formatException()
	fe.readerException()
}

func TestFormatException(t *testing.T) {
	var e error = NewFormatException()
	testFormatExceptionType(t, e)
}

func TestNewFormatException(t *testing.T) {
	var e error = NewFormatException("testmsg %d, %d", 10, 20)
	testFormatExceptionType(t, e)

	msg := e.Error()
	wants := "FormatException: testmsg 10, 20"
	if msg != wants {
		t.Fatalf("Error() = \"%s\", wants \"%s\"", msg, wants)
	}
}

func TestWrapFormatException(t *testing.T) {
	base := errors.New("newformatexception")
	var e error = WrapFormatException(base)

	testFormatExceptionType(t, e)

	if !errors.Is(e, base) {
		t.Fatalf("err(%v) is not base(%v)", e, base)
	}
}
