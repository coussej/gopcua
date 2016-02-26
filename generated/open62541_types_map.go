package gopcua

// This file was auto-generated from ../open62541.h. Do not change it manually! 

/*
#cgo CFLAGS: -std=c99 -I .
#include <stdlib.h>
#include "gopcua_client.h"
 
*/
import "C"

type UA_StatusCode C.UA_StatusCode // line: 413 
const (
	UA_STATUSCODE_GOOD = UA_StatusCode(0x00)
	UA_STATUSCODE_BADUNEXPECTEDERROR = UA_StatusCode(0x80010000) // An unexpected error occurred.
	UA_STATUSCODE_BADINTERNALERROR = UA_StatusCode(0x80020000) // An internal error occurred as a result of a programming or configuration error.
	UA_STATUSCODE_BADOUTOFMEMORY = UA_StatusCode(0x80030000) // Not enough memory to complete the operation.
	UA_STATUSCODE_BADRESOURCEUNAVAILABLE = UA_StatusCode(0x80040000) // An operating system resource is not available.
	UA_STATUSCODE_BADCOMMUNICATIONERROR = UA_StatusCode(0x80050000) // A low level communication error occurred.
	UA_STATUSCODE_BADENCODINGERROR = UA_StatusCode(0x80060000) // Encoding halted because of invalid data in the objects being serialized.
	UA_STATUSCODE_BADDECODINGERROR = UA_StatusCode(0x80070000) // Decoding halted because of invalid data in the stream.
	UA_STATUSCODE_BADENCODINGLIMITSEXCEEDED = UA_StatusCode(0x80080000) // The message encoding/decoding limits imposed by the stack have been exceeded.
	UA_STATUSCODE_BADREQUESTTOOLARGE = UA_StatusCode(0x80b80000) // The request message size exceeds limits set by the server.
	UA_STATUSCODE_BADRESPONSETOOLARGE = UA_StatusCode(0x80b90000) // The response message size exceeds limits set by the client.
	UA_STATUSCODE_BADUNKNOWNRESPONSE = UA_StatusCode(0x80090000) // An unrecognized response was received from the server.
	UA_STATUSCODE_BADTIMEOUT = UA_StatusCode(0x800a0000) // The operation timed out.
	UA_STATUSCODE_BADSERVICEUNSUPPORTED = UA_StatusCode(0x800b0000) // The server does not support the requested service.
	UA_STATUSCODE_BADSHUTDOWN = UA_StatusCode(0x800c0000) // The operation was cancelled because the application is shutting down.
	UA_STATUSCODE_BADSERVERNOTCONNECTED = UA_StatusCode(0x800d0000) // The operation could not complete because the client is not connected to the server.
	UA_STATUSCODE_BADSERVERHALTED = UA_StatusCode(0x800e0000) // The server has stopped and cannot process any requests.
	UA_STATUSCODE_BADNOTHINGTODO = UA_StatusCode(0x800f0000) // There was nothing to do because the client passed a list of operations with no elements.
	UA_STATUSCODE_BADTOOMANYOPERATIONS = UA_StatusCode(0x80100000) // The request could not be processed because it specified too many operations.
	UA_STATUSCODE_BADTOOMANYMONITOREDITEMS = UA_StatusCode(0x80db0000) // The request could not be processed because there are too many monitored items in the subscription.
	UA_STATUSCODE_BADDATATYPEIDUNKNOWN = UA_StatusCode(0x80110000) // The extension object cannot be (de)serialized because the data type id is not recognized.
	UA_STATUSCODE_BADCERTIFICATEINVALID = UA_StatusCode(0x80120000) // The certificate provided as a parameter is not valid.
	UA_STATUSCODE_BADSECURITYCHECKSFAILED = UA_StatusCode(0x80130000) // An error occurred verifying security.
	UA_STATUSCODE_BADCERTIFICATETIMEINVALID = UA_StatusCode(0x80140000) // The Certificate has expired or is not yet valid.
	UA_STATUSCODE_BADCERTIFICATEISSUERTIMEINVALID = UA_StatusCode(0x80150000) // An Issuer Certificate has expired or is not yet valid.
	UA_STATUSCODE_BADCERTIFICATEHOSTNAMEINVALID = UA_StatusCode(0x80160000) // The HostName used to connect to a Server does not match a HostName in the Certificate.
	UA_STATUSCODE_BADCERTIFICATEURIINVALID = UA_StatusCode(0x80170000) // The URI specified in the ApplicationDescription does not match the URI in the Certificate.
	UA_STATUSCODE_BADCERTIFICATEUSENOTALLOWED = UA_StatusCode(0x80180000) // The Certificate may not be used for the requested operation.
	UA_STATUSCODE_BADCERTIFICATEISSUERUSENOTALLOWED = UA_StatusCode(0x80190000) // The Issuer Certificate may not be used for the requested operation.
	UA_STATUSCODE_BADCERTIFICATEUNTRUSTED = UA_StatusCode(0x801a0000) // The Certificate is not trusted.
	UA_STATUSCODE_BADCERTIFICATEREVOCATIONUNKNOWN = UA_StatusCode(0x801b0000) // It was not possible to determine if the Certificate has been revoked.
	UA_STATUSCODE_BADCERTIFICATEISSUERREVOCATIONUNKNOWN = UA_StatusCode(0x801c0000) // It was not possible to determine if the Issuer Certificate has been revoked.
	UA_STATUSCODE_BADCERTIFICATEREVOKED = UA_StatusCode(0x801d0000) // The Certificate has been revoked.
	UA_STATUSCODE_BADCERTIFICATEISSUERREVOKED = UA_StatusCode(0x801e0000) // The Issuer Certificate has been revoked.
	UA_STATUSCODE_BADUSERACCESSDENIED = UA_StatusCode(0x801f0000) // User does not have permission to perform the requested operation.
	UA_STATUSCODE_BADIDENTITYTOKENINVALID = UA_StatusCode(0x80200000) // The user identity token is not valid.
	UA_STATUSCODE_BADIDENTITYTOKENREJECTED = UA_StatusCode(0x80210000) // The user identity token is valid but the server has rejected it.
	UA_STATUSCODE_BADSECURECHANNELIDINVALID = UA_StatusCode(0x80220000) // The specified secure channel is no longer valid.
	UA_STATUSCODE_BADINVALIDTIMESTAMP = UA_StatusCode(0x80230000) // The timestamp is outside the range allowed by the server.
	UA_STATUSCODE_BADNONCEINVALID = UA_StatusCode(0x80240000) // The nonce does appear to be not a random value or it is not the correct length.
	UA_STATUSCODE_BADSESSIONIDINVALID = UA_StatusCode(0x80250000) // The session id is not valid.
	UA_STATUSCODE_BADSESSIONCLOSED = UA_StatusCode(0x80260000) // The session was closed by the client.
	UA_STATUSCODE_BADSESSIONNOTACTIVATED = UA_StatusCode(0x80270000) // The session cannot be used because ActivateSession has not been called.
	UA_STATUSCODE_BADSUBSCRIPTIONIDINVALID = UA_StatusCode(0x80280000) // The subscription id is not valid.
	UA_STATUSCODE_BADREQUESTHEADERINVALID = UA_StatusCode(0x802a0000) // The header for the request is missing or invalid.
	UA_STATUSCODE_BADTIMESTAMPSTORETURNINVALID = UA_StatusCode(0x802b0000) // The timestamps to return parameter is invalid.
	UA_STATUSCODE_BADREQUESTCANCELLEDBYCLIENT = UA_StatusCode(0x802c0000) // The request was cancelled by the client.
	UA_STATUSCODE_GOODSUBSCRIPTIONTRANSFERRED = UA_StatusCode(0x002d0000) // The subscription was transferred to another session.
	UA_STATUSCODE_GOODCOMPLETESASYNCHRONOUSLY = UA_StatusCode(0x002e0000) // The processing will complete asynchronously.
	UA_STATUSCODE_GOODOVERLOAD = UA_StatusCode(0x002f0000) // Sampling has slowed down due to resource limitations.
	UA_STATUSCODE_GOODCLAMPED = UA_StatusCode(0x00300000) // The value written was accepted but was clamped.
	UA_STATUSCODE_BADNOCOMMUNICATION = UA_StatusCode(0x80310000) // Communication with the data source is defined, but not established, and there is no last known value available.
	UA_STATUSCODE_BADWAITINGFORINITIALDATA = UA_StatusCode(0x80320000) // Waiting for the server to obtain values from the underlying data source.
	UA_STATUSCODE_BADNODEIDINVALID = UA_StatusCode(0x80330000) // The syntax of the node id is not valid.
	UA_STATUSCODE_BADNODEIDUNKNOWN = UA_StatusCode(0x80340000) // The node id refers to a node that does not exist in the server address space.
	UA_STATUSCODE_BADATTRIBUTEIDINVALID = UA_StatusCode(0x80350000) // The attribute is not supported for the specified Node.
	UA_STATUSCODE_BADINDEXRANGEINVALID = UA_StatusCode(0x80360000) // The syntax of the index range parameter is invalid.
	UA_STATUSCODE_BADINDEXRANGENODATA = UA_StatusCode(0x80370000) // No data exists within the range of indexes specified.
	UA_STATUSCODE_BADDATAENCODINGINVALID = UA_StatusCode(0x80380000) // The data encoding is invalid.
	UA_STATUSCODE_BADDATAENCODINGUNSUPPORTED = UA_StatusCode(0x80390000) // The server does not support the requested data encoding for the node.
	UA_STATUSCODE_BADNOTREADABLE = UA_StatusCode(0x803a0000) // The access level does not allow reading or subscribing to the Node.
	UA_STATUSCODE_BADNOTWRITABLE = UA_StatusCode(0x803b0000) // The access level does not allow writing to the Node.
	UA_STATUSCODE_BADOUTOFRANGE = UA_StatusCode(0x803c0000) // The value was out of range.
	UA_STATUSCODE_BADNOTSUPPORTED = UA_StatusCode(0x803d0000) // The requested operation is not supported.
	UA_STATUSCODE_BADNOTFOUND = UA_StatusCode(0x803e0000) // A requested item was not found or a search operation ended without success.
	UA_STATUSCODE_BADOBJECTDELETED = UA_StatusCode(0x803f0000) // The object cannot be used because it has been deleted.
	UA_STATUSCODE_BADNOTIMPLEMENTED = UA_StatusCode(0x80400000) // Requested operation is not implemented.
	UA_STATUSCODE_BADMONITORINGMODEINVALID = UA_StatusCode(0x80410000) // The monitoring mode is invalid.
	UA_STATUSCODE_BADMONITOREDITEMIDINVALID = UA_StatusCode(0x80420000) // The monitoring item id does not refer to a valid monitored item.
	UA_STATUSCODE_BADMONITOREDITEMFILTERINVALID = UA_StatusCode(0x80430000) // The monitored item filter parameter is not valid.
	UA_STATUSCODE_BADMONITOREDITEMFILTERUNSUPPORTED = UA_StatusCode(0x80440000) // The server does not support the requested monitored item filter.
	UA_STATUSCODE_BADFILTERNOTALLOWED = UA_StatusCode(0x80450000) // A monitoring filter cannot be used in combination with the attribute specified.
	UA_STATUSCODE_BADSTRUCTUREMISSING = UA_StatusCode(0x80460000) // A mandatory structured parameter was missing or null.
	UA_STATUSCODE_BADEVENTFILTERINVALID = UA_StatusCode(0x80470000) // The event filter is not valid.
	UA_STATUSCODE_BADCONTENTFILTERINVALID = UA_StatusCode(0x80480000) // The content filter is not valid.
	UA_STATUSCODE_BADFILTEROPERATORINVALID = UA_StatusCode(0x80c10000) // An unrecognized operator was provided in a filter.
	UA_STATUSCODE_BADFILTEROPERATORUNSUPPORTED = UA_StatusCode(0x80c20000) // A valid operator was provided, but the server does not provide support for this filter operator.
	UA_STATUSCODE_BADFILTEROPERANDCOUNTMISMATCH = UA_StatusCode(0x80c30000) // The number of operands provided for the filter operator was less then expected for the operand provided.
	UA_STATUSCODE_BADFILTEROPERANDINVALID = UA_StatusCode(0x80490000) // The operand used in a content filter is not valid.
	UA_STATUSCODE_BADFILTERELEMENTINVALID = UA_StatusCode(0x80c40000) // The referenced element is not a valid element in the content filter.
	UA_STATUSCODE_BADFILTERLITERALINVALID = UA_StatusCode(0x80c50000) // The referenced literal is not a valid value.
	UA_STATUSCODE_BADCONTINUATIONPOINTINVALID = UA_StatusCode(0x804a0000) // The continuation point provide is longer valid.
	UA_STATUSCODE_BADNOCONTINUATIONPOINTS = UA_StatusCode(0x804b0000) // The operation could not be processed because all continuation points have been allocated.
	UA_STATUSCODE_BADREFERENCETYPEIDINVALID = UA_StatusCode(0x804c0000) // The operation could not be processed because all continuation points have been allocated.
	UA_STATUSCODE_BADBROWSEDIRECTIONINVALID = UA_StatusCode(0x804d0000) // The browse direction is not valid.
	UA_STATUSCODE_BADNODENOTINVIEW = UA_StatusCode(0x804e0000) // The node is not part of the view.
	UA_STATUSCODE_BADSERVERURIINVALID = UA_StatusCode(0x804f0000) // The ServerUri is not a valid URI.
	UA_STATUSCODE_BADSERVERNAMEMISSING = UA_StatusCode(0x80500000) // No ServerName was specified.
	UA_STATUSCODE_BADDISCOVERYURLMISSING = UA_StatusCode(0x80510000) // No DiscoveryUrl was specified.
	UA_STATUSCODE_BADSEMPAHOREFILEMISSING = UA_StatusCode(0x80520000) // The semaphore file specified by the client is not valid.
	UA_STATUSCODE_BADREQUESTTYPEINVALID = UA_StatusCode(0x80530000) // The security token request type is not valid.
	UA_STATUSCODE_BADSECURITYMODEREJECTED = UA_StatusCode(0x80540000) // The security mode does not meet the requirements set by the Server.
	UA_STATUSCODE_BADSECURITYPOLICYREJECTED = UA_StatusCode(0x80550000) // The security policy does not meet the requirements set by the Server.
	UA_STATUSCODE_BADTOOMANYSESSIONS = UA_StatusCode(0x80560000) // The server has reached its maximum number of sessions.
	UA_STATUSCODE_BADUSERSIGNATUREINVALID = UA_StatusCode(0x80570000) // The user token signature is missing or invalid.
	UA_STATUSCODE_BADAPPLICATIONSIGNATUREINVALID = UA_StatusCode(0x80580000) // The signature generated with the client certificate is missing or invalid.
	UA_STATUSCODE_BADNOVALIDCERTIFICATES = UA_StatusCode(0x80590000) // The client did not provide at least one software certificate that is valid and meets the profile requirements for the server.
	UA_STATUSCODE_BADIDENTITYCHANGENOTSUPPORTED = UA_StatusCode(0x80c60000) // The Server does not support changing the user identity assigned to the session.
	UA_STATUSCODE_BADREQUESTCANCELLEDBYREQUEST = UA_StatusCode(0x805a0000) // The request was canceled by the client with the Cancel service.
	UA_STATUSCODE_BADPARENTNODEIDINVALID = UA_StatusCode(0x805b0000) // The parent node id does not to refer to a valid node.
	UA_STATUSCODE_BADREFERENCENOTALLOWED = UA_StatusCode(0x805c0000) // The reference could not be created because it violates constraints imposed by the data model.
	UA_STATUSCODE_BADNODEIDREJECTED = UA_StatusCode(0x805d0000) // The requested node id was reject because it was either invalid or server does not allow node ids to be specified by the client.
	UA_STATUSCODE_BADNODEIDEXISTS = UA_StatusCode(0x805e0000) // The requested node id is already used by another node.
	UA_STATUSCODE_BADNODECLASSINVALID = UA_StatusCode(0x805f0000) // The node class is not valid.
	UA_STATUSCODE_BADBROWSENAMEINVALID = UA_StatusCode(0x80600000) // The browse name is invalid.
	UA_STATUSCODE_BADBROWSENAMEDUPLICATED = UA_StatusCode(0x80610000) // The browse name is not unique among nodes that share the same relationship with the parent.
	UA_STATUSCODE_BADNODEATTRIBUTESINVALID = UA_StatusCode(0x80620000) // The node attributes are not valid for the node class.
	UA_STATUSCODE_BADTYPEDEFINITIONINVALID = UA_StatusCode(0x80630000) // The type definition node id does not reference an appropriate type node.
	UA_STATUSCODE_BADSOURCENODEIDINVALID = UA_StatusCode(0x80640000) // The source node id does not reference a valid node.
	UA_STATUSCODE_BADTARGETNODEIDINVALID = UA_StatusCode(0x80650000) // The target node id does not reference a valid node.
	UA_STATUSCODE_BADDUPLICATEREFERENCENOTALLOWED = UA_StatusCode(0x80660000) // The reference type between the nodes is already defined.
	UA_STATUSCODE_BADINVALIDSELFREFERENCE = UA_StatusCode(0x80670000) // The server does not allow this type of self reference on this node.
	UA_STATUSCODE_BADREFERENCELOCALONLY = UA_StatusCode(0x80680000) // The reference type is not valid for a reference to a remote server.
	UA_STATUSCODE_BADNODELETERIGHTS = UA_StatusCode(0x80690000) // The server will not allow the node to be deleted.
	UA_STATUSCODE_UNCERTAINREFERENCENOTDELETED = UA_StatusCode(0x40bc0000) // The server was not able to delete all target references.
	UA_STATUSCODE_BADSERVERINDEXINVALID = UA_StatusCode(0x806a0000) // The server index is not valid.
	UA_STATUSCODE_BADVIEWIDUNKNOWN = UA_StatusCode(0x806b0000) // The view id does not refer to a valid view node.
	UA_STATUSCODE_BADVIEWTIMESTAMPINVALID = UA_StatusCode(0x80c90000) // The view timestamp is not available or not supported.
	UA_STATUSCODE_BADVIEWPARAMETERMISMATCH = UA_StatusCode(0x80ca0000) // The view parameters are not consistent with each other.
	UA_STATUSCODE_BADVIEWVERSIONINVALID = UA_StatusCode(0x80cb0000) // The view version is not available or not supported.
	UA_STATUSCODE_UNCERTAINNOTALLNODESAVAILABLE = UA_StatusCode(0x40c00000) // The list of references may not be complete because the underlying system is not available.
	UA_STATUSCODE_GOODRESULTSMAYBEINCOMPLETE = UA_StatusCode(0x00ba0000) // The server should have followed a reference to a node in a remote server but did not. The result set may be incomplete.
	UA_STATUSCODE_BADNOTTYPEDEFINITION = UA_StatusCode(0x80c80000) // The provided Nodeid was not a type definition nodeid.
	UA_STATUSCODE_UNCERTAINREFERENCEOUTOFSERVER = UA_StatusCode(0x406c0000) // One of the references to follow in the relative path references to a node in the address space in another server.
	UA_STATUSCODE_BADTOOMANYMATCHES = UA_StatusCode(0x806d0000) // The requested operation has too many matches to return.
	UA_STATUSCODE_BADQUERYTOOCOMPLEX = UA_StatusCode(0x806e0000) // The requested operation requires too many resources in the server.
	UA_STATUSCODE_BADNOMATCH = UA_StatusCode(0x806f0000) // The requested operation has no match to return.
	UA_STATUSCODE_BADMAXAGEINVALID = UA_StatusCode(0x80700000) // The max age parameter is invalid.
	UA_STATUSCODE_BADHISTORYOPERATIONINVALID = UA_StatusCode(0x80710000) // The history details parameter is not valid.
	UA_STATUSCODE_BADHISTORYOPERATIONUNSUPPORTED = UA_StatusCode(0x80720000) // The server does not support the requested operation.
	UA_STATUSCODE_BADINVALIDTIMESTAMPARGUMENT = UA_StatusCode(0x80bd0000) // The defined timestamp to return was invalid.
	UA_STATUSCODE_BADWRITENOTSUPPORTED = UA_StatusCode(0x80730000) // The server not does support writing the combination of value, status and timestamps provided.
	UA_STATUSCODE_BADTYPEMISMATCH = UA_StatusCode(0x80740000) // The value supplied for the attribute is not of the same type as the attribute's value.
	UA_STATUSCODE_BADMETHODINVALID = UA_StatusCode(0x80750000) // The method id does not refer to a method for the specified object.
	UA_STATUSCODE_BADARGUMENTSMISSING = UA_StatusCode(0x80760000) // The client did not specify all of the input arguments for the method.
	UA_STATUSCODE_BADTOOMANYSUBSCRIPTIONS = UA_StatusCode(0x80770000) // The server has reached its  maximum number of subscriptions.
	UA_STATUSCODE_BADTOOMANYPUBLISHREQUESTS = UA_StatusCode(0x80780000) // The server has reached the maximum number of queued publish requests.
	UA_STATUSCODE_BADNOSUBSCRIPTION = UA_StatusCode(0x80790000) // There is no subscription available for this session.
	UA_STATUSCODE_BADSEQUENCENUMBERUNKNOWN = UA_StatusCode(0x807a0000) // The sequence number is unknown to the server.
	UA_STATUSCODE_BADMESSAGENOTAVAILABLE = UA_StatusCode(0x807b0000) // The requested notification message is no longer available.
	UA_STATUSCODE_BADINSUFFICIENTCLIENTPROFILE = UA_StatusCode(0x807c0000) // The Client of the current Session does not support one or more Profiles that are necessary for the Subscription.
	UA_STATUSCODE_BADSTATENOTACTIVE = UA_StatusCode(0x80bf0000) // The sub-state machine is not currently active.
	UA_STATUSCODE_BADTCPSERVERTOOBUSY = UA_StatusCode(0x807d0000) // The server cannot process the request because it is too busy.
	UA_STATUSCODE_BADTCPMESSAGETYPEINVALID = UA_StatusCode(0x807e0000) // The type of the message specified in the header invalid.
	UA_STATUSCODE_BADTCPSECURECHANNELUNKNOWN = UA_StatusCode(0x807f0000) // The SecureChannelId and/or TokenId are not currently in use.
	UA_STATUSCODE_BADTCPMESSAGETOOLARGE = UA_StatusCode(0x80800000) // The size of the message specified in the header is too large.
	UA_STATUSCODE_BADTCPNOTENOUGHRESOURCES = UA_StatusCode(0x80810000) // There are not enough resources to process the request.
	UA_STATUSCODE_BADTCPINTERNALERROR = UA_StatusCode(0x80820000) // An internal error occurred.
	UA_STATUSCODE_BADTCPENDPOINTURLINVALID = UA_StatusCode(0x80830000) // The Server does not recognize the QueryString specified.
	UA_STATUSCODE_BADREQUESTINTERRUPTED = UA_StatusCode(0x80840000) // The request could not be sent because of a network interruption.
	UA_STATUSCODE_BADREQUESTTIMEOUT = UA_StatusCode(0x80850000) // Timeout occurred while processing the request.
	UA_STATUSCODE_BADSECURECHANNELCLOSED = UA_StatusCode(0x80860000) // The secure channel has been closed.
	UA_STATUSCODE_BADSECURECHANNELTOKENUNKNOWN = UA_StatusCode(0x80870000) // The token has expired or is not recognized.
	UA_STATUSCODE_BADSEQUENCENUMBERINVALID = UA_StatusCode(0x80880000) // The sequence number is not valid.
	UA_STATUSCODE_BADPROTOCOLVERSIONUNSUPPORTED = UA_StatusCode(0x80be0000) // The applications do not have compatible protocol versions.
	UA_STATUSCODE_BADCONFIGURATIONERROR = UA_StatusCode(0x80890000) // There is a problem with the configuration that affects the usefulness of the value.
	UA_STATUSCODE_BADNOTCONNECTED = UA_StatusCode(0x808a0000) // The variable should receive its value from another variable, but has never been configured to do so.
	UA_STATUSCODE_BADDEVICEFAILURE = UA_StatusCode(0x808b0000) // There has been a failure in the device/data source that generates the value that has affected the value.
	UA_STATUSCODE_BADSENSORFAILURE = UA_StatusCode(0x808c0000) // There has been a failure in the sensor from which the value is derived by the device/data source.
	UA_STATUSCODE_BADOUTOFSERVICE = UA_StatusCode(0x808d0000) // The source of the data is not operational.
	UA_STATUSCODE_BADDEADBANDFILTERINVALID = UA_StatusCode(0x808e0000) // The deadband filter is not valid.
	UA_STATUSCODE_UNCERTAINNOCOMMUNICATIONLASTUSABLEVALUE = UA_StatusCode(0x408f0000) // Communication to the data source has failed. The variable value is the last value that had a good quality.
	UA_STATUSCODE_UNCERTAINLASTUSABLEVALUE = UA_StatusCode(0x40900000) // Whatever was updating this value has stopped doing so.
	UA_STATUSCODE_UNCERTAINSUBSTITUTEVALUE = UA_StatusCode(0x40910000) // The value is an operational value that was manually overwritten.
	UA_STATUSCODE_UNCERTAININITIALVALUE = UA_StatusCode(0x40920000) // The value is an initial value for a variable that normally receives its value from another variable.
	UA_STATUSCODE_UNCERTAINSENSORNOTACCURATE = UA_StatusCode(0x40930000) // The value is at one of the sensor limits.
	UA_STATUSCODE_UNCERTAINENGINEERINGUNITSEXCEEDED = UA_StatusCode(0x40940000) // The value is outside of the range of values defined for this parameter.
	UA_STATUSCODE_UNCERTAINSUBNORMAL = UA_StatusCode(0x40950000) // The value is derived from multiple sources and has less than the required number of Good sources.
	UA_STATUSCODE_GOODLOCALOVERRIDE = UA_StatusCode(0x00960000) // The value has been overridden.
	UA_STATUSCODE_BADREFRESHINPROGRESS = UA_StatusCode(0x80970000) // This Condition refresh failed, a Condition refresh operation is already in progress.
	UA_STATUSCODE_BADCONDITIONALREADYDISABLED = UA_StatusCode(0x80980000) // This condition has already been disabled.
	UA_STATUSCODE_BADCONDITIONALREADYENABLED = UA_StatusCode(0x80cc0000) // This condition has already been enabled.
	UA_STATUSCODE_BADCONDITIONDISABLED = UA_StatusCode(0x80990000) // Property not available, this condition is disabled.
	UA_STATUSCODE_BADEVENTIDUNKNOWN = UA_StatusCode(0x809a0000) // The specified event id is not recognized.
	UA_STATUSCODE_BADEVENTNOTACKNOWLEDGEABLE = UA_StatusCode(0x80bb0000) // The event cannot be acknowledged.
	UA_STATUSCODE_BADDIALOGNOTACTIVE = UA_StatusCode(0x80cd0000) // The dialog condition is not active.
	UA_STATUSCODE_BADDIALOGRESPONSEINVALID = UA_StatusCode(0x80ce0000) // The response is not valid for the dialog.
	UA_STATUSCODE_BADCONDITIONBRANCHALREADYACKED = UA_StatusCode(0x80cf0000) // The condition branch has already been acknowledged.
	UA_STATUSCODE_BADCONDITIONBRANCHALREADYCONFIRMED = UA_StatusCode(0x80d00000) // The condition branch has already been confirmed.
	UA_STATUSCODE_BADCONDITIONALREADYSHELVED = UA_StatusCode(0x80d10000) // The condition has already been shelved.
	UA_STATUSCODE_BADCONDITIONNOTSHELVED = UA_StatusCode(0x80d20000) // The condition is not currently shelved.
	UA_STATUSCODE_BADSHELVINGTIMEOUTOFRANGE = UA_StatusCode(0x80d30000) // The shelving time not within an acceptable range.
	UA_STATUSCODE_BADNODATA = UA_StatusCode(0x809b0000) // No data exists for the requested time range or event filter.
	UA_STATUSCODE_BADBOUNDNOTFOUND = UA_StatusCode(0x80d70000) // No data found to provide upper or lower bound value.
	UA_STATUSCODE_BADBOUNDNOTSUPPORTED = UA_StatusCode(0x80d80000) // The server cannot retrieve a bound for the variable.
	UA_STATUSCODE_BADDATALOST = UA_StatusCode(0x809d0000) // Data is missing due to collection started/stopped/lost.
	UA_STATUSCODE_BADDATAUNAVAILABLE = UA_StatusCode(0x809e0000) // Expected data is unavailable for the requested time range due to an un-mounted volume, an off-line archive or tape, or similar reason for temporary unavailability.
	UA_STATUSCODE_BADENTRYEXISTS = UA_StatusCode(0x809f0000) // The data or event was not successfully inserted because a matching entry exists.
	UA_STATUSCODE_BADNOENTRYEXISTS = UA_StatusCode(0x80a00000) // The data or event was not successfully updated because no matching entry exists.
	UA_STATUSCODE_BADTIMESTAMPNOTSUPPORTED = UA_StatusCode(0x80a10000) // The client requested history using a timestamp format the server does not support (i.e requested ServerTimestamp when server only supports SourceTimestamp).
	UA_STATUSCODE_GOODENTRYINSERTED = UA_StatusCode(0x00a20000) // The data or event was successfully inserted into the historical database.
	UA_STATUSCODE_GOODENTRYREPLACED = UA_StatusCode(0x00a30000) // The data or event field was successfully replaced in the historical database.
	UA_STATUSCODE_UNCERTAINDATASUBNORMAL = UA_StatusCode(0x40a40000) // The value is derived from multiple values and has less than the required number of Good values.
	UA_STATUSCODE_GOODNODATA = UA_StatusCode(0x00a50000) // No data exists for the requested time range or event filter.
	UA_STATUSCODE_GOODMOREDATA = UA_StatusCode(0x00a60000) // The data or event field was successfully replaced in the historical database.
	UA_STATUSCODE_BADAGGREGATELISTMISMATCH = UA_StatusCode(0x80d40000) // The requested number of Aggregates does not match the requested number of NodeIds.
	UA_STATUSCODE_BADAGGREGATENOTSUPPORTED = UA_StatusCode(0x80d50000) // The requested Aggregate is not support by the server.
	UA_STATUSCODE_BADAGGREGATEINVALIDINPUTS = UA_StatusCode(0x80d60000) // The aggregate value could not be derived due to invalid data inputs.
	UA_STATUSCODE_BADAGGREGATECONFIGURATIONREJECTED = UA_StatusCode(0x80da0000) // The aggregate configuration is not valid for specified node.
	UA_STATUSCODE_GOODDATAIGNORED = UA_StatusCode(0x00d90000) // The request specifies fields which are not valid for the EventType or cannot be saved by the historian.
	UA_STATUSCODE_GOODCOMMUNICATIONEVENT = UA_StatusCode(0x00a70000) // The communication layer has raised an event.
	UA_STATUSCODE_GOODSHUTDOWNEVENT = UA_StatusCode(0x00a80000) // The system is shutting down.
	UA_STATUSCODE_GOODCALLAGAIN = UA_StatusCode(0x00a90000) // The operation is not finished and needs to be called again.
	UA_STATUSCODE_GOODNONCRITICALTIMEOUT = UA_StatusCode(0x00aa0000) // A non-critical timeout occurred.
	UA_STATUSCODE_BADINVALIDARGUMENT = UA_StatusCode(0x80ab0000) // One or more arguments are invalid.
	UA_STATUSCODE_BADCONNECTIONREJECTED = UA_StatusCode(0x80ac0000) // Could not establish a network connection to remote server.
	UA_STATUSCODE_BADDISCONNECT = UA_StatusCode(0x80ad0000) // The server has disconnected from the client.
	UA_STATUSCODE_BADCONNECTIONCLOSED = UA_StatusCode(0x80ae0000) // The network connection has been closed.
	UA_STATUSCODE_BADINVALIDSTATE = UA_StatusCode(0x80af0000) // The operation cannot be completed because the object is closed, uninitialized or in some other invalid state.
	UA_STATUSCODE_BADENDOFSTREAM = UA_StatusCode(0x80b00000) // Cannot move beyond end of the stream.
	UA_STATUSCODE_BADNODATAAVAILABLE = UA_StatusCode(0x80b10000) // No data is currently available for reading from a non-blocking stream.
	UA_STATUSCODE_BADWAITINGFORRESPONSE = UA_StatusCode(0x80b20000) // The asynchronous operation is waiting for a response.
	UA_STATUSCODE_BADOPERATIONABANDONED = UA_StatusCode(0x80b30000) // The asynchronous operation was abandoned by the caller.
	UA_STATUSCODE_BADEXPECTEDSTREAMTOBLOCK = UA_StatusCode(0x80b40000) // The stream did not return all data requested (possibly because it is a non-blocking stream).
	UA_STATUSCODE_BADWOULDBLOCK = UA_StatusCode(0x80b50000) // Non blocking behaviour is required and the operation would block.
	UA_STATUSCODE_BADSYNTAXERROR = UA_StatusCode(0x80b60000) // A value had an invalid syntax.
	UA_STATUSCODE_BADMAXCONNECTIONSREACHED = UA_StatusCode(0x80b70000) // The operation could not be finished because all available connections are in use.
)
type UA_Boolean C.UA_Boolean // line: 454 
type UA_SByte C.UA_SByte // line: 459 
type UA_Byte C.UA_Byte // line: 464 
type UA_Int16 C.UA_Int16 // line: 469 
type UA_UInt16 C.UA_UInt16 // line: 474 
type UA_Int32 C.UA_Int32 // line: 479 
type UA_UInt32 C.UA_UInt32 // line: 484 
type UA_Int64 C.UA_Int64 // line: 492 
type UA_UInt64 C.UA_UInt64 // line: 497 
type UA_Float C.UA_Float // line: 502 
type UA_Double C.UA_Double // line: 505 
type UA_String C.UA_String // line: 513 
type UA_DateTime C.UA_DateTime // line: 534 
type UA_DateTimeStruct C.UA_DateTimeStruct // line: 560 
type UA_Guid C.UA_Guid // line: 574 
type UA_ByteString C.UA_ByteString // line: 584 
type UA_XmlElement C.UA_XmlElement // line: 606 
type UA_NodeId C.UA_NodeId // line: 627 
type UA_ExpandedNodeId C.UA_ExpandedNodeId // line: 669 
type UA_QualifiedName C.UA_QualifiedName // line: 701 
type UA_LocalizedText C.UA_LocalizedText // line: 715 
type UA_DataType C.UA_DataType // line: 727 
type UA_ExtensionObject C.UA_ExtensionObject // line: 751 
type UA_Variant C.UA_Variant // line: 773 
type UA_NumericRangeDimension C.UA_NumericRangeDimension // line: 841 
type UA_NumericRange C.UA_NumericRange // line: 846 
type UA_DataValue C.UA_DataValue // line: 904 
type UA_DiagnosticInfo C.UA_DiagnosticInfo // line: 925 
type UA_DataTypeMember C.UA_DataTypeMember // line: 945 
type UA_AttributeId C.UA_AttributeId // line: 1074 
const (
    UA_ATTRIBUTEID_NODEID                  = UA_AttributeId(1)
    UA_ATTRIBUTEID_NODECLASS               = UA_AttributeId(2)
    UA_ATTRIBUTEID_BROWSENAME              = UA_AttributeId(3)
    UA_ATTRIBUTEID_DISPLAYNAME             = UA_AttributeId(4)
    UA_ATTRIBUTEID_DESCRIPTION             = UA_AttributeId(5)
    UA_ATTRIBUTEID_WRITEMASK               = UA_AttributeId(6)
    UA_ATTRIBUTEID_USERWRITEMASK           = UA_AttributeId(7)
    UA_ATTRIBUTEID_ISABSTRACT              = UA_AttributeId(8)
    UA_ATTRIBUTEID_SYMMETRIC               = UA_AttributeId(9)
    UA_ATTRIBUTEID_INVERSENAME             = UA_AttributeId(10)
    UA_ATTRIBUTEID_CONTAINSNOLOOPS         = UA_AttributeId(11)
    UA_ATTRIBUTEID_EVENTNOTIFIER           = UA_AttributeId(12)
    UA_ATTRIBUTEID_VALUE                   = UA_AttributeId(13)
    UA_ATTRIBUTEID_DATATYPE                = UA_AttributeId(14)
    UA_ATTRIBUTEID_VALUERANK               = UA_AttributeId(15)
    UA_ATTRIBUTEID_ARRAYDIMENSIONS         = UA_AttributeId(16)
    UA_ATTRIBUTEID_ACCESSLEVEL             = UA_AttributeId(17)
    UA_ATTRIBUTEID_USERACCESSLEVEL         = UA_AttributeId(18)
    UA_ATTRIBUTEID_MINIMUMSAMPLINGINTERVAL = UA_AttributeId(19)
    UA_ATTRIBUTEID_HISTORIZING             = UA_AttributeId(20)
    UA_ATTRIBUTEID_EXECUTABLE              = UA_AttributeId(21)
    UA_ATTRIBUTEID_USEREXECUTABLE          = UA_AttributeId(22)
)
type UA_AccessLevelMask C.UA_AccessLevelMask // line: 1082 
const (
    UA_ACCESSLEVELMASK_READ = UA_AccessLevelMask(0x01)
    UA_ACCESSLEVELMASK_WRITE = UA_AccessLevelMask(0x02)
    UA_ACCESSLEVELMASK_HISTORYREAD = UA_AccessLevelMask(0x4)
    UA_ACCESSLEVELMASK_HISTORYWRITE = UA_AccessLevelMask(0x08)
    UA_ACCESSLEVELMASK_SEMANTICCHANGE = UA_AccessLevelMask(0x10)
)
type UA_IdType C.UA_IdType // line: 2017 
const (
    UA_IDTYPE_NUMERIC = UA_IdType(0)
    UA_IDTYPE_STRING = UA_IdType(1)
    UA_IDTYPE_GUID = UA_IdType(2)
    UA_IDTYPE_OPAQUE = UA_IdType(3)
)
type UA_NodeClass C.UA_NodeClass // line: 2036 
const (
    UA_NODECLASS_UNSPECIFIED = UA_NodeClass(0)
    UA_NODECLASS_OBJECT = UA_NodeClass(1)
    UA_NODECLASS_VARIABLE = UA_NodeClass(2)
    UA_NODECLASS_METHOD = UA_NodeClass(4)
    UA_NODECLASS_OBJECTTYPE = UA_NodeClass(8)
    UA_NODECLASS_VARIABLETYPE = UA_NodeClass(16)
    UA_NODECLASS_REFERENCETYPE = UA_NodeClass(32)
    UA_NODECLASS_DATATYPE = UA_NodeClass(64)
    UA_NODECLASS_VIEW = UA_NodeClass(128)
)
type UA_ReferenceNode C.UA_ReferenceNode // line: 2049 
type UA_Argument C.UA_Argument // line: 2065 
type UA_ApplicationType C.UA_ApplicationType // line: 2079 
const (
    UA_APPLICATIONTYPE_SERVER = UA_ApplicationType(0)
    UA_APPLICATIONTYPE_CLIENT = UA_ApplicationType(1)
    UA_APPLICATIONTYPE_CLIENTANDSERVER = UA_ApplicationType(2)
    UA_APPLICATIONTYPE_DISCOVERYSERVER = UA_ApplicationType(3)
)
type UA_ApplicationDescription C.UA_ApplicationDescription // line: 2097 
type UA_RequestHeader C.UA_RequestHeader // line: 2114 
type UA_ResponseHeader C.UA_ResponseHeader // line: 2131 
type UA_ServiceFault C.UA_ServiceFault // line: 2142 
type UA_FindServersRequest C.UA_FindServersRequest // line: 2158 
type UA_FindServersResponse C.UA_FindServersResponse // line: 2171 
type UA_MessageSecurityMode C.UA_MessageSecurityMode // line: 2185 
const (
    UA_MESSAGESECURITYMODE_INVALID = UA_MessageSecurityMode(0)
    UA_MESSAGESECURITYMODE_NONE = UA_MessageSecurityMode(1)
    UA_MESSAGESECURITYMODE_SIGN = UA_MessageSecurityMode(2)
    UA_MESSAGESECURITYMODE_SIGNANDENCRYPT = UA_MessageSecurityMode(3)
)
type UA_UserTokenType C.UA_UserTokenType // line: 2199 
const (
    UA_USERTOKENTYPE_ANONYMOUS = UA_UserTokenType(0)
    UA_USERTOKENTYPE_USERNAME = UA_UserTokenType(1)
    UA_USERTOKENTYPE_CERTIFICATE = UA_UserTokenType(2)
    UA_USERTOKENTYPE_ISSUEDTOKEN = UA_UserTokenType(3)
)
type UA_UserTokenPolicy C.UA_UserTokenPolicy // line: 2214 
type UA_EndpointDescription C.UA_EndpointDescription // line: 2233 
type UA_GetEndpointsRequest C.UA_GetEndpointsRequest // line: 2249 
type UA_GetEndpointsResponse C.UA_GetEndpointsResponse // line: 2262 
type UA_SecurityTokenRequestType C.UA_SecurityTokenRequestType // line: 2274 
const (
    UA_SECURITYTOKENREQUESTTYPE_ISSUE = UA_SecurityTokenRequestType(0)
    UA_SECURITYTOKENREQUESTTYPE_RENEW = UA_SecurityTokenRequestType(1)
)
type UA_ChannelSecurityToken C.UA_ChannelSecurityToken // line: 2288 
type UA_OpenSecureChannelRequest C.UA_OpenSecureChannelRequest // line: 2304 
type UA_OpenSecureChannelResponse C.UA_OpenSecureChannelResponse // line: 2318 
type UA_CloseSecureChannelRequest C.UA_CloseSecureChannelRequest // line: 2329 
type UA_CloseSecureChannelResponse C.UA_CloseSecureChannelResponse // line: 2340 
type UA_SignedSoftwareCertificate C.UA_SignedSoftwareCertificate // line: 2352 
type UA_SignatureData C.UA_SignatureData // line: 2364 
type UA_CreateSessionRequest C.UA_CreateSessionRequest // line: 2383 
type UA_CreateSessionResponse C.UA_CreateSessionResponse // line: 2405 
type UA_UserIdentityToken C.UA_UserIdentityToken // line: 2416 
type UA_AnonymousIdentityToken C.UA_AnonymousIdentityToken // line: 2427 
type UA_UserNameIdentityToken C.UA_UserNameIdentityToken // line: 2441 
type UA_ActivateSessionRequest C.UA_ActivateSessionRequest // line: 2459 
type UA_ActivateSessionResponse C.UA_ActivateSessionResponse // line: 2475 
type UA_CloseSessionRequest C.UA_CloseSessionRequest // line: 2487 
type UA_CloseSessionResponse C.UA_CloseSessionResponse // line: 2498 
type UA_NodeAttributesMask C.UA_NodeAttributesMask // line: 2540 
const (
    UA_NODEATTRIBUTESMASK_NONE = UA_NodeAttributesMask(0)
    UA_NODEATTRIBUTESMASK_ACCESSLEVEL = UA_NodeAttributesMask(1)
    UA_NODEATTRIBUTESMASK_ARRAYDIMENSIONS = UA_NodeAttributesMask(2)
    UA_NODEATTRIBUTESMASK_BROWSENAME = UA_NodeAttributesMask(4)
    UA_NODEATTRIBUTESMASK_CONTAINSNOLOOPS = UA_NodeAttributesMask(8)
    UA_NODEATTRIBUTESMASK_DATATYPE = UA_NodeAttributesMask(16)
    UA_NODEATTRIBUTESMASK_DESCRIPTION = UA_NodeAttributesMask(32)
    UA_NODEATTRIBUTESMASK_DISPLAYNAME = UA_NodeAttributesMask(64)
    UA_NODEATTRIBUTESMASK_EVENTNOTIFIER = UA_NodeAttributesMask(128)
    UA_NODEATTRIBUTESMASK_EXECUTABLE = UA_NodeAttributesMask(256)
    UA_NODEATTRIBUTESMASK_HISTORIZING = UA_NodeAttributesMask(512)
    UA_NODEATTRIBUTESMASK_INVERSENAME = UA_NodeAttributesMask(1024)
    UA_NODEATTRIBUTESMASK_ISABSTRACT = UA_NodeAttributesMask(2048)
    UA_NODEATTRIBUTESMASK_MINIMUMSAMPLINGINTERVAL = UA_NodeAttributesMask(4096)
    UA_NODEATTRIBUTESMASK_NODECLASS = UA_NodeAttributesMask(8192)
    UA_NODEATTRIBUTESMASK_NODEID = UA_NodeAttributesMask(16384)
    UA_NODEATTRIBUTESMASK_SYMMETRIC = UA_NodeAttributesMask(32768)
    UA_NODEATTRIBUTESMASK_USERACCESSLEVEL = UA_NodeAttributesMask(65536)
    UA_NODEATTRIBUTESMASK_USEREXECUTABLE = UA_NodeAttributesMask(131072)
    UA_NODEATTRIBUTESMASK_USERWRITEMASK = UA_NodeAttributesMask(262144)
    UA_NODEATTRIBUTESMASK_VALUERANK = UA_NodeAttributesMask(524288)
    UA_NODEATTRIBUTESMASK_WRITEMASK = UA_NodeAttributesMask(1048576)
    UA_NODEATTRIBUTESMASK_VALUE = UA_NodeAttributesMask(2097152)
    UA_NODEATTRIBUTESMASK_ALL = UA_NodeAttributesMask(4194303)
    UA_NODEATTRIBUTESMASK_BASENODE = UA_NodeAttributesMask(1335396)
    UA_NODEATTRIBUTESMASK_OBJECT = UA_NodeAttributesMask(1335524)
    UA_NODEATTRIBUTESMASK_OBJECTTYPEORDATATYPE = UA_NodeAttributesMask(1337444)
    UA_NODEATTRIBUTESMASK_VARIABLE = UA_NodeAttributesMask(4026999)
    UA_NODEATTRIBUTESMASK_VARIABLETYPE = UA_NodeAttributesMask(3958902)
    UA_NODEATTRIBUTESMASK_METHOD = UA_NodeAttributesMask(1466724)
    UA_NODEATTRIBUTESMASK_REFERENCETYPE = UA_NodeAttributesMask(1371236)
    UA_NODEATTRIBUTESMASK_VIEW = UA_NodeAttributesMask(1335532)
)
type UA_NodeAttributes C.UA_NodeAttributes // line: 2555 
type UA_ObjectAttributes C.UA_ObjectAttributes // line: 2571 
type UA_VariableAttributes C.UA_VariableAttributes // line: 2595 
type UA_MethodAttributes C.UA_MethodAttributes // line: 2612 
type UA_ObjectTypeAttributes C.UA_ObjectTypeAttributes // line: 2628 
type UA_VariableTypeAttributes C.UA_VariableTypeAttributes // line: 2649 
type UA_ReferenceTypeAttributes C.UA_ReferenceTypeAttributes // line: 2667 
type UA_DataTypeAttributes C.UA_DataTypeAttributes // line: 2683 
type UA_ViewAttributes C.UA_ViewAttributes // line: 2700 
type UA_AddNodesItem C.UA_AddNodesItem // line: 2717 
type UA_AddNodesResult C.UA_AddNodesResult // line: 2729 
type UA_AddNodesRequest C.UA_AddNodesRequest // line: 2742 
type UA_AddNodesResponse C.UA_AddNodesResponse // line: 2757 
type UA_AddReferencesItem C.UA_AddReferencesItem // line: 2773 
type UA_AddReferencesRequest C.UA_AddReferencesRequest // line: 2786 
type UA_AddReferencesResponse C.UA_AddReferencesResponse // line: 2801 
type UA_DeleteNodesItem C.UA_DeleteNodesItem // line: 2813 
type UA_DeleteNodesRequest C.UA_DeleteNodesRequest // line: 2826 
type UA_DeleteNodesResponse C.UA_DeleteNodesResponse // line: 2841 
type UA_DeleteReferencesItem C.UA_DeleteReferencesItem // line: 2856 
type UA_DeleteReferencesRequest C.UA_DeleteReferencesRequest // line: 2869 
type UA_DeleteReferencesResponse C.UA_DeleteReferencesResponse // line: 2884 
type UA_BrowseDirection C.UA_BrowseDirection // line: 2897 
const (
    UA_BROWSEDIRECTION_FORWARD = UA_BrowseDirection(0)
    UA_BROWSEDIRECTION_INVERSE = UA_BrowseDirection(1)
    UA_BROWSEDIRECTION_BOTH = UA_BrowseDirection(2)
)
type UA_ViewDescription C.UA_ViewDescription // line: 2910 
type UA_BrowseDescription C.UA_BrowseDescription // line: 2926 
type UA_BrowseResultMask C.UA_BrowseResultMask // line: 2946 
const (
    UA_BROWSERESULTMASK_NONE = UA_BrowseResultMask(0)
    UA_BROWSERESULTMASK_REFERENCETYPEID = UA_BrowseResultMask(1)
    UA_BROWSERESULTMASK_ISFORWARD = UA_BrowseResultMask(2)
    UA_BROWSERESULTMASK_NODECLASS = UA_BrowseResultMask(4)
    UA_BROWSERESULTMASK_BROWSENAME = UA_BrowseResultMask(8)
    UA_BROWSERESULTMASK_DISPLAYNAME = UA_BrowseResultMask(16)
    UA_BROWSERESULTMASK_TYPEDEFINITION = UA_BrowseResultMask(32)
    UA_BROWSERESULTMASK_ALL = UA_BrowseResultMask(63)
    UA_BROWSERESULTMASK_REFERENCETYPEINFO = UA_BrowseResultMask(3)
    UA_BROWSERESULTMASK_TARGETINFO = UA_BrowseResultMask(60)
)
type UA_ReferenceDescription C.UA_ReferenceDescription // line: 2963 
type UA_BrowseResult C.UA_BrowseResult // line: 2977 
type UA_BrowseRequest C.UA_BrowseRequest // line: 2992 
type UA_BrowseResponse C.UA_BrowseResponse // line: 3007 
type UA_BrowseNextRequest C.UA_BrowseNextRequest // line: 3021 
type UA_BrowseNextResponse C.UA_BrowseNextResponse // line: 3036 
type UA_RelativePathElement C.UA_RelativePathElement // line: 3050 
type UA_RelativePath C.UA_RelativePath // line: 3062 
type UA_BrowsePath C.UA_BrowsePath // line: 3074 
type UA_BrowsePathTarget C.UA_BrowsePathTarget // line: 3086 
type UA_BrowsePathResult C.UA_BrowsePathResult // line: 3099 
type UA_TranslateBrowsePathsToNodeIdsRequest C.UA_TranslateBrowsePathsToNodeIdsRequest // line: 3112 
type UA_TranslateBrowsePathsToNodeIdsResponse C.UA_TranslateBrowsePathsToNodeIdsResponse // line: 3127 
type UA_RegisterNodesRequest C.UA_RegisterNodesRequest // line: 3140 
type UA_RegisterNodesResponse C.UA_RegisterNodesResponse // line: 3153 
type UA_UnregisterNodesRequest C.UA_UnregisterNodesRequest // line: 3166 
type UA_UnregisterNodesResponse C.UA_UnregisterNodesResponse // line: 3177 
type UA_QueryDataDescription C.UA_QueryDataDescription // line: 3189 
type UA_NodeTypeDescription C.UA_NodeTypeDescription // line: 3202 
type UA_FilterOperator C.UA_FilterOperator // line: 3229 
const (
    UA_FILTEROPERATOR_EQUALS = UA_FilterOperator(0)
    UA_FILTEROPERATOR_ISNULL = UA_FilterOperator(1)
    UA_FILTEROPERATOR_GREATERTHAN = UA_FilterOperator(2)
    UA_FILTEROPERATOR_LESSTHAN = UA_FilterOperator(3)
    UA_FILTEROPERATOR_GREATERTHANOREQUAL = UA_FilterOperator(4)
    UA_FILTEROPERATOR_LESSTHANOREQUAL = UA_FilterOperator(5)
    UA_FILTEROPERATOR_LIKE = UA_FilterOperator(6)
    UA_FILTEROPERATOR_NOT = UA_FilterOperator(7)
    UA_FILTEROPERATOR_BETWEEN = UA_FilterOperator(8)
    UA_FILTEROPERATOR_INLIST = UA_FilterOperator(9)
    UA_FILTEROPERATOR_AND = UA_FilterOperator(10)
    UA_FILTEROPERATOR_OR = UA_FilterOperator(11)
    UA_FILTEROPERATOR_CAST = UA_FilterOperator(12)
    UA_FILTEROPERATOR_INVIEW = UA_FilterOperator(13)
    UA_FILTEROPERATOR_OFTYPE = UA_FilterOperator(14)
    UA_FILTEROPERATOR_RELATEDTO = UA_FilterOperator(15)
    UA_FILTEROPERATOR_BITWISEAND = UA_FilterOperator(16)
    UA_FILTEROPERATOR_BITWISEOR = UA_FilterOperator(17)
)
type UA_QueryDataSet C.UA_QueryDataSet // line: 3242 
type UA_ContentFilterElement C.UA_ContentFilterElement // line: 3254 
type UA_ContentFilter C.UA_ContentFilter // line: 3265 
type UA_ContentFilterElementResult C.UA_ContentFilterElementResult // line: 3279 
type UA_ContentFilterResult C.UA_ContentFilterResult // line: 3292 
type UA_ParsingResult C.UA_ParsingResult // line: 3306 
type UA_QueryFirstRequest C.UA_QueryFirstRequest // line: 3322 
type UA_QueryFirstResponse C.UA_QueryFirstResponse // line: 3340 
type UA_QueryNextRequest C.UA_QueryNextRequest // line: 3352 
type UA_QueryNextResponse C.UA_QueryNextResponse // line: 3365 
type UA_TimestampsToReturn C.UA_TimestampsToReturn // line: 3378 
const (
    UA_TIMESTAMPSTORETURN_SOURCE = UA_TimestampsToReturn(0)
    UA_TIMESTAMPSTORETURN_SERVER = UA_TimestampsToReturn(1)
    UA_TIMESTAMPSTORETURN_BOTH = UA_TimestampsToReturn(2)
    UA_TIMESTAMPSTORETURN_NEITHER = UA_TimestampsToReturn(3)
)
type UA_ReadValueId C.UA_ReadValueId // line: 3391 
type UA_ReadRequest C.UA_ReadRequest // line: 3405 
type UA_ReadResponse C.UA_ReadResponse // line: 3419 
type UA_WriteValue C.UA_WriteValue // line: 3432 
type UA_WriteRequest C.UA_WriteRequest // line: 3444 
type UA_WriteResponse C.UA_WriteResponse // line: 3458 
type UA_CallMethodRequest C.UA_CallMethodRequest // line: 3471 
type UA_CallMethodResult C.UA_CallMethodResult // line: 3487 
type UA_CallRequest C.UA_CallRequest // line: 3499 
type UA_CallResponse C.UA_CallResponse // line: 3513 
type UA_MonitoringMode C.UA_MonitoringMode // line: 3525 
const (
    UA_MONITORINGMODE_DISABLED = UA_MonitoringMode(0)
    UA_MONITORINGMODE_SAMPLING = UA_MonitoringMode(1)
    UA_MONITORINGMODE_REPORTING = UA_MonitoringMode(2)
)
type UA_MonitoringParameters C.UA_MonitoringParameters // line: 3539 
type UA_MonitoredItemCreateRequest C.UA_MonitoredItemCreateRequest // line: 3551 
type UA_MonitoredItemCreateResult C.UA_MonitoredItemCreateResult // line: 3565 
type UA_CreateMonitoredItemsRequest C.UA_CreateMonitoredItemsRequest // line: 3579 
type UA_CreateMonitoredItemsResponse C.UA_CreateMonitoredItemsResponse // line: 3593 
type UA_CreateSubscriptionRequest C.UA_CreateSubscriptionRequest // line: 3609 
type UA_CreateSubscriptionResponse C.UA_CreateSubscriptionResponse // line: 3623 
type UA_SetPublishingModeRequest C.UA_SetPublishingModeRequest // line: 3636 
type UA_SetPublishingModeResponse C.UA_SetPublishingModeResponse // line: 3650 
type UA_NotificationMessage C.UA_NotificationMessage // line: 3663 
type UA_SubscriptionAcknowledgement C.UA_SubscriptionAcknowledgement // line: 3674 
type UA_PublishRequest C.UA_PublishRequest // line: 3686 
type UA_PublishResponse C.UA_PublishResponse // line: 3705 
type UA_DeleteSubscriptionsRequest C.UA_DeleteSubscriptionsRequest // line: 3717 
type UA_DeleteSubscriptionsResponse C.UA_DeleteSubscriptionsResponse // line: 3731 
type UA_BuildInfo C.UA_BuildInfo // line: 3746 
type UA_ServerState C.UA_ServerState // line: 3763 
const (
    UA_SERVERSTATE_RUNNING = UA_ServerState(0)
    UA_SERVERSTATE_FAILED = UA_ServerState(1)
    UA_SERVERSTATE_NOCONFIGURATION = UA_ServerState(2)
    UA_SERVERSTATE_SUSPENDED = UA_ServerState(3)
    UA_SERVERSTATE_SHUTDOWN = UA_ServerState(4)
    UA_SERVERSTATE_TEST = UA_ServerState(5)
    UA_SERVERSTATE_COMMUNICATIONFAULT = UA_ServerState(6)
    UA_SERVERSTATE_UNKNOWN = UA_ServerState(7)
)
type UA_ServerStatusDataType C.UA_ServerStatusDataType // line: 3778 
type UA_Connection C.UA_Connection // line: 3816 
type UA_Server C.UA_Server // line: 3819 
type UA_Job C.UA_Job // line: 3844 
type UA_ConnectionState C.UA_ConnectionState // line: 3878 
const (
    UA_CONNECTION_OPENING = iota ///< The socket is open, but the HEL/ACK handshake is not done
    UA_CONNECTION_ESTABLISHED ///< The socket is open and the connection configured
    UA_CONNECTION_CLOSED ///< The socket has been closed and the connection will be deleted
)
type UA_ConnectionConfig C.UA_ConnectionConfig // line: 3886 
type UA_SecureChannel C.UA_SecureChannel // line: 3892 
type UA_LogLevel C.UA_LogLevel // line: 4013 
const (
    UA_LOGLEVEL_TRACE = iota
    UA_LOGLEVEL_DEBUG
    UA_LOGLEVEL_INFO
    UA_LOGLEVEL_WARNING
    UA_LOGLEVEL_ERROR
    UA_LOGLEVEL_FATAL
)
type UA_LogCategory C.UA_LogCategory // line: 4022 
const (
    UA_LOGCATEGORY_NETWORK = iota
    UA_LOGCATEGORY_SECURECHANNEL
    UA_LOGCATEGORY_SESSION
    UA_LOGCATEGORY_SERVER
    UA_LOGCATEGORY_CLIENT
    UA_LOGCATEGORY_USERLAND
)
type UA_ServerNetworkLayer C.UA_ServerNetworkLayer // line: 4103 
type UA_UsernamePasswordLogin C.UA_UsernamePasswordLogin // line: 4108 
type UA_ServerConfig C.UA_ServerConfig // line: 4125 
type UA_DataSource C.UA_DataSource // line: 4266 
type UA_ValueCallback C.UA_ValueCallback // line: 4280 
type UA_ObjectLifecycleManagement C.UA_ObjectLifecycleManagement // line: 4291 
type UA_InstantiationCallback C.UA_InstantiationCallback // line: 4313 
type UA_ExternalNodeStore C.UA_ExternalNodeStore // line: 4689 
type UA_Client C.UA_Client // line: 4709 
type UA_ClientConfig C.UA_ClientConfig // line: 4715 
type UA_SubscriptionSettings C.UA_SubscriptionSettings // line: 5367 
