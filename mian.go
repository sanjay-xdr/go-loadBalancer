package main

import (
	"fmt"
	"log"
	"net/http"
)

var servers =[]string{"http://localhost:8080","http://localhost:8081"}

var round_robin int=-1;

func getServer() string{


	round_robin = (round_robin + 1);
	current_server:=(round_robin)%len(servers);


	return servers[current_server];


}

func forwardRequest(w http.ResponseWriter, r *http.Request){


	// fmt.Println(r,"request")

	fmt.Println("Forwarding the Request")

	serverUrl:=getServer();

	fmt.Println(serverUrl,"Server URL");


	req, err := http.NewRequest(r.Method,serverUrl, nil) //creating a new request
	if err!=nil{
		fmt.Print("Something went wrong");
	}

	
	req.Header = r.Header
	req.Host = r.Host
	req.RemoteAddr = r.RemoteAddr
	client := &http.Client{}
	resp, err := client.Do(req)

	if err!=nil{
		fmt.Print("Something went wrong with this")
	}

	defer resp.Body.Close();
	fmt.Printf("Response from server: %s %s\n\n", resp.Proto, resp.Status)

}

func main(){



http.HandleFunc("/",forwardRequest)

http.HandleFunc("/favicon.ico", doNothing) //handling browser favicon call

err:=http.ListenAndServe(":8000",nil);

if err!=nil{
	log.Fatal(err);
}
}

func doNothing(w http.ResponseWriter, r *http.Request){}