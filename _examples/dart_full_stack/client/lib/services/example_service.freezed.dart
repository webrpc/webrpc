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

_ExampleServiceUpdateNameArgs _$_ExampleServiceUpdateNameArgsFromJson(
    Map<String, dynamic> json) {
  return _ExampleServiceUpdateNameArgs_Freezed.fromJson(json);
}

class _$_ExampleServiceUpdateNameArgsTearOff {
  const _$_ExampleServiceUpdateNameArgsTearOff();

  _ExampleServiceUpdateNameArgs_Freezed call({@required String username}) {
    return _ExampleServiceUpdateNameArgs_Freezed(
      username: username,
    );
  }
}

// ignore: unused_element
const _$ExampleServiceUpdateNameArgs = _$_ExampleServiceUpdateNameArgsTearOff();

mixin _$_ExampleServiceUpdateNameArgs {
  String get username;

  Map<String, dynamic> toJson();
  _$ExampleServiceUpdateNameArgsCopyWith<_ExampleServiceUpdateNameArgs>
      get copyWith;
}

abstract class _$ExampleServiceUpdateNameArgsCopyWith<$Res> {
  factory _$ExampleServiceUpdateNameArgsCopyWith(
          _ExampleServiceUpdateNameArgs value,
          $Res Function(_ExampleServiceUpdateNameArgs) then) =
      __$ExampleServiceUpdateNameArgsCopyWithImpl<$Res>;
  $Res call({String username});
}

class __$ExampleServiceUpdateNameArgsCopyWithImpl<$Res>
    implements _$ExampleServiceUpdateNameArgsCopyWith<$Res> {
  __$ExampleServiceUpdateNameArgsCopyWithImpl(this._value, this._then);

  final _ExampleServiceUpdateNameArgs _value;
  // ignore: unused_field
  final $Res Function(_ExampleServiceUpdateNameArgs) _then;

  @override
  $Res call({
    Object username = freezed,
  }) {
    return _then(_value.copyWith(
      username: username == freezed ? _value.username : username as String,
    ));
  }
}

abstract class _$ExampleServiceUpdateNameArgs_FreezedCopyWith<$Res>
    implements _$ExampleServiceUpdateNameArgsCopyWith<$Res> {
  factory _$ExampleServiceUpdateNameArgs_FreezedCopyWith(
          _ExampleServiceUpdateNameArgs_Freezed value,
          $Res Function(_ExampleServiceUpdateNameArgs_Freezed) then) =
      __$ExampleServiceUpdateNameArgs_FreezedCopyWithImpl<$Res>;
  @override
  $Res call({String username});
}

class __$ExampleServiceUpdateNameArgs_FreezedCopyWithImpl<$Res>
    extends __$ExampleServiceUpdateNameArgsCopyWithImpl<$Res>
    implements _$ExampleServiceUpdateNameArgs_FreezedCopyWith<$Res> {
  __$ExampleServiceUpdateNameArgs_FreezedCopyWithImpl(
      _ExampleServiceUpdateNameArgs_Freezed _value,
      $Res Function(_ExampleServiceUpdateNameArgs_Freezed) _then)
      : super(_value, (v) => _then(v as _ExampleServiceUpdateNameArgs_Freezed));

  @override
  _ExampleServiceUpdateNameArgs_Freezed get _value =>
      super._value as _ExampleServiceUpdateNameArgs_Freezed;

  @override
  $Res call({
    Object username = freezed,
  }) {
    return _then(_ExampleServiceUpdateNameArgs_Freezed(
      username: username == freezed ? _value.username : username as String,
    ));
  }
}

