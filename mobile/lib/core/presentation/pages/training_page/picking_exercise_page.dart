import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:gym_genius/core/data/datasources/local/services/exercise_loader.dart';
import 'package:gym_genius/core/data/models/exercise_info_dto.dart';

import '/core/domain/entities/exercise_info_entity.dart';
import '/core/presentation/bloc/training_bloc.dart';
import '/core/presentation/bloc/training_event.dart';
import '/core/presentation/bloc/training_state.dart';
import '/core/presentation/shared/warnings.dart';
import '../../shared/custom_context_menu.dart';

import '/di.dart';

class PickingExercisePage extends StatefulWidget {
  final void Function(ExerciseInfoEntity) addExerciseCallback;
  const PickingExercisePage({super.key, required this.addExerciseCallback});

  @override
  State<PickingExercisePage> createState() => _PickingExercisePageState();
}

class _PickingExercisePageState extends State<PickingExercisePage> {
  List<ExerciseInfoEntity> exercises = [];

  // TODO: change into entity
  late Future<List<ExerciseInfoDTO>> fetchExercisesFuture;

  @override
  void initState() {
    super.initState();
    fetchExercisesFuture = getIt<ExerciseInfosLoader>().loadExercises();
  }

  @override
  Widget build(BuildContext context) {
    final double heightOfWidget = MediaQuery.of(context).size.height - 150;

    return BlocConsumer<TrainingBloc, TrainingState>(
      listener: (context, state) {
        switch (state.addExerciseStatus) {
          case AddExerciseStatus.duplicate:
            Warnings.showAlreadyHasExerciseWarning(context);
            break;
          case AddExerciseStatus.success:
            Navigator.pop(context);
          default:
            break;
        }
      },
      builder: (context, state) {
        final bloc = context.read<TrainingBloc>();

        return SafeArea(
          bottom: false,
          child: CupertinoPopupSurface(
            child: Material(
              child: CupertinoPageScaffold(
                child: SizedBox(
                  height: heightOfWidget,
                  child: CustomScrollView(
                    slivers: [
                      _buildNavBar(),
                      _buildSearchingSegment(),
                      _buildGrid(bloc),
                    ],
                  ),
                ),
              ),
            ),
          ),
        );
      },
    );
  }

  Widget _buildNavBar() {
    return CupertinoSliverNavigationBar(
      border: null,
      backgroundColor: Theme.of(context).colorScheme.surface,
      automaticBackgroundVisibility: false,
      largeTitle: Text(
        'Pick an Exercise',
        style: TextStyle().copyWith(
          color: Theme.of(context).colorScheme.onSecondary,
        ),
      ),
    );
  }

  Widget _buildSearchingSegment() {
    return PinnedHeaderSliver(
      child: Column(
        children: [
          Container(
            padding: EdgeInsets.symmetric(horizontal: 16),
            color: Theme.of(context).colorScheme.surface,
            child: Column(
              children: [
                Row(
                  children: [
                    Flexible(
                      flex: 9,
                      child: SizedBox(
                        child: CupertinoSearchTextField(),
                      ),
                    ),
                    CupertinoButton(
                      padding: EdgeInsets.all(0),
                      onPressed: () {
                        showCupertinoModalPopup(
                          context: context,
                          builder: (context) {
                            return _buildFilteringSegment();
                          },
                        );
                      },
                      child: Icon(Icons.filter_list),
                    ),
                  ],
                ),
                SizedBox(
                  height: 4,
                ),
              ],
            ),
          ),
          _buildBorder(),
        ],
      ),
    );
  }

  Widget _buildBorder() => Container(
        height: 1,
        width: double.infinity,
        color: CupertinoDynamicColor.withBrightness(
          color: Color(0x33000000),
          darkColor: Color(0x33FFFFFF),
        ),
      );

