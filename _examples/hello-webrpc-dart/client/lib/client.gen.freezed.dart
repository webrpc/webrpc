// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named

part of 'client.gen.dart';

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

class _$RpcStateTearOff {
  const _$RpcStateTearOff();

  _RpcStateIdle<T> idle<T>() {
    return _RpcStateIdle<T>();
  }

  _RpcStateLoading<T> loading<T>() {
    return _RpcStateLoading<T>();
  }

  _RpcStateOk<T> ok<T>({@required T data}) {
    return _RpcStateOk<T>(
      data: data,
    );
  }

  _RpcStateUnit<T> unit<T>() {
    return _RpcStateUnit<T>();
  }

  _RpcStateErr<T> err<T>(
      {@required String reason, @required int statusCode, String stackTrace}) {
    return _RpcStateErr<T>(
      reason: reason,
      statusCode: statusCode,
      stackTrace: stackTrace,
    );
  }
}

// ignore: unused_element
const $RpcState = _$RpcStateTearOff();

mixin _$RpcState<T> {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  });
}

abstract class $RpcStateCopyWith<T, $Res> {
  factory $RpcStateCopyWith(
          RpcState<T> value, $Res Function(RpcState<T>) then) =
      _$RpcStateCopyWithImpl<T, $Res>;
}

class _$RpcStateCopyWithImpl<T, $Res> implements $RpcStateCopyWith<T, $Res> {
  _$RpcStateCopyWithImpl(this._value, this._then);

  final RpcState<T> _value;
  // ignore: unused_field
  final $Res Function(RpcState<T>) _then;
}

abstract class _$RpcStateIdleCopyWith<T, $Res> {
  factory _$RpcStateIdleCopyWith(
          _RpcStateIdle<T> value, $Res Function(_RpcStateIdle<T>) then) =
      __$RpcStateIdleCopyWithImpl<T, $Res>;
}

class __$RpcStateIdleCopyWithImpl<T, $Res>
    extends _$RpcStateCopyWithImpl<T, $Res>
    implements _$RpcStateIdleCopyWith<T, $Res> {
  __$RpcStateIdleCopyWithImpl(
      _RpcStateIdle<T> _value, $Res Function(_RpcStateIdle<T>) _then)
      : super(_value, (v) => _then(v as _RpcStateIdle<T>));

  @override
  _RpcStateIdle<T> get _value => super._value as _RpcStateIdle<T>;
}

class _$_RpcStateIdle<T> implements _RpcStateIdle<T> {
  const _$_RpcStateIdle();

  @override
  String toString() {
    return 'RpcState<$T>.idle()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _RpcStateIdle<T>);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return idle();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (idle != null) {
      return idle();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return idle(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (idle != null) {
      return idle(this);
    }
    return orElse();
  }
}

abstract class _RpcStateIdle<T> implements RpcState<T> {
  const factory _RpcStateIdle() = _$_RpcStateIdle<T>;
}

abstract class _$RpcStateLoadingCopyWith<T, $Res> {
  factory _$RpcStateLoadingCopyWith(
          _RpcStateLoading<T> value, $Res Function(_RpcStateLoading<T>) then) =
      __$RpcStateLoadingCopyWithImpl<T, $Res>;
}

class __$RpcStateLoadingCopyWithImpl<T, $Res>
    extends _$RpcStateCopyWithImpl<T, $Res>
    implements _$RpcStateLoadingCopyWith<T, $Res> {
  __$RpcStateLoadingCopyWithImpl(
      _RpcStateLoading<T> _value, $Res Function(_RpcStateLoading<T>) _then)
      : super(_value, (v) => _then(v as _RpcStateLoading<T>));

  @override
  _RpcStateLoading<T> get _value => super._value as _RpcStateLoading<T>;
}

class _$_RpcStateLoading<T> implements _RpcStateLoading<T> {
  const _$_RpcStateLoading();

  @override
  String toString() {
    return 'RpcState<$T>.loading()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _RpcStateLoading<T>);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return loading();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (loading != null) {
      return loading();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return loading(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (loading != null) {
      return loading(this);
    }
    return orElse();
  }
}

abstract class _RpcStateLoading<T> implements RpcState<T> {
  const factory _RpcStateLoading() = _$_RpcStateLoading<T>;
}

abstract class _$RpcStateOkCopyWith<T, $Res> {
  factory _$RpcStateOkCopyWith(
          _RpcStateOk<T> value, $Res Function(_RpcStateOk<T>) then) =
      __$RpcStateOkCopyWithImpl<T, $Res>;
  $Res call({T data});
}

class __$RpcStateOkCopyWithImpl<T, $Res> extends _$RpcStateCopyWithImpl<T, $Res>
    implements _$RpcStateOkCopyWith<T, $Res> {
  __$RpcStateOkCopyWithImpl(
      _RpcStateOk<T> _value, $Res Function(_RpcStateOk<T>) _then)
      : super(_value, (v) => _then(v as _RpcStateOk<T>));

  @override
  _RpcStateOk<T> get _value => super._value as _RpcStateOk<T>;

  @override
  $Res call({
    Object data = freezed,
  }) {
    return _then(_RpcStateOk<T>(
      data: data == freezed ? _value.data : data as T,
    ));
  }
}

class _$_RpcStateOk<T> implements _RpcStateOk<T> {
  const _$_RpcStateOk({@required this.data}) : assert(data != null);

