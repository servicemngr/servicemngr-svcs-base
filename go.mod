module github.com/servicemngr/servicemngr-svcs-base

go 1.15

replace github.com/digineo/go-uci => github.com/servicemngr/go-uci v0.0.0-20200906220637-4e79be0b1d29

require (
	github.com/digineo/go-uci v0.0.0-20200907075334-cd86efb9cb51
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0
	github.com/servicemngr/core v1.0.0
	github.com/servicemngr/servicemngr-aux v1.0.0
	github.com/spf13/afero v1.5.1
)
