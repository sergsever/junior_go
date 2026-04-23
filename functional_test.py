import requests
import json

def main():
	base_url = "http://localhost:8080/api"
	print("start tests\n")
	print("test creation\n")

	data = {"ServiceName":'Yandex Plus',"Price":400,"UserId":'60601fee-2bf1-4721-ae6f-7636e79a0cba',"StartDate":'2025-01-07T23:20:50.52Z'}

	headers = {
	'Content-type': 'application/json'
	}

	jdata = json.dumps(data)

	url = base_url + "/create"
	response = requests.post(url, headers=headers, data=jdata).text
	print("response: ", response)
	print("test read")
	url = base_url + "/read/1"
	response = requests.get(url)
	print("response: ", response.text)
	print("test update")
	data = {"SubscriptionId":1, "ServiceName":'Yandex Plus',"Price":450,"UserId":'60601fee-2bf1-4721-ae6f-7636e79a0cba',"StartDate":'2025-01-07T23:20:50.52Z'}
	url = base_url + "/edit"
	jdata = json.dumps(data)
	response = requests.post(url, headers=headers, data=jdata).text
	print("response: ", response)
	print("test delete")
	url = base_url + "/delete/1"
	response = requests.get(url)
	print("response:", response)
	print("test list")
	url = base_url + "/list"
	response = requests.get(url)
	print("response: ", response.text)

	print("test sums")
	data_user = {"FromDate":"2025-01-07T23:20:50.52Z", "ToDate": "2028-01-07T23:20:50.52Z", "UserId": "60601fee-2bf1-4721-ae6f-7636e79a0cba" }
	data_service = {"FromDate":"2025-01-07T23:20:50.52Z", "ToDate": "2028-01-07T23:20:50.52Z", "ServiceName": "test" }

	url = base_url + "/sum"

	jdata = json.dumps(data_user)
	response = requests.post(url, headers=headers, data=jdata)
	print("response by user: ", response.text)
	jdata = json.dumps(data_service)
	response = requests.post(url, headers=headers, data=jdata)
	print("response by service: ", response.text)

 


if(__name__ == "__main__"):
	main()
