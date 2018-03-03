package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
	"fmt"
	"net"
	"bufio"
)

// The person Type (more like an object)
type LEDStatus struct {
    STATUS        string   `json:"status,omitempty"`
    DeviceID string   `json:"deviceid,omitempty"`
   
}


var LED []LEDStatus

// Display all from the LED var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(LED)
}

// Display a single data
func GetLEDStatus(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range LED {
        if item.STATUS == params["status"] {
            json.NewEncoder(w).Encode(item)
			//fmt.Println("Need to turn the server" + item.STATUS)
			conn, _ := net.Dial("tcp", "192.168.78.113:10002")
				  	
			if item.STATUS == "ON" {
				fmt.Println("Need to turn the server" + item.STATUS)
				//Raspberry love goes here 
				
					// connect to this socket
					//conn, _ := net.Dial("tcp", "192.168.78.113:10002")
				  	// read in input from stdin
					//reader := bufio.NewReader("1001-ON")
					fmt.Print("Text to send:1001-ON ")
					//text, _ := reader.ReadString('\n')
					// send to socket
					fmt.Fprintf(conn, "1001-ON" )//+ "\n")
					// listen for reply
					fmt.Print("Waiting for reply")
					message, _ := bufio.NewReader(conn).ReadString('\n')
					fmt.Print("Message from server: "+message)
					//conn.Close()
				
			} else if item.STATUS == "OFF" {
				//RAspberry love goes here 
				fmt.Println("Need to turn the server" + item.STATUS)
				
					// connect to this socket
					// read in input from stdin
					//reader := bufio.NewReader("1001-OFF")
					fmt.Print("Text to send:1001-OFF ")
					//text, _ := reader.ReadString('\n')
					// send to socket
					fmt.Fprintf(conn, "1001-OFF" ) // + "\n")
					// listen for reply
					fmt.Print("Waiting for reply")
					message2, _ := bufio.NewReader(conn).ReadString('\n')
					fmt.Print("Message from server: "+message2)
					//conn2.Close()
					
			} else {
				fmt.Println("Crapy request" + item.STATUS)
			}	
			
            return
        }
    }
    json.NewEncoder(w).Encode(&LEDStatus{})
}

// create a new item
func CreateLEDStatus(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person LEDStatus
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.STATUS = params["status"]
    LED = append(LED, person)
    json.NewEncoder(w).Encode(LED)
}

// Delete an item
func DeleteLEDStatus(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range LED {
        if item.STATUS == params["status"] {
            LED = append(LED[:index], LED[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(LED)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    LED = append(LED, LEDStatus{STATUS: "ON", DeviceID: "1001"})
    LED = append(LED, LEDStatus{STATUS: "OFF", DeviceID: "1001"})
    router.HandleFunc("/LED", GetPeople).Methods("GET")
    router.HandleFunc("/LED/{status}", GetLEDStatus).Methods("GET")
    router.HandleFunc("/LED/{status}", CreateLEDStatus).Methods("POST")
    router.HandleFunc("/LED/{status}", DeleteLEDStatus).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}


  