  @override
  final T data;

  @override
  String toString() {
    return 'RpcState<$T>.ok(data: $data)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _RpcStateOk<T> &&
            (identical(other.data, data) ||
                const DeepCollectionEquality().equals(other.data, data)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(data);

  @override
  _$RpcStateOkCopyWith<T, _RpcStateOk<T>> get copyWith =>
      __$RpcStateOkCopyWithImpl<T, _RpcStateOk<T>>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return ok(data);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (ok != null) {
      return ok(data);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return ok(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (ok != null) {
      return ok(this);
    }
    return orElse();
  }
}

abstract class _RpcStateOk<T> implements RpcState<T> {
  const factory _RpcStateOk({@required T data}) = _$_RpcStateOk<T>;

  T get data;
  _$RpcStateOkCopyWith<T, _RpcStateOk<T>> get copyWith;
}

abstract class _$RpcStateUnitCopyWith<T, $Res> {
  factory _$RpcStateUnitCopyWith(
          _RpcStateUnit<T> value, $Res Function(_RpcStateUnit<T>) then) =
      __$RpcStateUnitCopyWithImpl<T, $Res>;
}

class __$RpcStateUnitCopyWithImpl<T, $Res>
    extends _$RpcStateCopyWithImpl<T, $Res>
    implements _$RpcStateUnitCopyWith<T, $Res> {
  __$RpcStateUnitCopyWithImpl(
      _RpcStateUnit<T> _value, $Res Function(_RpcStateUnit<T>) _then)
      : super(_value, (v) => _then(v as _RpcStateUnit<T>));

  @override
  _RpcStateUnit<T> get _value => super._value as _RpcStateUnit<T>;
}

class _$_RpcStateUnit<T> implements _RpcStateUnit<T> {
  const _$_RpcStateUnit();

  @override
  String toString() {
    return 'RpcState<$T>.unit()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _RpcStateUnit<T>);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return unit();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (unit != null) {
      return unit();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return unit(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (unit != null) {
      return unit(this);
    }
    return orElse();
  }
}

abstract class _RpcStateUnit<T> implements RpcState<T> {
  const factory _RpcStateUnit() = _$_RpcStateUnit<T>;
}

abstract class _$RpcStateErrCopyWith<T, $Res> {
  factory _$RpcStateErrCopyWith(
          _RpcStateErr<T> value, $Res Function(_RpcStateErr<T>) then) =
      __$RpcStateErrCopyWithImpl<T, $Res>;
  $Res call({String reason, int statusCode, String stackTrace});
}

class __$RpcStateErrCopyWithImpl<T, $Res>
    extends _$RpcStateCopyWithImpl<T, $Res>
    implements _$RpcStateErrCopyWith<T, $Res> {
  __$RpcStateErrCopyWithImpl(
      _RpcStateErr<T> _value, $Res Function(_RpcStateErr<T>) _then)
      : super(_value, (v) => _then(v as _RpcStateErr<T>));

  @override
  _RpcStateErr<T> get _value => super._value as _RpcStateErr<T>;

  @override
  $Res call({
    Object reason = freezed,
    Object statusCode = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_RpcStateErr<T>(
      reason: reason == freezed ? _value.reason : reason as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

class _$_RpcStateErr<T> implements _RpcStateErr<T> {
  const _$_RpcStateErr(
      {@required this.reason, @required this.statusCode, this.stackTrace})
      : assert(reason != null),
        assert(statusCode != null);

  @override
  final String reason;
  @override
  final int statusCode;
  @override
  final String stackTrace;

  @override
  String toString() {
    return 'RpcState<$T>.err(reason: $reason, statusCode: $statusCode, stackTrace: $stackTrace)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _RpcStateErr<T> &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$RpcStateErrCopyWith<T, _RpcStateErr<T>> get copyWith =>
      __$RpcStateErrCopyWithImpl<T, _RpcStateErr<T>>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result idle(),
    @required Result loading(),
    @required Result ok(T data),
    @required Result unit(),
    @required Result err(String reason, int statusCode, String stackTrace),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return err(reason, statusCode, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result idle(),
    Result loading(),
    Result ok(T data),
    Result unit(),
    Result err(String reason, int statusCode, String stackTrace),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (err != null) {
      return err(reason, statusCode, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result idle(_RpcStateIdle<T> value),
    @required Result loading(_RpcStateLoading<T> value),
    @required Result ok(_RpcStateOk<T> value),
    @required Result unit(_RpcStateUnit<T> value),
    @required Result err(_RpcStateErr<T> value),
  }) {
    assert(idle != null);
    assert(loading != null);
    assert(ok != null);
    assert(unit != null);
    assert(err != null);
    return err(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result idle(_RpcStateIdle<T> value),
    Result loading(_RpcStateLoading<T> value),
    Result ok(_RpcStateOk<T> value),
    Result unit(_RpcStateUnit<T> value),
    Result err(_RpcStateErr<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (err != null) {
      return err(this);
    }
    return orElse();
  }
}

abstract class _RpcStateErr<T> implements RpcState<T> {
  const factory _RpcStateErr(
      {@required String reason,
      @required int statusCode,
      String stackTrace}) = _$_RpcStateErr<T>;

  String get reason;
  int get statusCode;
  String get stackTrace;
  _$RpcStateErrCopyWith<T, _RpcStateErr<T>> get copyWith;
}
