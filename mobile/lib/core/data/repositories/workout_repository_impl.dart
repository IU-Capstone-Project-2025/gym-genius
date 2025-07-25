import 'package:gym_genius/core/data/datasources/local/local_workout_datasource.dart';
import 'package:gym_genius/core/data/datasources/remote/remote_workout_datasource.dart';
import 'package:gym_genius/core/data/models/workout_dto.dart';

import '/core/domain/entities/workout_entity.dart';
import '../../domain/repositories/workout_repository.dart';

class WorkoutRepositoryImpl implements WorkoutRepository {
  final LocalWorkoutDatasource localWorkoutDatasource;
  final RemoteWorkoutDatasource remoteWorkoutDatasource;

  WorkoutRepositoryImpl(
    this.localWorkoutDatasource,
    this.remoteWorkoutDatasource,
  );

  @override
  Future<WorkoutEntity?> fetchWorkout(int workoutId) async {
    final dto = await localWorkoutDatasource.getWorkoutById(workoutId);
    return dto?.toEntity();
  }

  @override
  Future<List<WorkoutEntity>> fetchWorkouts() async {
    final dtos = await localWorkoutDatasource.getAllWorkouts();
    return dtos.map((dto) => dto.toEntity()).toList();
  }

  @override
  Future<String> getAIReview(WorkoutEntity workout) async {
    return remoteWorkoutDatasource
        .getAIDescription(WorkoutDTO.fromEntity(workout));
  }

  @override
  Future<void> saveWorkout(WorkoutEntity entity) async {
    final dto = WorkoutDTO.fromEntity(entity);
    localWorkoutDatasource.saveWorkout(dto);
    remoteWorkoutDatasource.saveWorkout(dto);
  }

  @override
  Future<List<WorkoutEntity>> fetchRemoteWorkouts() async {
    final dtos = await remoteWorkoutDatasource.fetchWorkouts();
    return dtos.map((dto) => dto.toEntity()).toList();
  }
}
