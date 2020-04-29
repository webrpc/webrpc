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

StatusResult _$StatusResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _StatusResult.fromJson(json);
    case 'failed':
      return _StatusResultFailed.fromJson(json);
    case 'pending':
      return _StatusResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$StatusResultTearOff {
  const _$StatusResultTearOff();

  _StatusResult data({@required bool status}) {
    return _StatusResult(
      status: status,
    );
  }

  _StatusResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _StatusResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _StatusResultPending pending() {
    return const _StatusResultPending();
  }
}

// ignore: unused_element
const $StatusResult = _$StatusResultTearOff();

mixin _$StatusResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_StatusResult value),
    @required Result failed(_StatusResultFailed value),
    @required Result pending(_StatusResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_StatusResult value),
    Result failed(_StatusResultFailed value),
    Result pending(_StatusResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $StatusResultCopyWith<$Res> {
  factory $StatusResultCopyWith(
          StatusResult value, $Res Function(StatusResult) then) =
      _$StatusResultCopyWithImpl<$Res>;
}

class _$StatusResultCopyWithImpl<$Res> implements $StatusResultCopyWith<$Res> {
  _$StatusResultCopyWithImpl(this._value, this._then);

  final StatusResult _value;
  // ignore: unused_field
  final $Res Function(StatusResult) _then;
}

abstract class _$StatusResultCopyWith<$Res> {
  factory _$StatusResultCopyWith(
          _StatusResult value, $Res Function(_StatusResult) then) =
      __$StatusResultCopyWithImpl<$Res>;
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

@JsonSerializable()
class _$_StatusResult with DiagnosticableTreeMixin implements _StatusResult {
  _$_StatusResult({@required this.status}) : assert(status != null);

  factory _$_StatusResult.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResultFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResult.data(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'StatusResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(status);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(status);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_StatusResult value),
    @required Result failed(_StatusResultFailed value),
    @required Result pending(_StatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_StatusResult value),
    Result failed(_StatusResultFailed value),
    Result pending(_StatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _StatusResult implements StatusResult {
  factory _StatusResult({@required bool status}) = _$_StatusResult;

  factory _StatusResult.fromJson(Map<String, dynamic> json) =
      _$_StatusResult.fromJson;

  bool get status;
  _$StatusResultCopyWith<_StatusResult> get copyWith;
}

abstract class _$StatusResultFailedCopyWith<$Res> {
  factory _$StatusResultFailedCopyWith(
          _StatusResultFailed value, $Res Function(_StatusResultFailed) then) =
      __$StatusResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$StatusResultFailedCopyWithImpl<$Res>
    extends _$StatusResultCopyWithImpl<$Res>
    implements _$StatusResultFailedCopyWith<$Res> {
  __$StatusResultFailedCopyWithImpl(
      _StatusResultFailed _value, $Res Function(_StatusResultFailed) _then)
      : super(_value, (v) => _then(v as _StatusResultFailed));

  @override
  _StatusResultFailed get _value => super._value as _StatusResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_StatusResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_StatusResultFailed
    with DiagnosticableTreeMixin
    implements _StatusResultFailed {
  _$_StatusResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_StatusResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'StatusResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _StatusResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$StatusResultFailedCopyWith<_StatusResultFailed> get copyWith =>
      __$StatusResultFailedCopyWithImpl<_StatusResultFailed>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_StatusResult value),
    @required Result failed(_StatusResultFailed value),
    @required Result pending(_StatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_StatusResult value),
    Result failed(_StatusResultFailed value),
    Result pending(_StatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _StatusResultFailed implements StatusResult {
  factory _StatusResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_StatusResultFailed;

  factory _StatusResultFailed.fromJson(Map<String, dynamic> json) =
      _$_StatusResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$StatusResultFailedCopyWith<_StatusResultFailed> get copyWith;
}

abstract class _$StatusResultPendingCopyWith<$Res> {
  factory _$StatusResultPendingCopyWith(_StatusResultPending value,
          $Res Function(_StatusResultPending) then) =
      __$StatusResultPendingCopyWithImpl<$Res>;
}

class __$StatusResultPendingCopyWithImpl<$Res>
    extends _$StatusResultCopyWithImpl<$Res>
    implements _$StatusResultPendingCopyWith<$Res> {
  __$StatusResultPendingCopyWithImpl(
      _StatusResultPending _value, $Res Function(_StatusResultPending) _then)
      : super(_value, (v) => _then(v as _StatusResultPending));

  @override
  _StatusResultPending get _value => super._value as _StatusResultPending;
}

@JsonSerializable()
class _$_StatusResultPending
    with DiagnosticableTreeMixin
    implements _StatusResultPending {
  const _$_StatusResultPending();

  factory _$_StatusResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_StatusResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'StatusResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'StatusResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _StatusResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_StatusResult value),
    @required Result failed(_StatusResultFailed value),
    @required Result pending(_StatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_StatusResult value),
    Result failed(_StatusResultFailed value),
    Result pending(_StatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_StatusResultPendingToJson(this)..['runtimeType'] = 'pending';
  }
}

abstract class _StatusResultPending implements StatusResult {
  const factory _StatusResultPending() = _$_StatusResultPending;

  factory _StatusResultPending.fromJson(Map<String, dynamic> json) =
      _$_StatusResultPending.fromJson;
}

VersionResult _$VersionResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _VersionResult.fromJson(json);
    case 'failed':
      return _VersionResultFailed.fromJson(json);
    case 'pending':
      return _VersionResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$VersionResultTearOff {
  const _$VersionResultTearOff();

  _VersionResult data({@required Version version}) {
    return _VersionResult(
      version: version,
    );
  }

  _VersionResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _VersionResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _VersionResultPending pending() {
    return const _VersionResultPending();
  }
}

// ignore: unused_element
const $VersionResult = _$VersionResultTearOff();

mixin _$VersionResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_VersionResult value),
    @required Result failed(_VersionResultFailed value),
    @required Result pending(_VersionResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_VersionResult value),
    Result failed(_VersionResultFailed value),
    Result pending(_VersionResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $VersionResultCopyWith<$Res> {
  factory $VersionResultCopyWith(
          VersionResult value, $Res Function(VersionResult) then) =
      _$VersionResultCopyWithImpl<$Res>;
}

class _$VersionResultCopyWithImpl<$Res>
    implements $VersionResultCopyWith<$Res> {
  _$VersionResultCopyWithImpl(this._value, this._then);

  final VersionResult _value;
  // ignore: unused_field
  final $Res Function(VersionResult) _then;
}

abstract class _$VersionResultCopyWith<$Res> {
  factory _$VersionResultCopyWith(
          _VersionResult value, $Res Function(_VersionResult) then) =
      __$VersionResultCopyWithImpl<$Res>;
  $Res call({Version version});

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
class _$_VersionResult with DiagnosticableTreeMixin implements _VersionResult {
  _$_VersionResult({@required this.version}) : assert(version != null);

  factory _$_VersionResult.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResultFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResult.data(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'VersionResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(version);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(version);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_VersionResult value),
    @required Result failed(_VersionResultFailed value),
    @required Result pending(_VersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_VersionResult value),
    Result failed(_VersionResultFailed value),
    Result pending(_VersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _VersionResult implements VersionResult {
  factory _VersionResult({@required Version version}) = _$_VersionResult;

  factory _VersionResult.fromJson(Map<String, dynamic> json) =
      _$_VersionResult.fromJson;

  Version get version;
  _$VersionResultCopyWith<_VersionResult> get copyWith;
}

abstract class _$VersionResultFailedCopyWith<$Res> {
  factory _$VersionResultFailedCopyWith(_VersionResultFailed value,
          $Res Function(_VersionResultFailed) then) =
      __$VersionResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$VersionResultFailedCopyWithImpl<$Res>
    extends _$VersionResultCopyWithImpl<$Res>
    implements _$VersionResultFailedCopyWith<$Res> {
  __$VersionResultFailedCopyWithImpl(
      _VersionResultFailed _value, $Res Function(_VersionResultFailed) _then)
      : super(_value, (v) => _then(v as _VersionResultFailed));

  @override
  _VersionResultFailed get _value => super._value as _VersionResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_VersionResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_VersionResultFailed
    with DiagnosticableTreeMixin
    implements _VersionResultFailed {
  _$_VersionResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_VersionResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'VersionResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _VersionResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$VersionResultFailedCopyWith<_VersionResultFailed> get copyWith =>
      __$VersionResultFailedCopyWithImpl<_VersionResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_VersionResult value),
    @required Result failed(_VersionResultFailed value),
    @required Result pending(_VersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_VersionResult value),
    Result failed(_VersionResultFailed value),
    Result pending(_VersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _VersionResultFailed implements VersionResult {
  factory _VersionResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_VersionResultFailed;

  factory _VersionResultFailed.fromJson(Map<String, dynamic> json) =
      _$_VersionResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$VersionResultFailedCopyWith<_VersionResultFailed> get copyWith;
}

abstract class _$VersionResultPendingCopyWith<$Res> {
  factory _$VersionResultPendingCopyWith(_VersionResultPending value,
          $Res Function(_VersionResultPending) then) =
      __$VersionResultPendingCopyWithImpl<$Res>;
}

class __$VersionResultPendingCopyWithImpl<$Res>
    extends _$VersionResultCopyWithImpl<$Res>
    implements _$VersionResultPendingCopyWith<$Res> {
  __$VersionResultPendingCopyWithImpl(
      _VersionResultPending _value, $Res Function(_VersionResultPending) _then)
      : super(_value, (v) => _then(v as _VersionResultPending));

  @override
  _VersionResultPending get _value => super._value as _VersionResultPending;
}

@JsonSerializable()
class _$_VersionResultPending
    with DiagnosticableTreeMixin
    implements _VersionResultPending {
  const _$_VersionResultPending();

  factory _$_VersionResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_VersionResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'VersionResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'VersionResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _VersionResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_VersionResult value),
    @required Result failed(_VersionResultFailed value),
    @required Result pending(_VersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_VersionResult value),
    Result failed(_VersionResultFailed value),
    Result pending(_VersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_VersionResultPendingToJson(this)..['runtimeType'] = 'pending';
  }
}

abstract class _VersionResultPending implements VersionResult {
  const factory _VersionResultPending() = _$_VersionResultPending;

  factory _VersionResultPending.fromJson(Map<String, dynamic> json) =
      _$_VersionResultPending.fromJson;
}

GetUserResult _$GetUserResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _GetUserResult.fromJson(json);
    case 'failed':
      return _GetUserResultFailed.fromJson(json);
    case 'pending':
      return _GetUserResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetUserResultTearOff {
  const _$GetUserResultTearOff();

  _GetUserResult data({@required int code, @required User user}) {
    return _GetUserResult(
      code: code,
      user: user,
    );
  }

  _GetUserResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _GetUserResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _GetUserResultPending pending() {
    return const _GetUserResultPending();
  }
}

// ignore: unused_element
const $GetUserResult = _$GetUserResultTearOff();

mixin _$GetUserResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUserResult value),
    @required Result failed(_GetUserResultFailed value),
    @required Result pending(_GetUserResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUserResult value),
    Result failed(_GetUserResultFailed value),
    Result pending(_GetUserResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetUserResultCopyWith<$Res> {
  factory $GetUserResultCopyWith(
          GetUserResult value, $Res Function(GetUserResult) then) =
      _$GetUserResultCopyWithImpl<$Res>;
}

class _$GetUserResultCopyWithImpl<$Res>
    implements $GetUserResultCopyWith<$Res> {
  _$GetUserResultCopyWithImpl(this._value, this._then);

  final GetUserResult _value;
  // ignore: unused_field
  final $Res Function(GetUserResult) _then;
}

abstract class _$GetUserResultCopyWith<$Res> {
  factory _$GetUserResultCopyWith(
          _GetUserResult value, $Res Function(_GetUserResult) then) =
      __$GetUserResultCopyWithImpl<$Res>;
  $Res call({int code, User user});

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
    return 'GetUserResult.data(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUserResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(code, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(code, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUserResult value),
    @required Result failed(_GetUserResultFailed value),
    @required Result pending(_GetUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUserResult value),
    Result failed(_GetUserResultFailed value),
    Result pending(_GetUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _GetUserResult implements GetUserResult {
  factory _GetUserResult({@required int code, @required User user}) =
      _$_GetUserResult;

  factory _GetUserResult.fromJson(Map<String, dynamic> json) =
      _$_GetUserResult.fromJson;

  int get code;
  User get user;
  _$GetUserResultCopyWith<_GetUserResult> get copyWith;
}

abstract class _$GetUserResultFailedCopyWith<$Res> {
  factory _$GetUserResultFailedCopyWith(_GetUserResultFailed value,
          $Res Function(_GetUserResultFailed) then) =
      __$GetUserResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$GetUserResultFailedCopyWithImpl<$Res>
    extends _$GetUserResultCopyWithImpl<$Res>
    implements _$GetUserResultFailedCopyWith<$Res> {
  __$GetUserResultFailedCopyWithImpl(
      _GetUserResultFailed _value, $Res Function(_GetUserResultFailed) _then)
      : super(_value, (v) => _then(v as _GetUserResultFailed));

  @override
  _GetUserResultFailed get _value => super._value as _GetUserResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_GetUserResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_GetUserResultFailed
    with DiagnosticableTreeMixin
    implements _GetUserResultFailed {
  _$_GetUserResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_GetUserResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUserResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUserResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$GetUserResultFailedCopyWith<_GetUserResultFailed> get copyWith =>
      __$GetUserResultFailedCopyWithImpl<_GetUserResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUserResult value),
    @required Result failed(_GetUserResultFailed value),
    @required Result pending(_GetUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUserResult value),
    Result failed(_GetUserResultFailed value),
    Result pending(_GetUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _GetUserResultFailed implements GetUserResult {
  factory _GetUserResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_GetUserResultFailed;

  factory _GetUserResultFailed.fromJson(Map<String, dynamic> json) =
      _$_GetUserResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$GetUserResultFailedCopyWith<_GetUserResultFailed> get copyWith;
}

abstract class _$GetUserResultPendingCopyWith<$Res> {
  factory _$GetUserResultPendingCopyWith(_GetUserResultPending value,
          $Res Function(_GetUserResultPending) then) =
      __$GetUserResultPendingCopyWithImpl<$Res>;
}

class __$GetUserResultPendingCopyWithImpl<$Res>
    extends _$GetUserResultCopyWithImpl<$Res>
    implements _$GetUserResultPendingCopyWith<$Res> {
  __$GetUserResultPendingCopyWithImpl(
      _GetUserResultPending _value, $Res Function(_GetUserResultPending) _then)
      : super(_value, (v) => _then(v as _GetUserResultPending));

  @override
  _GetUserResultPending get _value => super._value as _GetUserResultPending;
}

@JsonSerializable()
class _$_GetUserResultPending
    with DiagnosticableTreeMixin
    implements _GetUserResultPending {
  const _$_GetUserResultPending();

  factory _$_GetUserResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUserResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'GetUserResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _GetUserResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_GetUserResult value),
    @required Result failed(_GetUserResultFailed value),
    @required Result pending(_GetUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUserResult value),
    Result failed(_GetUserResultFailed value),
    Result pending(_GetUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUserResultPendingToJson(this)..['runtimeType'] = 'pending';
  }
}

abstract class _GetUserResultPending implements GetUserResult {
  const factory _GetUserResultPending() = _$_GetUserResultPending;

  factory _GetUserResultPending.fromJson(Map<String, dynamic> json) =
      _$_GetUserResultPending.fromJson;
}

FindUserResult _$FindUserResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _FindUserResult.fromJson(json);
    case 'failed':
      return _FindUserResultFailed.fromJson(json);
    case 'pending':
      return _FindUserResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$FindUserResultTearOff {
  const _$FindUserResultTearOff();

  _FindUserResult data({@required String name, @required User user}) {
    return _FindUserResult(
      name: name,
      user: user,
    );
  }

  _FindUserResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _FindUserResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _FindUserResultPending pending() {
    return const _FindUserResultPending();
  }
}

// ignore: unused_element
const $FindUserResult = _$FindUserResultTearOff();

mixin _$FindUserResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_FindUserResult value),
    @required Result failed(_FindUserResultFailed value),
    @required Result pending(_FindUserResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_FindUserResult value),
    Result failed(_FindUserResultFailed value),
    Result pending(_FindUserResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $FindUserResultCopyWith<$Res> {
  factory $FindUserResultCopyWith(
          FindUserResult value, $Res Function(FindUserResult) then) =
      _$FindUserResultCopyWithImpl<$Res>;
}

class _$FindUserResultCopyWithImpl<$Res>
    implements $FindUserResultCopyWith<$Res> {
  _$FindUserResultCopyWithImpl(this._value, this._then);

  final FindUserResult _value;
  // ignore: unused_field
  final $Res Function(FindUserResult) _then;
}

abstract class _$FindUserResultCopyWith<$Res> {
  factory _$FindUserResultCopyWith(
          _FindUserResult value, $Res Function(_FindUserResult) then) =
      __$FindUserResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

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
    return 'FindUserResult.data(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'FindUserResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(name, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(name, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_FindUserResult value),
    @required Result failed(_FindUserResultFailed value),
    @required Result pending(_FindUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_FindUserResult value),
    Result failed(_FindUserResultFailed value),
    Result pending(_FindUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _FindUserResult implements FindUserResult {
  factory _FindUserResult({@required String name, @required User user}) =
      _$_FindUserResult;

  factory _FindUserResult.fromJson(Map<String, dynamic> json) =
      _$_FindUserResult.fromJson;

  String get name;
  User get user;
  _$FindUserResultCopyWith<_FindUserResult> get copyWith;
}

abstract class _$FindUserResultFailedCopyWith<$Res> {
  factory _$FindUserResultFailedCopyWith(_FindUserResultFailed value,
          $Res Function(_FindUserResultFailed) then) =
      __$FindUserResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$FindUserResultFailedCopyWithImpl<$Res>
    extends _$FindUserResultCopyWithImpl<$Res>
    implements _$FindUserResultFailedCopyWith<$Res> {
  __$FindUserResultFailedCopyWithImpl(
      _FindUserResultFailed _value, $Res Function(_FindUserResultFailed) _then)
      : super(_value, (v) => _then(v as _FindUserResultFailed));

  @override
  _FindUserResultFailed get _value => super._value as _FindUserResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_FindUserResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_FindUserResultFailed
    with DiagnosticableTreeMixin
    implements _FindUserResultFailed {
  _$_FindUserResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_FindUserResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'FindUserResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'FindUserResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$FindUserResultFailedCopyWith<_FindUserResultFailed> get copyWith =>
      __$FindUserResultFailedCopyWithImpl<_FindUserResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_FindUserResult value),
    @required Result failed(_FindUserResultFailed value),
    @required Result pending(_FindUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_FindUserResult value),
    Result failed(_FindUserResultFailed value),
    Result pending(_FindUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _FindUserResultFailed implements FindUserResult {
  factory _FindUserResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_FindUserResultFailed;

  factory _FindUserResultFailed.fromJson(Map<String, dynamic> json) =
      _$_FindUserResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$FindUserResultFailedCopyWith<_FindUserResultFailed> get copyWith;
}

abstract class _$FindUserResultPendingCopyWith<$Res> {
  factory _$FindUserResultPendingCopyWith(_FindUserResultPending value,
          $Res Function(_FindUserResultPending) then) =
      __$FindUserResultPendingCopyWithImpl<$Res>;
}

class __$FindUserResultPendingCopyWithImpl<$Res>
    extends _$FindUserResultCopyWithImpl<$Res>
    implements _$FindUserResultPendingCopyWith<$Res> {
  __$FindUserResultPendingCopyWithImpl(_FindUserResultPending _value,
      $Res Function(_FindUserResultPending) _then)
      : super(_value, (v) => _then(v as _FindUserResultPending));

  @override
  _FindUserResultPending get _value => super._value as _FindUserResultPending;
}

@JsonSerializable()
class _$_FindUserResultPending
    with DiagnosticableTreeMixin
    implements _FindUserResultPending {
  const _$_FindUserResultPending();

  factory _$_FindUserResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'FindUserResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'FindUserResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _FindUserResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_FindUserResult value),
    @required Result failed(_FindUserResultFailed value),
    @required Result pending(_FindUserResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_FindUserResult value),
    Result failed(_FindUserResultFailed value),
    Result pending(_FindUserResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_FindUserResultPendingToJson(this)..['runtimeType'] = 'pending';
  }
}

abstract class _FindUserResultPending implements FindUserResult {
  const factory _FindUserResultPending() = _$_FindUserResultPending;

  factory _FindUserResultPending.fromJson(Map<String, dynamic> json) =
      _$_FindUserResultPending.fromJson;
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
  switch (json['runtimeType'] as String) {
    case 'data':
      return _AnotherExampleServiceStatusResult.fromJson(json);
    case 'failed':
      return _AnotherExampleServiceStatusResultFailed.fromJson(json);
    case 'pending':
      return _AnotherExampleServiceStatusResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$AnotherExampleServiceStatusResultTearOff {
  const _$AnotherExampleServiceStatusResultTearOff();

  _AnotherExampleServiceStatusResult data({@required bool status}) {
    return _AnotherExampleServiceStatusResult(
      status: status,
    );
  }

  _AnotherExampleServiceStatusResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _AnotherExampleServiceStatusResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _AnotherExampleServiceStatusResultPending pending() {
    return const _AnotherExampleServiceStatusResultPending();
  }
}

// ignore: unused_element
const $AnotherExampleServiceStatusResult =
    _$AnotherExampleServiceStatusResultTearOff();

mixin _$AnotherExampleServiceStatusResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_AnotherExampleServiceStatusResult value),
    @required Result failed(_AnotherExampleServiceStatusResultFailed value),
    @required Result pending(_AnotherExampleServiceStatusResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_AnotherExampleServiceStatusResult value),
    Result failed(_AnotherExampleServiceStatusResultFailed value),
    Result pending(_AnotherExampleServiceStatusResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $AnotherExampleServiceStatusResultCopyWith<$Res> {
  factory $AnotherExampleServiceStatusResultCopyWith(
          AnotherExampleServiceStatusResult value,
          $Res Function(AnotherExampleServiceStatusResult) then) =
      _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>;
}

class _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    implements $AnotherExampleServiceStatusResultCopyWith<$Res> {
  _$AnotherExampleServiceStatusResultCopyWithImpl(this._value, this._then);

  final AnotherExampleServiceStatusResult _value;
  // ignore: unused_field
  final $Res Function(AnotherExampleServiceStatusResult) _then;
}

abstract class _$AnotherExampleServiceStatusResultCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResultCopyWith(
          _AnotherExampleServiceStatusResult value,
          $Res Function(_AnotherExampleServiceStatusResult) then) =
      __$AnotherExampleServiceStatusResultCopyWithImpl<$Res>;
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

@JsonSerializable()
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
    return 'AnotherExampleServiceStatusResult.data(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(
          DiagnosticsProperty('type', 'AnotherExampleServiceStatusResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(status);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(status);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_AnotherExampleServiceStatusResult value),
    @required Result failed(_AnotherExampleServiceStatusResultFailed value),
    @required Result pending(_AnotherExampleServiceStatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_AnotherExampleServiceStatusResult value),
    Result failed(_AnotherExampleServiceStatusResultFailed value),
    Result pending(_AnotherExampleServiceStatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResultToJson(this)
      ..['runtimeType'] = 'data';
  }
}

abstract class _AnotherExampleServiceStatusResult
    implements AnotherExampleServiceStatusResult {
  factory _AnotherExampleServiceStatusResult({@required bool status}) =
      _$_AnotherExampleServiceStatusResult;

  factory _AnotherExampleServiceStatusResult.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResult.fromJson;

  bool get status;
  _$AnotherExampleServiceStatusResultCopyWith<
      _AnotherExampleServiceStatusResult> get copyWith;
}

abstract class _$AnotherExampleServiceStatusResultFailedCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResultFailedCopyWith(
          _AnotherExampleServiceStatusResultFailed value,
          $Res Function(_AnotherExampleServiceStatusResultFailed) then) =
      __$AnotherExampleServiceStatusResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$AnotherExampleServiceStatusResultFailedCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusResultFailedCopyWith<$Res> {
  __$AnotherExampleServiceStatusResultFailedCopyWithImpl(
      _AnotherExampleServiceStatusResultFailed _value,
      $Res Function(_AnotherExampleServiceStatusResultFailed) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceStatusResultFailed));

  @override
  _AnotherExampleServiceStatusResultFailed get _value =>
      super._value as _AnotherExampleServiceStatusResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_AnotherExampleServiceStatusResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_AnotherExampleServiceStatusResultFailed
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceStatusResultFailed {
  _$_AnotherExampleServiceStatusResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_AnotherExampleServiceStatusResultFailed.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AnotherExampleServiceStatusResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty(
          'type', 'AnotherExampleServiceStatusResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$AnotherExampleServiceStatusResultFailedCopyWith<
          _AnotherExampleServiceStatusResultFailed>
      get copyWith => __$AnotherExampleServiceStatusResultFailedCopyWithImpl<
          _AnotherExampleServiceStatusResultFailed>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_AnotherExampleServiceStatusResult value),
    @required Result failed(_AnotherExampleServiceStatusResultFailed value),
    @required Result pending(_AnotherExampleServiceStatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_AnotherExampleServiceStatusResult value),
    Result failed(_AnotherExampleServiceStatusResultFailed value),
    Result pending(_AnotherExampleServiceStatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResultFailedToJson(this)
      ..['runtimeType'] = 'failed';
  }
}

abstract class _AnotherExampleServiceStatusResultFailed
    implements AnotherExampleServiceStatusResult {
  factory _AnotherExampleServiceStatusResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_AnotherExampleServiceStatusResultFailed;

  factory _AnotherExampleServiceStatusResultFailed.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$AnotherExampleServiceStatusResultFailedCopyWith<
      _AnotherExampleServiceStatusResultFailed> get copyWith;
}

abstract class _$AnotherExampleServiceStatusResultPendingCopyWith<$Res> {
  factory _$AnotherExampleServiceStatusResultPendingCopyWith(
          _AnotherExampleServiceStatusResultPending value,
          $Res Function(_AnotherExampleServiceStatusResultPending) then) =
      __$AnotherExampleServiceStatusResultPendingCopyWithImpl<$Res>;
}

class __$AnotherExampleServiceStatusResultPendingCopyWithImpl<$Res>
    extends _$AnotherExampleServiceStatusResultCopyWithImpl<$Res>
    implements _$AnotherExampleServiceStatusResultPendingCopyWith<$Res> {
  __$AnotherExampleServiceStatusResultPendingCopyWithImpl(
      _AnotherExampleServiceStatusResultPending _value,
      $Res Function(_AnotherExampleServiceStatusResultPending) _then)
      : super(_value,
            (v) => _then(v as _AnotherExampleServiceStatusResultPending));

  @override
  _AnotherExampleServiceStatusResultPending get _value =>
      super._value as _AnotherExampleServiceStatusResultPending;
}

@JsonSerializable()
class _$_AnotherExampleServiceStatusResultPending
    with DiagnosticableTreeMixin
    implements _AnotherExampleServiceStatusResultPending {
  const _$_AnotherExampleServiceStatusResultPending();

  factory _$_AnotherExampleServiceStatusResultPending.fromJson(
          Map<String, dynamic> json) =>
      _$_$_AnotherExampleServiceStatusResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'AnotherExampleServiceStatusResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty(
          'type', 'AnotherExampleServiceStatusResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AnotherExampleServiceStatusResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(bool status),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(bool status),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_AnotherExampleServiceStatusResult value),
    @required Result failed(_AnotherExampleServiceStatusResultFailed value),
    @required Result pending(_AnotherExampleServiceStatusResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_AnotherExampleServiceStatusResult value),
    Result failed(_AnotherExampleServiceStatusResultFailed value),
    Result pending(_AnotherExampleServiceStatusResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_AnotherExampleServiceStatusResultPendingToJson(this)
      ..['runtimeType'] = 'pending';
  }
}

abstract class _AnotherExampleServiceStatusResultPending
    implements AnotherExampleServiceStatusResult {
  const factory _AnotherExampleServiceStatusResultPending() =
      _$_AnotherExampleServiceStatusResultPending;

  factory _AnotherExampleServiceStatusResultPending.fromJson(
          Map<String, dynamic> json) =
      _$_AnotherExampleServiceStatusResultPending.fromJson;
}

GetVersionResult _$GetVersionResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _GetVersionResult.fromJson(json);
    case 'failed':
      return _GetVersionResultFailed.fromJson(json);
    case 'pending':
      return _GetVersionResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetVersionResultTearOff {
  const _$GetVersionResultTearOff();

  _GetVersionResult data({@required Version version}) {
    return _GetVersionResult(
      version: version,
    );
  }

  _GetVersionResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _GetVersionResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _GetVersionResultPending pending() {
    return const _GetVersionResultPending();
  }
}

// ignore: unused_element
const $GetVersionResult = _$GetVersionResultTearOff();

mixin _$GetVersionResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetVersionResult value),
    @required Result failed(_GetVersionResultFailed value),
    @required Result pending(_GetVersionResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetVersionResult value),
    Result failed(_GetVersionResultFailed value),
    Result pending(_GetVersionResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetVersionResultCopyWith<$Res> {
  factory $GetVersionResultCopyWith(
          GetVersionResult value, $Res Function(GetVersionResult) then) =
      _$GetVersionResultCopyWithImpl<$Res>;
}

class _$GetVersionResultCopyWithImpl<$Res>
    implements $GetVersionResultCopyWith<$Res> {
  _$GetVersionResultCopyWithImpl(this._value, this._then);

  final GetVersionResult _value;
  // ignore: unused_field
  final $Res Function(GetVersionResult) _then;
}

abstract class _$GetVersionResultCopyWith<$Res> {
  factory _$GetVersionResultCopyWith(
          _GetVersionResult value, $Res Function(_GetVersionResult) then) =
      __$GetVersionResultCopyWithImpl<$Res>;
  $Res call({Version version});

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
    return 'GetVersionResult.data(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetVersionResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(version);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(version);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetVersionResult value),
    @required Result failed(_GetVersionResultFailed value),
    @required Result pending(_GetVersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetVersionResult value),
    Result failed(_GetVersionResultFailed value),
    Result pending(_GetVersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _GetVersionResult implements GetVersionResult {
  factory _GetVersionResult({@required Version version}) = _$_GetVersionResult;

  factory _GetVersionResult.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResult.fromJson;

  Version get version;
  _$GetVersionResultCopyWith<_GetVersionResult> get copyWith;
}

abstract class _$GetVersionResultFailedCopyWith<$Res> {
  factory _$GetVersionResultFailedCopyWith(_GetVersionResultFailed value,
          $Res Function(_GetVersionResultFailed) then) =
      __$GetVersionResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$GetVersionResultFailedCopyWithImpl<$Res>
    extends _$GetVersionResultCopyWithImpl<$Res>
    implements _$GetVersionResultFailedCopyWith<$Res> {
  __$GetVersionResultFailedCopyWithImpl(_GetVersionResultFailed _value,
      $Res Function(_GetVersionResultFailed) _then)
      : super(_value, (v) => _then(v as _GetVersionResultFailed));

  @override
  _GetVersionResultFailed get _value => super._value as _GetVersionResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_GetVersionResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_GetVersionResultFailed
    with DiagnosticableTreeMixin
    implements _GetVersionResultFailed {
  _$_GetVersionResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_GetVersionResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetVersionResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetVersionResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetVersionResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetVersionResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$GetVersionResultFailedCopyWith<_GetVersionResultFailed> get copyWith =>
      __$GetVersionResultFailedCopyWithImpl<_GetVersionResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetVersionResult value),
    @required Result failed(_GetVersionResultFailed value),
    @required Result pending(_GetVersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetVersionResult value),
    Result failed(_GetVersionResultFailed value),
    Result pending(_GetVersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _GetVersionResultFailed implements GetVersionResult {
  factory _GetVersionResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_GetVersionResultFailed;

  factory _GetVersionResultFailed.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$GetVersionResultFailedCopyWith<_GetVersionResultFailed> get copyWith;
}

abstract class _$GetVersionResultPendingCopyWith<$Res> {
  factory _$GetVersionResultPendingCopyWith(_GetVersionResultPending value,
          $Res Function(_GetVersionResultPending) then) =
      __$GetVersionResultPendingCopyWithImpl<$Res>;
}

class __$GetVersionResultPendingCopyWithImpl<$Res>
    extends _$GetVersionResultCopyWithImpl<$Res>
    implements _$GetVersionResultPendingCopyWith<$Res> {
  __$GetVersionResultPendingCopyWithImpl(_GetVersionResultPending _value,
      $Res Function(_GetVersionResultPending) _then)
      : super(_value, (v) => _then(v as _GetVersionResultPending));

  @override
  _GetVersionResultPending get _value =>
      super._value as _GetVersionResultPending;
}

@JsonSerializable()
class _$_GetVersionResultPending
    with DiagnosticableTreeMixin
    implements _GetVersionResultPending {
  const _$_GetVersionResultPending();

  factory _$_GetVersionResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_GetVersionResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetVersionResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'GetVersionResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _GetVersionResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(Version version),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(Version version),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_GetVersionResult value),
    @required Result failed(_GetVersionResultFailed value),
    @required Result pending(_GetVersionResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetVersionResult value),
    Result failed(_GetVersionResultFailed value),
    Result pending(_GetVersionResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetVersionResultPendingToJson(this)
      ..['runtimeType'] = 'pending';
  }
}

abstract class _GetVersionResultPending implements GetVersionResult {
  const factory _GetVersionResultPending() = _$_GetVersionResultPending;

  factory _GetVersionResultPending.fromJson(Map<String, dynamic> json) =
      _$_GetVersionResultPending.fromJson;
}

GetAccountResult _$GetAccountResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _GetAccountResult.fromJson(json);
    case 'failed':
      return _GetAccountResultFailed.fromJson(json);
    case 'pending':
      return _GetAccountResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetAccountResultTearOff {
  const _$GetAccountResultTearOff();

  _GetAccountResult data({@required int code, @required User user}) {
    return _GetAccountResult(
      code: code,
      user: user,
    );
  }

  _GetAccountResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _GetAccountResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _GetAccountResultPending pending() {
    return const _GetAccountResultPending();
  }
}

// ignore: unused_element
const $GetAccountResult = _$GetAccountResultTearOff();

mixin _$GetAccountResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetAccountResult value),
    @required Result failed(_GetAccountResultFailed value),
    @required Result pending(_GetAccountResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetAccountResult value),
    Result failed(_GetAccountResultFailed value),
    Result pending(_GetAccountResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetAccountResultCopyWith<$Res> {
  factory $GetAccountResultCopyWith(
          GetAccountResult value, $Res Function(GetAccountResult) then) =
      _$GetAccountResultCopyWithImpl<$Res>;
}

class _$GetAccountResultCopyWithImpl<$Res>
    implements $GetAccountResultCopyWith<$Res> {
  _$GetAccountResultCopyWithImpl(this._value, this._then);

  final GetAccountResult _value;
  // ignore: unused_field
  final $Res Function(GetAccountResult) _then;
}

abstract class _$GetAccountResultCopyWith<$Res> {
  factory _$GetAccountResultCopyWith(
          _GetAccountResult value, $Res Function(_GetAccountResult) then) =
      __$GetAccountResultCopyWithImpl<$Res>;
  $Res call({int code, User user});

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
    return 'GetAccountResult.data(code: $code, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetAccountResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(code, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(code, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetAccountResult value),
    @required Result failed(_GetAccountResultFailed value),
    @required Result pending(_GetAccountResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetAccountResult value),
    Result failed(_GetAccountResultFailed value),
    Result pending(_GetAccountResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _GetAccountResult implements GetAccountResult {
  factory _GetAccountResult({@required int code, @required User user}) =
      _$_GetAccountResult;

  factory _GetAccountResult.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResult.fromJson;

  int get code;
  User get user;
  _$GetAccountResultCopyWith<_GetAccountResult> get copyWith;
}

abstract class _$GetAccountResultFailedCopyWith<$Res> {
  factory _$GetAccountResultFailedCopyWith(_GetAccountResultFailed value,
          $Res Function(_GetAccountResultFailed) then) =
      __$GetAccountResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$GetAccountResultFailedCopyWithImpl<$Res>
    extends _$GetAccountResultCopyWithImpl<$Res>
    implements _$GetAccountResultFailedCopyWith<$Res> {
  __$GetAccountResultFailedCopyWithImpl(_GetAccountResultFailed _value,
      $Res Function(_GetAccountResultFailed) _then)
      : super(_value, (v) => _then(v as _GetAccountResultFailed));

  @override
  _GetAccountResultFailed get _value => super._value as _GetAccountResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_GetAccountResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_GetAccountResultFailed
    with DiagnosticableTreeMixin
    implements _GetAccountResultFailed {
  _$_GetAccountResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_GetAccountResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetAccountResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetAccountResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetAccountResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetAccountResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$GetAccountResultFailedCopyWith<_GetAccountResultFailed> get copyWith =>
      __$GetAccountResultFailedCopyWithImpl<_GetAccountResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetAccountResult value),
    @required Result failed(_GetAccountResultFailed value),
    @required Result pending(_GetAccountResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetAccountResult value),
    Result failed(_GetAccountResultFailed value),
    Result pending(_GetAccountResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _GetAccountResultFailed implements GetAccountResult {
  factory _GetAccountResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_GetAccountResultFailed;

  factory _GetAccountResultFailed.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$GetAccountResultFailedCopyWith<_GetAccountResultFailed> get copyWith;
}

abstract class _$GetAccountResultPendingCopyWith<$Res> {
  factory _$GetAccountResultPendingCopyWith(_GetAccountResultPending value,
          $Res Function(_GetAccountResultPending) then) =
      __$GetAccountResultPendingCopyWithImpl<$Res>;
}

class __$GetAccountResultPendingCopyWithImpl<$Res>
    extends _$GetAccountResultCopyWithImpl<$Res>
    implements _$GetAccountResultPendingCopyWith<$Res> {
  __$GetAccountResultPendingCopyWithImpl(_GetAccountResultPending _value,
      $Res Function(_GetAccountResultPending) _then)
      : super(_value, (v) => _then(v as _GetAccountResultPending));

  @override
  _GetAccountResultPending get _value =>
      super._value as _GetAccountResultPending;
}

@JsonSerializable()
class _$_GetAccountResultPending
    with DiagnosticableTreeMixin
    implements _GetAccountResultPending {
  const _$_GetAccountResultPending();

  factory _$_GetAccountResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_GetAccountResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetAccountResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'GetAccountResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _GetAccountResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(int code, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(int code, User user),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_GetAccountResult value),
    @required Result failed(_GetAccountResultFailed value),
    @required Result pending(_GetAccountResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetAccountResult value),
    Result failed(_GetAccountResultFailed value),
    Result pending(_GetAccountResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetAccountResultPendingToJson(this)
      ..['runtimeType'] = 'pending';
  }
}

abstract class _GetAccountResultPending implements GetAccountResult {
  const factory _GetAccountResultPending() = _$_GetAccountResultPending;

  factory _GetAccountResultPending.fromJson(Map<String, dynamic> json) =
      _$_GetAccountResultPending.fromJson;
}

GetUsersResult _$GetUsersResultFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'data':
      return _GetUsersResult.fromJson(json);
    case 'failed':
      return _GetUsersResultFailed.fromJson(json);
    case 'pending':
      return _GetUsersResultPending.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$GetUsersResultTearOff {
  const _$GetUsersResultTearOff();

  _GetUsersResult data({@required String name, @required User user}) {
    return _GetUsersResult(
      name: name,
      user: user,
    );
  }

  _GetUsersResultFailed failed(
      {@required int statusCode, @required String reason, String stackTrace}) {
    return _GetUsersResultFailed(
      statusCode: statusCode,
      reason: reason,
      stackTrace: stackTrace,
    );
  }

  _GetUsersResultPending pending() {
    return const _GetUsersResultPending();
  }
}

// ignore: unused_element
const $GetUsersResult = _$GetUsersResultTearOff();

mixin _$GetUsersResult {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUsersResult value),
    @required Result failed(_GetUsersResultFailed value),
    @required Result pending(_GetUsersResultPending value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUsersResult value),
    Result failed(_GetUsersResultFailed value),
    Result pending(_GetUsersResultPending value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $GetUsersResultCopyWith<$Res> {
  factory $GetUsersResultCopyWith(
          GetUsersResult value, $Res Function(GetUsersResult) then) =
      _$GetUsersResultCopyWithImpl<$Res>;
}

class _$GetUsersResultCopyWithImpl<$Res>
    implements $GetUsersResultCopyWith<$Res> {
  _$GetUsersResultCopyWithImpl(this._value, this._then);

  final GetUsersResult _value;
  // ignore: unused_field
  final $Res Function(GetUsersResult) _then;
}

abstract class _$GetUsersResultCopyWith<$Res> {
  factory _$GetUsersResultCopyWith(
          _GetUsersResult value, $Res Function(_GetUsersResult) then) =
      __$GetUsersResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

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
    return 'GetUsersResult.data(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUsersResult.data'))
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
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(name, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(name, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUsersResult value),
    @required Result failed(_GetUsersResultFailed value),
    @required Result pending(_GetUsersResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return data(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUsersResult value),
    Result failed(_GetUsersResultFailed value),
    Result pending(_GetUsersResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (data != null) {
      return data(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResultToJson(this)..['runtimeType'] = 'data';
  }
}

abstract class _GetUsersResult implements GetUsersResult {
  factory _GetUsersResult({@required String name, @required User user}) =
      _$_GetUsersResult;

  factory _GetUsersResult.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResult.fromJson;

  String get name;
  User get user;
  _$GetUsersResultCopyWith<_GetUsersResult> get copyWith;
}

abstract class _$GetUsersResultFailedCopyWith<$Res> {
  factory _$GetUsersResultFailedCopyWith(_GetUsersResultFailed value,
          $Res Function(_GetUsersResultFailed) then) =
      __$GetUsersResultFailedCopyWithImpl<$Res>;
  $Res call({int statusCode, String reason, String stackTrace});
}

class __$GetUsersResultFailedCopyWithImpl<$Res>
    extends _$GetUsersResultCopyWithImpl<$Res>
    implements _$GetUsersResultFailedCopyWith<$Res> {
  __$GetUsersResultFailedCopyWithImpl(
      _GetUsersResultFailed _value, $Res Function(_GetUsersResultFailed) _then)
      : super(_value, (v) => _then(v as _GetUsersResultFailed));

  @override
  _GetUsersResultFailed get _value => super._value as _GetUsersResultFailed;

  @override
  $Res call({
    Object statusCode = freezed,
    Object reason = freezed,
    Object stackTrace = freezed,
  }) {
    return _then(_GetUsersResultFailed(
      statusCode: statusCode == freezed ? _value.statusCode : statusCode as int,
      reason: reason == freezed ? _value.reason : reason as String,
      stackTrace:
          stackTrace == freezed ? _value.stackTrace : stackTrace as String,
    ));
  }
}

@JsonSerializable()
class _$_GetUsersResultFailed
    with DiagnosticableTreeMixin
    implements _GetUsersResultFailed {
  _$_GetUsersResultFailed(
      {@required this.statusCode, @required this.reason, this.stackTrace})
      : assert(statusCode != null),
        assert(reason != null);

  factory _$_GetUsersResultFailed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUsersResultFailedFromJson(json);

  @override
  final int statusCode;
  @override
  final String reason;
  @override
  final String stackTrace;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUsersResult.failed(statusCode: $statusCode, reason: $reason, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'GetUsersResult.failed'))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUsersResultFailed &&
            (identical(other.statusCode, statusCode) ||
                const DeepCollectionEquality()
                    .equals(other.statusCode, statusCode)) &&
            (identical(other.reason, reason) ||
                const DeepCollectionEquality().equals(other.reason, reason)) &&
            (identical(other.stackTrace, stackTrace) ||
                const DeepCollectionEquality()
                    .equals(other.stackTrace, stackTrace)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(statusCode) ^
      const DeepCollectionEquality().hash(reason) ^
      const DeepCollectionEquality().hash(stackTrace);

  @override
  _$GetUsersResultFailedCopyWith<_GetUsersResultFailed> get copyWith =>
      __$GetUsersResultFailedCopyWithImpl<_GetUsersResultFailed>(
          this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(statusCode, reason, stackTrace);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
    Result pending(),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(statusCode, reason, stackTrace);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result data(_GetUsersResult value),
    @required Result failed(_GetUsersResultFailed value),
    @required Result pending(_GetUsersResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return failed(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUsersResult value),
    Result failed(_GetUsersResultFailed value),
    Result pending(_GetUsersResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (failed != null) {
      return failed(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResultFailedToJson(this)..['runtimeType'] = 'failed';
  }
}

abstract class _GetUsersResultFailed implements GetUsersResult {
  factory _GetUsersResultFailed(
      {@required int statusCode,
      @required String reason,
      String stackTrace}) = _$_GetUsersResultFailed;

  factory _GetUsersResultFailed.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResultFailed.fromJson;

  int get statusCode;
  String get reason;
  String get stackTrace;
  _$GetUsersResultFailedCopyWith<_GetUsersResultFailed> get copyWith;
}

abstract class _$GetUsersResultPendingCopyWith<$Res> {
  factory _$GetUsersResultPendingCopyWith(_GetUsersResultPending value,
          $Res Function(_GetUsersResultPending) then) =
      __$GetUsersResultPendingCopyWithImpl<$Res>;
}

class __$GetUsersResultPendingCopyWithImpl<$Res>
    extends _$GetUsersResultCopyWithImpl<$Res>
    implements _$GetUsersResultPendingCopyWith<$Res> {
  __$GetUsersResultPendingCopyWithImpl(_GetUsersResultPending _value,
      $Res Function(_GetUsersResultPending) _then)
      : super(_value, (v) => _then(v as _GetUsersResultPending));

  @override
  _GetUsersResultPending get _value => super._value as _GetUsersResultPending;
}

@JsonSerializable()
class _$_GetUsersResultPending
    with DiagnosticableTreeMixin
    implements _GetUsersResultPending {
  const _$_GetUsersResultPending();

  factory _$_GetUsersResultPending.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUsersResultPendingFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'GetUsersResult.pending()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'GetUsersResult.pending'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is _GetUsersResultPending);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result data(String name, User user),
    @required Result failed(int statusCode, String reason, String stackTrace),
    @required Result pending(),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result data(String name, User user),
    Result failed(int statusCode, String reason, String stackTrace),
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
    @required Result data(_GetUsersResult value),
    @required Result failed(_GetUsersResultFailed value),
    @required Result pending(_GetUsersResultPending value),
  }) {
    assert(data != null);
    assert(failed != null);
    assert(pending != null);
    return pending(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result data(_GetUsersResult value),
    Result failed(_GetUsersResultFailed value),
    Result pending(_GetUsersResultPending value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pending != null) {
      return pending(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$_GetUsersResultPendingToJson(this)..['runtimeType'] = 'pending';
  }
}

abstract class _GetUsersResultPending implements GetUsersResult {
  const factory _GetUsersResultPending() = _$_GetUsersResultPending;

  factory _GetUsersResultPending.fromJson(Map<String, dynamic> json) =
      _$_GetUsersResultPending.fromJson;
}
