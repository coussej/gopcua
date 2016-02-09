#ifndef GOPCUA_CLIENT_H
#define GOPCUA_CLIENT_H

#include <open62541.h>

typedef struct gopcua_point_int32 {
  UA_Int32            value;
  UA_DateTimeStruct   server_ts;
  UA_DateTimeStruct   source_ts;
  UA_Boolean          has_value;
  UA_Boolean          good_value;
  UA_Boolean          has_source_ts;
  UA_Boolean          has_server_ts;
} gopcua_point_int32;

typedef struct gopcua_point_float32 {
  UA_Float            value;
  UA_DateTimeStruct   server_ts;
  UA_DateTimeStruct   source_ts;
  UA_Boolean          has_value;
  UA_Boolean          good_value;
  UA_Boolean          has_source_ts;
  UA_Boolean          has_server_ts;
} gopcua_point_float32;

extern UA_Client *gopcua_client_new();

extern int gopcua_client_connect(UA_Client *client, char connInfo[]);

extern struct gopcua_point_int32 gopcua_read_int32(UA_Client *client, char nodeid[]);

extern struct gopcua_point_float32 gopcua_read_float32(UA_Client *client, char nodeid[]);
  
#endif