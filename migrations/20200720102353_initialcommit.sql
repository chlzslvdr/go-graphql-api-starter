-- +goose Up
-- SQL in this section is executed when the migration is applied.
create extension "uuid-ossp";

CREATE TABLE provinces
(
  province_id INT PRIMARY KEY NOT NULL,
  province character varying(200) NOT NULL,
  region_id INT NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE availabilities
(
  availability_id INT PRIMARY KEY NOT NULL,
  availability character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE education_levels
(
  education_level_id INT PRIMARY KEY NOT NULL,
  education_level character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE employment_types
(
  employment_type_id INT PRIMARY KEY NOT NULL,
  employment_type character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE genders
(
  gender_id INT PRIMARY KEY NOT NULL,
  gender character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE impact_areas
(
  impact_area_id INT PRIMARY KEY NOT NULL,
  impact_area character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE importances
(
  importance_id INT PRIMARY KEY NOT NULL,
  importance character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE job_classifications
(
  job_classification_id INT PRIMARY KEY NOT NULL,
  job_classification character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE job_sub_classifications
(
  job_sub_classification_id INT PRIMARY KEY NOT NULL,
  job_sub_classification character varying(200) NOT NULL,
  job_classification_id INT NOT NULL REFERENCES job_classifications(job_classification_id),
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE job_levels
(
  job_level_id INT PRIMARY KEY NOT NULL,
  job_level character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE job_statuses
(
  job_status_id INT PRIMARY KEY NOT NULL,
  job_status character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE pay_periods
(
  pay_period_id INT PRIMARY KEY NOT NULL,
  pay_period character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE skills
(
  skill_id UUID PRIMARY KEY DEFAULT uuid_generate_v1() NOT NULL,
  skill character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE skill_levels
(
  skill_level_id INT PRIMARY KEY NOT NULL,
  skill_level character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL,
  created_when timestamp DEFAULT NOW() NOT NULL,
  updated_when timestamp
);

CREATE TABLE profile_visibilities
(
  profile_visibility_id INT PRIMARY KEY NOT NULL,
  profile_visibility character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL
);

CREATE TABLE cv_templates
(
  cv_template_id INT PRIMARY KEY NOT NULL,
  cv_template character varying(200) NOT NULL,
  cv_template_url character varying(200) NOT NULL,
  is_active boolean DEFAULT true NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS profile_visibilities;
DROP TABLE IF EXISTS cv_templates;
DROP TABLE IF EXISTS provinces
DROP TABLE IF EXISTS availabilities
DROP TABLE IF EXISTS education_levels
DROP TABLE IF EXISTS employment_types
DROP TABLE IF EXISTS genders
DROP TABLE IF EXISTS impact_areas
DROP TABLE IF EXISTS importances
DROP TABLE IF EXISTS job_classifications
DROP TABLE IF EXISTS job_sub_classifications
DROP TABLE IF EXISTS job_levels
DROP TABLE IF EXISTS job_statuses
DROP TABLE IF EXISTS pay_periods
DROP TABLE IF EXISTS skills
DROP TABLE IF EXISTS skill_levels