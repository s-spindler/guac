// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// CveGhsaObject is a union of CVE and GHSA.
type CveGhsaObject interface {
	IsCveGhsaObject()
}

// OsvCveGhsaObject is a union of OSV, CVE and GHSA. Any of these objects can be specified for vulnerability
type OsvCveGhsaObject interface {
	IsOsvCveGhsaObject()
}

// PkgArtObject is a union of Package and Artifact. Any of these objects can be specified
type PkgArtObject interface {
	IsPkgArtObject()
}

// PkgSrcArtObject is a union of Package, Source and Artifact. Any of these objects can be specified
type PkgSrcArtObject interface {
	IsPkgSrcArtObject()
}

// PkgSrcObject is a union of Package and Source. Any of these objects can be specified
type PkgSrcObject interface {
	IsPkgSrcObject()
}

// Artifact represents the artifact and contains a digest field
//
// algorithm is mandatory in the from strings.ToLower(string(checksum.Algorithm)) (sha256, sha1...etc)
// digest is mandatory in the form checksum.Value.
type Artifact struct {
	Algorithm string `json:"algorithm"`
	Digest    string `json:"digest"`
}

func (Artifact) IsPkgSrcArtObject() {}

func (Artifact) IsPkgArtObject() {}

// ArtifactSpec allows filtering the list of artifacts to return.
type ArtifactSpec struct {
	Algorithm *string `json:"algorithm"`
	Digest    *string `json:"digest"`
}

// Builder represents the builder such as (FRSCA or github actions) and contains a uri field
//
// uri is mandatory and represents the specific builder.
//
// This node is a singleton: backends guarantee that there is exactly one node with
// the same `uri` value.
type Builder struct {
	URI string `json:"uri"`
}

// BuilderSpec allows filtering the list of builders to return.
type BuilderSpec struct {
	URI *string `json:"uri"`
}

// CVE represents common vulnerabilities and exposures. It contains the year along
// with the CVE ID.
//
// year is mandatory.
//
// This node is a singleton: backends guarantee that there is exactly one node with
// the same `year` value.
type Cve struct {
	Year  string   `json:"year"`
	CveID []*CVEId `json:"cveId"`
}

func (Cve) IsOsvCveGhsaObject() {}

func (Cve) IsCveGhsaObject() {}

// CVEId is the actual ID that is given to a specific vulnerability
//
// id field is mandatory.
//
// This node can be referred to by other parts of GUAC.
type CVEId struct {
	ID string `json:"id"`
}

// CVESpec allows filtering the list of cves to return.
type CVESpec struct {
	Year  *string `json:"year"`
	CveID *string `json:"cveId"`
}

// CertifyBad is an attestation represents when a package, source or artifact is considered bad
//
// subject - union type that can be either a package, source or artifact object type
// justification (property) - string value representing why the subject is considered bad
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
//
// Note: Attestation must occur at the PackageName or the PackageVersion or at the SourceName.
type CertifyBad struct {
	Subject       PkgSrcArtObject `json:"subject"`
	Justification string          `json:"justification"`
	Origin        string          `json:"origin"`
	Collector     string          `json:"collector"`
}

// CertifyBadSpec allows filtering the list of CertifyBad to return.
// Note: Package, Source or artifact must be specified but not at the same time
// For package - a PackageName or PackageVersion must be specified (name or name, version, qualifiers and subpath)
// For source - a SourceName must be specified (name, tag or commit)
type CertifyBadSpec struct {
	Package       *PkgSpec      `json:"package"`
	Source        *SourceSpec   `json:"source"`
	Artifact      *ArtifactSpec `json:"artifact"`
	Justification *string       `json:"justification"`
	Origin        *string       `json:"origin"`
	Collector     *string       `json:"collector"`
}

// CertifyPkg is an attestation that represents when a package objects are similar
//
// packages (subject) - list of package objects
// justification (property) - string value representing why the packages are similar
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type CertifyPkg struct {
	Packages      []*Package `json:"packages"`
	Justification string     `json:"justification"`
	Origin        string     `json:"origin"`
	Collector     string     `json:"collector"`
}

