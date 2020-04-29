// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'example_service.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$USER _$_$USERFromJson(Map<String, dynamic> json) {
  return _$USER();
}

Map<String, dynamic> _$_$USERToJson(_$USER instance) => <String, dynamic>{};

_$ADMIN _$_$ADMINFromJson(Map<String, dynamic> json) {
  return _$ADMIN();
}

Map<String, dynamic> _$_$ADMINToJson(_$ADMIN instance) => <String, dynamic>{};

_$_Empty _$_$_EmptyFromJson(Map<String, dynamic> json) {
  return _$_Empty();
}

Map<String, dynamic> _$_$_EmptyToJson(_$_Empty instance) => <String, dynamic>{};

_$_User _$_$_UserFromJson(Map<String, dynamic> json) {
  return _$_User(
    id: json['id'] as int,
    username: json['USERNAME'] as String,
    role: json['role'] as String,
  );
}

Map<String, dynamic> _$_$_UserToJson(_$_User instance) => <String, dynamic>{
      'id': instance.id,
      'USERNAME': instance.username,
      'role': instance.role,
    };

_$_SearchFilter _$_$_SearchFilterFromJson(Map<String, dynamic> json) {
  return _$_SearchFilter(
    q: json['q'] as String,
  );
}

Map<String, dynamic> _$_$_SearchFilterToJson(_$_SearchFilter instance) =>
    <String, dynamic>{
      'q': instance.q,
    };

_$_Version _$_$_VersionFromJson(Map<String, dynamic> json) {
  return _$_Version(
    webrpcVersion: json['webrpcVersion'] as String,
    schemaVersion: json['schemaVersion'] as String,
    schemaHash: json['schemaHash'] as String,
  );
}

Map<String, dynamic> _$_$_VersionToJson(_$_Version instance) =>
    <String, dynamic>{
      'webrpcVersion': instance.webrpcVersion,
      'schemaVersion': instance.schemaVersion,
      'schemaHash': instance.schemaHash,
    };

