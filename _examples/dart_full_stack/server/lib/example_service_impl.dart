import 'dart:async';

import 'package:server/services/example_service.dart';

class ExampleServiceImpl implements ExampleService {
  @override
  FutureOr<AddUserResult> addUser({User user}) {
    // TODO: implement addUser
    return null;
  }

  @override
  FutureOr<DeleteUserResult> deleteUser({int id}) {
    // TODO: implement deleteUser
    return null;
  }

  @override
  FutureOr<FindUserByIdResult> findUserById({SearchFilter s}) {
    // TODO: implement findUserById
    return null;
  }

  @override
  FutureOr<GetUserResult> getUser({int userID}) {
    // TODO: implement getUser
    return null;
  }

  @override
  FutureOr<ListUsersResult> listUsers() {
    // TODO: implement listUsers
    return null;
  }

  @override
  FutureOr<void> ping() {
    // TODO: implement ping
    return null;
  }

  @override
  FutureOr<StatusResult> status() {
    // TODO: implement status
    return null;
  }

  @override
  FutureOr<UpdateNameResult> updateName({int id, String username}) {
    // TODO: implement updateName
    return null;
  }

  @override
  FutureOr<VersionResult> version() => VersionResult(
        version: Version(
          webrpcVersion: webRPCVersion(),
          schemaVersion: WebRPCSchemaVersion(),
          schemaHash: WebRPCSchemaHash(),
        ),
      );
}
