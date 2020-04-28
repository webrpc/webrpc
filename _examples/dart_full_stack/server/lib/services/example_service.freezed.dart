// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named

part of 'example_service.dart';

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

Empty _$EmptyFromJson(Map<String, dynamic> json) {
  return _Empty.fromJson(json);
}

class _$EmptyTearOff {
  const _$EmptyTearOff();

  _Empty call() {
    return _Empty();
  }
}

// ignore: unused_element
const $Empty = _$EmptyTearOff();

mixin _$Empty {
  Map<String, dynamic> toJson();
}

abstract class $EmptyCopyWith<$Res> {
  factory $EmptyCopyWith(Empty value, $Res Function(Empty) then) =
      _$EmptyCopyWithImpl<$Res>;
}

class _$EmptyCopyWithImpl<$Res> implements $EmptyCopyWith<$Res> {
  _$EmptyCopyWithImpl(this._value, this._then);

  final Empty _value;
  // ignore: unused_field
  final $Res Function(Empty) _then;
}

abstract class _$EmptyCopyWith<$Res> {
  factory _$EmptyCopyWith(_Empty value, $Res Function(_Empty) then) =
      __$EmptyCopyWithImpl<$Res>;
}

class __$EmptyCopyWithImpl<$Res> extends _$EmptyCopyWithImpl<$Res>
    implements _$EmptyCopyWith<$Res> {
  __$EmptyCopyWithImpl(_Empty _value, $Res Function(_Empty) _then)
      : super(_value, (v) => _then(v as _Empty));

  @override
  _Empty get _value => super._value as _Empty;
}

@JsonSerializable()
class _$_Empty implements _Empty {
  _$_Empty();

  factory _$_Empty.fromJson(Map<String, dynamic> json) =>
      _$_$_EmptyFromJson(json);

  @override
  String toString() {
    return 'Empty()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _Empty);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  Map<String, dynamic> toJson() {
    return _$_$_EmptyToJson(this);
  }
}

abstract class _Empty implements Empty {
  factory _Empty() = _$_Empty;

  factory _Empty.fromJson(Map<String, dynamic> json) = _$_Empty.fromJson;
}

User _$UserFromJson(Map<String, dynamic> json) {
  return _User.fromJson(json);
}

class _$UserTearOff {
  const _$UserTearOff();

  _User call(
      {@required int id,
      @required @JsonKey(name: 'USERNAME') String username,
      @required String role}) {
    return _User(
      id: id,
      username: username,
      role: role,
    );
  }
}

// ignore: unused_element
const $User = _$UserTearOff();

mixin _$User {
  int get id;
  @JsonKey(name: 'USERNAME')
  String get username;
  String get role;

  Map<String, dynamic> toJson();
  $UserCopyWith<User> get copyWith;
}

abstract class $UserCopyWith<$Res> {
  factory $UserCopyWith(User value, $Res Function(User) then) =
      _$UserCopyWithImpl<$Res>;
  $Res call({int id, @JsonKey(name: 'USERNAME') String username, String role});
}

class _$UserCopyWithImpl<$Res> implements $UserCopyWith<$Res> {
  _$UserCopyWithImpl(this._value, this._then);

  final User _value;
  // ignore: unused_field
  final $Res Function(User) _then;

  @override
  $Res call({
    Object id = freezed,
    Object username = freezed,
    Object role = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as int,
      username: username == freezed ? _value.username : username as String,
      role: role == freezed ? _value.role : role as String,
    ));
  }
}

abstract class _$UserCopyWith<$Res> implements $UserCopyWith<$Res> {
  factory _$UserCopyWith(_User value, $Res Function(_User) then) =
      __$UserCopyWithImpl<$Res>;
  @override
  $Res call({int id, @JsonKey(name: 'USERNAME') String username, String role});
}

class __$UserCopyWithImpl<$Res> extends _$UserCopyWithImpl<$Res>
    implements _$UserCopyWith<$Res> {
  __$UserCopyWithImpl(_User _value, $Res Function(_User) _then)
      : super(_value, (v) => _then(v as _User));

  @override
  _User get _value => super._value as _User;

  @override
  $Res call({
    Object id = freezed,
    Object username = freezed,
    Object role = freezed,
  }) {
    return _then(_User(
      id: id == freezed ? _value.id : id as int,
      username: username == freezed ? _value.username : username as String,
      role: role == freezed ? _value.role : role as String,
    ));
  }
}

@JsonSerializable()
class _$_User implements _User {
  _$_User(
      {@required this.id,
      @required @JsonKey(name: 'USERNAME') this.username,
      @required this.role})
      : assert(id != null),
        assert(username != null),
        assert(role != null);

  factory _$_User.fromJson(Map<String, dynamic> json) =>
      _$_$_UserFromJson(json);

  @override
  final int id;
  @override
  @JsonKey(name: 'USERNAME')
  final String username;
  @override
  final String role;

  @override
  String toString() {
    return 'User(id: $id, username: $username, role: $role)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _User &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.username, username) ||
                const DeepCollectionEquality()
                    .equals(other.username, username)) &&
            (identical(other.role, role) ||
                const DeepCollectionEquality().equals(other.role, role)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(username) ^
      const DeepCollectionEquality().hash(role);

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
      {@required int id,
      @required @JsonKey(name: 'USERNAME') String username,
      @required String role}) = _$_User;

  factory _User.fromJson(Map<String, dynamic> json) = _$_User.fromJson;

  @override
  int get id;
  @override
  @JsonKey(name: 'USERNAME')
  String get username;
  @override
  String get role;
  @override
  _$UserCopyWith<_User> get copyWith;
}

SearchFilter _$SearchFilterFromJson(Map<String, dynamic> json) {
  return _SearchFilter.fromJson(json);
}

class _$SearchFilterTearOff {
  const _$SearchFilterTearOff();

  _SearchFilter call({@required String q}) {
    return _SearchFilter(
      q: q,
    );
  }
}

// ignore: unused_element
const $SearchFilter = _$SearchFilterTearOff();

mixin _$SearchFilter {
  String get q;

  Map<String, dynamic> toJson();
  $SearchFilterCopyWith<SearchFilter> get copyWith;
}

abstract class $SearchFilterCopyWith<$Res> {
  factory $SearchFilterCopyWith(
          SearchFilter value, $Res Function(SearchFilter) then) =
      _$SearchFilterCopyWithImpl<$Res>;
  $Res call({String q});
}

class _$SearchFilterCopyWithImpl<$Res> implements $SearchFilterCopyWith<$Res> {
  _$SearchFilterCopyWithImpl(this._value, this._then);

  final SearchFilter _value;
  // ignore: unused_field
  final $Res Function(SearchFilter) _then;

  @override
  $Res call({
    Object q = freezed,
  }) {
    return _then(_value.copyWith(
      q: q == freezed ? _value.q : q as String,
    ));
  }
}

abstract class _$SearchFilterCopyWith<$Res>
    implements $SearchFilterCopyWith<$Res> {
  factory _$SearchFilterCopyWith(
          _SearchFilter value, $Res Function(_SearchFilter) then) =
      __$SearchFilterCopyWithImpl<$Res>;
  @override
  $Res call({String q});
}

class __$SearchFilterCopyWithImpl<$Res> extends _$SearchFilterCopyWithImpl<$Res>
    implements _$SearchFilterCopyWith<$Res> {
  __$SearchFilterCopyWithImpl(
      _SearchFilter _value, $Res Function(_SearchFilter) _then)
      : super(_value, (v) => _then(v as _SearchFilter));

  @override
  _SearchFilter get _value => super._value as _SearchFilter;

  @override
  $Res call({
    Object q = freezed,
  }) {
    return _then(_SearchFilter(
      q: q == freezed ? _value.q : q as String,
    ));
  }
}

@JsonSerializable()
class _$_SearchFilter implements _SearchFilter {
  _$_SearchFilter({@required this.q}) : assert(q != null);

  factory _$_SearchFilter.fromJson(Map<String, dynamic> json) =>
      _$_$_SearchFilterFromJson(json);

  @override
  final String q;

  @override
  String toString() {
    return 'SearchFilter(q: $q)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _SearchFilter &&
            (identical(other.q, q) ||
                const DeepCollectionEquality().equals(other.q, q)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(q);

  @override
  _$SearchFilterCopyWith<_SearchFilter> get copyWith =>
      __$SearchFilterCopyWithImpl<_SearchFilter>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_SearchFilterToJson(this);
  }
}

abstract class _SearchFilter implements SearchFilter {
  factory _SearchFilter({@required String q}) = _$_SearchFilter;

  factory _SearchFilter.fromJson(Map<String, dynamic> json) =
      _$_SearchFilter.fromJson;

  @override
  String get q;
  @override
  _$SearchFilterCopyWith<_SearchFilter> get copyWith;
}

Version _$VersionFromJson(Map<String, dynamic> json) {
  return _Version.fromJson(json);
}

class _$VersionTearOff {
  const _$VersionTearOff();

