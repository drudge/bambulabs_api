package bambulabs_api

// PreparationStage represents the current stage during print preparation
type PreparationStage int

const (
	PrepStageIdle              PreparationStage = 0  // Printing or Idle
	PrepStageAutoBedLeveling   PreparationStage = 1  // Auto Bed Leveling
	PrepStageHeatingBed        PreparationStage = 2  // Heatbed Preheating
	PrepStageNozzleCleaning    PreparationStage = 3  // Nozzle Cleaning
	PrepStageNozzleCalibrating PreparationStage = 4  // Nozzle Calibrating
	PrepStageSweepingXY        PreparationStage = 5  // Sweeping XY Mech Mode
	PrepStageBedLevelRefining  PreparationStage = 6  // Bed Leveling (Refining)
	PrepStageHeatingNozzle     PreparationStage = 7  // Nozzle Heating
	PrepStageCoolingNozzle     PreparationStage = 8  // Cooling Nozzle
	PrepStagePausedUser        PreparationStage = 9  // Paused (Wait for User)
	PrepStageChangingFilament  PreparationStage = 10 // Changing Filament (AMS Action)
	PrepStageFilamentPullback  PreparationStage = 11 // Filament Pulling Back
	PrepStagePurgingFilament   PreparationStage = 12 // Purging Filament
)

// String returns a human-readable description of the preparation stage
func (s PreparationStage) String() string {
	switch s {
	case PrepStageIdle:
		return "Idle"
	case PrepStageAutoBedLeveling:
		return "Auto Bed Leveling"
	case PrepStageHeatingBed:
		return "Heating Bed"
	case PrepStageNozzleCleaning:
		return "Cleaning Nozzle"
	case PrepStageNozzleCalibrating:
		return "Calibrating Nozzle"
	case PrepStageSweepingXY:
		return "Sweeping XY"
	case PrepStageBedLevelRefining:
		return "Refining Bed Level"
	case PrepStageHeatingNozzle:
		return "Heating Nozzle"
	case PrepStageCoolingNozzle:
		return "Cooling Nozzle"
	case PrepStagePausedUser:
		return "Waiting for User"
	case PrepStageChangingFilament:
		return "Changing Filament"
	case PrepStageFilamentPullback:
		return "Pulling Back Filament"
	case PrepStagePurgingFilament:
		return "Purging Filament"
	default:
		return "Unknown"
	}
}
