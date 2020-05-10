import 'package:client/services/example/example_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class DashBoard extends StatelessWidget {
  final ExampleServiceBloc bloc = ExampleServiceBloc(
    exampleService: ExampleServiceRpc(
      host: 'http://localhost:8080',
    ),
  );
  void dispose() {
    bloc.close();
  }

  DashBoard({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        pingMethodRow(),
        versionMethodRow(),
      ],
    );
  }

  Widget pingMethodRow() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: <Widget>[
        RaisedButton(
          onPressed: bloc.ping,
          child: Row(
            children: <Widget>[
              Icon(
                Icons.network_check,
              ),
              const Text(
                'Ping Server',
              ),
            ],
          ),
          color: Colors.blue,
        ),
        BlocBuilder<ExampleServiceBloc, RpcState<ExampleServiceState>>(
            bloc: bloc,
            condition: (previous, current) => current is RpcState<PingResult>,
            builder: (context, state) {
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(
                  state.when(
                    err: (reason, status, _) =>
                        'An error occured: $reason status: $status',
                    idle: () => 'click button to ping',
                    loading: () => 'Loading Results',
                    ok: (ExampleServiceState data) => 'Ok Ping Success!',
                    unit: () => 'Ping Sucess!',
                  ),
                ),
              );
            }),
      ],
    );
  }

  Widget versionMethodRow() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: <Widget>[
        RaisedButton(
          onPressed: bloc.version,
          child: Row(
            children: <Widget>[
              Icon(
                Icons.info,
              ),
              const Text(
                'Get Version',
              ),
            ],
          ),
          color: Colors.blue,
        ),
        BlocBuilder<ExampleServiceBloc, RpcState<ExampleServiceState>>(
            bloc: bloc,
            condition: (previous, current) =>
                current is RpcState<VersionResult>,
            builder: (context, state) {
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(
                  state.when(
                    err: (reason, status, _) =>
                        'An error occured: $reason status: $status',
                    idle: () => 'click button to get version',
                    loading: () => 'Loading Results',
                    ok: (ExampleServiceState data) =>
                        'webrpc version is: $data',
                    unit: () => 'version Sucess!',
                  ),
                ),
              );
            }),
      ],
    );
  }

  Widget _usersTable() {
    return StreamBuilder(builder: (context, snapshot) {
      return DataTable(
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
        ],
        rows: buildRows(),
      );
    });
  }

  DataRow _userRow(String id, String username, String role) {
    return DataRow(
      cells: [
        DataCell(
          Text(
            id,
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
      ],
    );
  }

  List<DataRow> buildRows() {
    return null;
  }
}