  _Version call(
      {@required String webrpcVersion,
      @required String schemaVersion,
      @required String schemaHash}) {
    return _Version(
      webrpcVersion: webrpcVersion,
      schemaVersion: schemaVersion,
      schemaHash: schemaHash,
    );
  }
}

// ignore: unused_element
const $Version = _$VersionTearOff();

mixin _$Version {
  String get webrpcVersion;
  String get schemaVersion;
  String get schemaHash;

  Map<String, dynamic> toJson();
  $VersionCopyWith<Version> get copyWith;
}

abstract class $VersionCopyWith<$Res> {
  factory $VersionCopyWith(Version value, $Res Function(Version) then) =
      _$VersionCopyWithImpl<$Res>;
  $Res call({String webrpcVersion, String schemaVersion, String schemaHash});
}

class _$VersionCopyWithImpl<$Res> implements $VersionCopyWith<$Res> {
  _$VersionCopyWithImpl(this._value, this._then);

  final Version _value;
  // ignore: unused_field
  final $Res Function(Version) _then;

  @override
  $Res call({
    Object webrpcVersion = freezed,
    Object schemaVersion = freezed,
    Object schemaHash = freezed,
  }) {
    return _then(_value.copyWith(
      webrpcVersion: webrpcVersion == freezed
          ? _value.webrpcVersion
          : webrpcVersion as String,
      schemaVersion: schemaVersion == freezed
          ? _value.schemaVersion
          : schemaVersion as String,
      schemaHash:
          schemaHash == freezed ? _value.schemaHash : schemaHash as String,
    ));
  }
}

abstract class _$VersionCopyWith<$Res> implements $VersionCopyWith<$Res> {
  factory _$VersionCopyWith(_Version value, $Res Function(_Version) then) =
      __$VersionCopyWithImpl<$Res>;
  @override
  $Res call({String webrpcVersion, String schemaVersion, String schemaHash});
}

class __$VersionCopyWithImpl<$Res> extends _$VersionCopyWithImpl<$Res>
    implements _$VersionCopyWith<$Res> {
  __$VersionCopyWithImpl(_Version _value, $Res Function(_Version) _then)
      : super(_value, (v) => _then(v as _Version));

  @override
  _Version get _value => super._value as _Version;

  @override
  $Res call({
    Object webrpcVersion = freezed,
    Object schemaVersion = freezed,
    Object schemaHash = freezed,
  }) {
    return _then(_Version(
      webrpcVersion: webrpcVersion == freezed
          ? _value.webrpcVersion
          : webrpcVersion as String,
      schemaVersion: schemaVersion == freezed
          ? _value.schemaVersion
          : schemaVersion as String,
      schemaHash:
          schemaHash == freezed ? _value.schemaHash : schemaHash as String,
    ));
  }
}

@JsonSerializable()
class _$_Version implements _Version {
  _$_Version(
      {@required this.webrpcVersion,
      @required this.schemaVersion,
      @required this.schemaHash})
      : assert(webrpcVersion != null),
        assert(schemaVersion != null),
        assert(schemaHash != null);

  factory _$_Version.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionFromJson(json);

  @override
  final String webrpcVersion;
  @override
  final String schemaVersion;
  @override
  final String schemaHash;

  @override
  String toString() {
    return 'Version(webrpcVersion: $webrpcVersion, schemaVersion: $schemaVersion, schemaHash: $schemaHash)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _Version &&
            (identical(other.webrpcVersion, webrpcVersion) ||
                const DeepCollectionEquality()
                    .equals(other.webrpcVersion, webrpcVersion)) &&
            (identical(other.schemaVersion, schemaVersion) ||
                const DeepCollectionEquality()
                    .equals(other.schemaVersion, schemaVersion)) &&
            (identical(other.schemaHash, schemaHash) ||
                const DeepCollectionEquality()
                    .equals(other.schemaHash, schemaHash)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(webrpcVersion) ^
      const DeepCollectionEquality().hash(schemaVersion) ^
      const DeepCollectionEquality().hash(schemaHash);

  @override
  _$VersionCopyWith<_Version> get copyWith =>
      __$VersionCopyWithImpl<_Version>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionToJson(this);
  }
}

abstract class _Version implements Version {
  factory _Version(
      {@required String webrpcVersion,
      @required String schemaVersion,
      @required String schemaHash}) = _$_Version;

  factory _Version.fromJson(Map<String, dynamic> json) = _$_Version.fromJson;

  @override
  String get webrpcVersion;
  @override
  String get schemaVersion;
  @override
  String get schemaHash;
  @override
  _$VersionCopyWith<_Version> get copyWith;
}

ComplexType _$ComplexTypeFromJson(Map<String, dynamic> json) {
  return _ComplexType.fromJson(json);
}

class _$ComplexTypeTearOff {
  const _$ComplexTypeTearOff();

  _ComplexType call(
      {@required Map<String, dynamic> meta,
      @required Map<String, Map<String, int>> metaNestedExample,
      @required List<String> namesList,
      @required List<int> numsList,
      @required List<List<String>> doubleArray,
      @required List<User> listOfUsers,
      @required Map<String, User> mapOfUsers,
      @required User user,
      List<Map<String, int>> listOfMaps}) {
    return _ComplexType(
      meta: meta,
      metaNestedExample: metaNestedExample,
      namesList: namesList,
      numsList: numsList,
      doubleArray: doubleArray,
      listOfUsers: listOfUsers,
      mapOfUsers: mapOfUsers,
      user: user,
      listOfMaps: listOfMaps,
    );
  }
}

// ignore: unused_element
const $ComplexType = _$ComplexTypeTearOff();

mixin _$ComplexType {
  Map<String, dynamic> get meta;
  Map<String, Map<String, int>> get metaNestedExample;
  List<String> get namesList;
  List<int> get numsList;
  List<List<String>> get doubleArray;
  List<User> get listOfUsers;
  Map<String, User> get mapOfUsers;
  User get user;
  List<Map<String, int>> get listOfMaps;

  Map<String, dynamic> toJson();
  $ComplexTypeCopyWith<ComplexType> get copyWith;
}

abstract class $ComplexTypeCopyWith<$Res> {
  factory $ComplexTypeCopyWith(
          ComplexType value, $Res Function(ComplexType) then) =
      _$ComplexTypeCopyWithImpl<$Res>;
  $Res call(
      {Map<String, dynamic> meta,
      Map<String, Map<String, int>> metaNestedExample,
      List<String> namesList,
      List<int> numsList,
      List<List<String>> doubleArray,
      List<User> listOfUsers,
      Map<String, User> mapOfUsers,
      User user,
      List<Map<String, int>> listOfMaps});

  $UserCopyWith<$Res> get user;
}

class _$ComplexTypeCopyWithImpl<$Res> implements $ComplexTypeCopyWith<$Res> {
  _$ComplexTypeCopyWithImpl(this._value, this._then);

  final ComplexType _value;
  // ignore: unused_field
  final $Res Function(ComplexType) _then;

