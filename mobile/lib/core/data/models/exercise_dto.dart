import 'dart:convert';
import 'dart:math';
import 'package:flutter/foundation.dart';

import '../../domain/entities/exercise_entity.dart';
import 'exercise_info_dto.dart';
import 'exercise_set_dto.dart';

class ExerciseDTO {
  final ExerciseInfoDTO exerciseInfo;
  List<ExerciseSetDTO> sets;

  ExerciseDTO({
    required this.exerciseInfo,
    this.sets = const [],
  });

  ExerciseDTO copyWith({
    ExerciseInfoDTO? exerciseInfo,
    List<ExerciseSetDTO>? sets,
  }) {
    return ExerciseDTO(
      exerciseInfo: exerciseInfo ?? this.exerciseInfo,
      sets: sets ?? this.sets,
    );
  }

  ExerciseEntity toEntity() {
    return ExerciseEntity(
      exerciseInfo: exerciseInfo.toEntity(),
      sets: sets.map((obj) => obj.toEntity()).toList(),
    );
  }

  factory ExerciseDTO.fake() {
    final rand = Random();

    return ExerciseDTO(
      exerciseInfo: ExerciseInfoDTO(
        id: rand.nextInt(28) + 1,
        name: "Random exercise",
      ),
      sets: List.generate(
        rand.nextInt(5) + 1,
        (id) => ExerciseSetDTO(
          weight: rand.nextInt(100) + 1,
          reps: rand.nextInt(20) + 1,
        ),
      ),
    );
  }

  factory ExerciseDTO.fromEntity(ExerciseEntity entity) {
    return ExerciseDTO(
      exerciseInfo: ExerciseInfoDTO.fromEntity(entity.exerciseInfo),
      sets: entity.sets.map((obj) => ExerciseSetDTO.fromEntity(obj)).toList(),
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'exercise_info': exerciseInfo.toMap().toString(),
      'sets': sets.map((x) => x.toMap()).toList(),
    };
  }

  Map<String, dynamic> toAnotherFuckingShit() {
    return <String, dynamic>{
      'exercise_id': exerciseInfo.id,
      'reps'
          'sets': sets.map((x) => x.toMap()).toList(),
    };
  }

  List<Map<String, num>> toShitList() {
    final List<Map<String, num>> list = [];

    for (final set in sets) {
      list.add({
        'exercise_id': exerciseInfo.id,
        'reps': set.reps,
        'weight': set.weight,
      });
    }

    return list;
  }

  Map<String, dynamic> toFuckingShit() {
    return <String, dynamic>{
      'exerciseInfo': exerciseInfo.description,
      'sets': sets.map((x) => x.toMap()).toList(),
      'muscleGroup': exerciseInfo.muscleGroups!.first,
      'name': exerciseInfo.name,
    };
  }

  factory ExerciseDTO.fromMap(Map<String, dynamic> map) {
    return ExerciseDTO(
      exerciseInfo: ExerciseInfoDTO.fromSQLMap(
        map['exercise_info'] as Map<String, dynamic>,
      ),
      sets: (map['sets'] as List)
          .map((x) => ExerciseSetDTO.fromMap(x as Map<String, dynamic>))
          .toList(),
    );
  }

  String toJson() => json.encode(toMap());

  factory ExerciseDTO.fromJson(String source) =>
      ExerciseDTO.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() => 'Exercise(exerciseInfo: $exerciseInfo, sets: $sets)';

  @override
  bool operator ==(covariant ExerciseDTO other) {
    if (identical(this, other)) return true;

    return other.exerciseInfo == exerciseInfo && listEquals(other.sets, sets);
  }

  @override
  int get hashCode => exerciseInfo.hashCode ^ sets.hashCode;
}
