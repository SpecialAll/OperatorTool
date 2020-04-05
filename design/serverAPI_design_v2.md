# serverAPI_Design

> The serverAPI design adopts the RESTful interface design pattern and designs the corresponding API interface according to the specific requirements. 

**From the perspective of requirements, related apis can be divided into the following points:**

## I. nameserver related:


*  **Query nameserver information:**
 `GET   /nameservers`  or 
 `GET    /nameserver/{NodeId}`

 
```json
-- 200 response
{
  "state": "success",
  "message": "success",
  "items": [
    {
      "nodeName": "node0",
      "nameServers": {
        "nameserver": [
          "10.4.192.27",
          "10.4.192.27"
        ],
        "search": [
          "domain.sensetime.com"
        ]
      }
    },
    {
      "nodeName": "node1",
      "nameServers": {
        "nameserver": [
          "10.4.192.27",
          "10.4.192.27"
        ],
        "search ": [
          "domain.sensetime.com"
        ]
      }
    }
  ]
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```

*  **Change the nameserver information:**`PUT /nameservers` or `PUT  /nameserver/{NodeId}`
```json
-- 200 response
{
  "state": "success",
  "message": "nameserver success added"
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```
 

 
## Ii. Images related:


*  **Query images information:** `GET  /images`  or `GET  /images/{NodeId}`
```json
-- 200 response
{
  "state": "success",
  "message": "success",
  "items": [
    {
      "repository": "nginx",
      "tag": "latest",
      "imageId": "c7460dfcab50",
      "created": "3 days ago",
      "size": "126MB"
    },
    {
      "repository": "nginx",
      "tag": "latest",
      "imageId": "c7460dfcab50",
      "created": "3 days ago",
      "size": "126MB"
    }
  ]
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```

*  **Change images information:**`PUT /images` or `PUT  /images/{NodeId}`
```json
-- 200 response
{
  "state": "success",
  "message": "image success added"
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```
*  **Delete images information:** `DELETE  /images/node?{image}`
```json
-- 200 response
{
  "state": "success",
  "messages": "delete success"
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```

## Iii. CPU related:


*  **Query CPU usage:**`GET /cpu` or `GET   /cpu/{NodeId}`
``` json
-- 200 response
{
  "state": "success",
  "message": "success",
  "items": [
    {
      "containerId": "9e79cc5cb9ae ",
      "containerName": "k8s_coredns_coredns-6955765f44-88hfb_kube-system_242e349b-de10-4342-8bb4-0a7c698e1c60_1",
      "cpu%": "0.12%",
      "MEM USAGE / LIMIT": "9.258MiB / 170MiB",
      "MEM%": "0.23%",
      "NET I/O": "2.05MB / 462kB",
      "BLOCK I/O": "262kB / 0B",
      "PIDS": "20"
    },
    {
      "containerId": "9e79cc5cb9ae ",
      "containerName": "k8s_coredns_coredns-6955765f44-88hfb_kube-system_242e349b-de10-4342-8bb4-0a7c698e1c60_1",
      "cpu%": "0.12%",
      "MEM USAGE / LIMIT": "9.258MiB / 170MiB",
      "MEM%": "0.23%",
      "NET I/O": "2.05MB / 462kB",
      "BLOCK I/O": "262kB / 0B",
      "PIDS": "20"
    }
  ]
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```
## Iiii. DNS related:

*  **Query DNS information:**`GET  /dns` or `GET   /dns/{NodeId}`
``` json
-- 200 response
{
    "state":"success",
    "message":"success",
    "items":[
        {
            "server":"10.4.192.27",
            "address":"10.4.192.27#53",
            "Non-authoritative answer":"www.baidu.com	canonical name = www.a.shifen.com.",
            "name":"www.a.shifen.com",
            "trueAddress":"61.135.169.125"
        },
        {
            "server":"10.4.192.27",
            "address":"10.4.192.27#53",
            "Non-authoritative answer":"www.baidu.com	canonical name = www.a.shifen.com.",
            "name":"www.a.shifen.com",
            "trueAddress":"61.135.169.125"
        }
        ]
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```
*  **Change DNS information:**`PUT /dns` or `PUT  /dns/{NodeId}`
``` json
-- 200 response
{
  "state": "success",
  "message": "dns success added"
}

-- 401 Unauthorized
{
  "state": "failed",
  "message": "Unauthorized Operation"
}
```
