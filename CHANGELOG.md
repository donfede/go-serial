# Changelog

## TIP

* Bugfix: `serial.PortError` methods now consistently using a pointer to the error struct (`*PortError`) as reciever, as
  of only pointer to error originally retured by `serial.Port` methods.
* Type `serial.PortError` is now compatible with go 1.13 errors models
    * Method `serial.PortError.Unwrap()` added.
    * Method `serial.PortError.Caused()` marked as deprected.
* Dependencies updated, re-generate syscalls code.
* `github.com/albenik/go-serial/enumerator` moved to `github.com/albenik-go/serial-enumerator`.
    * Current implementation of `enumerator` marked as deprecated and will not recieve updates.

## v2.4.0

* Bugfix: Fixed regress for `darwin` introduced in v2.3.0 (reported [xxx]())
* Add `darwin/arm64` support (by Martin Hebnes Pedersen).

## v2.3.0

* Backport changes from thttps://github.com/bugst/go-serial.

## v2.2.0

* Add `serial.PortError.Cause()` method (by [danielhongwoo](https://github.com/danielhongwoo)).

## v2.1.0

* MacOS: Extended baudrate support (by [smihica](https://github.com/smihica)).
* Internal code updated.

## v2.0.0

* New Go Module import path `github.com/albenik/go-serial/v2`
* `serial.Port` interface discared in favor of `serial.Port` structure (similar to `os.File`)
* `serial.Mode` discared and replaced with `serial.Option`
* `serial.Open()` method changed to use `serila.Option`)
* `port.SetMode(mode *Mode)` replaced with `port.Reconfigure(opts ...Option)`
* `Disable HUPCL by default` (see https://github.com/albenik/go-serial/pull/7)
* `WithHUPCL(bool)` option introduced
* Minor bugfix & refactoring

## v1.x.x

* Forked from https://github.com/bugst/go-serial
* Minor but incompatible interface & logic changes implemented
* Import path altered
