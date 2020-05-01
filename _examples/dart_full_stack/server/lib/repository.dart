import 'package:meta/meta.dart';

import './services/example_service.dart';

class UserRepository {
  final Set<User> users = {};

  bool addUser({@required User user}) => users.add(user);
  bool deleteUser({@required int id}) => users.remove(
        findUserById(id: id),
      );
  User findUserById({@required int id}) => users.firstWhere(
        (user) => user.id == id,
      );
  List<User> listUsers() => users.toList();
  bool updateUser({
    @required int id,
    @required String name,
  }) {
    final user = findUserById(id: id);
    final newUser = user.copyWith(username: name);
    if (!deleteUser(id: id)) {
      return false;
    }
    addUser(user: newUser);
    return true;
  }
}
