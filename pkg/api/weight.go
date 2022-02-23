package api

type WeightService interface {
	//New(request NewWeightRequest) error
	CalculateBMR(height, age, weight int, sex string) (int, error)
	DailyIntake(BMR, activityLevel int, weightGoal string) (int, error)
}

type WeightRepository interface {
	//CreateWeightEntry(w Weight) error
	//GetUser(userID int) (User, error)
}

type weightService struct {
	storage WeightRepository
}

func NewWeightService(weightRepo WeightRepository) WeightService {
	//return &weightService{storage: weightRepo}
	return nil
}

const (
	// Very Low Intensity activity multiplier - 1
	veryLowActivity = 1.2
	// Light exercise activity multiplier - 3-4x per week - 2
	lightActivity = 1.375
	// Moderate exercise activity multiplier - 3-5x per week 30-60 mins/session - 3
	moderateActivity = 1.55
	// High exercise activity multiplier - (6-7x per week for 45-60 mins) - 4
	highActivity = 1.725
	// Very high exercise activity multiplier - for people who is an Athlete - 5
	veryHighActivity = 1.9
)
