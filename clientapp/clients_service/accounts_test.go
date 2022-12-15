package accounts

import (
        "clientapp/models"
        "testing"
)

/////////////////////////////////////////////////////
//////////////   TEST ACCOUNT CREATION //////////////
/////////////////////////////////////////////////////
func TestAccountCreation_ValidRequest(t *testing.T) {
	client := NewClient()
	
	//Build the payload to be sent
	countryCode := "GB"
	version := new(int64)
	*version = 0

	accAttributes := data_models.AccountAttributes{
			BankID : "400300",
			Country: &countryCode,
			Bic: "NWBKGB22",
			Name: []string{"TEST"},
	}

	accData := data_models.AccountData{
			ID : "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type: "accounts",
			Attributes: &accAttributes,
			Version : version,
	}

	data := data_models.AccountApiPayload{
			Data : &accData,
	}

	//Delete the account that was created
	client.DeleteAccount(accData.ID, accData.Version)

	//Create the account
	res := client.CreateAccount(data)
	if (res != SUCCESS){
			t.Fatalf("Output from CreateAccount does not match expected value. Current : %d -- Expected : %d \n", res, SUCCESS)
	}

	//Delete the account that was created
	client.DeleteAccount(accData.ID, accData.Version)
}

/////////////////////////////////////////////////////
//////////////   TEST ACCOUNT CREATION //////////////
/////////////////////////////////////////////////////
func TestAccountCreation_EmptyPayload(t *testing.T) {
	client := NewClient()
	data := data_models.AccountApiPayload{}

	res := client.CreateAccount(data)
	if (res != ERROR){
			t.Fatalf("Output from CreateAccount does not match expected value. Current : %d -- Expected : %d \n", res, ERROR)
	}
}

/////////////////////////////////////////////////////
//////////////   TEST ACCOUNT CREATION //////////////
/////////////////////////////////////////////////////
func TestAccountCreation_InvalidPayload(t *testing.T) {
	client := NewClient()

	//Build the payload to be sent
	countryCode := "12"
	accAttributes := data_models.AccountAttributes{
			BankID : "ABCDEF",
			Country: &countryCode,
			Bic: "ALIEN",
			Name: []string{"12345"},
	}

	accData := data_models.AccountData{
			ID : "000",
			OrganisationID: "0000",
			Type: "INVALID",
			Attributes: &accAttributes,
	}

	data := data_models.AccountApiPayload{
			Data : &accData,
	}

	res := client.CreateAccount(data)
	if (res != ERROR){
			t.Fatalf("Output from CreateAccount does not match expected value. Current : %d -- Expected : %d \n", res, ERROR)
	}
}

/////////////////////////////////////////////////////
//////////////  TEST ACCOUNT FETCHING ///////////////
/////////////////////////////////////////////////////
func TestFetchingAccount_ValidRequest(t *testing.T){
	client := NewClient()
	
	//Build the payload to be sent
	countryCode := "UK"
	version := new(int64)
	*version = 0

	accAttributes := data_models.AccountAttributes{
			BankID : "400300",
			Country: &countryCode,
			Bic: "NWBKGB22",
			Name: []string{"TEST"},
	}

	accData := data_models.AccountData{
			ID : "be57e265-4156-abcd-a0e7-0007ea9cc8dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type: "accounts",
			Attributes: &accAttributes,
			Version : version,
	}

	data := data_models.AccountApiPayload{
			Data : &accData,
	}

	//Create the account
	res := client.CreateAccount(data)
	if (res != SUCCESS){
			t.Fatalf("Output from CreateAccount does not match expected value. Current : %d -- Expected : %d \n", res, SUCCESS)
	}

	fetchedAcc, fetchedRes := client.FetchAccount(accData.ID)
	if (fetchedRes != SUCCESS){
		t.Fatalf("Output from FetchAccount does not match expected value. Current : %d -- Expected : %d \n", fetchedRes, SUCCESS)
	}

	if (fetchedAcc.Data.ID != accData.ID){
		t.Fatalf("Output from FetchAccount does not match expected value. Current : %s -- Expected : %s \n", fetchedAcc.Data.ID, accData.ID)
	}

	//Delete the account that was created
	client.DeleteAccount(accData.ID, accData.Version)
}

/////////////////////////////////////////////////////
//////////////  TEST ACCOUNT FETCHING ///////////////
/////////////////////////////////////////////////////
func TestFetchingAccount_InvalidRequest(t *testing.T){
	client := NewClient()

	fetchedAcc, fetchedRes := client.FetchAccount("some-invalid-account-id")
	if (fetchedRes != ERROR){
		t.Fatalf("Output from FetchAccount does not match expected value. Current : %d -- Expected : %d \n", fetchedRes, ERROR)
	}

	if ((fetchedAcc.Data != nil)){
		t.Fatalf("Output from FetchAccount does not match expected value. Current : %p -- Expected : %v \n", fetchedAcc.Data, nil)
	}
}

/////////////////////////////////////////////////////
//////////////  TEST ACCOUNT DELETION ///////////////
/////////////////////////////////////////////////////
func TestDeletingAccount_ValidRequest(t *testing.T){
	client := NewClient()
	
	//Build the payload to be sent
	countryCode := "UK"
	version := new(int64)
	*version = 0

	accAttributes := data_models.AccountAttributes{
			BankID : "400300",
			Country: &countryCode,
			Bic: "NWBKGB22",
			Name: []string{"TEST"},
	}

	accData := data_models.AccountData{
			ID : "be57e265-4156-abcd-a0e7-0007ea9cc8dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type: "accounts",
			Attributes: &accAttributes,
			Version : version,
	}

	data := data_models.AccountApiPayload{
			Data : &accData,
	}

	//Create the account
	client.CreateAccount(data)

	res := client.DeleteAccount(accData.ID, accData.Version)
	if (res != SUCCESS){
		t.Fatalf("Output from DeleteAccount does not match expected value. Current : %d -- Expected : %d \n", res, SUCCESS)
	}
}

/////////////////////////////////////////////////////
//////////////  TEST ACCOUNT DELETION ///////////////
/////////////////////////////////////////////////////
func TestDeletingAccount_InvalidRequest(t *testing.T){
	client := NewClient()
	dummyId := "a1b2c3-q4w5e6r7-z0x9c8v7"
	dummyVersion := new(int64); *dummyVersion = 0 

	//We test with an unexisting account ID
	res := client.DeleteAccount(dummyId, dummyVersion)
	if (res != ERROR){
		t.Fatalf("Output from DeleteAccount does not match expected value. Current : %d -- Expected : %d \n", res, ERROR)
	}

	//We test with an invalid version
	*dummyVersion = -999 
	res = client.DeleteAccount(dummyId, dummyVersion)
	if (res != ERROR){
		t.Fatalf("Output from DeleteAccount does not match expected value. Current : %d -- Expected : %d \n", res, ERROR)
	}
}