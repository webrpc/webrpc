import 'package:server/example_service_impl.dart';
import 'package:server/repository.dart';
import 'package:server/services/example_service.dart';

Future<void> main(List<String> args) async {
  final server = WebRpcServer(
    exampleService: ExampleServiceImpl(
      userRepo: UserRepository(),
    ),
  );

  await server.serve(args);
}
