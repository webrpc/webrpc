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

@JsonSerializable(explicitToJson: true)
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

@JsonSerializable(explicitToJson: true)
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

  _SearchFilter call({@required int id}) {
    return _SearchFilter(
      id: id,
    );
  }
}

// ignore: unused_element
const $SearchFilter = _$SearchFilterTearOff();

mixin _$SearchFilter {
  int get id;

  Map<String, dynamic> toJson();
  $SearchFilterCopyWith<SearchFilter> get copyWith;
}

abstract class $SearchFilterCopyWith<$Res> {
  factory $SearchFilterCopyWith(
          SearchFilter value, $Res Function(SearchFilter) then) =
      _$SearchFilterCopyWithImpl<$Res>;
  $Res call({int id});
}

class _$SearchFilterCopyWithImpl<$Res> implements $SearchFilterCopyWith<$Res> {
  _$SearchFilterCopyWithImpl(this._value, this._then);

  final SearchFilter _value;
  // ignore: unused_field
  final $Res Function(SearchFilter) _then;

  @override
  $Res call({
    Object id = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as int,
    ));
  }
}

abstract class _$SearchFilterCopyWith<$Res>
    implements $SearchFilterCopyWith<$Res> {
  factory _$SearchFilterCopyWith(
          _SearchFilter value, $Res Function(_SearchFilter) then) =
      __$SearchFilterCopyWithImpl<$Res>;
  @override
  $Res call({int id});
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
    Object id = freezed,
  }) {
    return _then(_SearchFilter(
      id: id == freezed ? _value.id : id as int,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_SearchFilter implements _SearchFilter {
  _$_SearchFilter({@required this.id}) : assert(id != null);

  factory _$_SearchFilter.fromJson(Map<String, dynamic> json) =>
      _$_$_SearchFilterFromJson(json);

  @override
  final int id;

  @override
  String toString() {
    return 'SearchFilter(id: $id)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _SearchFilter &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(id);

  @override
  _$SearchFilterCopyWith<_SearchFilter> get copyWith =>
      __$SearchFilterCopyWithImpl<_SearchFilter>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_SearchFilterToJson(this);
  }
}

abstract class _SearchFilter implements SearchFilter {
  factory _SearchFilter({@required int id}) = _$_SearchFilter;

  factory _SearchFilter.fromJson(Map<String, dynamic> json) =
      _$_SearchFilter.fromJson;

  @override
  int get id;
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

@JsonSerializable(explicitToJson: true)
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

@JsonSerializable(explicitToJson: true)
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

_UpdateNameArgs _$_UpdateNameArgsFromJson(Map<String, dynamic> json) {
  return _UpdateNameArgs_Freezed.fromJson(json);
}

class _$_UpdateNameArgsTearOff {
  const _$_UpdateNameArgsTearOff();

  _UpdateNameArgs_Freezed call({@required int id, @required String username}) {
    return _UpdateNameArgs_Freezed(
      id: id,
      username: username,
    );
  }
}

// ignore: unused_element
const _$UpdateNameArgs = _$_UpdateNameArgsTearOff();

mixin _$_UpdateNameArgs {
  int get id;
  String get username;

  Map<String, dynamic> toJson();
  _$UpdateNameArgsCopyWith<_UpdateNameArgs> get copyWith;
}

abstract class _$UpdateNameArgsCopyWith<$Res> {
  factory _$UpdateNameArgsCopyWith(
          _UpdateNameArgs value, $Res Function(_UpdateNameArgs) then) =
      __$UpdateNameArgsCopyWithImpl<$Res>;
  $Res call({int id, String username});
}

class __$UpdateNameArgsCopyWithImpl<$Res>
    implements _$UpdateNameArgsCopyWith<$Res> {
  __$UpdateNameArgsCopyWithImpl(this._value, this._then);

  final _UpdateNameArgs _value;
  // ignore: unused_field
  final $Res Function(_UpdateNameArgs) _then;

  @override
  $Res call({
    Object id = freezed,
    Object username = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as int,
      username: username == freezed ? _value.username : username as String,
    ));
  }
}

abstract class _$UpdateNameArgs_FreezedCopyWith<$Res>
    implements _$UpdateNameArgsCopyWith<$Res> {
  factory _$UpdateNameArgs_FreezedCopyWith(_UpdateNameArgs_Freezed value,
          $Res Function(_UpdateNameArgs_Freezed) then) =
      __$UpdateNameArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({int id, String username});
}

class __$UpdateNameArgs_FreezedCopyWithImpl<$Res>
    extends __$UpdateNameArgsCopyWithImpl<$Res>
    implements _$UpdateNameArgs_FreezedCopyWith<$Res> {
  __$UpdateNameArgs_FreezedCopyWithImpl(_UpdateNameArgs_Freezed _value,
      $Res Function(_UpdateNameArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _UpdateNameArgs_Freezed));

  @override
  _UpdateNameArgs_Freezed get _value => super._value as _UpdateNameArgs_Freezed;

  @override
  $Res call({
    Object id = freezed,
    Object username = freezed,
  }) {
    return _then(_UpdateNameArgs_Freezed(
      id: id == freezed ? _value.id : id as int,
      username: username == freezed ? _value.username : username as String,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_UpdateNameArgs_Freezed implements _UpdateNameArgs_Freezed {
  _$_UpdateNameArgs_Freezed({@required this.id, @required this.username})
      : assert(id != null),
        assert(username != null);

  factory _$_UpdateNameArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_UpdateNameArgs_FreezedFromJson(json);

  @override
  final int id;
  @override
  final String username;

  @override
  String toString() {
    return '_UpdateNameArgs(id: $id, username: $username)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _UpdateNameArgs_Freezed &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.username, username) ||
                const DeepCollectionEquality()
                    .equals(other.username, username)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(username);

  @override
  _$UpdateNameArgs_FreezedCopyWith<_UpdateNameArgs_Freezed> get copyWith =>
      __$UpdateNameArgs_FreezedCopyWithImpl<_UpdateNameArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_UpdateNameArgs_FreezedToJson(this);
  }
}

abstract class _UpdateNameArgs_Freezed implements _UpdateNameArgs {
  factory _UpdateNameArgs_Freezed(
      {@required int id,
      @required String username}) = _$_UpdateNameArgs_Freezed;

  factory _UpdateNameArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_UpdateNameArgs_Freezed.fromJson;

  @override
  int get id;
  @override
  String get username;
  @override
  _$UpdateNameArgs_FreezedCopyWith<_UpdateNameArgs_Freezed> get copyWith;
}

_FindUserByIdArgs _$_FindUserByIdArgsFromJson(Map<String, dynamic> json) {
  return _FindUserByIdArgs_Freezed.fromJson(json);
}

class _$_FindUserByIdArgsTearOff {
  const _$_FindUserByIdArgsTearOff();

  _FindUserByIdArgs_Freezed call({@required SearchFilter s}) {
    return _FindUserByIdArgs_Freezed(
      s: s,
    );
  }
}

// ignore: unused_element
const _$FindUserByIdArgs = _$_FindUserByIdArgsTearOff();

mixin _$_FindUserByIdArgs {
  SearchFilter get s;

  Map<String, dynamic> toJson();
  _$FindUserByIdArgsCopyWith<_FindUserByIdArgs> get copyWith;
}

abstract class _$FindUserByIdArgsCopyWith<$Res> {
  factory _$FindUserByIdArgsCopyWith(
          _FindUserByIdArgs value, $Res Function(_FindUserByIdArgs) then) =
      __$FindUserByIdArgsCopyWithImpl<$Res>;
  $Res call({SearchFilter s});

  $SearchFilterCopyWith<$Res> get s;
}

class __$FindUserByIdArgsCopyWithImpl<$Res>
    implements _$FindUserByIdArgsCopyWith<$Res> {
  __$FindUserByIdArgsCopyWithImpl(this._value, this._then);

  final _FindUserByIdArgs _value;
  // ignore: unused_field
  final $Res Function(_FindUserByIdArgs) _then;

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

abstract class _$FindUserByIdArgs_FreezedCopyWith<$Res>
    implements _$FindUserByIdArgsCopyWith<$Res> {
  factory _$FindUserByIdArgs_FreezedCopyWith(_FindUserByIdArgs_Freezed value,
          $Res Function(_FindUserByIdArgs_Freezed) then) =
      __$FindUserByIdArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({SearchFilter s});

  @override
  $SearchFilterCopyWith<$Res> get s;
}

class __$FindUserByIdArgs_FreezedCopyWithImpl<$Res>
    extends __$FindUserByIdArgsCopyWithImpl<$Res>
    implements _$FindUserByIdArgs_FreezedCopyWith<$Res> {
  __$FindUserByIdArgs_FreezedCopyWithImpl(_FindUserByIdArgs_Freezed _value,
      $Res Function(_FindUserByIdArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _FindUserByIdArgs_Freezed));

  @override
  _FindUserByIdArgs_Freezed get _value =>
      super._value as _FindUserByIdArgs_Freezed;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_FindUserByIdArgs_Freezed(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_FindUserByIdArgs_Freezed implements _FindUserByIdArgs_Freezed {
  _$_FindUserByIdArgs_Freezed({@required this.s}) : assert(s != null);

  factory _$_FindUserByIdArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserByIdArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString() {
    return '_FindUserByIdArgs(s: $s)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserByIdArgs_Freezed &&
            (identical(other.s, s) ||
                const DeepCollectionEquality().equals(other.s, s)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(s);

  @override
  _$FindUserByIdArgs_FreezedCopyWith<_FindUserByIdArgs_Freezed> get copyWith =>
      __$FindUserByIdArgs_FreezedCopyWithImpl<_FindUserByIdArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserByIdArgs_FreezedToJson(this);
  }
}

abstract class _FindUserByIdArgs_Freezed implements _FindUserByIdArgs {
  factory _FindUserByIdArgs_Freezed({@required SearchFilter s}) =
      _$_FindUserByIdArgs_Freezed;

  factory _FindUserByIdArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_FindUserByIdArgs_Freezed.fromJson;

  @override
  SearchFilter get s;
  @override
  _$FindUserByIdArgs_FreezedCopyWith<_FindUserByIdArgs_Freezed> get copyWith;
}

_AddUserArgs _$_AddUserArgsFromJson(Map<String, dynamic> json) {
  return _AddUserArgs_Freezed.fromJson(json);
}

class _$_AddUserArgsTearOff {
  const _$_AddUserArgsTearOff();

  _AddUserArgs_Freezed call({@required User user}) {
    return _AddUserArgs_Freezed(
      user: user,
    );
  }
}

// ignore: unused_element
const _$AddUserArgs = _$_AddUserArgsTearOff();

mixin _$_AddUserArgs {
  User get user;

  Map<String, dynamic> toJson();
  _$AddUserArgsCopyWith<_AddUserArgs> get copyWith;
}

abstract class _$AddUserArgsCopyWith<$Res> {
  factory _$AddUserArgsCopyWith(
          _AddUserArgs value, $Res Function(_AddUserArgs) then) =
      __$AddUserArgsCopyWithImpl<$Res>;
  $Res call({User user});

  $UserCopyWith<$Res> get user;
}

class __$AddUserArgsCopyWithImpl<$Res> implements _$AddUserArgsCopyWith<$Res> {
  __$AddUserArgsCopyWithImpl(this._value, this._then);

  final _AddUserArgs _value;
  // ignore: unused_field
  final $Res Function(_AddUserArgs) _then;

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

abstract class _$AddUserArgs_FreezedCopyWith<$Res>
    implements _$AddUserArgsCopyWith<$Res> {
  factory _$AddUserArgs_FreezedCopyWith(_AddUserArgs_Freezed value,
          $Res Function(_AddUserArgs_Freezed) then) =
      __$AddUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$AddUserArgs_FreezedCopyWithImpl<$Res>
    extends __$AddUserArgsCopyWithImpl<$Res>
    implements _$AddUserArgs_FreezedCopyWith<$Res> {
  __$AddUserArgs_FreezedCopyWithImpl(
      _AddUserArgs_Freezed _value, $Res Function(_AddUserArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _AddUserArgs_Freezed));

  @override
  _AddUserArgs_Freezed get _value => super._value as _AddUserArgs_Freezed;

  @override
  $Res call({
    Object user = freezed,
  }) {
    return _then(_AddUserArgs_Freezed(
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_AddUserArgs_Freezed implements _AddUserArgs_Freezed {
  _$_AddUserArgs_Freezed({@required this.user}) : assert(user != null);

  factory _$_AddUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_AddUserArgs_FreezedFromJson(json);

  @override
  final User user;

  @override
  String toString() {
    return '_AddUserArgs(user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AddUserArgs_Freezed &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(user);

  @override
  _$AddUserArgs_FreezedCopyWith<_AddUserArgs_Freezed> get copyWith =>
      __$AddUserArgs_FreezedCopyWithImpl<_AddUserArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AddUserArgs_FreezedToJson(this);
  }
}

abstract class _AddUserArgs_Freezed implements _AddUserArgs {
  factory _AddUserArgs_Freezed({@required User user}) = _$_AddUserArgs_Freezed;

  factory _AddUserArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_AddUserArgs_Freezed.fromJson;

  @override
  User get user;
  @override
  _$AddUserArgs_FreezedCopyWith<_AddUserArgs_Freezed> get copyWith;
}

_DeleteUserArgs _$_DeleteUserArgsFromJson(Map<String, dynamic> json) {
  return _DeleteUserArgs_Freezed.fromJson(json);
}

class _$_DeleteUserArgsTearOff {
  const _$_DeleteUserArgsTearOff();

  _DeleteUserArgs_Freezed call({@required int id}) {
    return _DeleteUserArgs_Freezed(
      id: id,
    );
  }
}

// ignore: unused_element
const _$DeleteUserArgs = _$_DeleteUserArgsTearOff();

mixin _$_DeleteUserArgs {
  int get id;

  Map<String, dynamic> toJson();
  _$DeleteUserArgsCopyWith<_DeleteUserArgs> get copyWith;
}

abstract class _$DeleteUserArgsCopyWith<$Res> {
  factory _$DeleteUserArgsCopyWith(
          _DeleteUserArgs value, $Res Function(_DeleteUserArgs) then) =
      __$DeleteUserArgsCopyWithImpl<$Res>;
  $Res call({int id});
}

class __$DeleteUserArgsCopyWithImpl<$Res>
    implements _$DeleteUserArgsCopyWith<$Res> {
  __$DeleteUserArgsCopyWithImpl(this._value, this._then);

  final _DeleteUserArgs _value;
  // ignore: unused_field
  final $Res Function(_DeleteUserArgs) _then;

  @override
  $Res call({
    Object id = freezed,
  }) {
    return _then(_value.copyWith(
      id: id == freezed ? _value.id : id as int,
    ));
  }
}

abstract class _$DeleteUserArgs_FreezedCopyWith<$Res>
    implements _$DeleteUserArgsCopyWith<$Res> {
  factory _$DeleteUserArgs_FreezedCopyWith(_DeleteUserArgs_Freezed value,
          $Res Function(_DeleteUserArgs_Freezed) then) =
      __$DeleteUserArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({int id});
}

class __$DeleteUserArgs_FreezedCopyWithImpl<$Res>
    extends __$DeleteUserArgsCopyWithImpl<$Res>
    implements _$DeleteUserArgs_FreezedCopyWith<$Res> {
  __$DeleteUserArgs_FreezedCopyWithImpl(_DeleteUserArgs_Freezed _value,
      $Res Function(_DeleteUserArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _DeleteUserArgs_Freezed));

  @override
  _DeleteUserArgs_Freezed get _value => super._value as _DeleteUserArgs_Freezed;

  @override
  $Res call({
    Object id = freezed,
  }) {
    return _then(_DeleteUserArgs_Freezed(
      id: id == freezed ? _value.id : id as int,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_DeleteUserArgs_Freezed implements _DeleteUserArgs_Freezed {
  _$_DeleteUserArgs_Freezed({@required this.id}) : assert(id != null);

  factory _$_DeleteUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_DeleteUserArgs_FreezedFromJson(json);

  @override
  final int id;

  @override
  String toString() {
    return '_DeleteUserArgs(id: $id)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _DeleteUserArgs_Freezed &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(id);

  @override
  _$DeleteUserArgs_FreezedCopyWith<_DeleteUserArgs_Freezed> get copyWith =>
      __$DeleteUserArgs_FreezedCopyWithImpl<_DeleteUserArgs_Freezed>(
          this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_DeleteUserArgs_FreezedToJson(this);
  }
}

abstract class _DeleteUserArgs_Freezed implements _DeleteUserArgs {
  factory _DeleteUserArgs_Freezed({@required int id}) =
      _$_DeleteUserArgs_Freezed;

  factory _DeleteUserArgs_Freezed.fromJson(Map<String, dynamic> json) =
      _$_DeleteUserArgs_Freezed.fromJson;

  @override
  int get id;
  @override
  _$DeleteUserArgs_FreezedCopyWith<_DeleteUserArgs_Freezed> get copyWith;
}

StatusResult _$StatusResultFromJson(Map<String, dynamic> json) {
  return _StatusResult.fromJson(json);
}

class _$StatusResultTearOff {
  const _$StatusResultTearOff();

  _StatusResult call({@required bool status}) {
    return _StatusResult(
      status: status,
    );
  }
}

// ignore: unused_element
const $StatusResult = _$StatusResultTearOff();

mixin _$StatusResult {
  bool get status;

  Map<String, dynamic> toJson();
  $StatusResultCopyWith<StatusResult> get copyWith;
}

abstract class $StatusResultCopyWith<$Res> {
  factory $StatusResultCopyWith(
          StatusResult value, $Res Function(StatusResult) then) =
      _$StatusResultCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$StatusResultCopyWithImpl<$Res> implements $StatusResultCopyWith<$Res> {
  _$StatusResultCopyWithImpl(this._value, this._then);

  final StatusResult _value;
  // ignore: unused_field
  final $Res Function(StatusResult) _then;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_value.copyWith(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

abstract class _$StatusResultCopyWith<$Res>
    implements $StatusResultCopyWith<$Res> {
  factory _$StatusResultCopyWith(
          _StatusResult value, $Res Function(_StatusResult) then) =
      __$StatusResultCopyWithImpl<$Res>;
  @override
  $Res call({bool status});
}

class __$StatusResultCopyWithImpl<$Res> extends _$StatusResultCopyWithImpl<$Res>
    implements _$StatusResultCopyWith<$Res> {
  __$StatusResultCopyWithImpl(
      _StatusResult _value, $Res Function(_StatusResult) _then)
      : super(_value, (v) => _then(v as _StatusResult));

  @override
  _StatusResult get _value => super._value as _StatusResult;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_StatusResult(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_StatusResult implements _StatusResult {
  _$_StatusResult({@required this.status}) : assert(status != null);

  factory _$_StatusResult.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResultFromJson(json);

  @override
  final bool status;

  @override
  String toString() {
    return 'StatusResult(status: $status)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _StatusResult &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$StatusResultCopyWith<_StatusResult> get copyWith =>
      __$StatusResultCopyWithImpl<_StatusResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResultToJson(this);
  }
}

abstract class _StatusResult implements StatusResult {
  factory _StatusResult({@required bool status}) = _$_StatusResult;

  factory _StatusResult.fromJson(Map<String, dynamic> json) =
      _$_StatusResult.fromJson;

  @override
  bool get status;
  @override
  _$StatusResultCopyWith<_StatusResult> get copyWith;
}

VersionResult _$VersionResultFromJson(Map<String, dynamic> json) {
  return _VersionResult.fromJson(json);
}

class _$VersionResultTearOff {
  const _$VersionResultTearOff();

  _VersionResult call({@required Version version}) {
    return _VersionResult(
      version: version,
    );
  }
}

// ignore: unused_element
const $VersionResult = _$VersionResultTearOff();

mixin _$VersionResult {
  Version get version;

  Map<String, dynamic> toJson();
  $VersionResultCopyWith<VersionResult> get copyWith;
}

abstract class $VersionResultCopyWith<$Res> {
  factory $VersionResultCopyWith(
          VersionResult value, $Res Function(VersionResult) then) =
      _$VersionResultCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class _$VersionResultCopyWithImpl<$Res>
    implements $VersionResultCopyWith<$Res> {
  _$VersionResultCopyWithImpl(this._value, this._then);

  final VersionResult _value;
  // ignore: unused_field
  final $Res Function(VersionResult) _then;

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

abstract class _$VersionResultCopyWith<$Res>
    implements $VersionResultCopyWith<$Res> {
  factory _$VersionResultCopyWith(
          _VersionResult value, $Res Function(_VersionResult) then) =
      __$VersionResultCopyWithImpl<$Res>;
  @override
  $Res call({Version version});

  @override
  $VersionCopyWith<$Res> get version;
}

class __$VersionResultCopyWithImpl<$Res>
    extends _$VersionResultCopyWithImpl<$Res>
    implements _$VersionResultCopyWith<$Res> {
  __$VersionResultCopyWithImpl(
      _VersionResult _value, $Res Function(_VersionResult) _then)
      : super(_value, (v) => _then(v as _VersionResult));

  @override
  _VersionResult get _value => super._value as _VersionResult;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_VersionResult(
      version: version == freezed ? _value.version : version as Version,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_VersionResult implements _VersionResult {
  _$_VersionResult({@required this.version}) : assert(version != null);

  factory _$_VersionResult.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResultFromJson(json);

  @override
  final Version version;

  @override
  String toString() {
    return 'VersionResult(version: $version)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _VersionResult &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$VersionResultCopyWith<_VersionResult> get copyWith =>
      __$VersionResultCopyWithImpl<_VersionResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResultToJson(this);
  }
}

abstract class _VersionResult implements VersionResult {
  factory _VersionResult({@required Version version}) = _$_VersionResult;

  factory _VersionResult.fromJson(Map<String, dynamic> json) =
      _$_VersionResult.fromJson;

  @override
  Version get version;
  @override
  _$VersionResultCopyWith<_VersionResult> get copyWith;
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

UpdateNameResult _$UpdateNameResultFromJson(Map<String, dynamic> json) {
  return _UpdateNameResult.fromJson(json);
}

class _$UpdateNameResultTearOff {
  const _$UpdateNameResultTearOff();

  _UpdateNameResult call({@required bool wasUpdated}) {
    return _UpdateNameResult(
      wasUpdated: wasUpdated,
    );
  }
}

// ignore: unused_element
const $UpdateNameResult = _$UpdateNameResultTearOff();

mixin _$UpdateNameResult {
  bool get wasUpdated;

  Map<String, dynamic> toJson();
  $UpdateNameResultCopyWith<UpdateNameResult> get copyWith;
}

abstract class $UpdateNameResultCopyWith<$Res> {
  factory $UpdateNameResultCopyWith(
          UpdateNameResult value, $Res Function(UpdateNameResult) then) =
      _$UpdateNameResultCopyWithImpl<$Res>;
  $Res call({bool wasUpdated});
}

class _$UpdateNameResultCopyWithImpl<$Res>
    implements $UpdateNameResultCopyWith<$Res> {
  _$UpdateNameResultCopyWithImpl(this._value, this._then);

  final UpdateNameResult _value;
  // ignore: unused_field
  final $Res Function(UpdateNameResult) _then;

  @override
  $Res call({
    Object wasUpdated = freezed,
  }) {
    return _then(_value.copyWith(
      wasUpdated:
          wasUpdated == freezed ? _value.wasUpdated : wasUpdated as bool,
    ));
  }
}

abstract class _$UpdateNameResultCopyWith<$Res>
    implements $UpdateNameResultCopyWith<$Res> {
  factory _$UpdateNameResultCopyWith(
          _UpdateNameResult value, $Res Function(_UpdateNameResult) then) =
      __$UpdateNameResultCopyWithImpl<$Res>;
  @override
  $Res call({bool wasUpdated});
}

class __$UpdateNameResultCopyWithImpl<$Res>
    extends _$UpdateNameResultCopyWithImpl<$Res>
    implements _$UpdateNameResultCopyWith<$Res> {
  __$UpdateNameResultCopyWithImpl(
      _UpdateNameResult _value, $Res Function(_UpdateNameResult) _then)
      : super(_value, (v) => _then(v as _UpdateNameResult));

  @override
  _UpdateNameResult get _value => super._value as _UpdateNameResult;

  @override
  $Res call({
    Object wasUpdated = freezed,
  }) {
    return _then(_UpdateNameResult(
      wasUpdated:
          wasUpdated == freezed ? _value.wasUpdated : wasUpdated as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_UpdateNameResult implements _UpdateNameResult {
  _$_UpdateNameResult({@required this.wasUpdated}) : assert(wasUpdated != null);

  factory _$_UpdateNameResult.fromJson(Map<String, dynamic> json) =>
      _$_$_UpdateNameResultFromJson(json);

  @override
  final bool wasUpdated;

  @override
  String toString() {
    return 'UpdateNameResult(wasUpdated: $wasUpdated)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _UpdateNameResult &&
            (identical(other.wasUpdated, wasUpdated) ||
                const DeepCollectionEquality()
                    .equals(other.wasUpdated, wasUpdated)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasUpdated);

  @override
  _$UpdateNameResultCopyWith<_UpdateNameResult> get copyWith =>
      __$UpdateNameResultCopyWithImpl<_UpdateNameResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_UpdateNameResultToJson(this);
  }
}

abstract class _UpdateNameResult implements UpdateNameResult {
  factory _UpdateNameResult({@required bool wasUpdated}) = _$_UpdateNameResult;

  factory _UpdateNameResult.fromJson(Map<String, dynamic> json) =
      _$_UpdateNameResult.fromJson;

  @override
  bool get wasUpdated;
  @override
  _$UpdateNameResultCopyWith<_UpdateNameResult> get copyWith;
}

FindUserByIdResult _$FindUserByIdResultFromJson(Map<String, dynamic> json) {
  return _FindUserByIdResult.fromJson(json);
}

class _$FindUserByIdResultTearOff {
  const _$FindUserByIdResultTearOff();

  _FindUserByIdResult call({@required String name, @required User user}) {
    return _FindUserByIdResult(
      name: name,
      user: user,
    );
  }
}

// ignore: unused_element
const $FindUserByIdResult = _$FindUserByIdResultTearOff();

mixin _$FindUserByIdResult {
  String get name;
  User get user;

  Map<String, dynamic> toJson();
  $FindUserByIdResultCopyWith<FindUserByIdResult> get copyWith;
}

abstract class $FindUserByIdResultCopyWith<$Res> {
  factory $FindUserByIdResultCopyWith(
          FindUserByIdResult value, $Res Function(FindUserByIdResult) then) =
      _$FindUserByIdResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$FindUserByIdResultCopyWithImpl<$Res>
    implements $FindUserByIdResultCopyWith<$Res> {
  _$FindUserByIdResultCopyWithImpl(this._value, this._then);

  final FindUserByIdResult _value;
  // ignore: unused_field
  final $Res Function(FindUserByIdResult) _then;

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

abstract class _$FindUserByIdResultCopyWith<$Res>
    implements $FindUserByIdResultCopyWith<$Res> {
  factory _$FindUserByIdResultCopyWith(
          _FindUserByIdResult value, $Res Function(_FindUserByIdResult) then) =
      __$FindUserByIdResultCopyWithImpl<$Res>;
  @override
  $Res call({String name, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$FindUserByIdResultCopyWithImpl<$Res>
    extends _$FindUserByIdResultCopyWithImpl<$Res>
    implements _$FindUserByIdResultCopyWith<$Res> {
  __$FindUserByIdResultCopyWithImpl(
      _FindUserByIdResult _value, $Res Function(_FindUserByIdResult) _then)
      : super(_value, (v) => _then(v as _FindUserByIdResult));

  @override
  _FindUserByIdResult get _value => super._value as _FindUserByIdResult;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_FindUserByIdResult(
      name: name == freezed ? _value.name : name as String,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_FindUserByIdResult implements _FindUserByIdResult {
  _$_FindUserByIdResult({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_FindUserByIdResult.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserByIdResultFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString() {
    return 'FindUserByIdResult(name: $name, user: $user)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserByIdResult &&
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
  _$FindUserByIdResultCopyWith<_FindUserByIdResult> get copyWith =>
      __$FindUserByIdResultCopyWithImpl<_FindUserByIdResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserByIdResultToJson(this);
  }
}

abstract class _FindUserByIdResult implements FindUserByIdResult {
  factory _FindUserByIdResult({@required String name, @required User user}) =
      _$_FindUserByIdResult;

  factory _FindUserByIdResult.fromJson(Map<String, dynamic> json) =
      _$_FindUserByIdResult.fromJson;

  @override
  String get name;
  @override
  User get user;
  @override
  _$FindUserByIdResultCopyWith<_FindUserByIdResult> get copyWith;
}

AddUserResult _$AddUserResultFromJson(Map<String, dynamic> json) {
  return _AddUserResult.fromJson(json);
}

class _$AddUserResultTearOff {
  const _$AddUserResultTearOff();

  _AddUserResult call({@required bool wasAdded}) {
    return _AddUserResult(
      wasAdded: wasAdded,
    );
  }
}

// ignore: unused_element
const $AddUserResult = _$AddUserResultTearOff();

mixin _$AddUserResult {
  bool get wasAdded;

  Map<String, dynamic> toJson();
  $AddUserResultCopyWith<AddUserResult> get copyWith;
}

abstract class $AddUserResultCopyWith<$Res> {
  factory $AddUserResultCopyWith(
          AddUserResult value, $Res Function(AddUserResult) then) =
      _$AddUserResultCopyWithImpl<$Res>;
  $Res call({bool wasAdded});
}

class _$AddUserResultCopyWithImpl<$Res>
    implements $AddUserResultCopyWith<$Res> {
  _$AddUserResultCopyWithImpl(this._value, this._then);

  final AddUserResult _value;
  // ignore: unused_field
  final $Res Function(AddUserResult) _then;

  @override
  $Res call({
    Object wasAdded = freezed,
  }) {
    return _then(_value.copyWith(
      wasAdded: wasAdded == freezed ? _value.wasAdded : wasAdded as bool,
    ));
  }
}

abstract class _$AddUserResultCopyWith<$Res>
    implements $AddUserResultCopyWith<$Res> {
  factory _$AddUserResultCopyWith(
          _AddUserResult value, $Res Function(_AddUserResult) then) =
      __$AddUserResultCopyWithImpl<$Res>;
  @override
  $Res call({bool wasAdded});
}

class __$AddUserResultCopyWithImpl<$Res>
    extends _$AddUserResultCopyWithImpl<$Res>
    implements _$AddUserResultCopyWith<$Res> {
  __$AddUserResultCopyWithImpl(
      _AddUserResult _value, $Res Function(_AddUserResult) _then)
      : super(_value, (v) => _then(v as _AddUserResult));

  @override
  _AddUserResult get _value => super._value as _AddUserResult;

  @override
  $Res call({
    Object wasAdded = freezed,
  }) {
    return _then(_AddUserResult(
      wasAdded: wasAdded == freezed ? _value.wasAdded : wasAdded as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_AddUserResult implements _AddUserResult {
  _$_AddUserResult({@required this.wasAdded}) : assert(wasAdded != null);

  factory _$_AddUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$_AddUserResultFromJson(json);

  @override
  final bool wasAdded;

  @override
  String toString() {
    return 'AddUserResult(wasAdded: $wasAdded)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AddUserResult &&
            (identical(other.wasAdded, wasAdded) ||
                const DeepCollectionEquality()
                    .equals(other.wasAdded, wasAdded)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasAdded);

  @override
  _$AddUserResultCopyWith<_AddUserResult> get copyWith =>
      __$AddUserResultCopyWithImpl<_AddUserResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AddUserResultToJson(this);
  }
}

abstract class _AddUserResult implements AddUserResult {
  factory _AddUserResult({@required bool wasAdded}) = _$_AddUserResult;

  factory _AddUserResult.fromJson(Map<String, dynamic> json) =
      _$_AddUserResult.fromJson;

  @override
  bool get wasAdded;
  @override
  _$AddUserResultCopyWith<_AddUserResult> get copyWith;
}

ListUsersResult _$ListUsersResultFromJson(Map<String, dynamic> json) {
  return _ListUsersResult.fromJson(json);
}

class _$ListUsersResultTearOff {
  const _$ListUsersResultTearOff();

  _ListUsersResult call({@required List<User> users}) {
    return _ListUsersResult(
      users: users,
    );
  }
}

// ignore: unused_element
const $ListUsersResult = _$ListUsersResultTearOff();

mixin _$ListUsersResult {
  List<User> get users;

  Map<String, dynamic> toJson();
  $ListUsersResultCopyWith<ListUsersResult> get copyWith;
}

abstract class $ListUsersResultCopyWith<$Res> {
  factory $ListUsersResultCopyWith(
          ListUsersResult value, $Res Function(ListUsersResult) then) =
      _$ListUsersResultCopyWithImpl<$Res>;
  $Res call({List<User> users});
}

class _$ListUsersResultCopyWithImpl<$Res>
    implements $ListUsersResultCopyWith<$Res> {
  _$ListUsersResultCopyWithImpl(this._value, this._then);

  final ListUsersResult _value;
  // ignore: unused_field
  final $Res Function(ListUsersResult) _then;

  @override
  $Res call({
    Object users = freezed,
  }) {
    return _then(_value.copyWith(
      users: users == freezed ? _value.users : users as List<User>,
    ));
  }
}

abstract class _$ListUsersResultCopyWith<$Res>
    implements $ListUsersResultCopyWith<$Res> {
  factory _$ListUsersResultCopyWith(
          _ListUsersResult value, $Res Function(_ListUsersResult) then) =
      __$ListUsersResultCopyWithImpl<$Res>;
  @override
  $Res call({List<User> users});
}

class __$ListUsersResultCopyWithImpl<$Res>
    extends _$ListUsersResultCopyWithImpl<$Res>
    implements _$ListUsersResultCopyWith<$Res> {
  __$ListUsersResultCopyWithImpl(
      _ListUsersResult _value, $Res Function(_ListUsersResult) _then)
      : super(_value, (v) => _then(v as _ListUsersResult));

  @override
  _ListUsersResult get _value => super._value as _ListUsersResult;

  @override
  $Res call({
    Object users = freezed,
  }) {
    return _then(_ListUsersResult(
      users: users == freezed ? _value.users : users as List<User>,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_ListUsersResult implements _ListUsersResult {
  _$_ListUsersResult({@required this.users}) : assert(users != null);

  factory _$_ListUsersResult.fromJson(Map<String, dynamic> json) =>
      _$_$_ListUsersResultFromJson(json);

  @override
  final List<User> users;

  @override
  String toString() {
    return 'ListUsersResult(users: $users)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ListUsersResult &&
            (identical(other.users, users) ||
                const DeepCollectionEquality().equals(other.users, users)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(users);

  @override
  _$ListUsersResultCopyWith<_ListUsersResult> get copyWith =>
      __$ListUsersResultCopyWithImpl<_ListUsersResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ListUsersResultToJson(this);
  }
}

abstract class _ListUsersResult implements ListUsersResult {
  factory _ListUsersResult({@required List<User> users}) = _$_ListUsersResult;

  factory _ListUsersResult.fromJson(Map<String, dynamic> json) =
      _$_ListUsersResult.fromJson;

  @override
  List<User> get users;
  @override
  _$ListUsersResultCopyWith<_ListUsersResult> get copyWith;
}

DeleteUserResult _$DeleteUserResultFromJson(Map<String, dynamic> json) {
  return _DeleteUserResult.fromJson(json);
}

class _$DeleteUserResultTearOff {
  const _$DeleteUserResultTearOff();

  _DeleteUserResult call({@required bool wasDeleted}) {
    return _DeleteUserResult(
      wasDeleted: wasDeleted,
    );
  }
}

// ignore: unused_element
const $DeleteUserResult = _$DeleteUserResultTearOff();

mixin _$DeleteUserResult {
  bool get wasDeleted;

  Map<String, dynamic> toJson();
  $DeleteUserResultCopyWith<DeleteUserResult> get copyWith;
}

abstract class $DeleteUserResultCopyWith<$Res> {
  factory $DeleteUserResultCopyWith(
          DeleteUserResult value, $Res Function(DeleteUserResult) then) =
      _$DeleteUserResultCopyWithImpl<$Res>;
  $Res call({bool wasDeleted});
}

class _$DeleteUserResultCopyWithImpl<$Res>
    implements $DeleteUserResultCopyWith<$Res> {
  _$DeleteUserResultCopyWithImpl(this._value, this._then);

  final DeleteUserResult _value;
  // ignore: unused_field
  final $Res Function(DeleteUserResult) _then;

  @override
  $Res call({
    Object wasDeleted = freezed,
  }) {
    return _then(_value.copyWith(
      wasDeleted:
          wasDeleted == freezed ? _value.wasDeleted : wasDeleted as bool,
    ));
  }
}

abstract class _$DeleteUserResultCopyWith<$Res>
    implements $DeleteUserResultCopyWith<$Res> {
  factory _$DeleteUserResultCopyWith(
          _DeleteUserResult value, $Res Function(_DeleteUserResult) then) =
      __$DeleteUserResultCopyWithImpl<$Res>;
  @override
  $Res call({bool wasDeleted});
}

class __$DeleteUserResultCopyWithImpl<$Res>
    extends _$DeleteUserResultCopyWithImpl<$Res>
    implements _$DeleteUserResultCopyWith<$Res> {
  __$DeleteUserResultCopyWithImpl(
      _DeleteUserResult _value, $Res Function(_DeleteUserResult) _then)
      : super(_value, (v) => _then(v as _DeleteUserResult));

  @override
  _DeleteUserResult get _value => super._value as _DeleteUserResult;

  @override
  $Res call({
    Object wasDeleted = freezed,
  }) {
    return _then(_DeleteUserResult(
      wasDeleted:
          wasDeleted == freezed ? _value.wasDeleted : wasDeleted as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_DeleteUserResult implements _DeleteUserResult {
  _$_DeleteUserResult({@required this.wasDeleted}) : assert(wasDeleted != null);

  factory _$_DeleteUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$_DeleteUserResultFromJson(json);

  @override
  final bool wasDeleted;

  @override
  String toString() {
    return 'DeleteUserResult(wasDeleted: $wasDeleted)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _DeleteUserResult &&
            (identical(other.wasDeleted, wasDeleted) ||
                const DeepCollectionEquality()
                    .equals(other.wasDeleted, wasDeleted)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasDeleted);

  @override
  _$DeleteUserResultCopyWith<_DeleteUserResult> get copyWith =>
      __$DeleteUserResultCopyWithImpl<_DeleteUserResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_DeleteUserResultToJson(this);
  }
}

abstract class _DeleteUserResult implements DeleteUserResult {
  factory _DeleteUserResult({@required bool wasDeleted}) = _$_DeleteUserResult;

  factory _DeleteUserResult.fromJson(Map<String, dynamic> json) =
      _$_DeleteUserResult.fromJson;

  @override
  bool get wasDeleted;
  @override
  _$DeleteUserResultCopyWith<_DeleteUserResult> get copyWith;
}
