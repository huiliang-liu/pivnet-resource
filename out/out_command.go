package out

import (
	"fmt"
	"os"
	"path"

	pivnet "github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/go-pivnet/v2/logger"
	"github.com/pivotal-cf/pivnet-resource/concourse"
	"github.com/pivotal-cf/pivnet-resource/metadata"
)

type OutCommand struct {
	logger                       logger.Logger
	outDir                       string
	sourcesDir                   string
	globClient                   globber
	validation                   validation
	creator                      creator
	userGroupsUpdater            userGroupsUpdater
	releaseFileGroupsAdder       releaseFileGroupsAdder
	releaseDependenciesAdder     releaseDependenciesAdder
	dependencySpecifiersCreator  dependencySpecifiersCreator
	releaseUpgradePathsAdder     releaseUpgradePathsAdder
	upgradePathSpecifiersCreator upgradePathSpecifiersCreator
	finalizer                    finalizer
	uploader                     uploader
	m                            metadata.Metadata
	skipUpload                   bool
}

type OutCommandConfig struct {
	Logger                       logger.Logger
	OutDir                       string
	SourcesDir                   string
	GlobClient                   globber
	Validation                   validation
	Creator                      creator
	UserGroupsUpdater            userGroupsUpdater
	ReleaseFileGroupsAdder       releaseFileGroupsAdder
	ReleaseDependenciesAdder     releaseDependenciesAdder
	DependencySpecifiersCreator  dependencySpecifiersCreator
	ReleaseUpgradePathsAdder     releaseUpgradePathsAdder
	UpgradePathSpecifiersCreator upgradePathSpecifiersCreator
	Finalizer                    finalizer
	Uploader                     uploader
	M                            metadata.Metadata
	SkipUpload                   bool
}

func NewOutCommand(config OutCommandConfig) OutCommand {
	return OutCommand{
		logger:                       config.Logger,
		outDir:                       config.OutDir,
		sourcesDir:                   config.SourcesDir,
		globClient:                   config.GlobClient,
		validation:                   config.Validation,
		creator:                      config.Creator,
		userGroupsUpdater:            config.UserGroupsUpdater,
		releaseFileGroupsAdder:       config.ReleaseFileGroupsAdder,
		releaseDependenciesAdder:     config.ReleaseDependenciesAdder,
		dependencySpecifiersCreator:  config.DependencySpecifiersCreator,
		releaseUpgradePathsAdder:     config.ReleaseUpgradePathsAdder,
		upgradePathSpecifiersCreator: config.UpgradePathSpecifiersCreator,
		finalizer:                    config.Finalizer,
		uploader:                     config.Uploader,
		m:                            config.M,
		skipUpload:                   config.SkipUpload,
	}
}

//go:generate counterfeiter --fake-name Creator . creator
type creator interface {
	Create() (pivnet.Release, error)
}

//go:generate counterfeiter --fake-name Uploader . uploader
type uploader interface {
	Upload(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name UserGroupsUpdater . userGroupsUpdater
type userGroupsUpdater interface {
	UpdateUserGroups(release pivnet.Release) (pivnet.Release, error)
}

//go:generate counterfeiter --fake-name ReleaseFileGroupsAdder . releaseFileGroupsAdder
type releaseFileGroupsAdder interface {
	AddReleaseFileGroups(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name ReleaseDependenciesAdder . releaseDependenciesAdder
type releaseDependenciesAdder interface {
	AddReleaseDependencies(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name DependencySpecifiersCreator . dependencySpecifiersCreator
type dependencySpecifiersCreator interface {
	CreateDependencySpecifiers(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name ReleaseUpgradePathsAdder . releaseUpgradePathsAdder
type releaseUpgradePathsAdder interface {
	AddReleaseUpgradePaths(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name UpgradePathSpecifiersCreator . upgradePathSpecifiersCreator
type upgradePathSpecifiersCreator interface {
	CreateUpgradePathSpecifiers(release pivnet.Release) error
}

//go:generate counterfeiter --fake-name Finalizer . finalizer
type finalizer interface {
	Finalize(productSlug string, releaseVersion string) (concourse.OutResponse, error)
}

//go:generate counterfeiter --fake-name Validation . validation
type validation interface {
	Validate() error
}

//go:generate counterfeiter --fake-name Globber . globber
type globber interface {
	ExactGlobs() ([]string, error)
}

func (c OutCommand) findMissingFile(files []metadata.ProductFile) []string {
	var missingFiles []string
	for _, f := range files {

		_, err := os.Stat(path.Join(c.sourcesDir, f.File))
		if os.IsNotExist(err) {
			missingFiles = append(missingFiles, f.File)
		}
	}
	return missingFiles
}

func (c OutCommand) Run(input concourse.OutRequest) (concourse.OutResponse, error) {
	if c.outDir == "" {
		return concourse.OutResponse{}, fmt.Errorf("out dir must be provided")
	}

	err := c.validation.Validate()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	missingFiles := c.findMissingFile(c.m.ProductFiles)

	for _, g := range c.m.FileGroups {
		mf := c.findMissingFile(g.ProductFiles)
		missingFiles = append(missingFiles, mf...)
	}

	if len(missingFiles) > 0 {
		return concourse.OutResponse{},
			fmt.Errorf(
				"product files were provided in metadata that match no globs: %v",
				missingFiles,
			)
	}

	pivnetRelease, err := c.creator.Create()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	if c.skipUpload {
		c.logger.Info(
			"file glob not provided - skipping upload to s3")
	} else {
		err = c.uploader.Upload(pivnetRelease)
		if err != nil {
			return concourse.OutResponse{}, err
		}
	}

	err = c.releaseFileGroupsAdder.AddReleaseFileGroups(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	err = c.releaseUpgradePathsAdder.AddReleaseUpgradePaths(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	err = c.releaseDependenciesAdder.AddReleaseDependencies(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	err = c.upgradePathSpecifiersCreator.CreateUpgradePathSpecifiers(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	err = c.dependencySpecifiersCreator.CreateDependencySpecifiers(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	pivnetRelease, err = c.userGroupsUpdater.UpdateUserGroups(pivnetRelease)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	out, err := c.finalizer.Finalize(input.Source.ProductSlug, pivnetRelease.Version)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	c.logger.Info("Put complete")

	return out, nil
}
