// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'client.gen.dart';

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

_$_User _$_$_UserFromJson(Map<String, dynamic> json) {
  return _$_User(
    ID: json['id'] as int,
    username: json['USERNAME'] as String,
    role: json['role'] == null
        ? null
        : Kind.fromJson(json['role'] as Map<String, dynamic>),
    meta: json['meta'] as Map<String, dynamic>,
    internalID: json['internalID'] as int,
    createdAt: json['created_at'] == null
        ? null
        : DateTime.parse(json['created_at'] as String),
  );
}

Map<String, dynamic> _$_$_UserToJson(_$_User instance) => <String, dynamic>{
      'id': instance.ID,
      'USERNAME': instance.username,
      'role': instance.role?.toJson(),
      'meta': instance.meta,
      'internalID': instance.internalID,
      'created_at': instance.createdAt?.toIso8601String(),
    };

_$_Page _$_$_PageFromJson(Map<String, dynamic> json) {
  return _$_Page(
    num: json['num'] as int,
  );
}

Map<String, dynamic> _$_$_PageToJson(_$_Page instance) => <String, dynamic>{
      'num': instance.num,
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

_$_FindUsersArgs_Freezed _$_$_FindUsersArgs_FreezedFromJson(
    Map<String, dynamic> json) {
  return _$_FindUsersArgs_Freezed(
    q: json['q'] as String,
  );
}

Map<String, dynamic> _$_$_FindUsersArgs_FreezedToJson(
        _$_FindUsersArgs_Freezed instance) =>
    <String, dynamic>{
      'q': instance.q,
    };

_$_PingResult _$_$_PingResultFromJson(Map<String, dynamic> json) {
  return _$_PingResult(
    status: json['status'] as bool,
  );
}

Map<String, dynamic> _$_$_PingResultToJson(_$_PingResult instance) =>
    <String, dynamic>{
      'status': instance.status,
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

_$_FindUsersResult _$_$_FindUsersResultFromJson(Map<String, dynamic> json) {
  return _$_FindUsersResult(
    page: json['page'] == null
        ? null
        : Page.fromJson(json['page'] as Map<String, dynamic>),
    users: (json['users'] as List)
        ?.map(
            (e) => e == null ? null : User.fromJson(e as Map<String, dynamic>))
        ?.toList(),
  );
}

Map<String, dynamic> _$_$_FindUsersResultToJson(_$_FindUsersResult instance) =>
    <String, dynamic>{
      'page': instance.page?.toJson(),
      'users': instance.users?.map((e) => e?.toJson())?.toList(),
    };
