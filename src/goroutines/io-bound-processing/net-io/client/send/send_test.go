package send

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestPrimSendRequests(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			_, ok := r.(*url.Error)
			if ok {
				fmt.Println("You probably need to start the server.\n Error:\n", r)
				return
			}
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	howMany := 2

	responses := primSendRequests(howMany)

	// test the number of responses received
	numberOfResponses := len(responses)
	if numberOfResponses != howMany {
		t.Errorf("expected %d responses, got %d", numberOfResponses, howMany)
	}

	// test that the responses are correct
	for _, response := range responses {
		if !strings.Contains(response, "Request") {
			t.Errorf("response \"%v\" does not seem right", response)
		}
	}

}

func TestSplit_100_11(t *testing.T) {
	howMany := 100
	concurrent := 11

	expected := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 10}

	split := split(howMany, concurrent)

	// test that the spit is correct
	if len(split) != len(expected) {
		t.Errorf("number of splits %v not equal to numberof expected %v", len(split), len(expected))
	}
	for i := 0; i < concurrent; i++ {
		if split[i] != expected[i] {
			t.Errorf("expected %d splits, got %d", expected[i], split[i])
		}
	}

}
func TestSplit_999_10(t *testing.T) {
	howMany := 999
	concurrent := 10

	expected := []int{99, 99, 99, 99, 99, 99, 99, 99, 99, 108}

	split := split(howMany, concurrent)

	// test that the spit is correct
	if len(split) != len(expected) {
		t.Errorf("number of splits %v not equal to numberof expected %v", len(split), len(expected))
	}
	for i := 0; i < concurrent; i++ {
		if split[i] != expected[i] {
			t.Errorf("expected %d splits, got %d", expected[i], split[i])
		}
	}

}
func TestSplit_100(t *testing.T) {
	howMany := 100
	concurrent := 10

	split := split(howMany, concurrent)

	testSplit(howMany, concurrent, split, t)

}
func TestSplit_101(t *testing.T) {
	howMany := 101
	concurrent := 10

	split := split(howMany, concurrent)

	testSplit(howMany, concurrent, split, t)

}
func TestSplit_99(t *testing.T) {
	howMany := 99
	concurrent := 10

	split := split(howMany, concurrent)

	testSplit(howMany, concurrent, split, t)

}
func testSplit(howMany int, concurrent int, split []int, t *testing.T) {
	// test the number of splits
	if len(split) != concurrent {
		t.Errorf("expected %d splits, got %d", concurrent, len(split))
	}

	// test that the first concurrent - 1 splits are the same
	for i := 0; i < concurrent-1; i++ {
		if split[i] != howMany/concurrent {
			t.Errorf("expected %d splits, got %d", howMany/concurrent, split[i])
		}
	}

	// test that the sum of the splits is the same as howMany
	sum := 0
	for _, split := range split {
		sum += split
	}
	if sum != howMany {
		t.Errorf("expected %d splits, got %d", howMany, sum)
	}
}

func TestCallServer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			_, ok := r.(*url.Error)
			if ok {
				fmt.Println("You probably need to start the server.\n Error:\n", r)
				return
			}
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	howMany := 4
	concurrent := 2

	requestsProcessed := CallServer(howMany, concurrent)

	// test that the number of requests processed is correct
	if requestsProcessed != howMany {
		t.Errorf("expected %d requests processed, got %d", howMany, requestsProcessed)
	}

}
