import 'package:shelf/shelf.dart' as shelf;

abstract class WebRpcAppServer {
  Map<String, shelf.Handler> get handlersMap;
}
