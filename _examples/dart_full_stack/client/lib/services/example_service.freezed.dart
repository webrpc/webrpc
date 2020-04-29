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
class _$USER with DiagnosticableTreeMixin implements USER {
  const _$USER();

  factory _$USER.fromJson(Map<String, dynamic> json) => _$_$USERFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Kind.user()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'Kind.user'));
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
class _$ADMIN with DiagnosticableTreeMixin implements ADMIN {
  const _$ADMIN();

  factory _$ADMIN.fromJson(Map<String, dynamic> json) =>
      _$_$ADMINFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Kind.admin()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'Kind.admin'));
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
class _$_Empty with DiagnosticableTreeMixin implements _Empty {
  _$_Empty();

  factory _$_Empty.fromJson(Map<String, dynamic> json) =>
      _$_$_EmptyFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Empty()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'Empty'));
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
class _$_User with DiagnosticableTreeMixin implements _User {
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'User(id: $id, username: $username, role: $role)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'User'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('username', username))
      ..add(DiagnosticsProperty('role', role));
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
class _$_SearchFilter with DiagnosticableTreeMixin implements _SearchFilter {
  _$_SearchFilter({@required this.q}) : assert(q != null);

  factory _$_SearchFilter.fromJson(Map<String, dynamic> json) =>
      _$_$_SearchFilterFromJson(json);

