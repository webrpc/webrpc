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
        _pingMethodRow(),
        _versionMethodRow(),
        _statusMethodRow(),
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
