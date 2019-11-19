package main

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "fmt"
    "github.com/hyperledger/fabric/protos/peer"
    "encoding/json"
    "bytes"
)

type CouchDBChaincode struct { }

func (t *CouchDBChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func (t *CouchDBChaincode) Invoke (stub shim.ChaincodeStubInterface) peer.Response {
    fun, args := stub.GetFunctionAndParameters()
    if fun == "billInit" {
        return billInit(stub, args)
    } else if fun == "queryBills" {
        return queryBills(stub, args)
    } else if fun == "queryWaitBills" {
        return queryWaitBills(stub, args)
    }
    return shim.Error("非法操作，指定的函数名无效")
}

func billInit (stub shim.ChaincodeStubInterface, args []string) peer.Response {
    bill := BillStruct {
        ObjectType: "billObj",
        BillInfoID: "POC101",
        BillInfoAmt: "1000",
        BillInfoType: "111",
        BillIsseDate: "20100101",
        BillDueDate: "20100110",
        HolderAcct: "AAA",
        HolderCmID: "AAAID",
        WaitEndorseAcct: "",
        WaitEndorseCmID: "",
    }

    billByte, _ := json.Marshal(bill)
    err := stub.PutState(bill.BillInfoID, billByte)
    if err != nil {
        return shim.Error("初始化第1个票据失败: " + err.Error())
    }

    bill2 := BillStruct {
        ObjectType: "billObj",
        BillInfoID: "POC102222",
        BillInfoAmt: "1000",
        BillInfoType: "111",
        BillIsseDate: "20100101",
        BillDueDate: "20100110",
        HolderAcct: "BBB",
        HolderCmID: "BBBID",
        WaitEndorseAcct: "",
        WaitEndorseCmID: "",
    }

    billByte2, _ := json.Marshal(bill2)
    err = stub.PutState(bill2.BillInfoID, billByte2)
    if err != nil {
        return shim.Error("初始化第2个票据失败: " + err.Error())
    }

    bill3 := BillStruct {
        ObjectType: "billObj",
        BillInfoID: "POC104444",
        BillInfoAmt: "1000",
        BillInfoType: "111",
        BillIsseDate: "20100101",
        BillDueDate: "20100110",
        HolderAcct: "CCC",
        HolderCmID: "CCCID",
        WaitEndorseAcct: "",
        WaitEndorseCmID: "",
    }

    billByte3, _ := json.Marshal(bill3)
    err = stub.PutState(bill3.BillInfoID, billByte3)
    if err != nil {
        return shim.Error("初始化第3个票据失败: " + err.Error())
    }

    bill4 := BillStruct {
        ObjectType: "billObj",
        BillInfoID: "POC108888",
        BillInfoAmt: "1000",
        BillInfoType: "111",
        BillIsseDate: "20100101",
        BillDueDate: "20100110",
        HolderAcct: "DDD",
        HolderCmID: "DDDID",
        WaitEndorseAcct: "",
        WaitEndorseCmID: "",
    }

    billByte4, _ := json.Marshal(bill4)
    err = stub.PutState(bill4.BillInfoID, billByte4)
    if err != nil {
        return shim.Error("初始化第4个票据失败: " + err.Error())
    }

    return shim.Success([]byte("初始化票据成功"))
}

func queryBills (stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("必须且只能指定持票人的证件号码")
    }

    holderCmID := args[0]

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"billObj\",\"HolderCmID\":\"%s\"}}", holderCmID)

    result, err := getBillsByQueryString(stub, queryString)
    if err != nil {
        return shim.Error("根据持票人的证件号码批量查询持票人的持有票据列表时发生错误: " + err.Error())
    }

    return shim.Success(result)
}

func queryWaitBills(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("必须且只能指定待背书人的证件号码")
    }

    waitEndorseCmID := args[0]

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"billObj\",\"WaitEndorseCmID\":\"%s\"}}", waitEndorseCmID)

    result, err := getBillsByQueryString(stub, queryString)
    if err != nil {
        return shim.Error("根据待背书人的证件号码批量查询待背书的票据列表时发生错误: " + err.Error())
    }

    return shim.Success(result)
}

func getBillsByQueryString (stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
    iterator, err := stub.GetQueryResult(queryString)
    if err != nil {
        return nil, err
    }

    defer iterator.Close()

    var buffer bytes.Buffer
    var isSplit bool
    for iterator.HasNext() {
        result, err := iterator.Next()
        if err != nil {
            return nil, err
        }
        if isSplit {
            buffer.WriteString("; ")
        }
        buffer.WriteString("Key:")
        buffer.WriteString(result.Key)
        buffer.WriteString(", Value: ")
        buffer.WriteString(string(result.Value))
        isSplit = true
    }
    return buffer.Bytes(), nil
}

func main() {
    err := shim.Start(new(CouchDBChaincode))
    if err != nil {
        fmt.Errorf("启动链码失败: %v", err)
    }
}




