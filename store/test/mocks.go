/**
 * Copyright 2020 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/xmidt-org/argus/model"
	"github.com/xmidt-org/argus/store"
)

type MockDB struct {
	mock.Mock
}

func (s *MockDB) Push(key model.Key, item store.OwnableItem) error {
	args := s.Called(key, item)
	return args.Error(0)
}

func (s *MockDB) Get(key model.Key) (store.OwnableItem, error) {
	args := s.Called(key)
	return args.Get(0).(store.OwnableItem), args.Error(1)
}

func (s *MockDB) Delete(key model.Key) (store.OwnableItem, error) {
	args := s.Called(key)
	return args.Get(0).(store.OwnableItem), args.Error(1)
}

func (s *MockDB) GetAll(bucket string) (map[string]store.OwnableItem, error) {
	args := s.Called(bucket)
	return args.Get(0).(map[string]store.OwnableItem), args.Error(1)
}

func (s *MockDB) Close() {
	s.Called()
}

func (s *MockDB) Ping() error {
	args := s.Called()
	return args.Error(0)
}
