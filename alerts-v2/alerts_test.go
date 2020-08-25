package alerts


import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestListActiveFirings(t *testing.T) {
	defer gock.Off()

	mock := gock.New("https://akaa-baseurl-xxxxxxxxxxx-xxxxxxxxxxxxx.luna.akamaiapis.net")
	mock.
		Get("/alerts/v2/alert-firings/active").
		HeaderPresent("Authorization").
		Reply(200).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		BodyString(`
{
    "data": [
        {
            "firingId": "s@136523713",
            "definitionId": "s@175265",
            "name": "DNSSEC No DS record in parent",
            "startTime": "2016-06-02T07:21:18Z",
            "endTime": null,
            "service": "Edge DNS",
            "fieldMap": {
                "alertType": "No DS record in parent zone",
                "zone": "dnssec.akamaidemo.com",
                "emailTo": "dalawren@akamai.com"
            }
        },
        {
            "firingId": "s@140053407",
            "definitionId": "s@233655",
            "name": "New - Origin Object Not Found (HTTP 404 errors)",
            "startTime": "2016-10-28T03:45:34Z",
            "endTime": null,
            "service": "HTTP Content Delivery",
            "fieldMap": {
                "hits": "17",
                "cpCode": "130213",
                "alertConditionPctErrors": "22",
                "alertType": "Origin Object Not Found (HTTP 404 errors)",
                "alertThreshold": "5",
                "emailTo": "agoda-iat@akamai.com",
                "edgeIp": "96.17.180.42",
                "param1": "5",
                "errors": "4",
                "cpCodeDescription": "img.agoda.net"
            }
        },
        {
            "firingId": "s@140341187",
            "definitionId": "s@103155",
            "name": "Harmonie Web - Connection Failure",
            "startTime": "2016-11-07T20:55:06Z",
            "endTime": null,
            "service": "HTTP Content Delivery",
            "fieldMap": {
                "hits": "9",
                "cpCode": "16399",
                "alertConditionPctErrors": "33",
                "alertType": "Origin Connection Failure",
                "alertThreshold": "2",
                "emailTo": "jguest@akamai.com",
                "edgeIp": "104.86.110.115",
                "param1": "2",
                "errors": "3",
                "cpCodeDescription": "Akamai.Internal.3"
            }
        },
        {
            "firingId": "a@234929",
            "definitionId": "a@cp_140564_2672",
            "name": "cp 140564 Below model",
            "startTime": "2016-11-23T02:15:00Z",
            "endTime": null,
            "service": "all",
            "fieldMap": {
                "cpCode": "140564",
                "alertConditionBitsSec": "3750736.0",
                "accountName": "Akamai Technologies - Assets",
                "anomalyType": "Below model",
                "alertThresholdBitsSec": "1.64799595E7",
                "alertThresholdMbitsSec": "16.4799595",
                "expectedThresholdBitsSec": "6.3653717E7",
                "suppressed": false,
                "message": "cp 140564 Below model",
                "alertConditionMbitsSec": "3.750736",
                "expectedThresholdMbitsSec": "63.653717",
                "cpCodeDescription": "EIS - IPA CPCode"
            }
        },
        {
            "firingId": "a@235008",
            "definitionId": "a@cp_140564_2672",
            "name": "cp 140564 Above model",
            "startTime": "2016-11-23T21:00:00Z",
            "endTime": null,
            "service": "all",
            "fieldMap": {
                "cpCode": "140564",
                "alertConditionBitsSec": "1.36333084E8",
                "accountName": "Akamai Technologies - Assets",
                "anomalyType": "Above model",
                "alertThresholdBitsSec": "1.36189178E8",
                "alertThresholdMbitsSec": "136.189178",
                "expectedThresholdBitsSec": "9.58037033E7",
                "suppressed": false,
                "message": "cp 140564 Above model",
                "alertConditionMbitsSec": "136.333084",
                "expectedThresholdMbitsSec": "95.8037033",
                "cpCodeDescription": "EIS - IPA CPCode"
            }
        }
    ]
}
                `)

	Init(config)

	response, err := ListActiveFirings()
	assert.NoError(t, err)
	assert.Equal(t, assert.IsType(t, &FiringsList{}, response), true)
	assert.Equal(t, 5, len(response.Firings))

	firing := response.Firings[0]
	assert.Equal(t, "s@136523713", firing.FiringID)
	
}
