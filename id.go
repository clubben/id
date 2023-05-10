package id

// Used to create, manage, and parse unique IDs. This code creates stripe-like IDs,
// (e.g., xx_000000000000000000000000000). It allows for a prefix and a 27
// character ksuid separated by an underscore. The prefix makes it easy
// to visibly identify what the ID is for.
//
// Original Author: Kirk Morales, https://gist.github.com/knation/3541b4da1c5274eaf03ceafa6985bd0a

import (
	"fmt"
	"strings"

	"github.com/segmentio/ksuid"
)

// Define enum for ID types
type IDType int32

const (
	None IDType = iota
	User
	Company
)

// `idData` contains data for each TypeId
var idData = map[IDType]([]string){
	None:    {"", "none"},
	User:    {"u", "user"},
	Company: {"c", "company"},
}

// Used as a reverse lookup of TypeId by prefix
var (
	prefixLookup     map[string]IDType
	prefixLookupInit = false
)

// Populates `prefixLookup`
func createPrefixLookup() {
	prefixLookup = make(map[string]IDType)
	for key, val := range idData {
		prefixLookup[val[0]] = key
	}
}

// `String()` method for `TypeId`
func (typeId IDType) String() string {
	return idData[typeId][1]
}

type ID struct {
	t  IDType
	id ksuid.KSUID
}

// Creates an ID string for the given object type
func New(typeId IDType) ID {
	return ID{
		t:  typeId,
		id: ksuid.New(),
	}
}

func Parse(id string) (ID, error) {
	parts := strings.Split(id, "_")
	if len(parts) == 1 {
		parsed, err := ksuid.Parse(parts[0])
		if err != nil {
			return ID{}, err
		}

		return ID{
			t:  None,
			id: parsed,
		}, nil
	} else {
		parsed, err := ksuid.Parse(parts[1])
		if err != nil {
			return ID{}, err
		}

		if !prefixLookupInit {
			createPrefixLookup()
			prefixLookupInit = true
		}

		typeId := prefixLookup[parts[0]]

		return ID{
			t:  typeId,
			id: parsed,
		}, nil
	}
}

func (i ID) String() string {
	if i.t == None {
		return i.id.String()
	}
	prefix := idData[i.t][0]
	return fmt.Sprintf("%s_%s", prefix, i.id.String())
}

func (i ID) GetType() IDType {
	return i.t
}