// CertifyPkgSpec allows filtering the list of CertifyPkg to return.
//
// Specifying just the package allows to query for all similar packages (if they exist)
type CertifyPkgSpec struct {
	Packages      []*PkgSpec `json:"packages"`
	Justification *string    `json:"justification"`
	Origin        *string    `json:"origin"`
	Collector     *string    `json:"collector"`
}

// CertifyScorecard is an attestation represents the scorecard of a particular source
//
// source (subject) - the source object type that represents the source
// timeScanned (property) - timestamp when this was last scanned (exact time)
// aggregateScore (property) - overall scorecard score for the source
// checks (property) - individual scorecard check scores (Branch-Protection, Code-Review...etc)
// scorecardVersion (property) - version of the scorecard when the source was scanned
// scorecardCommit (property) - commit of scorecard when the source was scanned
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type CertifyScorecard struct {
	Source           *Source           `json:"source"`
	TimeScanned      string            `json:"timeScanned"`
	AggregateScore   float64           `json:"aggregateScore"`
	Checks           []*ScorecardCheck `json:"checks"`
	ScorecardVersion string            `json:"scorecardVersion"`
	ScorecardCommit  string            `json:"scorecardCommit"`
	Origin           string            `json:"origin"`
	Collector        string            `json:"collector"`
}

// CertifyScorecardSpec allows filtering the list of CertifyScorecard to return.
type CertifyScorecardSpec struct {
	Source           *SourceSpec           `json:"source"`
	TimeScanned      *string               `json:"timeScanned"`
	AggregateScore   *float64              `json:"aggregateScore"`
	Checks           []*ScorecardCheckSpec `json:"checks"`
	ScorecardVersion *string               `json:"scorecardVersion"`
	ScorecardCommit  *string               `json:"scorecardCommit"`
	Origin           *string               `json:"origin"`
	Collector        *string               `json:"collector"`
}

// CertifyVEXStatement is an attestation that represents when a package or artifact has a VEX about a specific vulnerability (CVE or GHSA)
//
// subject - union type that represents a package or artifact
// vulnerability (object) - union type that consists of cve or ghsa
// justification (property) - justification for VEX
// knownSince (property) - timestamp of the VEX (exact time)
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type CertifyVEXStatement struct {
	Subject       PkgArtObject  `json:"subject"`
	Vulnerability CveGhsaObject `json:"vulnerability"`
	Justification string        `json:"justification"`
	KnownSince    string        `json:"knownSince"`
	Origin        string        `json:"origin"`
	Collector     string        `json:"collector"`
}

// CertifyVEXStatementSpec allows filtering the list of CertifyVEXStatement to return.
// Only package or artifact and CVE or GHSA can be specified at once.
type CertifyVEXStatementSpec struct {
	Package       *PkgSpec      `json:"package"`
	Artifact      *ArtifactSpec `json:"artifact"`
	Cve           *CVESpec      `json:"cve"`
	Ghsa          *GHSASpec     `json:"ghsa"`
	Justification *string       `json:"justification"`
	KnownSince    *string       `json:"knownSince"`
	Origin        *string       `json:"origin"`
	Collector     *string       `json:"collector"`
}

// CertifyVuln is an attestation that represents when a package has a vulnerability
//
// package (subject) - the package object type that represents the package
// vulnerability (object) - union type that consists of osv, cve or ghsa
// timeScanned (property) - timestamp of when the package was last scanned
// dbUri (property) - scanner vulnerability database uri
// dbVersion (property) - scanner vulnerability database version
// scannerUri (property) - vulnerability scanner's uri
// scannerVersion (property) - vulnerability scanner version
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type CertifyVuln struct {
	Package        *Package         `json:"package"`
	Vulnerability  OsvCveGhsaObject `json:"vulnerability"`
	TimeScanned    string           `json:"timeScanned"`
	DbURI          string           `json:"dbUri"`
	DbVersion      string           `json:"dbVersion"`
	ScannerURI     string           `json:"scannerUri"`
	ScannerVersion string           `json:"scannerVersion"`
	Origin         string           `json:"origin"`
	Collector      string           `json:"collector"`
}

