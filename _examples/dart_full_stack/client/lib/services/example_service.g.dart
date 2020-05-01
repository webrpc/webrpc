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
    id: json['id'] as int,
  );
}

Map<String, dynamic> _$_$_SearchFilterToJson(_$_SearchFilter instance) =>
    <String, dynamic>{
      'id': instance.id,
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
      'listOfUsers': instance.listOfUsers?.map((e) => e?.toJson())?.toList(),
      'mapOfUsers':
          instance.mapOfUsers?.map((k, e) => MapEntry(k, e?.toJson())),
      'user': instance.user?.toJson(),
      'listOfMaps': instance.listOfMaps,
    };

_$_GetUserArgs_Freezed _$_$_GetUserArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_GetUserArgs_Freezed(
    userID: json['userID'] as int,
  );
}

Map<String, dynamic> _$_$_GetUserArgs_FreezedToJson(
        _$_GetUserArgs_Freezed instance) =>
    <String, dynamic>{
      'userID': instance.userID,
    };

_$_UpdateNameArgs_Freezed _$_$_UpdateNameArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_UpdateNameArgs_Freezed(
    id: json['id'] as int,
    username: json['username'] as String,
  );
}

Map<String, dynamic> _$_$_UpdateNameArgs_FreezedToJson(
        _$_UpdateNameArgs_Freezed instance) =>
    <String, dynamic>{
      'id': instance.id,
      'username': instance.username,
    };

_$_FindUserByIdArgs_Freezed _$_$_FindUserByIdArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_FindUserByIdArgs_Freezed(
    s: json['s'] == null
        ? null
        : SearchFilter.fromJson(json['s'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_FindUserByIdArgs_FreezedToJson(
        _$_FindUserByIdArgs_Freezed instance) =>
    <String, dynamic>{
      's': instance.s?.toJson(),
    };

_$_AddUserArgs_Freezed _$_$_AddUserArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_AddUserArgs_Freezed(
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AddUserArgs_FreezedToJson(
        _$_AddUserArgs_Freezed instance) =>
    <String, dynamic>{
      'user': instance.user?.toJson(),
    };

_$_DeleteUserArgs_Freezed _$_$_DeleteUserArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_DeleteUserArgs_Freezed(
    id: json['id'] as int,
  );
}

Map<String, dynamic> _$_$_DeleteUserArgs_FreezedToJson(
        _$_DeleteUserArgs_Freezed instance) =>
    <String, dynamic>{
      'id': instance.id,
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

_$_VersionResult _$_$_VersionResultFromJson(Map<String, dynamic> json) {
  return _$_VersionResult(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_VersionResultToJson(_$_VersionResult instance) =>
    <String, dynamic>{
      'version': instance.version?.toJson(),
    };

_$_GetUserResult _$_$_GetUserResultFromJson(Map<String, dynamic> json) {
  return _$_GetUserResult(
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_GetUserResultToJson(_$_GetUserResult instance) =>
    <String, dynamic>{
      'user': instance.user?.toJson(),
    };

_$_UpdateNameResult _$_$_UpdateNameResultFromJson(Map<String, dynamic> json) {
  return _$_UpdateNameResult(
    wasUpdated: json['wasUpdated'] as bool,
  );
}

Map<String, dynamic> _$_$_UpdateNameResultToJson(
        _$_UpdateNameResult instance) =>
    <String, dynamic>{
      'wasUpdated': instance.wasUpdated,
    };

_$_FindUserByIdResult _$_$_FindUserByIdResultFromJson(
    Map<String, dynamic> json) {
  return _$_FindUserByIdResult(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_FindUserByIdResultToJson(
        _$_FindUserByIdResult instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user?.toJson(),
    };

_$_AddUserResult _$_$_AddUserResultFromJson(Map<String, dynamic> json) {
  return _$_AddUserResult(
    wasAdded: json['wasAdded'] as bool,
  );
}

Map<String, dynamic> _$_$_AddUserResultToJson(_$_AddUserResult instance) =>
    <String, dynamic>{
      'wasAdded': instance.wasAdded,
    };

_$_ListUsersResult _$_$_ListUsersResultFromJson(Map<String, dynamic> json) {
  return _$_ListUsersResult(
    users: (json['users'] as List)
        ?.map(
            (e) => e == null ? null : User.fromJson(e as Map<String, dynamic>))
        ?.toList(),
  );
}

Map<String, dynamic> _$_$_ListUsersResultToJson(_$_ListUsersResult instance) =>
    <String, dynamic>{
      'users': instance.users?.map((e) => e?.toJson())?.toList(),
    };

_$_DeleteUserResult _$_$_DeleteUserResultFromJson(Map<String, dynamic> json) {
  return _$_DeleteUserResult(
    wasDeleted: json['wasDeleted'] as bool,
  );
}

Map<String, dynamic> _$_$_DeleteUserResultToJson(
        _$_DeleteUserResult instance) =>
    <String, dynamic>{
      'wasDeleted': instance.wasDeleted,
    };
