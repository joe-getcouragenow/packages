


current_dir = $(shell pwd)

FLU_LIB_NAME = ???
FLU_LIB_FSPATH = $(PWD)/$(FLU_LIB_NAME)

FLU_SAMPLE_NAME = ???
FLU_SSAMPLE_FSPATH = $(PWD)/$(FLU_SAMPLE_NAME)


## Prints the flutter settings
flu-print: ## print
	@echo -- FLU -- 
	@echo FLU_LIB_NAME: $(FLU_LIB_NAME)
	@echo FLU_LIB_FSPATH: $(FLU_LIB_FSPATH)
	@echo FLU_SAMPLE_NAME: $(FLU_SAMPLE_NAME)
	@echo FLU_SSAMPLE_FSPATH: $(FLU_SSAMPLE_FSPATH)
	@echo

	@echo -- I18N --
	@echo FLU_LANG_TEMPLATE_PATH: $(FLU_LANG_TEMPLATE_PATH)
	@echo FLU_LANG_FILE_PREFIX_OUT: $(FLU_LANG_FILE_PREFIX_OUT)
	@echo



### FLU

## Check for outdated packages.
flu-pub-outdated:
	cd $(FLU_SSAMPLE_FSPATH) && pub outdated

## Upgrade for all packages.
flu-pub-upgrade:
	cd $(FLU_SSAMPLE_FSPATH) && pub upgrade

## Recreates the Flutter scaffolding code. Useful after a SDK Channel change has occured.
flu-recreate:
	cd $(FLU_SSAMPLE_FSPATH) && flutter create .

## Configure flutter to correct CHANNEL.
flu-config: ## flu-config
	flutter channel dev
	flutter upgrade --force


## Runs Flutter Web.
flu-web-run: ## flu-web-run
	flutter config --enable-web
	cd $(FLU_SSAMPLE_FSPATH) && flutter run -d chrome

## Builds flutter web as a release version
flu-web-build: ## flu-web-build
	flutter config --enable-web
	cd $(FLU_SSAMPLE_FSPATH) && flutter build web

## Fixes a common bug with IOS
flu-mob-fix:
	# webrtc fix
	# MIGHT have to apply this to CI also. Lets see.
	cd $(FLU_SSAMPLE_FSPATH) && rm ios/Podfile.lock
	cd $(FLU_SSAMPLE_FSPATH) && rm -rf ios/Pods
	cd $(FLU_SSAMPLE_FSPATH) && flutter clean
	 
	# sign for ios local

	# 1. open https://developer.apple.com/account/resources/certificates/list
	# - go to Identifiers and make an BUNDLE ID ("org.getcouragenow.securesuite")
	# - open xcode project 
	# and GENERAL --> put in the BUNDLE ID (org.getcouragenow.securesuite)
	# and Signing & Capabilities -- Choose Team and put in the BUNDLE ID also

	# sign for CI
	#	* IOS_APP_BUNDLER_ID refer the iOS application bundle identifier, which can be found in the Info.plist.
		# Was done above...
	#	* IOS_P12_SIGN_KEY_BASE_64 refer to the p12 certificate encoded as base 64 which will be used to sign the build and archive.
	#   * IOS_P12_SIGN_KEY_PASSWORD refer to the p12 certificate password.
	#   * IOS_SIGNING_TYPE refer to the type of signing key that should be used, IOS_APP_DEVELOPMENT or IOS_APP_STORE.
	#   * IOS_SIGN_KEY_ISSUER_ID refer to the id of the certificate issuer.
	#   * IOS_SIGN_KEY_ID refer to id of private key of the certificate.
	#   * APP_STORE_CONNECT_PRIVATE_KEY refer to private key used to access app store connect api.


flu-mob-run: ## flu-mob-run
	cd $(FLU_SSAMPLE_FSPATH) && flutter run -d all

flu-mob-build: ## flu-mob-build
	# ios
	# https://flutter.dev/docs/deployment/ios
	# change to ios 11 in xcode for runner and pods
	cd $(FLU_SSAMPLE_FSPATH)/ios && pod install
	#cd $(FLU_SSAMPLE_FSPATH) && flutter build ios

	cd $(FLU_SSAMPLE_FSPATH) && 	flutter build ios --release --no-codesign

## Inits Flutter Desktop
flu-desk-init: ## flu-desk-init
	cd $(FLU_SSAMPLE_FSPATH) && hover init
	

## Runs Flutter Desktop
flu-desk-run: ## flu-desk-run
	cd $(FLU_SSAMPLE_FSPATH) && hover run


### GEN 

## Generates all code
flu-gen: ## flu-gen
	cd $(FLU_LIB_FSPATH) && flutter packages get
	$(MAKE) gen-icons
	$(MAKE) gen-hive
	$(MAKE) gen-proto
	cd $(FLU_LIB_FSPATH) && flutter analyze 

gen-icons:
	# mobile and web
	@echo
	@echo Generating icons for Flutter
	@echo
	cd $(FLU_LIB_FSPATH) && flutter pub run flutter_launcher_icons:main

	# desktop
	@echo
	@echo Copying icon-png from flutter assets into go assets, so hover can use it
	@echo
	cp $(FLU_LIB_FSPATH)/assets/icon/icon.png $(FLU_LIB_FSPATH)/go/assets

