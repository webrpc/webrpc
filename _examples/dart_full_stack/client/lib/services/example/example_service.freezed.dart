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

@JsonSerializable(explicitToJson: true)
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

@JsonSerializable(explicitToJson: true)
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
class _$_SearchFilter with DiagnosticableTreeMixin implements _SearchFilter {
  _$_SearchFilter({@required this.id}) : assert(id != null);

  factory _$_SearchFilter.fromJson(Map<String, dynamic> json) =>
      _$_$_SearchFilterFromJson(json);

  @override
  final int id;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'SearchFilter(id: $id)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'SearchFilter'))
      ..add(DiagnosticsProperty('id', id));
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

@JsonSerializable(explicitToJson: true)
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
class _$_GetUserArgs_Freezed
    with DiagnosticableTreeMixin
    implements _GetUserArgs_Freezed {
  _$_GetUserArgs_Freezed({@required this.userID}) : assert(userID != null);

  factory _$_GetUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_GetUserArgs_FreezedFromJson(json);

  @override
  final int userID;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_GetUserArgs(userID: $userID)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_GetUserArgs'))
      ..add(DiagnosticsProperty('userID', userID));
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
class _$_UpdateNameArgs_Freezed
    with DiagnosticableTreeMixin
    implements _UpdateNameArgs_Freezed {
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_UpdateNameArgs(id: $id, username: $username)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_UpdateNameArgs'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('username', username));
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
class _$_FindUserByIdArgs_Freezed
    with DiagnosticableTreeMixin
    implements _FindUserByIdArgs_Freezed {
  _$_FindUserByIdArgs_Freezed({@required this.s}) : assert(s != null);

  factory _$_FindUserByIdArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_FindUserByIdArgs_FreezedFromJson(json);

  @override
  final SearchFilter s;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_FindUserByIdArgs(s: $s)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_FindUserByIdArgs'))
      ..add(DiagnosticsProperty('s', s));
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
class _$_AddUserArgs_Freezed
    with DiagnosticableTreeMixin
    implements _AddUserArgs_Freezed {
  _$_AddUserArgs_Freezed({@required this.user}) : assert(user != null);

  factory _$_AddUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_AddUserArgs_FreezedFromJson(json);

  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_AddUserArgs(user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_AddUserArgs'))
      ..add(DiagnosticsProperty('user', user));
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
class _$_DeleteUserArgs_Freezed
    with DiagnosticableTreeMixin
    implements _DeleteUserArgs_Freezed {
  _$_DeleteUserArgs_Freezed({@required this.id}) : assert(id != null);

  factory _$_DeleteUserArgs_Freezed.fromJson(Map<String, dynamic> json) =>
      _$_$_DeleteUserArgs_FreezedFromJson(json);

  @override
  final int id;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return '_DeleteUserArgs(id: $id)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', '_DeleteUserArgs'))
      ..add(DiagnosticsProperty('id', id));
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

class _$ExampleServiceEventTearOff {
  const _$ExampleServiceEventTearOff();

  _PingEvent ping({Map<String, String> headers}) {
    return _PingEvent(
      headers: headers,
    );
  }

  _StatusEvent status({Map<String, String> headers}) {
    return _StatusEvent(
      headers: headers,
    );
  }

  _VersionEvent version({Map<String, String> headers}) {
    return _VersionEvent(
      headers: headers,
    );
  }

  _GetUserEvent getUser({@required int userID, Map<String, String> headers}) {
    return _GetUserEvent(
      userID: userID,
      headers: headers,
    );
  }

  _UpdateNameEvent updateName(
      {@required int id,
      @required String username,
      Map<String, String> headers}) {
    return _UpdateNameEvent(
      id: id,
      username: username,
      headers: headers,
    );
  }

  _FindUserByIdEvent findUserById(
      {@required SearchFilter s, Map<String, String> headers}) {
    return _FindUserByIdEvent(
      s: s,
      headers: headers,
    );
  }

  _AddUserEvent addUser({@required User user, Map<String, String> headers}) {
    return _AddUserEvent(
      user: user,
      headers: headers,
    );
  }

  _ListUsersEvent listUsers({Map<String, String> headers}) {
    return _ListUsersEvent(
      headers: headers,
    );
  }

  _DeleteUserEvent deleteUser({@required int id, Map<String, String> headers}) {
    return _DeleteUserEvent(
      id: id,
      headers: headers,
    );
  }
}

// ignore: unused_element
const $ExampleServiceEvent = _$ExampleServiceEventTearOff();

mixin _$ExampleServiceEvent {
  Map<String, String> get headers;

  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  });

  $ExampleServiceEventCopyWith<ExampleServiceEvent> get copyWith;
}

abstract class $ExampleServiceEventCopyWith<$Res> {
  factory $ExampleServiceEventCopyWith(
          ExampleServiceEvent value, $Res Function(ExampleServiceEvent) then) =
      _$ExampleServiceEventCopyWithImpl<$Res>;
  $Res call({Map<String, String> headers});
}

class _$ExampleServiceEventCopyWithImpl<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  _$ExampleServiceEventCopyWithImpl(this._value, this._then);

  final ExampleServiceEvent _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceEvent) _then;

  @override
  $Res call({
    Object headers = freezed,
  }) {
    return _then(_value.copyWith(
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

abstract class _$PingEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$PingEventCopyWith(
          _PingEvent value, $Res Function(_PingEvent) then) =
      __$PingEventCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> headers});
}

class __$PingEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$PingEventCopyWith<$Res> {
  __$PingEventCopyWithImpl(_PingEvent _value, $Res Function(_PingEvent) _then)
      : super(_value, (v) => _then(v as _PingEvent));

  @override
  _PingEvent get _value => super._value as _PingEvent;

  @override
  $Res call({
    Object headers = freezed,
  }) {
    return _then(_PingEvent(
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_PingEvent with DiagnosticableTreeMixin implements _PingEvent {
  _$_PingEvent({this.headers});

  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.ping(headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.ping'))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _PingEvent &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(headers);

  @override
  _$PingEventCopyWith<_PingEvent> get copyWith =>
      __$PingEventCopyWithImpl<_PingEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return ping(headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (ping != null) {
      return ping(headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return ping(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (ping != null) {
      return ping(this);
    }
    return orElse();
  }
}

abstract class _PingEvent implements ExampleServiceEvent {
  factory _PingEvent({Map<String, String> headers}) = _$_PingEvent;

  @override
  Map<String, String> get headers;
  @override
  _$PingEventCopyWith<_PingEvent> get copyWith;
}

abstract class _$StatusEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$StatusEventCopyWith(
          _StatusEvent value, $Res Function(_StatusEvent) then) =
      __$StatusEventCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> headers});
}

class __$StatusEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$StatusEventCopyWith<$Res> {
  __$StatusEventCopyWithImpl(
      _StatusEvent _value, $Res Function(_StatusEvent) _then)
      : super(_value, (v) => _then(v as _StatusEvent));

  @override
  _StatusEvent get _value => super._value as _StatusEvent;

  @override
  $Res call({
    Object headers = freezed,
  }) {
    return _then(_StatusEvent(
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_StatusEvent with DiagnosticableTreeMixin implements _StatusEvent {
  _$_StatusEvent({this.headers});

  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.status(headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.status'))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _StatusEvent &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(headers);

  @override
  _$StatusEventCopyWith<_StatusEvent> get copyWith =>
      __$StatusEventCopyWithImpl<_StatusEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return status(headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (status != null) {
      return status(headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return status(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (status != null) {
      return status(this);
    }
    return orElse();
  }
}

abstract class _StatusEvent implements ExampleServiceEvent {
  factory _StatusEvent({Map<String, String> headers}) = _$_StatusEvent;

  @override
  Map<String, String> get headers;
  @override
  _$StatusEventCopyWith<_StatusEvent> get copyWith;
}

abstract class _$VersionEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$VersionEventCopyWith(
          _VersionEvent value, $Res Function(_VersionEvent) then) =
      __$VersionEventCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> headers});
}

class __$VersionEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$VersionEventCopyWith<$Res> {
  __$VersionEventCopyWithImpl(
      _VersionEvent _value, $Res Function(_VersionEvent) _then)
      : super(_value, (v) => _then(v as _VersionEvent));

  @override
  _VersionEvent get _value => super._value as _VersionEvent;

  @override
  $Res call({
    Object headers = freezed,
  }) {
    return _then(_VersionEvent(
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_VersionEvent with DiagnosticableTreeMixin implements _VersionEvent {
  _$_VersionEvent({this.headers});

  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.version(headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.version'))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _VersionEvent &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(headers);

  @override
  _$VersionEventCopyWith<_VersionEvent> get copyWith =>
      __$VersionEventCopyWithImpl<_VersionEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return version(headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (version != null) {
      return version(headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return version(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (version != null) {
      return version(this);
    }
    return orElse();
  }
}

abstract class _VersionEvent implements ExampleServiceEvent {
  factory _VersionEvent({Map<String, String> headers}) = _$_VersionEvent;

  @override
  Map<String, String> get headers;
  @override
  _$VersionEventCopyWith<_VersionEvent> get copyWith;
}

abstract class _$GetUserEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$GetUserEventCopyWith(
          _GetUserEvent value, $Res Function(_GetUserEvent) then) =
      __$GetUserEventCopyWithImpl<$Res>;
  @override
  $Res call({int userID, Map<String, String> headers});
}

class __$GetUserEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$GetUserEventCopyWith<$Res> {
  __$GetUserEventCopyWithImpl(
      _GetUserEvent _value, $Res Function(_GetUserEvent) _then)
      : super(_value, (v) => _then(v as _GetUserEvent));

  @override
  _GetUserEvent get _value => super._value as _GetUserEvent;

  @override
  $Res call({
    Object userID = freezed,
    Object headers = freezed,
  }) {
    return _then(_GetUserEvent(
      userID: userID == freezed ? _value.userID : userID as int,
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_GetUserEvent with DiagnosticableTreeMixin implements _GetUserEvent {
  _$_GetUserEvent({@required this.userID, this.headers})
      : assert(userID != null);

  @override
  final int userID;
  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.getUser(userID: $userID, headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.getUser'))
      ..add(DiagnosticsProperty('userID', userID))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _GetUserEvent &&
            (identical(other.userID, userID) ||
                const DeepCollectionEquality().equals(other.userID, userID)) &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(userID) ^
      const DeepCollectionEquality().hash(headers);

  @override
  _$GetUserEventCopyWith<_GetUserEvent> get copyWith =>
      __$GetUserEventCopyWithImpl<_GetUserEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return getUser(userID, headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (getUser != null) {
      return getUser(userID, headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return getUser(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (getUser != null) {
      return getUser(this);
    }
    return orElse();
  }
}

abstract class _GetUserEvent implements ExampleServiceEvent {
  factory _GetUserEvent({@required int userID, Map<String, String> headers}) =
      _$_GetUserEvent;

  int get userID;
  @override
  Map<String, String> get headers;
  @override
  _$GetUserEventCopyWith<_GetUserEvent> get copyWith;
}

abstract class _$UpdateNameEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$UpdateNameEventCopyWith(
          _UpdateNameEvent value, $Res Function(_UpdateNameEvent) then) =
      __$UpdateNameEventCopyWithImpl<$Res>;
  @override
  $Res call({int id, String username, Map<String, String> headers});
}

class __$UpdateNameEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$UpdateNameEventCopyWith<$Res> {
  __$UpdateNameEventCopyWithImpl(
      _UpdateNameEvent _value, $Res Function(_UpdateNameEvent) _then)
      : super(_value, (v) => _then(v as _UpdateNameEvent));

  @override
  _UpdateNameEvent get _value => super._value as _UpdateNameEvent;

  @override
  $Res call({
    Object id = freezed,
    Object username = freezed,
    Object headers = freezed,
  }) {
    return _then(_UpdateNameEvent(
      id: id == freezed ? _value.id : id as int,
      username: username == freezed ? _value.username : username as String,
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_UpdateNameEvent
    with DiagnosticableTreeMixin
    implements _UpdateNameEvent {
  _$_UpdateNameEvent({@required this.id, @required this.username, this.headers})
      : assert(id != null),
        assert(username != null);

  @override
  final int id;
  @override
  final String username;
  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.updateName(id: $id, username: $username, headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.updateName'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('username', username))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _UpdateNameEvent &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.username, username) ||
                const DeepCollectionEquality()
                    .equals(other.username, username)) &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(username) ^
      const DeepCollectionEquality().hash(headers);

  @override
  _$UpdateNameEventCopyWith<_UpdateNameEvent> get copyWith =>
      __$UpdateNameEventCopyWithImpl<_UpdateNameEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return updateName(id, username, headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (updateName != null) {
      return updateName(id, username, headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return updateName(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (updateName != null) {
      return updateName(this);
    }
    return orElse();
  }
}

abstract class _UpdateNameEvent implements ExampleServiceEvent {
  factory _UpdateNameEvent(
      {@required int id,
      @required String username,
      Map<String, String> headers}) = _$_UpdateNameEvent;

  int get id;
  String get username;
  @override
  Map<String, String> get headers;
  @override
  _$UpdateNameEventCopyWith<_UpdateNameEvent> get copyWith;
}

abstract class _$FindUserByIdEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$FindUserByIdEventCopyWith(
          _FindUserByIdEvent value, $Res Function(_FindUserByIdEvent) then) =
      __$FindUserByIdEventCopyWithImpl<$Res>;
  @override
  $Res call({SearchFilter s, Map<String, String> headers});

  $SearchFilterCopyWith<$Res> get s;
}

class __$FindUserByIdEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$FindUserByIdEventCopyWith<$Res> {
  __$FindUserByIdEventCopyWithImpl(
      _FindUserByIdEvent _value, $Res Function(_FindUserByIdEvent) _then)
      : super(_value, (v) => _then(v as _FindUserByIdEvent));

  @override
  _FindUserByIdEvent get _value => super._value as _FindUserByIdEvent;

  @override
  $Res call({
    Object s = freezed,
    Object headers = freezed,
  }) {
    return _then(_FindUserByIdEvent(
      s: s == freezed ? _value.s : s as SearchFilter,
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
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

class _$_FindUserByIdEvent
    with DiagnosticableTreeMixin
    implements _FindUserByIdEvent {
  _$_FindUserByIdEvent({@required this.s, this.headers}) : assert(s != null);

  @override
  final SearchFilter s;
  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.findUserById(s: $s, headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.findUserById'))
      ..add(DiagnosticsProperty('s', s))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _FindUserByIdEvent &&
            (identical(other.s, s) ||
                const DeepCollectionEquality().equals(other.s, s)) &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(s) ^
      const DeepCollectionEquality().hash(headers);

  @override
  _$FindUserByIdEventCopyWith<_FindUserByIdEvent> get copyWith =>
      __$FindUserByIdEventCopyWithImpl<_FindUserByIdEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return findUserById(s, headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (findUserById != null) {
      return findUserById(s, headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return findUserById(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (findUserById != null) {
      return findUserById(this);
    }
    return orElse();
  }
}

abstract class _FindUserByIdEvent implements ExampleServiceEvent {
  factory _FindUserByIdEvent(
      {@required SearchFilter s,
      Map<String, String> headers}) = _$_FindUserByIdEvent;

  SearchFilter get s;
  @override
  Map<String, String> get headers;
  @override
  _$FindUserByIdEventCopyWith<_FindUserByIdEvent> get copyWith;
}

abstract class _$AddUserEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$AddUserEventCopyWith(
          _AddUserEvent value, $Res Function(_AddUserEvent) then) =
      __$AddUserEventCopyWithImpl<$Res>;
  @override
  $Res call({User user, Map<String, String> headers});

  $UserCopyWith<$Res> get user;
}

class __$AddUserEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$AddUserEventCopyWith<$Res> {
  __$AddUserEventCopyWithImpl(
      _AddUserEvent _value, $Res Function(_AddUserEvent) _then)
      : super(_value, (v) => _then(v as _AddUserEvent));

  @override
  _AddUserEvent get _value => super._value as _AddUserEvent;

  @override
  $Res call({
    Object user = freezed,
    Object headers = freezed,
  }) {
    return _then(_AddUserEvent(
      user: user == freezed ? _value.user : user as User,
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
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

class _$_AddUserEvent with DiagnosticableTreeMixin implements _AddUserEvent {
  _$_AddUserEvent({@required this.user, this.headers}) : assert(user != null);

  @override
  final User user;
  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.addUser(user: $user, headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.addUser'))
      ..add(DiagnosticsProperty('user', user))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _AddUserEvent &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)) &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(user) ^
      const DeepCollectionEquality().hash(headers);

  @override
  _$AddUserEventCopyWith<_AddUserEvent> get copyWith =>
      __$AddUserEventCopyWithImpl<_AddUserEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return addUser(user, headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (addUser != null) {
      return addUser(user, headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return addUser(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (addUser != null) {
      return addUser(this);
    }
    return orElse();
  }
}

abstract class _AddUserEvent implements ExampleServiceEvent {
  factory _AddUserEvent({@required User user, Map<String, String> headers}) =
      _$_AddUserEvent;

  User get user;
  @override
  Map<String, String> get headers;
  @override
  _$AddUserEventCopyWith<_AddUserEvent> get copyWith;
}

abstract class _$ListUsersEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$ListUsersEventCopyWith(
          _ListUsersEvent value, $Res Function(_ListUsersEvent) then) =
      __$ListUsersEventCopyWithImpl<$Res>;
  @override
  $Res call({Map<String, String> headers});
}

class __$ListUsersEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$ListUsersEventCopyWith<$Res> {
  __$ListUsersEventCopyWithImpl(
      _ListUsersEvent _value, $Res Function(_ListUsersEvent) _then)
      : super(_value, (v) => _then(v as _ListUsersEvent));

  @override
  _ListUsersEvent get _value => super._value as _ListUsersEvent;

  @override
  $Res call({
    Object headers = freezed,
  }) {
    return _then(_ListUsersEvent(
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_ListUsersEvent
    with DiagnosticableTreeMixin
    implements _ListUsersEvent {
  _$_ListUsersEvent({this.headers});

  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.listUsers(headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.listUsers'))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _ListUsersEvent &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(headers);

  @override
  _$ListUsersEventCopyWith<_ListUsersEvent> get copyWith =>
      __$ListUsersEventCopyWithImpl<_ListUsersEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return listUsers(headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (listUsers != null) {
      return listUsers(headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return listUsers(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (listUsers != null) {
      return listUsers(this);
    }
    return orElse();
  }
}

abstract class _ListUsersEvent implements ExampleServiceEvent {
  factory _ListUsersEvent({Map<String, String> headers}) = _$_ListUsersEvent;

  @override
  Map<String, String> get headers;
  @override
  _$ListUsersEventCopyWith<_ListUsersEvent> get copyWith;
}

abstract class _$DeleteUserEventCopyWith<$Res>
    implements $ExampleServiceEventCopyWith<$Res> {
  factory _$DeleteUserEventCopyWith(
          _DeleteUserEvent value, $Res Function(_DeleteUserEvent) then) =
      __$DeleteUserEventCopyWithImpl<$Res>;
  @override
  $Res call({int id, Map<String, String> headers});
}

class __$DeleteUserEventCopyWithImpl<$Res>
    extends _$ExampleServiceEventCopyWithImpl<$Res>
    implements _$DeleteUserEventCopyWith<$Res> {
  __$DeleteUserEventCopyWithImpl(
      _DeleteUserEvent _value, $Res Function(_DeleteUserEvent) _then)
      : super(_value, (v) => _then(v as _DeleteUserEvent));

  @override
  _DeleteUserEvent get _value => super._value as _DeleteUserEvent;

  @override
  $Res call({
    Object id = freezed,
    Object headers = freezed,
  }) {
    return _then(_DeleteUserEvent(
      id: id == freezed ? _value.id : id as int,
      headers:
          headers == freezed ? _value.headers : headers as Map<String, String>,
    ));
  }
}

class _$_DeleteUserEvent
    with DiagnosticableTreeMixin
    implements _DeleteUserEvent {
  _$_DeleteUserEvent({@required this.id, this.headers}) : assert(id != null);

  @override
  final int id;
  @override
  final Map<String, String> headers;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceEvent.deleteUser(id: $id, headers: $headers)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceEvent.deleteUser'))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('headers', headers));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is _DeleteUserEvent &&
            (identical(other.id, id) ||
                const DeepCollectionEquality().equals(other.id, id)) &&
            (identical(other.headers, headers) ||
                const DeepCollectionEquality().equals(other.headers, headers)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^
      const DeepCollectionEquality().hash(id) ^
      const DeepCollectionEquality().hash(headers);

  @override
  _$DeleteUserEventCopyWith<_DeleteUserEvent> get copyWith =>
      __$DeleteUserEventCopyWithImpl<_DeleteUserEvent>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result ping(Map<String, String> headers),
    @required Result status(Map<String, String> headers),
    @required Result version(Map<String, String> headers),
    @required Result getUser(int userID, Map<String, String> headers),
    @required
        Result updateName(int id, String username, Map<String, String> headers),
    @required Result findUserById(SearchFilter s, Map<String, String> headers),
    @required Result addUser(User user, Map<String, String> headers),
    @required Result listUsers(Map<String, String> headers),
    @required Result deleteUser(int id, Map<String, String> headers),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return deleteUser(id, headers);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result ping(Map<String, String> headers),
    Result status(Map<String, String> headers),
    Result version(Map<String, String> headers),
    Result getUser(int userID, Map<String, String> headers),
    Result updateName(int id, String username, Map<String, String> headers),
    Result findUserById(SearchFilter s, Map<String, String> headers),
    Result addUser(User user, Map<String, String> headers),
    Result listUsers(Map<String, String> headers),
    Result deleteUser(int id, Map<String, String> headers),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (deleteUser != null) {
      return deleteUser(id, headers);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result ping(_PingEvent value),
    @required Result status(_StatusEvent value),
    @required Result version(_VersionEvent value),
    @required Result getUser(_GetUserEvent value),
    @required Result updateName(_UpdateNameEvent value),
    @required Result findUserById(_FindUserByIdEvent value),
    @required Result addUser(_AddUserEvent value),
    @required Result listUsers(_ListUsersEvent value),
    @required Result deleteUser(_DeleteUserEvent value),
  }) {
    assert(ping != null);
    assert(status != null);
    assert(version != null);
    assert(getUser != null);
    assert(updateName != null);
    assert(findUserById != null);
    assert(addUser != null);
    assert(listUsers != null);
    assert(deleteUser != null);
    return deleteUser(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result ping(_PingEvent value),
    Result status(_StatusEvent value),
    Result version(_VersionEvent value),
    Result getUser(_GetUserEvent value),
    Result updateName(_UpdateNameEvent value),
    Result findUserById(_FindUserByIdEvent value),
    Result addUser(_AddUserEvent value),
    Result listUsers(_ListUsersEvent value),
    Result deleteUser(_DeleteUserEvent value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (deleteUser != null) {
      return deleteUser(this);
    }
    return orElse();
  }
}

abstract class _DeleteUserEvent implements ExampleServiceEvent {
  factory _DeleteUserEvent({@required int id, Map<String, String> headers}) =
      _$_DeleteUserEvent;

  int get id;
  @override
  Map<String, String> get headers;
  @override
  _$DeleteUserEventCopyWith<_DeleteUserEvent> get copyWith;
}

ExampleServiceState _$ExampleServiceStateFromJson(Map<String, dynamic> json) {
  switch (json['runtimeType'] as String) {
    case 'pingResult':
      return PingResult.fromJson(json);
    case 'statusResult':
      return StatusResult.fromJson(json);
    case 'versionResult':
      return VersionResult.fromJson(json);
    case 'getUserResult':
      return GetUserResult.fromJson(json);
    case 'updateNameResult':
      return UpdateNameResult.fromJson(json);
    case 'findUserByIdResult':
      return FindUserByIdResult.fromJson(json);
    case 'addUserResult':
      return AddUserResult.fromJson(json);
    case 'listUsersResult':
      return ListUsersResult.fromJson(json);
    case 'deleteUserResult':
      return DeleteUserResult.fromJson(json);

    default:
      throw FallThroughError();
  }
}

class _$ExampleServiceStateTearOff {
  const _$ExampleServiceStateTearOff();

  PingResult pingResult() {
    return PingResult();
  }

  StatusResult statusResult({@required bool status}) {
    return StatusResult(
      status: status,
    );
  }

  VersionResult versionResult({@required Version version}) {
    return VersionResult(
      version: version,
    );
  }

  GetUserResult getUserResult({@required User user}) {
    return GetUserResult(
      user: user,
    );
  }

  UpdateNameResult updateNameResult({@required bool wasUpdated}) {
    return UpdateNameResult(
      wasUpdated: wasUpdated,
    );
  }

  FindUserByIdResult findUserByIdResult(
      {@required String name, @required User user}) {
    return FindUserByIdResult(
      name: name,
      user: user,
    );
  }

  AddUserResult addUserResult({@required bool wasAdded}) {
    return AddUserResult(
      wasAdded: wasAdded,
    );
  }

  ListUsersResult listUsersResult({@required List<User> users}) {
    return ListUsersResult(
      users: users,
    );
  }

  DeleteUserResult deleteUserResult({@required bool wasDeleted}) {
    return DeleteUserResult(
      wasDeleted: wasDeleted,
    );
  }
}

// ignore: unused_element
const $ExampleServiceState = _$ExampleServiceStateTearOff();

mixin _$ExampleServiceState {
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  });
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  });
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  });
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  });
  Map<String, dynamic> toJson();
}

abstract class $ExampleServiceStateCopyWith<$Res> {
  factory $ExampleServiceStateCopyWith(
          ExampleServiceState value, $Res Function(ExampleServiceState) then) =
      _$ExampleServiceStateCopyWithImpl<$Res>;
}

class _$ExampleServiceStateCopyWithImpl<$Res>
    implements $ExampleServiceStateCopyWith<$Res> {
  _$ExampleServiceStateCopyWithImpl(this._value, this._then);

  final ExampleServiceState _value;
  // ignore: unused_field
  final $Res Function(ExampleServiceState) _then;
}

abstract class $PingResultCopyWith<$Res> {
  factory $PingResultCopyWith(
          PingResult value, $Res Function(PingResult) then) =
      _$PingResultCopyWithImpl<$Res>;
}

class _$PingResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $PingResultCopyWith<$Res> {
  _$PingResultCopyWithImpl(PingResult _value, $Res Function(PingResult) _then)
      : super(_value, (v) => _then(v as PingResult));

  @override
  PingResult get _value => super._value as PingResult;
}

@JsonSerializable()
class _$PingResult with DiagnosticableTreeMixin implements PingResult {
  _$PingResult();

  factory _$PingResult.fromJson(Map<String, dynamic> json) =>
      _$_$PingResultFromJson(json);

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.pingResult()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.pingResult'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) || (other is PingResult);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return pingResult();
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pingResult != null) {
      return pingResult();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return pingResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (pingResult != null) {
      return pingResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$PingResultToJson(this)..['runtimeType'] = 'pingResult';
  }
}

abstract class PingResult implements ExampleServiceState {
  factory PingResult() = _$PingResult;

  factory PingResult.fromJson(Map<String, dynamic> json) =
      _$PingResult.fromJson;
}

abstract class $StatusResultCopyWith<$Res> {
  factory $StatusResultCopyWith(
          StatusResult value, $Res Function(StatusResult) then) =
      _$StatusResultCopyWithImpl<$Res>;
  $Res call({bool status});
}

class _$StatusResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $StatusResultCopyWith<$Res> {
  _$StatusResultCopyWithImpl(
      StatusResult _value, $Res Function(StatusResult) _then)
      : super(_value, (v) => _then(v as StatusResult));

  @override
  StatusResult get _value => super._value as StatusResult;

  @override
  $Res call({
    Object status = freezed,
  }) {
    return _then(StatusResult(
      status: status == freezed ? _value.status : status as bool,
    ));
  }
}

@JsonSerializable()
class _$StatusResult with DiagnosticableTreeMixin implements StatusResult {
  _$StatusResult({@required this.status}) : assert(status != null);

  factory _$StatusResult.fromJson(Map<String, dynamic> json) =>
      _$_$StatusResultFromJson(json);

  @override
  final bool status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.statusResult(status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.statusResult'))
      ..add(DiagnosticsProperty('status', status));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is StatusResult &&
            (identical(other.status, status) ||
                const DeepCollectionEquality().equals(other.status, status)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(status);

  @override
  $StatusResultCopyWith<StatusResult> get copyWith =>
      _$StatusResultCopyWithImpl<StatusResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return statusResult(status);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (statusResult != null) {
      return statusResult(status);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return statusResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (statusResult != null) {
      return statusResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$StatusResultToJson(this)..['runtimeType'] = 'statusResult';
  }
}

abstract class StatusResult implements ExampleServiceState {
  factory StatusResult({@required bool status}) = _$StatusResult;

  factory StatusResult.fromJson(Map<String, dynamic> json) =
      _$StatusResult.fromJson;

  bool get status;
  $StatusResultCopyWith<StatusResult> get copyWith;
}

abstract class $VersionResultCopyWith<$Res> {
  factory $VersionResultCopyWith(
          VersionResult value, $Res Function(VersionResult) then) =
      _$VersionResultCopyWithImpl<$Res>;
  $Res call({Version version});

  $VersionCopyWith<$Res> get version;
}

class _$VersionResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $VersionResultCopyWith<$Res> {
  _$VersionResultCopyWithImpl(
      VersionResult _value, $Res Function(VersionResult) _then)
      : super(_value, (v) => _then(v as VersionResult));

  @override
  VersionResult get _value => super._value as VersionResult;

  @override
  $Res call({
    Object version = freezed,
  }) {
    return _then(VersionResult(
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
class _$VersionResult with DiagnosticableTreeMixin implements VersionResult {
  _$VersionResult({@required this.version}) : assert(version != null);

  factory _$VersionResult.fromJson(Map<String, dynamic> json) =>
      _$_$VersionResultFromJson(json);

  @override
  final Version version;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.versionResult(version: $version)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.versionResult'))
      ..add(DiagnosticsProperty('version', version));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is VersionResult &&
            (identical(other.version, version) ||
                const DeepCollectionEquality().equals(other.version, version)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(version);

  @override
  $VersionResultCopyWith<VersionResult> get copyWith =>
      _$VersionResultCopyWithImpl<VersionResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return versionResult(version);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (versionResult != null) {
      return versionResult(version);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return versionResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (versionResult != null) {
      return versionResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$VersionResultToJson(this)..['runtimeType'] = 'versionResult';
  }
}

abstract class VersionResult implements ExampleServiceState {
  factory VersionResult({@required Version version}) = _$VersionResult;

  factory VersionResult.fromJson(Map<String, dynamic> json) =
      _$VersionResult.fromJson;

  Version get version;
  $VersionResultCopyWith<VersionResult> get copyWith;
}

abstract class $GetUserResultCopyWith<$Res> {
  factory $GetUserResultCopyWith(
          GetUserResult value, $Res Function(GetUserResult) then) =
      _$GetUserResultCopyWithImpl<$Res>;
  $Res call({User user});

  $UserCopyWith<$Res> get user;
}

class _$GetUserResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $GetUserResultCopyWith<$Res> {
  _$GetUserResultCopyWithImpl(
      GetUserResult _value, $Res Function(GetUserResult) _then)
      : super(_value, (v) => _then(v as GetUserResult));

  @override
  GetUserResult get _value => super._value as GetUserResult;

  @override
  $Res call({
    Object user = freezed,
  }) {
    return _then(GetUserResult(
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
class _$GetUserResult with DiagnosticableTreeMixin implements GetUserResult {
  _$GetUserResult({@required this.user}) : assert(user != null);

  factory _$GetUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$GetUserResultFromJson(json);

  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.getUserResult(user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.getUserResult'))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is GetUserResult &&
            (identical(other.user, user) ||
                const DeepCollectionEquality().equals(other.user, user)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(user);

  @override
  $GetUserResultCopyWith<GetUserResult> get copyWith =>
      _$GetUserResultCopyWithImpl<GetUserResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return getUserResult(user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (getUserResult != null) {
      return getUserResult(user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return getUserResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (getUserResult != null) {
      return getUserResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$GetUserResultToJson(this)..['runtimeType'] = 'getUserResult';
  }
}

abstract class GetUserResult implements ExampleServiceState {
  factory GetUserResult({@required User user}) = _$GetUserResult;

  factory GetUserResult.fromJson(Map<String, dynamic> json) =
      _$GetUserResult.fromJson;

  User get user;
  $GetUserResultCopyWith<GetUserResult> get copyWith;
}

abstract class $UpdateNameResultCopyWith<$Res> {
  factory $UpdateNameResultCopyWith(
          UpdateNameResult value, $Res Function(UpdateNameResult) then) =
      _$UpdateNameResultCopyWithImpl<$Res>;
  $Res call({bool wasUpdated});
}

class _$UpdateNameResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $UpdateNameResultCopyWith<$Res> {
  _$UpdateNameResultCopyWithImpl(
      UpdateNameResult _value, $Res Function(UpdateNameResult) _then)
      : super(_value, (v) => _then(v as UpdateNameResult));

  @override
  UpdateNameResult get _value => super._value as UpdateNameResult;

  @override
  $Res call({
    Object wasUpdated = freezed,
  }) {
    return _then(UpdateNameResult(
      wasUpdated:
          wasUpdated == freezed ? _value.wasUpdated : wasUpdated as bool,
    ));
  }
}

@JsonSerializable()
class _$UpdateNameResult
    with DiagnosticableTreeMixin
    implements UpdateNameResult {
  _$UpdateNameResult({@required this.wasUpdated}) : assert(wasUpdated != null);

  factory _$UpdateNameResult.fromJson(Map<String, dynamic> json) =>
      _$_$UpdateNameResultFromJson(json);

  @override
  final bool wasUpdated;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.updateNameResult(wasUpdated: $wasUpdated)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.updateNameResult'))
      ..add(DiagnosticsProperty('wasUpdated', wasUpdated));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is UpdateNameResult &&
            (identical(other.wasUpdated, wasUpdated) ||
                const DeepCollectionEquality()
                    .equals(other.wasUpdated, wasUpdated)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasUpdated);

  @override
  $UpdateNameResultCopyWith<UpdateNameResult> get copyWith =>
      _$UpdateNameResultCopyWithImpl<UpdateNameResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return updateNameResult(wasUpdated);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (updateNameResult != null) {
      return updateNameResult(wasUpdated);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return updateNameResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (updateNameResult != null) {
      return updateNameResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$UpdateNameResultToJson(this)
      ..['runtimeType'] = 'updateNameResult';
  }
}

abstract class UpdateNameResult implements ExampleServiceState {
  factory UpdateNameResult({@required bool wasUpdated}) = _$UpdateNameResult;

  factory UpdateNameResult.fromJson(Map<String, dynamic> json) =
      _$UpdateNameResult.fromJson;

  bool get wasUpdated;
  $UpdateNameResultCopyWith<UpdateNameResult> get copyWith;
}

abstract class $FindUserByIdResultCopyWith<$Res> {
  factory $FindUserByIdResultCopyWith(
          FindUserByIdResult value, $Res Function(FindUserByIdResult) then) =
      _$FindUserByIdResultCopyWithImpl<$Res>;
  $Res call({String name, User user});

  $UserCopyWith<$Res> get user;
}

class _$FindUserByIdResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $FindUserByIdResultCopyWith<$Res> {
  _$FindUserByIdResultCopyWithImpl(
      FindUserByIdResult _value, $Res Function(FindUserByIdResult) _then)
      : super(_value, (v) => _then(v as FindUserByIdResult));

  @override
  FindUserByIdResult get _value => super._value as FindUserByIdResult;

  @override
  $Res call({
    Object name = freezed,
    Object user = freezed,
  }) {
    return _then(FindUserByIdResult(
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
class _$FindUserByIdResult
    with DiagnosticableTreeMixin
    implements FindUserByIdResult {
  _$FindUserByIdResult({@required this.name, @required this.user})
      : assert(name != null),
        assert(user != null);

  factory _$FindUserByIdResult.fromJson(Map<String, dynamic> json) =>
      _$_$FindUserByIdResultFromJson(json);

  @override
  final String name;
  @override
  final User user;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.findUserByIdResult(name: $name, user: $user)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(
          DiagnosticsProperty('type', 'ExampleServiceState.findUserByIdResult'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('user', user));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is FindUserByIdResult &&
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
  $FindUserByIdResultCopyWith<FindUserByIdResult> get copyWith =>
      _$FindUserByIdResultCopyWithImpl<FindUserByIdResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return findUserByIdResult(name, user);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (findUserByIdResult != null) {
      return findUserByIdResult(name, user);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return findUserByIdResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (findUserByIdResult != null) {
      return findUserByIdResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$FindUserByIdResultToJson(this)
      ..['runtimeType'] = 'findUserByIdResult';
  }
}

abstract class FindUserByIdResult implements ExampleServiceState {
  factory FindUserByIdResult({@required String name, @required User user}) =
      _$FindUserByIdResult;

  factory FindUserByIdResult.fromJson(Map<String, dynamic> json) =
      _$FindUserByIdResult.fromJson;

  String get name;
  User get user;
  $FindUserByIdResultCopyWith<FindUserByIdResult> get copyWith;
}

abstract class $AddUserResultCopyWith<$Res> {
  factory $AddUserResultCopyWith(
          AddUserResult value, $Res Function(AddUserResult) then) =
      _$AddUserResultCopyWithImpl<$Res>;
  $Res call({bool wasAdded});
}

class _$AddUserResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $AddUserResultCopyWith<$Res> {
  _$AddUserResultCopyWithImpl(
      AddUserResult _value, $Res Function(AddUserResult) _then)
      : super(_value, (v) => _then(v as AddUserResult));

  @override
  AddUserResult get _value => super._value as AddUserResult;

  @override
  $Res call({
    Object wasAdded = freezed,
  }) {
    return _then(AddUserResult(
      wasAdded: wasAdded == freezed ? _value.wasAdded : wasAdded as bool,
    ));
  }
}

@JsonSerializable()
class _$AddUserResult with DiagnosticableTreeMixin implements AddUserResult {
  _$AddUserResult({@required this.wasAdded}) : assert(wasAdded != null);

  factory _$AddUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$AddUserResultFromJson(json);

  @override
  final bool wasAdded;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.addUserResult(wasAdded: $wasAdded)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.addUserResult'))
      ..add(DiagnosticsProperty('wasAdded', wasAdded));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is AddUserResult &&
            (identical(other.wasAdded, wasAdded) ||
                const DeepCollectionEquality()
                    .equals(other.wasAdded, wasAdded)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasAdded);

  @override
  $AddUserResultCopyWith<AddUserResult> get copyWith =>
      _$AddUserResultCopyWithImpl<AddUserResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return addUserResult(wasAdded);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (addUserResult != null) {
      return addUserResult(wasAdded);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return addUserResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (addUserResult != null) {
      return addUserResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$AddUserResultToJson(this)..['runtimeType'] = 'addUserResult';
  }
}

abstract class AddUserResult implements ExampleServiceState {
  factory AddUserResult({@required bool wasAdded}) = _$AddUserResult;

  factory AddUserResult.fromJson(Map<String, dynamic> json) =
      _$AddUserResult.fromJson;

  bool get wasAdded;
  $AddUserResultCopyWith<AddUserResult> get copyWith;
}

abstract class $ListUsersResultCopyWith<$Res> {
  factory $ListUsersResultCopyWith(
          ListUsersResult value, $Res Function(ListUsersResult) then) =
      _$ListUsersResultCopyWithImpl<$Res>;
  $Res call({List<User> users});
}

class _$ListUsersResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $ListUsersResultCopyWith<$Res> {
  _$ListUsersResultCopyWithImpl(
      ListUsersResult _value, $Res Function(ListUsersResult) _then)
      : super(_value, (v) => _then(v as ListUsersResult));

  @override
  ListUsersResult get _value => super._value as ListUsersResult;

  @override
  $Res call({
    Object users = freezed,
  }) {
    return _then(ListUsersResult(
      users: users == freezed ? _value.users : users as List<User>,
    ));
  }
}

@JsonSerializable()
class _$ListUsersResult
    with DiagnosticableTreeMixin
    implements ListUsersResult {
  _$ListUsersResult({@required this.users}) : assert(users != null);

  factory _$ListUsersResult.fromJson(Map<String, dynamic> json) =>
      _$_$ListUsersResultFromJson(json);

  @override
  final List<User> users;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.listUsersResult(users: $users)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.listUsersResult'))
      ..add(DiagnosticsProperty('users', users));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is ListUsersResult &&
            (identical(other.users, users) ||
                const DeepCollectionEquality().equals(other.users, users)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(users);

  @override
  $ListUsersResultCopyWith<ListUsersResult> get copyWith =>
      _$ListUsersResultCopyWithImpl<ListUsersResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return listUsersResult(users);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (listUsersResult != null) {
      return listUsersResult(users);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return listUsersResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (listUsersResult != null) {
      return listUsersResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$ListUsersResultToJson(this)..['runtimeType'] = 'listUsersResult';
  }
}

abstract class ListUsersResult implements ExampleServiceState {
  factory ListUsersResult({@required List<User> users}) = _$ListUsersResult;

  factory ListUsersResult.fromJson(Map<String, dynamic> json) =
      _$ListUsersResult.fromJson;

  List<User> get users;
  $ListUsersResultCopyWith<ListUsersResult> get copyWith;
}

abstract class $DeleteUserResultCopyWith<$Res> {
  factory $DeleteUserResultCopyWith(
          DeleteUserResult value, $Res Function(DeleteUserResult) then) =
      _$DeleteUserResultCopyWithImpl<$Res>;
  $Res call({bool wasDeleted});
}

class _$DeleteUserResultCopyWithImpl<$Res>
    extends _$ExampleServiceStateCopyWithImpl<$Res>
    implements $DeleteUserResultCopyWith<$Res> {
  _$DeleteUserResultCopyWithImpl(
      DeleteUserResult _value, $Res Function(DeleteUserResult) _then)
      : super(_value, (v) => _then(v as DeleteUserResult));

  @override
  DeleteUserResult get _value => super._value as DeleteUserResult;

  @override
  $Res call({
    Object wasDeleted = freezed,
  }) {
    return _then(DeleteUserResult(
      wasDeleted:
          wasDeleted == freezed ? _value.wasDeleted : wasDeleted as bool,
    ));
  }
}

@JsonSerializable()
class _$DeleteUserResult
    with DiagnosticableTreeMixin
    implements DeleteUserResult {
  _$DeleteUserResult({@required this.wasDeleted}) : assert(wasDeleted != null);

  factory _$DeleteUserResult.fromJson(Map<String, dynamic> json) =>
      _$_$DeleteUserResultFromJson(json);

  @override
  final bool wasDeleted;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'ExampleServiceState.deleteUserResult(wasDeleted: $wasDeleted)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'ExampleServiceState.deleteUserResult'))
      ..add(DiagnosticsProperty('wasDeleted', wasDeleted));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other is DeleteUserResult &&
            (identical(other.wasDeleted, wasDeleted) ||
                const DeepCollectionEquality()
                    .equals(other.wasDeleted, wasDeleted)));
  }

  @override
  int get hashCode =>
      runtimeType.hashCode ^ const DeepCollectionEquality().hash(wasDeleted);

  @override
  $DeleteUserResultCopyWith<DeleteUserResult> get copyWith =>
      _$DeleteUserResultCopyWithImpl<DeleteUserResult>(this, _$identity);

  @override
  @optionalTypeArgs
  Result when<Result extends Object>({
    @required Result pingResult(),
    @required Result statusResult(bool status),
    @required Result versionResult(Version version),
    @required Result getUserResult(User user),
    @required Result updateNameResult(bool wasUpdated),
    @required Result findUserByIdResult(String name, User user),
    @required Result addUserResult(bool wasAdded),
    @required Result listUsersResult(List<User> users),
    @required Result deleteUserResult(bool wasDeleted),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return deleteUserResult(wasDeleted);
  }

  @override
  @optionalTypeArgs
  Result maybeWhen<Result extends Object>({
    Result pingResult(),
    Result statusResult(bool status),
    Result versionResult(Version version),
    Result getUserResult(User user),
    Result updateNameResult(bool wasUpdated),
    Result findUserByIdResult(String name, User user),
    Result addUserResult(bool wasAdded),
    Result listUsersResult(List<User> users),
    Result deleteUserResult(bool wasDeleted),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (deleteUserResult != null) {
      return deleteUserResult(wasDeleted);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  Result map<Result extends Object>({
    @required Result pingResult(PingResult value),
    @required Result statusResult(StatusResult value),
    @required Result versionResult(VersionResult value),
    @required Result getUserResult(GetUserResult value),
    @required Result updateNameResult(UpdateNameResult value),
    @required Result findUserByIdResult(FindUserByIdResult value),
    @required Result addUserResult(AddUserResult value),
    @required Result listUsersResult(ListUsersResult value),
    @required Result deleteUserResult(DeleteUserResult value),
  }) {
    assert(pingResult != null);
    assert(statusResult != null);
    assert(versionResult != null);
    assert(getUserResult != null);
    assert(updateNameResult != null);
    assert(findUserByIdResult != null);
    assert(addUserResult != null);
    assert(listUsersResult != null);
    assert(deleteUserResult != null);
    return deleteUserResult(this);
  }

  @override
  @optionalTypeArgs
  Result maybeMap<Result extends Object>({
    Result pingResult(PingResult value),
    Result statusResult(StatusResult value),
    Result versionResult(VersionResult value),
    Result getUserResult(GetUserResult value),
    Result updateNameResult(UpdateNameResult value),
    Result findUserByIdResult(FindUserByIdResult value),
    Result addUserResult(AddUserResult value),
    Result listUsersResult(ListUsersResult value),
    Result deleteUserResult(DeleteUserResult value),
    @required Result orElse(),
  }) {
    assert(orElse != null);
    if (deleteUserResult != null) {
      return deleteUserResult(this);
    }
    return orElse();
  }

  @override
  Map<String, dynamic> toJson() {
    return _$_$DeleteUserResultToJson(this)
      ..['runtimeType'] = 'deleteUserResult';
  }
}

abstract class DeleteUserResult implements ExampleServiceState {
  factory DeleteUserResult({@required bool wasDeleted}) = _$DeleteUserResult;

  factory DeleteUserResult.fromJson(Map<String, dynamic> json) =
      _$DeleteUserResult.fromJson;

  bool get wasDeleted;
  $DeleteUserResultCopyWith<DeleteUserResult> get copyWith;
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

class _$_RpcStateIdle<T>
    with DiagnosticableTreeMixin
    implements _RpcStateIdle<T> {
  const _$_RpcStateIdle();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcState<$T>.idle()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'RpcState<$T>.idle'));
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

class _$_RpcStateLoading<T>
    with DiagnosticableTreeMixin
    implements _RpcStateLoading<T> {
  const _$_RpcStateLoading();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcState<$T>.loading()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'RpcState<$T>.loading'));
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

class _$_RpcStateOk<T> with DiagnosticableTreeMixin implements _RpcStateOk<T> {
  const _$_RpcStateOk({@required this.data}) : assert(data != null);

  @override
  final T data;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcState<$T>.ok(data: $data)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'RpcState<$T>.ok'))
      ..add(DiagnosticsProperty('data', data));
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

class _$_RpcStateUnit<T>
    with DiagnosticableTreeMixin
    implements _RpcStateUnit<T> {
  const _$_RpcStateUnit();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcState<$T>.unit()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'RpcState<$T>.unit'));
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

class _$_RpcStateErr<T>
    with DiagnosticableTreeMixin
    implements _RpcStateErr<T> {
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
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'RpcState<$T>.err(reason: $reason, statusCode: $statusCode, stackTrace: $stackTrace)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'RpcState<$T>.err'))
      ..add(DiagnosticsProperty('reason', reason))
      ..add(DiagnosticsProperty('statusCode', statusCode))
      ..add(DiagnosticsProperty('stackTrace', stackTrace));
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
