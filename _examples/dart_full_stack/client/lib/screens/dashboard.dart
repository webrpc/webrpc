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
    return pingServiceRow();
  }

  Widget pingServiceRow() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: <Widget>[
        RaisedButton(
          onPressed: () {
            print('button clicked');
            bloc.add(
              ExampleServiceEvent.ping(),
            );
          },
          child: Row(
            children: <Widget>[
              Icon(
                Icons.network_check,
              ),
              const Text(
                ' Ping Server',
              ),
            ],
          ),
          color: Colors.blue,
        ),
        BlocBuilder<ExampleServiceBloc, RpcState<ExampleServiceState>>(
            bloc: bloc,
            builder: (context, state) {
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(
                  state.when(
                    err: (reason, status, _) =>
                        'An error occured: $reason status: $status',
                    idle: () => 'waiting to ping',
                    loading: () => 'Loading Results',
                    ok: (ExampleServiceState data) => 'Ok Ping Success!',
                    unit: () => 'Unit Ping Sucess!',
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