gen-hive:
	cd $(FLU_LIB_FSPATH) && flutter packages pub run build_runner build --delete-conflicting-outputs

gen-proto:
	pub global activate protoc_plugin
	mkdir -p $(FLU_LIB_FSPATH)/lib/api/v1/google/protobuf

ifeq ($(GO_OS), windows)
	@echo Windows detected
	protoc empty.proto timestamp.proto wrappers.proto --proto_path=$(LIB_FSPATH)/server/third_party/google/protobuf/ --plugin=$(HOME_PATH)/AppData/Roaming/Pub/Cache/bin/protoc-gen-dart.bat --dart_out=grpc:"$(PROTO_OUTPUT)/client/lib/chat_view/api/v1/google/protobuf"
	protoc chat.proto --proto_path=$(LIB_FSPATH)/server/api/proto/v1/ --plugin=$(HOME_PATH)/AppData/Roaming/Pub/Cache/bin/protoc-gen-dart.bat --dart_out=grpc:"$(PROTO_OUTPUT)/client/lib/chat_view/api/v1/"
else
	protoc empty.proto timestamp.proto wrappers.proto --proto_path=$(LIB_FSPATH)/server/third_party/google/protobuf --plugin=protoc-gen-dart=$(HOME)/.pub-cache/bin/protoc-gen-dart --dart_out=grpc:$(FLU_SSAMPLE_FSPATH)/lib/api/v1/google/protobuf
	protoc chat.proto --proto_path=$(LIB_FSPATH)/server/api/proto/v1/ --plugin=protoc-gen-dart=$(HOME)/.pub-cache/bin/protoc-gen-dart --dart_out=grpc:$(FLU_SSAMPLE_FSPATH)/client/lib/chat_view/api/v1/
endif


## LANG
# 0. make lang-dep => ensure you have the i18n tool
# 1. make lang-gen-flu => Generates everything
# OR
# 1. make lang-gen-flu-all = > Generates for maintemplate and all submodules

## LANG
# 0. 
# 1. make lang-dep => ensure you have the i18n tool
# 2. make lang-gen-flu => generate all files

FLU_LANG_LOCALES = en,fr,es,de,it,ur

FLU_LANG_DIR = $(FLU_LIB_FSPATH)/i18n
FLU_LANG_LOCALIZATION_DIR = $(FLU_LIB_FSPATH)/lib/core/i18n
FLU_LANG_GENERATED_DIR = $(FLU_LANG_LOCALIZATION_DIR)/generated
FLU_LANG_TEMPLATE_NAME = intl_messages.arb
FLU_LANG_TEMPLATE_PATH = $(FLU_LANG_DIR)/$(FLU_LANG_TEMPLATE_NAME)
FLU_LANG_FILE_PREFIX_OUT = lang

## prints flu-lang variables
flu-gen-lang-print: ## flu-gen-lang-print
	@echo
	@echo FLU_LANG_LOCALES: 			$(FLU_LANG_LOCALES)
	@echo FLU_LANG_DIR: 					$(FLU_LANG_DIR)
	@echo FLU_LANG_LOCALIZATION_DIR: 		$(FLU_LANG_LOCALIZATION_DIR)
	@echo FLU_LANG_GENERATED_DIR: 			$(FLU_LANG_GENERATED_DIR)
	@echo FLU_LANG_TEMPLATE_NAME: 			$(FLU_LANG_TEMPLATE_NAME)
	@echo FLU_LANG_FILE_PREFIX_OUT: 		$(FLU_LANG_FILE_PREFIX_OUT)
	@echo	
#	
flu-gen-lang-dep: ## flu-gen-lang-dep
	go get -u github.com/getcouragenow/bootstrap/tool/i18n

#
## Generates language file for maintemplate and all submodules
#flu-gen-lang-all: ## flu-gen-lang-all
#	$(MAKE) flu-gen-lang
#	$(MAKE) flu-gen-lang-dart
#	$(MAKE) flu-gen-lang-submodules
#
### Generates lang for sub modules
#flu-gen-lang-submodules: ## flu-gen-lang-submodules
#	#TODO joe make recursive
#	cd ../mod-account && make flu-gen-lang
#	cd ../mod-chat && make flu-gen-lang
#	cd ../mod-main && make flu-gen-lang
#	cd ../mod-geo && make flu-gen-lang
#	cd ../mod-chat && make flu-gen-lang
#	cd ../mod-write && make flu-gen-lang
#	cd ../mod-ion && make flu-gen-lang

