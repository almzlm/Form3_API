package accounts

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"clientapp/models"
)


/**
	Function    : sendDeleteRequest
	Description : Private helper function that handles sending the http request
	Params:     : string requestUrl --> The url pointing to the resource to delete
	Returns     : int indicating success (0) or failure (-1)
**/ 
func sendDeleteRequest(requestUrl string) (int){

	//Make DELETE request
	resp, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil { 
			fmt.Printf("accounts::DeleteAccount > Error creating DELETE request : %s \n", err)
			return ERROR
	}

	httpResp, httpErr := http.DefaultClient.Do(resp)
	if httpErr != nil{
		fmt.Printf("accounts::DeleteAccount > Error sending DELETE request : %s \n", httpErr)
		return ERROR
	}
	defer httpResp.Body.Close()

	//Validate HTTP response
	if (httpResp.StatusCode >= 300){
		fmt.Printf("accounts::DeleteAccount > Server returned invalid DELETE response. Returned status : %d \n", httpResp.StatusCode)
		return ERROR
	}

	return SUCCESS
}

/**
	Function    : sendFetchRequest
	Description : Private helper function that handles sending the http request
	Params:     : string requestUrl --> The URL to the resource to get
	Returns     : int indicating success (0) or failure (-1)
				  AccountData containing the retrieved account (empty AccountData is returned if no success)
**/ 
func sendFetchRequest(requestUrl string) (data_models.AccountApiPayload, int){

	resultAccount := data_models.AccountApiPayload {}
	
	//Make GET request
	resp, err := http.Get(requestUrl)
	if err != nil {
			fmt.Printf("accounts::FetchAccount > Error when sending GET request : %s \n", err)
			return resultAccount, ERROR
	}

	defer resp.Body.Close()

	//Validate HTTP response
	if (resp.StatusCode >= 300){
		fmt.Printf("accounts::CreateAccount > Server returned invalid POST request : %d \n", resp.StatusCode)
		return resultAccount, ERROR
	}

	getRespBody,_ := ioutil.ReadAll(resp.Body)

	//Unmarshal response into AccountApiPayload struct
	if err = json.Unmarshal(getRespBody, &resultAccount); err != nil {
		fmt.Printf("accounts::FetchAccount > Error Unmarshalling the JSON response : %s \n", err)
		return resultAccount, ERROR
	}

	return resultAccount, SUCCESS
}

/**
	Function    : sendCreateRequest
	Description : Private helper function that handles sending the http request
	Params:     : string requestUrl      --> The URL to the endpoint
				  AccountApiPayload data --> The account object to create
	Returns     : int indicating success (0) or failure (-1)
**/ 
func sendCreateRequest(data data_models.AccountApiPayload, requestUrl string) (int){

	j, err := json.Marshal(data)
	if err != nil{
			fmt.Printf("accounts::CreateAccount > Error while marshalling data : %s \n ", err)
			return ERROR
	}

	resBody := bytes.NewBuffer(j)

	//Send POST request
	resp, respErr := http.Post(requestUrl,"application/json", resBody)
	if (respErr != nil) {
			fmt.Printf("accounts::CreateAccount > Error sending POST request. Server returned error : %s. \n", respErr)
			return ERROR
	}
	defer resp.Body.Close()

	//Validate HTTP response code
	if (resp.StatusCode >= 300){
		fmt.Printf("accounts::CreateAccount > Server returned invalid POST request : %d \n ", resp.StatusCode)
		return ERROR
	}

	return SUCCESS
}