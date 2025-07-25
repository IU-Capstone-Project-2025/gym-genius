import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:gym_genius/core/data/datasources/mock_data.dart';
import 'package:gym_genius/core/domain/entities/workout_entity.dart';
import 'package:gym_genius/core/domain/repositories/workout_repository.dart';
import 'package:gym_genius/di.dart';

import 'stats_widgets.dart';

class StatsPage extends StatefulWidget {
  const StatsPage({super.key});

  @override
  State<StatsPage> createState() => _StatsPageState();
}

class _StatsPageState extends State<StatsPage> {
  late Future<List<WorkoutEntity>> _workoutsF;

  @override
  void initState() {
    super.initState();
    _workoutsF = getIt<WorkoutRepository>().fetchWorkouts();
  }

  @override
  Widget build(BuildContext context) {
    final schema = Theme.of(context).colorScheme;

    return Scaffold(
      appBar: CupertinoNavigationBar.large(
        border: null,
        backgroundColor: schema.surfaceContainer,
        largeTitle: Text(
          'Statistics',
          style: TextStyle().copyWith(color: schema.primary),
        ),
        trailing: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            CupertinoButton(
                padding: EdgeInsets.zero,
                onPressed: addMocks,
                child: Icon(
                  CupertinoIcons.add,
                  size: 32,
                )),
            CupertinoButton(
                padding: EdgeInsets.zero,
                onPressed: refresh,
                child: Icon(
                  CupertinoIcons.refresh,
                  size: 32,
                )),
          ],
        ),
      ),
      body: FutureBuilder<List<WorkoutEntity>>(
        future: _workoutsF,
        builder: (_, snap) {
          if (snap.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator.adaptive());
          }
          if (snap.hasError) {
            return Center(child: Text('Error: ${snap.error}'));
          }

          final workouts = snap.data!;

          // Normalize each startTime to a pure date
          final workoutDays = workouts
              .map((w) => DateTime(
                  w.startTime.year, w.startTime.month, w.startTime.day))
              .toSet();

          return SingleChildScrollView(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ActivityGrid(workoutDays: workoutDays),
                const SizedBox(height: 52),
                _buildWorkoutDurationChart(workouts, schema),
              ],
            ),
          );
        },
      ),
    );
  }

  void addMocks() {
    setState(() {
      _workoutsF = Future.value(getMockWorkouts(100));
    });
  }

  void refresh() {
    setState(() {
      _workoutsF = getIt<WorkoutRepository>().fetchRemoteWorkouts();
    });
  }

  Widget _buildWorkoutDurationChart(
      List<WorkoutEntity> workouts, ColorScheme schema) {
    if (workouts.isEmpty) {
      return const SizedBox(
        height: 200,
        child: Center(child: Text('No workout data available')),
      );
    }

    // Sort workouts by date
    final sortedWorkouts = List<WorkoutEntity>.from(workouts)
      ..sort((a, b) => a.startTime.compareTo(b.startTime));

    // Create data points for the chart
    final spots = <FlSpot>[];
    for (int i = 0; i < sortedWorkouts.length; i++) {
      final workout = sortedWorkouts[i];
      final durationMinutes = workout.duration.inMinutes.toDouble();
      spots.add(FlSpot(i.toDouble(), durationMinutes));
    }

    return Padding(
      padding: const EdgeInsets.all(16),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            'This is your workout duration over time',
            style: Theme.of(context).textTheme.titleMedium,
          ),
          const SizedBox(height: 16),
          SizedBox(
            height: 200,
            child: LineChart(
              LineChartData(
                gridData: FlGridData(
                  show: true,
                  drawVerticalLine: false,
                  horizontalInterval: 15,
                  getDrawingHorizontalLine: (value) {
                    return FlLine(
                      color: schema.outline.withOpacity(0.3),
                      strokeWidth: 1,
                    );
                  },
                ),
                titlesData: FlTitlesData(
                  show: true,
                  rightTitles: const AxisTitles(
                      sideTitles: SideTitles(showTitles: false)),
                  topTitles: const AxisTitles(
                      sideTitles: SideTitles(showTitles: false)),
                  bottomTitles: AxisTitles(
                    sideTitles: SideTitles(
                      showTitles: true,
                      reservedSize: 30,
                      interval: spots.length > 10
                          ? (spots.length / 5).floorToDouble()
                          : 1,
                      getTitlesWidget: (double value, TitleMeta meta) {
                        final index = value.toInt();
                        if (index >= 0 && index < sortedWorkouts.length) {
                          final date = sortedWorkouts[index].startTime;
                          return SideTitleWidget(
                            meta: meta, // Add the meta parameter here
                            child: Text(
                              '${date.month}/${date.day}',
                              style: TextStyle(
                                color: schema.onSurface.withOpacity(0.6),
                                fontSize: 12,
                              ),
                            ),
                          );
                        }
                        return const Text('');
                      },
                    ),
                  ),
                  leftTitles: AxisTitles(
                    sideTitles: SideTitles(
                      showTitles: true,
                      interval: 30,
                      reservedSize: 42,
                      getTitlesWidget: (double value, TitleMeta meta) {
                        return SideTitleWidget(
                          meta: meta, // Add the meta parameter here
                          child: Text(
                            '${value.toInt()}m',
                            style: TextStyle(
                              color: schema.onSurface.withOpacity(0.6),
                              fontSize: 12,
                            ),
                          ),
                        );
                      },
                    ),
                  ),
                ),
                borderData: FlBorderData(
                  show: true,
                  border: Border.all(color: schema.outline.withOpacity(0.3)),
                ),
                minX: 0,
                maxX: (spots.length - 1).toDouble(),
                minY: 0,
                maxY:
                    spots.map((e) => e.y).reduce((a, b) => a > b ? a : b) + 10,
                lineBarsData: [
                  LineChartBarData(
                    spots: spots,
                    isCurved: true,
                    gradient: LinearGradient(
                      colors: [
                        schema.primary.withOpacity(0.8),
                        schema.primary,
                      ],
                      begin: Alignment.centerLeft,
                      end: Alignment.centerRight,
                    ),
                    barWidth: 3,
                    isStrokeCapRound: true,
                    dotData: FlDotData(
                      show: true,
                      getDotPainter: (spot, percent, barData, index) {
                        return FlDotCirclePainter(
                          radius: 4,
                          color: schema.primary,
                          strokeWidth: 2,
                          strokeColor: schema.surface,
                        );
                      },
                    ),
                    belowBarData: BarAreaData(
                      show: true,
                      gradient: LinearGradient(
                        colors: [
                          schema.primary.withOpacity(0.1),
                          schema.primary.withOpacity(0.05),
                        ],
                        begin: Alignment.topCenter,
                        end: Alignment.bottomCenter,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
