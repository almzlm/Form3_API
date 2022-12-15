package accounts

import (
	"fmt"
	"clientapp/models"
)

const ACCOUNTSAPI_URL = "http://accountapi:8080/v1/organisation/accounts"
const MAX_RETRY_COUNT = 3

/** AccountsClient definition **/
type AccountsClient struct {
	Url         string
}

/** Return Codes **/
const (
	SUCCESS int =  0
	ERROR   int = -1
)

/** AccountsClient ctor **/
func NewClient() AccountsClient {
	return AccountsClient { Url : ACCOUNTSAPI_URL}
}


/**
	Function    : CreateAccount
	Description : Creates a new bank account in the Form3 accounts service
	Params:     : AccountCreationPayload data --> The payload to the GET request
	Returns     : int indicating success (0), internal failure (-1)
**/ 
func (client *AccountsClient) CreateAccount(data data_models.AccountApiPayload) (int){

	//Send the CREATE request
	res := sendCreateRequest(data, client.Url)

	//If the request failed, we retry
	if res != SUCCESS{
		for i := 1; i <= MAX_RETRY_COUNT; i++{
			fmt.Printf("Retrying (%d / %d),",i, MAX_RETRY_COUNT)
			res := sendCreateRequest(data, client.Url)
			if (res == SUCCESS){
				break
			}
		}

		//If still failed, return error
		return ERROR
	}

	return SUCCESS
}

/**
	Function    : FetchAccount
	Description : Retrieves an existing account in the Form3 accounts service
	Params:     : string data --> The accountId (ex : "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	Returns     : int indicating success (0) or failure (-1)
				  AccountData containing the retrieved account (empty AccountData is returned if no success)
**/ 
func (client *AccountsClient) FetchAccount(accountId string) (data_models.AccountApiPayload, int){

	if (len(accountId) == 0){
		fmt.Printf("accounts::FetchAccount > Invalid input accountId string \n")
		return data_models.AccountApiPayload {}, ERROR
	}
	
	requestUrl := client.Url +  "/" + accountId
	
	//Send the FETCH request
	acc, err := sendFetchRequest(requestUrl)

	//If the request failed, we retry
	if err != SUCCESS{
		for i := 1; i <= MAX_RETRY_COUNT; i++{
			fmt.Printf("Retrying (%d / %d),",i, MAX_RETRY_COUNT)
			acc, err = sendFetchRequest(requestUrl)
			if (err == SUCCESS){
				break
			}
		}

		//If still failed, return error
		return acc, ERROR
	}

	//Else success
	return acc, SUCCESS
}

/**
	Function    : DeleteAccount
	Description : Deletes an existing account in the Form3 accounts service
	Params:     : string accountId --> The accountId (ex : "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
				  int64  version   --> The version number  
	Returns     : int indicating success (0) or failure (-1)
**/ 
func (client *AccountsClient) DeleteAccount(accountId string, version *int64) (int){

	//Build the URL for the DELETE request since GO does not implement http.DELETE(...)
	requestUrl := fmt.Sprintf("%s%s%d", client.Url + "/", accountId + "?version=", *version )

	//Send the DELETE request
	res := sendDeleteRequest(requestUrl)

	//If the request failed, we retry
	if res != SUCCESS{
		for i := 1; i <= MAX_RETRY_COUNT; i++{
			fmt.Printf("Retrying (%d / %d),",i, MAX_RETRY_COUNT)
			res := sendDeleteRequest(requestUrl)
			if (res == SUCCESS){
				break
			}
		}

		//If still failed, return error
		return ERROR
	}

	return SUCCESS
}
