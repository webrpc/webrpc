import 'dart:html';
import 'dart:math';

import 'package:client/client.gen.dart';

void main() {
  final pingBtn = querySelector('#ping');
  final getUserBtn = querySelector('#getUser');
  final pingResultDiv = querySelector('#pingResult');
  final getUserResultDiv = querySelector('#getUserResult');
  final host = 'http://localhost:8080';
  int userId() {
    final min = 1;
    final max = 5;
    final rnd = Random();
    return min + rnd.nextInt(max - min);
  }

  final client = ExampleServiceRpc(host: host);
  pingBtn.onClick.listen((event) async {
    await for (final result in client.ping()) {
      result.maybeWhen(
          ok: (data) => pingResultDiv.innerHtml = data.status.toString(),
          orElse: () => pingResultDiv.innerHtml = 'server error occurred');
    }
  });

  getUserBtn.onClick.listen((event) async {
    final id = userId();
    await for (final result in client.getUser(
      userID: id,
    )) {
      result.maybeWhen(
        ok: (data) =>
            getUserResultDiv.innerHtml = '${data.user.username} received',
        orElse: () => getUserResultDiv.innerHtml = 'no user found for id: $id',
      );
    }
  });
}
