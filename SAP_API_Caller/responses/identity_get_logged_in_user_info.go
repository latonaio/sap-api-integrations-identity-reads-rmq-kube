package responses

type IdentityGetLoggedInUserInfo struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ObjectID                  string `json:"ObjectID"`
			UserID                    string `json:"UserID"`
			UserName                  string `json:"UserName"`
			UserAccountID             string `json:"UserAccountID"`
			UUID                      string `json:"UUID"`
			EmployeeID                string `json:"EmployeeID"`
			BusinessPartnerID         string `json:"BusinessPartnerID"`
			EmployeeUUID              string `json:"EmployeeUUID"`
			Email                     string `json:"Email"`
			DateFormatCode            string `json:"DateFormatCode"`
			DateFormatCodeText        string `json:"DateFormatCodeText"`
			DecimalFormatCode         string `json:"DecimalFormatCode"`
			DecimalFormatCodeText     string `json:"DecimalFormatCodeText"`
			LogonLanguageCode         string `json:"LogonLanguageCode"`
			LogonLanguageCodeText     string `json:"LogonLanguageCodeText"`
			TimeFormatCode            string `json:"TimeFormatCode"`
			TimeFormatCodeText        string `json:"TimeFormatCodeText"`
			TimeZoneCode              string `json:"TimeZoneCode"`
			TimeZoneCodeText          string `json:"TimeZoneCodeText"`
			TechnicalUserIndicator    bool   `json:"TechnicalUserIndicator"`
			KeyUserIndicator          bool   `json:"KeyUserIndicator"`
			InactiveIndicator         bool   `json:"InactiveIndicator"`
			PasswordInactiveIndicator bool   `json:"PasswordInactiveIndicator"`
			PasswordLockedIndicator   bool   `json:"PasswordLockedIndicator"`
			PasswordPolicyCode        string `json:"PasswordPolicyCode"`
			PasswordPolicyCodeText    string `json:"PasswordPolicyCodeText"`
			UserAccountTypeCode       string `json:"UserAccountTypeCode"`
			UserAccountTypeCodeText   string `json:"UserAccountTypeCodeText"`
			StatusCode                string `json:"StatusCode"`
			StatusCodeText            string `json:"StatusCodeText"`
			CreatedOn                 string `json:"CreatedOn"`
			CreatedBy                 string `json:"CreatedBy"`
			ChangedOn                 string `json:"ChangedOn"`
			ChangedBy                 string `json:"ChangedBy"`
			CreatedByUUID             string `json:"CreatedByUUID"`
			ChangedByUUID             string `json:"ChangedByUUID"`
			EntityLastChangedOn       string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}
