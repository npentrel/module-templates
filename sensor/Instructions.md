## Create your own sensor module

Viam provides built-in support for a variety of different components and services, but you can add support for your own custom resources by creating a module.

This is a module template that implements one sensor. You can use it as a starting point to write a module that adds support for the sensor you need.

### Instructions

#### 1. Copy this folder into your workspace

To begin copy the [`sensor`](../) folder into another folder on your machine and initialize git. You can rename the folder if you wish to describe your custom sensor better.

Inside your copied folder, you have the following files:

- [cmd](./cmd/)
  - [module](./module/)
    - [cmd.go](./cmd.go): The entry point main program file, which imports the API for the model, defines the `main()` function that registers the model with `viam-server` and creates and starts the module.
- [.gitignore](./.gitignore): A sensible default `.gitignore` file that ensures that no binaries are added to git.
- [customsensor.go](./customsensor.go): A template for a custom sensor implementation.
- [go.mod](./go.mod): Your go dependencies. For more information see [the Go documenation](https://go.dev/doc/tutorial/create-module).
- [go.sum](./go.sum): Checksums for the exact contents of each dependency at the time it is added to your module. For more information see [the Go documentation](https://go.dev/doc/tutorial/create-module).
- [Makefile](./Makefile): A default makefile with command `make module`, `make bin/customsensor`, `make lint` and more.
- [README.md](./README.md): A readme template for your module.

#### 2. Update `cmd/module/cmd.go`

Update the entry point file where there are TODOs. Each TODO explains what you need to do. Your updates will change the name `customsensor` to something more descriptive for the custom sensor you are implementing, for example `temperaturesensor`.

#### 3. Update `cmd/module/customsensor.go`

Update the custom sensor implementation where there are TODOs. Each TODO explains what you need to do. Your updates will:

- Give a name to the model. For more information, see the [docs on model naming](https://docs.viam.com/registry/create/#name-your-new-resource-model).
- Add any configurable values for the sensor's attributes field. For more information, see the [docs on component configuration](https://docs.viam.com/build/configure/#components). Additionally you will update the `Validate()` function to validate that config values meet requirements, for example that required fields exist.
- Update the constructor for your sensor as well as the `customSensor` struct.
- Update the implementation of the `Readings()` and optionally the `DoCommand` functions.

#### 5. Compile and build

If you renamed your package and are building locally you need to specify that your package name should be replaced with the local version of the package:

```
go mod edit -replace github.com/viam-labs/module-templates/sensor/customsensor=./customsensor
```

Then run the following commands to update the `rdk` to the latest build and synchronize and update the project's go dependencies:

```sh
make updaterdk
```

Next, run `make bin/customsensor` to build the executable.

#### 6. Test your module locally

Follow the instructions to [Add a local module](https://docs.viam.com/registry/configure/#add-a-local-module).

#### 7. Upload your module to the registry

To upload the module to the registry, which will allow you to add it to machines without manually moving the files to each machine and specifying the file path, follow the instructions to [Upload your Own Modules to the Viam Registry](https://docs.viam.com/registry/upload/).

#### 8. Update your Readme

Update the [README.md](./README.md) file with the configuration instructions for your module.