import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
part 'users_event.dart';
part 'users_state.dart';

class UsersBloc extends Bloc<UsersEvent, UsersState> {
  @override
  UsersState get initialState => UsersInitial();
  @override
  Stream<UsersState> mapEventToState(
    UsersEvent event,
  ) async* {
    // TODO: implement mapEventToState
  }
}
