package jsondb_test

import (
	"testing"

	"github.com/pt-arvind/gocleanarchitecture/domain"
	"github.com/pt-arvind/gocleanarchitecture/repository"
)

// TestMockService ensures the mock service works correctly.
func TestMockService(t *testing.T) {
	// Test the reading and writing.
	s := new(repository.MockService)

	// Test forced failures.
	s.WriteFail = true
	s.ReadFail = true
	AssertNotNil(t, s.Read())
	AssertNotNil(t, s.Write())

	// Test no failures.
	s.WriteFail = false
	s.ReadFail = false
	AssertEqual(t, s.Read(), nil)
	AssertEqual(t, s.Write(), nil)

	// Test adding a record and reading it.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	s.AddRecord(*u)
	AssertEqual(t, len(s.Records()), 1)
}
