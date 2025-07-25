import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class StartWorkoutPage extends StatelessWidget {
  const StartWorkoutPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: _buildNavBar(context),
      body: Center(
        child: CupertinoButton.tinted(
          child: Text(
            'Start a workout',
          ),
          onPressed: () {
            Navigator.pushNamed(context, '/training_process');
          },
        ),
      ),
    );
  }

  PreferredSizeWidget _buildNavBar(BuildContext context) {
    final schema = Theme.of(context).colorScheme;
    return CupertinoNavigationBar(
      backgroundColor: schema.secondary,
      middle: Text('Profile', style: TextStyle().copyWith(color: schema.onSecondary),),
      trailing: CupertinoButton(
          padding: EdgeInsets.zero,
          child: const Icon(CupertinoIcons.profile_circled),
          onPressed: () {
            Navigator.pushNamed(context, '/auth');
          }),
    );
  }
}
