// Package custommovementsensor implements a movementsensor where all methods are unimplemented.
// It extends the built-in resource subtype movementsensor and implements methods to handle resource construction and attribute configuration.

package custommovementsensor

import (
    "context"
    "errors"

    "go.viam.com/rdk/components/movementsensor"
    "go.viam.com/rdk/logging"
    "go.viam.com/rdk/resource"

    "go.viam.com/utils"
)

// Here is where we define your new model's colon-delimited-triplet (viam-labs:go-module-templates-movementsensor:custommovementsensor)
// viam-labs = namespace, go-module-templates-movementsensor = repo-name, custommovementsensor = model name.
// TODO: Change model namespace, family (often the repo-name), and model. For more information see https://docs.viam.com/registry/create/#name-your-new-resource-model
var (
    Model            = resource.NewModel("viam-labs", "go-module-templates-movementsensor", "custommovementsensor")
    errUnimplemented = errors.New("unimplemented")
)

func init() {
    resource.RegisterComponent(movementsensor.API, Model,
        resource.Registration[movementsensor.PowerSensor, *Config]{
            Constructor: newCustomPowerSensor,
        },
    )
}

// TODO: Change the Config struct to contain any values that you would like to be able to configure from the attributes field in the component
// configuration. For more information see https://docs.viam.com/build/configure/#components
type Config struct {
    ArgumentOne int    `json:"one"`
    ArgumentTwo string `json:"two"`
}

// Validate validates the config and returns implicit dependencies.
// TODO: Change the Validate function to validate any config variables.
func (cfg *Config) Validate(path string) ([]string, error) {
    if cfg.ArgumentOne == 0 {
        return nil, utils.NewConfigValidationFieldRequiredError(path, "one")
    }

    if cfg.ArgumentTwo == "" {
        return nil, utils.NewConfigValidationFieldRequiredError(path, "two")
    }

    // TODO: return implicit dependencies if needed as the first value
    return []string{}, nil
}

// Constructor for a custom movementsensor that creates and returns a customPowerSensor.
// TODO: update the customPowerSensor struct and the initialization.
func newCustomPowerSensor(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (movementsensor.PowerSensor, error) {
    // This takes the generic resource.Config passed down from the parent and converts it to the
    // model-specific (aka "native") Config structure defined above, making it easier to directly access attributes.
    conf, err := resource.NativeConfig[*Config](rawConf)
    if err != nil {
        return nil, err
    }

    // Create a cancelable context for custom movementsensor
    cancelCtx, cancelFunc := context.WithCancel(context.Background())

    s := &customPowerSensor{
        name:        rawConf.ResourceName(),
        logger:      logger,
        cfg:         conf,
        cancelCtx:   cancelCtx,
        cancelFunc:  cancelFunc,
    }

    // TODO: If your custom component has dependencies, perform any checks you need to on them.

    // The Reconfigure() method changes the values on the customPowerSensor based on the attributes in the component config
    if err := s.Reconfigure(ctx, deps, rawConf); err != nil {
        logger.Error("Error configuring module with ", rawConf)
        return nil, err
    }

    return s, nil
}

// TODO: update the customPowerSensor struct with any fields you require.
type customPowerSensor struct {
    name   resource.Name
    logger logging.Logger
    cfg    *Config

    cancelCtx  context.Context
    cancelFunc func()

    argumentOne int
    argumentTwo string
}

func (s *customPowerSensor) Name() resource.Name {
    return s.name
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the movementsensor, throw a NewMustBuildError.
func (s *customPowerSensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
    movementsensorConfig, err := resource.NativeConfig[*Config](conf)
    if err != nil {
        s.logger.Warn("Error reconfiguring module with ", err)
        return err
    }

    s.argumentOne = movementsensorConfig.ArgumentOne
    s.argumentTwo = movementsensorConfig.ArgumentTwo
    s.name = conf.ResourceName()
    s.logger.Info("one is now configured to: ", s.argumentOne)
    s.logger.Info("two is now configured to ", s.argumentTwo)

    return nil
}

func (s *customPowerSensor) Voltage(ctx context.Context, extra map[string]interface{}) (float64, bool, error) {
    // TODO: Obtain and return readings.
    s.logger.Error("Method unimplemented")
    return 0, false, errUnimplemented
}

func (s *customPowerSensor) Current(ctx context.Context, extra map[string]interface{}) (float64, bool, error) {
    // TODO: Obtain and return readings.
    s.logger.Error("Method unimplemented")
    return 0, false, errUnimplemented
}

func (s *customPowerSensor) Power(ctx context.Context, extra map[string]interface{}) (float64, error) {
    // TODO: Obtain and return readings.
    s.logger.Error("Method unimplemented")
    return 0, errUnimplemented
}

func (s *customPowerSensor) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
    // TODO: Obtain and return readings.
    s.logger.Error("Method unimplemented")
    return nil, errUnimplemented
}

// DoCommand is a place to add additional commands to extend the movementsensor API. This is optional.
func (s *customPowerSensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
    s.logger.Error("Method unimplemented")
    return nil, errUnimplemented
}

// Close closes the underlying generic.
func (s *customPowerSensor) Close(ctx context.Context) error {
    s.cancelFunc()
    return nil
}
