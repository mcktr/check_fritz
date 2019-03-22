# Changelog

## 1.0 (2019-03-22)

* First stable version of check_fritz

* [#38](https://github.com/mcktr/check_fritz/pull/38) (documentation): Add installation documentation
* [#37](https://github.com/mcktr/check_fritz/pull/37) (enhancement): Rename check methods to their final names
* [#36](https://github.com/mcktr/check_fritz/pull/36) (enhancement): Add --version command-line option
* [#35](https://github.com/mcktr/check_fritz/pull/35) (build): Provide SHA-256 checksum
* [#34](https://github.com/mcktr/check_fritz/pull/34) (build, tests): Add Go 1.12 to the CI tests
* [#33](https://github.com/mcktr/check_fritz/pull/33) (tests): Fix golint import path in CI tests
* [#31](https://github.com/mcktr/check_fritz/pull/31) (bug): Make password optional
* [#30](https://github.com/mcktr/check_fritz/pull/30) (bug): Fix wrong performance data value in socketpower
* [#29](https://github.com/mcktr/check_fritz/pull/29) (enhancement): Calculate human readable uptime
* [#28](https://github.com/mcktr/check_fritz/pull/28) (bug): Fix wrong new line statement in output
* [#27](https://github.com/mcktr/check_fritz/pull/27) (bug): Fix performance data output in downstream_max
* [#26](https://github.com/mcktr/check_fritz/pull/26) (bug): Fix null pointer reference
* [#24](https://github.com/mcktr/check_fritz/pull/24) (bug): Make Golint happy
* [#23](https://github.com/mcktr/check_fritz/pull/23) (enhancement): Add Travis CI integration
* [#22](https://github.com/mcktr/check_fritz/pull/22) (code-quality): Use go-cmdline library
* [#21](https://github.com/mcktr/check_fritz/pull/21) (code-quality): Make Golint happier
* [#20](https://github.com/mcktr/check_fritz/pull/20) (enhancement): Implement performance data output
* [#18](https://github.com/mcktr/check_fritz/pull/18) (bug): Fix wrong error message in CheckSmartThermometer
* [#17](https://github.com/mcktr/check_fritz/pull/17) (enhancement): Implement two check methods for smart sockets
* [#16](https://github.com/mcktr/check_fritz/pull/16) (documentation): Doc: Update parameter and methods
* [#15](https://github.com/mcktr/check_fritz/pull/15) (enhancement): AImplement check method for smart thermostats
* [#14](https://github.com/mcktr/check_fritz/pull/14) (bug): Set default thresholds to -1 and ignore them
* [#12](https://github.com/mcktr/check_fritz/pull/12) (code-quality): Refactor some variable names to make golint happy
* [#11](https://github.com/mcktr/check_fritz/pull/11) (code-quality): Add unit tests for package thresholds
* [#10](https://github.com/mcktr/check_fritz/pull/10) (enhancement): Implement interface_update method
* [#9](https://github.com/mcktr/check_fritz/pull/9) (enhancement): Implement thresholds
* [#8](https://github.com/mcktr/check_fritz/pull/8) (enhancement): Implement more methods
* [#7](https://github.com/mcktr/check_fritz/pull/7) (code-quality): Use naming convention for all files
* [#6](https://github.com/mcktr/check_fritz/pull/6) (enhancement): Add more check methods