  @override
  final String q;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'SearchFilter(q: $q)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'SearchFilter'))
      ..add(DiagnosticsProperty('q', q));
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
class _$_Version with DiagnosticableTreeMixin implements _Version {
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'Version(webrpcVersion: $webrpcVersion, schemaVersion: $schemaVersion, schemaHash: $schemaHash)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'Version'))
      ..add(DiagnosticsProperty('webrpcVersion', webrpcVersion))
      ..add(DiagnosticsProperty('schemaVersion', schemaVersion))
      ..add(DiagnosticsProperty('schemaHash', schemaHash));
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
class _$_ComplexType with DiagnosticableTreeMixin implements _ComplexType {
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ComplexType(meta: $meta, metaNestedExample: $metaNestedExample, namesList: $namesList, numsList: $numsList, doubleArray: $doubleArray, listOfUsers: $listOfUsers, mapOfUsers: $mapOfUsers, user: $user, listOfMaps: $listOfMaps)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ComplexType'))
      ..add(DiagnosticsProperty('meta', meta))
      ..add(DiagnosticsProperty('metaNestedExample', metaNestedExample))
      ..add(DiagnosticsProperty('namesList', namesList))
      ..add(DiagnosticsProperty('numsList', numsList))
      ..add(DiagnosticsProperty('doubleArray', doubleArray))
      ..add(DiagnosticsProperty('listOfUsers', listOfUsers))
      ..add(DiagnosticsProperty('mapOfUsers', mapOfUsers))
      ..add(DiagnosticsProperty('user', user))
      ..add(DiagnosticsProperty('listOfMaps', listOfMaps));
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
    with DiagnosticableTreeMixin
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_ExampleServiceGetUserArgs(header: $header, userID: $userID)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_ExampleServiceGetUserArgs'))
      ..add(DiagnosticsProperty('header', header))
      ..add(DiagnosticsProperty('userID', userID));
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
    with DiagnosticableTreeMixin
    implements _ExampleServiceFindUserArgs_Freezed {
  _$_ExampleServiceFindUserArgs_Freezed({@required this.s}) : assert(s != null);

  factory _$_ExampleServiceFindUserArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_ExampleServiceFindUserArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_ExampleServiceFindUserArgs(s: $s)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_ExampleServiceFindUserArgs'))
      ..add(DiagnosticsProperty('s', s));
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

StatusResponse _$StatusResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _StatusResponse.fromJson(json);
    case 'rpcErr':
      return _StatusResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$StatusResponseTearOff {
  const _$StatusResponseTearOff();

  _StatusResponse call({@required bool status}) {
    return _StatusResponse(
      status: status,
    );
  }

  _StatusResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _StatusResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $StatusResponse = _$StatusResponseTearOff();

mixin _$StatusResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_StatusResponse value), {
    @required Result rpcErr(_StatusResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_StatusResponse value), {
    Result rpcErr(_StatusResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $StatusResponseCopyWith<$Res> {
  factory $StatusResponseCopyWith(
          StatusResponse value, $Res Function(StatusResponse) then) =
      _$StatusResponseCopyWithImpl<$Res>;
}

class _$StatusResponseCopyWithImpl<$Res>
    implements $StatusResponseCopyWith<$Res> {
  _$StatusResponseCopyWithImpl(this._value, this._then);

  final StatusResponse _value;
  // ignore: unused_field
  final $Res Function(StatusResponse) _then;
}

abstract class _$StatusResponseCopyWith<$Res> {
  factory _$StatusResponseCopyWith(
          _StatusResponse value, $Res Function(_StatusResponse) then) =
      __$StatusResponseCopyWithImpl<$Res>;
  $Res call({bool status});
}

class __$StatusResponseCopyWithImpl<$Res>
    extends _$StatusResponseCopyWithImpl<$Res>
    implements _$StatusResponseCopyWith<$Res> {
  __$StatusResponseCopyWithImpl(
      _StatusResponse _value, $Res Function(_StatusResponse) _then)
      : super(_value, (v) => _then(v as _StatusResponse));

  @override
  _StatusResponse get _value => super._value as _StatusResponse;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_StatusResponse(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable()
class _$_StatusResponse
    with DiagnosticableTreeMixin
    implements _StatusResponse {
  _$_StatusResponse({@required this.status}) : assert(status != null);

  factory _$_StatusResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResponseFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResponse(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'StatusResponse'))
      ..add(DiagnosticsProperty('status', status));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _StatusResponse &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$StatusResponseCopyWith<_StatusResponse> get copyWith =>
      __$StatusResponseCopyWithImpl<_StatusResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(status);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(status);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_StatusResponse value), {
    @required Result rpcErr(_StatusResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_StatusResponse value), {
    Result rpcErr(_StatusResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _StatusResponse implements StatusResponse {
  factory _StatusResponse({@required bool status}) = _$_StatusResponse;

  factory _StatusResponse.fromJson(Map<String, dynamic> json) =
      _$_StatusResponse.fromJson;

  bool get status;
  _$StatusResponseCopyWith<_StatusResponse> get copyWith;
}

abstract class _$StatusResponseRpcErrCopyWith<$Res> {
  factory _$StatusResponseRpcErrCopyWith(_StatusResponseRpcErr value,
          $Res Function(_StatusResponseRpcErr) then) =
      __$StatusResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$StatusResponseRpcErrCopyWithImpl<$Res>
    extends _$StatusResponseCopyWithImpl<$Res>
    implements _$StatusResponseRpcErrCopyWith<$Res> {
  __$StatusResponseRpcErrCopyWithImpl(
      _StatusResponseRpcErr _value, $Res Function(_StatusResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _StatusResponseRpcErr));

  @override
  _StatusResponseRpcErr get _value => super._value as _StatusResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_StatusResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_StatusResponseRpcErr
    with DiagnosticableTreeMixin
    implements _StatusResponseRpcErr {
  _$_StatusResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_StatusResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'StatusResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _StatusResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$StatusResponseRpcErrCopyWith<_StatusResponseRpcErr> get copyWith =>
      __$StatusResponseRpcErrCopyWithImpl<_StatusResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_StatusResponse value), {
    @required Result rpcErr(_StatusResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_StatusResponse value), {
    Result rpcErr(_StatusResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResponseRpcErrToJson(this)..['runtimeType'] = 'rpcErr';
  }
}

abstract class _StatusResponseRpcErr implements StatusResponse {
  factory _StatusResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_StatusResponseRpcErr;

  factory _StatusResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_StatusResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$StatusResponseRpcErrCopyWith<_StatusResponseRpcErr> get copyWith;
}

VersionResponse _$VersionResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _VersionResponse.fromJson(json);
    case 'rpcErr':
      return _VersionResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$VersionResponseTearOff {
  const _$VersionResponseTearOff();

  _VersionResponse call({@required Version version}) {
    return _VersionResponse(
      version: version,
    );
  }

  _VersionResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _VersionResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $VersionResponse = _$VersionResponseTearOff();

mixin _$VersionResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_VersionResponse value), {
    @required Result rpcErr(_VersionResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_VersionResponse value), {
    Result rpcErr(_VersionResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $VersionResponseCopyWith<$Res> {
  factory $VersionResponseCopyWith(
          VersionResponse value, $Res Function(VersionResponse) then) =
      _$VersionResponseCopyWithImpl<$Res>;
}

class _$VersionResponseCopyWithImpl<$Res>
    implements $VersionResponseCopyWith<$Res> {
  _$VersionResponseCopyWithImpl(this._value, this._then);

  final VersionResponse _value;
  // ignore: unused_field
  final $Res Function(VersionResponse) _then;
}

abstract class _$VersionResponseCopyWith<$Res> {
  factory _$VersionResponseCopyWith(
          _VersionResponse value, $Res Function(_VersionResponse) then) =
      __$VersionResponseCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class __$VersionResponseCopyWithImpl<$Res>
    extends _$VersionResponseCopyWithImpl<$Res>
    implements _$VersionResponseCopyWith<$Res> {
  __$VersionResponseCopyWithImpl(
      _VersionResponse _value, $Res Function(_VersionResponse) _then)
      : super(_value, (v) => _then(v as _VersionResponse));

  @override
  _VersionResponse get _value => super._value as _VersionResponse;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_VersionResponse(
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

@JsonSerializable()
class _$_VersionResponse
    with DiagnosticableTreeMixin
    implements _VersionResponse {
  _$_VersionResponse({@required this.version}) : assert(version != null);

  factory _$_VersionResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResponseFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResponse(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'VersionResponse'))
      ..add(DiagnosticsProperty('version', version));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _VersionResponse &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$VersionResponseCopyWith<_VersionResponse> get copyWith =>
      __$VersionResponseCopyWithImpl<_VersionResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(version);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(version);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_VersionResponse value), {
    @required Result rpcErr(_VersionResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_VersionResponse value), {
    Result rpcErr(_VersionResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _VersionResponse implements VersionResponse {
  factory _VersionResponse({@required Version version}) = _$_VersionResponse;

  factory _VersionResponse.fromJson(Map<String, dynamic> json) =
      _$_VersionResponse.fromJson;

  Version get version;
  _$VersionResponseCopyWith<_VersionResponse> get copyWith;
}

abstract class _$VersionResponseRpcErrCopyWith<$Res> {
  factory _$VersionResponseRpcErrCopyWith(_VersionResponseRpcErr value,
          $Res Function(_VersionResponseRpcErr) then) =
      __$VersionResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$VersionResponseRpcErrCopyWithImpl<$Res>
    extends _$VersionResponseCopyWithImpl<$Res>
    implements _$VersionResponseRpcErrCopyWith<$Res> {
  __$VersionResponseRpcErrCopyWithImpl(_VersionResponseRpcErr _value,
      $Res Function(_VersionResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _VersionResponseRpcErr));

  @override
  _VersionResponseRpcErr get _value => super._value as _VersionResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_VersionResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_VersionResponseRpcErr
    with DiagnosticableTreeMixin
    implements _VersionResponseRpcErr {
  _$_VersionResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_VersionResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'VersionResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _VersionResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$VersionResponseRpcErrCopyWith<_VersionResponseRpcErr> get copyWith =>
      __$VersionResponseRpcErrCopyWithImpl<_VersionResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_VersionResponse value), {
    @required Result rpcErr(_VersionResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_VersionResponse value), {
    Result rpcErr(_VersionResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResponseRpcErrToJson(this)..['runtimeType'] = 'rpcErr';
  }
}

abstract class _VersionResponseRpcErr implements VersionResponse {
  factory _VersionResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_VersionResponseRpcErr;

  factory _VersionResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_VersionResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$VersionResponseRpcErrCopyWith<_VersionResponseRpcErr> get copyWith;
}

GetUserResponse _$GetUserResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _GetUserResponse.fromJson(json);
    case 'rpcErr':
      return _GetUserResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetUserResponseTearOff {
  const _$GetUserResponseTearOff();

  _GetUserResponse call({@required int code, @required User user}) {
    return _GetUserResponse(
      code: code,
      user: user,
    );
  }

  _GetUserResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _GetUserResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $GetUserResponse = _$GetUserResponseTearOff();

mixin _$GetUserResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUserResponse value), {
    @required Result rpcErr(_GetUserResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUserResponse value), {
    Result rpcErr(_GetUserResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetUserResponseCopyWith<$Res> {
  factory $GetUserResponseCopyWith(
          GetUserResponse value, $Res Function(GetUserResponse) then) =
      _$GetUserResponseCopyWithImpl<$Res>;
}

class _$GetUserResponseCopyWithImpl<$Res>
    implements $GetUserResponseCopyWith<$Res> {
  _$GetUserResponseCopyWithImpl(this._value, this._then);

  final GetUserResponse _value;
  // ignore: unused_field
  final $Res Function(GetUserResponse) _then;
}

abstract class _$GetUserResponseCopyWith<$Res> {
  factory _$GetUserResponseCopyWith(
          _GetUserResponse value, $Res Function(_GetUserResponse) then) =
      __$GetUserResponseCopyWithImpl<$Res>;
  $Res call({int code, User user});

  $UserCopyWith<$Res> get user;
}

class __$GetUserResponseCopyWithImpl<$Res>
    extends _$GetUserResponseCopyWithImpl<$Res>
    implements _$GetUserResponseCopyWith<$Res> {
  __$GetUserResponseCopyWithImpl(
      _GetUserResponse _value, $Res Function(_GetUserResponse) _then)
      : super(_value, (v) => _then(v as _GetUserResponse));

  @override
  _GetUserResponse get _value => super._value as _GetUserResponse;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_GetUserResponse(
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

@JsonSerializable()
class _$_GetUserResponse
    with DiagnosticableTreeMixin
    implements _GetUserResponse {
  _$_GetUserResponse({@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_GetUserResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResponseFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUserResponse(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUserResponse'))
      ..add(DiagnosticsProperty('code', code))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserResponse &&
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
  _$GetUserResponseCopyWith<_GetUserResponse> get copyWith =>
      __$GetUserResponseCopyWithImpl<_GetUserResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(code, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(code, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUserResponse value), {
    @required Result rpcErr(_GetUserResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUserResponse value), {
    Result rpcErr(_GetUserResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _GetUserResponse implements GetUserResponse {
  factory _GetUserResponse({@required int code, @required User user}) =
      _$_GetUserResponse;

  factory _GetUserResponse.fromJson(Map<String, dynamic> json) =
      _$_GetUserResponse.fromJson;

  int get code;
  User get user;
  _$GetUserResponseCopyWith<_GetUserResponse> get copyWith;
}

abstract class _$GetUserResponseRpcErrCopyWith<$Res> {
  factory _$GetUserResponseRpcErrCopyWith(_GetUserResponseRpcErr value,
          $Res Function(_GetUserResponseRpcErr) then) =
      __$GetUserResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$GetUserResponseRpcErrCopyWithImpl<$Res>
    extends _$GetUserResponseCopyWithImpl<$Res>
    implements _$GetUserResponseRpcErrCopyWith<$Res> {
  __$GetUserResponseRpcErrCopyWithImpl(_GetUserResponseRpcErr _value,
      $Res Function(_GetUserResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _GetUserResponseRpcErr));

  @override
  _GetUserResponseRpcErr get _value => super._value as _GetUserResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_GetUserResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_GetUserResponseRpcErr
    with DiagnosticableTreeMixin
    implements _GetUserResponseRpcErr {
  _$_GetUserResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_GetUserResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUserResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUserResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$GetUserResponseRpcErrCopyWith<_GetUserResponseRpcErr> get copyWith =>
      __$GetUserResponseRpcErrCopyWithImpl<_GetUserResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUserResponse value), {
    @required Result rpcErr(_GetUserResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUserResponse value), {
    Result rpcErr(_GetUserResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResponseRpcErrToJson(this)..['runtimeType'] = 'rpcErr';
  }
}

abstract class _GetUserResponseRpcErr implements GetUserResponse {
  factory _GetUserResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_GetUserResponseRpcErr;

  factory _GetUserResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_GetUserResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$GetUserResponseRpcErrCopyWith<_GetUserResponseRpcErr> get copyWith;
}

FindUserResponse _$FindUserResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _FindUserResponse.fromJson(json);
    case 'rpcErr':
      return _FindUserResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$FindUserResponseTearOff {
  const _$FindUserResponseTearOff();

  _FindUserResponse call({@required String name, @required User user}) {
    return _FindUserResponse(
      name: name,
      user: user,
    );
  }

  _FindUserResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _FindUserResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $FindUserResponse = _$FindUserResponseTearOff();

mixin _$FindUserResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_FindUserResponse value), {
    @required Result rpcErr(_FindUserResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_FindUserResponse value), {
    Result rpcErr(_FindUserResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $FindUserResponseCopyWith<$Res> {
  factory $FindUserResponseCopyWith(
          FindUserResponse value, $Res Function(FindUserResponse) then) =
      _$FindUserResponseCopyWithImpl<$Res>;
}

class _$FindUserResponseCopyWithImpl<$Res>
    implements $FindUserResponseCopyWith<$Res> {
  _$FindUserResponseCopyWithImpl(this._value, this._then);

  final FindUserResponse _value;
  // ignore: unused_field
  final $Res Function(FindUserResponse) _then;
}

abstract class _$FindUserResponseCopyWith<$Res> {
  factory _$FindUserResponseCopyWith(
          _FindUserResponse value, $Res Function(_FindUserResponse) then) =
      __$FindUserResponseCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class __$FindUserResponseCopyWithImpl<$Res>
    extends _$FindUserResponseCopyWithImpl<$Res>
    implements _$FindUserResponseCopyWith<$Res> {
  __$FindUserResponseCopyWithImpl(
      _FindUserResponse _value, $Res Function(_FindUserResponse) _then)
      : super(_value, (v) => _then(v as _FindUserResponse));

  @override
  _FindUserResponse get _value => super._value as _FindUserResponse;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_FindUserResponse(
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

@JsonSerializable()
class _$_FindUserResponse
    with DiagnosticableTreeMixin
    implements _FindUserResponse {
  _$_FindUserResponse({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_FindUserResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserResponseFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'FindUserResponse(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'FindUserResponse'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserResponse &&
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
  _$FindUserResponseCopyWith<_FindUserResponse> get copyWith =>
      __$FindUserResponseCopyWithImpl<_FindUserResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(name, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(name, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_FindUserResponse value), {
    @required Result rpcErr(_FindUserResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_FindUserResponse value), {
    Result rpcErr(_FindUserResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _FindUserResponse implements FindUserResponse {
  factory _FindUserResponse({@required String name, @required User user}) =
      _$_FindUserResponse;

  factory _FindUserResponse.fromJson(Map<String, dynamic> json) =
      _$_FindUserResponse.fromJson;

  String get name;
  User get user;
  _$FindUserResponseCopyWith<_FindUserResponse> get copyWith;
}

abstract class _$FindUserResponseRpcErrCopyWith<$Res> {
  factory _$FindUserResponseRpcErrCopyWith(_FindUserResponseRpcErr value,
          $Res Function(_FindUserResponseRpcErr) then) =
      __$FindUserResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$FindUserResponseRpcErrCopyWithImpl<$Res>
    extends _$FindUserResponseCopyWithImpl<$Res>
    implements _$FindUserResponseRpcErrCopyWith<$Res> {
  __$FindUserResponseRpcErrCopyWithImpl(_FindUserResponseRpcErr _value,
      $Res Function(_FindUserResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _FindUserResponseRpcErr));

  @override
  _FindUserResponseRpcErr get _value => super._value as _FindUserResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_FindUserResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_FindUserResponseRpcErr
    with DiagnosticableTreeMixin
    implements _FindUserResponseRpcErr {
  _$_FindUserResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_FindUserResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'FindUserResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'FindUserResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$FindUserResponseRpcErrCopyWith<_FindUserResponseRpcErr> get copyWith =>
      __$FindUserResponseRpcErrCopyWithImpl<_FindUserResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_FindUserResponse value), {
    @required Result rpcErr(_FindUserResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_FindUserResponse value), {
    Result rpcErr(_FindUserResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResponseRpcErrToJson(this)..['runtimeType'] = 'rpcErr';
  }
}

abstract class _FindUserResponseRpcErr implements FindUserResponse {
  factory _FindUserResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_FindUserResponseRpcErr;

  factory _FindUserResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_FindUserResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$FindUserResponseRpcErrCopyWith<_FindUserResponseRpcErr> get copyWith;
}

_AnotherExampleServiceGetAccountArgs
    _$_AnotherExampleServiceGetAccountArgsFromJson(Map<String, dynamic> json) {
  return _AnotherExampleServiceGetAccountArgs_Freezed.fromJson(json);
}

class _$_AnotherExampleServiceGetAccountArgsTearOff {
  const _$_AnotherExampleServiceGetAccountArgsTearOff();

  _AnotherExampleServiceGetAccountArgs_Freezed call(
      {@required Map<String, String> header, @required int userID}) {
    return _AnotherExampleServiceGetAccountArgs_Freezed(
      header: header,
      userID: userID,
    );
  }
}

// ignore: unused_element
const _$AnotherExampleServiceGetAccountArgs =
    _$_AnotherExampleServiceGetAccountArgsTearOff();

mixin _$_AnotherExampleServiceGetAccountArgs {
  Map<String, String> get header;
  int get userID;

  Map<String, dynamic> toJson();
  _$AnotherExampleServiceGetAccountArgsCopyWith<
      _AnotherExampleServiceGetAccountArgs> get copyWith;
}

abstract class _$AnotherExampleServiceGetAccountArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetAccountArgsCopyWith(
          _AnotherExampleServiceGetAccountArgs value,
          $Res Function(_AnotherExampleServiceGetAccountArgs) then) =
      __$AnotherExampleServiceGetAccountArgsCopyWithImpl<$Res>;
  $Res call({Map<String, String> header, int userID});
}

class __$AnotherExampleServiceGetAccountArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetAccountArgsCopyWith<$Res> {
  __$AnotherExampleServiceGetAccountArgsCopyWithImpl(this._value, this._then);

  final _AnotherExampleServiceGetAccountArgs _value;
  // ignore: unused_field
  final $Res Function(_AnotherExampleServiceGetAccountArgs) _then;

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

abstract class _$AnotherExampleServiceGetAccountArgs_FreezedCopyWith<$Res>
    implements _$AnotherExampleServiceGetAccountArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetAccountArgs_FreezedCopyWith(
          _AnotherExampleServiceGetAccountArgs_Freezed value,
          $Res Function(_AnotherExampleServiceGetAccountArgs_Freezed) then) =
      __$AnotherExampleServiceGetAccountArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> header, int userID});
}

class __$AnotherExampleServiceGetAccountArgs_FreezedCopyWithImpl<$Res>
    extends __$AnotherExampleServiceGetAccountArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetAccountArgs_FreezedCopyWith<$Res> {
  __$AnotherExampleServiceGetAccountArgs_FreezedCopyWithImpl(
      _AnotherExampleServiceGetAccountArgs_Freezed _value,
      $Res Function(_AnotherExampleServiceGetAccountArgs_Freezed) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceGetAccountArgs_Freezed));

  @override
  _AnotherExampleServiceGetAccountArgs_Freezed get _value =>
      super._value as _AnotherExampleServiceGetAccountArgs_Freezed;

  @override
  $Res call({
    Object header = freezed,
    Object userID = freezed,
  }) {
    return _then(_AnotherExampleServiceGetAccountArgs_Freezed(
      header: header == freezed ? _value.header : header as Map<String, String>,
      userID: userID == freezed ? _value.userID : userID as int,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceGetAccountArgs_Freezed
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceGetAccountArgs_Freezed {
  _$_AnotherExampleServiceGetAccountArgs_Freezed(
      {@required this.header, @required this.userID})
      : assert(header != null),
        assert(userID != null);

  factory _$_AnotherExampleServiceGetAccountArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceGetAccountArgs_FreezedFromJson(json);

  @override
  final Map<String, String> header;
  @override
  final int userID;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_AnotherExampleServiceGetAccountArgs(header: $header, userID: $userID)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_AnotherExampleServiceGetAccountArgs'))
      ..add(DiagnosticsProperty('header', header))
      ..add(DiagnosticsProperty('userID', userID));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceGetAccountArgs_Freezed &&
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
  _$AnotherExampleServiceGetAccountArgs_FreezedCopyWith<
          _AnotherExampleServiceGetAccountArgs_Freezed>
      get copyWith =>
          __$AnotherExampleServiceGetAccountArgs_FreezedCopyWithImpl<
              _AnotherExampleServiceGetAccountArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceGetAccountArgs_FreezedToJson(this);
  }
}

abstract class _AnotherExampleServiceGetAccountArgs_Freezed
    implements _AnotherExampleServiceGetAccountArgs {
  factory _AnotherExampleServiceGetAccountArgs_Freezed(
      {@required Map<String, String> header,
      @required int userID}) = _$_AnotherExampleServiceGetAccountArgs_Freezed;

  factory _AnotherExampleServiceGetAccountArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceGetAccountArgs_Freezed.fromJson;

  @override
  Map<String, String> get header;
  @override
  int get userID;
  @override
  _$AnotherExampleServiceGetAccountArgs_FreezedCopyWith<
      _AnotherExampleServiceGetAccountArgs_Freezed> get copyWith;
}

_AnotherExampleServiceGetUsersArgs _$_AnotherExampleServiceGetUsersArgsFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceGetUsersArgs_Freezed.fromJson(json);
}

class _$_AnotherExampleServiceGetUsersArgsTearOff {
  const _$_AnotherExampleServiceGetUsersArgsTearOff();

  _AnotherExampleServiceGetUsersArgs_Freezed call({@required SearchFilter s}) {
    return _AnotherExampleServiceGetUsersArgs_Freezed(
      s: s,
    );
  }
}

// ignore: unused_element
const _$AnotherExampleServiceGetUsersArgs =
    _$_AnotherExampleServiceGetUsersArgsTearOff();

mixin _$_AnotherExampleServiceGetUsersArgs {
  SearchFilter get s;

  Map<String, dynamic> toJson();
  _$AnotherExampleServiceGetUsersArgsCopyWith<
      _AnotherExampleServiceGetUsersArgs> get copyWith;
}

abstract class _$AnotherExampleServiceGetUsersArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetUsersArgsCopyWith(
          _AnotherExampleServiceGetUsersArgs value,
          $Res Function(_AnotherExampleServiceGetUsersArgs) then) =
      __$AnotherExampleServiceGetUsersArgsCopyWithImpl<$Res>;
  $Res call({SearchFilter s});

  $SearchFilterCopyWith<$Res> get s;
}

class __$AnotherExampleServiceGetUsersArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetUsersArgsCopyWith<$Res> {
  __$AnotherExampleServiceGetUsersArgsCopyWithImpl(this._value, this._then);

  final _AnotherExampleServiceGetUsersArgs _value;
  // ignore: unused_field
  final $Res Function(_AnotherExampleServiceGetUsersArgs) _then;

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

abstract class _$AnotherExampleServiceGetUsersArgs_FreezedCopyWith<$Res>
    implements _$AnotherExampleServiceGetUsersArgsCopyWith<$Res> {
  factory _$AnotherExampleServiceGetUsersArgs_FreezedCopyWith(
          _AnotherExampleServiceGetUsersArgs_Freezed value,
          $Res Function(_AnotherExampleServiceGetUsersArgs_Freezed) then) =
      __$AnotherExampleServiceGetUsersArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({SearchFilter s});

  @override
  $SearchFilterCopyWith<$Res> get s;
}

class __$AnotherExampleServiceGetUsersArgs_FreezedCopyWithImpl<$Res>
    extends __$AnotherExampleServiceGetUsersArgsCopyWithImpl<$Res>
    implements _$AnotherExampleServiceGetUsersArgs_FreezedCopyWith<$Res> {
  __$AnotherExampleServiceGetUsersArgs_FreezedCopyWithImpl(
      _AnotherExampleServiceGetUsersArgs_Freezed _value,
      $Res Function(_AnotherExampleServiceGetUsersArgs_Freezed) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceGetUsersArgs_Freezed));

  @override
  _AnotherExampleServiceGetUsersArgs_Freezed get _value =>
      super._value as _AnotherExampleServiceGetUsersArgs_Freezed;

  @override
  $Res call({
    Object s = freezed,
  }) {
    return _then(_AnotherExampleServiceGetUsersArgs_Freezed(
      s: s == freezed ? _value.s : s as SearchFilter,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceGetUsersArgs_Freezed
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceGetUsersArgs_Freezed {
  _$_AnotherExampleServiceGetUsersArgs_Freezed({@required this.s})
      : assert(s != null);

  factory _$_AnotherExampleServiceGetUsersArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceGetUsersArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_AnotherExampleServiceGetUsersArgs(s: $s)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_AnotherExampleServiceGetUsersArgs'))
      ..add(DiagnosticsProperty('s', s));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceGetUsersArgs_Freezed &&
            (identical(other.s, s) ||
                const DeepCollectionEquality().equals(other.s, s)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(s);

  @override
  _$AnotherExampleServiceGetUsersArgs_FreezedCopyWith<
          _AnotherExampleServiceGetUsersArgs_Freezed>
      get copyWith => __$AnotherExampleServiceGetUsersArgs_FreezedCopyWithImpl<
          _AnotherExampleServiceGetUsersArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceGetUsersArgs_FreezedToJson(this);
  }
}

abstract class _AnotherExampleServiceGetUsersArgs_Freezed
    implements _AnotherExampleServiceGetUsersArgs {
  factory _AnotherExampleServiceGetUsersArgs_Freezed(
          {@required SearchFilter s}) =
      _$_AnotherExampleServiceGetUsersArgs_Freezed;

  factory _AnotherExampleServiceGetUsersArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceGetUsersArgs_Freezed.fromJson;

  @override
  SearchFilter get s;
  @override
  _$AnotherExampleServiceGetUsersArgs_FreezedCopyWith<
      _AnotherExampleServiceGetUsersArgs_Freezed> get copyWith;
}

AnotherExampleServiceStatusResponse
    _$AnotherExampleServiceStatusResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _AnotherExampleServiceStatusResponse.fromJson(json);
    case 'rpcErr':
      return _AnotherExampleServiceStatusResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$AnotherExampleServiceStatusResponseTearOff {
  const _$AnotherExampleServiceStatusResponseTearOff();

  _AnotherExampleServiceStatusResponse call({@required bool status}) {
    return _AnotherExampleServiceStatusResponse(
      status: status,
    );
  }

  _AnotherExampleServiceStatusResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _AnotherExampleServiceStatusResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceStatusResponse =
    _$AnotherExampleServiceStatusResponseTearOff();

mixin _$AnotherExampleServiceStatusResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    @required Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $AnotherExampleServiceStatusResponseCopyWith<$Res> {
  factory $AnotherExampleServiceStatusResponseCopyWith(
          AnotherExampleServiceStatusResponse value,
          $Res Function(AnotherExampleServiceStatusResponse) then) =
      _$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>;
}

class _$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>
    implements $AnotherExampleServiceStatusResponseCopyWith<$Res> {
  _$AnotherExampleServiceStatusResponseCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceStatusResponse _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceStatusResponse) _then;
}

abstract class _$AnotherExampleServiceStatusResponseCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResponseCopyWith(
          _AnotherExampleServiceStatusResponse value,
          $Res Function(_AnotherExampleServiceStatusResponse) then) =
      __$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>;
  $Res call({bool status});
}

class __$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusResponseCopyWith<$Res> {
  __$AnotherExampleServiceStatusResponseCopyWithImpl(
      _AnotherExampleServiceStatusResponse _value,
      $Res Function(_AnotherExampleServiceStatusResponse) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceStatusResponse));

  @override
  _AnotherExampleServiceStatusResponse get _value =>
      super._value as _AnotherExampleServiceStatusResponse;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_AnotherExampleServiceStatusResponse(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceStatusResponse
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceStatusResponse {
  _$_AnotherExampleServiceStatusResponse({@required this.status})
      : assert(status != null);

  factory _$_AnotherExampleServiceStatusResponse.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusResponseFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AnotherExampleServiceStatusResponse(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'AnotherExampleServiceStatusResponse'))
      ..add(DiagnosticsProperty('status', status));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusResponse &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$AnotherExampleServiceStatusResponseCopyWith<
          _AnotherExampleServiceStatusResponse>
      get copyWith => __$AnotherExampleServiceStatusResponseCopyWithImpl<
          _AnotherExampleServiceStatusResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(status);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(status);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    @required Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResponseToJson(this)
      ..['runtimeType'] = 'default';
  }
}

abstract class _AnotherExampleServiceStatusResponse
    implements AnotherExampleServiceStatusResponse {
  factory _AnotherExampleServiceStatusResponse({@required bool status}) =
      _$_AnotherExampleServiceStatusResponse;

  factory _AnotherExampleServiceStatusResponse.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResponse.fromJson;

  bool get status;
  _$AnotherExampleServiceStatusResponseCopyWith<
      _AnotherExampleServiceStatusResponse> get copyWith;
}

abstract class _$AnotherExampleServiceStatusResponseRpcErrCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResponseRpcErrCopyWith(
          _AnotherExampleServiceStatusResponseRpcErr value,
          $Res Function(_AnotherExampleServiceStatusResponseRpcErr) then) =
      __$AnotherExampleServiceStatusResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$AnotherExampleServiceStatusResponseRpcErrCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusResponseCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusResponseRpcErrCopyWith<$Res> {
  __$AnotherExampleServiceStatusResponseRpcErrCopyWithImpl(
      _AnotherExampleServiceStatusResponseRpcErr _value,
      $Res Function(_AnotherExampleServiceStatusResponseRpcErr) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceStatusResponseRpcErr));

  @override
  _AnotherExampleServiceStatusResponseRpcErr get _value =>
      super._value as _AnotherExampleServiceStatusResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_AnotherExampleServiceStatusResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceStatusResponseRpcErr
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceStatusResponseRpcErr {
  _$_AnotherExampleServiceStatusResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_AnotherExampleServiceStatusResponseRpcErr.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AnotherExampleServiceStatusResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty(
          'type', 'AnotherExampleServiceStatusResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$AnotherExampleServiceStatusResponseRpcErrCopyWith<
          _AnotherExampleServiceStatusResponseRpcErr>
      get copyWith => __$AnotherExampleServiceStatusResponseRpcErrCopyWithImpl<
          _AnotherExampleServiceStatusResponseRpcErr>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(bool status), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(bool status), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    @required Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_AnotherExampleServiceStatusResponse value), {
    Result rpcErr(_AnotherExampleServiceStatusResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResponseRpcErrToJson(this)
      ..['runtimeType'] = 'rpcErr';
  }
}

abstract class _AnotherExampleServiceStatusResponseRpcErr
    implements AnotherExampleServiceStatusResponse {
  factory _AnotherExampleServiceStatusResponseRpcErr(
          {@required String message,
          @required String route,
          @required int statusCode,
          @required DateTime timeStamp}) =
      _$_AnotherExampleServiceStatusResponseRpcErr;

  factory _AnotherExampleServiceStatusResponseRpcErr.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$AnotherExampleServiceStatusResponseRpcErrCopyWith<
      _AnotherExampleServiceStatusResponseRpcErr> get copyWith;
}

GetVersionResponse _$GetVersionResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _GetVersionResponse.fromJson(json);
    case 'rpcErr':
      return _GetVersionResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetVersionResponseTearOff {
  const _$GetVersionResponseTearOff();

  _GetVersionResponse call({@required Version version}) {
    return _GetVersionResponse(
      version: version,
    );
  }

  _GetVersionResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _GetVersionResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $GetVersionResponse = _$GetVersionResponseTearOff();

mixin _$GetVersionResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    @required Result rpcErr(_GetVersionResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    Result rpcErr(_GetVersionResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetVersionResponseCopyWith<$Res> {
  factory $GetVersionResponseCopyWith(
          GetVersionResponse value, $Res Function(GetVersionResponse) then) =
      _$GetVersionResponseCopyWithImpl<$Res>;
}

class _$GetVersionResponseCopyWithImpl<$Res>
    implements $GetVersionResponseCopyWith<$Res> {
  _$GetVersionResponseCopyWithImpl(this._value, this._then);

  final GetVersionResponse _value;
  // ignore: unused_field
  final $Res Function(GetVersionResponse) _then;
}

abstract class _$GetVersionResponseCopyWith<$Res> {
  factory _$GetVersionResponseCopyWith(
          _GetVersionResponse value, $Res Function(_GetVersionResponse) then) =
      __$GetVersionResponseCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class __$GetVersionResponseCopyWithImpl<$Res>
    extends _$GetVersionResponseCopyWithImpl<$Res>
    implements _$GetVersionResponseCopyWith<$Res> {
  __$GetVersionResponseCopyWithImpl(
      _GetVersionResponse _value, $Res Function(_GetVersionResponse) _then)
      : super(_value, (v) => _then(v as _GetVersionResponse));

  @override
  _GetVersionResponse get _value => super._value as _GetVersionResponse;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_GetVersionResponse(
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

@JsonSerializable()
class _$_GetVersionResponse
    with DiagnosticableTreeMixin
    implements _GetVersionResponse {
  _$_GetVersionResponse({@required this.version}) : assert(version != null);

  factory _$_GetVersionResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_GetVersionResponseFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetVersionResponse(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetVersionResponse'))
      ..add(DiagnosticsProperty('version', version));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetVersionResponse &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$GetVersionResponseCopyWith<_GetVersionResponse> get copyWith =>
      __$GetVersionResponseCopyWithImpl<_GetVersionResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(version);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(version);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    @required Result rpcErr(_GetVersionResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    Result rpcErr(_GetVersionResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _GetVersionResponse implements GetVersionResponse {
  factory _GetVersionResponse({@required Version version}) =
      _$_GetVersionResponse;

  factory _GetVersionResponse.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResponse.fromJson;

  Version get version;
  _$GetVersionResponseCopyWith<_GetVersionResponse> get copyWith;
}

abstract class _$GetVersionResponseRpcErrCopyWith<$Res> {
  factory _$GetVersionResponseRpcErrCopyWith(_GetVersionResponseRpcErr value,
          $Res Function(_GetVersionResponseRpcErr) then) =
      __$GetVersionResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$GetVersionResponseRpcErrCopyWithImpl<$Res>
    extends _$GetVersionResponseCopyWithImpl<$Res>
    implements _$GetVersionResponseRpcErrCopyWith<$Res> {
  __$GetVersionResponseRpcErrCopyWithImpl(_GetVersionResponseRpcErr _value,
      $Res Function(_GetVersionResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _GetVersionResponseRpcErr));

  @override
  _GetVersionResponseRpcErr get _value =>
      super._value as _GetVersionResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_GetVersionResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_GetVersionResponseRpcErr
    with DiagnosticableTreeMixin
    implements _GetVersionResponseRpcErr {
  _$_GetVersionResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_GetVersionResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_GetVersionResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetVersionResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetVersionResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetVersionResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$GetVersionResponseRpcErrCopyWith<_GetVersionResponseRpcErr> get copyWith =>
      __$GetVersionResponseRpcErrCopyWithImpl<_GetVersionResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(Version version), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(Version version), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    @required Result rpcErr(_GetVersionResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetVersionResponse value), {
    Result rpcErr(_GetVersionResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResponseRpcErrToJson(this)
      ..['runtimeType'] = 'rpcErr';
  }
}

abstract class _GetVersionResponseRpcErr implements GetVersionResponse {
  factory _GetVersionResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_GetVersionResponseRpcErr;

  factory _GetVersionResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$GetVersionResponseRpcErrCopyWith<_GetVersionResponseRpcErr> get copyWith;
}

GetAccountResponse _$GetAccountResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _GetAccountResponse.fromJson(json);
    case 'rpcErr':
      return _GetAccountResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetAccountResponseTearOff {
  const _$GetAccountResponseTearOff();

  _GetAccountResponse call({@required int code, @required User user}) {
    return _GetAccountResponse(
      code: code,
      user: user,
    );
  }

  _GetAccountResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _GetAccountResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $GetAccountResponse = _$GetAccountResponseTearOff();

mixin _$GetAccountResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    @required Result rpcErr(_GetAccountResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    Result rpcErr(_GetAccountResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetAccountResponseCopyWith<$Res> {
  factory $GetAccountResponseCopyWith(
          GetAccountResponse value, $Res Function(GetAccountResponse) then) =
      _$GetAccountResponseCopyWithImpl<$Res>;
}

class _$GetAccountResponseCopyWithImpl<$Res>
    implements $GetAccountResponseCopyWith<$Res> {
  _$GetAccountResponseCopyWithImpl(this._value, this._then);

  final GetAccountResponse _value;
  // ignore: unused_field
  final $Res Function(GetAccountResponse) _then;
}

abstract class _$GetAccountResponseCopyWith<$Res> {
  factory _$GetAccountResponseCopyWith(
          _GetAccountResponse value, $Res Function(_GetAccountResponse) then) =
      __$GetAccountResponseCopyWithImpl<$Res>;
  $Res call({int code, User user});

  $UserCopyWith<$Res> get user;
}

class __$GetAccountResponseCopyWithImpl<$Res>
    extends _$GetAccountResponseCopyWithImpl<$Res>
    implements _$GetAccountResponseCopyWith<$Res> {
  __$GetAccountResponseCopyWithImpl(
      _GetAccountResponse _value, $Res Function(_GetAccountResponse) _then)
      : super(_value, (v) => _then(v as _GetAccountResponse));

  @override
  _GetAccountResponse get _value => super._value as _GetAccountResponse;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_GetAccountResponse(
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

@JsonSerializable()
class _$_GetAccountResponse
    with DiagnosticableTreeMixin
    implements _GetAccountResponse {
  _$_GetAccountResponse({@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_GetAccountResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_GetAccountResponseFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetAccountResponse(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetAccountResponse'))
      ..add(DiagnosticsProperty('code', code))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetAccountResponse &&
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
  _$GetAccountResponseCopyWith<_GetAccountResponse> get copyWith =>
      __$GetAccountResponseCopyWithImpl<_GetAccountResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(code, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(code, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    @required Result rpcErr(_GetAccountResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    Result rpcErr(_GetAccountResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _GetAccountResponse implements GetAccountResponse {
  factory _GetAccountResponse({@required int code, @required User user}) =
      _$_GetAccountResponse;

  factory _GetAccountResponse.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResponse.fromJson;

  int get code;
  User get user;
  _$GetAccountResponseCopyWith<_GetAccountResponse> get copyWith;
}

abstract class _$GetAccountResponseRpcErrCopyWith<$Res> {
  factory _$GetAccountResponseRpcErrCopyWith(_GetAccountResponseRpcErr value,
          $Res Function(_GetAccountResponseRpcErr) then) =
      __$GetAccountResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$GetAccountResponseRpcErrCopyWithImpl<$Res>
    extends _$GetAccountResponseCopyWithImpl<$Res>
    implements _$GetAccountResponseRpcErrCopyWith<$Res> {
  __$GetAccountResponseRpcErrCopyWithImpl(_GetAccountResponseRpcErr _value,
      $Res Function(_GetAccountResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _GetAccountResponseRpcErr));

  @override
  _GetAccountResponseRpcErr get _value =>
      super._value as _GetAccountResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_GetAccountResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_GetAccountResponseRpcErr
    with DiagnosticableTreeMixin
    implements _GetAccountResponseRpcErr {
  _$_GetAccountResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_GetAccountResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_GetAccountResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetAccountResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetAccountResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetAccountResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$GetAccountResponseRpcErrCopyWith<_GetAccountResponseRpcErr> get copyWith =>
      __$GetAccountResponseRpcErrCopyWithImpl<_GetAccountResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(int code, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(int code, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    @required Result rpcErr(_GetAccountResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetAccountResponse value), {
    Result rpcErr(_GetAccountResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResponseRpcErrToJson(this)
      ..['runtimeType'] = 'rpcErr';
  }
}

abstract class _GetAccountResponseRpcErr implements GetAccountResponse {
  factory _GetAccountResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_GetAccountResponseRpcErr;

  factory _GetAccountResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$GetAccountResponseRpcErrCopyWith<_GetAccountResponseRpcErr> get copyWith;
}

GetUsersResponse _$GetUsersResponseFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'default':
      return _GetUsersResponse.fromJson(json);
    case 'rpcErr':
      return _GetUsersResponseRpcErr.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetUsersResponseTearOff {
  const _$GetUsersResponseTearOff();

  _GetUsersResponse call({@required String name, @required User user}) {
    return _GetUsersResponse(
      name: name,
      user: user,
    );
  }

  _GetUsersResponseRpcErr rpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) {
    return _GetUsersResponseRpcErr(
      message: message,
      route: route,
      statusCode: statusCode,
      timeStamp: timeStamp,
    );
  }
}

// ignore: unused_element
const $GetUsersResponse = _$GetUsersResponseTearOff();

mixin _$GetUsersResponse {
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    @required Result rpcErr(_GetUsersResponseRpcErr value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    Result rpcErr(_GetUsersResponseRpcErr value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetUsersResponseCopyWith<$Res> {
  factory $GetUsersResponseCopyWith(
          GetUsersResponse value, $Res Function(GetUsersResponse) then) =
      _$GetUsersResponseCopyWithImpl<$Res>;
}

class _$GetUsersResponseCopyWithImpl<$Res>
    implements $GetUsersResponseCopyWith<$Res> {
  _$GetUsersResponseCopyWithImpl(this._value, this._then);

  final GetUsersResponse _value;
  // ignore: unused_field
  final $Res Function(GetUsersResponse) _then;
}

abstract class _$GetUsersResponseCopyWith<$Res> {
  factory _$GetUsersResponseCopyWith(
          _GetUsersResponse value, $Res Function(_GetUsersResponse) then) =
      __$GetUsersResponseCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class __$GetUsersResponseCopyWithImpl<$Res>
    extends _$GetUsersResponseCopyWithImpl<$Res>
    implements _$GetUsersResponseCopyWith<$Res> {
  __$GetUsersResponseCopyWithImpl(
      _GetUsersResponse _value, $Res Function(_GetUsersResponse) _then)
      : super(_value, (v) => _then(v as _GetUsersResponse));

  @override
  _GetUsersResponse get _value => super._value as _GetUsersResponse;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_GetUsersResponse(
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

@JsonSerializable()
class _$_GetUsersResponse
    with DiagnosticableTreeMixin
    implements _GetUsersResponse {
  _$_GetUsersResponse({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_GetUsersResponse.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUsersResponseFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUsersResponse(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUsersResponse'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUsersResponse &&
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
  _$GetUsersResponseCopyWith<_GetUsersResponse> get copyWith =>
      __$GetUsersResponseCopyWithImpl<_GetUsersResponse>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(name, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(name, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    @required Result rpcErr(_GetUsersResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return $default(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    Result rpcErr(_GetUsersResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if ($default != null) {
      return $default(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResponseToJson(this)..['runtimeType'] = 'default';
  }
}

abstract class _GetUsersResponse implements GetUsersResponse {
  factory _GetUsersResponse({@required String name, @required User user}) =
      _$_GetUsersResponse;

  factory _GetUsersResponse.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResponse.fromJson;

  String get name;
  User get user;
  _$GetUsersResponseCopyWith<_GetUsersResponse> get copyWith;
}

abstract class _$GetUsersResponseRpcErrCopyWith<$Res> {
  factory _$GetUsersResponseRpcErrCopyWith(_GetUsersResponseRpcErr value,
          $Res Function(_GetUsersResponseRpcErr) then) =
      __$GetUsersResponseRpcErrCopyWithImpl<$Res>;
  $Res call({String message, String route, int statusCode, DateTime timeStamp});
}

class __$GetUsersResponseRpcErrCopyWithImpl<$Res>
    extends _$GetUsersResponseCopyWithImpl<$Res>
    implements _$GetUsersResponseRpcErrCopyWith<$Res> {
  __$GetUsersResponseRpcErrCopyWithImpl(_GetUsersResponseRpcErr _value,
      $Res Function(_GetUsersResponseRpcErr) _then)
      : super(_value, (v) => _then(v as _GetUsersResponseRpcErr));

  @override
  _GetUsersResponseRpcErr get _value => super._value as _GetUsersResponseRpcErr;

  @override
  $Res call({
    Object message = freezed,
    Object route = freezed,
    Object statusCode = freezed,
    Object timeStamp = freezed,
  }) {
    return _then(_GetUsersResponseRpcErr(
      message: message == freezed ? _value.message : message as String,
      route: route == freezed ? _value.route : route as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      timeStamp:
          timeStamp == freezed ? _value.timeStamp : timeStamp as DateTime,
    ));
  }
}

@JsonSerializable()
class _$_GetUsersResponseRpcErr
    with DiagnosticableTreeMixin
    implements _GetUsersResponseRpcErr {
  _$_GetUsersResponseRpcErr(
      {@required this.message,
      @required this.route,
      @required this.statusCode,
      @required this.timeStamp})
      : assert(message != null),
        assert(route != null),
        assert(statusCode != null),
        assert(timeStamp != null);

  factory _$_GetUsersResponseRpcErr.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUsersResponseRpcErrFromJson(json);

  @override
  final String message;
  @override
  final String route;
  @override
  final int statusCode;
  @override
  final DateTime timeStamp;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUsersResponse.rpcErr(message: $message, route: $route, statusCode: $statusCode, timeStamp: $timeStamp)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUsersResponse.rpcErr'))
      ..add(DiagnosticsProperty('message', message))
      ..add(DiagnosticsProperty('route', route))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('timeStamp', timeStamp));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUsersResponseRpcErr &&
            (identical(other.message, message) ||
                const DeepCollectionEquality()
                    .equals(other.message, message)) &&
            (identical(other.route, route) ||
                const DeepCollectionEquality().equals(other.route, route)) &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.timeStamp, timeStamp) ||
                const DeepCollectionEquality()
                    .equals(other.timeStamp, timeStamp)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(message) ^
      const DeepCollectionEquality().hash(route) ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(timeStamp);

  @override
  _$GetUsersResponseRpcErrCopyWith<_GetUsersResponseRpcErr> get copyWith =>
      __$GetUsersResponseRpcErrCopyWithImpl<_GetUsersResponseRpcErr>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>(
    Result $default(String name, User user), {
    @required
        Result rpcErr(
            String message, String route, int statusCode, DateTime timeStamp),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(message, route, statusCode, timeStamp);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>(
    Result $default(String name, User user), {
    Result rpcErr(
        String message, String route, int statusCode, DateTime timeStamp),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(message, route, statusCode, timeStamp);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    @required Result rpcErr(_GetUsersResponseRpcErr value),
  }) {
    assert($default != null);
    assert(rpcErr != null);
    return rpcErr(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>(
    Result $default(_GetUsersResponse value), {
    Result rpcErr(_GetUsersResponseRpcErr value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (rpcErr != null) {
      return rpcErr(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResponseRpcErrToJson(this)..['runtimeType'] = 'rpcErr';
  }
}

abstract class _GetUsersResponseRpcErr implements GetUsersResponse {
  factory _GetUsersResponseRpcErr(
      {@required String message,
      @required String route,
      @required int statusCode,
      @required DateTime timeStamp}) = _$_GetUsersResponseRpcErr;

  factory _GetUsersResponseRpcErr.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResponseRpcErr.fromJson;

  String get message;
  String get route;
  int get statusCode;
  DateTime get timeStamp;
  _$GetUsersResponseRpcErrCopyWith<_GetUsersResponseRpcErr> get copyWith;
}