@JsonSerializable()
class _$_ExampleServiceUpdateNameArgs_Freezed
    with DiagnosticableTreeMixin
    implements _ExampleServiceUpdateNameArgs_Freezed {
  _$_ExampleServiceUpdateNameArgs_Freezed({@required this.username})
      : assert(username != null);

  factory _$_ExampleServiceUpdateNameArgs_Freezed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_ExampleServiceUpdateNameArgs_FreezedFromJson(json);

  @override
  final String username;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_ExampleServiceUpdateNameArgs(username: $username)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_ExampleServiceUpdateNameArgs'))
      ..add(DiagnosticsProperty('username', username));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ExampleServiceUpdateNameArgs_Freezed &&
            (identical(other.username, username) ||
                const DeepCollectionEquality()
                    .equals(other.username, username)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(username);

  @override
  _$ExampleServiceUpdateNameArgs_FreezedCopyWith<
          _ExampleServiceUpdateNameArgs_Freezed>
      get copyWith => __$ExampleServiceUpdateNameArgs_FreezedCopyWithImpl<
          _ExampleServiceUpdateNameArgs_Freezed>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_ExampleServiceUpdateNameArgs_FreezedToJson(this);
  }
}

abstract class _ExampleServiceUpdateNameArgs_Freezed
    implements _ExampleServiceUpdateNameArgs {
  factory _ExampleServiceUpdateNameArgs_Freezed({@required String username}) =
      _$_ExampleServiceUpdateNameArgs_Freezed;

  factory _ExampleServiceUpdateNameArgs_Freezed.fromJson(
          Map<String, dynamic> json) =
      _$_ExampleServiceUpdateNameArgs_Freezed.fromJson;

  @override
  String get username;
  @override
  _$ExampleServiceUpdateNameArgs_FreezedCopyWith<
      _ExampleServiceUpdateNameArgs_Freezed> get copyWith;
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
class _$_StatusResult with DiagnosticableTreeMixin implements _StatusResult {
  _$_StatusResult({@required this.status}) : assert(status != null);

  factory _$_StatusResult.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResultFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResult(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'StatusResult'))
      ..add(DiagnosticsProperty('status', status));
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
class _$_VersionResult with DiagnosticableTreeMixin implements _VersionResult {
  _$_VersionResult({@required this.version}) : assert(version != null);

  factory _$_VersionResult.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResultFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResult(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'VersionResult'))
      ..add(DiagnosticsProperty('version', version));
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

  _GetUserResult call({@required int code, @required User user}) {
    return _GetUserResult(
      code: code,
      user: user,
    );
  }
}

// ignore: unused_element
const $GetUserResult = _$GetUserResultTearOff();

mixin _$GetUserResult {
  int get code;
  User get user;

  Map<String, dynamic> toJson();
  $GetUserResultCopyWith<GetUserResult> get copyWith;
}

