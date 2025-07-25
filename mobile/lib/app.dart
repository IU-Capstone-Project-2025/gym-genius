import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

import '../config/routes.dart' show routes;
import '../config/themes.dart' show brightAppTheme, darkAppTheme;

import 'package:http/http.dart' as http;

class MainApp extends StatefulWidget {
  const MainApp({super.key});

  @override
  State<MainApp> createState() => _MainAppState();
}

class _MainAppState extends State<MainApp> {
  @override
  void initState() {
    super.initState();
    fetchData();
  }

  Future<void> fetchData() async {
    // Define the URL for the POST request
    final url =
        Uri.parse('http://api.xn--b1ab5acc.site/statistics/add_activity');
    // Create the JSON body for the request
    Map<String, dynamic> body = {
      "date": DateTime.now().toString(),
      "user_id": 1,
      // Add other key-value pairs as needed
    };
    try {
      // Make the POST request
      final response = await http.post(
        url,
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: json.encode(body),
      );
      if (response.statusCode == 200) {
        // Parse the response
        var data = json.decode(response.body);
        print(data);
      } else {
        throw Exception('Failed to load data');
      }
    } catch (e) {
      print('Error: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      localizationsDelegates: const <LocalizationsDelegate<dynamic>>[
        DefaultMaterialLocalizations.delegate,
        DefaultCupertinoLocalizations.delegate,
      ],
      debugShowCheckedModeBanner: false,
      theme: brightAppTheme,
      darkTheme: darkAppTheme,
      initialRoute: '/',
      routes: routes,
      themeMode: ThemeMode.system,
    );
  }
}
