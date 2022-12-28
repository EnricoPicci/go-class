package main

import (
	"sort"
	"testing"
)

func TestTriangleRowGenerator_10(t *testing.T) {
	rowsCh := TriangleRowGenerator(10)

	rows := [][]int{}

	for r := range rowsCh {
		rows = append(rows, r)
	}

	var expectedRows = [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	checkRows(expectedRows, rows, t)

}

func TestTriangleRowGenerator_11(t *testing.T) {
	rowsCh := TriangleRowGenerator(11)

	rows := [][]int{}

	for r := range rowsCh {
		rows = append(rows, r)
	}

	var expectedRows = [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
		{11, 0, 0, 0, 0},
	}
	checkRows(expectedRows, rows, t)

}

func checkRows(expectedRows [][]int, gotRows [][]int, t *testing.T) {
	expectedNumOfRows := len(expectedRows)
	gotNumOfRows := len(gotRows)
	if gotNumOfRows != expectedNumOfRows {
		t.Fatalf("expected %v, got %v", expectedNumOfRows, gotNumOfRows)
	}

	for i, gotRow := range gotRows {
		var expectedRow = expectedRows[i]
		if len(expectedRow) != len(gotRow) {
			t.Errorf("Row %v ==>> expected %v elements, got %v elements == row got %v", i, len(expectedRow), len(gotRow), gotRow)
		}
		for j, v := range gotRow {
			if expectedRow[j] != v {
				t.Errorf("Row %v at position %v ==>> expected %v, got %v", i, j, expectedRow[j], v)
			}
		}
	}
}

func TestNewWelcomeEnvelope(t *testing.T) {
	row := []int{7, 8, 9, 10}
	busNumber := 1
	passengerNumber := 4

	welcomeEnvelope := NewWelcomeEnvelope(busNumber, passengerNumber, row)

	gotRoomNumber := welcomeEnvelope.RoomNumber
	expectedRoomNumber := 10
	if gotRoomNumber != expectedRoomNumber {
		t.Errorf("Room number ==>> expected %v, got %v", expectedRoomNumber, gotRoomNumber)
	}
}

func TestWelcomeEnvelopesForBus(t *testing.T) {
	rowsCh := TriangleRowGenerator(11)
	hilbertCh := make(chan WelcomeEnvelope)
	firstBusNumber := 1

	go WelcomeEnvelopesForBus(rowsCh, hilbertCh, firstBusNumber)

	envelopes := []WelcomeEnvelope{}
	for envelope := range hilbertCh {
		envelopes = append(envelopes, envelope)
	}

	expectedEnvelopes := []WelcomeEnvelope{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{1, 4, 10},
		{2, 1, 2},
		{2, 2, 5},
		{2, 3, 9},
		{3, 1, 4},
		{3, 2, 8},
		{4, 1, 7},
		{5, 1, 11},
	}

	expectedNumOfEnvelopes := len(expectedEnvelopes)
	gotNumOfEnvelopes := len(envelopes)
	if gotNumOfEnvelopes != expectedNumOfEnvelopes {
		t.Errorf("expected %v, got %v", expectedNumOfEnvelopes, gotNumOfEnvelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i].BusNumber == envelopes[j].BusNumber {
			return envelopes[i].PassengerNumber < envelopes[j].PassengerNumber
		}
		return envelopes[i].BusNumber < envelopes[j].BusNumber
	})

	for i, gotEnvelop := range envelopes {
		expectedEnvelop := expectedEnvelopes[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in envelope %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}

func TestGoHilbert(t *testing.T) {
	GoHilbert(16)
}
