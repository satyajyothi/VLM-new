package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func printCar(t *testing.T, payload []byte) {
	var car CarStruct
	err := json.Unmarshal(payload, &car)
	if err != nil {
		t.Fatalf("MockInvoke error: Issue with Car json unmarshaling")
		return
	}
	fmt.Println("Car ->", car)
}

func testCreateCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testCreateCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("createCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)

}

func testTransferCar(t *testing.T, stub *shim.MockStub, chasisNo string, newOwner string) {
	fmt.Println("Entering testTransferCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("transferCar"), []byte(chasisNo), []byte(newOwner)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testSellnRegisterCar(t *testing.T,
	stub *shim.MockStub,
	chasisNo string,
	newOwner string,
	registrationNo string,
	registrationExp string) {
	fmt.Println("Entering testSellnRegisterCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("sellnRegisterCar"),
			[]byte(chasisNo),
			[]byte(newOwner),
			[]byte(registrationNo),
			[]byte(registrationExp)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)
}

func testScrapCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testScrapCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("scrapCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	printCar(t, valAsbytes)

}

func testGetCar(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testGetCar")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("getCar"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	} else {
		printCar(t, result.Payload)
	}
}

func testGetCarHistory(t *testing.T, stub *shim.MockStub, chasisNo string) {
	fmt.Println("Entering testGetCarHistory")

	// The first parameter is the function we are invoking
	result := stub.MockInvoke("001",
		[][]byte{[]byte("getCarHistory"), []byte(chasisNo)})

	// We expect a shim.ok if all goes well
	if result.Status != shim.OK {
		t.Fatalf("MockInvoke error: %s", result.Message)
	}

	// Assert
	valAsbytes, err := stub.GetState(chasisNo)
	if err != nil {
		t.Errorf("Failed to get state: %s", err.Error())
	} else if valAsbytes == nil {
		t.Errorf("Value does not exist.")
	}
	fmt.Println(valAsbytes)
}

func TestSmartContract(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("vlm", scc)

	testCreateCar(t, stub, "1000")
	testTransferCar(t, stub, "1000", "Rahul")
	testSellnRegisterCar(t, stub, "1000", "Vaibhav", "TS071", "31DEC2025")
	testScrapCar(t, stub, "1000")
	testSellnRegisterCar(t, stub, "1000", "Jyothi", "TS071", "31DEC2025")
	testGetCar(t, stub, "1000")
	// testGetCarHistory(t, stub, "1000")
}