// CertifyVulnSpec allows filtering the list of CertifyVuln to return.
//
// Specifying just the package allows to query for all vulnerabilities associated with the package.
// Only OSV, CVE or GHSA can be specified at once
type CertifyVulnSpec struct {
	Package        *PkgSpec  `json:"package"`
	Osv            *OSVSpec  `json:"osv"`
	Cve            *CVESpec  `json:"cve"`
	Ghsa           *GHSASpec `json:"ghsa"`
	TimeScanned    *string   `json:"timeScanned"`
	DbURI          *string   `json:"dbUri"`
	DbVersion      *string   `json:"dbVersion"`
	ScannerURI     *string   `json:"scannerUri"`
	ScannerVersion *string   `json:"scannerVersion"`
	Origin         *string   `json:"origin"`
	Collector      *string   `json:"collector"`
}

// GHSA represents github security advisory. It contains the ghsa ID (GHSA-pgvh-p3g4-86jw)
type Ghsa struct {
	GhsaID []*GHSAId `json:"ghsaId"`
}

func (Ghsa) IsOsvCveGhsaObject() {}

func (Ghsa) IsCveGhsaObject() {}

// GHSAId is the actual ID that is given to a specific vulnerability on github
//
// id field is mandatory.
//
// This node can be referred to by other parts of GUAC.
type GHSAId struct {
	ID string `json:"id"`
}

// GHSASpec allows filtering the list of ghsa to return.
type GHSASpec struct {
	GhsaID *string `json:"ghsaId"`
}

// HasSBOM is an attestation represents that a package object or source object has an SBOM associated with a uri
//
// subject - union type that can be either a package or source object type
// uri (property) - identifier string for the SBOM
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
//
// Note: Only package object or source object can be defined. Not both.
type HasSbom struct {
	Subject   PkgSrcObject `json:"subject"`
	URI       string       `json:"uri"`
	Origin    string       `json:"origin"`
	Collector string       `json:"collector"`
}

// HashEqualSpec allows filtering the list of HasSBOM to return.
//
// Only the package or source can be added, not both. HasSourceAt will be used to create the package to source
// relationship.
type HasSBOMSpec struct {
	Package   *PkgSpec    `json:"package"`
	Source    *SourceSpec `json:"source"`
	URI       *string     `json:"uri"`
	Origin    *string     `json:"origin"`
	Collector *string     `json:"collector"`
}

// HasSLSA is an attestation represents that the subject has a SLSA attestation associated with it.
//
// subject - an union type that consists of package, source or artifact
// builtFrom (object) - list of union types that consists of the package, source or artifact that the subject was build from
// builtBy (object) - represents the builder that was used to build the subject
// buildType (property) - individual scorecard check scores (Branch-Protection, Code-Review...etc)
// slsaPredicate (property) - a list of key value pair that consist of the keys and values of the SLSA predicate
// slsaVersion (property) - version of the SLSA predicate
// startedOn (property) - timestamp when the SLSA predicate was recorded during the build time of the subject
// finishedOn (property) - timestamp when the SLSA predicate was completed during the build time of the subject
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type HasSlsa struct {
	Subject       PkgSrcArtObject   `json:"subject"`
	BuiltFrom     []PkgSrcArtObject `json:"builtFrom"`
	BuiltBy       *Builder          `json:"builtBy"`
	BuildType     string            `json:"buildType"`
	SlsaPredicate []*SLSAPredicate  `json:"slsaPredicate"`
	SlsaVersion   string            `json:"slsaVersion"`
	StartedOn     string            `json:"startedOn"`
	FinishedOn    string            `json:"finishedOn"`
	Origin        string            `json:"origin"`
	Collector     string            `json:"collector"`
}

