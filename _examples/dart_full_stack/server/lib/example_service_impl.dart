import 'dart:async';

import 'package:meta/meta.dart';
import 'package:server/repository.dart';
import 'package:server/services/example_service.dart';

class ExampleServiceImpl implements ExampleService {
  final UserRepository _userRepo;

  ExampleServiceImpl({@required UserRepository userRepo})
      : _userRepo = userRepo;
  @override
  FutureOr<AddUserResult> addUser({@required User user}) => AddUserResult(
        wasAdded: _userRepo.addUser(user: user),
      );

  @override
  FutureOr<DeleteUserResult> deleteUser({int id}) => DeleteUserResult(
        wasDeleted: _userRepo.deleteUser(id: id),
      );

  @override
  FutureOr<FindUserByIdResult> findUserById({SearchFilter s}) {
    final user = _userRepo.findUserById(id: s.id);
    return FindUserByIdResult(
      name: user?.username,
      user: user,
    );
  }

  @override
  FutureOr<GetUserResult> getUser({int userID}) => GetUserResult(
        user: _userRepo.findUserById(id: userID),
      );

  @override
  FutureOr<ListUsersResult> listUsers() => ListUsersResult(
        users: _userRepo.listUsers(),
      );

  @override
  FutureOr<void> ping() => null;

  @override
  FutureOr<StatusResult> status() => StatusResult(status: true);

  @override
  FutureOr<UpdateNameResult> updateName({int id, String username}) =>
      UpdateNameResult(
        wasUpdated: _userRepo.updateUser(id: id, name: username),
      );

  @override
  FutureOr<VersionResult> version() => VersionResult(
        version: Version(
          webrpcVersion: webRPCVersion(),
          schemaVersion: WebRPCSchemaVersion(),
          schemaHash: WebRPCSchemaHash(),
        ),
      );
}
