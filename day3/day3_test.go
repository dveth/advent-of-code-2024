package main

import "testing"

func TestRemovePrefix(t *testing.T) {
	line := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	want := "mul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got, err := removePrefix(line)

	if err != nil {
		t.Fatalf("In removePrefix, received error: %s", err.Error())
	}

	if got != want {
		t.Fatalf("In removePrefix, wanted %s, got %s", want, got)
	}

	line = "4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	want = "mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	got, err = removePrefix(line)

	if err != nil {
		t.Fatalf("In removePrefix, received error: %s", err.Error())
	}

	if got != want {
		t.Fatalf("In removePrefix, wanted %s, got %s", want, got)
	}

}

func TestGetMulInstructionResult(t *testing.T) {
	line := "mul(3,9)"

	want := 27
	got, err := getMulInstructionResult(line)

	if err != nil {
		t.Fatalf("In getMulInstructionResult, received error: %s", err.Error())
	}

	if got != want {
		t.Fatalf("In getMulInstructionResult, wanted %d, got %d", want, got)
	}

	line = "mul[3,7]!@^do_not_mul(5,5)"
	_, err = getMulInstructionResult(line)
	if err == nil {
		t.Fatal("In getMulInstructionResult, should've received error and did not.")
	}

	line = "mul(3,7]!@^do_not_mul(5,5)"
	_, err = getMulInstructionResult(line)
	if err == nil {
		t.Fatal("In getMulInstructionResult, should've received error and did not.")
	}
}
