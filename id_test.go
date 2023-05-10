package id

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	type test struct {
		name         string
		input        []string
		expectedType IDType
	}

	tables := []test{
		{
			name: "User",
			input: []string{
				"u_2PbxA3jxwlfu6WMV8B45QaFA6nH",
				"u_2PbxA5xIuKz6zeudIBfWHZfzNHN",
				"u_2PbxA3a9gPGQO1vlJ7x1BJOLEhZ",
			},
			expectedType: User,
		},
		{
			name: "Company",
			input: []string{
				"c_2PbxHSJYsiaWyuprke168Do5vxy",
				"c_2PbxHUMUJBaQTTstfMOhG0oKtA9",
				"c_2PbxHUtqHcp5FhLwuKc1ntARjmn",
			},
			expectedType: Company,
		},
		{
			name: "None",
			input: []string{
				"2PbxV6hn3ceANtvwr6Rs7lRgwra",
				"2PbxV8MUKuqeOBgrVXtDPkVB5h2",
				"2PbxV1iCfsulW1KxLcBurfRjaWX",
			},
			expectedType: None,
		},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			for _, in := range tt.input {
				id, err := Parse(in)
				if err != nil {
					t.Error("Expected err to be nil. got: ", err)
				}
				if tt.expectedType != id.GetType() {
					t.Errorf("Expected %s to equal %s", id.GetType(), tt.expectedType)
				}
			}
		})
	}
}

func TestNewWithTime(t *testing.T) {
	id1, err := NewWithTime(User, time.Now().Add(-time.Minute))
	if err != nil {
		t.Error("Expected err to be nil. got: ", err)
	}
	id2, err := NewWithTime(User, time.Now().Add(-2*time.Minute))
	if err != nil {
		t.Error("Expected err to be nil. got: ", err)
	}
	id3, err := NewWithTime(User, time.Now().Add(-3*time.Minute))
	if err != nil {
		t.Error("Expected err to be nil. got: ", err)
	}

	if id1.Time().Unix() < id2.Time().Unix() {
		t.Errorf("Expected %v to be greator than %v", id1.Time().Unix(), id2.Time().Unix())
	}
	if id2.Time().Unix() < id3.Time().Unix() {
		t.Errorf("Expected %v to be greator than %v", id2.Time().Unix(), id3.Time().Unix())
	}
}
