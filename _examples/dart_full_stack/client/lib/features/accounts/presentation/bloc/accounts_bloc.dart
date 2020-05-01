import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
part 'accounts_event.dart';
part 'accounts_state.dart';
class AccountsBloc extends Bloc<AccountsEvent, AccountsState> {
  @override
  AccountsState get initialState => AccountsInitial();
  @override
  Stream<AccountsState> mapEventToState(
    AccountsEvent event,
  ) async* {
    // TODO: implement mapEventToState
  }
}