// HasSLSASpec allows filtering the list of HasSLSA to return.
type HasSLSASpec struct {
	Package           *PkgSpec             `json:"package"`
	Source            *SourceSpec          `json:"source"`
	Artifact          *ArtifactSpec        `json:"artifact"`
	BuiltFromPackages []*PkgSpec           `json:"builtFromPackages"`
	BuiltFromSource   []*SourceSpec        `json:"builtFromSource"`
	BuiltFromArtifact []*ArtifactSpec      `json:"builtFromArtifact"`
	BuiltBy           *BuilderSpec         `json:"builtBy"`
	BuildType         *string              `json:"buildType"`
	Predicate         []*SLSAPredicateSpec `json:"predicate"`
	SlsaVersion       *string              `json:"slsaVersion"`
	StartedOn         *string              `json:"startedOn"`
	FinishedOn        *string              `json:"finishedOn"`
	Origin            *string              `json:"origin"`
	Collector         *string              `json:"collector"`
}

// HasSourceAt is an attestation represents that a package object has a source object since a timestamp
//
// package (subject) - the package object type that represents the package
// source (object) - the source object type that represents the source
// knownSince (property) - timestamp when this was last checked (exact time)
// justification (property) - string value representing why the package has a source specified
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type HasSourceAt struct {
	Package       *Package `json:"package"`
	Source        *Source  `json:"source"`
	KnownSince    string   `json:"knownSince"`
	Justification string   `json:"justification"`
	Origin        string   `json:"origin"`
	Collector     string   `json:"collector"`
}

// HasSourceAtSpec allows filtering the list of HasSourceAt to return.
type HasSourceAtSpec struct {
	Package       *PkgSpec    `json:"package"`
	Source        *SourceSpec `json:"source"`
	KnownSince    *string     `json:"knownSince"`
	Justification *string     `json:"justification"`
	Origin        *string     `json:"origin"`
	Collector     *string     `json:"collector"`
}

// HashEqual is an attestation that represents when two artifact hash are similar based on a justification.
//
// artifacts (subject) - the artifacts (represented by algorithm and digest) that are equal
// justification (property) - string value representing why the artifacts are the equal
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type HashEqual struct {
	Artifacts     []*Artifact `json:"artifacts"`
	Justification string      `json:"justification"`
	Origin        string      `json:"origin"`
	Collector     string      `json:"collector"`
}

// HashEqualSpec allows filtering the list of HashEqual to return.
//
// Specifying just the artifacts allows to query for all equivalent artifacts (if they exist)
type HashEqualSpec struct {
	Justification *string         `json:"justification"`
	Artifacts     []*ArtifactSpec `json:"artifacts"`
	Origin        *string         `json:"origin"`
	Collector     *string         `json:"collector"`
}

// IsDependency is an attestation that represents when a package is dependent on another package
//
// package (subject) - the package object type that represents the package
// dependentPackage (object) - the package object type that represents the packageName (cannot be to the packageVersion)
// versionRange (property) - string value for version range that applies to the dependent package
// justification (property) - string value representing why the artifacts are the equal
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type IsDependency struct {
	Package          *Package `json:"package"`
	DependentPackage *Package `json:"dependentPackage"`
	VersionRange     string   `json:"versionRange"`
	Justification    string   `json:"justification"`
	Origin           string   `json:"origin"`
	Collector        string   `json:"collector"`
}

// IsDependencySpec allows filtering the list of IsDependency to return.
//
// Note: the package object must be defined to return its dependent packages.
// Dependent Packages must represent the packageName (cannot be the packageVersion)
type IsDependencySpec struct {
	Package          *PkgSpec     `json:"package"`
	DependentPackage *PkgNameSpec `json:"dependentPackage"`
	VersionRange     *string      `json:"versionRange"`
	Justification    *string      `json:"justification"`
	Origin           *string      `json:"origin"`
	Collector        *string      `json:"collector"`
}

