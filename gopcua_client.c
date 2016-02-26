#include <stdio.h>
#include <signal.h>

#include <open62541.h>
#include <gopcua_client.h>

/**
  * Creates a new UA_client
  *  @return *UA_Client
*/
UA_Client *gopcua_client_new(UA_ClientConfig config) {
  return UA_Client_new(config, Logger_Stdout);
}

/**
  * Connects a UA_Client client to the UA_Server
  *  @param connInfo - the connection string of the OPC UA server 
  *  @return statusCode - 0 if ok.
*/
int gopcua_client_connect(UA_Client *client, char connInfo[]) {
  return UA_Client_connect(client, UA_ClientConnectionTCP, connInfo);
}

/**
  * Prepares an UA_ReadRequest for a single read to an UA_Server
  *  @param *r - the UA_ReadRequest 
  *  @param nodeid - the string node id to be read
*/
void UA_ReadRequest_PrepareSingleRead(UA_ReadRequest *r, char nodeid[]) {
  r->nodesToRead =  UA_Array_new(1, &UA_TYPES[UA_TYPES_READVALUEID]);
  r->nodesToReadSize = 1;
  r->nodesToRead[0].nodeId = UA_NODEID_STRING_ALLOC(1, nodeid);
  r->nodesToRead[0].attributeId = UA_ATTRIBUTEID_VALUE;
}

/**
  * Checks if a UA_ReadResponse is a valid scalar response.
  *  @param *r - the UA_ReadResponse 
  *  @param type - the numeric type to match agains the response
  *  @return true if valid.
*/
UA_Boolean UA_ReadResponse_IsValidScalar(UA_ReadResponse *r, const UA_DataType *type) {
  if (r->responseHeader.serviceResult == UA_STATUSCODE_GOOD &&
        r->resultsSize > 0 && r->results[0].hasValue &&
        UA_Variant_isScalar(&r->results[0].value) &&
        r->results[0].value.type == type) {
    return 1;
  }
  return 0;
}

/**
  * Attempts to read an Int32 value using the given UA_Client
  *  @param *r - the UA_ReadResponse 
  *  @param type - the numeric type to match agains the response
  *  @return true if valid.
*/
struct gopcua_point_int32 gopcua_read_int32(UA_Client *client, char nodeid[]) {
  UA_Int32 value = 0;
  
  UA_ReadRequest rReq;
  UA_ReadRequest_init(&rReq);
  UA_ReadRequest_PrepareSingleRead(&rReq, nodeid);
  
  UA_ReadResponse rResp = UA_Client_Service_read(client, rReq);
  
  if (UA_ReadResponse_IsValidScalar(&rResp, &UA_TYPES[UA_TYPES_INT32])) {
    value = *(UA_Int32*)rResp.results[0].value.data;
  }
  
  gopcua_point_int32 p =  {
    .value =      value,
    .source_ts =  UA_DateTime_toStruct(rResp.results[0].sourceTimestamp),
    .server_ts =  UA_DateTime_toStruct(rResp.results[0].serverTimestamp),
    .has_value =  rResp.results[0].hasValue,   
    .good_value = (rResp.results[0].status == UA_STATUSCODE_GOOD) ? 1 : 0, 
    .has_source_ts = rResp.results[0].hasSourceTimestamp,
    .has_server_ts = rResp.results[0].hasServerTimestamp
  };
  
  UA_ReadRequest_deleteMembers(&rReq);
  UA_ReadResponse_deleteMembers(&rResp);

  return p;
}

/**
  * Attempts to read a Float value using the given UA_Client
  *  @param *r - the UA_ReadResponse 
  *  @param type - the numeric type to match agains the response
  *  @return true if valid.
*/
struct gopcua_point_float32 gopcua_read_float32(UA_Client *client, char nodeid[]) {
  UA_Float value = 0;
  
  UA_ReadRequest rReq;
  UA_ReadRequest_init(&rReq);
  UA_ReadRequest_PrepareSingleRead(&rReq, nodeid);
  
  UA_ReadResponse rResp = UA_Client_Service_read(client, rReq);
  
  if (UA_ReadResponse_IsValidScalar(&rResp, &UA_TYPES[UA_TYPES_FLOAT])) {
    value = *(UA_Float*)rResp.results[0].value.data;
  }
  
  gopcua_point_float32 p =  {
    .value =      value,
    .source_ts =  UA_DateTime_toStruct(rResp.results[0].sourceTimestamp),
    .server_ts =  UA_DateTime_toStruct(rResp.results[0].serverTimestamp),
    .has_value =  rResp.results[0].hasValue,   
    .good_value = (rResp.results[0].status == UA_STATUSCODE_GOOD) ? 1 : 0, 
    .has_source_ts = rResp.results[0].hasSourceTimestamp,
    .has_server_ts = rResp.results[0].hasServerTimestamp
  };
  
  UA_ReadRequest_deleteMembers(&rReq);
  UA_ReadResponse_deleteMembers(&rResp);

  return p;
}