  Widget _buildFilteringSegment() {
    return Material(
      type: MaterialType.transparency,
      child: Center(
        child: Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(12),
            color: Colors.transparent.withAlpha(100),
          ),
          height: 200,
          width: 300,
          child: DefaultTextStyle(
            style: TextStyle().copyWith(color: Colors.white),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Padding(
                  padding: EdgeInsets.all(16),
                  child: Text('Pick Filtering Options'),
                ),
                Expanded(
                  child: Placeholder(
                    child: Text('COMING SOON'),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildGrid(TrainingBloc bloc) {
    final double bottomPadding = MediaQuery.of(context).size.width / 4;

    return FutureBuilder(
      future: fetchExercisesFuture,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.done) {
          if (snapshot.hasError) {
            return SliverToBoxAdapter(
              child: Center(
                child: SizedBox(
                    height: 500, child: Text('Error: ${snapshot.error}')),
              ),
            );
          }

          List<ExerciseInfoDTO> exercises = snapshot.data!;

          return SliverPadding(
            padding: EdgeInsets.only(
              top: 16,
              left: 16,
              right: 16,
              bottom: bottomPadding,
            ),
            sliver: SliverGrid.builder(
              itemCount: exercises.length,
              gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                crossAxisSpacing: 16,
                mainAxisSpacing: 16,
              ),
              itemBuilder: (context, val) {
                ExerciseInfoEntity exerciseInfo = exercises[val].toEntity();
                return _buildExerciseContainer(exerciseInfo, bloc);
              },
            ),
          );
        } else {
          return SliverFillRemaining(
            child: CircularProgressIndicator.adaptive(),
          );
        }
      },
    );
  }

  Widget _buildExerciseContainer(
    ExerciseInfoEntity exerciseInfo,
    TrainingBloc bloc,
  ) {
    final double focusedContainerHeight =
        MediaQuery.of(context).size.width / 1.5;

    return CustomContextMenu(
      actions: [
        CupertinoContextMenuAction(
          isDefaultAction: true,
          child: Text(
            '${exerciseInfo.name}. ${exerciseInfo.description}',
          ),
        ),
        CupertinoContextMenuAction(
          onPressed: () {
            Navigator.pop(context); // Close the context menu
            // Also it should locally add ex to favorite.
          },
          trailingIcon: CupertinoIcons.heart,
          child: const Text('Favorite'),
        ),
      ],
      // TODO: Fix taps on focused
      child: GestureDetector(
        onTap: () {
          print("Added ex to bloc");
          bloc.add(AddExercise(exerciseInfo));
        },

        /// Container itself
        child: Container(
          height: focusedContainerHeight,
          width: focusedContainerHeight,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(12),
            color: Theme.of(context).colorScheme.primaryContainer,
            image:
                exerciseInfo.imagePath != null && exerciseInfo.imagePath != ""
                    ? DecorationImage(
                        image: AssetImage(exerciseInfo.imagePath!),
                        fit: BoxFit.cover)
                    : null,
          ),
          alignment: Alignment.bottomLeft,

          /// Little Text with decoration
          child: Material(
            type: MaterialType.transparency,
            child: Container(
              padding: EdgeInsets.all(4),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.only(topRight: Radius.circular(4)),
                color: Theme.of(context).colorScheme.secondary.withAlpha(230),
              ),
              child: Text(
                exerciseInfo.name,
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
            ),
          ),
        ),
      ),
    );
  }
}

/// Might be extracted to a different page
class CustomSegmentedControl extends StatefulWidget {
  const CustomSegmentedControl({super.key});

  @override
  State<CustomSegmentedControl> createState() => _CustomSegmentedControlState();
}

class _CustomSegmentedControlState extends State<CustomSegmentedControl> {
  int _selectedFilter = 1;

  @override
  Widget build(BuildContext context) {
    return CupertinoSegmentedControl<int>(
      groupValue: _selectedFilter,
      children: <int, Widget>{
        0: Padding(
          padding: EdgeInsets.symmetric(
            horizontal: 20,
          ),
          child: Text('Chest'),
        ),
        1: Padding(
          padding: EdgeInsets.symmetric(
            horizontal: 0,
          ),
          child: Text('Back'),
        ),
        2: Padding(
          padding: EdgeInsets.symmetric(
            horizontal: 0,
          ),
          child: Text('Legs'),
        ),
      },
      onValueChanged: (value) {
        setState(() {
          _selectedFilter = value;
        });
      },
    );
  }
}
