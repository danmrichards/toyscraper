package schema

// JobRequirement represents a specific requirement for a job posting.
type JobRequirement struct {
	Category string   `json:"category" description:"Category of the requirement (e.g., Technical, Soft Skills)"`
	Items    []string `json:"items" description:"List of specific requirements in this category"`
	Priority string   `json:"priority" description:"Priority level (Required/Preferred) based on the HTML class or context"`
}

// Company represents the company information in a job posting.
type Company struct {
	Name             string            `json:"name" description:"Company name"`
	Description      string            `json:"description" description:"Company description or overview"`
	Website          string            `json:"website" description:"Company website URL"`
	LogoURL          string            `json:"logo_url" description:"URL to the company logo"`
	SocialMediaLinks map[string]string `json:"social_media_links" description:"Company social media links (e.g., Github, LinkedIn, Twitter, X etc) as key-value pairs"`
	Culture          []string          `json:"culture" description:"List of company-related information, such as values, mission, or culture"`
	ContactInfo      map[string]string `json:"contact_info" description:"Contact information from header, footer or contact section"`
}

// Benefits represents the benefits or perks offered by the company in a job posting.
type Benefits struct {
	Items              []string `json:"items" description:"List of job benefits or perks"`
	AdditionalInfo     string   `json:"additional_info" description:"Additional information about the benefits or perks, if available"`
	AdditionalInfoLink string   `json:"additional_info_link" description:"Link to more information about the benefits or perks, if provided"`
	Sponsorship        string   `json:"sponsorship" description:"Sponsorship information, if applicable (e.g., visa sponsorship)"`
}

// PositionType represents the type of position in a job posting.
type PositionType string

const (
	FullTime       PositionType = "Full Time"
	PartTime       PositionType = "Part Time"
	Contract       PositionType = "Contract"
	Internship     PositionType = "Internship"
	Temporary      PositionType = "Temporary"
	Volunteer      PositionType = "Volunteer"
	Freelance      PositionType = "Freelance"
	Apprenticeship PositionType = "Apprenticeship"
	Other          PositionType = "Other"
)

// ApplicationProcess represents the application process details in a job posting.
type ApplicationProcess struct {
	Deadline        string `json:"deadline" description:"Application deadline if specified"`
	Link            string `json:"link" description:"Link to the application page"`
	Details         string `json:"details" description:"Details about the application process, if available"`
	NonSolicitation string `json:"non_solicitation" description:"Non-solicitation clause or information, if available"`
	ManagedBy       string `json:"managed_by" description:"Information about the person, recruitment agency or platform managing the application, if available"`
}

// Package represents the compensation package details in a job posting.
type Package struct {
	SalaryRange          string   `json:"salary_range" description:"Salary range if specified"`
	Bonus                string   `json:"bonus" description:"Bonus information if specified"`
	Equity               string   `json:"equity" description:"Equity information if specified"`
	Commission           string   `json:"commission" description:"Commission information if specified"`
	SalaryCurrency       string   `json:"salary_currency" description:"Currency of the salary range"`
	SalaryCurrencySymbol string   `json:"salary_currency_symbol" description:"Currency symbol of the salary range"`
	OnCallAllowance      string   `json:"on_call_allowance" description:"On-call allowance information if specified"`
	Benefits             Benefits `json:"benefits" description:"Additional benefits or perks such as medical, dental, vision, pension, 401k etc."`
	BenefitsLink         string   `json:"benefits_link" description:"Link to more information about the benefits or perks, if provided"`
	Sponsorship          string   `json:"sponsorship" description:"Sponsorship information, if applicable (e.g., visa sponsorship)"`
}

// JobPosting represents a job posting with various details.
type JobPosting struct {
	Title              string             `json:"title" description:"Job title, not including the company name, location, position type or hours"`
	Description        string             `json:"description" description:"Detailed description of the job or role, must not include company description, mission or history"`
	Location           string             `json:"location" description:"Job location, including remote options such as hybrid or fully remote"`
	Department         string             `json:"department" description:"Department or team"`
	Package            Package            `json:"package" description:"Compensation details, including salary range and currency, if specified"`
	Requirements       []JobRequirement   `json:"requirements" description:"Categorized job requirements"`
	Responsibilities   []string           `json:"responsibilities" description:"List of job responsibilities"`
	PositionType       PositionType       `json:"position_type" description:"The type of position (e.g. Full Time, Part Time, Contract, etc.)"`
	Hours              string             `json:"hours" description:"Working hours or schedule, including any flexibility"`
	PostingDate        string             `json:"posting_date" description:"Date when the job was posted"`
	ApplicationProcess ApplicationProcess `json:"application_process" description:"Details about the application process"`
	Company            Company            `json:"company" description:"Company information"`
}
