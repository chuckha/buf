package buf

import (
	"context"
	"fmt"
	"time"

	"github.com/bufbuild/buf/internal/buf/ext/extio"
	"github.com/bufbuild/buf/internal/pkg/app"
	"github.com/bufbuild/buf/internal/pkg/app/appflag"
	"github.com/bufbuild/buf/internal/pkg/app/applog"
	"github.com/spf13/pflag"
)

const (
	imageBuildInputFlagName            = "source"
	imageBuildConfigFlagName           = "source-config"
	imageBuildOutputFlagName           = "output"
	checkLintInputFlagName             = "input"
	checkLintConfigFlagName            = "input-config"
	checkBreakingInputFlagName         = "input"
	checkBreakingConfigFlagName        = "input-config"
	checkBreakingAgainstInputFlagName  = "against-input"
	checkBreakingAgainstConfigFlagName = "against-input-config"
	checkLsCheckersConfigFlagName      = "config"
	checkLsCheckersFormatFlagName      = "format"
	lsFilesInputFlagName               = "input"
	lsFilesConfigFlagName              = "input-config"
	errorFormatFlagName                = "error-format"
	experimentalGitCloneFlagName       = "experimental-git-clone"
)

// flags are the flags.
type flags struct {
	Config               string
	AgainstConfig        string
	Input                string
	AgainstInput         string
	Output               string
	AsFileDescriptorSet  bool
	ExcludeImports       bool
	ExcludeSourceInfo    bool
	Files                []string
	LimitToInputFiles    bool
	CheckerAll           bool
	CheckerCategories    []string
	ErrorFormat          string
	Format               string
	ExperimentalGitClone bool
}

// newFlags returns a new flags.
func newFlags() *flags {
	return &flags{}
}

// container is a container.
type container struct {
	applog.Container
	*flags
}

// newContainer returns a new container.
func newContainer(logContainer applog.Container, flags *flags) *container {
	return &container{
		Container: logContainer,
		flags:     flags,
	}
}

// builder is a builder.
type builder struct {
	appflag.Builder
	*flags
}

// newBuilder returns a new builder.
func newBuilder() *builder {
	return &builder{
		Builder: appflag.NewBuilder(appflag.BuilderWithTimeout(10 * time.Second)),
		flags:   newFlags(),
	}
}

// newRunFunc creates a new run function.
func (b *builder) newRunFunc(fn func(context.Context, *container) error) func(context.Context, app.Container) error {
	return b.Builder.NewRunFunc(
		func(ctx context.Context, container applog.Container) error {
			return fn(ctx, newContainer(container, b.flags))
		},
	)
}

func (b *builder) bindImageBuildInput(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Input, imageBuildInputFlagName, ".", fmt.Sprintf(`The source to build. Must be one of format %s.`, extio.SourceFormatsToString()))
}

func (b *builder) bindImageBuildConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Config, imageBuildConfigFlagName, "", `The config file or data to use.`)
}

func (b *builder) bindImageBuildOutput(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&b.Output, imageBuildOutputFlagName, "o", "", fmt.Sprintf(`Required. The location to write the image. Must be one of format %s.`, extio.ImageFormatsToString()))
}

func (b *builder) bindImageBuildAsFileDescriptorSet(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.AsFileDescriptorSet, "as-file-descriptor-set", false, `Output as a google.protobuf.FileDescriptorSet instead of an image.

Note that images are wire-compatible with FileDescriptorSets, however this flag will strip
the additional metadata added for Buf usage.`)
}

func (b *builder) bindImageBuildExcludeImports(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.ExcludeImports, "exclude-imports", false, "Exclude imports.")
}

func (b *builder) bindImageBuildExcludeSourceInfo(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.ExcludeSourceInfo, "exclude-source-info", false, "Exclude source info.")
}

func (b *builder) bindImageBuildErrorFormat(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.ErrorFormat, errorFormatFlagName, "text", "The format for build errors, printed to stderr. Must be one of [text,json].")
}

func (b *builder) bindCheckLintInput(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Input, checkLintInputFlagName, ".", fmt.Sprintf(`The source or image to lint. Must be one of format %s.`, extio.InputFormatsToString()))
}

func (b *builder) bindCheckLintConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Config, checkLintConfigFlagName, "", `The config file or data to use.`)
}

func (b *builder) bindCheckBreakingInput(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Input, checkBreakingInputFlagName, ".", fmt.Sprintf(`The source or image to check for breaking changes. Must be one of format %s.`, extio.InputFormatsToString()))
}

func (b *builder) bindCheckBreakingConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Config, checkBreakingConfigFlagName, "", `The config file or data to use.`)
}

func (b *builder) bindCheckBreakingAgainstInput(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.AgainstInput, checkBreakingAgainstInputFlagName, "", fmt.Sprintf(`Required. The source or image to check against. Must be one of format %s.`, extio.InputFormatsToString()))
}

func (b *builder) bindCheckBreakingAgainstConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.AgainstConfig, checkBreakingAgainstConfigFlagName, "", `The config file or data to use for the against source or image.`)
}

func (b *builder) bindCheckBreakingLimitToInputFiles(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.LimitToInputFiles, "limit-to-input-files", false, `Only run breaking checks against the files in the input.
This has the effect of filtering the against input to only contain the files in the input.
Overrides --file.`)
}

func (b *builder) bindCheckBreakingExcludeImports(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.ExcludeImports, "exclude-imports", false, "Exclude imports from breaking change detection.")
}

func (b *builder) bindCheckFiles(flagSet *pflag.FlagSet) {
	flagSet.StringSliceVar(&b.Files, "file", nil, `Limit to specific files. This is an advanced feature and is not recommended.`)
}

func (b *builder) bindCheckBreakingErrorFormat(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.ErrorFormat, errorFormatFlagName, "text", "The format for build errors or check violations, printed to stdout. Must be one of [text,json].")
}

func (b *builder) bindCheckLintErrorFormat(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.ErrorFormat, errorFormatFlagName, "text", "The format for build errors or check violations, printed to stdout. Must be one of [text,json,config-ignore-yaml].")
}

func (b *builder) bindLsFilesInput(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Input, lsFilesInputFlagName, ".", fmt.Sprintf(`The source or image to list the files from. Must be one of format %s.`, extio.InputFormatsToString()))
}

func (b *builder) bindLsFilesConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Config, lsFilesConfigFlagName, "", `The config file or data to use.`)
}

func (b *builder) bindCheckLsCheckersConfig(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Config, checkLsCheckersConfigFlagName, "", `The config file or data to use. If --all is specified, this is ignored.`)
}

func (b *builder) bindCheckLsCheckersAll(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.CheckerAll, "all", false, "List all checkers and not just those currently configured.")
}

func (b *builder) bindCheckLsCheckersCategories(flagSet *pflag.FlagSet) {
	flagSet.StringSliceVar(&b.CheckerCategories, "category", nil, "Only list the checkers in these categories.")
}

func (b *builder) bindCheckLsCheckersFormat(flagSet *pflag.FlagSet) {
	flagSet.StringVar(&b.Format, checkLsCheckersFormatFlagName, "text", "The format to print checkers as. Must be one of [text,json].")
}

func (b *builder) bindExperimentalGitClone(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(&b.ExperimentalGitClone, experimentalGitCloneFlagName, false, "Use the git binary to clone instead of the internal git library.")
}
