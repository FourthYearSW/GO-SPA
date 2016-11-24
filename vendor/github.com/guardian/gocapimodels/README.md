#gocapimodels
Generated code for Content API models for Go.

The current release is from:

- [content-api-models v9.18](https://github.com/guardian/content-api-models/releases/tag/v9.18)
- [story-packages-model v1.0.3](https://github.com/guardian/story-packages-model/releases/tag/v1.0.3)
- [content-atom v2.4.3](https://github.com/guardian/content-atom/releases/tag/v2.4.3)

##Details

This package was created as follows:

`thrift -r --gen go:package_prefix=github.com/guardian/gocapimodels/ content.thrift` 

The source `Thrift` can be found in `/thrift` in this repo.

##Warning

Some fields are missing from `Content` as they used reserved words - specifically "end".
