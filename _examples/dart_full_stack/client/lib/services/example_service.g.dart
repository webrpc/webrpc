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

_$_StatusResult _$_$_StatusResultFromJson(Map<String, dynamic> json) {
  return _$_StatusResult(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_StatusResultToJson(_$_StatusResult instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_StatusResultFailed _$_$_StatusResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_StatusResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_StatusResultFailedToJson(
        _$_StatusResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_StatusResultPending _$_$_StatusResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_StatusResultPending();
}

Map<String, dynamic> _$_$_StatusResultPendingToJson(
        _$_StatusResultPending instance) =>
    <String, dynamic>{};

_$_VersionResult _$_$_VersionResultFromJson(Map<String, dynamic> json) {
  return _$_VersionResult(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_VersionResultToJson(_$_VersionResult instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_VersionResultFailed _$_$_VersionResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_VersionResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_VersionResultFailedToJson(
        _$_VersionResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_VersionResultPending _$_$_VersionResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_VersionResultPending();
}

Map<String, dynamic> _$_$_VersionResultPendingToJson(
        _$_VersionResultPending instance) =>
    <String, dynamic>{};

_$_GetUserResult _$_$_GetUserResultFromJson(Map<String, dynamic> json) {
  return _$_GetUserResult(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetUserResultToJson(_$_GetUserResult instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_GetUserResultFailed _$_$_GetUserResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_GetUserResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_GetUserResultFailedToJson(
        _$_GetUserResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_GetUserResultPending _$_$_GetUserResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_GetUserResultPending();
}

Map<String, dynamic> _$_$_GetUserResultPendingToJson(
        _$_GetUserResultPending instance) =>
    <String, dynamic>{};

_$_FindUserResult _$_$_FindUserResultFromJson(Map<String, dynamic> json) {
  return _$_FindUserResult(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_FindUserResultToJson(_$_FindUserResult instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };

_$_FindUserResultFailed _$_$_FindUserResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_FindUserResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_FindUserResultFailedToJson(
        _$_FindUserResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_FindUserResultPending _$_$_FindUserResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_FindUserResultPending();
}

Map<String, dynamic> _$_$_FindUserResultPendingToJson(
        _$_FindUserResultPending instance) =>
    <String, dynamic>{};

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

_$_AnotherExampleServiceStatusResult
    _$_$_AnotherExampleServiceStatusResultFromJson(Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusResult(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusResultToJson(
        _$_AnotherExampleServiceStatusResult instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_AnotherExampleServiceStatusResultFailed
    _$_$_AnotherExampleServiceStatusResultFailedFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusResultFailedToJson(
        _$_AnotherExampleServiceStatusResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_AnotherExampleServiceStatusResultPending
    _$_$_AnotherExampleServiceStatusResultPendingFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusResultPending();
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusResultPendingToJson(
        _$_AnotherExampleServiceStatusResultPending instance) =>
    <String, dynamic>{};

_$_GetVersionResult _$_$_GetVersionResultFromJson(Map<String, dynamic> json) {
  return _$_GetVersionResult(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetVersionResultToJson(
        _$_GetVersionResult instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_GetVersionResultFailed _$_$_GetVersionResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_GetVersionResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_GetVersionResultFailedToJson(
        _$_GetVersionResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_GetVersionResultPending _$_$_GetVersionResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_GetVersionResultPending();
}

Map<String, dynamic> _$_$_GetVersionResultPendingToJson(
        _$_GetVersionResultPending instance) =>
    <String, dynamic>{};

_$_GetAccountResult _$_$_GetAccountResultFromJson(Map<String, dynamic> json) {
  return _$_GetAccountResult(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetAccountResultToJson(
        _$_GetAccountResult instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_GetAccountResultFailed _$_$_GetAccountResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_GetAccountResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_GetAccountResultFailedToJson(
        _$_GetAccountResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_GetAccountResultPending _$_$_GetAccountResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_GetAccountResultPending();
}

Map<String, dynamic> _$_$_GetAccountResultPendingToJson(
        _$_GetAccountResultPending instance) =>
    <String, dynamic>{};

_$_GetUsersResult _$_$_GetUsersResultFromJson(Map<String, dynamic> json) {
  return _$_GetUsersResult(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetUsersResultToJson(_$_GetUsersResult instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };

_$_GetUsersResultFailed _$_$_GetUsersResultFailedFromJson(
    Map<String, dynamic> json) {
  return _$_GetUsersResultFailed(
    statusCode: json['statusCode'] as int,
    reason: json['reason'] as String,
    stackTrace: json['stackTrace'] as String,
  );
}

Map<String, dynamic> _$_$_GetUsersResultFailedToJson(
        _$_GetUsersResultFailed instance) =>
    <String, dynamic>{
      'statusCode': instance.statusCode,
      'reason': instance.reason,
      'stackTrace': instance.stackTrace,
    };

_$_GetUsersResultPending _$_$_GetUsersResultPendingFromJson(
    Map<String, dynamic> json) {
  return _$_GetUsersResultPending();
}

Map<String, dynamic> _$_$_GetUsersResultPendingToJson(
        _$_GetUsersResultPending instance) =>
    <String, dynamic>{};
