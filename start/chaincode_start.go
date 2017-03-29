package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode tipo de objecto
type SimpleChaincode struct {
}
// declacaracion de parametros de entrada y salida
//entrada:stub     tipo shim.ChaincodeStubInterface
//        function tipo string
//        args     tipo []string   
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	var A, B string    // declacion de entidades 
	var Aval, Bval int // valor de cada entidad
	var err error      // valores de error
    // si la variable args tiene mas o menos de 4 argumentos , se define un error
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")//creacion del error tipo string
	}

	// asignacion de variables 
	A = args[0] // nombre de la entidad
	Aval, err = strconv.Atoi(args[1])	//Validacion del valor de la entidad A y asignacion del valor
  
	if err != nil {// si hubo error en la validacion del numero de entidad A, publica q tipo de error hubo
		return nil, errors.New("Expecting integer value for asset holding")
	}
	B = args[2]// asignacion de variables 
	Bval, err = strconv.Atoi(args[3])// nombre de la entidad
	
	if err != nil {// si hubo error en la validacion del numero de entidad A, publica q tipo de error hubo
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)// muestra la el valor de las entidades

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))//registra el las actividades con un id unico
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))//registra el las actividades con un id unico
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Transaction makes payment of X units from A to B//creacion de transaccionesiones  de A - B
//recibe 2 variables 
// Stub
// args
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	var A, B string    // declacion de entidades 
	var Aval, Bval int // valor de cada entidad
	var X int          // valor de transaccion
	var err error      // valores de error
 // si la variable args tiene mas o menos de 4 argumentos , se define un error
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")//creacion del error tipo string
	}
//asignacion de las variables de las variables de entrada al nombre de las entidades
	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)//devuelve la cadena de bytes asociados a la llave(A)
	
	if err != nil {
		return nil, errors.New("Failed to get state")//crea mensaje de error de la obtencion de la llave
	}
	
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")// no puede encontrar la entidad buscada
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))//extrae el valor de la variable Avalbytes

	Bvalbytes, err := stub.GetState(B)//devuelve la cadena de bytes asociados a la llave(B)
	if err != nil {
		return nil, errors.New("Failed to get state")//crea mensaje de error de la obtencion de la llave
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")// no puede encontrar la entidad buscada
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))//extrae el valor de la variable Avalbytes

	// Perform the execution
	X, err = strconv.Atoi(args[2])//extraer el valor de transaccion
	Aval = Aval - X // se realiza la sustraccion de los elementos de A
	Bval = Bval + X // se realiza la adiccion de los elementos de B
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)// se muestra los nuevos valores de la trasaccion 

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))//se guarda los nuevos valores de las entidades para A
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))//se guarda los nuevos valores de las entidades para B
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running delete")
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}



func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}



// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")
	
	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}