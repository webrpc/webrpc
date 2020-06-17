// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named

part of 'hello_api.gen.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;
Kind _$KindFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'user':
      return USER.fromJson(json);
    case 'admin':
      return ADMIN.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$KindTearOff {
  const _$KindTearOff();

  USER user() {
    return const USER();
  }

  ADMIN admin() {
    return const ADMIN();
  }
}

// ignore: unused_element
const $Kind = _$KindTearOff();

mixin _$Kind {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result user(),
    @required Result admin(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result user(),
    Result admin(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result user(USER value),
    @required Result admin(ADMIN value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result user(USER value),
    Result admin(ADMIN value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $KindCopyWith<$Res> {
  factory $KindCopyWith(Kind value, $Res Function(Kind) then) =
      _$KindCopyWithImpl<$Res>;
}

class _$KindCopyWithImpl<$Res> implements $KindCopyWith<$Res> {
  _$KindCopyWithImpl(this._value, this._then);

  final Kind _value;
  // ignore: unused_field
  final $Res Function(Kind) _then;
}

abstract class $USERCopyWith<$Res> {
  factory $USERCopyWith(USER value, $Res Function(USER) then) =
      _$USERCopyWithImpl<$Res>;
}

class _$USERCopyWithImpl<$Res> extends _$KindCopyWithImpl<$Res>
    implements $USERCopyWith<$Res> {
  _$USERCopyWithImpl(USER _value, $Res Function(USER) _then)
      : super(_value, (v) => _then(v as USER));

  @override
  USER get _value => super._value as USER;
}

@JsonSerializable()
class _$USER implements USER {
  const _$USER();

  factory _$USER.fromJson(Map<String, dynamic> json) => _$_$USERFromJson(json);

  @override
  String toString() {
    return 'Kind.user()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is USER);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result user(),
    @required Result admin(),
  }) {
    assert(user != null);
    assert(admin != null);
    return user();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result user(),
    Result admin(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (user != null) {
      return user();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result user(USER value),
    @required Result admin(ADMIN value),
  }) {
    assert(user != null);
    assert(admin != null);
    return user(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result user(USER value),
    Result admin(ADMIN value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (user != null) {
      return user(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$USERToJson(this)..['runtimeType'] = 'user';
  }
}

abstract class USER implements Kind {
  const factory USER() = _$USER;

  factory USER.fromJson(Map<String, dynamic> json) = _$USER.fromJson;
}

abstract class $ADMINCopyWith<$Res> {
  factory $ADMINCopyWith(ADMIN value, $Res Function(ADMIN) then) =
      _$ADMINCopyWithImpl<$Res>;
}

class _$ADMINCopyWithImpl<$Res> extends _$KindCopyWithImpl<$Res>
    implements $ADMINCopyWith<$Res> {
  _$ADMINCopyWithImpl(ADMIN _value, $Res Function(ADMIN) _then)
      : super(_value, (v) => _then(v as ADMIN));

  @override
  ADMIN get _value => super._value as ADMIN;
}

@JsonSerializable()
class _$ADMIN implements ADMIN {
  const _$ADMIN();

  factory _$ADMIN.fromJson(Map<String, dynamic> json) =>
      _$_$ADMINFromJson(json);

  @override
  String toString() {
    return 'Kind.admin()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is ADMIN);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result user(),
    @required Result admin(),
  }) {
    assert(user != null);
    assert(admin != null);
    return admin();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result user(),
    Result admin(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (admin != null) {
      return admin();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result user(USER value),
    @required Result admin(ADMIN value),
  }) {
    assert(user != null);
    assert(admin != null);
    return admin(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result user(USER value),
    Result admin(ADMIN value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (admin != null) {
      return admin(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$ADMINToJson(this)..['runtimeType'] = 'admin';
  }
}

abstract class ADMIN implements Kind {
  const factory ADMIN() = _$ADMIN;

  factory ADMIN.fromJson(Map<String, dynamic> json) = _$ADMIN.fromJson;
}

User _$UserFromJson(Map<String, dynamic> json) {
  return _User.fromJson(json);
}

class _$UserTearOff {
  const _$UserTearOff();

  _User call(
      {@required @JsonKey(name: 'id') int ID,
      @required @JsonKey(name: 'USERNAME') String username,
      @required Kind role,
      @required Map<String, dynamic> meta,
      @required int internalID,
      @JsonKey(name: 'created_at') DateTime createdAt}) {
    return _User(
      ID: ID,
      username: username,
      role: role,
      meta: meta,
      internalID: internalID,
      createdAt: createdAt,
    );
  }
}

// ignore: unused_element
const $User = _$UserTearOff();

mixin _$User {
  @JsonKey(name: 'id')
  int get ID;
  @JsonKey(name: 'USERNAME')
  String get username;
  Kind get role;
  Map<String, dynamic> get meta;
  int get internalID;
  @JsonKey(name: 'created_at')
  DateTime get createdAt;

  Map<String, dynamic> toJson();
  $UserCopyWith<User> get copyWith;
}

abstract class $UserCopyWith<$Res> {
  factory $UserCopyWith(User value, $Res Function(User) then) =
      _$UserCopyWithImpl<$Res>;
  $Res call(
      {@JsonKey(name: 'id') int ID,
      @JsonKey(name: 'USERNAME') String username,
      Kind role,
      Map<String, dynamic> meta,
      int internalID,
      @JsonKey(name: 'created_at') DateTime createdAt});

  $KindCopyWith<$Res> get role;
}

class _$UserCopyWithImpl<$Res> implements $UserCopyWith<$Res> {
  _$UserCopyWithImpl(this._value, this._then);

  final User _value;
  // ignore: unused_field
  final $Res Function(User) _then;

  @override
  $Res call({
    Object ID = freezed,
    Object username = freezed,
    Object role = freezed,
    Object meta = freezed,
    Object internalID = freezed,
    Object createdAt = freezed,
  }) {
    return _then(_value.copyWith(
      ID: ID == freezed ? _value.ID : ID as int,
      username: username == freezed ? _value.username : username as String,
      role: role == freezed ? _value.role : role as Kind,
      meta: meta == freezed ? _value.meta : meta as Map<String, dynamic>,
      internalID: internalID == freezed ? _value.internalID : internalID as int,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
    ));
  }

  @override
  $KindCopyWith<$Res> get role {
    if (_value.role == null) {
      return null;
    }
    return $KindCopyWith<$Res>(_value.role, (value) {
      return _then(_value.copyWith(role: value));
    });
  }
}

abstract class _$UserCopyWith<$Res> implements $UserCopyWith<$Res> {
  factory _$UserCopyWith(_User value, $Res Function(_User) then) =
      __$UserCopyWithImpl<$Res>;
  @override
  $Res call(
      {@JsonKey(name: 'id') int ID,
      @JsonKey(name: 'USERNAME') String username,
      Kind role,
      Map<String, dynamic> meta,
      int internalID,
      @JsonKey(name: 'created_at') DateTime createdAt});

  @override
  $KindCopyWith<$Res> get role;
}

class __$UserCopyWithImpl<$Res> extends _$UserCopyWithImpl<$Res>
    implements _$UserCopyWith<$Res> {
  __$UserCopyWithImpl(_User _value, $Res Function(_User) _then)
      : super(_value, (v) => _then(v as _User));

  @override
  _User get _value => super._value as _User;

  @override
  $Res call({
    Object ID = freezed,
    Object username = freezed,
    Object role = freezed,
    Object meta = freezed,
    Object internalID = freezed,
    Object createdAt = freezed,
  }) {
    return _then(_User(
      ID: ID == freezed ? _value.ID : ID as int,
      username: username == freezed ? _value.username : username as String,
      role: role == freezed ? _value.role : role as Kind,
      meta: meta == freezed ? _value.meta : meta as Map<String, dynamic>,
      internalID: internalID == freezed ? _value.internalID : internalID as int,
      createdAt:
          createdAt == freezed ? _value.createdAt : createdAt as DateTime,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_User implements _User {
  _$_User(
      {@required @JsonKey(name: 'id') this.ID,
      @required @JsonKey(name: 'USERNAME') this.username,
      @required this.role,
      @required this.meta,
      @required this.internalID,
      @JsonKey(name: 'created_at') this.createdAt})
      : assert(ID != null),
        assert(username != null),
        assert(role != null),
        assert(meta != null),
        assert(internalID != null);

  factory _$_User.fromJson(Map<String, dynamic> json) =>
      _$_$_UserFromJson(json);

  @override
  @JsonKey(name: 'id')
  final int ID;
  @override
  @JsonKey(name: 'USERNAME')
  final String username;
  @override
  final Kind role;
  @override
  final Map<String, dynamic> meta;
  @override
  final int internalID;
  @override
  @JsonKey(name: 'created_at')
  final DateTime createdAt;

  @override
  String toString() {
    return 'User(ID: $ID, username: $username, role: $role, meta: $meta, internalID: $internalID, createdAt: $createdAt)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _User &&
            (identical(other.ID, ID) ||
                const DeepCollectionEquality().equals(other.ID, ID)) &&
            (identical(other.username, username) ||
                const DeepCollectionEquality()
                    .equals(other.username, username)) &&
            (identical(other.role, role) ||
                const DeepCollectionEquality().equals(other.role, role)) &&
            (identical(other.meta, meta) ||
                const DeepCollectionEquality().equals(other.meta, meta)) &&
            (identical(other.internalID, internalID) ||
                const DeepCollectionEquality()
                    .equals(other.internalID, internalID)) &&
            (identical(other.createdAt, createdAt) ||
                const DeepCollectionEquality()
                    .equals(other.createdAt, createdAt)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(ID) ^
      const DeepCollectionEquality().hash(username) ^
      const DeepCollectionEquality().hash(role) ^
      const DeepCollectionEquality().hash(meta) ^
      const DeepCollectionEquality().hash(internalID) ^
      const DeepCollectionEquality().hash(createdAt);

  @override
  _$UserCopyWith<_User> get copyWith =>
      __$UserCopyWithImpl<_User>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_UserToJson(this);
  }
}

abstract class _User implements User {
  factory _User(
      {@required @JsonKey(name: 'id') int ID,
      @required @JsonKey(name: 'USERNAME') String username,
      @required Kind role,
      @required Map<String, dynamic> meta,
      @required int internalID,
      @JsonKey(name: 'created_at') DateTime createdAt}) = _$_User;

  factory _User.fromJson(Map<String, dynamic> json) = _$_User.fromJson;

  @override
  @JsonKey(name: 'id')
  int get ID;
  @override
  @JsonKey(name: 'USERNAME')
  String get username;
  @override
  Kind get role;
  @override
  Map<String, dynamic> get meta;
  @override
  int get internalID;
  @override
  @JsonKey(name: 'created_at')
  DateTime get createdAt;
  @override
  _$UserCopyWith<_User> get copyWith;
}

Page _$PageFromJson(Map<String, dynamic> json) {
  return _Page.fromJson(json);
}

class _$PageTearOff {
  const _$PageTearOff();

  _Page call({@required int num}) {
    return _Page(
      num: num,
    );
  }
}

// ignore: unused_element
const $Page = _$PageTearOff();

mixin _$Page {
  int get num;

  Map<String, dynamic> toJson();
  $PageCopyWith<Page> get copyWith;
}

abstract class $PageCopyWith<$Res> {
  factory $PageCopyWith(Page value, $Res Function(Page) then) =
      _$PageCopyWithImpl<$Res>;
  $Res call({int num});
}

class _$PageCopyWithImpl<$Res> implements $PageCopyWith<$Res> {
  _$PageCopyWithImpl(this._value, this._then);

  final Page _value;
  // ignore: unused_field
  final $Res Function(Page) _then;

  @override
  $Res call({
    Object num = freezed,
  }) {
    return _then(_value.copyWith(
      num: num == freezed ? _value.num : num as int,
    ));
  }
}

abstract class _$PageCopyWith<$Res> implements $PageCopyWith<$Res> {
  factory _$PageCopyWith(_Page value, $Res Function(_Page) then) =
      __$PageCopyWithImpl<$Res>;
  @override
  $Res call({int num});
}

class __$PageCopyWithImpl<$Res> extends _$PageCopyWithImpl<$Res>
    implements _$PageCopyWith<$Res> {
  __$PageCopyWithImpl(_Page _value, $Res Function(_Page) _then)
      : super(_value, (v) => _then(v as _Page));

  @override
  _Page get _value => super._value as _Page;

  @override
  $Res call({
    Object num = freezed,
  }) {
    return _then(_Page(
      num: num == freezed ? _value.num : num as int,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_Page implements _Page {
  _$_Page({@required this.num}) : assert(num != null);

  factory _$_Page.fromJson(Map<String, dynamic> json) =>
      _$_$_PageFromJson(json);

  @override
  final int num;

  @override
  String toString() {
    return 'Page(num: $num)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _Page &&
            (identical(other.num, num) ||
                const DeepCollectionEquality().equals(other.num, num)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(num);

  @override
  _$PageCopyWith<_Page> get copyWith =>
      __$PageCopyWithImpl<_Page>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_PageToJson(this);
  }
}

abstract class _Page implements Page {
  factory _Page({@required int num}) = _$_Page;

  factory _Page.fromJson(Map<String, dynamic> json) = _$_Page.fromJson;

  @override
  int get num;
  @override
  _$PageCopyWith<_Page> get copyWith;
}

_GetUserArgs _$_GetUserArgsFromJson(Map<String, dynamic> json) {
  return _GetUserArgs_Freezed.fromJson(json);
}

class _$_GetUserArgsTearOff {
  const _$_GetUserArgsTearOff();

  _GetUserArgs_Freezed call({@required int userID}) {
    return _GetUserArgs_Freezed(
      userID: userID,
    );
  }
}

// ignore: unused_element
const _$GetUserArgs = _$_GetUserArgsTearOff();

mixin _$_GetUserArgs {
  int get userID;

  Map<String, dynamic> toJson();
  _$GetUserArgsCopyWith<_GetUserArgs> get copyWith;
}

abstract class _$GetUserArgsCopyWith<$Res> {
  factory _$GetUserArgsCopyWith(
          _GetUserArgs value, $Res Function(_GetUserArgs) then) =
      __$GetUserArgsCopyWithImpl<$Res>;
  $Res call({int userID});
}

class __$GetUserArgsCopyWithImpl<$Res> implements _$GetUserArgsCopyWith<$Res> {
  __$GetUserArgsCopyWithImpl(this._value, this._then);

  final _GetUserArgs _value;
  // ignore: unused_field
  final $Res Function(_GetUserArgs) _then;

  @override
  $Res call({
    Object userID = freezed,
  }) {
    return _then(_value.copyWith(
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

abstract class _$GetUserArgs_FreezedCopyWith<$Res>
    implements _$GetUserArgsCopyWith<$Res> {
  factory _$GetUserArgs_FreezedCopyWith(_GetUserArgs_Freezed value,
          $Res Function(_GetUserArgs_Freezed) then) =
      __$GetUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({int userID});
}

class __$GetUserArgs_FreezedCopyWithImpl<$Res>
    extends __$GetUserArgsCopyWithImpl<$Res>
    implements _$GetUserArgs_FreezedCopyWith<$Res> {
  __$GetUserArgs_FreezedCopyWithImpl(
      _GetUserArgs_Freezed _value, $Res Function(_GetUserArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _GetUserArgs_Freezed));

  @override
  _GetUserArgs_Freezed get _value => super._value as _GetUserArgs_Freezed;

  @override
  $Res call({
    Object userID = freezed,
  }) {
    return _then(_GetUserArgs_Freezed(
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetUserArgs_Freezed implements _GetUserArgs_Freezed {
  _$_GetUserArgs_Freezed({@required this.userID}) : assert(userID != null);

  factory _$_GetUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserArgs_FreezedFromJson(json);

  @override
  final int userID;

  @override
  String toString() {
    return '_GetUserArgs(userID: $userID)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserArgs_Freezed &&
            (identical(other.userID, userID) ||
                const DeepCollectionEquality().equals(other.userID, userID)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(userID);

  @override
  _$GetUserArgs_FreezedCopyWith<_GetUserArgs_Freezed> get copyWith =>
      __$GetUserArgs_FreezedCopyWithImpl<_GetUserArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserArgs_FreezedToJson(this);
  }
}

abstract class _GetUserArgs_Freezed implements _GetUserArgs {
  factory _GetUserArgs_Freezed({@required int userID}) = _$_GetUserArgs_Freezed;

  factory _GetUserArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_GetUserArgs_Freezed.fromJson;

  @override
  int get userID;
  @override
  _$GetUserArgs_FreezedCopyWith<_GetUserArgs_Freezed> get copyWith;
}

_FindUsersArgs _$_FindUsersArgsFromJson(Map<String, dynamic> json) {
  return _FindUsersArgs_Freezed.fromJson(json);
}

class _$_FindUsersArgsTearOff {
  const _$_FindUsersArgsTearOff();

  _FindUsersArgs_Freezed call({@required String q}) {
    return _FindUsersArgs_Freezed(
      q: q,
    );
  }
}

// ignore: unused_element
const _$FindUsersArgs = _$_FindUsersArgsTearOff();

mixin _$_FindUsersArgs {
  String get q;

  Map<String, dynamic> toJson();
  _$FindUsersArgsCopyWith<_FindUsersArgs> get copyWith;
}

abstract class _$FindUsersArgsCopyWith<$Res> {
  factory _$FindUsersArgsCopyWith(
          _FindUsersArgs value, $Res Function(_FindUsersArgs) then) =
      __$FindUsersArgsCopyWithImpl<$Res>;
  $Res call({String q});
}

class __$FindUsersArgsCopyWithImpl<$Res>
    implements _$FindUsersArgsCopyWith<$Res> {
  __$FindUsersArgsCopyWithImpl(this._value, this._then);

  final _FindUsersArgs _value;
  // ignore: unused_field
  final $Res Function(_FindUsersArgs) _then;

  @override
  $Res call({
    Object q = freezed,
  }) {
    return _then(_value.copyWith(
      q: q == freezed ? _value.q : q as String,
    ));
  }
}

abstract class _$FindUsersArgs_FreezedCopyWith<$Res>
    implements _$FindUsersArgsCopyWith<$Res> {
  factory _$FindUsersArgs_FreezedCopyWith(_FindUsersArgs_Freezed value,
          $Res Function(_FindUsersArgs_Freezed) then) =
      __$FindUsersArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({String q});
}

class __$FindUsersArgs_FreezedCopyWithImpl<$Res>
    extends __$FindUsersArgsCopyWithImpl<$Res>
    implements _$FindUsersArgs_FreezedCopyWith<$Res> {
  __$FindUsersArgs_FreezedCopyWithImpl(_FindUsersArgs_Freezed _value,
      $Res Function(_FindUsersArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _FindUsersArgs_Freezed));

  @override
  _FindUsersArgs_Freezed get _value => super._value as _FindUsersArgs_Freezed;

  @override
  $Res call({
    Object q = freezed,
  }) {
    return _then(_FindUsersArgs_Freezed(
      q: q == freezed ? _value.q : q as String,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_FindUsersArgs_Freezed implements _FindUsersArgs_Freezed {
  _$_FindUsersArgs_Freezed({@required this.q}) : assert(q != null);

  factory _$_FindUsersArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUsersArgs_FreezedFromJson(json);

  @override
  final String q;

  @override
  String toString() {
    return '_FindUsersArgs(q: $q)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUsersArgs_Freezed &&
            (identical(other.q, q) ||
                const DeepCollectionEquality().equals(other.q, q)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(q);

  @override
  _$FindUsersArgs_FreezedCopyWith<_FindUsersArgs_Freezed> get copyWith =>
      __$FindUsersArgs_FreezedCopyWithImpl<_FindUsersArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUsersArgs_FreezedToJson(this);
  }
}

abstract class _FindUsersArgs_Freezed implements _FindUsersArgs {
  factory _FindUsersArgs_Freezed({@required String q}) =
      _$_FindUsersArgs_Freezed;

  factory _FindUsersArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_FindUsersArgs_Freezed.fromJson;

  @override
  String get q;
  @override
  _$FindUsersArgs_FreezedCopyWith<_FindUsersArgs_Freezed> get copyWith;
}

PingResult _$PingResultFromJson(Map<String, dynamic> json) {
  return _PingResult.fromJson(json);
}

class _$PingResultTearOff {
  const _$PingResultTearOff();

  _PingResult call({@required bool status}) {
    return _PingResult(
      status: status,
    );
  }
}

// ignore: unused_element
const $PingResult = _$PingResultTearOff();

mixin _$PingResult {
  bool get status;

  Map<String, dynamic> toJson();
  $PingResultCopyWith<PingResult> get copyWith;
}

abstract class $PingResultCopyWith<$Res> {
  factory $PingResultCopyWith(
          PingResult value, $Res Function(PingResult) then) =
      _$PingResultCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$PingResultCopyWithImpl<$Res> implements $PingResultCopyWith<$Res> {
  _$PingResultCopyWithImpl(this._value, this._then);

  final PingResult _value;
  // ignore: unused_field
  final $Res Function(PingResult) _then;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_value.copyWith(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

abstract class _$PingResultCopyWith<$Res> implements $PingResultCopyWith<$Res> {
  factory _$PingResultCopyWith(
          _PingResult value, $Res Function(_PingResult) then) =
      __$PingResultCopyWithImpl<$Res>;
  @override
  $Res call({bool status});
}

class __$PingResultCopyWithImpl<$Res> extends _$PingResultCopyWithImpl<$Res>
    implements _$PingResultCopyWith<$Res> {
  __$PingResultCopyWithImpl(
      _PingResult _value, $Res Function(_PingResult) _then)
      : super(_value, (v) => _then(v as _PingResult));

  @override
  _PingResult get _value => super._value as _PingResult;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_PingResult(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_PingResult implements _PingResult {
  _$_PingResult({@required this.status}) : assert(status != null);

  factory _$_PingResult.fromJson(Map<String, dynamic> json) =>
      _$_$_PingResultFromJson(json);

  @override
  final bool status;

  @override
  String toString() {
    return 'PingResult(status: $status)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _PingResult &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$PingResultCopyWith<_PingResult> get copyWith =>
      __$PingResultCopyWithImpl<_PingResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_PingResultToJson(this);
  }
}

abstract class _PingResult implements PingResult {
  factory _PingResult({@required bool status}) = _$_PingResult;

  factory _PingResult.fromJson(Map<String, dynamic> json) =
      _$_PingResult.fromJson;

  @override
  bool get status;
  @override
  _$PingResultCopyWith<_PingResult> get copyWith;
}

GetUserResult _$GetUserResultFromJson(Map<String, dynamic> json) {
  return _GetUserResult.fromJson(json);
}

class _$GetUserResultTearOff {
  const _$GetUserResultTearOff();

  _GetUserResult call({@required User user}) {
    return _GetUserResult(
      user: user,
    );
  }
}

// ignore: unused_element
const $GetUserResult = _$GetUserResultTearOff();

mixin _$GetUserResult {
  User get user;

  Map<String, dynamic> toJson();
  $GetUserResultCopyWith<GetUserResult> get copyWith;
}

abstract class $GetUserResultCopyWith<$Res> {
  factory $GetUserResultCopyWith(
          GetUserResult value, $Res Function(GetUserResult) then) =
      _$GetUserResultCopyWithImpl<$Res>;
  $Res call({User user});

  $UserCopyWith<$Res> get user;
}

class _$GetUserResultCopyWithImpl<$Res>
    implements $GetUserResultCopyWith<$Res> {
  _$GetUserResultCopyWithImpl(this._value, this._then);

  final GetUserResult _value;
  // ignore: unused_field
  final $Res Function(GetUserResult) _then;

  @override
  $Res call({
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      user: user == freezed ? _value.user : user as User,
    ));
  }

  @override
  $UserCopyWith<$Res> get user {
    if (_value.user == null) {
      return null;
    }
    return $UserCopyWith<$Res>(_value.user, (value) {
      return _then(_value.copyWith(user: value));
    });
  }
}

abstract class _$GetUserResultCopyWith<$Res>
    implements $GetUserResultCopyWith<$Res> {
  factory _$GetUserResultCopyWith(
          _GetUserResult value, $Res Function(_GetUserResult) then) =
      __$GetUserResultCopyWithImpl<$Res>;
  @override
  $Res call({User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$GetUserResultCopyWithImpl<$Res>
    extends _$GetUserResultCopyWithImpl<$Res>
    implements _$GetUserResultCopyWith<$Res> {
  __$GetUserResultCopyWithImpl(
      _GetUserResult _value, $Res Function(_GetUserResult) _then)
      : super(_value, (v) => _then(v as _GetUserResult));

  @override
  _GetUserResult get _value => super._value as _GetUserResult;

  @override
  $Res call({
    Object user = freezed,
  }) {
    return _then(_GetUserResult(
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetUserResult implements _GetUserResult {
  _$_GetUserResult({@required this.user}) : assert(user != null);

  factory _$_GetUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResultFromJson(json);

  @override
  final User user;

  @override
  String toString() {
    return 'GetUserResult(user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserResult &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(user);

  @override
  _$GetUserResultCopyWith<_GetUserResult> get copyWith =>
      __$GetUserResultCopyWithImpl<_GetUserResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResultToJson(this);
  }
}

abstract class _GetUserResult implements GetUserResult {
  factory _GetUserResult({@required User user}) = _$_GetUserResult;

  factory _GetUserResult.fromJson(Map<String, dynamic> json) =
      _$_GetUserResult.fromJson;

  @override
  User get user;
  @override
  _$GetUserResultCopyWith<_GetUserResult> get copyWith;
}

FindUsersResult _$FindUsersResultFromJson(Map<String, dynamic> json) {
  return _FindUsersResult.fromJson(json);
}

class _$FindUsersResultTearOff {
  const _$FindUsersResultTearOff();

  _FindUsersResult call({@required Page page, @required List<User> users}) {
    return _FindUsersResult(
      page: page,
      users: users,
    );
  }
}

// ignore: unused_element
const $FindUsersResult = _$FindUsersResultTearOff();

mixin _$FindUsersResult {
  Page get page;
  List<User> get users;

  Map<String, dynamic> toJson();
  $FindUsersResultCopyWith<FindUsersResult> get copyWith;
}

abstract class $FindUsersResultCopyWith<$Res> {
  factory $FindUsersResultCopyWith(
          FindUsersResult value, $Res Function(FindUsersResult) then) =
      _$FindUsersResultCopyWithImpl<$Res>;
  $Res call({Page page, List<User> users});

  $PageCopyWith<$Res> get page;
}

class _$FindUsersResultCopyWithImpl<$Res>
    implements $FindUsersResultCopyWith<$Res> {
  _$FindUsersResultCopyWithImpl(this._value, this._then);

  final FindUsersResult _value;
  // ignore: unused_field
  final $Res Function(FindUsersResult) _then;

  @override
  $Res call({
    Object page = freezed,
    Object users = freezed,
  }) {
    return _then(_value.copyWith(
      page: page == freezed ? _value.page : page as Page,
      users: users == freezed ? _value.users : users as List<User>,
    ));
  }

  @override
  $PageCopyWith<$Res> get page {
    if (_value.page == null) {
      return null;
    }
    return $PageCopyWith<$Res>(_value.page, (value) {
      return _then(_value.copyWith(page: value));
    });
  }
}

abstract class _$FindUsersResultCopyWith<$Res>
    implements $FindUsersResultCopyWith<$Res> {
  factory _$FindUsersResultCopyWith(
          _FindUsersResult value, $Res Function(_FindUsersResult) then) =
      __$FindUsersResultCopyWithImpl<$Res>;
  @override
  $Res call({Page page, List<User> users});

  @override
  $PageCopyWith<$Res> get page;
}

class __$FindUsersResultCopyWithImpl<$Res>
    extends _$FindUsersResultCopyWithImpl<$Res>
    implements _$FindUsersResultCopyWith<$Res> {
  __$FindUsersResultCopyWithImpl(
      _FindUsersResult _value, $Res Function(_FindUsersResult) _then)
      : super(_value, (v) => _then(v as _FindUsersResult));

  @override
  _FindUsersResult get _value => super._value as _FindUsersResult;

  @override
  $Res call({
    Object page = freezed,
    Object users = freezed,
  }) {
    return _then(_FindUsersResult(
      page: page == freezed ? _value.page : page as Page,
      users: users == freezed ? _value.users : users as List<User>,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_FindUsersResult implements _FindUsersResult {
  _$_FindUsersResult({@required this.page, @required this.users})
      : assert(page != null),
        assert(users != null);

  factory _$_FindUsersResult.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUsersResultFromJson(json);

  @override
  final Page page;
  @override
  final List<User> users;

  @override
  String toString() {
    return 'FindUsersResult(page: $page, users: $users)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUsersResult &&
            (identical(other.page, page) ||
                const DeepCollectionEquality().equals(other.page, page)) &&
            (identical(other.users, users) ||
                const DeepCollectionEquality().equals(other.users, users)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(page) ^
      const DeepCollectionEquality().hash(users);

  @override
  _$FindUsersResultCopyWith<_FindUsersResult> get copyWith =>
      __$FindUsersResultCopyWithImpl<_FindUsersResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUsersResultToJson(this);
  }
}

abstract class _FindUsersResult implements FindUsersResult {
  factory _FindUsersResult({@required Page page, @required List<User> users}) =
      _$_FindUsersResult;

  factory _FindUsersResult.fromJson(Map<String, dynamic> json) =
      _$_FindUsersResult.fromJson;

  @override
  Page get page;
  @override
  List<User> get users;
  @override
  _$FindUsersResultCopyWith<_FindUsersResult> get copyWith;
}
