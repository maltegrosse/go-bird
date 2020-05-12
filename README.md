Go-Richard E. Bird - Clear Sky Broadband Solar Radiation Model
=======================================
[![Go Report Card](https://goreportcard.com/badge/github.com/maltegrosse/go-bird)](https://goreportcard.com/report/github.com/maltegrosse/go-bird)
[![GoDoc](https://godoc.org/github.com/maltegrosse/go-bird?status.svg)](https://pkg.go.dev/github.com/maltegrosse/go-bird)
![Go](https://github.com/maltegrosse/go-bird/workflows/Go/badge.svg) 

:warning: **results not validated**

The Bird Clear Sky Model, (original implemented by Afshin Michael Andreas/NREL, adopted in go lang) authored by Richard Bird, is a broadband algorithm that produces estimates of clear sky direct beam, hemispherical diffuse, and total hemispherical solar radiation on a horizontal surface.
The model is based on comparisons with results from rigorous radiative transfer codes. It is composed of simple algebraic expressions with 10 user-provided inputs. Model results should be expected to agree within ±10% with rigorous radiative transfer codes. The model computes hourly average solar radiation for every hour of the year, based on the 10 user input parameters; however, variable atmospheric parameters such as aerosol optical depth, ozone, and water vapor are fixed for the entire year.
## Installation

This packages requires Go 1.13. If you installed it and set up your GOPATH, just run:

`go get -u github.com/maltegrosse/go-bird`

## Usage

You can find some examples in the [examples](examples) directory.

Please visit https://midcdmz.nrel.gov/sampa/ for additional information.

Some additional helper functions have been added to the original application logic.



## License
**[NREL BIRD License](https://midcdmz.nrel.gov/sampa/#license)**

Adoption in Golang under **[MIT license](http://opensource.org/licenses/mit-license.php)** 2020 © Malte Grosse.

