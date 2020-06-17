import 'dart:async';

import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:server/hello_api.gen.dart';

Future<void> main(List<String> args) async {
  final server = WebRpcServer(
    exampleService: ExampleServiceImpl(
      userRepo: UserRepository(),
    ),
  );

  await server.serve(args);
}

class ExampleServiceImpl implements ExampleService {
  final UserRepository userRepo;

  @override
  FutureOr<FindUsersResult> findUsers({String q}) => FindUsersResult(
        page: null,
        users: userRepo.findUsers(q),
      );

  @override
  FutureOr<GetUserResult> getUser({int userID}) => GetUserResult(
        user: userRepo.getUser(userID),
      );

  @override
  FutureOr<PingResult> ping() => PingResult(status: true);

  ExampleServiceImpl({@required this.userRepo});
}

List<User> initUsers() {
  final ids = [1, 2, 3, 4];
  return ids.map(
    (id) => User(
      ID: id,
      username: 'user$id',
      role: const Kind.user(),
      meta: {},
      internalID: id,
      createdAt: DateTime.now(),
    ),
  );
}

class UserRepository {
  final Set<User> users = Set.from(
    initUsers(),
  );

  List<User> findUsers(String query) => users.where(
        (user) => user.username.contains(query),
      );

  User getUser(int id) => users.firstWhere((user) => user.ID == id);
}