// IsOccurrence is an attestation represents when either a package or source is represented by an artifact
//
// subject - union type that can be either a package or source object type
// occurrenceArtifacts (object) - list of artifacts that represent the the package or source
// justification (property) - string value representing why the package or source is represented by the specified artifact
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
//
// Note: Package or Source must be specified but not both at the same time.
// Attestation must occur at the PackageName or the PackageVersion or at the SourceName.
//
// IsOccurrence does not connect a package with a source.
// HasSourceAt attestation will be used to connect a package with a source
type IsOccurrence struct {
	Subject             PkgSrcObject `json:"subject"`
	OccurrenceArtifacts []*Artifact  `json:"occurrenceArtifacts"`
	Justification       string       `json:"justification"`
	Origin              string       `json:"origin"`
	Collector           string       `json:"collector"`
}

// IsOccurrenceSpec allows filtering the list of IsOccurrence to return.
// Note: Package or Source must be specified but not both at the same time
// For package - a PackageName or PackageVersion must be specified (name or name, version, qualifiers and subpath)
// For source - a SourceName must be specified (name, tag or commit)
type IsOccurrenceSpec struct {
	Justification *string         `json:"justification"`
	Package       *PkgSpec        `json:"package"`
	Source        *SourceSpec     `json:"source"`
	Artifacts     []*ArtifactSpec `json:"artifacts"`
	Origin        *string         `json:"origin"`
	Collector     *string         `json:"collector"`
}

// IsVulnerability is an attestation that represents when an OSV ID represents a CVE or GHSA
//
// osv (subject) - the osv object type that represents OSV and its ID
// vulnerability (object) - union type that consists of cve or ghsa
// justification (property) - the reason why the osv ID represents the cve or ghsa
// origin (property) - where this attestation was generated from (based on which document)
// collector (property) - the GUAC collector that collected the document that generated this attestation
type IsVulnerability struct {
	Osv           *Osv          `json:"osv"`
	Vulnerability CveGhsaObject `json:"vulnerability"`
	Justification string        `json:"justification"`
	Origin        string        `json:"origin"`
	Collector     string        `json:"collector"`
}

// IsVulnerabilitySpec allows filtering the list of IsVulnerability to return.
// Only CVE or GHSA can be specified at once.
type IsVulnerabilitySpec struct {
	Osv           *OSVSpec  `json:"osv"`
	Cve           *CVESpec  `json:"cve"`
	Ghsa          *GHSASpec `json:"ghsa"`
	Justification *string   `json:"justification"`
	Origin        *string   `json:"origin"`
	Collector     *string   `json:"collector"`
}

// OSV represents Open Source Vulnerability . It contains a OSV ID.
type Osv struct {
	OsvID []*OSVId `json:"osvId"`
}

func (Osv) IsOsvCveGhsaObject() {}

// OSVId is the actual ID that is given to a specific vulnerability
//
// id field is mandatory. This maps to a GHSA or CVE ID
//
// This node can be referred to by other parts of GUAC.
type OSVId struct {
	ID string `json:"id"`
}

// OSVSpec allows filtering the list of osv to return.
type OSVSpec struct {
	OsvID *string `json:"osvId"`
}

// Package represents a package.
//
// In the pURL representation, each Package matches a `pkg:<type>` partial pURL.
// The `type` field matches the pURL types but we might also use `"guac"` for the
// cases where the pURL representation is not complete or when we have custom
// rules.
//
// This node is a singleton: backends guarantee that there is exactly one node
// with the same `type` value.
//
// Also note that this is named `Package`, not `PackageType`. This is only to make
// queries more readable.
type Package struct {
	Type       string              `json:"type"`
	Namespaces []*PackageNamespace `json:"namespaces"`
}

func (Package) IsPkgSrcArtObject() {}

func (Package) IsPkgArtObject() {}

func (Package) IsPkgSrcObject() {}

// PackageName is a name for packages.
//
// In the pURL representation, each PackageName matches the
// `pkg:<type>/<namespace>/<name>` pURL.
//
// Names are always mandatory.
//
// This is the first node in the trie that can be referred to by other parts of
// GUAC.
type PackageName struct {
	Name     string            `json:"name"`
	Versions []*PackageVersion `json:"versions"`
}

