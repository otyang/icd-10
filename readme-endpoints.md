## Get diagnosis by ID

### Request
GET /icd/**:fullCode**


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `fullCode | `string` | **Required**. The fullcode of the diagnosis |


```
curl --request GET \
  --url http://localhost:3000/icd/A000
```

### Response
```
HTTP/1.1 200 OK
Date: Tue, 02 May 2023 07:36:05 GMT
Content-Type: application/json
Content-Length: 362
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful",
  "data": {
    "categoryCode": "A00",
    "diagnosisCode": "0",
    "fullCode": "A000",
    "abbreviatedDeScription": "Cholera due to Vibrio cholerae 01, biovar cholerae",
    "fullDecription": "Cholera due to Vibrio cholerae 01, biovar cholerae",
    "categoryTitle": "Cholera",
    "createdAt": "2023-05-01T03:36:28Z",
    "updatedAt": "2023-05-01T03:36:28Z"
  }
}
```



## List Diagnosis code

### Request
GET /icd?limit=**limitPerPage**&cursor=**fullcode**


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `limit | `int` | **Optional**. default is 20 |
| `fullCode | `string` | **Optional**. diagnosis fullcode. default is 0 |


```
curl --request GET \
  --url 'http://localhost:3000/icd?limit=20&cursor=0'
```

### Response
```
HTTP/1.1 200 OK
Date: Tue, 02 May 2023 08:03:01 GMT
Content-Type: application/json
Content-Length: 5472
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful",
  "data": {
    "nextPageCursor": "A0223",
    "prevPageCursor": null,
    "records": [
      {
        "categoryCode": "A00",
        "diagnosisCode": "0",
        "fullCode": "A000",
        "abbreviatedDeScription": "Cholera due to Vibrio cholerae 01, biovar cholerae",
        "fullDecription": "Cholera due to Vibrio cholerae 01, biovar cholerae",
        "categoryTitle": "Cholera",
        "createdAt": "2023-05-01T03:36:28Z",
        "updatedAt": "2023-05-01T03:36:28Z"
      },
      .....
      ]
    }
}
```



## Create a Diagnosis Code

### Request
POST /icd


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `categoryCode | `string` | **required**. Must be 3 letters and above code |
| `diagnosisCode | `string` | **optional**. usually numbers max of 1 digit |
| `abbreviatedDescription | `string` | **required**. minimum of 4 letters |
| `fullDescription | `string` | **required**. minimum of 4 letters |
| `categoryTitle | `string` | **required**. minimum of 3 letters |


```
curl --request POST \
  --url http://localhost:3000/icd \
  --header 'content-type: application/json' \
  --data '{"categoryCode": "A00000","diagnosisCode": "1","abbreviatedDescription": "New diseases sars","fullDescription": "sars disease, airbone","categoryTitle": "airborne"}'
```

### Response
```
HTTP/1.1 201 Created
Date: Tue, 02 May 2023 08:23:42 GMT
Content-Type: application/json
Content-Length: 331
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful",
  "data": {
    "categoryCode": "A00000",
    "diagnosisCode": "1",
    "fullCode": "A000001",
    "abbreviatedDeScription": "New diseases sars",
    "fullDecription": "sars disease, airbone",
    "categoryTitle": "airborne",
    "createdAt": "2023-05-02T09:23:42.481114+01:00",
    "updatedAt": "2023-05-02T09:23:42.481115+01:00"
  }
}
```




## Edit an existing diagnosis code record
 
### Request
PUT /icd/**:fullCode**


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `categoryCode | `string` | **required**. Must be 3 letters and above code |
| `diagnosisCode | `string` | **optional**. usually numbers max of 1 digit |
| `abbreviatedDescription | `string` | **required**. minimum of 4 letters |
| `fullDescription | `string` | **required**. minimum of 4 letters |
| `categoryTitle | `string` | **required**. minimum of 3 letters |


```
curl --request PUT \
  --url http://localhost:3000/icd/A000001 \
  --header 'content-type: application/json' \
  --header 'user-agent: vscode-restclient' \
  --data '{"categoryCode": "A00000","diagnosisCode": "1","abbreviatedDescription": "Edited - New diseases sars","fullDescription": "Edited - sars disease, airbone","categoryTitle": "Edited - airborne"}'
```

### Response
```
HTTP/1.1 200 OK
Date: Tue, 02 May 2023 08:27:32 GMT
Content-Type: application/json
Content-Length: 353
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful",
  "data": {
    "categoryCode": "A00000",
    "diagnosisCode": "1",
    "fullCode": "A000001",
    "abbreviatedDeScription": "Edited - New diseases sars",
    "fullDecription": "Edited - sars disease, airbone",
    "categoryTitle": "Edited - airborne",
    "createdAt": "2023-05-02T08:23:42.481114Z",
    "updatedAt": "2023-05-02T09:27:32.615988+01:00"
  }
}
```




## Delete a diagnosis code by ID
 
### Request
DELETE /icd/**:fullCode**


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `fullCode | `string` | **Required**. The fullcode of the diagnosis |


```
curl --request DELETE \
  --url http://localhost:3000/icd/A000001 
```

### Response
```
HTTP/1.1 200 OK
Date: Tue, 02 May 2023 08:29:57 GMT
Content-Type: application/json
Content-Length: 50
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful"
}
```


##  Uploading ICD CSV files with diagnosis code records
 
### Request
POST /icd-upload


| Parameter | Type | Description |
| :--- | :--- | :--- |
| `email | `string` | **Required**. email of the uploader. This is for notification purpose when upload is completed |
| `csv-file | `file` | **Required**. the csv file to upload |


```
curl --request POST \
  --url http://localhost:3000/icd-upload \
  --header 'content-type: multipart/form-data; boundary=----theDelimiter' \
  --data '------theDelimiter
Content-Disposition: form-data; name="email"

email-address-of-uploader@domain.com
------theDelimiter
Content-Disposition: form-data; name="csv-file"; filename="1.csv"
Content-Type: text/csv

A00,0,A000,"Cholera due to Vibrio cholerae 01, biovar cholerae","Cholera due to Vibrio cholerae 01, biovar cholerae","Cholera"
A00,1,A001,"Cholera due to Vibrio cholerae 01, biovar eltor","Cholera due to Vibrio cholerae 01, biovar eltor","Cholera"
------theDelimiter--'
```

### Response
```
HTTP/1.1 200 OK
Date: Tue, 02 May 2023 08:29:57 GMT
Content-Type: application/json
Content-Length: 50
Vary: Origin
Access-Control-Allow-Origin: *
Connection: close

{
  "success": true,
  "message": "Request was succesful"
}
```