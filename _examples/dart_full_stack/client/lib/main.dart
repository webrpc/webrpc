import 'package:client/screens/dashboard.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'WebRpc Example',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

Widget buildTitle() => Row(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: <Widget>[
        Container(
          margin: const EdgeInsets.only(left: 32),
          child: Text(
            "WebRpc Example",
            style: TextStyle(
              fontSize: 24,
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontFamily: 'HelveticaNeue',
            ),
          ),
        ),
      ],
    );

Widget buildDrawer() => Drawer(
      child: ListView(
        children: <Widget>[
          DrawerHeader(
            child: const Text(
              'ExampleService Methods',
            ),
            decoration: BoxDecoration(
              color: Colors.blue,
            ),
          ),
          const ListTile(
            title: Text(
              'Ping()',
            ),
          ),
          const ListTile(
            title: Text(
              'Version()',
            ),
          ),
          const ListTile(
            title: Text(
              'Status()',
            ),
          ),
          const ListTile(
            title: Text(
              'GetUser()',
            ),
          ),
          const ListTile(
            title: Text(
              'UpdateName()',
            ),
          ),
          const ListTile(
            title: Text(
              'ListUsers()',
            ),
          ),
          const ListTile(
            title: Text(
              'FindUserById()',
            ),
          ),
          const ListTile(
            title: Text(
              'AddUser()',
            ),
          ),
          const ListTile(
            title: Text(
              'DeleteUser()',
            ),
          ),
        ],
      ),
    );

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading:
            MediaQuery.of(context).size.width < 1300 ? true : false,
        title: buildTitle(),
      ),
      drawer: buildDrawer(),
      body: DashBoard(),
    );
  }
}