// PackageNamespace is a namespace for packages.
//
// In the pURL representation, each PackageNamespace matches the
// `pkg:<type>/<namespace>/` partial pURL.
//
// Namespaces are optional and type specific. Because they are optional, we use
// empty string to denote missing namespaces.
type PackageNamespace struct {
	Namespace string         `json:"namespace"`
	Names     []*PackageName `json:"names"`
}

// PackageQualifier is a qualifier for a package, a key-value pair.
//
// In the pURL representation, it is a part of the `<qualifiers>` part of the
// `pkg:<type>/<namespace>/<name>@<version>?<qualifiers>` pURL.
//
// Qualifiers are optional, each Package type defines own rules for handling them,
// and multiple qualifiers could be attached to the same package.
//
// This node cannot be directly referred by other parts of GUAC.
type PackageQualifier struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PackageQualifierInputSpec is the same as PackageQualifier, but usable as
// mutation input.
//
// GraphQL does not allow input types to contain composite types and does not allow
// composite types to contain input types. So, although in this case these two
// types are semantically the same, we have to duplicate the definition.
//
// Both fields are mandatory.
type PackageQualifierInputSpec struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PackageQualifierSpec is the same as PackageQualifier, but usable as query
// input.
//
// GraphQL does not allow input types to contain composite types and does not allow
// composite types to contain input types. So, although in this case these two
// types are semantically the same, we have to duplicate the definition.
//
// Keys are mandatory, but values could also be `null` if we want to match all
// values for a specific key.
//
// TODO(mihaimaruseac): Formalize empty vs null when the schema is fully done
type PackageQualifierSpec struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}

// PackageVersion is a package version.
//
// In the pURL representation, each PackageName matches the
// `pkg:<type>/<namespace>/<name>@<version>` pURL.
//
// Versions are optional and each Package type defines own rules for handling them.
// For this level of GUAC, these are just opaque strings.
//
// This node can be referred to by other parts of GUAC.
//
// Subpath and qualifiers are optional. Lack of qualifiers is represented by an
// empty list and lack of subpath by empty string (to be consistent with
// optionality of namespace and version). Two nodes that have different qualifiers
// and/or subpath but the same version mean two different packages in the trie
// (they are different). Two nodes that have same version but qualifiers of one are
// a subset of the qualifier of the other also mean two different packages in the
// trie.
type PackageVersion struct {
	Version    string              `json:"version"`
	Qualifiers []*PackageQualifier `json:"qualifiers"`
	Subpath    string              `json:"subpath"`
}

// PkgInputSpec specifies a package for a mutation.
//
// This is different than PkgSpec because we want to encode mandatatory fields:
// `type` and `name`. All optional fields are given empty default values.
type PkgInputSpec struct {
	Type       string                       `json:"type"`
	Namespace  *string                      `json:"namespace"`
	Name       string                       `json:"name"`
	Version    *string                      `json:"version"`
	Qualifiers []*PackageQualifierInputSpec `json:"qualifiers"`
	Subpath    *string                      `json:"subpath"`
}

// PkgNameSpec is used for IsDependency to input dependent packages. This is different from PkgSpec
// as the IsDependency attestation should only be allowed to be made to the packageName node and not the
// packageVersion node. Versions will be handled by the version_range in the IsDependency attestation node.
type PkgNameSpec struct {
	Type      *string `json:"type"`
	Namespace *string `json:"namespace"`
	Name      *string `json:"name"`
}

// PkgSpec allows filtering the list of packages to return.
//
// Each field matches a qualifier from pURL. Use `null` to match on all values at
// that level. For example, to get all packages in GUAC backend, use a PkgSpec
// where every field is `null`.
//
// Empty string at a field means matching with the empty string. If passing in
// qualifiers, all of the values in the list must match. Since we want to return
// nodes with any number of qualifiers if no qualifiers are passed in the input, we
// must also return the same set of nodes it the qualifiers list is empty. To match
// on nodes that don't contain any qualifier, set `matchOnlyEmptyQualifiers` to
// true. If this field is true, then the qualifiers argument is ignored.
type PkgSpec struct {
	Type                     *string                 `json:"type"`
	Namespace                *string                 `json:"namespace"`
	Name                     *string                 `json:"name"`
	Version                  *string                 `json:"version"`
	Qualifiers               []*PackageQualifierSpec `json:"qualifiers"`
	MatchOnlyEmptyQualifiers *bool                   `json:"matchOnlyEmptyQualifiers"`
	Subpath                  *string                 `json:"subpath"`
}

