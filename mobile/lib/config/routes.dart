import 'package:flutter/cupertino.dart';
import 'package:gym_genius/core/presentation/pages/auth_page/auth_page.dart';

import '/core/presentation/pages/stats_page/stats_page.dart';
import '/core/presentation/pages/training_page/start_workout_page.dart';
import '/core/presentation/pages/training_page/training_process_page.dart';

final routes = {
  '/': (context) => const TabBarPage(),
  '/training_process': (context) => const TrainingProcessPage(),
  '/auth': (context) => const AuthPage(),
};

class TabBarPage extends StatelessWidget {
  const TabBarPage({super.key});

  @override
  Widget build(BuildContext context) {
    return CupertinoTabScaffold(
      tabBar: CupertinoTabBar(
        items: const [
          BottomNavigationBarItem(icon: Icon(CupertinoIcons.flame), label: 'Train'),
          BottomNavigationBarItem(icon: Icon(CupertinoIcons.chart_bar), label: 'Stats'),
        ],
      ),
      tabBuilder: (ctx, i) {
        switch (i) {
          case 0:
            return StartWorkoutPage();
          case 1:
            return StatsPage();
          default:
            return const SizedBox();
        }
      },
    );
  }
}