  @override
  $Res call({
    Object meta = freezed,
    Object metaNestedExample = freezed,
    Object namesList = freezed,
    Object numsList = freezed,
    Object doubleArray = freezed,
    Object listOfUsers = freezed,
    Object mapOfUsers = freezed,
    Object user = freezed,
    Object listOfMaps = freezed,
  }) {
    return _then(_value.copyWith(
      meta: meta == freezed ? _value.meta : meta as Map<String, dynamic>,
      metaNestedExample: metaNestedExample == freezed
          ? _value.metaNestedExample
          : metaNestedExample as Map<String, Map<String, int>>,
      namesList:
          namesList == freezed ? _value.namesList : namesList as List<String>,
      numsList: numsList == freezed ? _value.numsList : numsList as List<int>,
      doubleArray: doubleArray == freezed
          ? _value.doubleArray
          : doubleArray as List<List<String>>,
      listOfUsers: listOfUsers == freezed
          ? _value.listOfUsers
          : listOfUsers as List<User>,
      mapOfUsers: mapOfUsers == freezed
          ? _value.mapOfUsers
          : mapOfUsers as Map<String, User>,
      user: user == freezed ? _value.user : user as User,
      listOfMaps: listOfMaps == freezed
          ? _value.listOfMaps
          : listOfMaps as List<Map<String, int>>,
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

abstract class _$ComplexTypeCopyWith<$Res>
    implements $ComplexTypeCopyWith<$Res> {
  factory _$ComplexTypeCopyWith(
          _ComplexType value, $Res Function(_ComplexType) then) =
      __$ComplexTypeCopyWithImpl<$Res>;
  @override
  $Res call(
      {Map<String, dynamic> meta,
      Map<String, Map<String, int>> metaNestedExample,
      List<String> namesList,
      List<int> numsList,
      List<List<String>> doubleArray,
      List<User> listOfUsers,
      Map<String, User> mapOfUsers,
      User user,
      List<Map<String, int>> listOfMaps});

  @override
  $UserCopyWith<$Res> get user;
}

class __$ComplexTypeCopyWithImpl<$Res> extends _$ComplexTypeCopyWithImpl<$Res>
    implements _$ComplexTypeCopyWith<$Res> {
  __$ComplexTypeCopyWithImpl(
      _ComplexType _value, $Res Function(_ComplexType) _then)
      : super(_value, (v) => _then(v as _ComplexType));

  @override
  _ComplexType get _value => super._value as _ComplexType;

  @override
  $Res call({
    Object meta = freezed,
    Object metaNestedExample = freezed,
    Object namesList = freezed,
    Object numsList = freezed,
    Object doubleArray = freezed,
    Object listOfUsers = freezed,
    Object mapOfUsers = freezed,
    Object user = freezed,
    Object listOfMaps = freezed,
  }) {
    return _then(_ComplexType(
      meta: meta == freezed ? _value.meta : meta as Map<String, dynamic>,
      metaNestedExample: metaNestedExample == freezed
          ? _value.metaNestedExample
          : metaNestedExample as Map<String, Map<String, int>>,
      namesList:
          namesList == freezed ? _value.namesList : namesList as List<String>,
      numsList: numsList == freezed ? _value.numsList : numsList as List<int>,
      doubleArray: doubleArray == freezed
          ? _value.doubleArray
          : doubleArray as List<List<String>>,
      listOfUsers: listOfUsers == freezed
          ? _value.listOfUsers
          : listOfUsers as List<User>,
      mapOfUsers: mapOfUsers == freezed
          ? _value.mapOfUsers
          : mapOfUsers as Map<String, User>,
      user: user == freezed ? _value.user : user as User,
      listOfMaps: listOfMaps == freezed
          ? _value.listOfMaps
          : listOfMaps as List<Map<String, int>>,
    ));
  }
}

@JsonSerializable()
class _$_ComplexType implements _ComplexType {
  _$_ComplexType(
      {@required this.meta,
      @required this.metaNestedExample,
      @required this.namesList,
      @required this.numsList,
      @required this.doubleArray,
      @required this.listOfUsers,
      @required this.mapOfUsers,
      @required this.user,
      this.listOfMaps})
      : assert(meta != null),
        assert(metaNestedExample != null),
        assert(namesList != null),
        assert(numsList != null),
        assert(doubleArray != null),
        assert(listOfUsers != null),
        assert(mapOfUsers != null),
        assert(user != null);

  factory _$_ComplexType.fromJson(Map<String, dynamic> json) =>
      _$_$_ComplexTypeFromJson(json);

  @override
  final Map<String, dynamic> meta;
  @override
  final Map<String, Map<String, int>> metaNestedExample;
  @override
  final List<String> namesList;
  @override
  final List<int> numsList;
  @override
  final List<List<String>> doubleArray;
  @override
  final List<User> listOfUsers;
  @override
  final Map<String, User> mapOfUsers;
  @override
  final User user;
  @override
  final List<Map<String, int>> listOfMaps;

  @override
  String toString() {
    return 'ComplexType(meta: $meta, metaNestedExample: $metaNestedExample, namesList: $namesList, numsList: $numsList, doubleArray: $doubleArray, listOfUsers: $listOfUsers, mapOfUsers: $mapOfUsers, user: $user, listOfMaps: $listOfMaps)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ComplexType &&
            (identical(other.meta, meta) ||
                const DeepCollectionEquality().equals(other.meta, meta)) &&
            (identical(other.metaNestedExample, metaNestedExample) ||
                const DeepCollectionEquality()
                    .equals(other.metaNestedExample, metaNestedExample)) &&
            (identical(other.namesList, namesList) ||
                const DeepCollectionEquality()
                    .equals(other.namesList, namesList)) &&
            (identical(other.numsList, numsList) ||
                const DeepCollectionEquality()
                    .equals(other.numsList, numsList)) &&
            (identical(other.doubleArray, doubleArray) ||
                const DeepCollectionEquality()
                    .equals(other.doubleArray, doubleArray)) &&
            (identical(other.listOfUsers, listOfUsers) ||
                const DeepCollectionEquality()
                    .equals(other.listOfUsers, listOfUsers)) &&
            (identical(other.mapOfUsers, mapOfUsers) ||
                const DeepCollectionEquality()
                    .equals(other.mapOfUsers, mapOfUsers)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)) &&
            (identical(other.listOfMaps, listOfMaps) ||
                const DeepCollectionEquality()
                    .equals(other.listOfMaps, listOfMaps)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(meta) ^
      const DeepCollectionEquality().hash(metaNestedExample) ^
      const DeepCollectionEquality().hash(namesList) ^
      const DeepCollectionEquality().hash(numsList) ^
      const DeepCollectionEquality().hash(doubleArray) ^
      const DeepCollectionEquality().hash(listOfUsers) ^
      const DeepCollectionEquality().hash(mapOfUsers) ^
      const DeepCollectionEquality().hash(user) ^
      const DeepCollectionEquality().hash(listOfMaps);

  @override
  _$ComplexTypeCopyWith<_ComplexType> get copyWith =>
      __$ComplexTypeCopyWithImpl<_ComplexType>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ComplexTypeToJson(this);
  }
}

abstract class _ComplexType implements ComplexType {
  factory _ComplexType(
      {@required Map<String, dynamic> meta,
      @required Map<String, Map<String, int>> metaNestedExample,
      @required List<String> namesList,
      @required List<int> numsList,
      @required List<List<String>> doubleArray,
      @required List<User> listOfUsers,
      @required Map<String, User> mapOfUsers,
      @required User user,
      List<Map<String, int>> listOfMaps}) = _$_ComplexType;

  factory _ComplexType.fromJson(Map<String, dynamic> json) =
      _$_ComplexType.fromJson;

  @override
  Map<String, dynamic> get meta;
  @override
  Map<String, Map<String, int>> get metaNestedExample;
  @override
  List<String> get namesList;
  @override
  List<int> get numsList;
  @override
  List<List<String>> get doubleArray;
  @override
  List<User> get listOfUsers;
  @override
  Map<String, User> get mapOfUsers;
  @override
  User get user;
  @override
  List<Map<String, int>> get listOfMaps;
  @override
  _$ComplexTypeCopyWith<_ComplexType> get copyWith;
}

_ExampleServiceGetUserArgs _$_ExampleServiceGetUserArgsFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceGetUserArgs_Freezed.fromJson(json);
}

class _$_ExampleServiceGetUserArgsTearOff {
  const _$_ExampleServiceGetUserArgsTearOff();

  _ExampleServiceGetUserArgs_Freezed call(
      {@required Map<String, String> header, @required int userID}) {
    return _ExampleServiceGetUserArgs_Freezed(
      header: header,
      userID: userID,
    );
  }
}

// ignore: unused_element
const _$ExampleServiceGetUserArgs = _$_ExampleServiceGetUserArgsTearOff();

mixin _$_ExampleServiceGetUserArgs {
  Map<String, String> get header;
  int get userID;

  Map<String, dynamic> toJson();
  _$ExampleServiceGetUserArgsCopyWith<_ExampleServiceGetUserArgs> get copyWith;
}

abstract class _$ExampleServiceGetUserArgsCopyWith<$Res> {
  factory _$ExampleServiceGetUserArgsCopyWith(_ExampleServiceGetUserArgs value,
          $Res Function(_ExampleServiceGetUserArgs) then) =
      __$ExampleServiceGetUserArgsCopyWithImpl<$Res>;
  $Res call({Map<String, String> header, int userID});
}

class __$ExampleServiceGetUserArgsCopyWithImpl<$Res>
    implements _$ExampleServiceGetUserArgsCopyWith<$Res> {
  __$ExampleServiceGetUserArgsCopyWithImpl(this._value, this._then);

  final _ExampleServiceGetUserArgs _value;
  // ignore: unused_field
  final $Res Function(_ExampleServiceGetUserArgs) _then;

  @override
  $Res call({
    Object header = freezed,
    Object userID = freezed,
  }) {
    return _then(_value.copyWith(
      header: header == freezed ? _value.header : header as Map<String, String>,
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

abstract class _$ExampleServiceGetUserArgs_FreezedCopyWith<$Res>
    implements _$ExampleServiceGetUserArgsCopyWith<$Res> {
  factory _$ExampleServiceGetUserArgs_FreezedCopyWith(
          _ExampleServiceGetUserArgs_Freezed value,
          $Res Function(_ExampleServiceGetUserArgs_Freezed) then) =
      __$ExampleServiceGetUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> header, int userID});
}

class __$ExampleServiceGetUserArgs_FreezedCopyWithImpl<$Res>
    extends __$ExampleServiceGetUserArgsCopyWithImpl<$Res>
    implements _$ExampleServiceGetUserArgs_FreezedCopyWith<$Res> {
  __$ExampleServiceGetUserArgs_FreezedCopyWithImpl(
      _ExampleServiceGetUserArgs_Freezed _value,
      $Res Function(_ExampleServiceGetUserArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _ExampleServiceGetUserArgs_Freezed));

  @override
  _ExampleServiceGetUserArgs_Freezed get _value =>
      super._value as _ExampleServiceGetUserArgs_Freezed;

  @override
  $Res call({
    Object header = freezed,
    Object userID = freezed,
  }) {
    return _then(_ExampleServiceGetUserArgs_Freezed(
      header: header == freezed ? _value.header : header as Map<String, String>,
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceGetUserArgs_Freezed
    implements _ExampleServiceGetUserArgs_Freezed {
  _$_ExampleServiceGetUserArgs_Freezed(
      {@required this.header, @required this.userID})
      : assert(header != null),
        assert(userID != null);

  factory _$_ExampleServiceGetUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_ExampleServiceGetUserArgs_FreezedFromJson(json);

  @override
  final Map<String, String> header;
  @override
  final int userID;

  @override
  String toString() {
    return '_ExampleServiceGetUserArgs(header: $header, userID: $userID)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceGetUserArgs_Freezed &&
            (identical(other.header, header) ||
                const DeepCollectionEquality().equals(other.header, header)) &&
            (identical(other.userID, userID) ||
                const DeepCollectionEquality().equals(other.userID, userID)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(header) ^
      const DeepCollectionEquality().hash(userID);

  @override
  _$ExampleServiceGetUserArgs_FreezedCopyWith<
          _ExampleServiceGetUserArgs_Freezed>
      get copyWith => __$ExampleServiceGetUserArgs_FreezedCopyWithImpl<
          _ExampleServiceGetUserArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceGetUserArgs_FreezedToJson(this);
  }
}

abstract class _ExampleServiceGetUserArgs_Freezed
    implements _ExampleServiceGetUserArgs {
  factory _ExampleServiceGetUserArgs_Freezed(
      {@required Map<String, String> header,
      @required int userID}) = _$_ExampleServiceGetUserArgs_Freezed;

  factory _ExampleServiceGetUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_ExampleServiceGetUserArgs_Freezed.fromJson;

  @override
  Map<String, String> get header;
  @override
  int get userID;
  @override
  _$ExampleServiceGetUserArgs_FreezedCopyWith<
      _ExampleServiceGetUserArgs_Freezed> get copyWith;
}

_ExampleServiceFindUserArgs _$_ExampleServiceFindUserArgsFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceFindUserArgs_Freezed.fromJson(json);
}

class _$_ExampleServiceFindUserArgsTearOff {
  const _$_ExampleServiceFindUserArgsTearOff();

  _ExampleServiceFindUserArgs_Freezed call({@required SearchFilter s}) {
    return _ExampleServiceFindUserArgs_Freezed(
      s: s,
    );
  }
}

// ignore: unused_element
const _$ExampleServiceFindUserArgs = _$_ExampleServiceFindUserArgsTearOff();

mixin _$_ExampleServiceFindUserArgs {
  SearchFilter get s;

  Map<String, dynamic> toJson();
  _$ExampleServiceFindUserArgsCopyWith<_ExampleServiceFindUserArgs>
      get copyWith;
}

abstract class _$ExampleServiceFindUserArgsCopyWith<$Res> {
  factory _$ExampleServiceFindUserArgsCopyWith(
          _ExampleServiceFindUserArgs value,
          $Res Function(_ExampleServiceFindUserArgs) then) =
      __$ExampleServiceFindUserArgsCopyWithImpl<$Res>;
  $Res call({SearchFilter s});

  $SearchFilterCopyWith<$Res> get s;
}

class __$ExampleServiceFindUserArgsCopyWithImpl<$Res>
    implements _$ExampleServiceFindUserArgsCopyWith<$Res> {
  __$ExampleServiceFindUserArgsCopyWithImpl(this._value, this._then);

  final _ExampleServiceFindUserArgs _value;
  // ignore: unused_field
  final $Res Function(_ExampleServiceFindUserArgs) _then;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_value.copyWith(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }

  @override
  $SearchFilterCopyWith<$Res> get s {
    if (_value.s == null) {
      return null;
    }
    return $SearchFilterCopyWith<$Res>(_value.s, (value) {
      return _then(_value.copyWith(s: value));
    });
  }
}

abstract class _$ExampleServiceFindUserArgs_FreezedCopyWith<$Res>
    implements _$ExampleServiceFindUserArgsCopyWith<$Res> {
  factory _$ExampleServiceFindUserArgs_FreezedCopyWith(
          _ExampleServiceFindUserArgs_Freezed value,
          $Res Function(_ExampleServiceFindUserArgs_Freezed) then) =
      __$ExampleServiceFindUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({SearchFilter s});

  @override
  $SearchFilterCopyWith<$Res> get s;
}

class __$ExampleServiceFindUserArgs_FreezedCopyWithImpl<$Res>
    extends __$ExampleServiceFindUserArgsCopyWithImpl<$Res>
    implements _$ExampleServiceFindUserArgs_FreezedCopyWith<$Res> {
  __$ExampleServiceFindUserArgs_FreezedCopyWithImpl(
      _ExampleServiceFindUserArgs_Freezed _value,
      $Res Function(_ExampleServiceFindUserArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _ExampleServiceFindUserArgs_Freezed));

  @override
  _ExampleServiceFindUserArgs_Freezed get _value =>
      super._value as _ExampleServiceFindUserArgs_Freezed;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_ExampleServiceFindUserArgs_Freezed(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceFindUserArgs_Freezed
    implements _ExampleServiceFindUserArgs_Freezed {
  _$_ExampleServiceFindUserArgs_Freezed({@required this.s}) : assert(s != null);

  factory _$_ExampleServiceFindUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_ExampleServiceFindUserArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString() {
    return '_ExampleServiceFindUserArgs(s: $s)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceFindUserArgs_Freezed &&
            (identical(other.s, s) ||
                const DeepCollectionEquality().equals(other.s, s)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(s);

  @override
  _$ExampleServiceFindUserArgs_FreezedCopyWith<
          _ExampleServiceFindUserArgs_Freezed>
      get copyWith => __$ExampleServiceFindUserArgs_FreezedCopyWithImpl<
          _ExampleServiceFindUserArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceFindUserArgs_FreezedToJson(this);
  }
}

abstract class _ExampleServiceFindUserArgs_Freezed
    implements _ExampleServiceFindUserArgs {
  factory _ExampleServiceFindUserArgs_Freezed({@required SearchFilter s}) =
      _$_ExampleServiceFindUserArgs_Freezed;

  factory _ExampleServiceFindUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_ExampleServiceFindUserArgs_Freezed.fromJson;

  @override
  SearchFilter get s;
  @override
  _$ExampleServiceFindUserArgs_FreezedCopyWith<
      _ExampleServiceFindUserArgs_Freezed> get copyWith;
}

ExampleServiceStatusReturn _$ExampleServiceStatusReturnFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceStatusReturn.fromJson(json);
}

class _$ExampleServiceStatusReturnTearOff {
  const _$ExampleServiceStatusReturnTearOff();

  _ExampleServiceStatusReturn call({@required bool status}) {
    return _ExampleServiceStatusReturn(
      status: status,
    );
  }
}

// ignore: unused_element
const $ExampleServiceStatusReturn = _$ExampleServiceStatusReturnTearOff();

mixin _$ExampleServiceStatusReturn {
  bool get status;

  Map<String, dynamic> toJson();
  $ExampleServiceStatusReturnCopyWith<ExampleServiceStatusReturn> get copyWith;
}

abstract class $ExampleServiceStatusReturnCopyWith<$Res> {
  factory $ExampleServiceStatusReturnCopyWith(ExampleServiceStatusReturn value,
          $Res Function(ExampleServiceStatusReturn) then) =
      _$ExampleServiceStatusReturnCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$ExampleServiceStatusReturnCopyWithImpl<$Res>
    implements $ExampleServiceStatusReturnCopyWith<$Res> {
  _$ExampleServiceStatusReturnCopyWithImpl(this._value, this._then);

  final ExampleServiceStatusReturn _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceStatusReturn) _then;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_value.copyWith(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

abstract class _$ExampleServiceStatusReturnCopyWith<$Res>
    implements $ExampleServiceStatusReturnCopyWith<$Res> {
  factory _$ExampleServiceStatusReturnCopyWith(
          _ExampleServiceStatusReturn value,
          $Res Function(_ExampleServiceStatusReturn) then) =
      __$ExampleServiceStatusReturnCopyWithImpl<$Res>;
  @override
  $Res call({bool status});
}

class __$ExampleServiceStatusReturnCopyWithImpl<$Res>
    extends _$ExampleServiceStatusReturnCopyWithImpl<$Res>
    implements _$ExampleServiceStatusReturnCopyWith<$Res> {
  __$ExampleServiceStatusReturnCopyWithImpl(_ExampleServiceStatusReturn _value,
      $Res Function(_ExampleServiceStatusReturn) _then)
      : super(_value, (v) => _then(v as _ExampleServiceStatusReturn));

  @override
  _ExampleServiceStatusReturn get _value =>
      super._value as _ExampleServiceStatusReturn;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_ExampleServiceStatusReturn(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceStatusReturn implements _ExampleServiceStatusReturn {
  _$_ExampleServiceStatusReturn({@required this.status})
      : assert(status != null);

  factory _$_ExampleServiceStatusReturn.fromJson(Map<String, dynamic> json) =>
      _$_$_ExampleServiceStatusReturnFromJson(json);

  @override
  final bool status;

  @override
  String toString() {
    return 'ExampleServiceStatusReturn(status: $status)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceStatusReturn &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$ExampleServiceStatusReturnCopyWith<_ExampleServiceStatusReturn>
      get copyWith => __$ExampleServiceStatusReturnCopyWithImpl<
          _ExampleServiceStatusReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceStatusReturnToJson(this);
  }
}

abstract class _ExampleServiceStatusReturn
    implements ExampleServiceStatusReturn {
  factory _ExampleServiceStatusReturn({@required bool status}) =
      _$_ExampleServiceStatusReturn;

  factory _ExampleServiceStatusReturn.fromJson(Map<String, dynamic> json) =
      _$_ExampleServiceStatusReturn.fromJson;

  @override
  bool get status;
  @override
  _$ExampleServiceStatusReturnCopyWith<_ExampleServiceStatusReturn>
      get copyWith;
}

ExampleServiceVersionReturn _$ExampleServiceVersionReturnFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceVersionReturn.fromJson(json);
}

class _$ExampleServiceVersionReturnTearOff {
  const _$ExampleServiceVersionReturnTearOff();

  _ExampleServiceVersionReturn call({@required Version version}) {
    return _ExampleServiceVersionReturn(
      version: version,
    );
  }
}

// ignore: unused_element
const $ExampleServiceVersionReturn = _$ExampleServiceVersionReturnTearOff();

mixin _$ExampleServiceVersionReturn {
  Version get version;

  Map<String, dynamic> toJson();
  $ExampleServiceVersionReturnCopyWith<ExampleServiceVersionReturn>
      get copyWith;
}

abstract class $ExampleServiceVersionReturnCopyWith<$Res> {
  factory $ExampleServiceVersionReturnCopyWith(
          ExampleServiceVersionReturn value,
          $Res Function(ExampleServiceVersionReturn) then) =
      _$ExampleServiceVersionReturnCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class _$ExampleServiceVersionReturnCopyWithImpl<$Res>
    implements $ExampleServiceVersionReturnCopyWith<$Res> {
  _$ExampleServiceVersionReturnCopyWithImpl(this._value, this._then);

  final ExampleServiceVersionReturn _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceVersionReturn) _then;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_value.copyWith(
      version: version == freezed ? _value.version : version as Version,
    ));
  }

  @override
  $VersionCopyWith<$Res> get version {
    if (_value.version == null) {
      return null;
    }
    return $VersionCopyWith<$Res>(_value.version, (value) {
      return _then(_value.copyWith(version: value));
    });
  }
}

abstract class _$ExampleServiceVersionReturnCopyWith<$Res>
    implements $ExampleServiceVersionReturnCopyWith<$Res> {
  factory _$ExampleServiceVersionReturnCopyWith(
          _ExampleServiceVersionReturn value,
          $Res Function(_ExampleServiceVersionReturn) then) =
      __$ExampleServiceVersionReturnCopyWithImpl<$Res>;
  @override
  $Res call({Version version});

  @override
  $VersionCopyWith<$Res> get version;
}

class __$ExampleServiceVersionReturnCopyWithImpl<$Res>
    extends _$ExampleServiceVersionReturnCopyWithImpl<$Res>
    implements _$ExampleServiceVersionReturnCopyWith<$Res> {
  __$ExampleServiceVersionReturnCopyWithImpl(
      _ExampleServiceVersionReturn _value,
      $Res Function(_ExampleServiceVersionReturn) _then)
      : super(_value, (v) => _then(v as _ExampleServiceVersionReturn));

  @override
  _ExampleServiceVersionReturn get _value =>
      super._value as _ExampleServiceVersionReturn;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_ExampleServiceVersionReturn(
      version: version == freezed ? _value.version : version as Version,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceVersionReturn implements _ExampleServiceVersionReturn {
  _$_ExampleServiceVersionReturn({@required this.version})
      : assert(version != null);

  factory _$_ExampleServiceVersionReturn.fromJson(Map<String, dynamic> json) =>
      _$_$_ExampleServiceVersionReturnFromJson(json);

  @override
  final Version version;

  @override
  String toString() {
    return 'ExampleServiceVersionReturn(version: $version)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceVersionReturn &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$ExampleServiceVersionReturnCopyWith<_ExampleServiceVersionReturn>
      get copyWith => __$ExampleServiceVersionReturnCopyWithImpl<
          _ExampleServiceVersionReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceVersionReturnToJson(this);
  }
}

abstract class _ExampleServiceVersionReturn
    implements ExampleServiceVersionReturn {
  factory _ExampleServiceVersionReturn({@required Version version}) =
      _$_ExampleServiceVersionReturn;

  factory _ExampleServiceVersionReturn.fromJson(Map<String, dynamic> json) =
      _$_ExampleServiceVersionReturn.fromJson;

  @override
  Version get version;
  @override
  _$ExampleServiceVersionReturnCopyWith<_ExampleServiceVersionReturn>
      get copyWith;
}

ExampleServiceGetUserReturn _$ExampleServiceGetUserReturnFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceGetUserReturn.fromJson(json);
}

class _$ExampleServiceGetUserReturnTearOff {
  const _$ExampleServiceGetUserReturnTearOff();

  _ExampleServiceGetUserReturn call({@required int code, @required User user}) {
    return _ExampleServiceGetUserReturn(
      code: code,
      user: user,
    );
  }
}

// ignore: unused_element
const $ExampleServiceGetUserReturn = _$ExampleServiceGetUserReturnTearOff();

mixin _$ExampleServiceGetUserReturn {
  int get code;
  User get user;

  Map<String, dynamic> toJson();
  $ExampleServiceGetUserReturnCopyWith<ExampleServiceGetUserReturn>
      get copyWith;
}

abstract class $ExampleServiceGetUserReturnCopyWith<$Res> {
  factory $ExampleServiceGetUserReturnCopyWith(
          ExampleServiceGetUserReturn value,
          $Res Function(ExampleServiceGetUserReturn) then) =
      _$ExampleServiceGetUserReturnCopyWithImpl<$Res>;
  $Res call({int code, User user});

  $UserCopyWith<$Res> get user;
}

class _$ExampleServiceGetUserReturnCopyWithImpl<$Res>
    implements $ExampleServiceGetUserReturnCopyWith<$Res> {
  _$ExampleServiceGetUserReturnCopyWithImpl(this._value, this._then);

  final ExampleServiceGetUserReturn _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceGetUserReturn) _then;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      code: code == freezed ? _value.code : code as int,
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

abstract class _$ExampleServiceGetUserReturnCopyWith<$Res>
    implements $ExampleServiceGetUserReturnCopyWith<$Res> {
  factory _$ExampleServiceGetUserReturnCopyWith(
          _ExampleServiceGetUserReturn value,
          $Res Function(_ExampleServiceGetUserReturn) then) =
      __$ExampleServiceGetUserReturnCopyWithImpl<$Res>;
  @override
  $Res call({int code, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$ExampleServiceGetUserReturnCopyWithImpl<$Res>
    extends _$ExampleServiceGetUserReturnCopyWithImpl<$Res>
    implements _$ExampleServiceGetUserReturnCopyWith<$Res> {
  __$ExampleServiceGetUserReturnCopyWithImpl(
      _ExampleServiceGetUserReturn _value,
      $Res Function(_ExampleServiceGetUserReturn) _then)
      : super(_value, (v) => _then(v as _ExampleServiceGetUserReturn));

  @override
  _ExampleServiceGetUserReturn get _value =>
      super._value as _ExampleServiceGetUserReturn;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_ExampleServiceGetUserReturn(
      code: code == freezed ? _value.code : code as int,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceGetUserReturn implements _ExampleServiceGetUserReturn {
  _$_ExampleServiceGetUserReturn({@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_ExampleServiceGetUserReturn.fromJson(Map<String, dynamic> json) =>
      _$_$_ExampleServiceGetUserReturnFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString() {
    return 'ExampleServiceGetUserReturn(code: $code, user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceGetUserReturn &&
            (identical(other.code, code) ||
                const DeepCollectionEquality().equals(other.code, code)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(code) ^
      const DeepCollectionEquality().hash(user);

  @override
  _$ExampleServiceGetUserReturnCopyWith<_ExampleServiceGetUserReturn>
      get copyWith => __$ExampleServiceGetUserReturnCopyWithImpl<
          _ExampleServiceGetUserReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceGetUserReturnToJson(this);
  }
}

abstract class _ExampleServiceGetUserReturn
    implements ExampleServiceGetUserReturn {
  factory _ExampleServiceGetUserReturn(
      {@required int code,
      @required User user}) = _$_ExampleServiceGetUserReturn;

  factory _ExampleServiceGetUserReturn.fromJson(Map<String, dynamic> json) =
      _$_ExampleServiceGetUserReturn.fromJson;

  @override
  int get code;
  @override
  User get user;
  @override
  _$ExampleServiceGetUserReturnCopyWith<_ExampleServiceGetUserReturn>
      get copyWith;
}

ExampleServiceFindUserReturn _$ExampleServiceFindUserReturnFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceFindUserReturn.fromJson(json);
}

class _$ExampleServiceFindUserReturnTearOff {
  const _$ExampleServiceFindUserReturnTearOff();

  _ExampleServiceFindUserReturn call(
      {@required String name, @required User user}) {
    return _ExampleServiceFindUserReturn(
      name: name,
      user: user,
    );
  }
}

// ignore: unused_element
const $ExampleServiceFindUserReturn = _$ExampleServiceFindUserReturnTearOff();

mixin _$ExampleServiceFindUserReturn {
  String get name;
  User get user;

  Map<String, dynamic> toJson();
  $ExampleServiceFindUserReturnCopyWith<ExampleServiceFindUserReturn>
      get copyWith;
}

abstract class $ExampleServiceFindUserReturnCopyWith<$Res> {
  factory $ExampleServiceFindUserReturnCopyWith(
          ExampleServiceFindUserReturn value,
          $Res Function(ExampleServiceFindUserReturn) then) =
      _$ExampleServiceFindUserReturnCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$ExampleServiceFindUserReturnCopyWithImpl<$Res>
    implements $ExampleServiceFindUserReturnCopyWith<$Res> {
  _$ExampleServiceFindUserReturnCopyWithImpl(this._value, this._then);

  final ExampleServiceFindUserReturn _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceFindUserReturn) _then;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      name: name == freezed ? _value.name : name as String,
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

abstract class _$ExampleServiceFindUserReturnCopyWith<$Res>
    implements $ExampleServiceFindUserReturnCopyWith<$Res> {
  factory _$ExampleServiceFindUserReturnCopyWith(
          _ExampleServiceFindUserReturn value,
          $Res Function(_ExampleServiceFindUserReturn) then) =
      __$ExampleServiceFindUserReturnCopyWithImpl<$Res>;
  @override
  $Res call({String name, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$ExampleServiceFindUserReturnCopyWithImpl<$Res>
    extends _$ExampleServiceFindUserReturnCopyWithImpl<$Res>
    implements _$ExampleServiceFindUserReturnCopyWith<$Res> {
  __$ExampleServiceFindUserReturnCopyWithImpl(
      _ExampleServiceFindUserReturn _value,
      $Res Function(_ExampleServiceFindUserReturn) _then)
      : super(_value, (v) => _then(v as _ExampleServiceFindUserReturn));

  @override
  _ExampleServiceFindUserReturn get _value =>
      super._value as _ExampleServiceFindUserReturn;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_ExampleServiceFindUserReturn(
      name: name == freezed ? _value.name : name as String,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceFindUserReturn implements _ExampleServiceFindUserReturn {
  _$_ExampleServiceFindUserReturn({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_ExampleServiceFindUserReturn.fromJson(Map<String, dynamic> json) =>
      _$_$_ExampleServiceFindUserReturnFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString() {
    return 'ExampleServiceFindUserReturn(name: $name, user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceFindUserReturn &&
            (identical(other.name, name) ||
                const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(user);

  @override
  _$ExampleServiceFindUserReturnCopyWith<_ExampleServiceFindUserReturn>
      get copyWith => __$ExampleServiceFindUserReturnCopyWithImpl<
          _ExampleServiceFindUserReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceFindUserReturnToJson(this);
  }
}

abstract class _ExampleServiceFindUserReturn
    implements ExampleServiceFindUserReturn {
  factory _ExampleServiceFindUserReturn(
      {@required String name,
      @required User user}) = _$_ExampleServiceFindUserReturn;

  factory _ExampleServiceFindUserReturn.fromJson(Map<String, dynamic> json) =
      _$_ExampleServiceFindUserReturn.fromJson;

  @override
  String get name;
  @override
  User get user;
  @override
  _$ExampleServiceFindUserReturnCopyWith<_ExampleServiceFindUserReturn>
      get copyWith;
}

_AnotherExampleServiceGetUserArgs _$_AnotherExampleServiceGetUserArgsFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceGetUserArgs_Freezed.fromJson(json);
}

class _$_AnotherExampleServiceGetUserArgsTearOff {
  const _$_AnotherExampleServiceGetUserArgsTearOff();

  _AnotherExampleServiceGetUserArgs_Freezed call(
      {@required Map<String, String> header, @required int userID}) {
    return _AnotherExampleServiceGetUserArgs_Freezed(
      header: header,
      userID: userID,
    );
  }
}

// ignore: unused_element
const _$AnotherExampleServiceGetUserArgs =
    _$_AnotherExampleServiceGetUserArgsTearOff();

mixin _$_AnotherExampleServiceGetUserArgs {
  Map<String, String> get header;
  int get userID;

  Map<String, dynamic> toJson();
  _$AnotherExampleServiceGetUserArgsCopyWith<_AnotherExampleServiceGetUserArgs>
      get copyWith;
}

abstract class _$AnotherExampleServiceGetUserArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetUserArgsCopyWith(
          _AnotherExampleServiceGetUserArgs value,
          $Res Function(_AnotherExampleServiceGetUserArgs) then) =
      __$AnotherExampleServiceGetUserArgsCopyWithImpl<$Res>;
  $Res call({Map<String, String> header, int userID});
}

class __$AnotherExampleServiceGetUserArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetUserArgsCopyWith<$Res> {
  __$AnotherExampleServiceGetUserArgsCopyWithImpl(this._value, this._then);

  final _AnotherExampleServiceGetUserArgs _value;
  // ignore: unused_field
  final $Res Function(_AnotherExampleServiceGetUserArgs) _then;

  @override
  $Res call({
    Object header = freezed,
    Object userID = freezed,
  }) {
    return _then(_value.copyWith(
      header: header == freezed ? _value.header : header as Map<String, String>,
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

abstract class _$AnotherExampleServiceGetUserArgs_FreezedCopyWith<$Res>
    implements _$AnotherExampleServiceGetUserArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetUserArgs_FreezedCopyWith(
          _AnotherExampleServiceGetUserArgs_Freezed value,
          $Res Function(_AnotherExampleServiceGetUserArgs_Freezed) then) =
      __$AnotherExampleServiceGetUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> header, int userID});
}

class __$AnotherExampleServiceGetUserArgs_FreezedCopyWithImpl<$Res>
    extends __$AnotherExampleServiceGetUserArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetUserArgs_FreezedCopyWith<$Res> {
  __$AnotherExampleServiceGetUserArgs_FreezedCopyWithImpl(
      _AnotherExampleServiceGetUserArgs_Freezed _value,
      $Res Function(_AnotherExampleServiceGetUserArgs_Freezed) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceGetUserArgs_Freezed));

  @override
  _AnotherExampleServiceGetUserArgs_Freezed get _value =>
      super._value as _AnotherExampleServiceGetUserArgs_Freezed;

  @override
  $Res call({
    Object header = freezed,
    Object userID = freezed,
  }) {
    return _then(_AnotherExampleServiceGetUserArgs_Freezed(
      header: header == freezed ? _value.header : header as Map<String, String>,
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceGetUserArgs_Freezed
    implements _AnotherExampleServiceGetUserArgs_Freezed {
  _$_AnotherExampleServiceGetUserArgs_Freezed(
      {@required this.header, @required this.userID})
      : assert(header != null),
        assert(userID != null);

  factory _$_AnotherExampleServiceGetUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceGetUserArgs_FreezedFromJson(json);

  @override
  final Map<String, String> header;
  @override
  final int userID;

  @override
  String toString() {
    return '_AnotherExampleServiceGetUserArgs(header: $header, userID: $userID)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceGetUserArgs_Freezed &&
            (identical(other.header, header) ||
                const DeepCollectionEquality().equals(other.header, header)) &&
            (identical(other.userID, userID) ||
                const DeepCollectionEquality().equals(other.userID, userID)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(header) ^
      const DeepCollectionEquality().hash(userID);

  @override
  _$AnotherExampleServiceGetUserArgs_FreezedCopyWith<
          _AnotherExampleServiceGetUserArgs_Freezed>
      get copyWith => __$AnotherExampleServiceGetUserArgs_FreezedCopyWithImpl<
          _AnotherExampleServiceGetUserArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceGetUserArgs_FreezedToJson(this);
  }
}

abstract class _AnotherExampleServiceGetUserArgs_Freezed
    implements _AnotherExampleServiceGetUserArgs {
  factory _AnotherExampleServiceGetUserArgs_Freezed(
      {@required Map<String, String> header,
      @required int userID}) = _$_AnotherExampleServiceGetUserArgs_Freezed;

  factory _AnotherExampleServiceGetUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceGetUserArgs_Freezed.fromJson;

  @override
  Map<String, String> get header;
  @override
  int get userID;
  @override
  _$AnotherExampleServiceGetUserArgs_FreezedCopyWith<
      _AnotherExampleServiceGetUserArgs_Freezed> get copyWith;
}

_AnotherExampleServiceFindUserArgs _$_AnotherExampleServiceFindUserArgsFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceFindUserArgs_Freezed.fromJson(json);
}

class _$_AnotherExampleServiceFindUserArgsTearOff {
  const _$_AnotherExampleServiceFindUserArgsTearOff();

  _AnotherExampleServiceFindUserArgs_Freezed call({@required SearchFilter s}) {
    return _AnotherExampleServiceFindUserArgs_Freezed(
      s: s,
    );
  }
}

// ignore: unused_element
const _$AnotherExampleServiceFindUserArgs =
    _$_AnotherExampleServiceFindUserArgsTearOff();

mixin _$_AnotherExampleServiceFindUserArgs {
  SearchFilter get s;

  Map<String, dynamic> toJson();
  _$AnotherExampleServiceFindUserArgsCopyWith<
      _AnotherExampleServiceFindUserArgs> get copyWith;
}

abstract class _$AnotherExampleServiceFindUserArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceFindUserArgsCopyWith(
          _AnotherExampleServiceFindUserArgs value,
          $Res Function(_AnotherExampleServiceFindUserArgs) then) =
      __$AnotherExampleServiceFindUserArgsCopyWithImpl<$Res>;
  $Res call({SearchFilter s});

  $SearchFilterCopyWith<$Res> get s;
}

class __$AnotherExampleServiceFindUserArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceFindUserArgsCopyWith<$Res> {
  __$AnotherExampleServiceFindUserArgsCopyWithImpl(this._value, this._then);

  final _AnotherExampleServiceFindUserArgs _value;
  // ignore: unused_field
  final $Res Function(_AnotherExampleServiceFindUserArgs) _then;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_value.copyWith(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }

  @override
  $SearchFilterCopyWith<$Res> get s {
    if (_value.s == null) {
      return null;
    }
    return $SearchFilterCopyWith<$Res>(_value.s, (value) {
      return _then(_value.copyWith(s: value));
    });
  }
}

abstract class _$AnotherExampleServiceFindUserArgs_FreezedCopyWith<$Res>
    implements _$AnotherExampleServiceFindUserArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceFindUserArgs_FreezedCopyWith(
          _AnotherExampleServiceFindUserArgs_Freezed value,
          $Res Function(_AnotherExampleServiceFindUserArgs_Freezed) then) =
      __$AnotherExampleServiceFindUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({SearchFilter s});

  @override
  $SearchFilterCopyWith<$Res> get s;
}

class __$AnotherExampleServiceFindUserArgs_FreezedCopyWithImpl<$Res>
    extends __$AnotherExampleServiceFindUserArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceFindUserArgs_FreezedCopyWith<$Res> {
  __$AnotherExampleServiceFindUserArgs_FreezedCopyWithImpl(
      _AnotherExampleServiceFindUserArgs_Freezed _value,
      $Res Function(_AnotherExampleServiceFindUserArgs_Freezed) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceFindUserArgs_Freezed));

  @override
  _AnotherExampleServiceFindUserArgs_Freezed get _value =>
      super._value as _AnotherExampleServiceFindUserArgs_Freezed;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_AnotherExampleServiceFindUserArgs_Freezed(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceFindUserArgs_Freezed
    implements _AnotherExampleServiceFindUserArgs_Freezed {
  _$_AnotherExampleServiceFindUserArgs_Freezed({@required this.s})
      : assert(s != null);

  factory _$_AnotherExampleServiceFindUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceFindUserArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString() {
    return '_AnotherExampleServiceFindUserArgs(s: $s)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceFindUserArgs_Freezed &&
            (identical(other.s, s) ||
                const DeepCollectionEquality().equals(other.s, s)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(s);

  @override
  _$AnotherExampleServiceFindUserArgs_FreezedCopyWith<
          _AnotherExampleServiceFindUserArgs_Freezed>
      get copyWith => __$AnotherExampleServiceFindUserArgs_FreezedCopyWithImpl<
          _AnotherExampleServiceFindUserArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceFindUserArgs_FreezedToJson(this);
  }
}

abstract class _AnotherExampleServiceFindUserArgs_Freezed
    implements _AnotherExampleServiceFindUserArgs {
  factory _AnotherExampleServiceFindUserArgs_Freezed(
          {@required SearchFilter s}) =
      _$_AnotherExampleServiceFindUserArgs_Freezed;

  factory _AnotherExampleServiceFindUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceFindUserArgs_Freezed.fromJson;

  @override
  SearchFilter get s;
  @override
  _$AnotherExampleServiceFindUserArgs_FreezedCopyWith<
      _AnotherExampleServiceFindUserArgs_Freezed> get copyWith;
}

AnotherExampleServiceStatusReturn _$AnotherExampleServiceStatusReturnFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceStatusReturn.fromJson(json);
}

class _$AnotherExampleServiceStatusReturnTearOff {
  const _$AnotherExampleServiceStatusReturnTearOff();

  _AnotherExampleServiceStatusReturn call({@required bool status}) {
    return _AnotherExampleServiceStatusReturn(
      status: status,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceStatusReturn =
    _$AnotherExampleServiceStatusReturnTearOff();

mixin _$AnotherExampleServiceStatusReturn {
  bool get status;

  Map<String, dynamic> toJson();
  $AnotherExampleServiceStatusReturnCopyWith<AnotherExampleServiceStatusReturn>
      get copyWith;
}

abstract class $AnotherExampleServiceStatusReturnCopyWith<$Res> {
  factory $AnotherExampleServiceStatusReturnCopyWith(
          AnotherExampleServiceStatusReturn value,
          $Res Function(AnotherExampleServiceStatusReturn) then) =
      _$AnotherExampleServiceStatusReturnCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$AnotherExampleServiceStatusReturnCopyWithImpl<$Res>
    implements $AnotherExampleServiceStatusReturnCopyWith<$Res> {
  _$AnotherExampleServiceStatusReturnCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceStatusReturn _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceStatusReturn) _then;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_value.copyWith(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

abstract class _$AnotherExampleServiceStatusReturnCopyWith<$Res>
    implements $AnotherExampleServiceStatusReturnCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusReturnCopyWith(
          _AnotherExampleServiceStatusReturn value,
          $Res Function(_AnotherExampleServiceStatusReturn) then) =
      __$AnotherExampleServiceStatusReturnCopyWithImpl<$Res>;
  @override
  $Res call({bool status});
}

class __$AnotherExampleServiceStatusReturnCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusReturnCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusReturnCopyWith<$Res> {
  __$AnotherExampleServiceStatusReturnCopyWithImpl(
      _AnotherExampleServiceStatusReturn _value,
      $Res Function(_AnotherExampleServiceStatusReturn) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceStatusReturn));

  @override
  _AnotherExampleServiceStatusReturn get _value =>
      super._value as _AnotherExampleServiceStatusReturn;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_AnotherExampleServiceStatusReturn(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceStatusReturn
    implements _AnotherExampleServiceStatusReturn {
  _$_AnotherExampleServiceStatusReturn({@required this.status})
      : assert(status != null);

  factory _$_AnotherExampleServiceStatusReturn.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusReturnFromJson(json);

  @override
  final bool status;

  @override
  String toString() {
    return 'AnotherExampleServiceStatusReturn(status: $status)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusReturn &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$AnotherExampleServiceStatusReturnCopyWith<
          _AnotherExampleServiceStatusReturn>
      get copyWith => __$AnotherExampleServiceStatusReturnCopyWithImpl<
          _AnotherExampleServiceStatusReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusReturnToJson(this);
  }
}

abstract class _AnotherExampleServiceStatusReturn
    implements AnotherExampleServiceStatusReturn {
  factory _AnotherExampleServiceStatusReturn({@required bool status}) =
      _$_AnotherExampleServiceStatusReturn;

  factory _AnotherExampleServiceStatusReturn.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusReturn.fromJson;

  @override
  bool get status;
  @override
  _$AnotherExampleServiceStatusReturnCopyWith<
      _AnotherExampleServiceStatusReturn> get copyWith;
}

AnotherExampleServiceVersionReturn _$AnotherExampleServiceVersionReturnFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceVersionReturn.fromJson(json);
}

class _$AnotherExampleServiceVersionReturnTearOff {
  const _$AnotherExampleServiceVersionReturnTearOff();

  _AnotherExampleServiceVersionReturn call({@required Version version}) {
    return _AnotherExampleServiceVersionReturn(
      version: version,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceVersionReturn =
    _$AnotherExampleServiceVersionReturnTearOff();

mixin _$AnotherExampleServiceVersionReturn {
  Version get version;

  Map<String, dynamic> toJson();
  $AnotherExampleServiceVersionReturnCopyWith<
      AnotherExampleServiceVersionReturn> get copyWith;
}

abstract class $AnotherExampleServiceVersionReturnCopyWith<$Res> {
  factory $AnotherExampleServiceVersionReturnCopyWith(
          AnotherExampleServiceVersionReturn value,
          $Res Function(AnotherExampleServiceVersionReturn) then) =
      _$AnotherExampleServiceVersionReturnCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class _$AnotherExampleServiceVersionReturnCopyWithImpl<$Res>
    implements $AnotherExampleServiceVersionReturnCopyWith<$Res> {
  _$AnotherExampleServiceVersionReturnCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceVersionReturn _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceVersionReturn) _then;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_value.copyWith(
      version: version == freezed ? _value.version : version as Version,
    ));
  }

  @override
  $VersionCopyWith<$Res> get version {
    if (_value.version == null) {
      return null;
    }
    return $VersionCopyWith<$Res>(_value.version, (value) {
      return _then(_value.copyWith(version: value));
    });
  }
}

abstract class _$AnotherExampleServiceVersionReturnCopyWith<$Res>
    implements $AnotherExampleServiceVersionReturnCopyWith<$Res> {
  factory _$AnotherExampleServiceVersionReturnCopyWith(
          _AnotherExampleServiceVersionReturn value,
          $Res Function(_AnotherExampleServiceVersionReturn) then) =
      __$AnotherExampleServiceVersionReturnCopyWithImpl<$Res>;
  @override
  $Res call({Version version});

  @override
  $VersionCopyWith<$Res> get version;
}

class __$AnotherExampleServiceVersionReturnCopyWithImpl<$Res>
    extends _$AnotherExampleServiceVersionReturnCopyWithImpl<$Res>
    implements _$AnotherExampleServiceVersionReturnCopyWith<$Res> {
  __$AnotherExampleServiceVersionReturnCopyWithImpl(
      _AnotherExampleServiceVersionReturn _value,
      $Res Function(_AnotherExampleServiceVersionReturn) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceVersionReturn));

  @override
  _AnotherExampleServiceVersionReturn get _value =>
      super._value as _AnotherExampleServiceVersionReturn;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_AnotherExampleServiceVersionReturn(
      version: version == freezed ? _value.version : version as Version,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceVersionReturn
    implements _AnotherExampleServiceVersionReturn {
  _$_AnotherExampleServiceVersionReturn({@required this.version})
      : assert(version != null);

  factory _$_AnotherExampleServiceVersionReturn.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceVersionReturnFromJson(json);

  @override
  final Version version;

  @override
  String toString() {
    return 'AnotherExampleServiceVersionReturn(version: $version)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceVersionReturn &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$AnotherExampleServiceVersionReturnCopyWith<
          _AnotherExampleServiceVersionReturn>
      get copyWith => __$AnotherExampleServiceVersionReturnCopyWithImpl<
          _AnotherExampleServiceVersionReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceVersionReturnToJson(this);
  }
}

abstract class _AnotherExampleServiceVersionReturn
    implements AnotherExampleServiceVersionReturn {
  factory _AnotherExampleServiceVersionReturn({@required Version version}) =
      _$_AnotherExampleServiceVersionReturn;

  factory _AnotherExampleServiceVersionReturn.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceVersionReturn.fromJson;

  @override
  Version get version;
  @override
  _$AnotherExampleServiceVersionReturnCopyWith<
      _AnotherExampleServiceVersionReturn> get copyWith;
}

AnotherExampleServiceGetUserReturn _$AnotherExampleServiceGetUserReturnFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceGetUserReturn.fromJson(json);
}

class _$AnotherExampleServiceGetUserReturnTearOff {
  const _$AnotherExampleServiceGetUserReturnTearOff();

  _AnotherExampleServiceGetUserReturn call(
      {@required int code, @required User user}) {
    return _AnotherExampleServiceGetUserReturn(
      code: code,
      user: user,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceGetUserReturn =
    _$AnotherExampleServiceGetUserReturnTearOff();

mixin _$AnotherExampleServiceGetUserReturn {
  int get code;
  User get user;

  Map<String, dynamic> toJson();
  $AnotherExampleServiceGetUserReturnCopyWith<
      AnotherExampleServiceGetUserReturn> get copyWith;
}

abstract class $AnotherExampleServiceGetUserReturnCopyWith<$Res> {
  factory $AnotherExampleServiceGetUserReturnCopyWith(
          AnotherExampleServiceGetUserReturn value,
          $Res Function(AnotherExampleServiceGetUserReturn) then) =
      _$AnotherExampleServiceGetUserReturnCopyWithImpl<$Res>;
  $Res call({int code, User user});

  $UserCopyWith<$Res> get user;
}

class _$AnotherExampleServiceGetUserReturnCopyWithImpl<$Res>
    implements $AnotherExampleServiceGetUserReturnCopyWith<$Res> {
  _$AnotherExampleServiceGetUserReturnCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceGetUserReturn _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceGetUserReturn) _then;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      code: code == freezed ? _value.code : code as int,
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

abstract class _$AnotherExampleServiceGetUserReturnCopyWith<$Res>
    implements $AnotherExampleServiceGetUserReturnCopyWith<$Res> {
  factory _$AnotherExampleServiceGetUserReturnCopyWith(
          _AnotherExampleServiceGetUserReturn value,
          $Res Function(_AnotherExampleServiceGetUserReturn) then) =
      __$AnotherExampleServiceGetUserReturnCopyWithImpl<$Res>;
  @override
  $Res call({int code, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$AnotherExampleServiceGetUserReturnCopyWithImpl<$Res>
    extends _$AnotherExampleServiceGetUserReturnCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetUserReturnCopyWith<$Res> {
  __$AnotherExampleServiceGetUserReturnCopyWithImpl(
      _AnotherExampleServiceGetUserReturn _value,
      $Res Function(_AnotherExampleServiceGetUserReturn) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceGetUserReturn));

  @override
  _AnotherExampleServiceGetUserReturn get _value =>
      super._value as _AnotherExampleServiceGetUserReturn;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_AnotherExampleServiceGetUserReturn(
      code: code == freezed ? _value.code : code as int,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceGetUserReturn
    implements _AnotherExampleServiceGetUserReturn {
  _$_AnotherExampleServiceGetUserReturn(
      {@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_AnotherExampleServiceGetUserReturn.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceGetUserReturnFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString() {
    return 'AnotherExampleServiceGetUserReturn(code: $code, user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceGetUserReturn &&
            (identical(other.code, code) ||
                const DeepCollectionEquality().equals(other.code, code)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(code) ^
      const DeepCollectionEquality().hash(user);

  @override
  _$AnotherExampleServiceGetUserReturnCopyWith<
          _AnotherExampleServiceGetUserReturn>
      get copyWith => __$AnotherExampleServiceGetUserReturnCopyWithImpl<
          _AnotherExampleServiceGetUserReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceGetUserReturnToJson(this);
  }
}

abstract class _AnotherExampleServiceGetUserReturn
    implements AnotherExampleServiceGetUserReturn {
  factory _AnotherExampleServiceGetUserReturn(
      {@required int code,
      @required User user}) = _$_AnotherExampleServiceGetUserReturn;

  factory _AnotherExampleServiceGetUserReturn.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceGetUserReturn.fromJson;

  @override
  int get code;
  @override
  User get user;
  @override
  _$AnotherExampleServiceGetUserReturnCopyWith<
      _AnotherExampleServiceGetUserReturn> get copyWith;
}

AnotherExampleServiceFindUserReturn
    _$AnotherExampleServiceFindUserReturnFromJson(Map<String, dynamic> json) {
  return _AnotherExampleServiceFindUserReturn.fromJson(json);
}

class _$AnotherExampleServiceFindUserReturnTearOff {
  const _$AnotherExampleServiceFindUserReturnTearOff();

  _AnotherExampleServiceFindUserReturn call(
      {@required String name, @required User user}) {
    return _AnotherExampleServiceFindUserReturn(
      name: name,
      user: user,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceFindUserReturn =
    _$AnotherExampleServiceFindUserReturnTearOff();

mixin _$AnotherExampleServiceFindUserReturn {
  String get name;
  User get user;

  Map<String, dynamic> toJson();
  $AnotherExampleServiceFindUserReturnCopyWith<
      AnotherExampleServiceFindUserReturn> get copyWith;
}

abstract class $AnotherExampleServiceFindUserReturnCopyWith<$Res> {
  factory $AnotherExampleServiceFindUserReturnCopyWith(
          AnotherExampleServiceFindUserReturn value,
          $Res Function(AnotherExampleServiceFindUserReturn) then) =
      _$AnotherExampleServiceFindUserReturnCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$AnotherExampleServiceFindUserReturnCopyWithImpl<$Res>
    implements $AnotherExampleServiceFindUserReturnCopyWith<$Res> {
  _$AnotherExampleServiceFindUserReturnCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceFindUserReturn _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceFindUserReturn) _then;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_value.copyWith(
      name: name == freezed ? _value.name : name as String,
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

abstract class _$AnotherExampleServiceFindUserReturnCopyWith<$Res>
    implements $AnotherExampleServiceFindUserReturnCopyWith<$Res> {
  factory _$AnotherExampleServiceFindUserReturnCopyWith(
          _AnotherExampleServiceFindUserReturn value,
          $Res Function(_AnotherExampleServiceFindUserReturn) then) =
      __$AnotherExampleServiceFindUserReturnCopyWithImpl<$Res>;
  @override
  $Res call({String name, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$AnotherExampleServiceFindUserReturnCopyWithImpl<$Res>
    extends _$AnotherExampleServiceFindUserReturnCopyWithImpl<$Res>
    implements _$AnotherExampleServiceFindUserReturnCopyWith<$Res> {
  __$AnotherExampleServiceFindUserReturnCopyWithImpl(
      _AnotherExampleServiceFindUserReturn _value,
      $Res Function(_AnotherExampleServiceFindUserReturn) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceFindUserReturn));

  @override
  _AnotherExampleServiceFindUserReturn get _value =>
      super._value as _AnotherExampleServiceFindUserReturn;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_AnotherExampleServiceFindUserReturn(
      name: name == freezed ? _value.name : name as String,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceFindUserReturn
    implements _AnotherExampleServiceFindUserReturn {
  _$_AnotherExampleServiceFindUserReturn(
      {@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_AnotherExampleServiceFindUserReturn.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceFindUserReturnFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString() {
    return 'AnotherExampleServiceFindUserReturn(name: $name, user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceFindUserReturn &&
            (identical(other.name, name) ||
                const DeepCollectionEquality().equals(other.name, name)) &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(name) ^
      const DeepCollectionEquality().hash(user);

  @override
  _$AnotherExampleServiceFindUserReturnCopyWith<
          _AnotherExampleServiceFindUserReturn>
      get copyWith => __$AnotherExampleServiceFindUserReturnCopyWithImpl<
          _AnotherExampleServiceFindUserReturn>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceFindUserReturnToJson(this);
  }
}

abstract class _AnotherExampleServiceFindUserReturn
    implements AnotherExampleServiceFindUserReturn {
  factory _AnotherExampleServiceFindUserReturn(
      {@required String name,
      @required User user}) = _$_AnotherExampleServiceFindUserReturn;

  factory _AnotherExampleServiceFindUserReturn.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceFindUserReturn.fromJson;

  @override
  String get name;
  @override
  User get user;
  @override
  _$AnotherExampleServiceFindUserReturnCopyWith<
      _AnotherExampleServiceFindUserReturn> get copyWith;
}
