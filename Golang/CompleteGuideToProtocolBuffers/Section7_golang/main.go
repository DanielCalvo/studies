package main

import (
	"fmt"
	complexpb "github.com/DanielCalvo/studies/Golang/CompleteGuideToProtocolBuffers/Section7_golang/complex"
	enumpb "github.com/DanielCalvo/studies/Golang/CompleteGuideToProtocolBuffers/Section7_golang/enum_example"
	simplepb "github.com/DanielCalvo/studies/Golang/CompleteGuideToProtocolBuffers/Section7_golang/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {

	fmt.Println("hi")
	sm := doSimple()
	fmt.Println(sm)
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm3 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm3)
	fmt.Println(sm3)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}
	fmt.Println(em)

}

func fromJSON(s string, pb proto.Message) {
	err := jsonpb.UnmarshalString(s, pb)
	if err != nil {
		log.Fatalln("Can't unmarshal string to pb:", err)
	}
}

func toJSON(pb proto.Message) string { //For debugging, mostly
	marshaler := jsonpb.Marshaler{}
	s, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't unmarshal pb to string:", err)
		return ""
	}
	return s
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes:", err)
	}

	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Can't write to file:", err)
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read from file:", err)
		return err
	}

	err = proto.Unmarshal(bs, pb)
	if err != nil {
		log.Fatalln("Can't unmarshall byte[] into pb:", err)
		return err
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		Name:       "Simplesimple",
		SampleList: []int32{1, 22, 2, 2, 2, 23, 4},
	}

	sm.Id = 111
	fmt.Println(sm.Id)
	fmt.Println(sm.GetId()) //Always use the getter as it's nil safe

	return &sm

}