// SLSAPredicate are the values from the SLSA predicate in key-value pair form.
// // Predicate:
// "predicateType": "https://slsa.dev/provenance/v1",
//
//	"predicate": {
//	    "buildDefinition": {
//	        "buildType": string,
//	        "externalParameters": object,
//	        "systemParameters": object,
//	        "resolvedDependencies": [ ...#ArtifactReference ],
//	    },
//	    "runDetails": {
//	        "builder": {
//	            "id": string,
//	            "version": string,
//	            "builderDependencies": [ ...#ArtifactReference ],
//	        },
//	        "metadata": {
//	            "invocationId": string,
//	            "startedOn": #Timestamp,
//	            "finishedOn": #Timestamp,
//	        },
//	        "byproducts": [ ...#ArtifactReference ],
//	    }
//	}
//
// where
//
//	"externalParameters": {
//	    "repository": "https://github.com/octocat/hello-world",
//	    "ref": "refs/heads/main"
//	},
//
// For example: key = "buildDefinition.externalParameters.repository" value = "https://github.com/octocat/hello-world"
// This node cannot be directly referred by other parts of GUAC.
type SLSAPredicate struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SLSAPredicateSpec is the same as SLSAPredicateSpec, but usable as query
// input.
type SLSAPredicateSpec struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ScorecardCheck are the individual checks from scorecard and their values, a key-value pair.
// For example:  Branch-Protection, Code-Review...etc
// Based off scorecard's:
//
//	type jsonCheckResultV2 struct {
//		Details []string                 `json:"details"`
//		Score   int                      `json:"score"`
//		Reason  string                   `json:"reason"`
//		Name    string                   `json:"name"`
//		Doc     jsonCheckDocumentationV2 `json:"documentation"`
//	}
//
// This node cannot be directly referred by other parts of GUAC.
type ScorecardCheck struct {
	Check string `json:"check"`
	Score int    `json:"score"`
}

// ScorecardCheckSpec is the same as ScorecardCheck, but usable as query
// input.
type ScorecardCheckSpec struct {
	Check string `json:"check"`
	Score int    `json:"score"`
}

// Source represents a source.
//
// This can be the version control system that is being used.
//
// This node is a singleton: backends guarantee that there is exactly one node
// with the same `type` value.
//
// Also note that this is named `Source`, not `SourceType`. This is only to make
// queries more readable.
type Source struct {
	Type       string             `json:"type"`
	Namespaces []*SourceNamespace `json:"namespaces"`
}

func (Source) IsPkgSrcArtObject() {}

func (Source) IsPkgSrcObject() {}

// SourceName is a url of the repository and its tag or commit.
//
// The `name` field is mandatory. The `tag` and `commit` fields are optional, but
// it is an error to specify both.
//
// This is the only source trie node that can be referenced by other parts of
// GUAC.
type SourceName struct {
	Name   string  `json:"name"`
	Tag    *string `json:"tag"`
	Commit *string `json:"commit"`
}

// SourceNamespace is a namespace for sources.
//
// This is the location of the repository (such as github/gitlab/bitbucket).
//
// The `namespace` field is mandatory.
type SourceNamespace struct {
	Namespace string        `json:"namespace"`
	Names     []*SourceName `json:"names"`
}

// SourceSpec allows filtering the list of sources to return.
//
// Empty string at a field means matching with the empty string. Missing field
// means retrieving all possible matches.
//
// It is an error to specify both tag and commit fields, except it both are set as
// empty string (in which case the returned sources are only those for which there
// is no tag/commit information).
type SourceSpec struct {
	Type      *string `json:"type"`
	Namespace *string `json:"namespace"`
	Name      *string `json:"name"`
	Tag       *string `json:"tag"`
	Commit    *string `json:"commit"`
}
