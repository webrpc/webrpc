import 'dart:convert';

import 'package:collection/collection.dart';

class MyType {
  final String name;
  final int age;
  final Map<String, int> data;
  final List<int> list;
  final bool hasMore;
  MyType({
    this.name,
    this.age,
    this.data,
    this.list,
    this.hasMore,
  });

  MyType copyWith({
    String name,
    int age,
    Map<String, int> data,
    List<int> list,
    bool hasMore,
  }) {
    return MyType(
      name: name ?? this.name,
      age: age ?? this.age,
      data: data ?? this.data,
      list: list ?? this.list,
      hasMore: hasMore ?? this.hasMore,
    );
  }

  Map<String, dynamic> toMap() {
    return {
      'name': name,
      'age': age,
      'data': data.toMap(),
      'list': List<dynamic>.from(list.map((x) => x)),
      'hasMore': hasMore,
    };
  }

  static MyType fromMap(Map<String, dynamic> map) {
    if (map == null) return null;

    return MyType(
      name: map['name'],
      age: map['age'],
      data: Map<String, int>.fromMap(map['data']),
      list: List<int>.from(map['list']),
      hasMore: map['hasMore'],
    );
  }

  String toJson() => json.encode(toMap());

  static MyType fromJson(String source) => fromMap(json.decode(source));

  @override
  String toString() {
    return 'MyType(name: $name, age: $age, data: $data, list: $list, hasMore: $hasMore)';
  }

  @override
  bool operator ==(Object o) {
    if (identical(this, o)) return true;
    final collectionEquals = const DeepCollectionEquality().equals;

    return o is MyType &&
        o.name == name &&
        o.age == age &&
        collectionEquals(o.data, data) &&
        collectionEquals(o.list, list) &&
        o.hasMore == hasMore;
  }

  @override
  int get hashCode {
    return name.hashCode ^
        age.hashCode ^
        data.hashCode ^
        list.hashCode ^
        hasMore.hashCode;
  }
}
