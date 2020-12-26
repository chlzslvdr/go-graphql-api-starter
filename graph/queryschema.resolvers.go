package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/go-graphl-server-lookup-api/graph/generated"
	"github.com/go-graphl-server-lookup-api/models"
)

func (r *lookupsResolver) Availabilities(ctx context.Context, obj *models.Lookups) ([]*models.Availability, error) {
	availabilities := []*models.Availability{}
	r.DB.Where(models.Availability{IsActive: isActive}).Find(&availabilities)

	return availabilities, nil
}

func (r *lookupsResolver) EducationLevels(ctx context.Context, obj *models.Lookups) ([]*models.EducationLevel, error) {
	educationLevels := []*models.EducationLevel{}
	r.DB.Where(models.EducationLevel{IsActive: isActive}).Find(&educationLevels)

	return educationLevels, nil
}

func (r *lookupsResolver) EmploymentTypes(ctx context.Context, obj *models.Lookups) ([]*models.EmploymentType, error) {
	employmentTypes := []*models.EmploymentType{}
	r.DB.Where(models.EmploymentType{IsActive: isActive}).Find(&employmentTypes)

	return employmentTypes, nil
}

func (r *lookupsResolver) Genders(ctx context.Context, obj *models.Lookups) ([]*models.Gender, error) {
	genders := []*models.Gender{}
	r.DB.Where(models.Gender{IsActive: isActive}).Find(&genders)

	return genders, nil
}

func (r *lookupsResolver) ImpactAreas(ctx context.Context, obj *models.Lookups) ([]*models.ImpactArea, error) {
	impactAreas := []*models.ImpactArea{}
	r.DB.Where(models.ImpactArea{IsActive: isActive}).Order("impact_area").Find(&impactAreas)

	return impactAreas, nil
}

func (r *lookupsResolver) Importances(ctx context.Context, obj *models.Lookups) ([]*models.Importance, error) {
	importances := []*models.Importance{}
	r.DB.Where(models.Importance{IsActive: isActive}).Find(&importances)

	return importances, nil
}

func (r *lookupsResolver) JobClassifications(ctx context.Context, obj *models.Lookups) ([]*models.JobClassification, error) {
	jobClassifications := []*models.JobClassification{}
	r.DB.Where(models.JobClassification{IsActive: isActive}).Order("job_classification").Find(&jobClassifications)

	return jobClassifications, nil
}

func (r *lookupsResolver) JobLevels(ctx context.Context, obj *models.Lookups) ([]*models.JobLevel, error) {
	jobLevels := []*models.JobLevel{}
	r.DB.Where(models.JobLevel{IsActive: isActive}).Find(&jobLevels)

	return jobLevels, nil
}

func (r *lookupsResolver) JobStatuses(ctx context.Context, obj *models.Lookups) ([]*models.JobStatus, error) {
	jobStatuses := []*models.JobStatus{}
	r.DB.Where(models.JobStatus{IsActive: isActive}).Find(&jobStatuses)

	return jobStatuses, nil
}

func (r *lookupsResolver) JobSubClassifications(ctx context.Context, obj *models.Lookups) ([]*models.JobSubClassification, error) {
	jobSubClassifications := []*models.JobSubClassification{}
	r.DB.Where(models.JobSubClassification{IsActive: isActive}).Find(&jobSubClassifications)

	return jobSubClassifications, nil
}

func (r *lookupsResolver) JobSubClassificationsBy(ctx context.Context, obj *models.Lookups, jobClassificationID int) ([]*models.JobSubClassification, error) {
	jobSubClassifications := []*models.JobSubClassification{}
	r.DB.Where(models.JobSubClassification{IsActive: isActive, JobClassificationID: &jobClassificationID}).Find(&jobSubClassifications)

	return jobSubClassifications, nil
}

func (r *lookupsResolver) PayPeriods(ctx context.Context, obj *models.Lookups) ([]*models.PayPeriod, error) {
	payPeriods := []*models.PayPeriod{}
	r.DB.Where(models.PayPeriod{IsActive: isActive}).Find(&payPeriods)

	return payPeriods, nil
}

func (r *lookupsResolver) Provinces(ctx context.Context, obj *models.Lookups) ([]*models.Province, error) {
	provinces := []*models.Province{}
	r.DB.Where(models.Province{IsActive: isActive}).Order("province").Find(&provinces)

	return provinces, nil
}

func (r *lookupsResolver) SkillsBy(ctx context.Context, obj *models.Lookups, keyword string) ([]*models.Skill, error) {
	skills := []*models.Skill{}

	skill := "%" + keyword + "%"
	skill = strings.ToLower(skill)

	r.DB.Where("LOWER( skill ) LIKE ?", skill).Order("skill").Find(&skills)

	return skills, nil
}

func (r *lookupsResolver) SkillLevels(ctx context.Context, obj *models.Lookups) ([]*models.SkillLevel, error) {
	skillLevels := []*models.SkillLevel{}
	r.DB.Where(models.SkillLevel{IsActive: isActive}).Find(&skillLevels)

	return skillLevels, nil
}

func (r *lookupsResolver) SkillCompetencies(ctx context.Context, obj *models.Lookups) ([]*models.Skill, error) {
	skillCompetencies := []*models.Skill{}
	r.DB.Where(models.Skill{IsActive: isActive}).Order("skill").Find(&skillCompetencies)

	return skillCompetencies, nil
}

func (r *lookupsResolver) SkillCompetenciesByKeyword(ctx context.Context, obj *models.Lookups, keyword string) ([]*models.Skill, error) {
	skillCompetencies := []*models.Skill{}

	skillCompetency := "%" + keyword + "%"
	skillCompetency = strings.ToLower(skillCompetency)

	r.DB.Where("LOWER( skill ) LIKE ?", skillCompetency).Order("skill").Find(&skillCompetencies)

	return skillCompetencies, nil
}

func (r *lookupsResolver) SkillCompetencyLevels(ctx context.Context, obj *models.Lookups) ([]*models.SkillLevel, error) {
	skillCompetencyLevels := []*models.SkillLevel{}
	r.DB.Where(models.SkillLevel{IsActive: isActive}).Find(&skillCompetencyLevels)

	return skillCompetencyLevels, nil
}

func (r *lookupsResolver) ProfileVisibilities(ctx context.Context, obj *models.Lookups) ([]*models.ProfileVisibility, error) {
	profileVisibilities := []*models.ProfileVisibility{}
	r.DB.Where(models.ProfileVisibility{IsActive: isActive}).Find(&profileVisibilities)

	return profileVisibilities, nil
}

func (r *lookupsResolver) CvTemplates(ctx context.Context, obj *models.Lookups) ([]*models.CVTemplate, error) {
	cvTemplates := []*models.CVTemplate{}
	r.DB.Where(models.CVTemplate{IsActive: isActive}).Find(&cvTemplates)

	return cvTemplates, nil
}

func (r *queryResolver) Lookups(ctx context.Context) (*models.Lookups, error) {
	lookups := models.Lookups{}
	return &lookups, nil
}

// Lookups returns generated.LookupsResolver implementation.
func (r *Resolver) Lookups() generated.LookupsResolver { return &lookupsResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type lookupsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
