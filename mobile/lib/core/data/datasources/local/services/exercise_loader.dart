import 'dart:convert';
import 'package:flutter/services.dart' show rootBundle;

import '/core/data/models/exercise_info_dto.dart';

abstract class ExerciseInfosLoader {
  Future<List<ExerciseInfoDTO>> loadExercises();
}

class JsonExerciseInfosLoader implements ExerciseInfosLoader {
  static const String _jsonLoadingPath = "assets/exercises.json";
  List<ExerciseInfoDTO>? exercisesInfo;

  @override
  Future<List<ExerciseInfoDTO>> loadExercises() async {
    if (exercisesInfo != null) {
      return exercisesInfo!;
    }
    // Load JSON as a String
    String jsonExercises;
    try {
      jsonExercises = await rootBundle.loadString(_jsonLoadingPath);
    } catch (e) {
      throw Exception("Failed to load JSON asset: ${e.toString()}");
    }

    // Parse it as a list
    List<dynamic> jsonExercisesList;
    try {
      jsonExercisesList = json.decode(jsonExercises);
    } catch (e) {
      throw Exception("Failed to decode JSON ${e.toString()}");
    }

    // Get class instances and return them
    try {
      exercisesInfo = jsonExercisesList.map((e) {
        final ex = ExerciseInfoDTO.fromJsonMap(e);
        return ex;
      }).toList();
    } catch (e) {
      throw Exception("Failed to transfer DTO to Entities ${e.toString()}");
    }

    return exercisesInfo!;
  }
}
