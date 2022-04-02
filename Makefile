# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


LOCAL_REGISTRY="leonharetd"
LOCAL_CONTROLLER_IMAGE="sample-k8s-device-plugin"

.PHONY: clean
clean:
	rm rm ./build

.PHONY: sample
sample:
	docker build -t $(LOCAL_REGISTRY)/$(LOCAL_CONTROLLER_IMAGE):dev -f ./Dockerfile .
	docker push $(LOCAL_REGISTRY)/$(LOCAL_CONTROLLER_IMAGE):dev 
