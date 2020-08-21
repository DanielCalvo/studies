package main

//type Person struct {
//	Name string
//	Age  int
//}
//
//func personCreate(w http.ResponseWriter, r *http.Request) {
//	var p Person
//	err := json.NewDecoder(r.Body).Decode(&p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	fmt.Fprintf(w, "Person: %+v", p)
//	fmt.Println("Received:",p)
//}
//
//func main(){
//	fmt.Println("hai")
//	http.HandleFunc("/person/create", personCreate)
//	log.Fatal(http.ListenAndServe(":8080", nil)) //Using the default serveMux
//	fmt.Println("done")
//}

//curl --header "Content-Type: application/json" --request POST --data '{"Name":"Bob","Age":10}' http://localhost:8080/person/create