import 'package:gym_genius/core/domain/entities/workout_entity.dart';

import '/core/domain/entities/exercise_info_entity.dart';

sealed class TrainingEvent {
  const TrainingEvent();
}

class TestGetRandomWorkout extends TrainingEvent {
  const TestGetRandomWorkout();
}

class AddExercise extends TrainingEvent {
  final ExerciseInfoEntity info;
  const AddExercise(this.info);
}

class RemoveExercise extends TrainingEvent {
  final int exerciseID;
  const RemoveExercise(this.exerciseID);
}

class SubmitTraining extends TrainingEvent {
  final Duration workoutDuration;
  const SubmitTraining(this.workoutDuration);
}

class GetAIReview extends TrainingEvent {
  final WorkoutEntity workout;
  const GetAIReview(this.workout);
}
