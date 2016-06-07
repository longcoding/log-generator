package main

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
	"strconv"
)

func GetLogLine() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return (createTimeStamp() + "||" +
		createUUID() + "||" +
		createAssetId() + "||" +
		createServiceId() + "||" +
		createOpenApiId() + "||" +
		createApplicationId() + "||" +
		createClientIp() + "||" +
		createLogDataSize() + "||" +
		createResponseCode() + "||" +
		createResponseTime() + "||" +
		createUserAgent() + "||" +
		createApiStnCode())
}

func createTimeStamp() string {
	var currentTime = time.Now()
	var timeStamp string = strconv.Itoa(currentTime.Year())

	standard := int(currentTime.Month())
	timeStamp += AppendCharacter(standard) + AppendCharacter(currentTime.Day());
	timeStamp += strconv.Itoa(currentTime.Hour()) + strconv.Itoa(currentTime.Minute()) +
		strconv.Itoa(currentTime.Second()) + strconv.Itoa(currentTime.Nanosecond() / 1000000)
	return timeStamp
}

func AppendCharacter(standard int) string {
	var returnValue string
	if standard < 10 {
		returnValue = "0"
	}
	returnValue += strconv.Itoa(standard)
	return returnValue
}

func createUUID() string {
	return uuid.NewV4().String()
}

func createAssetId() string {
	var assetList [5]string = [5]string{"10023", "10040", "10030", "10042", "10073"}
	return assetList[rand.Intn(5)]
}

func createServiceId() string {
	var serviceList = make([]string, 5)
	serviceList[0] = "10023"
	serviceList[1] = "10004"
	serviceList[2] = "10030"
	serviceList[3] = "10042"
	serviceList[4] = "10073"
	return serviceList[rand.Intn(5)]
}

func createOpenApiId() string {
	apiId := make(map[int]string)
	apiId[0] = "520004"
	apiId[1] = "520010"
	apiId[2] = "530020"
	apiId[3] = "530060"
	apiId[4] = "530070"
	apiId[5] = "530075"
	apiId[6] = "540090"

	randomValue := rand.Intn(len(apiId))
	for key, value := range apiId {
		if key == randomValue {
			return value
		}
	}
	return ""
}

func createApplicationId() string {
	applicationId := map[int]string{
		0 : "400002323",
		1 : "400001540",
		2 : "400023201",
		3 : "400102021",
		4 : "400102030",
		5 : "400001020",
		6 : "400102032",
		7 : "400203021",
		8 : "400502021",
		9 : "400602012",
	}

	randomValue := rand.Intn(len(applicationId))
	for key, value := range applicationId {
		if key == randomValue {
			return value
		}
	}
	return ""
}

func createClientIp() string {
	remoteIp := []string{
		"127.0.0.1",
		"255.255.255.0",
		"130.255.111.2",
		"162.211.1.13",
		"192.188.2.123",
		"211.155.12.233",
		"11.120.22.21",
		"150,23.11.0",
		"34.121.222.111",
	}
	return remoteIp[rand.Intn(len(remoteIp))]
}

func createLogDataSize() string {
	return strconv.Itoa(rand.Intn(100000))
}

func createResponseCode() string {
	responseCode := []string{
		"200-",
		"401-3102",
		"400-9401",
		"403-9404",
		"502-9504",
	}
	return responseCode[rand.Intn(len(responseCode))]
}

func createResponseTime() string {
	return strconv.Itoa(rand.Intn(1000))
}

func createUserAgent() string {
	userAgent := map[int]string{
		0: "Nescape",
		1: "Android",
		2: "iOS",
		3: "fireFox",
	}
	return userAgent[rand.Intn(len(userAgent))]
}

func createApiStnCode() string {
	apiStnCode := map[int]string{
		0: "0",
		1: "1",
	}
	randomValue := rand.Intn(1)
	for key, value := range apiStnCode {
		if key == randomValue {
			return value
		}
	}
	return ""
}
