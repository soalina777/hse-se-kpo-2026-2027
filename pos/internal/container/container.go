package container

import (
	"pos/internal/application"
	"pos/internal/domain"
	"pos/internal/infrastructure"
)

type Container struct {
	Facade *application.POSFacade
}

func NewContainer() *Container {
	studentRepo := infrastructure.NewMemoryStudentRepository()
	criterionRepo := infrastructure.NewMemoryCriterionRepository()
	evaluationRepo := infrastructure.NewMemoryEvaluationRepository()

	evaluationFactory := domain.EvaluationFactory{}
	calculator := domain.NewScoreCalculator(
		domain.WeightedSumStrategy{},
	)

	csvExporter := infrastructure.NewCSVExporter()

	managementService := application.NewManagementService(
		studentRepo,
		criterionRepo,
	)

	evaluationService := application.NewEvaluationService(
		studentRepo,
		criterionRepo,
		evaluationRepo,
		evaluationFactory,
	)

	resultService := application.NewResultService(
		studentRepo,
		evaluationRepo,
		criterionRepo,
		calculator,
	)

	exportService := application.NewExportService(
		studentRepo,
		criterionRepo,
		evaluationRepo,
		calculator,
		csvExporter,
	)

	facade := application.NewPOSFacade(
		managementService,
		evaluationService,
		resultService,
		exportService,
	)

	return &Container{
		Facade: facade,
	}
}