## generates language code
flu-gen-lang: ## flu-gen-lang
	@echo -- Running flutter pub get first
	cd $(FLU_LIB_FSPATH) && flutter pub get

	@echo -- Creating: $(FLU_LANG_GENERATED_DIR)
	mkdir -p $(FLU_LANG_GENERATED_DIR)
	
	@echo -- Creating: $(FLU_LANG_DIR)
	mkdir -p $(FLU_LANG_DIR)

	@echo -- Extracting: To ARB
	cd $(FLU_LIB_FSPATH) && flutter pub run intl_translation:extract_to_arb --output-dir=$(FLU_LANG_DIR) $(FLU_LANG_LOCALIZATION_DIR)/translations.dart
	
	@echo -- Translating Text
	i18n flutter --dir $(FLU_LANG_DIR) --template $(FLU_LANG_TEMPLATE_PATH) --prefix $(FLU_LANG_FILE_PREFIX_OUT) --languages $(FLU_LANG_LOCALES) -f
	i18n flutter --dir $(FLU_LANG_DIR)

## generates dart code out of arb files
flu-gen-lang-dart: ## flu-gen-lang-dart
	cd $(FLU_LIB_FSPATH) && flutter pub run intl_translation:generate_from_arb --output-dir=$(FLU_LANG_GENERATED_DIR) $(FLU_LANG_LOCALIZATION_DIR)/translations.dart $(FLU_LANG_DIR)/*.arb




### ASTI

# 1. dep
# 2. build
# 3. run (for the desktop your on)
# 4. pack (for the desktop your on)
# 5. sign (for the desktop your on)

go-desk-dep: ## go-desk-dep
	# as per the readme: https://github.com/asticode/go-astilectron-bundler#installation  !!
	go get -u github.com/asticode/go-astilectron-bundler/...
	go get -u github.com/asticode/go-astilectron-bootstrap
	cd $(current_dir)/desktop && go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
	which astilectron-bundler

go-desk-build: ## go-desk-build
	# build flutter web
	# flutter config --enable-web
	# cd $(FLU_SSAMPLE_FSPATH) && flutter build web

	# copy flutter web build to desktop/resources/app
	rm -rf $(FLU_LIB_FSPATH)desktop/resources/app
	mkdir -p $(FLU_LIB_FSPATH)desktop/resources/app/
	cp -r $(FLU_LIB_FSPATH)build/web/* $(FLU_LIB_FSPATH)desktop/resources/app

	# -d is for darwin
	cd $(FLU_LIB_FSPATH)desktop && astilectron-bundler

############## Linux ##############
go-desk-run-lin: ## go-desk-run-lin
	cd $(FLU_LIB_FSPATH) && desktop/output/linux-amd64/ION\ Desktop\ App

go-desk-pack-lin: ## go-desk-pack-lin
	# deb
	cd $(FLU_LIB_FSPATH)packer/ && make go-pack-deb

go-desk-sign-lin:
	# https://blog.packagecloud.io/eng/2014/10/28/howto-gpg-sign-verify-deb-packages-apt-repositories/
	
	# GPG_KEY_ID is the ID of gpg key
	cd packer/build/linux-deb && dpkg-sig -k $(GPG_KEY_ID) --sign builder *.deb
	cd packer/build/linux-deb && dpkg-sig --verify *.deb

############## Windows ##############
go-desk-run-win: ## go-desk-run-win
	# SOMEONE on Windows fix this to be correct pathing style
	$(current_dir)/desktop/output/windows-386/ION\ Desktop\ App.exe

go-desk-pack-win:
	cd $(FLU_LIB_FSPATH)packer/ && make go-pack-win

go-desk-sign-win:
	# https://github.com/itchio/itch-setup/blob/master/scripts/ci-build.sh#L73
	#TODO

############## Windows ##############
go-desk-run-mac: ## go-desk-run-mac
	open $(current_dir)/desktop/output/darwin-amd64/ION\ Desktop\ App.app 

go-desk-pack-mac:  ## go-desk-pack-mac
	#darwin
	#cd $(FLU_SSAMPLE_FSPATH) && hover init-packaging darwin-bundle

	# skaffold the file system with the pkg templates.
	#cd $(FLU_SSAMPLE_FSPATH) && hover init-packaging darwin-pkg

	# inject out .app into it.
	cp -rf $(current_dir)/desktop/output/darwin-amd64/ION\ Desktop\ App.app ./Applications/maintemplate.app


TARGET_MAC=$(current_dir)/desktop/output/darwin-amd64/ION\ Desktop\ App.app 
SIGNKEY_MAC="Developer ID Application: Amos Wenger (B2N6FSRTPV)"
# From: https://developer.apple.com/account
# NOTE: waiting on Apple to approve me...
go-desk-sign-mac: ## go-desk-sign-mac
	# sign *after* packing
	echo TARGET_MAC : $(TARGET_MAC)
	echo SIGNKEY_MAC : $(SIGNKEY_MAC)
	codesign --deep --force --verbose --sign ${SIGNKEY_MAC} "${TARGET}"
	codesign --verify -vvvv "${TARGET}"

go-desk-clean:
	rm -rf $(FLU_LIB_FSPATH)desktop/resources/app
	rm -rf $(FLU_LIB_FSPATH)desktop/output
	rm -rf $(FLU_LIB_FSPATH)desktop/bind_darwin_amd64.go
	rm -rf $(FLU_LIB_FSPATH)desktop/bind_linux_amd64.go
	rm -rf $(FLU_LIB_FSPATH)desktop/bind_windows_386.go
	rm -rf $(FLU_LIB_FSPATH)desktop/windows.syso


