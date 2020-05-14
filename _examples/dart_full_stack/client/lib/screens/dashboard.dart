import 'package:client/services/example/example_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class DashBoard extends StatefulWidget {
  DashBoard({Key key}) : super(key: key);

  @override
  _DashBoardState createState() => _DashBoardState();
}

class _DashBoardState extends State<DashBoard> {
  final _userFormKey = GlobalKey<FormState>();
  final _findUserFormKey = GlobalKey<FormState>();
  int _userToFindID;
  String _username = '';
  int _userID = 0;
  int get userID => _userID = _userID + 1;

  final ExampleServiceBloc bloc = ExampleServiceBloc(
    exampleService: ExampleServiceRpc(
      host: 'http://localhost:8080',
    ),
  );

  @override
  void dispose() {
    bloc.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    bloc.listUsers();
    return Column(
      children: <Widget>[
        _pingMethodRow(),
        _versionMethodRow(),
        _statusMethodRow(),
        _addUserRow(),
        _findUserRow(),
        _usersTable(),
      ],
    );
  }

  Widget _simpleRow({
    IconData buttonIcon,
    String buttonText,
    void Function() onPressed,
    bool Function(RpcState<ExampleServiceState>) blocRebuildCondition,
    String Function(ExampleServiceState) onOk,
    String unitText,
    String idleText,
  }) =>
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: <Widget>[
          Container(
            alignment: Alignment.topLeft,
            child: RaisedButton(
              onPressed: onPressed,
              child: Row(
                children: <Widget>[
                  Icon(
                    buttonIcon,
                  ),
                  Text(
                    buttonText,
                  ),
                ],
              ),
              color: Colors.blue,
            ),
          ),
          BlocBuilder<ExampleServiceBloc, RpcState<ExampleServiceState>>(
              bloc: bloc,
              condition: (previous, current) => blocRebuildCondition(current),
              builder: (context, state) {
                return Container(
                  width: 150.0,
                  alignment: Alignment.topRight,
                  child: state.when(
                    err: (reason, status, _) => Text(
                      'An error occured: $reason status: $status',
                      textAlign: TextAlign.right,
                    ),
                    idle: () => Text(
                      idleText,
                      textAlign: TextAlign.right,
                    ),
                    loading: () => const CircularProgressIndicator(),
                    ok: (ExampleServiceState data) => Text(
                      onOk(data),
                      textAlign: TextAlign.right,
                    ),
                    unit: () => Text(
                      unitText,
                      textAlign: TextAlign.right,
                    ),
                  ),
                );
              }),
        ],
      );

  Widget _pingMethodRow() => _simpleRow(
        buttonIcon: Icons.network_check,
        buttonText: 'Ping Server',
        onPressed: bloc.ping,
        blocRebuildCondition: (current) => current is RpcState<PingResult>,
        onOk: (data) => 'Ok Ping Success!',
        unitText: 'Ping Sucess!',
        idleText: 'click button to ping',
      );

  Widget _versionMethodRow() => _simpleRow(
        buttonIcon: Icons.info_outline,
        buttonText: 'Get Version',
        onPressed: bloc.version,
        blocRebuildCondition: (current) => current is RpcState<VersionResult>,
        onOk: (data) => data.maybeWhen(
          versionResult: (version) => version.webrpcVersion,
          orElse: () => 'failed to get version',
        ),
        unitText: 'Got Version!',
        idleText: 'click button to get version',
      );

  Widget _statusMethodRow() => _simpleRow(
        buttonIcon: Icons.cloud,
        buttonText: 'Check Server Status',
        onPressed: bloc.status,
        blocRebuildCondition: (current) => current is RpcState<StatusResult>,
        onOk: (data) => data.maybeWhen(
            statusResult: (status) => 'server up? : $status',
            orElse: () => 'server is down.'),
        unitText: 'Server is up.',
        idleText: 'press button to check server status',
      );
  Widget _addUserRow() {
    final showAlert = (String username) => showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: const Text(
              'Success!',
            ),
            content: Text(
              'User with username: $username added successfully.',
            ),
            actions: <Widget>[
              RaisedButton(
                onPressed: () {
                  Navigator.of(context).pop();
                  bloc.listUsers();
                },
                child: const Text(
                  'Ok',
                ),
                color: Colors.green,
              ),
            ],
          ),
        );
    return Form(
      key: _userFormKey,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: <Widget>[
          const SizedBox(),
          RaisedButton(
            onPressed: () {
              if (_userFormKey.currentState.validate()) {
                _userFormKey.currentState.save();
                bloc.addUser(
                  user: User(
                    id: userID,
                    username: _username,
                    role: 'user',
                  ),
                );
                _userFormKey.currentState.reset();
              }
            },
            color: Colors.blue,
            child: Row(
              children: <Widget>[
                Icon(
                  Icons.add,
                ),
                const Text(
                  'Add User',
                  textAlign: TextAlign.right,
                ),
              ],
            ),
          ),
          const SizedBox(),
          ConstrainedBox(
            constraints: const BoxConstraints(
              maxWidth: 150.0,
            ),
            child: TextFormField(
              decoration: const InputDecoration(
                labelText: "Enter Username",
                hintText: "username",
              ),
              validator: (val) =>
                  val.length > 6 ? null : 'username must be 6 char or more.',
              onSaved: (username) => _username = username,
            ),
          ),
          BlocListener<ExampleServiceBloc, RpcState<ExampleServiceState>>(
            bloc: bloc,
            child: const SizedBox(),
            condition: (previous, current) =>
                current is RpcState<AddUserResult>,
            listener: (context, state) => state.maybeWhen(
              ok: (data) => data.maybeWhen(
                addUserResult: (wasAdded) {
                  _userFormKey.currentState.reset();
                  return showAlert(_username);
                },
                orElse: () => 'failed to add user.',
              ),
              orElse: () => 'failed to add user',
            ),
          ),
        ],
      ),
    );
  }

  Widget _usersTable() {
    return BlocBuilder<ExampleServiceBloc, RpcState<ExampleServiceState>>(
        bloc: bloc,
        condition: (prev, current) => current is RpcState<ListUsersResult>,
        builder: (context, snapshot) {
          return DataTable(
            horizontalMargin: 400.0,
            columns: [
              const DataColumn(
                label: Text(
                  'id',
                ),
              ),
              const DataColumn(
                label: Text(
                  'UserName',
                ),
              ),
              const DataColumn(
                label: Text(
                  'Role',
                ),
              ),
              const DataColumn(
                label: Text(
                  'Delete',
                ),
              ),
            ],
            rows: buildRows(snapshot),
          );
        });
  }

  DataRow _userRow(int id, String username, String role) {
    return DataRow(
      cells: [
        DataCell(
          Text(
            '$id',
          ),
        ),
        DataCell(
          Text(
            username,
          ),
        ),
        DataCell(
          Text(
            role,
          ),
        ),
        DataCell(
          IconButton(
            tooltip: 'Delete User',
            onPressed: () {
              bloc.deleteUser(
                id: id,
              );
              bloc.listUsers();
            },
            color: Colors.red,
            icon: Icon(
              Icons.delete,
            ),
          ),
        ),
      ],
    );
  }

  List<DataRow> buildRows(RpcState<ExampleServiceState> userList) {
    return userList.when(
      err: (err, status, stack) => [],
      idle: () => [],
      loading: () => [],
      unit: () => [],
      ok: (data) => data.maybeWhen(
        listUsersResult: (data) => [
          for (final user in data) _userRow(user.id, user.username, user.role),
        ],
        orElse: () => [],
      ),
    );
  }

  Widget _findUserRow() {
    final showAlert = (User user) => showDialog(
          context: context,
          builder: (context) => AlertDialog(
            title: const Text(
              'User Found.',
            ),
            content: Text(
              user.toJson().toString(),
            ),
            actions: <Widget>[
              RaisedButton(
                onPressed: () {
                  Navigator.of(context).pop();
                },
                child: const Text(
                  'Ok',
                ),
                color: Colors.green,
              ),
            ],
          ),
        );
    return Form(
      key: _findUserFormKey,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: <Widget>[
          const SizedBox(),
          RaisedButton(
            onPressed: () {
              if (_findUserFormKey.currentState.validate()) {
                _findUserFormKey.currentState.save();
                bloc.findUserById(
                  s: SearchFilter(
                    id: _userToFindID,
                  ),
                );
              }
            },
            color: Colors.blue,
            child: Row(
              children: <Widget>[
                Icon(
                  Icons.search,
                ),
                const Text(
                  'Find User',
                  textAlign: TextAlign.right,
                ),
              ],
            ),
          ),
          const SizedBox(),
          ConstrainedBox(
            constraints: const BoxConstraints(
              maxWidth: 150.0,
            ),
            child: TextFormField(
              decoration: const InputDecoration(
                labelText: "Enter User Id",
                hintText: "user Id",
              ),
              validator: (val) =>
                  int.tryParse(val) != null ? null : 'User Id must be a number',
              onSaved: (id) {
                try {
                  _userToFindID = int.parse(id);
                } on Exception catch (e) {
                  print(e);
                }
              },
            ),
          ),
          BlocListener<ExampleServiceBloc, RpcState<ExampleServiceState>>(
            bloc: bloc,
            child: const SizedBox(),
            condition: (previous, current) =>
                current is RpcState<FindUserByIdResult>,
            listener: (context, state) => state.maybeWhen(
              ok: (data) => data.maybeWhen(
                findUserByIdResult: (name, user) {
                  _userFormKey.currentState.reset();
                  return showAlert(user);
                },
                orElse: () => 'Could not find user',
              ),
              orElse: () => 'Could not find user',
            ),
          ),
        ],
      ),
    );
  }
}
