package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-identity-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToIdentityGetLoggedInUserInfo(raw []byte, l *logger.Logger) ([]IdentityGetLoggedInUserInfo, error) {
	pm := &responses.IdentityGetLoggedInUserInfo{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to LoggedInUserInfoCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	identityGetLoggedInUserInfo := make([]IdentityGetLoggedInUserInfo, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		identityGetLoggedInUserInfo = append(identityGetLoggedInUserInfo, IdentityGetLoggedInUserInfo{
			ObjectID:                  data.ObjectID,
			UserID:                    data.UserID,
			UserName:                  data.UserName,
			UserAccountID:             data.UserAccountID,
			UUID:                      data.UUID,
			EmployeeID:                data.EmployeeID,
			BusinessPartnerID:         data.BusinessPartnerID,
			EmployeeUUID:              data.EmployeeUUID,
			Email:                     data.Email,
			DateFormatCode:            data.DateFormatCode,
			DateFormatCodeText:        data.DateFormatCodeText,
			DecimalFormatCode:         data.DecimalFormatCode,
			DecimalFormatCodeText:     data.DecimalFormatCodeText,
			LogonLanguageCode:         data.LogonLanguageCode,
			LogonLanguageCodeText:     data.LogonLanguageCodeText,
			TimeFormatCode:            data.TimeFormatCode,
			TimeFormatCodeText:        data.TimeFormatCodeText,
			TimeZoneCode:              data.TimeZoneCode,
			TimeZoneCodeText:          data.TimeZoneCodeText,
			TechnicalUserIndicator:    data.TechnicalUserIndicator,
			KeyUserIndicator:          data.KeyUserIndicator,
			InactiveIndicator:         data.InactiveIndicator,
			PasswordInactiveIndicator: data.PasswordInactiveIndicator,
			PasswordLockedIndicator:   data.PasswordLockedIndicator,
			PasswordPolicyCode:        data.PasswordPolicyCode,
			PasswordPolicyCodeText:    data.PasswordPolicyCodeText,
			UserAccountTypeCode:       data.UserAccountTypeCode,
			UserAccountTypeCodeText:   data.UserAccountTypeCodeText,
			StatusCode:                data.StatusCode,
			StatusCodeText:            data.StatusCodeText,
			CreatedOn:                 data.CreatedOn,
			CreatedBy:                 data.CreatedBy,
			ChangedOn:                 data.ChangedOn,
			ChangedBy:                 data.ChangedBy,
			CreatedByUUID:             data.CreatedByUUID,
			ChangedByUUID:             data.ChangedByUUID,
			EntityLastChangedOn:       data.EntityLastChangedOn,
		})
	}

	return identityGetLoggedInUserInfo, nil
}
