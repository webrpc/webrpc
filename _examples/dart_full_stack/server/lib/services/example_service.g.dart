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

_$_ExampleServiceStatusReturn _$_$_ExampleServiceStatusReturnFromJson(
    Map<String, dynamic> json) {
  return _$_ExampleServiceStatusReturn(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_ExampleServiceStatusReturnToJson(
        _$_ExampleServiceStatusReturn instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_ExampleServiceVersionReturn _$_$_ExampleServiceVersionReturnFromJson(
    Map<String, dynamic> json) {
  return _$_ExampleServiceVersionReturn(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_ExampleServiceVersionReturnToJson(
        _$_ExampleServiceVersionReturn instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_ExampleServiceGetUserReturn _$_$_ExampleServiceGetUserReturnFromJson(
    Map<String, dynamic> json) {
  return _$_ExampleServiceGetUserReturn(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_ExampleServiceGetUserReturnToJson(
        _$_ExampleServiceGetUserReturn instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_ExampleServiceFindUserReturn _$_$_ExampleServiceFindUserReturnFromJson(
    Map<String, dynamic> json) {
  return _$_ExampleServiceFindUserReturn(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_ExampleServiceFindUserReturnToJson(
        _$_ExampleServiceFindUserReturn instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };

_$_AnotherExampleServiceGetUserArgs_Freezed
    _$_$_AnotherExampleServiceGetUserArgs_FreezedFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceGetUserArgs_Freezed(
    header: (json['header'] as Map<String, dynamic>)?.map(
      (k, e) => MapEntry(k, e as String),
    ),
    userID: json['userID'] as int,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceGetUserArgs_FreezedToJson(
        _$_AnotherExampleServiceGetUserArgs_Freezed instance) =>
    <String, dynamic>{
      'header': instance.header,
      'userID': instance.userID,
    };

_$_AnotherExampleServiceFindUserArgs_Freezed
    _$_$_AnotherExampleServiceFindUserArgs_FreezedFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceFindUserArgs_Freezed(
    s: json['s'] == null
        ? null
        : SearchFilter.fromJson(json['s'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceFindUserArgs_FreezedToJson(
        _$_AnotherExampleServiceFindUserArgs_Freezed instance) =>
    <String, dynamic>{
      's': instance.s,
    };

_$_AnotherExampleServiceStatusReturn
    _$_$_AnotherExampleServiceStatusReturnFromJson(Map<String, dynamic> json) {
  return _$_AnotherExampleServiceStatusReturn(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceStatusReturnToJson(
        _$_AnotherExampleServiceStatusReturn instance) =>
    <String, dynamic>{
      'status': instance.status,
    };

_$_AnotherExampleServiceVersionReturn
    _$_$_AnotherExampleServiceVersionReturnFromJson(Map<String, dynamic> json) {
  return _$_AnotherExampleServiceVersionReturn(
    version: json['version'] == null
        ? null
        : Version.fromJson(json['version'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceVersionReturnToJson(
        _$_AnotherExampleServiceVersionReturn instance) =>
    <String, dynamic>{
      'version': instance.version,
    };

_$_AnotherExampleServiceGetUserReturn
    _$_$_AnotherExampleServiceGetUserReturnFromJson(Map<String, dynamic> json) {
  return _$_AnotherExampleServiceGetUserReturn(
    code: json['code'] as int,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceGetUserReturnToJson(
        _$_AnotherExampleServiceGetUserReturn instance) =>
    <String, dynamic>{
      'code': instance.code,
      'user': instance.user,
    };

_$_AnotherExampleServiceFindUserReturn
    _$_$_AnotherExampleServiceFindUserReturnFromJson(
        Map<String, dynamic> json) {
  return _$_AnotherExampleServiceFindUserReturn(
    name: json['name'] as String,
    user: json['user'] == null
        ? null
        : User.fromJson(json['user'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$_$_AnotherExampleServiceFindUserReturnToJson(
        _$_AnotherExampleServiceFindUserReturn instance) =>
    <String, dynamic>{
      'name': instance.name,
      'user': instance.user,
    };