abstract class $GetUserResultCopyWith<$Res> {
  factory $GetUserResultCopyWith(
          GetUserResult value, $Res Function(GetUserResult) then) =
      _$GetUserResultCopyWithImpl<$Res>;
  $Res call({int code, User user});

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

abstract class _$GetUserResultCopyWith<$Res>
    implements $GetUserResultCopyWith<$Res> {
  factory _$GetUserResultCopyWith(
          _GetUserResult value, $Res Function(_GetUserResult) then) =
      __$GetUserResultCopyWithImpl<$Res>;
  @override
  $Res call({int code, User user});

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
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_GetUserResult(
      code: code == freezed ? _value.code : code as int,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetUserResult with DiagnosticableTreeMixin implements _GetUserResult {
  _$_GetUserResult({@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_GetUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResultFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUserResult(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUserResult'))
      ..add(DiagnosticsProperty('code', code))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserResult &&
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
  _$GetUserResultCopyWith<_GetUserResult> get copyWith =>
      __$GetUserResultCopyWithImpl<_GetUserResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResultToJson(this);
  }
}

abstract class _GetUserResult implements GetUserResult {
  factory _GetUserResult({@required int code, @required User user}) =
      _$_GetUserResult;

  factory _GetUserResult.fromJson(Map<String, dynamic> json) =
      _$_GetUserResult.fromJson;

  @override
  int get code;
  @override
  User get user;
  @override
  _$GetUserResultCopyWith<_GetUserResult> get copyWith;
}

FindUserResult _$FindUserResultFromJson(Map<String, dynamic> json) {
  return _FindUserResult.fromJson(json);
}

class _$FindUserResultTearOff {
  const _$FindUserResultTearOff();

  _FindUserResult call({@required String name, @required User user}) {
    return _FindUserResult(
      name: name,
      user: user,
    );
  }
}

// ignore: unused_element
const $FindUserResult = _$FindUserResultTearOff();

mixin _$FindUserResult {
  String get name;
  User get user;

  Map<String, dynamic> toJson();
  $FindUserResultCopyWith<FindUserResult> get copyWith;
}

abstract class $FindUserResultCopyWith<$Res> {
  factory $FindUserResultCopyWith(
          FindUserResult value, $Res Function(FindUserResult) then) =
      _$FindUserResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$FindUserResultCopyWithImpl<$Res>
    implements $FindUserResultCopyWith<$Res> {
  _$FindUserResultCopyWithImpl(this._value, this._then);

  final FindUserResult _value;
  // ignore: unused_field
  final $Res Function(FindUserResult) _then;

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

abstract class _$FindUserResultCopyWith<$Res>
    implements $FindUserResultCopyWith<$Res> {
  factory _$FindUserResultCopyWith(
          _FindUserResult value, $Res Function(_FindUserResult) then) =
      __$FindUserResultCopyWithImpl<$Res>;
  @override
  $Res call({String name, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$FindUserResultCopyWithImpl<$Res>
    extends _$FindUserResultCopyWithImpl<$Res>
    implements _$FindUserResultCopyWith<$Res> {
  __$FindUserResultCopyWithImpl(
      _FindUserResult _value, $Res Function(_FindUserResult) _then)
      : super(_value, (v) => _then(v as _FindUserResult));

  @override
  _FindUserResult get _value => super._value as _FindUserResult;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_FindUserResult(
      name: name == freezed ? _value.name : name as String,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_FindUserResult
    with DiagnosticableTreeMixin
    implements _FindUserResult {
  _$_FindUserResult({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_FindUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserResultFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'FindUserResult(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'FindUserResult'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserResult &&
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
  _$FindUserResultCopyWith<_FindUserResult> get copyWith =>
      __$FindUserResultCopyWithImpl<_FindUserResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResultToJson(this);
  }
}

abstract class _FindUserResult implements FindUserResult {
  factory _FindUserResult({@required String name, @required User user}) =
      _$_FindUserResult;

  factory _FindUserResult.fromJson(Map<String, dynamic> json) =
      _$_FindUserResult.fromJson;

  @override
  String get name;
  @override
  User get user;
  @override
  _$FindUserResultCopyWith<_FindUserResult> get copyWith;
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

AnotherExampleServiceStatusResult _$AnotherExampleServiceStatusResultFromJson(
    Map<String, dynamic> json) {
  return _AnotherExampleServiceStatusResult.fromJson(json);
}

class _$AnotherExampleServiceStatusResultTearOff {
  const _$AnotherExampleServiceStatusResultTearOff();

  _AnotherExampleServiceStatusResult call({@required bool status}) {
    return _AnotherExampleServiceStatusResult(
      status: status,
    );
  }
}

// ignore: unused_element
const $AnotherExampleServiceStatusResult =
    _$AnotherExampleServiceStatusResultTearOff();

mixin _$AnotherExampleServiceStatusResult {
  bool get status;

  Map<String, dynamic> toJson();
  $AnotherExampleServiceStatusResultCopyWith<AnotherExampleServiceStatusResult>
      get copyWith;
}

abstract class $AnotherExampleServiceStatusResultCopyWith<$Res> {
  factory $AnotherExampleServiceStatusResultCopyWith(
          AnotherExampleServiceStatusResult value,
          $Res Function(AnotherExampleServiceStatusResult) then) =
      _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    implements $AnotherExampleServiceStatusResultCopyWith<$Res> {
  _$AnotherExampleServiceStatusResultCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceStatusResult _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceStatusResult) _then;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_value.copyWith(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

abstract class _$AnotherExampleServiceStatusResultCopyWith<$Res>
    implements $AnotherExampleServiceStatusResultCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResultCopyWith(
          _AnotherExampleServiceStatusResult value,
          $Res Function(_AnotherExampleServiceStatusResult) then) =
      __$AnotherExampleServiceStatusResultCopyWithImpl<$Res>;
  @override
  $Res call({bool status});
}

class __$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusResultCopyWith<$Res> {
  __$AnotherExampleServiceStatusResultCopyWithImpl(
      _AnotherExampleServiceStatusResult _value,
      $Res Function(_AnotherExampleServiceStatusResult) _then)
      : super(_value, (v) => _then(v as _AnotherExampleServiceStatusResult));

  @override
  _AnotherExampleServiceStatusResult get _value =>
      super._value as _AnotherExampleServiceStatusResult;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(_AnotherExampleServiceStatusResult(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_AnotherExampleServiceStatusResult
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceStatusResult {
  _$_AnotherExampleServiceStatusResult({@required this.status})
      : assert(status != null);

  factory _$_AnotherExampleServiceStatusResult.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusResultFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AnotherExampleServiceStatusResult(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'AnotherExampleServiceStatusResult'))
      ..add(DiagnosticsProperty('status', status));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusResult &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  _$AnotherExampleServiceStatusResultCopyWith<
          _AnotherExampleServiceStatusResult>
      get copyWith => __$AnotherExampleServiceStatusResultCopyWithImpl<
          _AnotherExampleServiceStatusResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResultToJson(this);
  }
}

abstract class _AnotherExampleServiceStatusResult
    implements AnotherExampleServiceStatusResult {
  factory _AnotherExampleServiceStatusResult({@required bool status}) =
      _$_AnotherExampleServiceStatusResult;

  factory _AnotherExampleServiceStatusResult.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResult.fromJson;

  @override
  bool get status;
  @override
  _$AnotherExampleServiceStatusResultCopyWith<
      _AnotherExampleServiceStatusResult> get copyWith;
}

GetVersionResult _$GetVersionResultFromJson(Map<String, dynamic> json) {
  return _GetVersionResult.fromJson(json);
}

class _$GetVersionResultTearOff {
  const _$GetVersionResultTearOff();

  _GetVersionResult call({@required Version version}) {
    return _GetVersionResult(
      version: version,
    );
  }
}

// ignore: unused_element
const $GetVersionResult = _$GetVersionResultTearOff();

mixin _$GetVersionResult {
  Version get version;

  Map<String, dynamic> toJson();
  $GetVersionResultCopyWith<GetVersionResult> get copyWith;
}

abstract class $GetVersionResultCopyWith<$Res> {
  factory $GetVersionResultCopyWith(
          GetVersionResult value, $Res Function(GetVersionResult) then) =
      _$GetVersionResultCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class _$GetVersionResultCopyWithImpl<$Res>
    implements $GetVersionResultCopyWith<$Res> {
  _$GetVersionResultCopyWithImpl(this._value, this._then);

  final GetVersionResult _value;
  // ignore: unused_field
  final $Res Function(GetVersionResult) _then;

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

abstract class _$GetVersionResultCopyWith<$Res>
    implements $GetVersionResultCopyWith<$Res> {
  factory _$GetVersionResultCopyWith(
          _GetVersionResult value, $Res Function(_GetVersionResult) then) =
      __$GetVersionResultCopyWithImpl<$Res>;
  @override
  $Res call({Version version});

  @override
  $VersionCopyWith<$Res> get version;
}

class __$GetVersionResultCopyWithImpl<$Res>
    extends _$GetVersionResultCopyWithImpl<$Res>
    implements _$GetVersionResultCopyWith<$Res> {
  __$GetVersionResultCopyWithImpl(
      _GetVersionResult _value, $Res Function(_GetVersionResult) _then)
      : super(_value, (v) => _then(v as _GetVersionResult));

  @override
  _GetVersionResult get _value => super._value as _GetVersionResult;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(_GetVersionResult(
      version: version == freezed ? _value.version : version as Version,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetVersionResult
    with DiagnosticableTreeMixin
    implements _GetVersionResult {
  _$_GetVersionResult({@required this.version}) : assert(version != null);

  factory _$_GetVersionResult.fromJson(Map<String, dynamic> json) =>
      _$_$_GetVersionResultFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetVersionResult(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetVersionResult'))
      ..add(DiagnosticsProperty('version', version));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetVersionResult &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  _$GetVersionResultCopyWith<_GetVersionResult> get copyWith =>
      __$GetVersionResultCopyWithImpl<_GetVersionResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResultToJson(this);
  }
}

abstract class _GetVersionResult implements GetVersionResult {
  factory _GetVersionResult({@required Version version}) = _$_GetVersionResult;

  factory _GetVersionResult.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResult.fromJson;

  @override
  Version get version;
  @override
  _$GetVersionResultCopyWith<_GetVersionResult> get copyWith;
}

GetAccountResult _$GetAccountResultFromJson(Map<String, dynamic> json) {
  return _GetAccountResult.fromJson(json);
}

class _$GetAccountResultTearOff {
  const _$GetAccountResultTearOff();

  _GetAccountResult call({@required int code, @required User user}) {
    return _GetAccountResult(
      code: code,
      user: user,
    );
  }
}

// ignore: unused_element
const $GetAccountResult = _$GetAccountResultTearOff();

mixin _$GetAccountResult {
  int get code;
  User get user;

  Map<String, dynamic> toJson();
  $GetAccountResultCopyWith<GetAccountResult> get copyWith;
}

abstract class $GetAccountResultCopyWith<$Res> {
  factory $GetAccountResultCopyWith(
          GetAccountResult value, $Res Function(GetAccountResult) then) =
      _$GetAccountResultCopyWithImpl<$Res>;
  $Res call({int code, User user});

  $UserCopyWith<$Res> get user;
}

class _$GetAccountResultCopyWithImpl<$Res>
    implements $GetAccountResultCopyWith<$Res> {
  _$GetAccountResultCopyWithImpl(this._value, this._then);

  final GetAccountResult _value;
  // ignore: unused_field
  final $Res Function(GetAccountResult) _then;

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

abstract class _$GetAccountResultCopyWith<$Res>
    implements $GetAccountResultCopyWith<$Res> {
  factory _$GetAccountResultCopyWith(
          _GetAccountResult value, $Res Function(_GetAccountResult) then) =
      __$GetAccountResultCopyWithImpl<$Res>;
  @override
  $Res call({int code, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$GetAccountResultCopyWithImpl<$Res>
    extends _$GetAccountResultCopyWithImpl<$Res>
    implements _$GetAccountResultCopyWith<$Res> {
  __$GetAccountResultCopyWithImpl(
      _GetAccountResult _value, $Res Function(_GetAccountResult) _then)
      : super(_value, (v) => _then(v as _GetAccountResult));

  @override
  _GetAccountResult get _value => super._value as _GetAccountResult;

  @override
  $Res call({
    Object code = freezed,
    Object user = freezed,
  }) {
    return _then(_GetAccountResult(
      code: code == freezed ? _value.code : code as int,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetAccountResult
    with DiagnosticableTreeMixin
    implements _GetAccountResult {
  _$_GetAccountResult({@required this.code, @required this.user})
      : assert(code != null),
        assert(user != null);

  factory _$_GetAccountResult.fromJson(Map<String, dynamic> json) =>
      _$_$_GetAccountResultFromJson(json);

  @override
  final int code;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetAccountResult(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetAccountResult'))
      ..add(DiagnosticsProperty('code', code))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetAccountResult &&
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
  _$GetAccountResultCopyWith<_GetAccountResult> get copyWith =>
      __$GetAccountResultCopyWithImpl<_GetAccountResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResultToJson(this);
  }
}

abstract class _GetAccountResult implements GetAccountResult {
  factory _GetAccountResult({@required int code, @required User user}) =
      _$_GetAccountResult;

  factory _GetAccountResult.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResult.fromJson;

  @override
  int get code;
  @override
  User get user;
  @override
  _$GetAccountResultCopyWith<_GetAccountResult> get copyWith;
}

GetUsersResult _$GetUsersResultFromJson(Map<String, dynamic> json) {
  return _GetUsersResult.fromJson(json);
}

class _$GetUsersResultTearOff {
  const _$GetUsersResultTearOff();

  _GetUsersResult call({@required String name, @required User user}) {
    return _GetUsersResult(
      name: name,
      user: user,
    );
  }
}

// ignore: unused_element
const $GetUsersResult = _$GetUsersResultTearOff();

mixin _$GetUsersResult {
  String get name;
  User get user;

  Map<String, dynamic> toJson();
  $GetUsersResultCopyWith<GetUsersResult> get copyWith;
}

abstract class $GetUsersResultCopyWith<$Res> {
  factory $GetUsersResultCopyWith(
          GetUsersResult value, $Res Function(GetUsersResult) then) =
      _$GetUsersResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$GetUsersResultCopyWithImpl<$Res>
    implements $GetUsersResultCopyWith<$Res> {
  _$GetUsersResultCopyWithImpl(this._value, this._then);

  final GetUsersResult _value;
  // ignore: unused_field
  final $Res Function(GetUsersResult) _then;

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

abstract class _$GetUsersResultCopyWith<$Res>
    implements $GetUsersResultCopyWith<$Res> {
  factory _$GetUsersResultCopyWith(
          _GetUsersResult value, $Res Function(_GetUsersResult) then) =
      __$GetUsersResultCopyWithImpl<$Res>;
  @override
  $Res call({String name, User user});

  @override
  $UserCopyWith<$Res> get user;
}

class __$GetUsersResultCopyWithImpl<$Res>
    extends _$GetUsersResultCopyWithImpl<$Res>
    implements _$GetUsersResultCopyWith<$Res> {
  __$GetUsersResultCopyWithImpl(
      _GetUsersResult _value, $Res Function(_GetUsersResult) _then)
      : super(_value, (v) => _then(v as _GetUsersResult));

  @override
  _GetUsersResult get _value => super._value as _GetUsersResult;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(_GetUsersResult(
      name: name == freezed ? _value.name : name as String,
      user: user == freezed ? _value.user : user as User,
    ));
  }
}

@JsonSerializable(explicitToJson: true)
class _$_GetUsersResult
    with DiagnosticableTreeMixin
    implements _GetUsersResult {
  _$_GetUsersResult({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$_GetUsersResult.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUsersResultFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUsersResult(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUsersResult'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUsersResult &&
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
  _$GetUsersResultCopyWith<_GetUsersResult> get copyWith =>
      __$GetUsersResultCopyWithImpl<_GetUsersResult>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResultToJson(this);
  }
}

abstract class _GetUsersResult implements GetUsersResult {
  factory _GetUsersResult({@required String name, @required User user}) =
      _$_GetUsersResult;

  factory _GetUsersResult.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResult.fromJson;

  @override
  String get name;
  @override
  User get user;
  @override
  _$GetUsersResultCopyWith<_GetUsersResult> get copyWith;
}

class _$RpcResponseTearOff {
  const _$RpcResponseTearOff();

  _RpcResponseOk<T> ok<T>({@required T data}) {
    return _RpcResponseOk<T>(
      data: data,
    );
  }

  _RpcResonseErr<T> err<T>(
      {@required String reason, @required int statusCode, String stackTrace}) {
    return _RpcResonseErr<T>(
      reason: reason,
      statusCode: statusCode,
      stackTrace: stackTrace,
    );
  }

  _RpcResponsePending<T> pending<T>() {
    return _RpcResponsePending<T>();
  }
}

// ignore: unused_element
const $RpcResponse = _$RpcResponseTearOff();

mixin _$RpcResponse<T> {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ok(T data),
    @required Result err(String reason, int statusCode, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ok(T data),
    Result err(String reason, int statusCode, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ok(_RpcResponseOk<T> value),
    @required Result err(_RpcResonseErr<T> value),
    @required Result pending(_RpcResponsePending<T> value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ok(_RpcResponseOk<T> value),
    Result err(_RpcResonseErr<T> value),
    Result pending(_RpcResponsePending<T> value),
    @required Result orElse(),
  });
}

abstract class $RpcResponseCopyWith<T, $Res> {
  factory $RpcResponseCopyWith(
          RpcResponse<T> value, $Res Function(RpcResponse<T>) then) =
      _$RpcResponseCopyWithImpl<T, $Res>;
}

class _$RpcResponseCopyWithImpl<T, $Res>
    implements $RpcResponseCopyWith<T, $Res> {
  _$RpcResponseCopyWithImpl(this._value, this._then);

  final RpcResponse<T> _value;
  // ignore: unused_field
  final $Res Function(RpcResponse<T>) _then;
}

abstract class _$RpcResponseOkCopyWith<T, $Res> {
  factory _$RpcResponseOkCopyWith(
          _RpcResponseOk<T> value, $Res Function(_RpcResponseOk<T>) then) =
      __$RpcResponseOkCopyWithImpl<T, $Res>;
  $Res call({T data});
}

class __$RpcResponseOkCopyWithImpl<T, $Res>
    extends _$RpcResponseCopyWithImpl<T, $Res>
    implements _$RpcResponseOkCopyWith<T, $Res> {
  __$RpcResponseOkCopyWithImpl(
      _RpcResponseOk<T> _value, $Res Function(_RpcResponseOk<T>) _then)
      : super(_value, (v) => _then(v as _RpcResponseOk<T>));

  @override
  _RpcResponseOk<T> get _value => super._value as _RpcResponseOk<T>;

  @override
  $Res call({
    Object data = freezed,
  }) {
    return _then(_RpcResponseOk<T>(
      data: data == freezed ? _value.data : data as T,
    ));
  }
}

class _$_RpcResponseOk<T>
    with DiagnosticableTreeMixin
    implements _RpcResponseOk<T> {
  _$_RpcResponseOk({@required this.data}) : assert(data != null);

  @override
  final T data;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcResponse<$T>.ok(data: $data)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'RpcResponse<$T>.ok'))
      ..add(DiagnosticsProperty('data', data));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _RpcResponseOk<T> &&
            (identical(other.data, data) ||
                const DeepCollectionEquality().equals(other.data, data)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(data);

  @override
  _$RpcResponseOkCopyWith<T, _RpcResponseOk<T>> get copyWith =>
      __$RpcResponseOkCopyWithImpl<T, _RpcResponseOk<T>>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ok(T data),
    @required Result err(String reason, int statusCode, String stackTrace),
    @required Result pending(),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return ok(data);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ok(T data),
    Result err(String reason, int statusCode, String stackTrace),
    Result pending(),
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
    @required Result ok(_RpcResponseOk<T> value),
    @required Result err(_RpcResonseErr<T> value),
    @required Result pending(_RpcResponsePending<T> value),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return ok(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ok(_RpcResponseOk<T> value),
    Result err(_RpcResonseErr<T> value),
    Result pending(_RpcResponsePending<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (ok != null) {
      return ok(this);
    }
    return orElse();
  }
}

abstract class _RpcResponseOk<T> implements RpcResponse<T> {
  factory _RpcResponseOk({@required T data}) = _$_RpcResponseOk<T>;

  T get data;
  _$RpcResponseOkCopyWith<T, _RpcResponseOk<T>> get copyWith;
}

abstract class _$RpcResonseErrCopyWith<T, $Res> {
  factory _$RpcResonseErrCopyWith(
          _RpcResonseErr<T> value, $Res Function(_RpcResonseErr<T>) then) =
      __$RpcResonseErrCopyWithImpl<T, $Res>;
  $Res call({String reason, int statusCode, String stackTrace});
}

class __$RpcResonseErrCopyWithImpl<T, $Res>
    extends _$RpcResponseCopyWithImpl<T, $Res>
    implements _$RpcResonseErrCopyWith<T, $Res> {
  __$RpcResonseErrCopyWithImpl(
      _RpcResonseErr<T> _value, $Res Function(_RpcResonseErr<T>) _then)
      : super(_value, (v) => _then(v as _RpcResonseErr<T>));

  @override
  _RpcResonseErr<T> get _value => super._value as _RpcResonseErr<T>;

  @override
  $Res call({
    Object reason = freezed,
    Object statusCode = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_RpcResonseErr<T>(
      reason: reason == freezed ? _value.reason : reason as String,
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

class _$_RpcResonseErr<T>
    with DiagnosticableTreeMixin
    implements _RpcResonseErr<T> {
  _$_RpcResonseErr(
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcResponse<$T>.err(reason: $reason, statusCode: $statusCode, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'RpcResponse<$T>.err'))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _RpcResonseErr<T> &&
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
  _$RpcResonseErrCopyWith<T, _RpcResonseErr<T>> get copyWith =>
      __$RpcResonseErrCopyWithImpl<T, _RpcResonseErr<T>>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ok(T data),
    @required Result err(String reason, int statusCode, String stackTrace),
    @required Result pending(),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return err(reason, statusCode, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ok(T data),
    Result err(String reason, int statusCode, String stackTrace),
    Result pending(),
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
    @required Result ok(_RpcResponseOk<T> value),
    @required Result err(_RpcResonseErr<T> value),
    @required Result pending(_RpcResponsePending<T> value),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return err(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ok(_RpcResponseOk<T> value),
    Result err(_RpcResonseErr<T> value),
    Result pending(_RpcResponsePending<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (err != null) {
      return err(this);
    }
    return orElse();
  }
}

abstract class _RpcResonseErr<T> implements RpcResponse<T> {
  factory _RpcResonseErr(
      {@required String reason,
      @required int statusCode,
      String stackTrace}) = _$_RpcResonseErr<T>;

  String get reason;
  int get statusCode;
  String get stackTrace;
  _$RpcResonseErrCopyWith<T, _RpcResonseErr<T>> get copyWith;
}

abstract class _$RpcResponsePendingCopyWith<T, $Res> {
  factory _$RpcResponsePendingCopyWith(_RpcResponsePending<T> value,
          $Res Function(_RpcResponsePending<T>) then) =
      __$RpcResponsePendingCopyWithImpl<T, $Res>;
}

class __$RpcResponsePendingCopyWithImpl<T, $Res>
    extends _$RpcResponseCopyWithImpl<T, $Res>
    implements _$RpcResponsePendingCopyWith<T, $Res> {
  __$RpcResponsePendingCopyWithImpl(_RpcResponsePending<T> _value,
      $Res Function(_RpcResponsePending<T>) _then)
      : super(_value, (v) => _then(v as _RpcResponsePending<T>));

  @override
  _RpcResponsePending<T> get _value => super._value as _RpcResponsePending<T>;
}

class _$_RpcResponsePending<T>
    with DiagnosticableTreeMixin
    implements _RpcResponsePending<T> {
  const _$_RpcResponsePending();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcResponse<$T>.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'RpcResponse<$T>.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _RpcResponsePending<T>);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ok(T data),
    @required Result err(String reason, int statusCode, String stackTrace),
    @required Result pending(),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ok(T data),
    Result err(String reason, int statusCode, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ok(_RpcResponseOk<T> value),
    @required Result err(_RpcResonseErr<T> value),
    @required Result pending(_RpcResponsePending<T> value),
  }) {
    assert(ok != null);
    assert(err != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ok(_RpcResponseOk<T> value),
    Result err(_RpcResonseErr<T> value),
    Result pending(_RpcResponsePending<T> value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }
}

abstract class _RpcResponsePending<T> implements RpcResponse<T> {
  const factory _RpcResponsePending() = _$_RpcResponsePending<T>;
}
