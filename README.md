# noaa-exporter
[![Golangci-lint](https://github.com/TawR1024/noaa-exporter/actions/workflows/golangci-linter.yaml/badge.svg?branch=master)](https://github.com/TawR1024/noaa-exporter/actions/workflows/golangci-linter.yaml)
[![build](https://github.com/TawR1024/noaa-exporter/actions/workflows/build-binary.yaml/badge.svg?branch=master)](https://github.com/TawR1024/noaa-exporter/actions/workflows/build-binary.yaml)


This app provides simple way to scrape solar wind parameters from [NOAA](https://www.noaa.gov/) services.
noaa-exporter gets data from NOAA and store it to victoria-metrics database.

already implemented functionality:
* get daily solar-wind-parameters