_$_ComplexType _$_$_ComplexTypeFromJson(Map<String, dynamic> json) {
  return _$_ComplexType(
    meta: json['meta'] as Map<String, dynamic>,
    metaNestedExample: (json['metaNestedExample'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(
          k,
          (e as Map<String, dynamic>)?.map(
            (k, e) => MapEntry(k, e as int),
          )),
    ),
    namesList: (json['namesList'] as List)?.map((e) => e as String)?.toList(),
    numsList: (json['numsList'] as List)?.map((e) => e as int)?.toList(),
    doubleArray: (json['doubleArray'] as List)
        ?.map((e) => (e as List)?.map((e) => e as String)?.toList())
        ?.toList(),
    listOfUsers: (json['listOfUsers'] as List)
        ?.map(
            (e) => e == null ? null : User.fromJson(e as Map<String, dynamic>))
        ?.toList(),
    mapOfUsers: (json['mapOfUsers'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(
          k, e == null ? null : User.fromJson(e as Map<String, dynamic>)),
    ),
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
    listOfMaps: (json['listOfMaps'] as List)
        ?.map((e) => (e as Map<String, dynamic>)?.map(
              (k, e) => MapEntry(k, e as int),
            ))
        ?.toList(),
  );
}

Map<String, dynamic> _$_$_ComplexTypeToJson(_$_ComplexType instance) =>
    <String, dynamic>{
      'meta': instance.meta,
      'metaNestedExample': instance.metaNestedExample,
      'namesList': instance.namesList,
      'numsList': instance.numsList,
      'doubleArray': instance.doubleArray,
      'listOfUsers': instance.listOfUsers,
      'mapOfUsers': instance.mapOfUsers,
      'user': instance.user,
      'listOfMaps': instance.listOfMaps,
    };

_$_ExampleServiceGetUserArgs_Freezed
    _$_$_ExampleServiceGetUserArgs_FreezedFromJson(Map<String, dynamic> json) {
  return _$_ExampleServiceGetUserArgs_Freezed(
    header: (json['header'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(k, e as String),
    ),
    userID: json['userID'] as int,
  );
}

Map<String, dynamic> _$_$_ExampleServiceGetUserArgs_FreezedToJson(
        _$_ExampleServiceGetUserArgs_Freezed instance) =>
    <String, dynamic>{
      'header': instance.header,
      'userID': instance.userID,
    };

_$_ExampleServiceFindUserArgs_Freezed
    _$_$_ExampleServiceFindUserArgs_FreezedFromJson(Map<String, dynamic> json) {
  return _$_ExampleServiceFindUserArgs_Freezed(
    s: json['s'] == null
        ? null
        : SearchFilter.fromJson(json['s'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_ExampleServiceFindUserArgs_FreezedToJson(
        _$_ExampleServiceFindUserArgs_Freezed instance) =>
    <String, dynamic>{
      's': instance.s,
    };

_$_StatusResponse _$_$_StatusResponseFromJson(Map<String, dynamic> json) {
  return _$_StatusResponse(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_StatusResponseToJson(_$_StatusResponse instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_StatusResponseRpcErr _$_$_StatusResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_StatusResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_StatusResponseRpcErrToJson(
        _$_StatusResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_VersionResponse _$_$_VersionResponseFromJson(Map<String, dynamic> json) {
  return _$_VersionResponse(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_VersionResponseToJson(_$_VersionResponse instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_VersionResponseRpcErr _$_$_VersionResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_VersionResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_VersionResponseRpcErrToJson(
        _$_VersionResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_GetUserResponse _$_$_GetUserResponseFromJson(Map<String, dynamic> json) {
  return _$_GetUserResponse(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetUserResponseToJson(_$_GetUserResponse instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_GetUserResponseRpcErr _$_$_GetUserResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_GetUserResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_GetUserResponseRpcErrToJson(
        _$_GetUserResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_FindUserResponse _$_$_FindUserResponseFromJson(Map<String, dynamic> json) {
  return _$_FindUserResponse(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_FindUserResponseToJson(
        _$_FindUserResponse instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };

_$_FindUserResponseRpcErr _$_$_FindUserResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_FindUserResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_FindUserResponseRpcErrToJson(
        _$_FindUserResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_AnotherExampleServiceGetAccountArgs_Freezed
    _$_$_AnotherExampleServiceGetAccountArgs_FreezedFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceGetAccountArgs_Freezed(
    header: (json['header'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(k, e as String),
    ),
    userID: json['userID'] as int,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceGetAccountArgs_FreezedToJson(
        _$_AnotherExampleServiceGetAccountArgs_Freezed instance) =>
    <String, dynamic>{
      'header': instance.header,
      'userID': instance.userID,
    };

_$_AnotherExampleServiceGetUsersArgs_Freezed
    _$_$_AnotherExampleServiceGetUsersArgs_FreezedFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceGetUsersArgs_Freezed(
    s: json['s'] == null
        ? null
        : SearchFilter.fromJson(json['s'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceGetUsersArgs_FreezedToJson(
        _$_AnotherExampleServiceGetUsersArgs_Freezed instance) =>
    <String, dynamic>{
      's': instance.s,
    };

_$_AnotherExampleServiceStatusResponse
    _$_$_AnotherExampleServiceStatusResponseFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusResponse(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusResponseToJson(
        _$_AnotherExampleServiceStatusResponse instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_AnotherExampleServiceStatusResponseRpcErr
    _$_$_AnotherExampleServiceStatusResponseRpcErrFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusResponseRpcErrToJson(
        _$_AnotherExampleServiceStatusResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_GetVersionResponse _$_$_GetVersionResponseFromJson(
    Map<String, dynamic> json) {
  return _$_GetVersionResponse(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetVersionResponseToJson(
        _$_GetVersionResponse instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_GetVersionResponseRpcErr _$_$_GetVersionResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_GetVersionResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_GetVersionResponseRpcErrToJson(
        _$_GetVersionResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_GetAccountResponse _$_$_GetAccountResponseFromJson(
    Map<String, dynamic> json) {
  return _$_GetAccountResponse(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetAccountResponseToJson(
        _$_GetAccountResponse instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_GetAccountResponseRpcErr _$_$_GetAccountResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_GetAccountResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_GetAccountResponseRpcErrToJson(
        _$_GetAccountResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };

_$_GetUsersResponse _$_$_GetUsersResponseFromJson(Map<String, dynamic> json) {
  return _$_GetUsersResponse(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetUsersResponseToJson(
        _$_GetUsersResponse instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };

_$_GetUsersResponseRpcErr _$_$_GetUsersResponseRpcErrFromJson(
    Map<String, dynamic> json) {
  return _$_GetUsersResponseRpcErr(
    message: json['message'] as String,
    route: json['route'] as String,
    statusCode: json['statusCode'] as int,
    timeStamp: json['timeStamp'] == null
        ? null
        : DateTime.parse(json['timeStamp'] as String),
  );
}

Map<String, dynamic> _$_$_GetUsersResponseRpcErrToJson(
        _$_GetUsersResponseRpcErr instance) =>
    <String, dynamic>{
      'message': instance.message,
      'route': instance.route,
      'statusCode': instance.statusCode,
      'timeStamp': instance.timeStamp?.toIso8601String(),
    };
