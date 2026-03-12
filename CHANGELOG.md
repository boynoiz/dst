# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [unreleased]
### Features

- [`#3a21510`](https://github.com/boynoiz/dst/commit/3a21510670bb8450083fc1d55753a819357a351e) Support Go 1.24 generic type aliases (TypeSpec.Assign ordering)
- [`#26becc9`](https://github.com/boynoiz/dst/commit/26becc99afcfe721dde3b6f1e32675f0da5b313a) Add GoVersion field to dst.File (Go 1.21+)

### Bug Fixes

- [`#2ddad76`](https://github.com/boynoiz/dst/commit/2ddad76be0f05637a078ba360a9f276507d1535a) Return `os.FileInfo` in resolver test for compatibility
- [`#260707f`](https://github.com/boynoiz/dst/commit/260707ff69d06a9811ec53a069ca7266094ee622) Restore SliceExpr fixture in positions.go
- [`#f01cd06`](https://github.com/boynoiz/dst/commit/f01cd069706027198969566af928b9d6a14b9e26) *(decorator)* Only update space type if new value is larger
- [`#8096f70`](https://github.com/boynoiz/dst/commit/8096f70d1385218d39b39a7889f1afbf1e2d7650) Resolve struct literal field order and missing errors imports

### Refactor

- [`#1f0da69`](https://github.com/boynoiz/dst/commit/1f0da696b7ca2068de04b86065be4a71877463ad) Simplify `ValueSpec` declarations and clean up `SliceExpr` fixtures
- [`#881f50e`](https://github.com/boynoiz/dst/commit/881f50e278eec1b9d8a37e5640cb7a0935f88d13) Align permissions format, clean up unused `return` statements, and migrate from `ioutil` to `os` APIs

### Testing

- [`#62da692`](https://github.com/boynoiz/dst/commit/62da692730fc22b6336487a6f3115dc0635c5430) Skip known upstream bugs in TestPositions and TestLoadStdLibAll

### Miscellaneous Tasks

- [`#3a30d7d`](https://github.com/boynoiz/dst/commit/3a30d7d1163734cf36765c5327a7c17edc907f1f) Rename module path from dave/dst to boynoiz/dst

## [0.27.3] - 2022-12-09
### Other

- [`#5fa8d6e`](https://github.com/boynoiz/dst/commit/5fa8d6ebe49a6b04afa15ee8f982d210f8a00b80) *(other)* Fixes #69

## [0.27.2] - 2022-10-13
### Other

- [`#bc71a76`](https://github.com/boynoiz/dst/commit/bc71a763e6c47fced1dade1267b78666d1251843) *(other)* Upgrade golang.org/x/tools to v0.1.12

## [0.27.1] - 2022-09-26
### Other

- [`#5bf3201`](https://github.com/boynoiz/dst/commit/5bf3201c491bdf49e216794604240b2ce8e51197) *(other)* Merge pull request #70 from stevekuznetsov/skuznets/better-error
- [`#b6f3447`](https://github.com/boynoiz/dst/commit/b6f34475d92af6ee92fc2068a53b3302ea113005) *(other)* Make an error more verbose

## [0.27.0] - 2022-05-29
### Other

- [`#fc8c7a2`](https://github.com/boynoiz/dst/commit/fc8c7a23962d90f905c04d344ab95d39e178d8e3) *(other)* Fixed link in readme
- [`#b312ee9`](https://github.com/boynoiz/dst/commit/b312ee95cb643c24ed9687d545325b4cb83be8cc) *(other)* Tweaked contributing section and moved to separate readme
- [`#2f51957`](https://github.com/boynoiz/dst/commit/2f519576342f0e6c1579beec1ed57770c797a290) *(other)* Add building/developing/contributing documentation
- [`#fd0f1be`](https://github.com/boynoiz/dst/commit/fd0f1bedaf52df04501dbab6db5f242554755a12) *(other)* Added tests
- [`#9649d2a`](https://github.com/boynoiz/dst/commit/9649d2a02fb229e2b691a14f3ac3ff637d518599) *(other)* Fixed bugs in rewrite.go and walk.go
- [`#918a64e`](https://github.com/boynoiz/dst/commit/918a64e5bb46e5adf2b0ac89e869b8b192ddd350) *(other)* Improved tests that check positions.go for completeness. Simplified some examples in positions.go.
- [`#cd0f424`](https://github.com/boynoiz/dst/commit/cd0f42470589925f01324eca49cc63d6872b0682) *(other)* Improve tests to fail when no packages / source is found
- [`#e0c1e01`](https://github.com/boynoiz/dst/commit/e0c1e013d3b1a1f1e53d41ffa3150a3c97a753c0) *(other)* Updates to dst.go to get up to date with ast.go 0605bf6052807e71e52fc3864b18b221ce61b047
- [`#43ee57f`](https://github.com/boynoiz/dst/commit/43ee57f6bda4a2c66afd37501fe32f7940becd90) *(other)* Add support for Generics (requires Go 1.18+)

## [0.26.2] - 2020-10-21
### Other

- [`#4f484ea`](https://github.com/boynoiz/dst/commit/4f484ea8fc7f5a6532bd5b227389d369eff6a332) *(other)* Fix for #53
- [`#44b2a2d`](https://github.com/boynoiz/dst/commit/44b2a2dfb1fc3e75527908a2eea488c10b19e05f) *(other)* Badge
- [`#ff4555c`](https://github.com/boynoiz/dst/commit/ff4555c0d498cf60ddc3c7a27440e61c141b23dc) *(other)* Badge

## [0.26.1] - 2020-10-21
### Other

- [`#aa5dfbf`](https://github.com/boynoiz/dst/commit/aa5dfbfba0140ffae85fb711456a99bf80f19d6e) *(other)* Merge pull request #55 from jpap/fix-missing-functype-mapping
- [`#e19822a`](https://github.com/boynoiz/dst/commit/e19822ac991006fab13f2e0507590cd87318413c) *(other)* Add missing FuncType ast<>dst mapping.

## [0.26.0] - 2020-08-23
### Other

- [`#e325bff`](https://github.com/boynoiz/dst/commit/e325bff7482d31df9947eccecee533552567b351) *(other)* Exclude crypto/x509/x509.go from TestLoadStdLibAll and update comment. See https://github.com/dave/dst/issues/45
- [`#58f25f1`](https://github.com/boynoiz/dst/commit/58f25f1ea10bac650525c8bcac219f9e69981222) *(other)* Add error check
- [`#b061a62`](https://github.com/boynoiz/dst/commit/b061a62cd01b0c5506b25c1530d5c0f2b4e5f619) *(other)* Merge pull request #51 from myshkin5/include-local-pkg
- [`#4d9bf57`](https://github.com/boynoiz/dst/commit/4d9bf5761b05cf959ffb537fcf5b47b6ba3e331c) *(other)* PR feedback
- [`#0a9c789`](https://github.com/boynoiz/dst/commit/0a9c7893d5532d6ec63a58b6f31c1586642f314e) *(other)* Add include local package config to decorator

## [0.25.5] - 2020-06-24
### Other

- [`#bce17d0`](https://github.com/boynoiz/dst/commit/bce17d0c83acd585b79d2a66855e6024edb0ba97) *(other)* Tweaked example

## [0.25.4] - 2020-06-24
### Other

- [`#027b101`](https://github.com/boynoiz/dst/commit/027b101a90e97df4810419161307a2efa0b15569) *(other)* Tweaked example

## [0.25.3] - 2020-06-24
### Other

- [`#d68e822`](https://github.com/boynoiz/dst/commit/d68e82273b4ab1d52ed576161b159f79fab7400c) *(other)* Readme
- [`#6813fa0`](https://github.com/boynoiz/dst/commit/6813fa040c1dafa6b46bfe71234d0a563218dc20) *(other)* Readme
- [`#d982238`](https://github.com/boynoiz/dst/commit/d982238666570da946c6ac0865dc13aa4140406d) *(other)* Readme

## [0.25.2] - 2020-06-24
### Other

- [`#3e8a506`](https://github.com/boynoiz/dst/commit/3e8a50627b892bd276d1e3e2a8495c4e2988ac18) *(other)* Improved readme

## [0.25.0] - 2020-06-24
### Other

- [`#edd10b4`](https://github.com/boynoiz/dst/commit/edd10b4435ae703f226502f871cf97d452e0b35a) *(other)* Added Decorations, DecorationPoint to dstutil package, example, documentation.

## [0.24.0] - 2020-05-12
### Other

- [`#ce1c8af`](https://github.com/boynoiz/dst/commit/ce1c8af3ca7fccd4405ec0594a8b68d40370dda0) *(other)* Revert verbose tests
- [`#7344d16`](https://github.com/boynoiz/dst/commit/7344d167aa095c603969cf11a256073b5f3c83c5) *(other)* Turn on verbose tests
- [`#da8b371`](https://github.com/boynoiz/dst/commit/da8b371bf516817e6d8ff777471c1e0be69c7177) *(other)* Removed courtney from travis config
- [`#abc6b91`](https://github.com/boynoiz/dst/commit/abc6b91b1dd63f55060e45bdf0e8643863d854e7) *(other)* Fix go modules test
- [`#89a63d6`](https://github.com/boynoiz/dst/commit/89a63d6fa30a93a07d06b9f15f9206ee140402c0) *(other)* Skipping two stdlib tests
- [`#4552165`](https://github.com/boynoiz/dst/commit/4552165b95c623b3119d0067dc11e55324e71dc0) *(other)* Improve HasNoPos generation
- [`#0fe0655`](https://github.com/boynoiz/dst/commit/0fe0655a34b800d91a3c6db07fdcc98a99f4b159) *(other)* Add RbraceHasNoPos to clone
- [`#a38c508`](https://github.com/boynoiz/dst/commit/a38c5083607306bde8ea516d7033a1f79f74e344) *(other)* Fix TestLoadStdLibAll
- [`#76249bd`](https://github.com/boynoiz/dst/commit/76249bdb3d61d6f93513ebba64a6217390ae34fc) *(other)* Fix module bug for go 1.14
- [`#7c5c1d3`](https://github.com/boynoiz/dst/commit/7c5c1d3526ef23b5a45767da56109a2ea1a435fe) *(other)* Update to new version of golang.org/x/tools
- [`#691fbb4`](https://github.com/boynoiz/dst/commit/691fbb4a220492be554c6a2098cac022f86d160d) *(other)* Fix BadXXX formatting bug
- [`#343b2c5`](https://github.com/boynoiz/dst/commit/343b2c502fda99fee04a5555719e34a3a87cf868) *(other)* Readme changes. Added sourcegraph badge. Changed stability to "stable".

## [0.23.1] - 2019-02-05
### Other

- [`#e3c2080`](https://github.com/boynoiz/dst/commit/e3c208037a64131003377785f5ccddc79a887a29) *(other)* Panic if restorer encounters an Ident with Path set, but without a Resolver

## [0.23.0] - 2018-12-16
### Other

- [`#d9f5ecd`](https://github.com/boynoiz/dst/commit/d9f5ecd83011abdf5288918d7bc19af9d37f8699) *(other)* Adds Hints to gobuild and gopackages resolvers

## [0.22.4] - 2018-12-06
### Other

- [`#e7be21f`](https://github.com/boynoiz/dst/commit/e7be21f386387c0b629e7cac36ed886df03fe231) *(other)* Lint fixes
- [`#58eeaf7`](https://github.com/boynoiz/dst/commit/58eeaf70faeadaaf6df22446bc18a67487849d46) *(other)* Change codecov badge

## [0.22.3] - 2018-12-06
### Other

- [`#55644ce`](https://github.com/boynoiz/dst/commit/55644ce9f22962cc50a1aafe3f11b4afc5f7d711) *(other)* Added test for Package.Save, SaveWithResolver

## [0.22.2] - 2018-12-06
### Other

- [`#144504e`](https://github.com/boynoiz/dst/commit/144504e3e50ac8846472ab42ff6933c86b452b3d) *(other)* Tests for bad nodes and package decoration

## [0.22.1] - 2018-12-06
### Other

- [`#07ba42f`](https://github.com/boynoiz/dst/commit/07ba42fa7eb21a651b26b51411d62404aac24e8f) *(other)* Improve test coverage for dstutil.Apply

## [0.22.0] - 2018-12-06
### Other

- [`#374a06a`](https://github.com/boynoiz/dst/commit/374a06a548ab573e00d7ceae0ff542bc1df62054) *(other)* Added Restorer.Extras flag and disabled resorting Scopes and Objects by default.

## [0.21.0] - 2018-12-04
### Other

- [`#9d1cc8e`](https://github.com/boynoiz/dst/commit/9d1cc8e5fd8c7f2d9688c703dd6b104f17e90d5f) *(other)* Revert "Removed Scope and Object"

## [0.20.0] - 2018-12-03
### Other

- [`#9f0c941`](https://github.com/boynoiz/dst/commit/9f0c941d1ca6615eb1d0ad096c60f700c80ef0df) *(other)* Removed Scope and Object

## [0.19.1] - 2018-12-03
### Other

- [`#5a24b9a`](https://github.com/boynoiz/dst/commit/5a24b9afae7a29acde78bf6b9de22bb26d40c185) *(other)* Removed dummy testing package

## [0.19.0] - 2018-12-02
### Other

- [`#ab0d940`](https://github.com/boynoiz/dst/commit/ab0d94079c9b476551b51bd263e4e8f2fc9af98f) *(other)* Added SaveWithResolver

## [0.18.0] - 2018-11-30
### Other

- [`#842c200`](https://github.com/boynoiz/dst/commit/842c2000680bae4b03450681484b2bbb27b4b3c6) *(other)* Renamed decorator.Package.Files to Syntax, decorate imported packages
- [`#dd8dcad`](https://github.com/boynoiz/dst/commit/dd8dcad72f31258f218530ffe9ea355b9b05c22a) *(other)* Added File.Imports to decoration / clone. Excluded from restore.
- [`#3db6114`](https://github.com/boynoiz/dst/commit/3db6114f64a574fce102a6d30d3e8cffbccd76a7) *(other)* Added livense from go repo
- [`#15bc431`](https://github.com/boynoiz/dst/commit/15bc431c86598d10c96ca25366cca5e194458d0a) *(other)* Readme
- [`#805c3b4`](https://github.com/boynoiz/dst/commit/805c3b4ca333671f2b399362bdd88049792d888d) *(other)* Testing
- [`#17d1ba3`](https://github.com/boynoiz/dst/commit/17d1ba3b6a2db44a88d321f20e8634217f62d8d2) *(other)* Test coverage badge
- [`#2735e40`](https://github.com/boynoiz/dst/commit/2735e401b1790a308aacc018003f8028f0c12bf0) *(other)* Clone tests
- [`#dc5e1a9`](https://github.com/boynoiz/dst/commit/dc5e1a99246900928e58822fd12dda7f7fd1edbc) *(other)* Testing
- [`#093887d`](https://github.com/boynoiz/dst/commit/093887d60c6792981b57c410b8cc345c7c9c4577) *(other)* Comments
- [`#a8dc998`](https://github.com/boynoiz/dst/commit/a8dc998c45ff9aabf3a6ab622ce29572c33527b5) *(other)* Remove unused function

## [0.17.3] - 2018-11-29
### Other

- [`#4f68f03`](https://github.com/boynoiz/dst/commit/4f68f03a192e369710e70edc5f01d2385fedeeb2) *(other)* Comment

## [0.17.2] - 2018-11-29
### Other

- [`#298f20e`](https://github.com/boynoiz/dst/commit/298f20e3441b48d4971400db3a50e0e1395ea405) *(other)* Refactored decorateSelectorExpr to improve decoration merging

## [0.17.1] - 2018-11-29
### Other

- [`#e697f38`](https://github.com/boynoiz/dst/commit/e697f3884b92e8a9f937e3cfdac062a0cbee8040) *(other)* Fix SelectorExpr to Ident decoration merge edge case bug

## [0.17.0] - 2018-11-29
### Other

- [`#6fd11d6`](https://github.com/boynoiz/dst/commit/6fd11d6884267935e143053e0f9db9c596cee098) *(other)* Fix bug where dot-imported ident was not resolved when in the X of a SelectorExpr

## [0.16.4] - 2018-11-28
### Other

- [`#cd559e5`](https://github.com/boynoiz/dst/commit/cd559e52024a6b5e910521f6991f914f77a76a53) *(other)* Refactored tests

## [0.16.3] - 2018-11-28
### Other

- [`#3d91f12`](https://github.com/boynoiz/dst/commit/3d91f12c6ebc3f4a157fd76094afa9c679f14912) *(other)* Updated go.mod with tests

## [0.16.2] - 2018-11-28
### Other

- [`#e2019de`](https://github.com/boynoiz/dst/commit/e2019deecc5fe426503ebba2e874b7853c379886) *(other)* Removed duplicate tests
- [`#a21d760`](https://github.com/boynoiz/dst/commit/a21d7600576b9ea005353e4d433fae78077190fe) *(other)* Refactored Restorer.updateImports

## [0.16.1] - 2018-11-28
### Other

- [`#f5a2113`](https://github.com/boynoiz/dst/commit/f5a21133d5753614d6943d26e9dd6a312ea9d425) *(other)* Added go.mod
- [`#921f3dd`](https://github.com/boynoiz/dst/commit/921f3ddb47393d753ab024d95c7f3c2173532f2f) *(other)* Changed diff library

## [0.16.0] - 2018-11-28
### Other

- [`#e598522`](https://github.com/boynoiz/dst/commit/e598522a0b79b7e2b667b8daa791ed158fc1e6af) *(other)* Renamed resolver interfaces
- [`#4a3d315`](https://github.com/boynoiz/dst/commit/4a3d315b24700f42b8438dbbc19d542ea0ff5e71) *(other)* Simplify imports example
- [`#15db69d`](https://github.com/boynoiz/dst/commit/15db69d7a5423c37b75bd21a8cb1bac9a9144312) *(other)* Changes to API, import management more deliberate
- [`#5b972d1`](https://github.com/boynoiz/dst/commit/5b972d1ddf469c3eacc322bfdec661e596d50ed8) *(other)* Alias example
- [`#81ff017`](https://github.com/boynoiz/dst/commit/81ff0177a9e1a24f6e35e6e2ac577ac24a14f896) *(other)* Readme

## [0.15.0] - 2018-11-27
### Other

- [`#ccc6168`](https://github.com/boynoiz/dst/commit/ccc61684bd8850aa3503b596bc57c47759699fbf) *(other)* Readme
- [`#b05dea3`](https://github.com/boynoiz/dst/commit/b05dea3ccbc35ee07d63c84d1ac9a179d61df266) *(other)* Moved Path from interface to Decorator, enforced Path set when using Resolver
- [`#455f57d`](https://github.com/boynoiz/dst/commit/455f57dd1095e53b90074d49d7b8c73fa0478c6b) *(other)* Merge pull request #33 from dave/ident
- [`#d653da7`](https://github.com/boynoiz/dst/commit/d653da7bb60880c18b50819403c361ba7c7c44e8) *(other)* Reverted Def/Ref split but kept ident normalization
- [`#68771bc`](https://github.com/boynoiz/dst/commit/68771bcb85284ca464619c743c58b18a0ae7d7ee) *(other)* Split Ident into Def and Ref, automatic SelectorExpr -> Ref conversion on decorate, etc.
- [`#77fc5e1`](https://github.com/boynoiz/dst/commit/77fc5e1123c27b61f88d03f54e995208b68efb37) *(other)* Moved Path from Decorator to resolver implementations
- [`#7e97510`](https://github.com/boynoiz/dst/commit/7e9751063172ca6ab768c7564fad02e8ccf865c4) *(other)* Readme

## [0.14.0] - 2018-11-25
### Other

- [`#695da8a`](https://github.com/boynoiz/dst/commit/695da8a226a214b2a9cfac2202e8c484441b502f) *(other)* Restorer error handling
- [`#03df45d`](https://github.com/boynoiz/dst/commit/03df45d88da0669e579584376f4370bb36477813) *(other)* Decorator error handling
- [`#32ad1ba`](https://github.com/boynoiz/dst/commit/32ad1ba83c023560b23b89f0a4b8a1c1e6cd9b2c) *(other)* Readme, example
- [`#a38a00e`](https://github.com/boynoiz/dst/commit/a38a00e89db94531332fde7a8166f394d5ec0155) *(other)* Readme

## [0.13.0] - 2018-11-25
### Other

- [`#a354612`](https://github.com/boynoiz/dst/commit/a354612ab7e1bbcbe6b51faf0a39badf9b5e19c2) *(other)* Resolver refactoring and documentation

## [0.12.2] - 2018-11-25
### Other

- [`#e6551ce`](https://github.com/boynoiz/dst/commit/e6551cec102229b0c076a13ee498f4adac36dc8d) *(other)* Optimisations, goast.IdentResolver, prevent Ident.Path from being set when not needed

## [0.12.1] - 2018-11-24
### Other

- [`#263215f`](https://github.com/boynoiz/dst/commit/263215f935adc1b762cf6e1ad2eb64d3f8ec5c15) *(other)* Add ParseComments flag to mode in ParseFile and ParseDir

## [0.12.0] - 2018-11-24
### Other

- [`#fb22fad`](https://github.com/boynoiz/dst/commit/fb22fad24bf56531ec49e7789f68c9e432e95872) *(other)* Rename Space to Before

## [0.11.1] - 2018-11-22
### Other

- [`#e3d108f`](https://github.com/boynoiz/dst/commit/e3d108f69ff6d7305391b5514a6ba792258c1a5f) *(other)* Strip vendor before comparing paths
- [`#4d0e43c`](https://github.com/boynoiz/dst/commit/4d0e43c64d3114799a795e5f42203c16af27255f) *(other)* Readme, examples

## [0.11.0] - 2018-11-22
### Other

- [`#1e56045`](https://github.com/boynoiz/dst/commit/1e56045b8e2109e57b554d97dfd4a561fa3db834) *(other)* Merge pull request #32 from dave/imports
- [`#a0990f3`](https://github.com/boynoiz/dst/commit/a0990f3106e1d31f1a57c257244a85b7ec706cef) *(other)* Readme, examples
- [`#88f1992`](https://github.com/boynoiz/dst/commit/88f1992561ab6ff260b84e4cd0d60f17ee2bbe4a) *(other)* Tweaks to test
- [`#8afb843`](https://github.com/boynoiz/dst/commit/8afb843b181ad69ecb6ec258d136d109133cd825) *(other)* TestLoadStdLibAll tests pass
- [`#dd487bb`](https://github.com/boynoiz/dst/commit/dd487bb445912c79800d10d0cddf55b063104c99) *(other)* Reverted DecoratorRestorer and streamlined Decorator / Restorer
- [`#9c0f73e`](https://github.com/boynoiz/dst/commit/9c0f73eb39b7527eee9fba7f4ce69e6d9be80f56) *(other)* Decorator and Restorer merged into DecoratorRestorer
- [`#0c264da`](https://github.com/boynoiz/dst/commit/0c264da6782aff72b047389e4a1c415acbb5eaf9) *(other)* Finishedtest cases for import manager
- [`#412fee1`](https://github.com/boynoiz/dst/commit/412fee19c4f70703d181bd3aa794dc29965238b2) *(other)* Import manager test coverage
- [`#12648bd`](https://github.com/boynoiz/dst/commit/12648bdb3209d297770712210fc331780fd09192) *(other)* More restorer resolver tests
- [`#67ed2bd`](https://github.com/boynoiz/dst/commit/67ed2bd71f4e68f3a617fb87acb020d0abd2bee1) *(other)* Decorator and Restorer import management, PackageResolver, IdentResolver

## [0.10.0] - 2018-10-25
### Other

- [`#3ad3633`](https://github.com/boynoiz/dst/commit/3ad36334796efd30e5176501326101cef4b8fbf4) *(other)* Refactor Decorator, fragger etc.

## [0.9.3] - 2018-10-25
### Other

- [`#3a7803b`](https://github.com/boynoiz/dst/commit/3a7803b9e045b4a905a8dd03707de48c9fafb651) *(other)* Tidy up handling of BadDecl, BadStmt and BadExpr to exclude newlines

## [0.9.2] - 2018-10-24
### Other

- [`#8cb762d`](https://github.com/boynoiz/dst/commit/8cb762d85880e2dc0460648796901ded9123f49f) *(other)* Added logic to decorate / restore Ignored sections (for BadExpr, BadStmt and BadDecl)

## [0.9.1] - 2018-10-24
### Other

- [`#1dba731`](https://github.com/boynoiz/dst/commit/1dba73119f7446d59fa147ff07f6415445dde201) *(other)* Readme

## [0.9.0] - 2018-10-24
### Other

- [`#2e576bf`](https://github.com/boynoiz/dst/commit/2e576bf5fd0be558f689d39de6e207c3d241a876) *(other)* Refactored ast-dst mapping in decorator and restorer

## [0.8.1] - 2018-10-24
### Other

- [`#f14008b`](https://github.com/boynoiz/dst/commit/f14008b33fb4fb9d9346a33d364fc4ccff4bf194) *(other)* Removed name parameter from applyDecorations

## [0.8.0] - 2018-10-23
### Other

- [`#3f984e9`](https://github.com/boynoiz/dst/commit/3f984e931909309a565a91fa43621f909c8f03db) *(other)* Merge pull request #29 from dave/decoration-points
- [`#3a90ff1`](https://github.com/boynoiz/dst/commit/3a90ff18714e41dfedf8229c06416adad614e7d3) *(other)* Remove unnecessary decoration attachment points (after lists of nodes)
- [`#be242e7`](https://github.com/boynoiz/dst/commit/be242e7916547bfde411a03ec869c948586dc0e9) *(other)* Add End decoration to Fle but disable in decorator / fragger

## [0.7.1] - 2018-10-23
### Other

- [`#71b48ed`](https://github.com/boynoiz/dst/commit/71b48ed28257a42a97fd6d2599a2b005bbe85724) *(other)* Render existing FuncType decorations after adding to FuncDecl

## [0.7.0] - 2018-10-23
### Other

- [`#c8a70cb`](https://github.com/boynoiz/dst/commit/c8a70cb33d15bbc6b9d563c72da8891c24829002) *(other)* Merge pull request #26 from dave/node-reuse
- [`#173bafd`](https://github.com/boynoiz/dst/commit/173bafd3edbd65f009ee63c75abebe0938507c1d) *(other)* Added Clone function, tests and documentation
- [`#873abcc`](https://github.com/boynoiz/dst/commit/873abcc234ea62079339846c78a171e38d1a52a4) *(other)* Added TestRestorerApply with failing test case for node reuse
- [`#700d848`](https://github.com/boynoiz/dst/commit/700d84874d5ddd8ff1907ffb999b673da3523518) *(other)* Readme
- [`#010fa9e`](https://github.com/boynoiz/dst/commit/010fa9ec64b1e6ee087fa0b632d5115ff2a6757f) *(other)* Refactored gendst package and added readme
- [`#7608960`](https://github.com/boynoiz/dst/commit/76089607b3ca38ebf1f609752a3f85da1a14cabe) *(other)* Fixed comments in decorations-types-generated.go
- [`#8ffa783`](https://github.com/boynoiz/dst/commit/8ffa783390e0e8ec0d4c8160714cbd78accae521) *(other)* Readme, types example
- [`#8077fc9`](https://github.com/boynoiz/dst/commit/8077fc947c0bcec354cbcfa08955f280335e1b4c) *(other)* Fixed examples, Go report card badge
- [`#2a3bc9f`](https://github.com/boynoiz/dst/commit/2a3bc9f9535056fa712c90b45cf404692debff9c) *(other)* Readme, examples
- [`#1eb0e5a`](https://github.com/boynoiz/dst/commit/1eb0e5a0dea8e45c4d015ac8bf5deff4d058c25a) *(other)* Readme
- [`#c7820dc`](https://github.com/boynoiz/dst/commit/c7820dce0b7311cfa02b070eec7d2f5ddec30afa) *(other)* Merge pull request #23 from dave/node-decs

## [0.6.0] - 2018-10-21
### Other

- [`#5df9ce4`](https://github.com/boynoiz/dst/commit/5df9ce4273d50ee6806133abb36c038efa45bdf5) *(other)* Readme and examples
- [`#97a9a70`](https://github.com/boynoiz/dst/commit/97a9a70d2e0066ad2a7d407a2db80e8ca768b562) *(other)* Refactored interfaces and decoration API
- [`#789bc50`](https://github.com/boynoiz/dst/commit/789bc5050c75ff17466abddbba907bc67d9fa479) *(other)* Added godoc badge
- [`#3e615c5`](https://github.com/boynoiz/dst/commit/3e615c5af213ed3c6978aa6b666a5676d23d6427) *(other)* Added Go docs, made Fragger private
- [`#be9d3d9`](https://github.com/boynoiz/dst/commit/be9d3d9a0f2d2c2603d81629b4518684b964efe6) *(other)* Added test to check postests compiles correctly
- [`#7797dd6`](https://github.com/boynoiz/dst/commit/7797dd6a5b2c1aa9c0da3dc4f0a7de7267184ff8) *(other)* Readme

## [0.5.0] - 2018-10-20
### Other

- [`#50416b8`](https://github.com/boynoiz/dst/commit/50416b8bd2038a0932eec3c833bd1d1a45967484) *(other)* Readme
- [`#61a8c0d`](https://github.com/boynoiz/dst/commit/61a8c0dfadfa21feea65b5a0fed549ea8f1fa082) *(other)* Readme
- [`#dd22502`](https://github.com/boynoiz/dst/commit/dd22502ee8244619dd195dcc5b565891872782c8) *(other)* Add two-stage travis build, short test mode, badges
- [`#b4ee268`](https://github.com/boynoiz/dst/commit/b4ee26861fa67bf5d7d147437e427d018b508fca) *(other)* Add Travis CI
- [`#525f962`](https://github.com/boynoiz/dst/commit/525f96212a89f9c908d7952e0486a237692935ca) *(other)* Tidy tests and filenames
- [`#49cfe08`](https://github.com/boynoiz/dst/commit/49cfe08a8a424620d83653a6823bd07b0da59f4c) *(other)* Correctly attach comments at end of line to Comment field

## [0.4.0] - 2018-10-20
### Other

- [`#63c30ba`](https://github.com/boynoiz/dst/commit/63c30bae464606c0ac923449ae9a935f810ede68) *(other)* Merge pull request #19 from dave/hanging-indent
- [`#a330855`](https://github.com/boynoiz/dst/commit/a330855689ac3fec1b9ffdc615823e3e5af6001d) *(other)* Fix hanging-indent general case
- [`#7513279`](https://github.com/boynoiz/dst/commit/751327950da7a249d3279886d7711b0fe794342d) *(other)* Re-enabled End decoration on CommClause and CaseClause
- [`#2473a82`](https://github.com/boynoiz/dst/commit/2473a828b53d0fef8677cf9650fd3621773db3f1) *(other)* Added tests for #18
- [`#522645e`](https://github.com/boynoiz/dst/commit/522645ef933a1338c4fb46528262cdad93963356) *(other)* Advance cursor one extra space when adding a newline to ensure it doesn't go before comma in lists
- [`#c96fb48`](https://github.com/boynoiz/dst/commit/c96fb48b7c4b11e54c45d0a74583637a727e97d4) *(other)* Correctly reconstruct newlines inside multi-line back-quoted string literals

## [0.3.0] - 2018-10-15
### Other

- [`#e7ca46b`](https://github.com/boynoiz/dst/commit/e7ca46bc3a04dccb22bbeed94384a87971c27e0d) *(other)* Merge pull request #16 from dave/case-clause-2
- [`#0ea31d1`](https://github.com/boynoiz/dst/commit/0ea31d1ee27d9751465342cff3737df6f110947f) *(other)* Apply correct indents in output stream
- [`#cd8e5a8`](https://github.com/boynoiz/dst/commit/cd8e5a8ae7f52fc2d792075a7742bc66e430f563) *(other)* Accach comments to correct case clauses...
- [`#e9b0633`](https://github.com/boynoiz/dst/commit/e9b0633cc175e82da69df60cc4f6052badecce2d) *(other)* Record line offset of all comments that come directly after new or empty lines
- [`#35f88a8`](https://github.com/boynoiz/dst/commit/35f88a8d025a36355c8ab97f009b0c0a3ca2f55f) *(other)* Record line offset of all node "Start" positions that come directly after new or empty lines
- [`#82f0086`](https://github.com/boynoiz/dst/commit/82f0086194e3ae60cb42a0439f175ac3e6de9ad0) *(other)* Types example
- [`#04aa755`](https://github.com/boynoiz/dst/commit/04aa755a74f0663f17ba364395d4a69535ee5b60) *(other)* Readme

## [0.2.1] - 2018-10-11
### Other

- [`#a76e9fb`](https://github.com/boynoiz/dst/commit/a76e9fbfd9747a3ddc6e13c8cd81eab685e6334b) *(other)* Space example
- [`#22b4086`](https://github.com/boynoiz/dst/commit/22b4086d057c9f5d8f12295d35c6f040fee1e8d2) *(other)* Examples / readme
- [`#23bc729`](https://github.com/boynoiz/dst/commit/23bc729f395d8dc98ba8f8d265c63041bc051e7c) *(other)* Added back in before and after spacing to nodes. Before = "Space", after = "After"
- [`#08ab36f`](https://github.com/boynoiz/dst/commit/08ab36f3edd3dcfb78cbecb639396b306f1ec9ae) *(other)* Add End to CaseClause, CommClause. Simplify interfaces, add Decorated.
- [`#bc6acb6`](https://github.com/boynoiz/dst/commit/bc6acb63fd7e6b30587478c5116a91ea9d8c8e0f) *(other)* Renamed generated-decs.go, readme
- [`#3ba1af6`](https://github.com/boynoiz/dst/commit/3ba1af60b2fd0692b8666876d7eb0f5cbf1da8fd) *(other)* Readme and examples
- [`#ead4bbf`](https://github.com/boynoiz/dst/commit/ead4bbf10e5550327a7ea0b1e5a916bfab01af3c) *(other)* Readme
- [`#7e875ce`](https://github.com/boynoiz/dst/commit/7e875ce036574c67e0d0af5935e5c8dafd7a1f58) *(other)* Merge pull request #15 from dave/interfaces
- [`#8902b6c`](https://github.com/boynoiz/dst/commit/8902b6ce1cc776f84b30f5460f3892b3e02a46ef) *(other)* Added Starter, Ender and Spacer interfaces
- [`#ca07362`](https://github.com/boynoiz/dst/commit/ca073626e4651f59d0c0ff3a174b3dbf9fae1405) *(other)* Renamed all decoration points to lose the "After" prefix

## [0.1.0] - 2018-10-10
### Other

- [`#cb0e30f`](https://github.com/boynoiz/dst/commit/cb0e30fa2661c824d9f9fce0d1ba2c80178712c6) *(other)* Readme / examples
- [`#eca8154`](https://github.com/boynoiz/dst/commit/eca81544a15d7e8e8574119ffba075c16b685080) *(other)* Refactored to remove Before space and rename After to Space
- [`#eafdd42`](https://github.com/boynoiz/dst/commit/eafdd42941e78604564971861f96a88a35676357) *(other)* Begin and After SpaceType, working and all tests pass
- [`#302be03`](https://github.com/boynoiz/dst/commit/302be03f322ce40fa872bf7a7eeeeb1d66e6c24b) *(other)* Added code to Decorator to associate space types with nodes
- [`#d15fba0`](https://github.com/boynoiz/dst/commit/d15fba061f5b209edb0c8718e653fe58af3fbcac) *(other)* Added Before, After SpaceType to decoration types and test helpers
- [`#1348546`](https://github.com/boynoiz/dst/commit/1348546526dfb7cf5d434f58cccee6397fb23b0f) *(other)* Corrected comment attachment bug, Fragger marks empty NewlineFragments
- [`#f262385`](https://github.com/boynoiz/dst/commit/f26238566367ff16e5546780ec6a1963300b7244) *(other)* Added Info to hold filenames
- [`#7e5777d`](https://github.com/boynoiz/dst/commit/7e5777d6036ba0a552921428803327fdf4fc0be9) *(other)* Fixed Object Decl / Data Node bug
- [`#700d92a`](https://github.com/boynoiz/dst/commit/700d92ad4a557b373b92deed7b9ad6abf453a7cb) *(other)* Added decorating and restoring Objects and Scopes
- [`#78c8d6e`](https://github.com/boynoiz/dst/commit/78c8d6e8a116bc18838dc7670da3e163558c2f4a) *(other)* Refactor fragger, convenience functions, package compatibility
- [`#fe8b797`](https://github.com/boynoiz/dst/commit/fe8b797ef3c4611fd70bec0f089ef3138d01645c) *(other)* Added DecorateWithNodes
- [`#77c93bb`](https://github.com/boynoiz/dst/commit/77c93bb277100bea1ffc4d44002b95ad4971135f) *(other)* Simplified examples

## [0.0.1] - 2018-10-03
### Other

- [`#cb9d7aa`](https://github.com/boynoiz/dst/commit/cb9d7aaffaaa0e0ec55c33f3c09a073f428961fa) *(other)* Added convenience functions, made decorator and restorer private, examples, readme
- [`#22e924c`](https://github.com/boynoiz/dst/commit/22e924ce2720fe4c97506939f8ea8df55d74a610) *(other)* Bug fixes, example, readme
- [`#d1a3eda`](https://github.com/boynoiz/dst/commit/d1a3eda2b2d14623b5ef63949c0d063eeb42dd29) *(other)* Add "non indented comment in case" test case
- [`#a9848e2`](https://github.com/boynoiz/dst/commit/a9848e24b4083289095c399f124e42a1139efb28) *(other)* Added dstutil package with Apply function
- [`#29c6fe2`](https://github.com/boynoiz/dst/commit/29c6fe24dc28ae09542c398459ec217ca0993599) *(other)* Added Restore and Decorate convenience functions
- [`#b42a467`](https://github.com/boynoiz/dst/commit/b42a46750287addc7add26c1fff47d89f7fdfa26) *(other)* Removed Get methods
- [`#2891b15`](https://github.com/boynoiz/dst/commit/2891b15620690db9bffa7e7996d7a2dfff7dc138) *(other)* Merge pull request #11 from dave/decorations
- [`#3392cde`](https://github.com/boynoiz/dst/commit/3392cde9769fc21a76886b6a07d634c5027396d9) *(other)* Typed accessors for decoration positions, tests pass
- [`#b0b0088`](https://github.com/boynoiz/dst/commit/b0b0088f90b1fd53a56e04a4bf045037e861acea) *(other)* Updated readme
- [`#3fde0ab`](https://github.com/boynoiz/dst/commit/3fde0aba4c1b3e0e99097f2de7c017acbc941d76) *(other)* Merge pull request #10 from dave/remove-doc-comment
- [`#641707f`](https://github.com/boynoiz/dst/commit/641707f3334f8a2e9d9bc71c1001e40dad58d3da) *(other)* Restorer tests pass
- [`#c3a9ed3`](https://github.com/boynoiz/dst/commit/c3a9ed32f396ecb3874af08ab991d5ad724fc278) *(other)* Added TestPositions
- [`#910150a`](https://github.com/boynoiz/dst/commit/910150a7d74670c26104eb59589b5858df28684e) *(other)* Decorator tests pass
- [`#1aa82b9`](https://github.com/boynoiz/dst/commit/1aa82b9f8d512544e681202ddbe0d860ce7f6dab) *(other)* Added fragment package with Info data
- [`#16d26bf`](https://github.com/boynoiz/dst/commit/16d26bf424a008c43870f5142662b40306e0a2e0) *(other)* Remove customLength function
- [`#aa185da`](https://github.com/boynoiz/dst/commit/aa185da06848f5ace802ef91c1f3f6aadca16510) *(other)* Rename NodeToDst to DecorateNode
- [`#49f519b`](https://github.com/boynoiz/dst/commit/49f519bde2e593b3348890272c4aa45b819736aa) *(other)* Removed Comment, CommentGroup, *.Doc, *.Comment, File.Comments. Added DecorationStmt, DecorationDecl.
- [`#377db68`](https://github.com/boynoiz/dst/commit/377db68dada3dd8867043a2f72878ec16e550bd3) *(other)* Postition tests and added to TestStdLib
- [`#f83e579`](https://github.com/boynoiz/dst/commit/f83e5795d120fc8a9d4725ac9f1d4d85901ba0e2) *(other)* RangeStmt length tweaks
- [`#4695ec2`](https://github.com/boynoiz/dst/commit/4695ec210a33c7679495788615056818b1f37218) *(other)* Fix ChanType length
- [`#d7a5234`](https://github.com/boynoiz/dst/commit/d7a52340caa555b23a6acba2f2fa520be8885c23) *(other)* Fixed comment-at-end-of-file bug
- [`#939a8ad`](https://github.com/boynoiz/dst/commit/939a8adb7c7971d9fb085e9653ddc6a29586946b) *(other)* Fixed last comment in block bug
- [`#5148b60`](https://github.com/boynoiz/dst/commit/5148b60fb63a36c3ef5c331ec45c10c13e36f454) *(other)* Bug fixes, add StdLib test
- [`#520e83e`](https://github.com/boynoiz/dst/commit/520e83ea149cd239aa99e3c6e6694b0f493a2ac2) *(other)* Run code gen
- [`#4e6a38a`](https://github.com/boynoiz/dst/commit/4e6a38a06e09aa719313447241c56bd19b23bdde) *(other)* Add special case for CaseClause.Case
- [`#2fd3790`](https://github.com/boynoiz/dst/commit/2fd37906564a862889f96f439a6b96719c6fd2bf) *(other)* Add bools for FieldList.Opening and FieldList.Closing
- [`#c73d4d8`](https://github.com/boynoiz/dst/commit/c73d4d8ebe54a810bc3d570f7f004914c7ad6c37) *(other)* Updated readme
- [`#049ff22`](https://github.com/boynoiz/dst/commit/049ff224213c6619b86c8166efff6b01fcd171b8) *(other)* Big refactor. All tests pass.
- [`#45f9e60`](https://github.com/boynoiz/dst/commit/45f9e609a097bd6207957fe5643e1f29b985858a) *(other)* Refactor gendst
- [`#85effa2`](https://github.com/boynoiz/dst/commit/85effa2387deac852effb0f36e6bd99c43d213e8) *(other)* Remove special case todos (now in an issue)
- [`#4bfd70b`](https://github.com/boynoiz/dst/commit/4bfd70bc46eedca6c7c1da0dfbc64c5cd1370ae0) *(other)* Add decorator test for FuncDecl
- [`#c5e80d2`](https://github.com/boynoiz/dst/commit/c5e80d2d0468c7af3ecb762da94f72c8ac682bf8) *(other)* Added readme
- [`#e8a55f2`](https://github.com/boynoiz/dst/commit/e8a55f21602cc238360bfe26cbe02da7292a8ec9) *(other)* Prototype
- [`#e7ec5bb`](https://github.com/boynoiz/dst/commit/e7ec5bb967e6fdd6ad5cbb0633227e5aa8621422) *(other)* Rename package to dst
- [`#75390d2`](https://github.com/boynoiz/dst/commit/75390d20b8449b085a00b342ef6a7025e7b43f26) *(other)* Add Go 1.11 go/ast package
- [`#714401c`](https://github.com/boynoiz/dst/commit/714401c3bdfefa58adffb093ad44dcf63b48619d) *(other)* First commit

[unreleased]: https://github.com/boynoiz/dst/compare/v0.27.3..HEAD
[0.27.3]: https://github.com/boynoiz/dst/compare/v0.27.2..v0.27.3
[0.27.2]: https://github.com/boynoiz/dst/compare/v0.27.1..v0.27.2
[0.27.1]: https://github.com/boynoiz/dst/compare/v0.27.0..v0.27.1
[0.27.0]: https://github.com/boynoiz/dst/compare/v0.26.2..v0.27.0
[0.26.2]: https://github.com/boynoiz/dst/compare/v0.26.1..v0.26.2
[0.26.1]: https://github.com/boynoiz/dst/compare/v0.26.0..v0.26.1
[0.26.0]: https://github.com/boynoiz/dst/compare/v0.25.5..v0.26.0
[0.25.5]: https://github.com/boynoiz/dst/compare/v0.25.4..v0.25.5
[0.25.4]: https://github.com/boynoiz/dst/compare/v0.25.3..v0.25.4
[0.25.3]: https://github.com/boynoiz/dst/compare/v0.25.2..v0.25.3
[0.25.2]: https://github.com/boynoiz/dst/compare/v0.25.0..v0.25.2
[0.25.0]: https://github.com/boynoiz/dst/compare/v0.24.0..v0.25.0
[0.24.0]: https://github.com/boynoiz/dst/compare/v0.23.1..v0.24.0
[0.23.1]: https://github.com/boynoiz/dst/compare/v0.23.0..v0.23.1
[0.23.0]: https://github.com/boynoiz/dst/compare/v0.22.4..v0.23.0
[0.22.4]: https://github.com/boynoiz/dst/compare/v0.22.3..v0.22.4
[0.22.3]: https://github.com/boynoiz/dst/compare/v0.22.2..v0.22.3
[0.22.2]: https://github.com/boynoiz/dst/compare/v0.22.1..v0.22.2
[0.22.1]: https://github.com/boynoiz/dst/compare/v0.22.0..v0.22.1
[0.22.0]: https://github.com/boynoiz/dst/compare/v0.21.0..v0.22.0
[0.21.0]: https://github.com/boynoiz/dst/compare/v0.20.0..v0.21.0
[0.20.0]: https://github.com/boynoiz/dst/compare/v0.19.1..v0.20.0
[0.19.1]: https://github.com/boynoiz/dst/compare/v0.19.0..v0.19.1
[0.19.0]: https://github.com/boynoiz/dst/compare/v0.18.0..v0.19.0
[0.18.0]: https://github.com/boynoiz/dst/compare/v0.17.3..v0.18.0
[0.17.3]: https://github.com/boynoiz/dst/compare/v0.17.2..v0.17.3
[0.17.2]: https://github.com/boynoiz/dst/compare/v0.17.1..v0.17.2
[0.17.1]: https://github.com/boynoiz/dst/compare/v0.17.0..v0.17.1
[0.17.0]: https://github.com/boynoiz/dst/compare/v0.16.4..v0.17.0
[0.16.4]: https://github.com/boynoiz/dst/compare/v0.16.3..v0.16.4
[0.16.3]: https://github.com/boynoiz/dst/compare/v0.16.2..v0.16.3
[0.16.2]: https://github.com/boynoiz/dst/compare/v0.16.1..v0.16.2
[0.16.1]: https://github.com/boynoiz/dst/compare/v0.16.0..v0.16.1
[0.16.0]: https://github.com/boynoiz/dst/compare/v0.15.0..v0.16.0
[0.15.0]: https://github.com/boynoiz/dst/compare/v0.14.0..v0.15.0
[0.14.0]: https://github.com/boynoiz/dst/compare/v0.13.0..v0.14.0
[0.13.0]: https://github.com/boynoiz/dst/compare/v0.12.2..v0.13.0
[0.12.2]: https://github.com/boynoiz/dst/compare/v0.12.1..v0.12.2
[0.12.1]: https://github.com/boynoiz/dst/compare/v0.12.0..v0.12.1
[0.12.0]: https://github.com/boynoiz/dst/compare/v0.11.1..v0.12.0
[0.11.1]: https://github.com/boynoiz/dst/compare/v0.11.0..v0.11.1
[0.11.0]: https://github.com/boynoiz/dst/compare/v0.10.0..v0.11.0
[0.10.0]: https://github.com/boynoiz/dst/compare/v0.9.3..v0.10.0
[0.9.3]: https://github.com/boynoiz/dst/compare/v0.9.2..v0.9.3
[0.9.2]: https://github.com/boynoiz/dst/compare/v0.9.1..v0.9.2
[0.9.1]: https://github.com/boynoiz/dst/compare/v0.9.0..v0.9.1
[0.9.0]: https://github.com/boynoiz/dst/compare/v0.8.1..v0.9.0
[0.8.1]: https://github.com/boynoiz/dst/compare/v0.8.0..v0.8.1
[0.8.0]: https://github.com/boynoiz/dst/compare/v0.7.1..v0.8.0
[0.7.1]: https://github.com/boynoiz/dst/compare/v0.7.0..v0.7.1
[0.7.0]: https://github.com/boynoiz/dst/compare/v0.6.0..v0.7.0
[0.6.0]: https://github.com/boynoiz/dst/compare/v0.5.0..v0.6.0
[0.5.0]: https://github.com/boynoiz/dst/compare/v0.4.0..v0.5.0
[0.4.0]: https://github.com/boynoiz/dst/compare/v0.3.0..v0.4.0
[0.3.0]: https://github.com/boynoiz/dst/compare/v0.2.1..v0.3.0
[0.2.1]: https://github.com/boynoiz/dst/compare/v0.1.0..v0.2.1
[0.1.0]: https://github.com/boynoiz/dst/compare/v0.0.1..v0.1.0

<!-- generated by git-cliff -->
