package main

import (
	sap_api_caller "sap-api-integrations-identity-reads-rmq-kube/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-identity-reads-rmq-kube/SAP_API_Input_Reader"
	"sap-api-integrations-identity-reads-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.QueueTo())
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()

	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		conf.RMQ.QueueTo(),
		rmq,
		l,
	)

	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		err = callProcess(caller, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func callProcess(caller *sap_api_caller.SAPAPICaller, msg rabbitmq.RabbitmqMessage) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = xerrors.Errorf("error occurred: %w", e)
			return
		}
	}()
	userID := extractData(msg.Data())
	accepter := getAccepter(msg.Data())
	caller.AsyncGetIdentityGetLoggedInUserInfo(userID, accepter)
	return nil
}

func extractData(data map[string]interface{}) (userID string) {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	userID = sdc.IdentityGetLoggedInUserInfo.UserID
	return
}

func getAccepter(data map[string]interface{}) []string {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"IdentityGetLoggedInUserInfo",
		}
	}
	return accepter
}

