
### Distributed Image Store

1. Image are made base64 encoded and stored as key-values pairs
2. There are 3 database servers acting as shards and each shard stores key-values in JSON format


### For Backend
1. The main server will be 127:0.0.1
2. db servers will be 127:0.0.2, 127:0.0.3, 127:0.0.4 (get/set for each)
3. Main server routes requests to db servers and returns the response
4. The method for dividing data shards is -> hash(key) % 3, this will equally divide data into 3 shards

### For Frontend
1. Encoding is being processed in frontend and then sent to backend
2. A unique key is associated with the image while storing so that it can be retrieved later

# FLow

1. If store is not created while setting, we create it. The data from file is first read into a store variable



### Questions ?
1. What distribution means - we store the data in multiple servers. Why? cause one server
   might not be able to store all data